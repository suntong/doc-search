////////////////////////////////////////////////////////////////////////////
// Program: doc-search
// Purpose: Doc search
// Authors: Tong Sun (c) 2021, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	//  	"fmt"
	//  	"os"

	"github.com/mkideal/cli"
	//  	"github.com/mkideal/cli/clis"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

//==========================================================================
// doc-search

type rootT struct {
	cli.Helper
	BaseFolder string      `cli:"*B,base" usage:"base directory holding all indexes (mandatory)" dft:"$DS_BASE"`
	Group      string      `cli:"*G,group" usage:"index group all same doc belong (mandatory)" dft:"$DS_GROUP"`
	Verbose    cli.Counter `cli:"v,verbose" usage:"Verbose mode (Multiple -v options increase the verbosity)\n"`
}

var root = &cli.Command{
	Name: "doc-search",
	Desc: "Doc search\nVersion " + version + " built on " + date +
		"\nCopyright (C) 2021, Tong Sun",
	Text:   "CLI tool to do indexed full-text search on doc archives",
	Global: true,
	Argv:   func() interface{} { return new(rootT) },
	Fn:     DocSearch,

	NumArg: cli.AtLeast(1),
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
//  type OptsT struct {
//  	BaseFolder	string
//  	Group	string
//  	Verbose	cli.Counter
//  	Verbose int
//  }

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "doc-search"
//          version   = "0.1.0"
//          date = "2021-01-17"

//  	rootArgv *rootT
//  	// Opts store all the configurable options
//  	Opts OptsT
//  )

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.DenseNormalStyle)
//  	if err := cli.Root(root,
//  		cli.Tree(indexDef),
//  		cli.Tree(searchDef)).Run(os.Args[1:]); err != nil {
//  		fmt.Fprintln(os.Stderr, err)
//  		os.Exit(1)
//  	}
//  	fmt.Println("")
//  }

// Template for main dispatcher starts here
//==========================================================================
// Dumb root handler

// DocSearch - main dispatcher dumb handler
//  func DocSearch(ctx *cli.Context) error {
//  	ctx.JSON(ctx.RootArgv())
//  	ctx.JSON(ctx.Argv())
//  	fmt.Println()

//  	return nil
//  }

// Template for CLI handling starts here

////////////////////////////////////////////////////////////////////////////
// index

//  func indexCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*indexT)
//  	clis.Setup(fmt.Sprintf("%s::%s", progname, ctx.Path()), rootArgv.Verbose.Value())
//  	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
//  	Opts.BaseFolder, Opts.Group, Opts.Verbose, Opts.Verbose =
//  		rootArgv.BaseFolder, rootArgv.Group, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	// argv.IndexFolder, argv.IndexType, argv.Chinese,
//  	//return nil
//  	return DoIndex()
//  }
//
// DoIndex implements the business logic of command `index`
//  func DoIndex() error {
//  	fmt.Fprintf(os.Stderr, "Doc-search - Index doc archives\n")
//  	// fmt.Fprintf(os.Stderr, "Copyright (C) 2021, Tong Sun\n\n")
//  	// err := ...
//  	// clis.WarnOn("Doing Index", err)
//  	// or,
//  	// clis.AbortOn("Doing Index", err)
//  	return nil
//  }

type indexT struct {
	IndexFolder string `cli:"*d,dir" usage:"directory of the doc archive (mandatory)"`
	IndexType   string `cli:"t,type" usage:"type of files of the doc archive to index" dft:"txt,md"`
	Chinese     bool   `cli:"cc" usage:"index Chinese/CJK files"`
}

var indexDef = &cli.Command{
	Name: "index",
	Desc: "Doc-search - Index doc archives",
	Text: "Usage:\n  ds index [Options]\n\nExamples:\n  ds index -B ~/.ds -G blogs -d myBlogs\n  DS_BASE=~/.ds DS_GROUP=blogs ds index -d myBlogs -t md --cc",
	Argv: func() interface{} { return new(indexT) },
	Fn:   indexCLI,

	NumOption: cli.AtLeast(1),
}

////////////////////////////////////////////////////////////////////////////
// search

//  func searchCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*searchT)
//  	clis.Setup(fmt.Sprintf("%s::%s", progname, ctx.Path()), rootArgv.Verbose.Value())
//  	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
//  	Opts.BaseFolder, Opts.Group, Opts.Verbose, Opts.Verbose =
//  		rootArgv.BaseFolder, rootArgv.Group, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	// argv.Query, argv.FileOnly, argv.DeepSearch,
//  	//return nil
//  	return DoSearch()
//  }
//
// DoSearch implements the business logic of command `search`
//  func DoSearch() error {
//  	fmt.Fprintf(os.Stderr, "Doc-search - Search the indexed doc archive\n")
//  	// fmt.Fprintf(os.Stderr, "Copyright (C) 2021, Tong Sun\n\n")
//  	// err := ...
//  	// clis.WarnOn("Doing Search", err)
//  	// or,
//  	// clis.AbortOn("Doing Search", err)
//  	return nil
//  }

type searchT struct {
	Query      string `cli:"*q,query" usage:"query in bleve search syntax (mandatory)"`
	FileOnly   bool   `cli:"l,files" usage:"like grep -l, print the file name & suppress content output"`
	DeepSearch bool   `cli:"g,grep" usage:"use grep to search files for more hits than the first match"`
}

var searchDef = &cli.Command{
	Name: "search",
	Desc: "Doc-search - Search the indexed doc archive",
	Text: "Usage:\n  ds search [Options]",
	Argv: func() interface{} { return new(searchT) },
	Fn:   searchCLI,

	NumOption: cli.AtLeast(1),
}
