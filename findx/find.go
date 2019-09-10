package findx

import (
	"bufio"
	//"container/heap"
	"log"
	"os"
	"strconv"
	"strings"
)

const INT64_MAX = int64(uint64(1)<<63 - 1)

// FindX return first non-repeated word in text file
func FindX(filename string) *WordInfo {
	var firstNonRepeatWInfo *WordInfo
	strToWordInfoMap := make(map[string]*WordInfo)
	//wh := NewWordInfoHeap()
	srcFile, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	srcScanner := bufio.NewScanner(srcFile)
	defer func() {
		err = srcScanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		err = srcFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for srcScanner.Scan() {
		curStr := srcScanner.Text()
		strAndIdx := strings.Split(curStr, ",") // [str, idx]
		idx, _ := strconv.ParseInt(strAndIdx[1], 10, 64)
		wInfo, ok := strToWordInfoMap[strAndIdx[0]]
		if !ok {
			wInfo = NewWordInfo(strAndIdx[0], idx, int64(1))
			strToWordInfoMap[strAndIdx[0]] = wInfo
		} else {
			wInfo.Freq++
		}
		//heap.Push(&wh, wInfo)
	}
	minIdx := INT64_MAX
	for _, wi := range strToWordInfoMap {
		if wi.Freq == int64(1) && wi.Idx < minIdx {
			minIdx = wi.Idx
			firstNonRepeatWInfo = wi
		}
	}
	// topWord := heap.Pop(&wh)

	return firstNonRepeatWInfo
}
