package comparisonStrategySolution

import "strings"

type Comparator interface {
	Compare(key1 string, key2 string) bool
}

func NewComparator() Comparator {
	return defaultComparator{}
}

func NewCaseInsensitiveComparator() Comparator {
	return caseInsensitiveComparator{}
}

type defaultComparator struct {
}

func (d defaultComparator) Compare(key1 string, key2 string) bool {
	return key1 == key2
}

type caseInsensitiveComparator struct {
}

func (c caseInsensitiveComparator) Compare(key1 string, key2 string) bool {
	return strings.EqualFold(key1, key2)
}
