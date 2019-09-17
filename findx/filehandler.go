package findx

import (
	sutil "FindFirstNonrepeatWord/Util"
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

// SplitLargeFile split large size file
// numFile: split large file into numFile
// srcFilePath: source file path
// desFilePath: destination file path
// filename: small file prefix
func SplitLargeFile(numFile uint64, srcFilePath string, desFilePath string, filename string) uint64 {

	var err error
	var srcScanner *bufio.Scanner
	var index int64 // global Increasing index use to record the line number of every string
	var smallFile []*os.File
	var bufWriters []*bufio.Writer
	var hashMap map[string]int
	var hashMapNum int64 // amount of all key-value store in hasMap
	var hashMapMem int64 // memory of all key-value store in hashMap
	var hashMapMemLimit int64

	// use hashmap to reduce duplication string written into small file
	// set capacity of hashMap to 0.2G, so hashMap can store 0.2G key-value pair without grow
	hashMap = make(map[string]int, (oneK*oneK*oneK)/5)
	smallFile = make([]*os.File, numFile)
	bufWriters = make([]*bufio.Writer, numFile)
	if !sutil.PathIsExist(desFilePath) {
		mkDirErr := os.Mkdir(desFilePath, 0777)
		if mkDirErr != nil {
			log.Fatal(mkDirErr)
		}
	}
	for i := 0; uint64(i) < numFile; i++ {
		smallFile[i], err = os.OpenFile(desFilePath+"\\"+filename+"_"+strconv.Itoa(i)+".txt", os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			log.Fatal(err)
		}
		// set writer buffer to 1M
		bufWriters[i] = bufio.NewWriterSize(smallFile[i], 1*1024*1024)
	}

	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := srcFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	srcScanner = bufio.NewScanner(srcFile)
	index = int64(1)
	hashMapMem = int64(0)
	hashMapMemLimit = int64(oneK * oneK * oneK * 14) // set the max size of hashmap in memory to 14G
	hashMapNum = int64(0)
	for srcScanner.Scan() {
		curStr := srcScanner.Text()
		hashcode := sutil.BKDRHash(curStr)
		_, ok := hashMap[curStr]
		// the size of hashmap equal to all key-value pair size plus the the amount of key-value pair (hashMapMem+hashMapNum)
		// how to infer this equation?
		// see the struct of map in rumtime/map.go
		if !ok && (hashMapMem+hashMapNum) < hashMapMemLimit {
			hashMap[curStr] = 1
			hashMapMem += int64(len(curStr)) + int64(8)
			hashMapNum++
		} else if ok {
			hashMap[curStr]++
		}

		// if the frequency is smaller than _two_(else discard it cause two same string with smaller index be written into corresponding small file)
		// or
		// hashMap is oversize
		// write this string to corresponding small file.
		if times, ok := hashMap[curStr]; times <= 2 || !ok && (hashMapMem+hashMapNum) > hashMapMemLimit {
			writeTo := hashcode % uint64(numFile)
			_, err = bufWriters[writeTo].WriteString(curStr + "," + strconv.FormatInt(index, 10) + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
		index++
	}
	scanErr := srcScanner.Err()
	if scanErr != nil {
		log.Fatal(err)
	}
	for i := 0; uint64(i) < numFile; i++ {
		err = bufWriters[i].Flush()
		if err != nil {
			log.Fatal(err)
		}
	}
	return numFile
}

// SplitOverSizeSmallFile re-split small size file if it's oversize
func SplitOverSizeSmallFile(filename string, smallFileMemLimit uint64) uint64 {
	var splitNum uint64 = 1
	var nanoFile []*os.File
	var bufWriters []*bufio.Writer
	fileSuffix := filepath.Ext(filename)
	filenameOnly := strings.TrimSuffix(filename, fileSuffix)
	smallFileSize := sutil.GetFileSize(filename)

	if smallFileSize > smallFileMemLimit {
		splitNum = (smallFileSize / smallFileMemLimit) + 1
		nanoFile = make([]*os.File, splitNum)
		bufWriters = make([]*bufio.Writer, splitNum)
		for i := 0; uint64(i) < splitNum; i++ {
			var nanoFileErr error
			nanoFilename := filenameOnly + "_" + strconv.Itoa(i) + ".txt"
			nanoFile[i], nanoFileErr = os.OpenFile(nanoFilename, os.O_CREATE|os.O_APPEND, 0777)
			if nanoFileErr != nil {
				log.Fatal(nanoFileErr)
			}
			bufWriters[i] = bufio.NewWriterSize(nanoFile[i], int(oneK*oneK))
		}
		smallFile, err := os.OpenFile(filename, os.O_RDONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			for i := 0; uint64(i) < splitNum; i++ {
				flushErr := bufWriters[i].Flush()
				if flushErr != nil {
					log.Fatal(flushErr)
				}
				nanoFileCloseErr := nanoFile[i].Close()
				if nanoFileCloseErr != nil {
					log.Fatal(nanoFileCloseErr)
				}
			}

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
			hashcode := sutil.BKDRHash(textAndIdx[0])
			hashNum := hashcode % splitNum
			_, err := bufWriters[hashNum].WriteString(curStr + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return splitNum
}
