package prototype

import "github.com/rookie-xy/hubble/types"

// prototype pattern
type Prototype interface {
    Clone() types.Object
}