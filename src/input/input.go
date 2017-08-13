package input

import (
    "github.com/rookie-xy/worker/src/prototype"
    "github.com/rookie-xy/worker/src/observer"
)

const Name = "input"
/*
type Input struct {
    observers []observer.Observer
    Data prototype.Object
    Configure chan prototype.Object
}

func New() *Input {
    return &Input{}
}

func (r *Input) Attach(obs observer.Observer) {
    r.observers = append(r.observers, obs)
}

func (r *Input) Notify() {
    for _, observer := range r.observers {
         observer.Update(r.Data)
    }
}

func (r *Input) Update(configure prototype.Object) {
    if configure != nil {
        r.Configure <- configure
    }

    // 初始化命令
    // load 模块
}

func (r *Input) Create() prototype.Object {
    //c := <- r.Configure
    //return c
    return nil
}
*/
