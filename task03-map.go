package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var keys []int
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	resul := make([]string, 0, len(input))
	for _, k := range keys {
		resul = append(resul, input[k])
	}
	return resul
}
