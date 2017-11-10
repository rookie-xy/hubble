package value

import (
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/configure"
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

func (r *Value) GetInt() (int, error) {
    if obj := r.Object; obj != nil {
        return obj.(int), nil
    }

    return -1, fmt.Errorf("Not found object")
}

func (r *Value) GetUint64() (uint64, error) {
    if obj := r.Object; obj != nil {
        return obj.(uint64), nil
    }

    // TODO 修改错误信息
    return 0, fmt.Errorf("Not found object")
}

func (r *Value) GetBool() bool {
    if obj := r.Object; obj != nil {
        return obj.(bool)
    }

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
            return types.ARRAY

        case map[interface{}]interface{}:
            return types.MAP

        case string:
            return types.STRING

        case int:
            return types.INTEGER
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

    case types.ARRAY:
        if iterms := r.GetArray(); iterms != nil {
            for _, iterm := range iterms {
                r.Object = iterm
                r.GetIterator(c)
            }
        }

    case types.MAP:
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

func (r *Value) GetDuration() (time.Duration, error) {
    if obj := r.Object; obj != nil {

        switch r.GetType() {

        case types.STRING:
            duration, err := time.ParseDuration(r.GetString())
            if err != nil {
                return duration, err
            }
            return duration, nil

        default:
            return obj.(time.Duration), nil
        }
    }

    // TODO 修改错误信息
    return -1, fmt.Errorf("get duration object is nil")
}
