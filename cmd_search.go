////////////////////////////////////////////////////////////////////////////
// Program: doc-search
// Purpose: Doc search
// Authors: Tong Sun (c) 2021, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/blevesearch/bleve/v2"
	_ "github.com/blevesearch/bleve/v2/search/highlight/highlighter/ansi"

	"github.com/mkideal/cli"
	"github.com/mkideal/cli/clis"
)

const (
	MaxInt64 = 1<<63 - 1
)

////////////////////////////////////////////////////////////////////////////
// search

func searchCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*searchT)
	clis.Setup(fmt.Sprintf("%s::%s", progname, ctx.Path()), rootArgv.Verbose.Value())
	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
	Opts.BaseFolder, Opts.Group, Opts.Verbose =
		rootArgv.BaseFolder, rootArgv.Group, rootArgv.Verbose.Value()
	return DoSearch(getIdx(Opts.BaseFolder, Opts.Group), argv.Query)
}

//
// DoSearch implements the business logic of command `search`
func DoSearch(idx bleve.Index, query string) error {
	fmt.Fprintf(os.Stderr, "Doc-search - Search the indexed doc archive\n")
	q := bleve.NewQueryStringQuery(query)
	sreq := bleve.NewSearchRequestOptions(q, MaxInt64, 0, true)
	sreq.Highlight = bleve.NewHighlightWithStyle("ansi")
	sres, err := idx.Search(sreq)
	clis.AbortOn("Doing Search", err)
	fmt.Println(sres)
	return nil
}
