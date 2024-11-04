package comparisonStrategySolution

type DefaultReader struct {
	store             map[string]string
	defaultComparator Comparator
}

// NewReader returns a new DefaultReader with a default comparator.
// The default comparator does a case-sensitive comparison of strings.
func NewReader() DefaultReader {
	c := NewComparator()
	store := getStore()
	return DefaultReader{
		store:             store,
		defaultComparator: c,
	}
}

// WithComparator returns a new DefaultReader with the given comparator.
// The comparator is used in the Get method to compare the given key with the
// keys in the store.
func (d DefaultReader) WithComparator(comparator Comparator) DefaultReader {
	d.defaultComparator = comparator
	return d
}

func (d DefaultReader) Get(key string) string {
	// Before
	//return d.store[key]

	for k, v := range d.store {
		if d.defaultComparator.Compare(k, key) {
			return v
		}
	}

	return ""
}

func getStore() map[string]string {
	store := make(map[string]string)
	store["hello"] = "world"
	return store
}
