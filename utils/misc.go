package utils

type KVEntry struct {
	Key   string
	Value int
}

// func SortMapByKey(input []KVEntry) []KVEntry {
// 	keys := make([]string, 0, len(input))

// 	for k := range input {
// 		keys = append(keys, k)
// 	}
// 	sort.Strings(keys)

// 	var result map[string]int = make(map[string]int, 0)

// 	for _, k := range keys {
// 		// fmt.Println(k, input[k])
// 		result[k] = input[k]
// 	}

// 	return result
// }
