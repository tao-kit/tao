package conf

import (
	"fmt"
	"github.com/sllt/tao/core/jsonx"
	"github.com/sllt/tao/core/mapping"
	"github.com/sllt/tao/internal/encoding"
	"log"
	"os"
	"path"
	"reflect"
	"strings"
)

const distanceBetweenUpperAndLower = 32

var loaders = map[string]func([]byte, interface{}) error{
	".json": LoadFromJsonBytes,
	".toml": LoadFromTomlBytes,
	".yaml": LoadFromYamlBytes,
	".yml":  LoadFromYamlBytes,
}

type fieldInfo struct {
	name     string
	kind     reflect.Kind
	children map[string]fieldInfo
}

// Load loads config into v from file, .json, .yaml and .yml are acceptable.
func Load(file string, v interface{}, opts ...Option) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	loader, ok := loaders[strings.ToLower(path.Ext(file))]
	if !ok {
		return fmt.Errorf("unrecognized file type: %s", file)
	}

	var opt options
	for _, o := range opts {
		o(&opt)
	}

	if opt.env {
		return loader([]byte(os.ExpandEnv(string(content))), v)
	}

	return loader(content, v)
}

// LoadConfig loads config into v from file, .json, .yaml and .yml are acceptable.
// Deprecated: use Load instead.
func LoadConfig(file string, v interface{}, opts ...Option) error {
	return Load(file, v, opts...)
}

// LoadFromJsonBytes loads config into v from content json bytes.
func LoadFromJsonBytes(content []byte, v interface{}) error {
	var m map[string]interface{}
	if err := jsonx.Unmarshal(content, &m); err != nil {
		return err
	}

	finfo := buildFieldsInfo(reflect.TypeOf(v))
	camelCaseKeyMap := toCamelCaseKeyMap(m, finfo)

	return mapping.UnmarshalJsonMap(camelCaseKeyMap, v, mapping.WithCanonicalKeyFunc(toCamelCase))
}

// LoadConfigFromJsonBytes loads config into v from content json bytes.
// Deprecated: use LoadFromJsonBytes instead.
func LoadConfigFromJsonBytes(content []byte, v interface{}) error {
	return LoadFromJsonBytes(content, v)
}

// LoadFromTomlBytes loads config into v from content toml bytes.
func LoadFromTomlBytes(content []byte, v interface{}) error {
	b, err := encoding.TomlToJson(content)
	if err != nil {
		return err
	}

	return LoadFromJsonBytes(b, v)
}

// LoadFromYamlBytes loads config into v from content yaml bytes.
func LoadFromYamlBytes(content []byte, v interface{}) error {
	b, err := encoding.YamlToJson(content)
	if err != nil {
		return err
	}

	return LoadFromJsonBytes(b, v)
}

// LoadConfigFromYamlBytes loads config into v from content yaml bytes.
// Deprecated: use LoadFromYamlBytes instead.
func LoadConfigFromYamlBytes(content []byte, v interface{}) error {
	return LoadFromYamlBytes(content, v)
}

// MustLoad loads config into v from path, exits on error.
func MustLoad(path string, v interface{}, opts ...Option) {
	if err := Load(path, v, opts...); err != nil {
		log.Fatalf("error: config file %s, %s", path, err.Error())
	}
}

func buildFieldsInfo(tp reflect.Type) map[string]fieldInfo {
	tp = mapping.Deref(tp)

	switch tp.Kind() {
	case reflect.Struct:
		return buildStructFieldsInfo(tp)
	case reflect.Array, reflect.Slice:
		return buildFieldsInfo(mapping.Deref(tp.Elem()))
	default:
		return nil
	}
}

func buildStructFieldsInfo(tp reflect.Type) map[string]fieldInfo {
	info := make(map[string]fieldInfo)

	for i := 0; i < tp.NumField(); i++ {
		field := tp.Field(i)
		name := field.Name
		ccName := toCamelCase(name)
		ft := mapping.Deref(field.Type)

		// flatten anonymous fields
		if field.Anonymous {
			if ft.Kind() == reflect.Struct {
				fields := buildFieldsInfo(ft)
				for k, v := range fields {
					info[k] = v
				}
			} else {
				info[ccName] = fieldInfo{
					name: name,
					kind: ft.Kind(),
				}
			}
			continue
		}

		var fields map[string]fieldInfo
		switch ft.Kind() {
		case reflect.Struct:
			fields = buildFieldsInfo(ft)
		case reflect.Array, reflect.Slice:
			fields = buildFieldsInfo(ft.Elem())
		case reflect.Map:
			fields = buildFieldsInfo(ft.Elem())
		}

		info[ccName] = fieldInfo{
			name:     name,
			kind:     ft.Kind(),
			children: fields,
		}
	}

	return info
}

func toCamelCase(s string) string {
	var buf strings.Builder
	buf.Grow(len(s))
	var capNext bool
	boundary := true
	for _, v := range s {
		isCap := v >= 'A' && v <= 'Z'
		isLow := v >= 'a' && v <= 'z'
		if boundary && (isCap || isLow) {
			if capNext {
				if isLow {
					v -= distanceBetweenUpperAndLower
				}
			} else {
				if isCap {
					v += distanceBetweenUpperAndLower
				}
			}
			boundary = false
		}
		if isCap || isLow {
			buf.WriteRune(v)
			capNext = false
			continue
		}

		switch v {
		// '.' is used for chained keys, e.g. "grand.parent.child"
		case ' ', '.', '\t':
			buf.WriteRune(v)
			capNext = false
			boundary = true
		case '_':
			capNext = true
			boundary = true
		default:
			buf.WriteRune(v)
			capNext = true
		}
	}

	return buf.String()
}

func toCamelCaseInterface(v interface{}, info map[string]fieldInfo) interface{} {
	switch vv := v.(type) {
	case map[string]interface{}:
		return toCamelCaseKeyMap(vv, info)
	case []interface{}:
		var arr []interface{}
		for _, vvv := range vv {
			arr = append(arr, toCamelCaseInterface(vvv, info))
		}
		return arr
	default:
		return v
	}
}

func toCamelCaseKeyMap(m map[string]interface{}, info map[string]fieldInfo) map[string]interface{} {
	res := make(map[string]interface{})

	for k, v := range m {
		ti, ok := info[k]
		if ok {
			res[k] = toCamelCaseInterface(v, ti.children)
			continue
		}

		cck := toCamelCase(k)
		if ti, ok = info[cck]; ok {
			res[toCamelCase(k)] = toCamelCaseInterface(v, ti.children)
		} else {
			res[k] = v
		}
	}

	return res
}
