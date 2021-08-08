package main

import (
	"fmt"
	"github.com/ozonva/ova-location-api/internal/utils"
)

func main() {
	fmt.Println("Разбиваем слайс на слайсы")
	{
		sourceSlice := []string{"a", "b", "c", "d", "e"}
		fmt.Println(sourceSlice)
		fmt.Println(utils.SliceSplit(sourceSlice, 2))
	}

	fmt.Println("Фильтруем слайс по слайсу")
	{
		sourceSlice := []string{"a", "b", "c", "d", "e"}
		fmt.Println(sourceSlice)
		excludeSlice := []string{"b", "d"}
		fmt.Println(excludeSlice)
		fmt.Println(utils.SliceFilter(sourceSlice, excludeSlice))
	}

	fmt.Println("Меняем в мапе местами ключ и значение")
	{
		sourceMap := map[string]string{
			"BCN": "Barcelona",
			"LED": "Saint-Petersburg",
			"MSK": "Moscow",
		}
		fmt.Println(sourceMap)
		fmt.Println(utils.MapFlip(sourceMap))
	}
}
