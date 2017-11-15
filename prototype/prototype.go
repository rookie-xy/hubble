package prototype

import "github.com/rookie-xy/hubble/types"

type Prototype interface {
    Clone() types.Object
}