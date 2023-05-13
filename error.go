package main

import "fmt"

func catPathAndErr(path string, summary string, err error) error {
	return fmt.Errorf("error: %v %v: %v", summary, path, err)
}
