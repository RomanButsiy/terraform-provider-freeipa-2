package main

import (
	"strings"
)

func splitID(id string) (a, b, c string) {
	resource_map := make([]string, 3)
	copy(resource_map, strings.Split(id, "/"))
	return resource_map[0], resource_map[1], resource_map[2]
}

func utilsGetArry(itemsRaw []interface{}) []string {
	res := make([]string, len(itemsRaw))
	for i, raw := range itemsRaw {
		res[i] = raw.(string)
	}
	return res
}
