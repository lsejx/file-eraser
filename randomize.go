package main

import (
	"crypto/rand"
	"io"
	"os"
)

func randomize(dst io.Writer, length int64) error {
	buf := make([]byte, length)

	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		return err
	}

	_, err = dst.Write(buf)
	if err != nil {
		return err
	}

	return nil
}

func randomizeFile(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY, 0600)
	if err != nil {
		return catPathAndErr(path, err)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return catPathAndErr(path, err)
	}

	err = randomize(f, stat.Size())
	if err != nil {
		return catPathAndErr(path, err)
	}
	return nil
}
