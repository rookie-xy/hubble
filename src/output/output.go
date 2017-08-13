package output

// proxy
/*
type Output interface {
    Send()
}
*/

type Factory func(*Config) (Output, error)

type Output interface {
}

var Ouputs = map[string]Factory{}

type Config struct {
    Name      string
    configure string
}
