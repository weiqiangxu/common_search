package main

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
)

type ConditionActionEnum string

type ConditionItemValue interface {
}

const (
	Between ConditionActionEnum = "between"
	In      ConditionActionEnum = "in"
	EQ      ConditionActionEnum = "eq"
	LT      ConditionActionEnum = "lt"
	GT      ConditionActionEnum = "gt"
)

type ConditionItem struct {
	Key    string              `json:"key"`
	Action ConditionActionEnum `json:"action"`
	Value  ConditionItemValue  `json:"value"`
}

type Where struct {
	Condition []ConditionItem
	Sort      string
	OR        [][]ConditionItem
}

type ConditionHandlerFunc func(c *ConditionItem) string

func ConditionHandler(c *ConditionItem, opt ...ConditionHandlerFunc) (SQLItem string) {
	for _, v := range opt {
		tmp := v(c)
		if tmp != "" {
			SQLItem = tmp
		}
	}
	return
}

func setBetweenHandler() ConditionHandlerFunc {
	return func(c *ConditionItem) (SQL string) {
		if c.Action == Between {
			l := gconv.SliceStr(c.Value)
			if len(l) >= 2 {
				SQL = fmt.Sprintf("%s between (%s,%s)", c.Key, l[0], l[1])
			}
		}
		return SQL
	}
}

func setEQHandler() ConditionHandlerFunc {
	return func(c *ConditionItem) (SQL string) {
		if c.Action == EQ {
			SQL = fmt.Sprintf("%s = %s", c.Key, gconv.String(c.Value))
		}
		return
	}
}
