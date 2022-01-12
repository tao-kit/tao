package xcopy

import (
	"errors"
	"reflect"
)

const (
	noDepthLimited = -1
	noTagLimit     = ""
)

type deepCopy struct {
	dst      interface{}
	src      interface{}
	tagName  string
	maxDepth int

	err error
}

func Copy(dst, src interface{}) error {
	if dst == nil || src == nil {
		return errors.New("Unsupported type:nil")
	}

	d := deepCopy{
		maxDepth: noDepthLimited,
		dst:      dst,
		src:      src,
	}

	return d.deepCopy(reflect.ValueOf(d.dst), reflect.ValueOf(d.src), 0)
}

func (d *deepCopy) RegisterTagName(tagName string) *deepCopy {
	d.tagName = tagName
	return d
}

func haveTagName(curTabName string) bool {
	return len(curTabName) > 0
}

func isArraySlice(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		return true
	}
	return false
}

func (d *deepCopy) cpySliceArray(dst, src reflect.Value, depth int) error {

	if !isArraySlice(dst) {
		return nil
	}

	l := src.Len()
	if dst.Len() > 0 && dst.Len() < src.Len() {
		l = dst.Len()
	}

	if dst.Kind() == reflect.Slice && dst.Len() == 0 && src.Len() > 0 {
		elemType := src.Type().Elem()
		if dst.Type().Elem().Kind() != elemType.Kind() {
			return nil
		}

		newDst := reflect.MakeSlice(reflect.SliceOf(elemType), l, l)
		dst.Set(newDst)
	}

	for i := 0; i < l; i++ {
		if err := d.deepCopy(dst.Index(i), src.Index(i), depth); err != nil {
			return err
		}
	}
	return nil
}

func (d *deepCopy) cpyMap(dst, src reflect.Value, depth int) error {
	if dst.Kind() != src.Kind() {
		return nil
	}

	if !dst.CanSet() {
		return nil
	}

	if dst.Type().Elem().Kind() != src.Type().Elem().Kind() {
		return nil
	}

	if dst.Type().Key().Kind() != src.Type().Key().Kind() {
		return nil
	}

	if dst.IsNil() {

		newMap := reflect.MakeMap(src.Type())
		dst.Set(newMap)
	}

	iter := src.MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()

		newVal := reflect.New(v.Type()).Elem()
		if err := d.deepCopy(newVal, v, depth); err != nil {
			return err
		}

		dst.SetMapIndex(k, newVal)
	}
	return nil
}

func (d *deepCopy) cpyFunc(dst, src reflect.Value, depth int) error {
	if dst.Kind() != src.Kind() {
		return nil
	}

	dst.Set(src)
	return nil
}

func (d *deepCopy) cpyStruct(dst, src reflect.Value, depth int) error {

	if dst.Kind() != src.Kind() {
		if dst.Kind() == reflect.Ptr && !dst.IsNil() {
			dst = dst.Elem()
			return d.cpyStruct(dst, src, depth)
		}
		return nil
	}

	typ := src.Type()
	for i, n := 0, src.NumField(); i < n; i++ {
		sf := typ.Field(i)
		if sf.PkgPath != "" && !sf.Anonymous {
			continue
		}

		if len(d.tagName) > 0 && !haveTagName(sf.Tag.Get(d.tagName)) {
			continue
		}

		dstValue := dst.FieldByName(sf.Name)

		if !dstValue.IsValid() {
			continue
		}

		sField := src.Field(i)

		if err := d.deepCopy(dstValue, sField, depth+1); err != nil {
			return err
		}
	}

	return nil
}

func (d *deepCopy) cpyInterface(dst, src reflect.Value, depth int) error {

	if dst.Kind() != src.Kind() {
		return nil
	}

	src = src.Elem()
	newDst := reflect.New(src.Type()).Elem()

	if err := d.deepCopy(newDst, src, depth); err != nil {
		return err
	}

	dst.Set(newDst)
	return nil
}

func (d *deepCopy) cpyPtr(dst, src reflect.Value, depth int) error {

	if dst.Kind() == reflect.Ptr && dst.IsNil() {
		if !dst.CanSet() {
			return nil
		}
		p := reflect.New(src.Type().Elem())
		dst.Set(p)
	}

	if src.Kind() == reflect.Ptr {
		src = src.Elem()
	}

	if dst.Kind() == reflect.Ptr {
		dst = dst.Elem()
	}

	return d.deepCopy(dst, src, depth)
}

func (d *deepCopy) cpyDefault(dst, src reflect.Value, depth int) error {
	if dst.Kind() != src.Kind() {
		return nil
	}
	dst.Set(src)
	return nil
}

func (d *deepCopy) deepCopy(dst, src reflect.Value, depth int) error {
	if d.err != nil {
		return d.err
	}

	if src.IsZero() {
		return nil
	}

	if d.maxDepth != noDepthLimited && depth > d.maxDepth {
		return nil
	}

	switch src.Kind() {
	case reflect.Slice, reflect.Array:
		return d.cpySliceArray(dst, src, depth)

	case reflect.Map:
		return d.cpyMap(dst, src, depth)

	case reflect.Func:
		return d.cpyFunc(dst, src, depth)

	case reflect.Struct:
		return d.cpyStruct(dst, src, depth)

	case reflect.Interface:
		return d.cpyInterface(dst, src, depth)

	case reflect.Ptr:
		return d.cpyPtr(dst, src, depth)

	default:
		return d.cpyDefault(dst, src, depth)
	}
}
