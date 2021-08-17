package utils

import (
	"github.com/ozonva/ova-location-api/internal/location"
	"io/ioutil"
	"os"
)

func LocationSliceSplit(sourceSlice []location.Location, chunkSize int) [][]location.Location {
	if chunkSize <= 0 {
		return nil
	}
	chunksCount := (len(sourceSlice) + chunkSize - 1) / chunkSize
	resultChunks := make([][]location.Location, chunksCount)
	for i := 0; i < chunksCount; i++ {
		first := i * chunkSize
		last := min(first+chunkSize, len(sourceSlice))
		resultChunks[i] = sourceSlice[first:last]
	}
	return resultChunks
}

func LocationSliceToMap(sourceSlice []location.Location) map[uint64]location.Location {
	resultMap := make(map[uint64]location.Location, len(sourceSlice))
	for _, model := range sourceSlice {
		if _, found := resultMap[model.Id]; found {
			continue
		}
		resultMap[model.Id] = model
	}
	return resultMap
}

func ReadFiles(paths []string) ([]string, error) {
	reader := func(path string) ([]byte, error) {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	result := make([]string, 0, len(paths))
	for _, path := range paths {
		data, err := reader(path)
		if err != nil {
			return nil, err
		}
		result = append(result, string(data))
	}
	return result, nil
}

func SliceSplit(sourceSlice []string, chunkSize int) [][]string {
	if chunkSize <= 0 {
		return nil
	}
	chunksCount := (len(sourceSlice) + chunkSize - 1) / chunkSize
	resultChunks := make([][]string, chunksCount)
	for i := 0; i < chunksCount; i++ {
		first := i * chunkSize
		last := min(first+chunkSize, len(sourceSlice))
		resultChunks[i] = sourceSlice[first:last]
	}
	return resultChunks
}

func SliceFilter(sourceSlice []string, excludeSlice []string) []string {
	excludeMap := convertSliceToMap(excludeSlice)
	var resultSlice []string
	for _, element := range sourceSlice {
		if _, found := excludeMap[element]; !found {
			resultSlice = append(resultSlice, element)
		}
	}
	return resultSlice
}

func MapFlip(sourceMap map[string]string) map[string]string {
	resultMap := make(map[string]string, len(sourceMap))
	for key, value := range sourceMap {
		if _, found := resultMap[value]; found {
			continue
		}
		resultMap[value] = key
	}
	return resultMap
}

func convertSliceToMap(sourceSlice []string) map[string]string {
	resultMap := map[string]string{}
	for _, element := range sourceSlice {
		resultMap[element] = element
	}
	return resultMap
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
