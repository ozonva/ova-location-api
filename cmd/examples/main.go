package main

import (
	"fmt"
	"github.com/ozonva/ova-location-api/internal/location"
	"github.com/ozonva/ova-location-api/internal/utils"
	"time"
)

func main() {
	fmt.Println("Читаем файлы")
	{
		paths := []string{
			"/Users/optimusfrai/projects/ozon/ova-location-api/cmd/examples/.config",
		}
		fmt.Println(paths)
		data, _ := utils.ReadFiles(paths)
		fmt.Println(data)
	}
	fmt.Println("Разбиваем слайс на слайсы")
	{
		sourceSlice := []string{"a", "b", "c", "d", "e"}
		fmt.Println(sourceSlice)
		fmt.Println(utils.SliceSplit(sourceSlice, 2))
	}

	fmt.Println("Разбиваем слайс локаций на слайсы")
	{
		sourceSlice := []location.Location{
			{1, 1, "Some address 1", 20, 30, time.Now()},
			{2, 2, "Some address 2", 22, 34, time.Now()},
		}
		fmt.Println(sourceSlice)
		fmt.Println(utils.LocationSliceSplit(sourceSlice, 1))
	}

	fmt.Println("Превращаем слайс локаций в мапу по Id")
	{
		sourceSlice := []location.Location{
			{1, 1, "Some address 1", 20, 30, time.Now()},
			{2, 2, "Some address 2", 22, 34, time.Now()},
		}
		fmt.Println(sourceSlice)
		fmt.Println(utils.LocationSliceToMap(sourceSlice))
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
