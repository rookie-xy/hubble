package value

import (
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/configure"
    "github.com/rookie-xy/hubble/state"
    "fmt"
    "time"
)

type Value struct {
    types.Object
}

func New(o types.Object) *Value {
    return &Value{
        Object: o,
    }
}

func (r *Value) GetString() string {
    if obj := r.Object; obj != nil {
        return obj.(string)
    }

    return ""
}

func (r *Value) GetBytes() []byte {
    if obj := r.Object; obj != nil {
        return obj.([]byte)
    }

    return nil
}

func (r *Value) GetInt() int {
    if obj := r.Object; obj != nil {
        return obj.(int)
    }

    return state.Error
}

func (r *Value) GetUint64() uint64 {
    if obj := r.Object; obj != nil {
        return obj.(uint64)
    }

    // TODO 修改错误信息
    return state.Ok
}

func (r *Value) GetBool() bool {
    if obj := r.Object; obj != nil {
        return obj.(bool)
    }

    fmt.Println("bool value is not found")

    return false
}

func (r *Value) GetArray() []interface{} {
    if obj := r.Object; obj != nil {
        return obj.([]interface{})
    }

    return nil
}

func (r *Value) GetMap() map[interface{}]interface{} {
    if obj := r.Object; obj != nil {
        return obj.(map[interface{}]interface{})
    }

    return nil
}

func (r *Value) GetType() int {

    if Value := r.Object; Value != nil {

        switch Value.(type) {

        case []interface{}:
            return types.Array

        case map[interface{}]interface{}:
            return types.Map

        case string:
            return types.String

        case int:
            return types.Int
        }
    }

    return -1
}

func (r *Value) GetIterator(cfg types.Object) types.Iterator {
    // TODO add log
    if cfg == nil {
        cfg = configure.New(nil)
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
                c.Add(New(k), v)
            }
        }

        return c.Iterator()

    default:
        return nil
    }

    return c.Iterator()
}

func (r *Value) GetDuration() time.Duration {
    if obj := r.Object; obj != nil {

        switch r.GetType() {

        case types.String:
            duration, err := time.ParseDuration(r.GetString())
            if err != nil {
                return state.Error
            }
            return duration

        default:
            return obj.(time.Duration)
        }
    }

    // TODO 修改错误信息
    return state.Error
}
