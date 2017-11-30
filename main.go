package main

import (
    "os"
    "os/signal"
    "syscall"
    "fmt"

    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/paths"
    "github.com/rookie-xy/hubble/module"
    "github.com/rookie-xy/hubble/command"
    "github.com/rookie-xy/hubble/builder"
    "github.com/rookie-xy/hubble/log/level"

  _ "github.com/rookie-xy/modules"
)

var (
    version = command.New("-v", "version",  "0.0.1",   "Display engine version, golang version, " +
                                                                            "system architecture and other information"    )
    help    = command.New("-?",  "help",    "",        "Assist information on how to use the system"  )
    check   = command.New("-cc", "check",   false,     "Pre check before system startup"              )
    home    = command.New("-h",  "home",     paths.Home(),    "Program root path"                           )
)

var commands = []command.Item{

    { version,
      command.LINE,
      module.Worker,
      "main",
      command.Version,
      nil },

    { help,
      command.LINE,
      module.Worker,
      "main",
      command.Help,
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
        command.Line(help.GetFlag(), help.GetFlag())
        exit(0)
    }

    for i := 1; i < argc; i++ {
        if argv[i][0] != '-' {
            fmt.Println("format is failure")
            exit(0)
        }

        j := i
        if j = i + 1; j >= argc {
            j = i
        }

        flag, value := argv[i], argv[j]
        if err := command.Line(flag, value); err != nil {
            fmt.Println(err)
            exit(0)
        }

        i++
    }
}

func main() {
    log := log.New()

    if value := home.GetValue(); value != nil {
        paths.Init(value)
    } else {
        log.Print(level.INFO, "not found home paths")
    }

    component:= []string{
        module.Proxys,
        module.Agents,
    }

    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

    module := module.New(log)

    builder := builder.New(log)
    if err := builder.Director(module); err != nil {
        log.Print(level.FATAL, err.Error())
        exit(0)
	}

    if err := builder.Construct(component); err != nil {
        log.Print(level.FATAL, err.Error())
        exit(0)
    }

    module.Init()

    if value := check.GetValue(); value != nil {
        if value.GetBool() {
            exit(-1)
        }
    }

    module.Main()
    <-signalChan
    module.Exit(0)
}

func exit(code int) {
    os.Exit(code)
}
