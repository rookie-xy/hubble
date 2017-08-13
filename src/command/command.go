package command

import (
    "github.com/rookie-xy/worker/src/prototype"
    "fmt"
    "github.com/rookie-xy/worker/src/state"
)

const (
    LINE = 1
    FILE = 2
)

type SetFunc func(cmd *Item, meta *Meta, val prototype.Object) int

type Item struct {
    Meta   *Meta
    Type    int
    Scope   string
    Set     SetFunc
    State   bool
    Offset  uintptr
    Load    prototype.Object
}

type Meta struct {
    Flag     string
    Key      string
    Value    prototype.Object
    Details  string
}

func New(flag string, key string, value prototype.Object, details string) *Meta {
    return &Meta{ flag, key, value, details }
}

var Pool []Item

func Setup(flag, value string) int {
    for _, item := range Pool {

        if item.Type != LINE || item.Meta.Flag != flag {
            continue
        }

        return item.Set(&item, item.Meta, value)
    }

    return state.Error
}

func File(nameSpace, key string, value prototype.Object) int {
    for _, item := range Pool {

        if item.Scope != nameSpace ||
                 item.Type != FILE || item.Meta.Key != key {

            continue
        }

        return item.Set(&item, item.Meta, value)
    }

    return state.Error
}

func List(_ *Item, _ *Meta, _ prototype.Object) int {
    for _, item := range Pool {
        if item.Type != LINE {
            continue
        }

        if meta := item.Meta; meta != nil {
            fmt.Printf("%s\t%s\t\t%s\n", meta.Flag, meta.Key, meta.Details)
        }
    }

    return state.Done
}

func Display(_ *Item, meta *Meta, _ prototype.Object) int {
    if meta != nil {
        fmt.Println(meta.Details)
    }

    return state.Done
}

func SetObject(_ *Item, meta *Meta, value prototype.Object) int {
    if meta == nil || value == nil {
        return state.Error
    }

    meta.Value = value

    return state.Ok
}
/*
func SetArray(_ *Item, meta *Meta, value prototype.Object) int {
    if meta == nil || value == nil {
        return state.Error
    }

    for i, v := range value.([]interface{}) {
        fmt.Println("yyyyyyyyyyhhhhhhhhhhhhhhhhhh", i, v)
    }

    //meta.Value = value

    return state.Ok
}
*/

/*
type Command interface {
   	Execute() string
}

type Option struct {

}

type Command interface {
	Execute() string
}

type ToggleOnCommand struct {
	receiver *Receiver
}

func (self *ToggleOnCommand) Execute() string {
	return self.receiver.ToggleOn()
}

type ToggleOffCommand struct {
	receiver *Receiver
}

func (self *ToggleOffCommand) Execute() string {
	return self.receiver.ToggleOff()
}

type Receiver struct {
}

func (self *Receiver) ToggleOn() string {
	return "Toggle On"
}

func (self *Receiver) ToggleOff() string {
	return "Toggle Off"
}

type Invoker struct {
	commands []Command
}

func (self *Invoker) StoreCommand(command Command) {
	self.commands = append(self.commands, command)
}

func (self *Invoker) UnStoreCommand() {
	if len(self.commands) != 0 {
		self.commands = self.commands[:len(self.commands)-1]
	}
}

func (self *Invoker) Execute() string {
	var result string
	for _, command := range self.commands {
		result += command.Execute() + "\n"
	}
	return result
}
*/
