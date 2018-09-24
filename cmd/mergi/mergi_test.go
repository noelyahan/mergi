package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetJobMap(t *testing.T) {
	s := "0#resize 0#crop 0#watermark 1#crop 3#crop 3#resize 4#final 4#crop 4#resize"
	arr := strings.Split(s, " ")
	m := getJobMap(arr)
	for k, v := range m {
		fmt.Println(k, v)
	}
}
