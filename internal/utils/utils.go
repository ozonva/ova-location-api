package utils

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
