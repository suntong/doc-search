////////////////////////////////////////////////////////////////////////////
// Program: doc-search
// Purpose: Doc search
// Authors: Tong Sun (c) 2021, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
	_ "github.com/leopku/bleve-gse-tokenizer"

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
	return DoIndex(getIdx(Opts.BaseFolder, Opts.Group, argv.Chinese),
		argv.IndexFolder, argv.IndexType, argv.AbsPath)
}

//
// DoIndex implements the business logic of command `index`
func DoIndex(idx bleve.Index, indexFolder string, indexType string,
	absPath bool) error {
	fmt.Fprintf(os.Stderr, "Doc-search - Index doc archives\n")
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2021, Tong Sun\n\n")
	defer idx.Close()
	if absPath {
		realpath, err := Realpath(indexFolder)
		clis.AbortOn("Get realpath", err)
		indexFolder = realpath
	}

	var types []string
	for _, t := range strings.Split(indexType, ",") {
		types = append(types, strings.TrimSpace(t))
	}
	err := indexingFiles(idx, indexFolder, types)
	// clis.WarnOn("Doing Index", err)
	clis.AbortOn("Creating Index", err)
	fmt.Fprintf(os.Stderr, "Done\n")
	return nil
}

func getIdx(baseFolder, group string, isChinese bool) bleve.Index {
	idxPath := filepath.Join(baseFolder, group)
	idx, err := bleve.Open(idxPath)
	if err == bleve.ErrorIndexPathDoesNotExist {
		idx, err = createIndex(idxPath, isChinese)
	}
	clis.Verbose(1, "Index '%s' opened", idxPath)
	return idx
}

func createIndex(idxPath string, isChinese bool) (bleve.Index, error) {
	m, err := buildMapping(isChinese)
	clis.AbortOn("Build mapping", err)
	return bleve.New(idxPath, m)
}

func buildMapping(isChinese bool) (mapping.IndexMapping, error) {
	// open a new index
	m := bleve.NewIndexMapping()
	if isChinese {
		clis.Verbose(1, "Chinese mode is on")
		err := m.AddCustomTokenizer("gse", map[string]interface{}{
			"type":       "gse",
			"user_dicts": "",
		})
		clis.AbortOn("Add custom tokenizer", err)
		err = m.AddCustomAnalyzer("gse", map[string]interface{}{
			"type":      "gse",
			"tokenizer": "gse",
		})
		clis.AbortOn("Add custom analyzer", err)
		m.DefaultAnalyzer = "gse"
	}
	dm := bleve.NewDocumentMapping()
	tf := bleve.NewTextFieldMapping()
	dm.AddFieldMappingsAt("Content", tf)
	m.AddDocumentMapping("Doc", dm)
	return m, nil
}

func indexingFiles(idx bleve.Index, dir string, types []string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !toIndex(info.Name(), types) {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		buf, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		d := Doc{
			Path:    path,
			Content: string(buf),
		}
		return idx.Index(path, d)
	})
}

func toIndex(file string, exts []string) bool {
	for _, ext := range exts {
		if strings.HasSuffix(file, ext) {
			return true
		}
	}
	return false
}

type Doc struct {
	Path    string
	Content string
}
