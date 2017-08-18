package value

import (
    "github.com/rookie-xy/hubble/src/prototype"
    "github.com/rookie-xy/hubble/src/types"
    "github.com/rookie-xy/hubble/src/configure"
)

type value struct {
    prototype.Object
}

func New(o prototype.Object) *value {
    return &value{
        Object: o,
    }
}

func (r *value) GetString() string {
    return r.Object.(string)
}

func (r *value) GetInt() int {
    return r.Object.(int)
}

func (r *value) GetArray() []interface{} {
    return r.Object.([]interface{})
}

func (r *value) GetMap() map[interface{}]interface{} {
    return r.Object.(map[interface{}]interface{})
}

func (r *value) GetType() int {

    if value := r.Object; value != nil {

        switch value.(type) {

        case []interface{}:
            return types.Array

        case map[interface{}]interface{}:
            return types.Map
        }
    }

    return -1
}

func (r *value) GetIterator(cfg prototype.Object) types.Iterator {
    if cfg == nil {
        cfg = configure.New()
    }

    c := cfg.(*configure.Configure)

    switch r.GetType() {

    case types.Array:
        if iterms := r.GetArray(); iterms != nil {
            for _, iterm := range iterms {
                r.Object = iterm
                r.GetIterator(c)
            }
        }

    case types.Map:
        if iterms := r.GetMap(); iterms != nil {
            for k, v := range iterms {
                c.Add(New(k), New(v))
            }
        }

        return c.Iterator()

    default:
        return nil
    }

    return c.Iterator()
}
