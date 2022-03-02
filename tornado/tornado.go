package tornado

import (
	"time"

	"github.com/bcyxy/tornado/executor"
	"github.com/bcyxy/tornado/interact"
)

func Run() {
	for {
		tasks := interact.Interact.PullTasks(10)
		if len(tasks) == 0 {
			time.Sleep(time.Second)
		}

		executor.DoJobs(tasks)
	}
}
