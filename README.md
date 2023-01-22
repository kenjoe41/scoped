# Scoped 
Scoped is a tool that takes in a list of domains, and filters them based on whether they are in or out of scope. It is useful for bug bounty hunters and penetration testers to quickly identify which domains they should focus on.

## Installation

```bash
go install github.com/kenjoe41/scoped@latest
```

## Usage 

```bash
$ scoped -h

Usage of scoped:
  -df string
        Domain input file.
  -exclude-subs
        Exclude subdomains when matching.
  -if string
        In scope domains file.
  -of string
        Out of scope domains file.

```

## Examples

- Read domains from a file and filter out the out-of-scope domains:
```bash
scoped -df domains.txt -of outofscope.txt
```
This command reads the list of domains from the file 'domains.txt' and filters out the domains specified in the outofscope.txt file, and output the filtered domains to the console

- Read domains from stdin and filter out the out-of-scope domains:
```bash
cat domains.txt | scoped -of outofscope.txt
```

- Read domains from a file and filter out the out-of-scope domains and only print in-scope domains:
```bash
scoped -df domains.txt -of outofscope.txt -if inscope.txt
```
This command reads the list of domains from the file 'domains.txt' and filters out the domains specified in the outofscope.txt file, then it filters only the domains present in inscope.txt file, and output the filtered domains to the console

- Read domains from a file and filter out the out-of-scope domains and exclude subdomains:
```bash
scoped -df domains.txt -of outofscope.txt -exclude-subs
```
This command reads the list of domains from the standard input and filters out the domains specified in the outofscope.txt file, also it excludes all subdomains of the filtered domains, and output the filtered domains to the console


## Dependencies
* Golang 1.11 or later

## Contributing

1. Fork the repository
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request

## License

Scoped is released under the MIT License. See [LICENSE](https://github.com/kenjoe41/scoped/blob/main/LICENSE) for more details.
