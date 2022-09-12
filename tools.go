package main

func getJSON() string {
	where := map[string]interface{}{
		"condition": []interface{}{
			map[string]interface{}{
				"key":    "key1",
				"action": "eq",
				"value":  19,
			},
			map[string]interface{}{
				"key":    "key2",
				"action": "between",
				"value": []string{
					"2021-02-01 01:01:01",
					"2021-02-05 01:01:01",
				},
			},
			map[string]interface{}{
				"key":    "key3",
				"action": "lt",
				"value":  19,
			},
		},
		"sort": "key1 desc",
		"or": []interface{}{
			[]interface{}{
				map[string]interface{}{
					"key":    "key1",
					"action": "eq",
					"value":  19,
				},
				map[string]interface{}{
					"key":    "key2",
					"action": "eq",
					"value":  19,
				},
			},
			[]interface{}{
				map[string]interface{}{
					"key":    "key3",
					"action": "eq",
					"value":  19,
				},
				map[string]interface{}{
					"key":    "key4",
					"action": "eq",
					"value":  19,
				},
			},
		},
	}
	s, _ := json.Marshal(where)
	return string(s)
}