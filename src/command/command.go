package command

import (
    "github.com/rookie-xy/hubble/src/prototype"
    "fmt"
    "github.com/rookie-xy/hubble/src/state"
    "github.com/rookie-xy/hubble/src/plugin"
//    "github.com/rookie-xy/hubble/src/module"
    "strings"
)

const (
    LINE = 1
    FILE = 2
)

type SetFunc func(cmd *Item, meta *command, val prototype.Object) int

type Item struct {
    Command  *command
    Type      int
    Scope     string
    Set       SetFunc
    State     bool
    Offset    uintptr
    Load      prototype.Object
}

type command struct {
    Flag     string
    Key      string
    Value    prototype.Object
    Details  string
}

func New(flag string, key string, value prototype.Object, details string) *command {
    return &command{ flag, key, value, details }
}

func (r *command) GetFlag() string {
    return r.Flag
}

func (r *command) GetKey() string {
    return r.Key
}

func (r *command) GetDetails() string {
    key := ""
    if v := r.Details; v != key {
        return v
    }
    return key
}

func (r *command) GetString() string {
    if v := r.Value; v != nil {
        return v.(string)
    }

    return ""
}

func (r *command) GetInt() int {
    if v := r.Value; v != nil {
        return v.(int)
    }

    return state.Error
}
/*
func (r *command) GetObject() prototype.Object {
    if v := r.Value; v != nil {
        switch v {

        case map[interface{}]interface{}:
            return r.GetMap()

        case []interface{}:
            return r.GetArray()
        }
    }

    return nil
}
*/

func (r *command) GetMap() map[interface{}]interface{} {
    if v := r.Value; v != nil {
        return v.(map[interface{}]interface{})
    }

    return nil
}

func (r *command) GetArray() []interface{} {
    if v := r.Value; v != nil {
        return v.([]interface{})
    }

    return nil
}

var Pool []Item

func Setup(flag, value string) int {
    for _, item := range Pool {

        if item.Type != LINE || item.Command.Flag != flag {
            continue
        }

        return item.Set(&item, item.Command, value)
    }

    return state.Error
}

func File(nameSpace, key string, value prototype.Object) int {
    for _, item := range Pool {

        if item.Scope != nameSpace || item.Type != FILE {
            continue
        }

        if item.Command.Key != key {
            prefix := item.Command.Key
            if n := strings.Index(prefix, "."); n > -1 {
                prefix = prefix[0:n]
            }

            if item.Command.Flag == plugin.Flag {
                if strings.HasPrefix(key, prefix) {
                    item.Command.Key = key
                } else {
                    continue
                }

            } else {
                continue
            }
        }

        return item.Set(&item, item.Command, value)
    }

    return state.Error
}

func List(_ *Item, _ *command, _ prototype.Object) int {
    for _, item := range Pool {
        if item.Type != LINE {
            continue
        }

        if command := item.Command; command != nil {
            fmt.Printf("%s\t%s\t\t%s\n", command.Flag, command.Key, command.Details)
        }
    }

    return state.Done
}

func Display(_ *Item, meta *command, _ prototype.Object) int {
    if meta != nil {
        fmt.Println(meta.Details)
    }

    return state.Done
}

func SetObject(_ *Item, c *command, value prototype.Object) int {
    if c == nil || value == nil {
        return state.Error
    }

    c.Value = value

    return state.Ok
}
/*
func SetArray(_ *Item, meta *command, value prototype.Object) int {
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
