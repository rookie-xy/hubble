package state

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

/*
type ModileAlertStater interface {
	Alert() string
}

type ModileAlert struct {
	state ModileAlertStater
}

func (self *ModileAlert) Alert() string {
	return self.state.Alert()
}

func (self *ModileAlert) SetState(state ModileAlertStater) {
	self.state = state
}

func NewModileAlert() *ModileAlert {
	return &ModileAlert{state: &MobileAlertVibration{}}
}

type MobileAlertVibration struct {
}

func (self *MobileAlertVibration) Alert() string {
	return "Vrrr... Brrr... Vrrr..."
}

type MobileAlertSong struct {
}

func (self *MobileAlertSong) Alert() string {
	return "Белые розы, Белые розы. Беззащитны шипы..."
}
*/
