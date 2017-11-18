package file

type ID interface {
    Same(uint64, uint64) bool
    String() string
}
