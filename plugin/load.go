package plugin

import (
    "errors"
    "plugin"
)

func load(path string) error {
    p, err := plugin.Open(path)
    if err != nil {
        return err
    }

    sym, err := p.Lookup("Bundle")
    if err != nil {
        return err
    }

    ptr, ok := sym.(*map[string][]interface{})
    if !ok {
        return errors.New("invalid bundle type")
    }

    bundle := *ptr
    for name, plugins := range bundle {
        loader := registry[name]
        if loader == nil {
            continue
        }

        for _, plugin := range plugins {
            if err := loader(plugin); err != nil {
                return err
            }
        }
    }

    return nil
}
