package prototype

import "github.com/rookie-xy/hubble/pipeline"

type PipelinePrototype interface {
	Prototype
	pipeline.Queue
}

func Pipeline(this pipeline.Queue) pipeline.Queue {
    prototype := this.(PipelinePrototype)
    return prototype.Clone().(pipeline.Queue)
}
