package main

import "fmt"

func catPathAndErr(path string, summary string, err error) error {
	return fmt.Errorf("%v: %v, %v", path, summary, err)
}
