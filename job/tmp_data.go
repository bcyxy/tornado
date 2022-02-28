package job

var tmpItemConfs = []string{
	`{
		"key": "frxxz_main",
		"input": {
			"type": "CMD",
			"param": {
				"cmd": "seq 2 5"
			}
		},
		"parse": {
			"type": "REG_CUT",
			"param": {
				"reg": "(\\d+)",
				"keys": [
					"chapter"
				]
			}
		},
		"output": {
			"type": "JOB",
			"data": "PARSE",
			"param": {
				"key": "frxxz_mp3"
			}
		}
	}`,
	`{
		"key": "frxxz_mp3",
		"input": {
			"type": "HTTP_GET",
			"interval": 30,
			"param": {
				"url": "https://zerohour.tcbbw.com/70ts_mp3/{chapter:%04d}.mp3",
				"header": {
					"User-Agent": "Mozilla/5.0"
				},
				"data": {}
			}
		},
		"output": {
			"type": "FILE",
			"data": "ORIGINAL",
			"param": {
				"path": "/home/data/{chapter:%04d}"
			}
		}
	}`,
}
