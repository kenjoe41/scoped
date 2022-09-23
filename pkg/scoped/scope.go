package scoped

import (
	"bufio"
	"os"
	"strings"
)

func Contains(dSlice *[]string, domain string, excludeSubs bool) bool {

	for _, dmn := range *dSlice {
		// Check if domain is a suffix of itered (sub)domain; true for google.com, www.google.com
		if domain == dmn {
			return true
		} else if excludeSubs && strings.HasSuffix(domain, dmn) {
			return true
		}
	}
	return false
}

func ReadFileToSlice(path string, dSlice *[]string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*dSlice = append(*dSlice, scanner.Text())
	}
	return scanner.Err()
}

func ReadFileToChan(path string, dChan chan string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dChan <- scanner.Text()
	}
	return scanner.Err()
}
