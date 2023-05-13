package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

func eraseFile(path string, op option) error {
	err := randomizeFile(path, op.interactive)
	if err != nil {
		return err
	}
	if op.keep {
		return nil
	}
	err = os.Remove(path)
	if err != nil {
		return catPathAndErr(path, "remove error", err)
	}
	return nil
}

func eraseDir(path string, op option, errWriter io.Writer) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	var errOccurred atomic.Bool
	errOccurred.Store(false)

	for _, entry := range entries {
		if entry.IsDir() {
			wg.Add(1)
			go func(path string, op option) {
				defer wg.Done()
				if eraseDir(path, op, errWriter) != nil {
					errOccurred.CompareAndSwap(false, true)
				}
			}(filepath.Join(path, entry.Name()), op)
			continue
		}
		wg.Add(1)
		go func(path string, op option) {
			defer wg.Done()
			if err := eraseFile(path, op); err != nil {
				errOccurred.CompareAndSwap(false, true)
				fmt.Fprintln(errWriter, err)
			}
		}(filepath.Join(path, entry.Name()), op)
	}

	wg.Wait()
	if errOccurred.Load() {
		return errors.New("error occurred")
	}
	if op.keep {
		return nil
	}
	return os.Remove(path)
}
