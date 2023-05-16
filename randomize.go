package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func randomize(dst io.Writer, length int64) error {
	buf := make([]byte, length)

	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return err
	}

	if _, err := dst.Write(buf); err != nil {
		return err
	}

	return nil
}

func randomizeFile(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("%v: randomization error: %v", path, parsePathErr(err))
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return fmt.Errorf("%v: randomization error: %v", path, parsePathErr(err))
	}

	if err = randomize(f, stat.Size()); err != nil {
		return fmt.Errorf("%v: randomization error: %v", path, err)
	}
	return nil
}
