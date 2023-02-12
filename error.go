package main

import "fmt"

func catPathAndErr(path string, err error) error {
	return fmt.Errorf("%v: %v", path, err)
}
