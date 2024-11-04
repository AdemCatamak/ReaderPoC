package inheritanceSolution

// Reader : This is a method signature in Go that defines a method named `Get` which takes a `string` parameter `key` and returns a `string` value.
type Reader interface {
	Get(key string) string
}

func NewReader() Reader {
	store := getStore()
	return defaultReader{
		store: store,
	}
}

func NewCaseInsensitiveReader() Reader {
	store := getStore()
	reader := defaultReader{
		store: store,
	}

	return caseInsensitiveReader{
		defaultReader: reader,
	}
}

func getStore() map[string]string {
	store := make(map[string]string)
	store["hello"] = "world"
	return store
}
