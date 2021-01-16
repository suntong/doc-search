////////////////////////////////////////////////////////////////////////////
// Program: doc-search
// Purpose: Doc search
// Authors: Tong Sun (c) 2021, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
	"github.com/mkideal/cli/clis"
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
	// argv.Query,
	//return nil
	return DoSearch()
}

//
// DoSearch implements the business logic of command `search`
func DoSearch() error {
	fmt.Fprintf(os.Stderr,
		"%s v%s search - Search the indexed doc archive\n",
		progname, version)
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2021, Tong Sun\n\n")
	// err := ...
	// clis.WarnOn("Doing Search", err)
	// or,
	// clis.AbortOn("Doing Search", err)
	return nil
}
