package main

import (
	"fmt"
	"os"
)

func writeFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	defer func() {
		_ = f.Close()
	}()

	if err != nil {
		return fmt.Errorf("open file error: %w", err)
	}
	if _, err = fmt.Fprint(f, content); err != nil {
		return fmt.Errorf("write content error: %w", err)
	}
	return nil
}
