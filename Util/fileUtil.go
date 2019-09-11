package Util

import (
	"bufio"
	"crypto/rand"
	"log"
	"math/big"
	"os"
	"path/filepath"
)

const (
	StrLenLimit = int64(100)
	constString = "abcdefghizklmnopqrstuvwsyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890~!@#$%^&*()"
)

func PathIsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil{
		if os.IsExist(err){
			return true
		}
		if os.IsNotExist(err){
			return false
		}
		return false
	}
	return true
}

func RandStringGenerator() string {
	// strRand := rand.New(rand.NewSource(time.Now().Unix()))
	strLen, _ := rand.Int(rand.Reader, big.NewInt(StrLenLimit))
	strByte := make([]byte, strLen.Int64())
	for i := 0; i < int(strLen.Int64()); i++ {
		randIdx, _ := rand.Int(rand.Reader, big.NewInt(73))
		strByte[i] = constString[randIdx.Int64()]
	}
	return string(strByte)
}

func CreateLargeFile(sizeLimit uint64, path string, filename string) {

	fileSize := uint64(0)
	if !PathIsExist(path){
		mkDirErr := os.Mkdir(path, 0711)
		if mkDirErr != nil {
			log.Fatal(mkDirErr)
		}
	}
	desFile, err := os.OpenFile(path + "\\" + filename, os.O_APPEND|os.O_CREATE, 0600)
	desBufWriter := bufio.NewWriterSize(desFile, 1024*1024)
	defer func() {
		err := desBufWriter.Flush()
		if err != nil {
			log.Fatal(err)
		}
		err = desFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}
	for fileSize < sizeLimit {
		curStr := RandStringGenerator()
		_, err := desBufWriter.WriteString(curStr + "\n")
		if err != nil {
			log.Fatal(err)
		}
		fileSize += uint64(len(curStr) + 2)
	}
	flushErr := desBufWriter.Flush()
	if flushErr != nil {
		log.Fatal(err)
	}
}

func GetFileSize(filename string) uint64 {
	var result uint64
	err := filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = uint64(f.Size())
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return result
}
