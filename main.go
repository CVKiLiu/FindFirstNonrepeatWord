package FindFirstNonrepeatWord

import (
	"FindFirstNonrepeatWord/findx"
	"fmt"
	"strconv"
)

func main() {
	var numFile int
	var srcFilePath string
	var desFilePath string
	var filename string
	var smallFileMem int64
	findx.CreateLargeFile(1024, srcFilePath)
	findx.SplitLargeFile(numFile, srcFilePath, desFilePath, filename, smallFileMem)
	wordInfos := make([]*findx.WordInfo, numFile)
	minIdx := findx.INT64_MAX
	var minWordInfo *findx.WordInfo
	for i := 0; i < numFile; i++ {
		smallFileName := desFilePath + "\\" + desFilePath + "\\" + filename + "_" + strconv.Itoa(i) + ".txt"
		wordInfos[i] = findx.FindX(smallFileName)
		if minIdx > wordInfos[i].Idx {
			minIdx = wordInfos[i].Idx
			minWordInfo = wordInfos[i]
		}
	}
	fmt.Println(minWordInfo.Word)
}
