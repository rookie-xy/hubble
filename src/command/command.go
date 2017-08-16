package command

import (
    "github.com/rookie-xy/hubble/src/prototype"
    "fmt"
    "github.com/rookie-xy/hubble/src/state"
    "github.com/rookie-xy/hubble/src/plugin"
//    "github.com/rookie-xy/hubble/src/module"
    "strings"
    "github.com/rookie-xy/hubble/src/value"
    "github.com/rookie-xy/hubble/src/types/array"
)

const (
    LINE = 1
    FILE = 2
)

type SetFunc func(cmd *Item, meta *Command, val prototype.Object) int

type Item struct {
    Command  *Command
    Type      int
    Scope     string
    Set       SetFunc
    State     bool
    Offset    uintptr
    Load      prototype.Object
}

type Command struct {
    Flag     string
    Key      string
    Value    prototype.Object
    Details  string
}

func New(flag string, key string, value prototype.Object, details string) *Command {
    return &Command{ flag, key, value, details }
}

func (r *Command) GetFlag() string {
    return r.Flag
}

func (r *Command) GetKey() string {
    return r.Key
}

func (r *Command) GetDetails() string {
    key := ""
    if v := r.Details; v != key {
        return v
    }
    return key
}

func (r *Command) GetString() string {
    if v := r.Value; v != nil {
        return v.(string)
    }

    return ""
}

func (r *Command) GetInt() int {
    if v := r.Value; v != nil {
        return v.(int)
    }

    return state.Error
}

func (r *Command) GetMap() map[interface{}]interface{} {
    if v := r.Value; v != nil {
        return v.(map[interface{}]interface{})
    }

    return nil
}

func (r *Command) GetArray() []interface{} {
    if v := r.Value; v != nil {
        return v.([]interface{})
    }

    return nil
}

func (r *Command) Get() value.Value {

    if v := r.Value; v != nil {
        switch v.(type) {

        case []interface{}:
            return array.New(v)

        case map[interface{}]interface{}:
            return map.New(v)

        }
    }

    return nil
}

func (r *Command) Clear() {
    r.Value = nil
}

/*
func (r *Command) GetArrays() []types.Array {
    if v := r.Value; v != nil {
        for k, v := range v {

        }

        return v.([]interface{})
    }

    return nil
}
*/

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

func List(_ *Item, _ *Command, _ prototype.Object) int {
    for _, item := range Pool {
        if item.Type != LINE {
            continue
        }

        if Command := item.Command; Command != nil {
            fmt.Printf("%s\t%s\t\t%s\n", Command.Flag, Command.Key, Command.Details)
        }
    }

    return state.Done
}

func Display(_ *Item, meta *Command, _ prototype.Object) int {
    if meta != nil {
        fmt.Println(meta.Details)
    }

    return state.Done
}

func SetObject(_ *Item, c *Command, value prototype.Object) int {
    if c == nil || value == nil {
        return state.Error
    }

    c.Value = value

    return state.Ok
}
/*
func SetArray(_ *Item, meta *Command, value prototype.Object) int {
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
	Commands []Command
}

func (self *Invoker) StoreCommand(Command Command) {
	self.Commands = append(self.Commands, Command)
}

func (self *Invoker) UnStoreCommand() {
	if len(self.Commands) != 0 {
		self.Commands = self.Commands[:len(self.Commands)-1]
	}
}

func (self *Invoker) Execute() string {
	var result string
	for _, Command := range self.Commands {
		result += Command.Execute() + "\n"
	}
	return result
}
*/
