////////////////////////////////////////////////////////////////////////////
// Program: doc-search
// Purpose: Doc search
// Authors: Tong Sun (c) 2021, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

//go:generate sh -v doc-search_cliGen.sh

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
	//  	"github.com/mkideal/cli/clis"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
type OptsT struct {
	BaseFolder string
	Group      string
	Verbose    int
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "doc-search"
	version  = "0.1.0"
	date     = "2021-01-16"

	rootArgv *rootT
	// Opts store all the configurable options
	Opts OptsT
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	cli.SetUsageStyle(cli.DenseNormalStyle)
	//NOTE: You can set any writer implements io.Writer
	// default writer is os.Stdout
	if err := cli.Root(root,
		cli.Tree(indexDef),
		cli.Tree(searchDef)).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("")
}

//==========================================================================
// Dumb root handler

// DocSearch - main dispatcher dumb handler
func DocSearch(ctx *cli.Context) error {
	ctx.JSON(ctx.RootArgv())
	ctx.JSON(ctx.Argv())
	fmt.Println()

	return nil
}
