package state

type State struct {
	On bool
}

func New() State {
	return State{
		On: false,
	}
}
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
