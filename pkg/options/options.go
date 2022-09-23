package options

import "flag"

type Options struct {
	OutofScopeFile string
	InScopeFile    string
	DomainsFile    string
	ExcludeSubs    bool
	Others         bool
}

func Scanflags() Options {
	outofscopePtr := flag.String("of", "", "File with Out of Scope domains to exclude from domains supplied.")
	inscopePtr := flag.String("if", "", "File containing in scope domains to include.")
	domainsPtr := flag.String("df", "", "File containing all domains to process.")
	excludeSubsPtr := flag.Bool("exsubs", false, "Exclude subdomains of supplied domains.")
	othersPtr := flag.Bool("others", false, "Print other non-Out of Scope domains if InScope file is specified.")

	flag.Parse()

	results := Options{
		*outofscopePtr,
		*inscopePtr,
		*domainsPtr,
		*excludeSubsPtr,
		*othersPtr,
	}

	return results
}

func Usage() {
	flag.Usage()
}
