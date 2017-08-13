package register

import (
    "fmt"
    "github.com/rookie-xy/worker/src/observer"
)

func Observer(name string, o observer.Observer) {
    if _, exists := observer.Observers[name]; exists {
        panic(fmt.Sprintf("this observer '%v' already registered ", name))
    }

    observer.Observers[name] = o
}

func Subject(name string, o observer.Subject) {
    if _, exists := observer.Subjects[name]; exists {
        panic(fmt.Sprintf("this subject '%v' already registered ", name))
    }

    observer.Subjects[name] = o
}
