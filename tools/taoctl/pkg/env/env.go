package env

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"manlu.org/tao/tools/taoctl/internal/version"
	sortedmap "manlu.org/tao/tools/taoctl/pkg/collection"
	"manlu.org/tao/tools/taoctl/pkg/protoc"
	"manlu.org/tao/tools/taoctl/pkg/protocgengo"
	"manlu.org/tao/tools/taoctl/pkg/protocgengogrpc"
	"manlu.org/tao/tools/taoctl/util/pathx"
)

var taoctlEnv *sortedmap.SortedMap

const (
	TaoctlOS               = "GOCTL_OS"
	TaoctlArch             = "GOCTL_ARCH"
	TaoctlHome             = "GOCTL_HOME"
	TaoctlDebug            = "GOCTL_DEBUG"
	TaoctlCache            = "GOCTL_CACHE"
	TaoctlVersion          = "GOCTL_VERSION"
	ProtocVersion          = "PROTOC_VERSION"
	ProtocGenGoVersion     = "PROTOC_GEN_GO_VERSION"
	ProtocGenGoGRPCVersion = "PROTO_GEN_GO_GRPC_VERSION"

	envFileDir = "env"
)

// init initializes the taoctl environment variables, the environment variables of the function are set in order,
// please do not change the logic order of the code.
func init() {
	defaultTaoctlHome, err := pathx.GetDefaultTaoctlHome()
	if err != nil {
		log.Fatalln(err)
	}
	taoctlEnv = sortedmap.New()
	taoctlEnv.SetKV(TaoctlOS, runtime.GOOS)
	taoctlEnv.SetKV(TaoctlArch, runtime.GOARCH)
	existsEnv := readEnv(defaultTaoctlHome)
	if existsEnv != nil {
		taoctlHome, ok := existsEnv.GetString(TaoctlHome)
		if ok && len(taoctlHome) > 0 {
			taoctlEnv.SetKV(TaoctlHome, taoctlHome)
		}
		if debug := existsEnv.GetOr(TaoctlDebug, "").(string); debug != "" {
			if strings.EqualFold(debug, "true") || strings.EqualFold(debug, "false") {
				taoctlEnv.SetKV(TaoctlDebug, debug)
			}
		}
		if value := existsEnv.GetStringOr(TaoctlCache, ""); value != "" {
			taoctlEnv.SetKV(TaoctlCache, value)
		}
	}
	if !taoctlEnv.HasKey(TaoctlHome) {
		taoctlEnv.SetKV(TaoctlHome, defaultTaoctlHome)
	}
	if !taoctlEnv.HasKey(TaoctlDebug) {
		taoctlEnv.SetKV(TaoctlDebug, "False")
	}

	if !taoctlEnv.HasKey(TaoctlCache) {
		cacheDir, _ := pathx.GetCacheDir()
		taoctlEnv.SetKV(TaoctlCache, cacheDir)
	}

	taoctlEnv.SetKV(TaoctlVersion, version.BuildVersion)
	protocVer, _ := protoc.Version()
	taoctlEnv.SetKV(ProtocVersion, protocVer)

	protocGenGoVer, _ := protocgengo.Version()
	taoctlEnv.SetKV(ProtocGenGoVersion, protocGenGoVer)

	protocGenGoGrpcVer, _ := protocgengogrpc.Version()
	taoctlEnv.SetKV(ProtocGenGoGRPCVersion, protocGenGoGrpcVer)
}

func Print() string {
	return strings.Join(taoctlEnv.Format(), "\n")
}

func Get(key string) string {
	return GetOr(key, "")
}

func GetOr(key, def string) string {
	return taoctlEnv.GetStringOr(key, def)
}

func readEnv(taoctlHome string) *sortedmap.SortedMap {
	envFile := filepath.Join(taoctlHome, envFileDir)
	data, err := ioutil.ReadFile(envFile)
	if err != nil {
		return nil
	}
	dataStr := string(data)
	lines := strings.Split(dataStr, "\n")
	sm := sortedmap.New()
	for _, line := range lines {
		_, _, err = sm.SetExpression(line)
		if err != nil {
			continue
		}
	}
	return sm
}

func WriteEnv(kv []string) error {
	defaultTaoctlHome, err := pathx.GetDefaultTaoctlHome()
	if err != nil {
		log.Fatalln(err)
	}
	data := sortedmap.New()
	for _, e := range kv {
		_, _, err := data.SetExpression(e)
		if err != nil {
			return err
		}
	}
	data.RangeIf(func(key, value interface{}) bool {
		switch key.(string) {
		case TaoctlHome, TaoctlCache:
			path := value.(string)
			if !pathx.FileExists(path) {
				err = fmt.Errorf("[writeEnv]: path %q is not exists", path)
				return false
			}
		}
		if taoctlEnv.HasKey(key) {
			taoctlEnv.SetKV(key, value)
			return true
		} else {
			err = fmt.Errorf("[writeEnv]: invalid key: %v", key)
			return false
		}
	})
	if err != nil {
		return err
	}
	envFile := filepath.Join(defaultTaoctlHome, envFileDir)
	return ioutil.WriteFile(envFile, []byte(strings.Join(taoctlEnv.Format(), "\n")), 0o777)
}
