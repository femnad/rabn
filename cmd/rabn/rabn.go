package main

import (
	"github.com/alexflint/go-arg"
	"github.com/femnad/mare"
	"github.com/femnad/rabn/pkg/rabn"
)

var args struct{
	HistoryFile string `arg:"-H,required"`
	PathSpec []string `arg:"-p,required,separate"`
	Selection string `arg:"positional" default:""`
}

func main() {
	arg.MustParse(&args)
	prefix :=  rabn.FindLongestCommonPrefix(args.PathSpec)
	h, err := rabn.HistoryFromFile(args.HistoryFile, prefix)
	mare.PanicIfErr(err)
	if args.Selection == "" {
		rabn.ListPathContentsWithHistory(h, args.PathSpec, prefix)
	} else {
		rabn.AddToHistory(h, args.Selection)
	}
}
