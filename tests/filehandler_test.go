package tests

import (
	sutil "FindFirstNonrepeatWord/Util"
	"FindFirstNonrepeatWord/findx"
	"fmt"
	"testing"
)

func TestCreateLargeFile(t *testing.T) {
	path := "text"
	sutil.CreateLargeFile(1024*1024, path, "TestCreateLargeFile.txt")
}

func TestSplitLargeFile(t *testing.T) {
	path := "text"
	largeFileName := "TestSplitLargeFile.txt"
	smallFileName := "TestSplitLargeFile_small"
	sutil.CreateLargeFile(10*1024, path, largeFileName)
	findx.SplitLargeFile(uint64(10), largeFileName, "textGenerate", smallFileName)
}

func TestSplitOverSizeSmallFile(t *testing.T) {
	path := "text"
	overSizeFile := "oversizeSmallFile.txt"
	sutil.CreateLargeFile(10*1024, path, overSizeFile)
	findx.SplitOverSizeSmallFile(overSizeFile, 5*1024)
}

func TestBKDRHash(t *testing.T) {
	fmt.Println(sutil.BKDRHash("1234567890abcdefghijklmnopqrstuvwsyz"))
	fmt.Println(sutil.BKDRHash("1234567890abcdefghijklmnopqrstuvwsyz"))
	fmt.Println(sutil.BKDRHash("1234567890abcdefghijklmnopqrstuvwsy_"))
}

func TestGetFileSize(t *testing.T) {
	filename := "TestSplitLargeFile.txt"
	fmt.Println(sutil.GetFileSize(filename))
}
