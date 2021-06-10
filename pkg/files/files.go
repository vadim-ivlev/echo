package files

import (
	"io/ioutil"
	"os"
)

func ListFiles() ([]string, error) {
	ss := make([]string, 0)
	dir, _ := os.Getwd()
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		ss = append(ss, file.Name())
	}
	return ss, nil
}

func ReadTextFile(fileName string) (string, error) {
	dat, err := ioutil.ReadFile("txt.txt")
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
