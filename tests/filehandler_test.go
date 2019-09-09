package tests
import (
	"FindFirstNonrepeatWord"
	"fmt"
	"testing"
)

func TestCreateLargeFile(t *testing.T){
	FindFirstNonrepeatWord.CreateLargeFile(1024, "TestCreateLargeFile.txt")
}

func TestSplitLargeFile(t *testing.T){
	largeFileName := "TestSplitLargeFile.txt"
	smallFileName := "TestSplitLargeFile_small"
	FindFirstNonrepeatWord.CreateLargeFile(1024 * 1024 * 1024, largeFileName)
	FindFirstNonrepeatWord.SplitLargeFile(10, largeFileName, "", smallFileName, 1024 * 1024 * 1024)
}

func TestBKDRHash(t *testing.T) {
	fmt.Println(FindFirstNonrepeatWord.BKDRHash("1234567890abcdefghijklmnopqrstuvwsyz"))
}