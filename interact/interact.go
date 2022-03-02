package interact

var Interact InteractIf = &localStore{}

type InteractIf interface {
	PullTasks(n uint) []string
	PushTasks(jobs []string) error
	OutputRet(ret interface{}) error
}
