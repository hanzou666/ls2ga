package main

import (
	"fmt"
	"github.com/docopt/docopt.go"
	"os"
)

const version = "0.1.0-alpha"
const usage = `Last-Split To Global Alignment on graph structure

Usage:
  ls2ga [--min-aln-size=<N>] [--gfa=<FILE>] [--json=<FILE>] <fastaFile> <mafFile>
  ls2ga --version | -V
  ls2ga --help | -h

Options:
  -g --gfa=<FILE>        output GFA file [default: graph.gfa]
  -j --json=<FILE>       output alignment file [default: aln.json]
  -M --node-max=<N>      maximum node sequence size [default: 1000]
  -m --min-aln-size=<N>  minimum alignment block size [default: 1000]
  -V --version           show version
  -h --help              show this document
`

// Command line arguments
type Args struct {
	MinAlnSize  int    `docopt:"--min-aln-size"`
	NodeMaxSize int    `docopt:"node-max"`
	GfaPath     string `docopt:"--gfa"`
	JsonPath    string `docopt:"--json"`
	FastaPath   string `docopt:"<fastaFile>"`
	MafPath     string `docopt:"<mafFile>"`
}

// Parse command line arguments
func parseArgument(argv []string) Args {
	parser := &docopt.Parser{}

	opts, err := parser.ParseArgs(usage, argv, version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	var args Args
	err = opts.Bind(&args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	return args
}

func run(args Args)  {
	fmt.Println(args)
}


// Entry point
func main() {
	run(parseArgument(os.Args[1:]))
}