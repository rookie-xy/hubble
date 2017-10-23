package adapter

import "github.com/rookie-xy/hubble/pipeline"

type MessageQueue interface {
	pipeline.Queue

	Publish()
	Subscribe()
}