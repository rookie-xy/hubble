package configure

import (
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/module"
)

// the build
type build func(name string, i types.Iterator, l module.Load) error

var Build build = nil
