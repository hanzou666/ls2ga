package main

import (
	"fmt"
	"github.com/docopt/docopt.go"
	"log"
	"os"
)

const version = "0.1.0-alpha"
const usage = `Last-Split To Global Alignment on graph structure

Usage:
  ls2ga [--node-max N] [--min-aln-size N] [--gfa FILE] [--json FILE] <target.fa> <query.fa> <aln.maf>

Options:
  -g --gfa FILE        output GFA file [default: graph.gfa]
  -j --json FILE       output alignment file [default: aln.json]
  -M --node-max N      maximum node sequence size [default: 1000]
  -m --min-aln-size N  minimum alignment block size [default: 1000]
  -V --version         show version
  -h --help            show this help message
`

// Command line arguments
type Args struct {
	MinAlnSize      int    `docopt:"--min-aln-size"`
	NodeMaxSize     int    `docopt:"--node-max"`
	GfaPath         string `docopt:"--gfa"`
	JsonPath        string `docopt:"--json"`
	TargetFastaPath string `docopt:"<target.fa>"`
	QueryFastaPath  string `docopt:"<query.fa>"`
	MafPath         string `docopt:"<aln.maf>"`
}

// Parse command line arguments
func parseArgument(argv []string) Args {
	if len(argv) == 0 {
		fmt.Print(usage)
		os.Exit(0)
	}

	opts, err := (&docopt.Parser{}).ParseArgs(usage, argv, version)
	if err != nil {
		log.Fatalln(err)
	}

	var args Args
	err = opts.Bind(&args)
	if err != nil {
		log.Fatalln(err)
	}
	return args
}

func run(args Args) {
	graph, err := fa2vg(args.TargetFastaPath, args.NodeMaxSize)
	if err != nil {
		log.Fatalln(err)
	}
	if err = graph.WriteGFA(args.GfaPath); err != nil {
		log.Fatalln(err)
	}
}

// Entry point
func main() {
	run(parseArgument(os.Args[1:]))
}
