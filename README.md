# wrq
A simple worker queue, as explained in [this](http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/) blog by [marcio](http://marcio.io/) with funcionality to gracefully finish the jobs when the worker exits

## Usage
```go
package main

import (
  "time"
  "fmt"
  "github.com/sankalpjonn/wrq"
)

type job struct {
  name string
  delay time.Duration
}

func(j *job) Name() string {
  return j.name
}

func(j *job) Execute() error {
  time.Sleep(j.delay)
  fmt.Println("executed job:", j.name)
  return nil
}

func main() {
  // returns a dispatcher with 100 workers and a queue size of 100
  w := wrq.New()
  defer w.Stop()

  // to customise the number of workers and queue size, use
  // w := wrq.NewWithSettings(NAME, QUEUE_SIZE, MAX_WORKERS)
  
  j1 := &job{
    name: "test job 1 sec",
    delay: time.Second * 1,
  }

  j2 := &job{
    name: "test job 2 sec",
    delay: time.Second * 2,
  }

  w.AddJob(j1)
  w.AddJob(j2)

  // wait for jobs to execute
  time.Sleep(time.Second * 5)
}
```
