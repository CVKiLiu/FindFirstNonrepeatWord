package FindFirstNonrepeatWord

import (
	"os"
	"log"
	"bufio"
	"strconv"
)

func SplitLargeFile(num_file int, srcFilePath string, desFilePath string, filename string, smallFileMem int){

	var oneK = int64(1024)
	var err error
	var smallFile [num_file]*os.File
	var bufWriters [num_file]*bufio.Writer
	hashMap := make(map[string]string, (oneK * oneK * oneK * 3) >> 1)  //apply 1.5G capacity avoid hash grow

	for i := 0; i < num_file; i++{
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
	for srcScanner.Scan(){
		curStr := srcScanner.Text()
		_, ok := hashMap[curStr]
		if !ok {

		}else{

		}
	}
	defer func(){
		err := srcScanner.Err()
		if err != nil{
			log.Fatal(err)
		}
	}()
	
}