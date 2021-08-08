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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
