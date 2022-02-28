package job

var JobTable = make(map[string]Item)

type Item struct {
	Key    string      `json:"key"`
	Input  interface{} `json:"input"`
	Parse  interface{} `json:"parse"`
	Output interface{} `json:"output"`
}

// Load xxx
func Load() (err error) {
	for _, itStr := range tmpItemConfs {
		loadIt(itStr)
	}
	return
}

// loadIt 加载
func loadIt(itStr string) (it Item, err error) {
	return
}
