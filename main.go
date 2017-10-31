package main

import (
    "os"

    "github.com/rookie-xy/hubble/command"
    "github.com/rookie-xy/hubble/module"
    "github.com/rookie-xy/hubble/builder"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/state"
    "github.com/rookie-xy/hubble/paths"

  _ "github.com/rookie-xy/modules"
)

var (
    version = command.New("-v", "version", "0.0.1",    "Display engine version, golang version, " +
                                                       "system architecture and other information"    )
    help    = command.New("-?",  "help",    "",        "Assist information on how to use the system"  )
    check   = command.New("-cc", "check",  false,      "Pre check before system startup"              )
    home    = command.New("-h",  "home",   paths.Home(), "Program root path"                           )
)

var commands = []command.Item{

    { version,
      command.LINE,
      module.Worker,
      "main",
      command.Display,
      nil },

    { help,
      command.LINE,
      module.Worker,
      "main",
      command.List,
      nil },

    { check,
      command.LINE,
      module.Worker,
      "main",
      command.SetObject,
      nil },

    { home,
      command.LINE,
      module.Worker,
      "main",
      command.SetObject,
      nil },

}

func init() {
    for _, item := range commands {
        command.Pool = append(command.Pool, item)
    }

    argc, argv := len(os.Args), os.Args
    if argc <= 1 {
        command.Setup(help.GetFlag(), "")
        exit(state.Done)
    }

    for i := 1; i < argc; i++ {
        if argv[i][0] != '-' {
            exit(-1)
        }

        j := i
        if j = i + 1; j >= argc {
            j = i
        }

        flag, value := argv[i], argv[j]
        if status := command.Setup(flag, value); status != state.Ok {
            exit(status)
        }

        i++
    }
}

func main() {
    log := log.New()

    if value := home.GetValue(); value != nil {
        paths.Init(value)
    }

    core := []string{
        module.Proxy,
        module.Agents,
    }

    module := module.New(log)

    director := builder.Directors(module)
    if director == nil {
        exit(state.Error)
    }

    director.Construct(core)

    module.Init()

    if value := check.GetValue(); value != nil {
        if value.GetBool() {
            //TODO if module init exp, then notice
            return
        }
    }

    module.Main()

    module.Exit(state.Ok)
}

func exit(code int) {
    os.Exit(code)
}
