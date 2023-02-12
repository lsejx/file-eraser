package main

import (
	"errors"
	"io"
	"os"

	rand "example.com/me/hashed_random"
)

func randomize(dst io.Writer, length int64) error {
	buf := make([]byte, length)

	randomSrc, err := rand.NewHashedRandomReader()
	if err != nil {
		return err
	}
	n, err := io.ReadFull(randomSrc, buf)
	if err != nil {
		return err
	}
	if n != len(buf) {
		return errors.New("failed to read enough from random source")
	}

	n, err = dst.Write(buf)
	if err != nil {
		return err
	}
	if int64(n) != length {
		return errors.New("failed to fully randomize the file")
	}

	return nil
}

func randomizeFile(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY, 0600)
	if err != nil {
		return catPathAndErr(path, err)
	}

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
