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

func eraseFile(path string, interactive bool) error {
	if interactive {
		yes, err := interacter.ask(path)
		if err != nil {
			return err
		}
		if !yes {
			return nil
		}
	}
	err := randomizeFile(path)
	if err != nil {
		return err
	}
	err = os.Remove(path)
	if err != nil {
		return fmt.Errorf("%v: failed to remove: %v", path, err)
	}
	return nil
}

func eraseDir(path string, interactive bool, errWriter io.Writer) error {
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
			go func(path string) {
				defer wg.Done()
				if eraseDir(path, interactive, errWriter) != nil {
					errOccurred.CompareAndSwap(false, true)
				}
			}(filepath.Join(path, entry.Name()))
			continue
		}
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			err := eraseFile(path, interactive)
			if err != nil {
				errOccurred.CompareAndSwap(false, true)
				fmt.Fprintf(errWriter, "error: %v\n", err)
			}
		}(filepath.Join(path, entry.Name()))
	}

	wg.Wait()
	if errOccurred.Load() {
		return errors.New("error occured")
	}
	return os.Remove(path)
}
