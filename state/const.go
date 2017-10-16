package state

type State interface {
    Enable()
    Disable()
}

const (
    Enable  = true
    Disable = false
)

const (
    Ok       = 0
    Error    = -1
    Again    = -2
    Busy     = -3
    Done     = -4
    Declined = -5
    Ignore   = -6
    Abort    = -7
)

const (
    RELOAD = 1
    RECONFIGURE = 2
    EXIT = 3
)
