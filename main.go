package main
import (
	"FindFirstNonrepeatWord/Util"
	"FindFirstNonrepeatWord/findx"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)
const (
	oneK = uint64(1024)
)
func main() {
	var numFile uint64
	var srcFileName string
	var desFilePath string
	var filename string
	var smallFileMem uint64
	var path string

	numFile = 42
	path = "text"
	srcFileName = "sroucetext.txt"
	srcFilePath := path +"\\" + srcFileName
	filename = "smallFile"
	desFilePath = "smallText"
	smallFileMem = (5 * oneK * oneK) >> 2

	Util.CreateLargeFile(oneK, path, srcFileName)
	findx.SplitLargeFile(numFile, srcFilePath, desFilePath, filename)
	for i := 0; uint64(i) < numFile; i++ {
		smallFileName := desFilePath + "\\" + filename + "_" + strconv.Itoa(i) + ".txt"
		cutCount := findx.SplitOverSizeSmallFile(smallFileName, smallFileMem)
		numFile += cutCount - 1
	}
	wordInfos := make([]*findx.WordInfo, numFile)
	minIdx := findx.INT64_MAX
	var minWordInfo *findx.WordInfo
	var firstStr string
	dirList, e := ioutil.ReadDir(desFilePath)
	if e != nil{
		log.Fatal(e)
	}
	for i, smallFilename := range dirList {
		var str string
		str, wordInfos[i] = findx.FindX(desFilePath + "\\" + smallFilename.Name())
		if minIdx > wordInfos[i].Idx {
			minIdx = wordInfos[i].Idx
			minWordInfo = wordInfos[i]
			firstStr = str
		}
	}
	fmt.Println(firstStr)
	fmt.Println(minWordInfo.Idx)
	fmt.Println(minWordInfo.Freq)
}
