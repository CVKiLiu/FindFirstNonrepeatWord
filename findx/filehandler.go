package findx

import (
	"FindFirstNonrepeatWord/Util"
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func SplitLargeFile(numFile int, srcFilePath string, desFilePath string, filename string, smallFileMem int64){

	var oneK = int64(1024)
	var err error
	var smallFile []*os.File
	var bufWriters []*bufio.Writer
	var hashMapNum int64 // amount of all key-value store in hasMap
	var hashMapMem int64 // memory of all key-value store in hashMap
	var hashMapMemLimit int64

	_, err = os.Stat(desFilePath)
	if os.IsNotExist(err) {
		os.Mkdir(desFilePath, 0711)
	}
	smallFile = make([]*os.File, numFile)
	bufWriters = make([]*bufio.Writer, numFile)
	for i := 0; i < numFile; i++{
		smallFile[i], err = os.OpenFile(desFilePath + "\\" + filename + "_"+ strconv.Itoa(i)+".txt", os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			log.Fatal(err)
		}
		bufWriters[i] = bufio.NewWriterSize(smallFile[i], 1 * 1024 * 1024)
	}

	srcFile, err := os.Open(srcFilePath)
	if err != nil{
		log.Fatal(err)
	}
	defer func(){
		err := srcFile.Close()
		if err != nil{
			log.Fatal(err)
		}
	}()

	srcScanner := bufio.NewScanner(srcFile)
	var index int64
	index = int64(0)
	hashMapMem = int64(0)
	hashMapMemLimit = int64(oneK * oneK * oneK * 14)
	hashMapNum = int64(0)
	hashMap := make(map[string]int, (oneK * oneK * oneK ) / 5 )  //hashMap can store 0.2G key-value pair without grow
	for srcScanner.Scan() {
		curStr := srcScanner.Text()
		hashcode := Util.BKDRHash(curStr)
		_, ok := hashMap[curStr]
		if !ok && (hashMapMem + hashMapNum) < hashMapMemLimit {
			hashMap[curStr] = 1
			hashMapMem += int64(len(curStr)) + int64(8)
			hashMapNum++
		}else if ok {
			hashMap[curStr]++
		}

		if times, ok := hashMap[curStr]; times == 1 || !ok && (hashMapMem+hashMapNum) < hashMapMemLimit {
			writeTo := hashcode % uint64(numFile)
			_, err = bufWriters[writeTo].WriteString(curStr + "," + strconv.FormatInt(index, 10) + "\n" )
			if err != nil {
				log.Fatal(err)
			}
		}
		index++
	}
	for i:=0; i < numFile; i++ {
		err = bufWriters[i].Flush()
		if err != nil {
			log.Fatal(err)
		}
	}
	defer func() {
		err := srcScanner.Err()
		if err != nil{
			log.Fatal(err)
		}
	}()
}

func splitOverSizeSmallFile(filename string, smallFileMemLimit int64) int64{
	var splitNum int64 = 1
	fileSuffix := filepath.Ext(filename)
	filenameOnly := strings.TrimSuffix(filename, fileSuffix)
	smallFileSize := Util.GetFileSize(filename)
	if smallFileSize > smallFileMemLimit {
		splitNum = (smallFileSize % smallFileMemLimit) + 1
		smallFile, err := os.OpenFile(filename, os.O_RDONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}
		smallFileScanner :=
	}
	return splitNum
}