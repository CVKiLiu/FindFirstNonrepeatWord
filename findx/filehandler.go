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
const (
	oneK = int64(1024)
)
func SplitLargeFile(numFile uint64, srcFilePath string, desFilePath string, filename string) uint64{

	var err error
	var smallFile []*os.File
	var bufWriters []*bufio.Writer
	var hashMapNum int64 // amount of all key-value store in hasMap
	var hashMapMem int64 // memory of all key-value store in hashMap
	var hashMapMemLimit int64

	smallFile = make([]*os.File, numFile)
	bufWriters = make([]*bufio.Writer, numFile)
	if !Util.PathIsExist(desFilePath) {
		mkDirErr := os.Mkdir(desFilePath, 0711)
		if mkDirErr != nil {
			log.Fatal(mkDirErr)
		}
	}
	for i := 0; uint64(i) < numFile; i++{
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
	for i:=0; uint64(i) < numFile; i++ {
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

	return numFile
}

func SplitOverSizeSmallFile(filename string, smallFileMemLimit uint64) uint64{
	var splitNum uint64 = 1
	var nanoFile []*os.File
	var bufWriters []*bufio.Writer
	fileSuffix := filepath.Ext(filename)
	filenameOnly := strings.TrimSuffix(filename, fileSuffix)
	smallFileSize := Util.GetFileSize(filename)

	if smallFileSize > smallFileMemLimit {
		splitNum = (smallFileSize / smallFileMemLimit) + 1
		nanoFile = make([]*os.File, splitNum)
		bufWriters = make([]*bufio.Writer, splitNum)
		for i := 0 ; uint64(i) < splitNum; i++ {
			var nanoFileErr error
			nanoFilename := filenameOnly+ "_" + strconv.Itoa(i) + ".txt"
			nanoFile[i], nanoFileErr = os.OpenFile(nanoFilename, os.O_CREATE | os.O_APPEND, 0777)
			if nanoFileErr != nil{
				log.Fatal(nanoFileErr)
			}
			bufWriters[i] = bufio.NewWriterSize(nanoFile[i], int(oneK * oneK))
		}
		defer func() {
			for i := 0 ; uint64(i) < splitNum; i++ {
				flushErr := bufWriters[i].Flush()
				if flushErr != nil{
					log.Fatal(flushErr)
				}
				closeErr := nanoFile[i].Close()
				if closeErr != nil {
					log.Fatal(closeErr)
				}
			}
		}()

		smallFile, err := os.OpenFile(filename, os.O_RDONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			closeErr := smallFile.Close()
			if closeErr != nil {
				log.Fatal(closeErr)
			}
			removeSrcErr := os.Remove(filename)
			if removeSrcErr != nil {
				log.Fatal(removeSrcErr)
			}
		}()

		smallFileScanner := bufio.NewScanner(smallFile)
		for smallFileScanner.Scan() {
			curStr := smallFileScanner.Text()
			textAndIdx := strings.Split(curStr, ",")
			hashcode := Util.BKDRHash(textAndIdx[0])
			hashNum := hashcode % splitNum
			_, err := bufWriters[hashNum].WriteString(curStr+"\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return splitNum
}