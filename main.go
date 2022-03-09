package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

var m = make(map[string]int)

func walk(s string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !d.IsDir() {
		f, err := os.Open(s)
		if err != nil {
			log.Fatal(err)
		}

		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}
		func() {
			f.Close()
		}()
		hsum := hex.EncodeToString(h.Sum(nil))
		if _, in := m[hsum]; in {
			err := os.Remove(s)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			m[hsum] = 1
		}
	}
	return nil
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	filepath.WalkDir(path, walk)
}
