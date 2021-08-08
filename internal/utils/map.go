package utils

import "fmt"

func MapFlip(sourceMap map[string]string) map[string]string {
	resultMap := make(map[string]string, len(sourceMap))
	for key, value := range sourceMap {
		if _, found := resultMap[value]; found {
			panic(fmt.Sprintf("Пара с таким ключом уже присутствует в словаре: %v", value))
		}
		resultMap[value] = key
	}
	return resultMap
}
