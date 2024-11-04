package inheritanceSolution

import "strings"

type caseInsensitiveReader struct {
	defaultReader
}

// Get : This method is overriding the default Get method to provide case-insensitive search of the key
func (c caseInsensitiveReader) Get(key string) string {
	for k, v := range c.store {
		if strings.EqualFold(k, key) {
			return v
		}
	}
	return ""
}
