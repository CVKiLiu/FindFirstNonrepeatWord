package FindFirstNonrepeatWord

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const StrLenLimit = 100

func RandStringGenerator() string {
	strRand := rand.New(rand.NewSource(time.Now().Unix()))
	strLen := strRand.Intn(StrLenLimit)
	strByte := make([]byte, strLen)
	for i := 0; i < strLen; i++ {
		strByte[i] = byte('a'+(strRand.Intn(128)%26))
	}
	return string(strByte)
}

func CreateLargeFile(sizeLimit int64, filename string) {
	fileSize := int64(0)
	desFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0600)
	desBufWriter := bufio.NewWriterSize(desFile, 1024 * 1024)
	defer func() {
		err := desBufWriter.Flush()
		if err != nil {
			log.Fatal(err)
		}
		err = desFile.Close()
		if err != nil{
			log.Fatal(err)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}
	for fileSize < sizeLimit {
		curStr := RandStringGenerator()
		_, err := desBufWriter.WriteString(curStr+"\n")
		if err != nil {
			log.Fatal(err)
		}
		fileSize += int64(len(curStr) + 2)
	}
}

func SplitLargeFile(numFile int, srcFilePath string, desFilePath string, filename string, smallFileMem int){

	var oneK = int64(1024)
	var err error
	var smallFile []*os.File
	var bufWriters []*bufio.Writer
	var hashMapNum int64 // amount of all key-value store in hasMap
	var hashMapMem int64 // memory of all key-value store in hashMap
	var hashMapMemLimit int64

	hashMap := make(map[string]int, (oneK * oneK * oneK ) / 5 )  //hashMap can store 0.2G key-value pair without grow

	smallFile = make([]*os.File, numFile)
	bufWriters = make([]*bufio.Writer, numFile)
	for i := 0; i < numFile; i++{
		smallFile[i], err = os.OpenFile(filename + "_"+ strconv.Itoa(i)+".txt", os.O_CREATE|os.O_APPEND, 0777)
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
	for srcScanner.Scan() {
		curStr := srcScanner.Text()
		hashcode := BKDRHash(curStr)
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
			_, err = bufWriters[writeTo].WriteString(curStr + "," + strconv.FormatInt(index, 10) )
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

func BKDRHash(str string) uint64 {
	var hashcode uint64 = 0
	var seed uint64 = 131
	for i := 0; i < len(str); i++ {
		hashcode = hashcode * seed + uint64(str[i])
	}
	return hashcode
}