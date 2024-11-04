package inheritanceSolution

type defaultReader struct {
	store map[string]string
}

// Get returns the value associated with the given key in the map.
// If the key is not present in the map, an empty string is returned.
func (d defaultReader) Get(key string) string {
	return d.store[key]
}
