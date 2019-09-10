package tests

import (
	"FindFirstNonrepeatWord/findx"
	"fmt"
	"testing"
)

func TestWordInfoHeapPop(t *testing.T) {
	filename := "textGenerate/TestSplitLargeFile_small_5.txt"
	firstX := findx.FindX(filename)
	fmt.Println(firstX)
}
