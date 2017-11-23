package register

import (
    "fmt"
    "github.com/rookie-xy/hubble/observer"
)

func Observer(name string, o observer.Observer) {
    if _, exists := observer.Observers[name]; exists {
        fmt.Printf("This observer '%v' already registered\n", name)
        return
    }

    observer.Observers[name] = o
}

func Subject(name string, o observer.Subject) {
    if _, exists := observer.Subjects[name]; exists {
        fmt.Printf("This subject '%v' already registered\n", name)
        return
    }

    observer.Subjects[name] = o
}
