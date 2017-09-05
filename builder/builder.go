package builder

import (
    "github.com/rookie-xy/hubble/module"
)

type Builder interface {
    Configure(m module.Template) int
    module.Module
}


