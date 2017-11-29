package register

import (
	"fmt"
	"github.com/rookie-xy/hubble/module"
)

func Modules(key string, f module.Factory) error {
    if _, exist := module.Pool[key]; !exist {
        module.Pool[key] = &f
        return nil
    }

    return fmt.Errorf("The %s is exist", key)
}
