package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/kenjoe41/scoped/pkg/options"
	"github.com/kenjoe41/scoped/pkg/scoped"
)

func main() {
	flags := options.Scanflags()

	outOfScopeSlice := []string{}
	inScopeSlice := []string{}

	// If we have an out of scope file, let's read it to a slice.
	if flags.OutofScopeFile != "" {
		err := scoped.ReadFileToSlice(flags.OutofScopeFile, &outOfScopeSlice)
		if err != nil {
			log.Fatalf("error reading Out of Scope domains file: %s", err)
		}

	}

	// if we have an in scope domains file, let's read them to a slice.
	if flags.InScopeFile != "" {
		err := scoped.ReadFileToSlice(flags.InScopeFile, &inScopeSlice)
		if err != nil {
			log.Fatalf("error reading in scope domains file: %s", err)
		}
	}

	domainsChan := make(chan string)
	outputChan := make(chan string)

	var domainsWG sync.WaitGroup
	domainsWG.Add(1)

	go func() {
		defer domainsWG.Done()

		for domain := range domainsChan {

			// Check if domain is in Out of Scope Slice
			if !scoped.Contains(&outOfScopeSlice, domain, flags.ExcludeSubs) {

				// If inScopeSlice is not empty, then we have to print only inscope domains.
				if len(inScopeSlice) > 0 {
					if scoped.Contains(&inScopeSlice, domain, flags.ExcludeSubs) {
						outputChan <- domain
					}
				} else {
					// inScopeSlice is empty so we print the domain
					outputChan <- domain
				}
			} else if scoped.Contains(&inScopeSlice, domain, true) {
				outputChan <- domain
			}
		}

	}()

	var outputWG sync.WaitGroup
	outputWG.Add(1)
	go func() {
		defer outputWG.Done()

		for domain := range outputChan {
			fmt.Println(domain)
		}
	}()

	// Close the Output Chan after domain worker is done.
	go func() {
		domainsWG.Wait()
		close(outputChan)
	}()

	// Check if we have a domains input file.
	// And read it to domainsChan
	if flags.DomainsFile != "" {
		err := scoped.ReadFileToChan(flags.DomainsFile, domainsChan)
		if err != nil {
			log.Fatal(err)
		}
	} else {

		// Check for stdin input
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			fmt.Fprintln(os.Stderr, "No domains or urls detected. Hint: cat domains.txt | scoped -of outofscope.txt")
			flag.Usage()
			os.Exit(1)
		}

		sc := bufio.NewScanner(os.Stdin)

		for sc.Scan() {
			domainsChan <- sc.Text()
		}

		// check there were no errors reading stdin (unlikely)
		if err := sc.Err(); err != nil {
			log.Fatal(err)
		}
	}

	// Close DomainsChan and wait until the output waitgroup is done
	close(domainsChan)
	outputWG.Wait()
}
