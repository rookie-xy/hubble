package job

import (
    uuid "github.com/satori/go.uuid"
)

// bridge pattern
type Job interface {
    ID() uuid.UUID
    Run() error
    Wait()
    Stop()
}
