package options

import "flag"

type Options struct {
	OutofScopeFile string
	InScopeFile    string
	DomainsFile    string
}

func Scanflags() Options {
	outofscopePtr := flag.String("of", "", "File with Out of Scope domains to exclude from domains supplied.")
	inscopePtr := flag.String("if", "", "File containing in scope domains to include.")
	domainsPtr := flag.String("df", "", "File containing all domains to process.")

	flag.Parse()

	results := Options{
		*outofscopePtr,
		*inscopePtr,
		*domainsPtr,
	}

	return results
}

func Usage() {
	flag.Usage()
}
