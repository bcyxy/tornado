package task

type Task struct {
	Key string

	Param interface{}
}

type HttpTaskParam struct {
	URL    string
	Header map[string]string
	Date   map[string]string
}
