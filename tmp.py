{
	"key": "tkd",
	"input": {
		"type": "HTTP_GET",
		"param": {
			"url": "{0}://{1}",
			"header": {},
			"data": {}
		}
	},
	"handle": [
		{
			"parse": {
				"type": "REG_CUT",
				"param": {
					"reg": "(https?)://([\\w\\.\\-]+)"
				}
			},
			"output": {
				"type": "NEW_JOB",
				"param": {
					"exe": "tkd"
				}
			}
		},
		{
			"parse": {
				"type": "NONE"
			},
			"output": {
				"type": "FILE",
				"param": {
					"path": "/home/data/{1}"
				}
			}
		}
	]
}
