// filescaner project filescaner.go
package filescaner

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"math"
	"os"
	"path/filepath"
)

const (
	MEGABYTE  = 1024 * 1024
	filechunk = 8192
)

type fInfo struct {
	path string
	hsum string
}


func readFiles(dirPath string, readOnlyGhtmlFiles bool) ([]fInfo, []error) {
	var data []fInfo
	var errors []error
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err == nil {

			isFile := false
			if readOnlyGhtmlFiles {
				isFile = !info.IsDir() && filepath.Ext(info.Name()) == ".ghtml"
			}else{
				isFile = !info.IsDir() 
			}

			if isFile {

				fileinfo, err := checkSum(path)
				if err != nil {
					errors = append(errors, err)
					return nil
				}
				data = append(data, fileinfo)
			}
			return nil
		}
		errors = append(errors, err)
		return nil
	})
	return data, errors
}

func checkSum(filepath string) (fInfo, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return fInfo{}, err
	}
	info, _ := file.Stat()
	filesize := info.Size()
	blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))
	hash := md5.New()
	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(filesize-int64(i*filechunk))))
		buf := make([]byte, blocksize)
		file.Read(buf)
		io.WriteString(hash, string(buf))
	}
	fileinfo := fInfo{path: file.Name(), hsum: hex.EncodeToString(hash.Sum(nil))}
	file.Close()
	return fileinfo, nil
}

func ScanFS(path string) (map[string]string, []error) {
	mapOfInfo := make(map[string]string)
	r, err := readFiles(path,false)
	for _, v := range r {
		mapOfInfo[v.path] = v.hsum
	}
	return mapOfInfo, err

}

func ScanGhtmlFilesOnly(path string) (map[string]string, []error) {
	mapOfInfo := make(map[string]string)
	r, err := readFiles(path, true)
	for _, v := range r {
		mapOfInfo[v.path] = v.hsum
	}
	return mapOfInfo, err

}
