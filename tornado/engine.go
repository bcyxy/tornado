package tornado

import (
	"context"
	"time"

	"github.com/bcyxy/tornado/executor"
	"github.com/bcyxy/tornado/interact"
)

var engine = engineT{}

type engineT struct {
}

func (sf *engineT) Start(ctx context.Context, name string, msgCh chan<- string) error {
	go sf.run()
	return nil
}

func (sf *engineT) run() {
	for {
		tasks := interact.Interact.PullTasks(10)
		if len(tasks) == 0 {
			time.Sleep(time.Second)
		}

		executor.DoJobs(tasks)
	}
}
