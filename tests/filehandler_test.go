package tests

import (
	"FindFirstNonrepeatWord/Util"
	"FindFirstNonrepeatWord/findx"
	"fmt"
	"testing"
)

func TestCreateLargeFile(t *testing.T) {
	Util.CreateLargeFile(1024*1024, "TestCreateLargeFile.txt")
}

func TestSplitLargeFile(t *testing.T) {
	largeFileName := "TestSplitLargeFile.txt"
	smallFileName := "TestSplitLargeFile_small"
	Util.CreateLargeFile(10*1024, largeFileName)
	findx.SplitLargeFile(10, largeFileName, "textGenerate", smallFileName, 1024*1024)
}

func TestBKDRHash(t *testing.T) {
	fmt.Println(Util.BKDRHash("1234567890abcdefghijklmnopqrstuvwsyz"))
	fmt.Println(Util.BKDRHash("1234567890abcdefghijklmnopqrstuvwsyz"))
	fmt.Println(Util.BKDRHash("1234567890abcdefghijklmnopqrstuvwsy_"))
}

func TestGetFileSize(t *testing.T) {
	filename := "TestSplitLargeFile.txt"
	fmt.Println(Util.GetFileSize(filename))
}
