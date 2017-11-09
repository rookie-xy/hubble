package state

type State interface {
    On() bool
    Off()
}
/*
type ModileAlertStater interface {
	Alert() string
}

type ModileAlert struct {
	file ModileAlertStater
}

func (self *ModileAlert) Alert() string {
	return self.file.Alert()
}

func (self *ModileAlert) SetState(file ModileAlertStater) {
	self.file = file
}

func NewModileAlert() *ModileAlert {
	return &ModileAlert{file: &MobileAlertVibration{}}
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
