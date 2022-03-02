package interact

type localStore struct{}

func (sf *localStore) PullTasks(n uint) (jobs []string) {
	return []string{"job1", "job2"}
}

func (sf *localStore) PushTasks(jobs []string) (err error) {
	return
}

func (sf *localStore) OutputRet(ret interface{}) (err error) {
	return
}
