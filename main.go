package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func ValidNewName(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		// We want a NotExist
		if os.IsNotExist(err) {
			return true, nil
		}
	}

	return false, err
}

func NewName(path string) (string, error) {
	fInfo, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	datestring := time.Now().Format("20060102")
	newname := fmt.Sprintf("%s-%s", fInfo.Name(), datestring)

	d := filepath.Dir(path)
	newPath := filepath.Join(d, newname)

	v, err := ValidNewName(newPath)
	if err != nil {
		return "", err
	}

	if v {
		return newPath, nil
	}

	// At this point, another yyyymmdd Oldfile exists,
	// so append n until we have a new filename
	n := 0
	for {
		d := filepath.Dir(path)
		newNumName := fmt.Sprintf("%s-%d", newname, n)
		newPath := filepath.Join(d, newNumName)
		v, err := ValidNewName(newPath)
		if err != nil {
			return "", err
		}

		if v {
			return newPath, nil
		}

		n++
	}
}

func Rename(path string) error {
	newPath, err := NewName(path)
	if err != nil {
		return err
	}

	fmt.Printf("moving %s to %s\n", path, newPath)
	return os.Rename(path, newPath)
}

func main() {
	flag.Parse()

	for _, file := range flag.Args() {
		err := Rename(file)
		if err != nil {
			log.Print(err)
		}
	}

}
