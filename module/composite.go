package module

// composite
type Module interface {
    Load(module Template)
    Template
}

