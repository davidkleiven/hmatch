package pkg

func groupby[K comparable, E any](s []E, keyFunc func(E) K) map[K][]E {
	result := make(map[K][]E)

	for _, v := range s {
		key := keyFunc(v)
		result[key] = append(result[key], v)
	}
	return result
}
