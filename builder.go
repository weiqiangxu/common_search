package main

import "fmt"

func Builder(w *Where) string {
	if w == nil {
		return ""
	}
	var queryList []string
	for _, v := range w.Condition {
		tmp := ConditionHandler(&v, setBetweenHandler(), setEQHandler())
		if tmp != "" {
			queryList = append(queryList, tmp)
		}
	}
	var SQLString string
	if len(queryList) != 0 {
		SQLString = "1"
		for _, v := range queryList {
			SQLString += fmt.Sprintf(" and %s", v)
		}
	}
	return SQLString
}
