package scoped

import (
	"bufio"
	"os"
)

func ReadFileToSlice(path string, dSlice []string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dSlice = append(dSlice, scanner.Text())
	}
	return scanner.Err()
}
