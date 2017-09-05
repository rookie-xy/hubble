package job

import (
    "github.com/satori/go.uuid"
)

// bridge pattern
type Job interface {
    ID() uuid.UUID
    Run() error
    Stop()
}
