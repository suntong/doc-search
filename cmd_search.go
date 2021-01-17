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
	fmt.Println(searchResult{sres})
	return nil
}

type searchResult struct {
	*bleve.SearchResult
}

func (sr searchResult) String() string {
	rv := ""
	if sr.Total > 0 {
		if sr.Request.Size > 0 {
			rv = fmt.Sprintf("%d matches, showing %d through %d, took %s\n", sr.Total, sr.Request.From+1, sr.Request.From+len(sr.Hits), sr.Took)
			for i, hit := range sr.Hits {
				rv += fmt.Sprintf("%5d. %s (%f)\n", i+sr.Request.From+1, hit.ID, hit.Score)
				for fragmentField, fragments := range hit.Fragments {
					rv += fmt.Sprintf("\t%s\n", fragmentField)
					for _, fragment := range fragments {
						rv += fmt.Sprintf("\t\t%s\n", fragment)
					}
				}
				for otherFieldName, otherFieldValue := range hit.Fields {
					if _, ok := hit.Fragments[otherFieldName]; !ok {
						rv += fmt.Sprintf("\t%s\n", otherFieldName)
						rv += fmt.Sprintf("\t\t%v\n", otherFieldValue)
					}
				}
			}
		} else {
			rv = fmt.Sprintf("%d matches, took %s\n", sr.Total, sr.Took)
		}
	} else {
		rv = "No matches"
	}
	if len(sr.Facets) > 0 {
		rv += fmt.Sprintf("Facets:\n")
		for fn, f := range sr.Facets {
			rv += fmt.Sprintf("%s(%d)\n", fn, f.Total)
			for _, t := range f.Terms {
				rv += fmt.Sprintf("\t%s(%d)\n", t.Term, t.Count)
			}
			if f.Other != 0 {
				rv += fmt.Sprintf("\tOther(%d)\n", f.Other)
			}
		}
	}
	return rv
}
