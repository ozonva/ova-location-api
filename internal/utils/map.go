package utils

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
