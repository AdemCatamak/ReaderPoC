package main

import (
	"ReaderPoC/comparisonStrategySolution"
	"ReaderPoC/inheritanceSolution"
	"fmt"
)

func main() {

	readerWithDefaultComparator := comparisonStrategySolution.NewReader()
	readerWithCaseInsensitiveComparator := comparisonStrategySolution.NewReader().WithComparator(comparisonStrategySolution.NewCaseInsensitiveComparator())

	fmt.Printf("readerWithDefaultComparator [hello]: %s\n", readerWithDefaultComparator.Get("hello"))
	fmt.Printf("readerWithDefaultComparator [heLLo]: %s\n", readerWithDefaultComparator.Get("heLLo"))

	fmt.Printf("readerWithCaseInsensitiveComparator [hello]: %s\n", readerWithCaseInsensitiveComparator.Get("hello"))
	fmt.Printf("readerWithCaseInsensitiveComparator [heLLo]: %s\n", readerWithCaseInsensitiveComparator.Get("heLLo"))

	defaultReader := inheritanceSolution.NewReader()
	caseInsensitiveReader := inheritanceSolution.NewCaseInsensitiveReader()

	fmt.Printf("defaultReader [hello]: %s\n", defaultReader.Get("hello"))
	fmt.Printf("defaultReader [heLLo]: %s\n", defaultReader.Get("heLLo"))

	fmt.Printf("caseInsensitiveReader [hello]: %s\n", caseInsensitiveReader.Get("hello"))
	fmt.Printf("caseInsensitiveReader [heLLo]: %s\n", caseInsensitiveReader.Get("heLLo"))

}
