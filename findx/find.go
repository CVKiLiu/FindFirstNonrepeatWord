package findx

import (
	"bufio"
	//"container/heap"
	"log"
	"os"
	"strconv"
	"strings"
)

// Int64Max The max value of int64
const Int64Max = int64(uint64(1)<<63 - 1)

// FindX return first non-repeated word in text file
func FindX(filename string) (string, *WordInfo) {
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

	// read all string and use hashmap to save their frequence and the first idx
	for srcScanner.Scan() {
		curStr := srcScanner.Text()
		strAndIdx := strings.Split(curStr, ",") // [str, idx]
		idx, _ := strconv.ParseInt(strAndIdx[1], 10, 64)
		wInfo, ok := strToWordInfoMap[strAndIdx[0]]
		if !ok {
			wInfo = NewWordInfo(idx, int64(1))
			strToWordInfoMap[strAndIdx[0]] = wInfo
		} else {
			wInfo.Freq++
		}
	}

	// find the first non-repeated word
	minIdx := Int64Max
	firstStr := ""
	for str, wi := range strToWordInfoMap {
		if wi.Freq == int64(1) && wi.Idx < minIdx {
			minIdx = wi.Idx
			firstNonRepeatWInfo = wi
			firstStr = str
		}
	}

	return firstStr, firstNonRepeatWInfo
}
