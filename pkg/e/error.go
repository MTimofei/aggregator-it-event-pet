package e

import "fmt"

func Err(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func IfErr(msg string, err error) error {
	if err == nil {
		return nil
	}
	return Err(msg, err)
}
