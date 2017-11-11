package state

type State interface {
    On() bool
    Off()
}
