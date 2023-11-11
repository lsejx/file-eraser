package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"

	fpath "github.com/lsejx/go-filepath"
)

func handleAskResult(operation string, path string) error {
	yes, err := interacter.ask(operation, path)
	if err != nil {
		return fmt.Errorf("%v: input error: %v", path, err)
	}
	if !yes {
		return errCanceled
	}
	return nil
}

func eraseFile(path string, op option) error {
	if op.interactive {
		if err := handleAskResult("erase", path); err != nil {
			return err
		}
	}
	if err := randomizeFile(path); err != nil {
		return err
	}
	if op.keep {
		return nil
	}
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("%v: removal error: %v", path, parsePathErr(err))
	}
	return nil
}

func eraseDir(path string, op option, errWriter io.Writer) error {
	if op.interactive {
		if err := handleAskResult("descend into", path); err != nil {
			return err
		}
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("%v: erasure error: %v", path, parsePathErr(err))
	}

	var wg sync.WaitGroup

	var errOccurred atomic.Bool
	errOccurred.Store(false)

	for _, entry := range entries {
		ePath := filepath.Join(path, entry.Name())
		if !fl.isNew(ePath) {
			continue
		}

		if entry.IsDir() {
			// directory
			wg.Add(1)
			go func(path string, op option) {
				defer wg.Done()
				if eraseDir(path, op, errWriter) != nil {
					errOccurred.CompareAndSwap(false, true)
				}
			}(ePath, op)
			continue
		}

		tp, err := fpath.GetType(ePath)
		if err != nil {
			errOccurred.CompareAndSwap(false, true)
			fmt.Fprintln(errWriter, fmt.Errorf("%v: %v", ePath, err))
		}
		if tp.IsRegularFile() {
			// file
			wg.Add(1)
			go func(path string, op option) {
				defer wg.Done()
				if err := eraseFile(path, op); err != nil {
					errOccurred.CompareAndSwap(false, true)
					if !errors.Is(err, errCanceled) {
						fmt.Fprintln(errWriter, err)
					}
				}
			}(ePath, op)
			continue
		}

		errOccurred.CompareAndSwap(false, true)
		fmt.Fprintln(errWriter, fmt.Errorf("%v: is neither a regular file nor a directory", ePath))
	}

	wg.Wait()
	if errOccurred.Load() {
		return errErrOccurred
	}
	if op.keep {
		return nil
	}
	if op.interactive {
		if err = handleAskResult("remove", path); err != nil {
			return err
		}
	}
	if err = os.Remove(path); err != nil {
		return fmt.Errorf("%v: removal error: %v", path, parsePathErr(err))
	}
	return nil
}
