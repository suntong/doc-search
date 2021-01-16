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
// index

func indexCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*indexT)
	clis.Setup(fmt.Sprintf("%s::%s", progname, ctx.Path()), rootArgv.Verbose.Value())
	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
	Opts.BaseFolder, Opts.Group, Opts.Verbose =
		rootArgv.BaseFolder, rootArgv.Group, rootArgv.Verbose.Value()
	// argv.IndexFolder, argv.IndexType, argv.ChineseChar,
	//return nil
	return DoIndex()
}

//
// DoIndex implements the business logic of command `index`
func DoIndex() error {
	fmt.Fprintf(os.Stderr,
		"%s v%s index - Index doc archives\n",
		progname, version)
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2021, Tong Sun\n\n")
	// err := ...
	// clis.WarnOn("Doing Index", err)
	// or,
	// clis.AbortOn("Doing Index", err)
	return nil
}
