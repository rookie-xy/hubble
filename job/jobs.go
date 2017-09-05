package job

import (
    "fmt"
    "sync"

    "github.com/rookie-xy/hubble/log"
    "github.com/satori/go.uuid"
)

type Jobs struct {
    sync.RWMutex

    jobs  map[uuid.UUID]Job
    wg    sync.WaitGroup
    done  chan struct{}
    log   log.Log
}

// NewJobs creates a new registry object
func New(log log.Log) *Jobs {
    return &Jobs{
        log: log,
        jobs: map[uuid.UUID]Job{},
        done: make(chan struct{}),
    }
}

func (r *Jobs) remove(j Job) {
    r.Lock()
    defer r.Unlock()

    delete(r.jobs, j.ID())
}

func (r *Jobs) add(j Job) {
    r.Lock()
    defer r.Unlock()

    r.jobs[j.ID()] = j
}

// Stop stops all jobs in the registry
func (r *Jobs) Stop() {
    r.Lock()
    defer func() {
        r.Unlock()
        r.WaitForCompletion()
    }()

    // Makes sure no new jobs are added during stopping
    close(r.done)

    for _, job := range r.jobs {
        go func(j Job) {
            j.Stop()
        }(job)
    }
}

// WaitForCompletion can be used to wait until all jobs are stopped
func (r *Jobs) WaitForCompletion() {
    r.wg.Wait()
}

// Start starts the given harvester and add its to the registry
func (r *Jobs) Start(j Job) {
    // Make sure stop is not called during starting a harvester
    r.Lock()
    defer r.Unlock()

    // Make sure no new jobs are started after stop was called
    if !r.active() {
        return
    }

    r.wg.Add(1)
    go func() {
        defer func() {
            r.remove(j)
            r.wg.Done()
        }()

        r.add(j)

        // Starts harvester and picks the right type. In case
        // type is not set, set it to default (log)
        err := j.Run()
        if err != nil {
            //r.log.Print("Error running prospector: %v", err)
            fmt.Println("Error running prospector: %v", err)
        }
    }()
}

// Len returns the current number of jobs in the registry
func (r *Jobs) Len() uint64 {
    r.RLock()
    defer r.RUnlock()

    return uint64(len(r.jobs))
}

func (r *Jobs) active() bool {
    select {

    case <-r.done:
        return false

    default:
        return true
    }
}
