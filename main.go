package main

import (
    "os"
    "os/signal"
    "syscall"

    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/paths"
    "github.com/rookie-xy/hubble/module"
    "github.com/rookie-xy/hubble/command"
    "github.com/rookie-xy/hubble/builder"

  //_ "github.com/rookie-xy/modules"
)

var (
    version = command.New("-v", "version",  "0.0.1",   "Display engine version, golang version, " +
                                                                            "system architecture and other information"    )
    help    = command.New("-?",  "help",    "",        "Assist information on how to use the system"  )
    check   = command.New("-cc", "check",   false,     "Pre check before system startup"              )
    home    = command.New("-h",  "home",     paths.Home(),    "Program root path"                           )

    verbose = command.New("-V", "verbose", true,       "output detail info")
    level   = command.New("-l", "level",   "info",     "output detail info")
    prefix  = command.New("-p", "prefix",  "[hubble]", "output detail info")
)

var commands = []command.Item{

    { version,
      command.LINE,
      module.Worker,
      Name,
      command.Display,
      nil },

    { help,
      command.LINE,
      module.Worker,
      Name,
      command.List,
      nil },

    { check,
      command.LINE,
      module.Worker,
      Name,
      command.SetObject,
      nil },

    { home,
      command.LINE,
      module.Worker,
      Name,
      command.SetObject,
      nil },

    { verbose,
      command.LINE,
      module.Worker,
      Name,
      command.SetObject,
      nil },

    { level,
      command.LINE,
      module.Worker,
      Name,
      command.SetObject,
      nil },

    { prefix,
      command.LINE,
      module.Worker,
      Name,
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
        exit(-1)
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
        if err := command.Setup(flag, value); err != nil {
            exit(0)
        }

        i++
    }
}

func main() {
    log := log.New()
    if err := log.Init(prefix.GetValue(), verbose.GetValue(),
                       level.GetValue()); err != nil {
        exit(1)
    }

    if value := home.GetValue(); value != nil {
        paths.Init(value)
    }

    core := []string{
        module.Proxys,
        module.Agents,
    }

    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

    module := module.New(log)

    builder := builder.New(log)
    if err := builder.Director(module); err != nil {
        exit(-1)
	}

    builder.Construct(core)

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
