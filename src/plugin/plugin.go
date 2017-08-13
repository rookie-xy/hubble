package plugin

import "fmt"

const Flag = "plugin"

type Loader func(p interface{}) error

var registry = map[string]Loader{}

func Bundle(bundles ...map[string][]interface{}) map[string][]interface{} {
    ret := map[string][]interface{}{}

    for _, bundle := range bundles {
        for name, plugins := range bundle {
            ret[name] = append(ret[name], plugins...)
        }
    }

    return ret
}

func Make(key string, ifc interface{}) map[string][]interface{} {
    return map[string][]interface{}{
        key: {ifc},
    }
}

func MustRegisterLoader(name string, loader Loader) {
    err := RegisterLoader(name, loader)
    if err != nil {
        panic(err)
    }
}

func RegisterLoader(name string, loader Loader) error {
    if loader := registry[name]; loader != nil {
        return fmt.Errorf("plugin loader '%v' already registered", name)
    }

    registry[name] = loader
    return nil
}

func Load(path string) error {
	// TODO: add flag to enable/disable plugins?
    return load(path)
}
