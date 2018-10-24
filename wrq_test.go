package wrq

import (
	"testing"
	"time"
)

type job struct {
  name  string
  delay time.Duration
}

func newJob(name string, delay time.Duration) *job {
  return &job{name:name, delay:delay}
}

func(j *job) Name() string {
  return j.name
}

func(j *job) Execute() error {
  time.Sleep(j.delay)
  return nil
}

func TestDispatcher_Run_Test_10m(t *testing.T) {
	delay := 10 * time.Millisecond
	jobs := []*job{
		newJob("1", delay),
		newJob("2", delay),
		newJob("3", delay),
		newJob("4", delay),
		newJob("5", delay),
		newJob("6", delay),
		newJob("7", delay),
		newJob("8", delay),
		newJob("9", delay),
		newJob("10", delay),
	}

  w := New()
  w.Run()
  defer w.Stop()

	for _, j := range jobs {
		w.AddJob(j)
	}
	time.Sleep(1 * time.Millisecond)
}
