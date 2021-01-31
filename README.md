
[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/suntong/doc-search?status.svg)](http://godoc.org/github.com/suntong/doc-search)
[![Go Report Card](https://goreportcard.com/badge/github.com/suntong/doc-search)](https://goreportcard.com/report/github.com/suntong/doc-search)
[![travis Status](https://travis-ci.org/suntong/doc-search.svg?branch=master)](https://travis-ci.org/suntong/doc-search)
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-R.svg)](http://godoc.org/github.com/go-easygen/wireframe)

## TOC
- [doc-search, ds - Doc search cli tool](#doc-search-ds---doc-search-cli-tool)
- [Usage](#usage)
  - [$ doc-search](#-doc-search)
    - [$ doc-search index](#-doc-search-index)
    - [$ doc-search search](#-doc-search-search)
  - [Examples](#examples)
    - [Index, English](#index-english)
    - [Search, English](#search-english)
  - [Index, Chinese](#index-chinese)
    - [Search, Chinese](#search-chinese)
  - [How to get](#how-to-get)
    - [Debian package](#debian-package)
    - [Install Source](#install-source)
  - [Credits](#credits)
  - [Author(s) & Contributor(s)](#author(s)-&-contributor(s))

## doc-search, ds - Doc search cli tool

The `doc-search` is the solution to indexed files searching, especially for Chinese, where there is no good Chinese indexed searching tools so far, apart from the Java based elastic-search.

I have personal reservation to all Java based tools so I decided to write one in Go instead.


## Usage

### $ doc-search
```sh
Doc search
Version 1.0.0 built on 2021-01-30
Copyright (C) 2021, Tong Sun

CLI tool to do indexed full-text search on doc archives

Options:

  -h, --help      display help information 
  -B, --base     *base directory holding all indexes (mandatory) [=$DS_BASE]
  -G, --group    *index group all same doc belong (mandatory) [=$DS_GROUP]
  -v, --verbose   Verbose mode (Multiple -v options increase the verbosity) 


Commands:

  index    Doc-search - Index doc archives
  search   Doc-search - Search the indexed doc archive
```

#### $ doc-search index
```sh
Doc-search - Index doc archives

Usage:
  ds index [Options]

Examples:
  ds index -B ~/.ds -G blogs -d myBlogs
  DS_BASE=~/.ds DS_GROUP=blogs ds index -d myBlogs -t md --cc

Options:

  -h, --help      display help information 
  -B, --base     *base directory holding all indexes (mandatory) [=$DS_BASE]
  -G, --group    *index group all same doc belong (mandatory) [=$DS_GROUP]
  -v, --verbose   Verbose mode (Multiple -v options increase the verbosity) 

  -d, --dir      *directory of the doc archive (mandatory) 
  -t, --type      type of files of the doc archive to index [=txt,md]
  -a, --abs       use abs-path as file names 
      --cc        index Chinese/CJK files
```

#### $ doc-search search
```sh
Doc-search - Search the indexed doc archive

Usage:
  ds search [Options]

Options:

  -h, --help      display help information 
  -B, --base     *base directory holding all indexes (mandatory) [=$DS_BASE]
  -G, --group    *index group all same doc belong (mandatory) [=$DS_GROUP]
  -v, --verbose   Verbose mode (Multiple -v options increase the verbosity) 

  -q, --query    *query in bleve search syntax (mandatory) 
  -l, --files     like grep -l, print the file name & suppress content output 
  -g, --grep      use grep to search files for more hits than the first match
```

Note, `doc-search` also has a shorter version of the command name -- `ds`. For example:

    DS_BASE=~/.ds DS_GROUP=blogs ds search KW


## Examples

### Index, English

```sh
rm -rf ~/.ds/ds-index

$ ds index -B ~/.ds -G ds-index -d . -t go
Doc-search - Index doc archives
Done
```

If you see a warning output of `Could not load dictionaries: "./data/dict/zh/dict.txt", open ./data/dict/zh/dict.txt: no such file or directory`, that's OK. It doesn't affect the program functioning.

### Search, English

```sh
$ DS_BASE=~/.ds ds search -G ds-index -q clis
Doc-search - Search the indexed doc archive
4 matches, showing 1 through 4, took 1.030794ms

======
doc-search_main.go
		…-search_cliGen.sh

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
	//  	"github.com/mkideal/cli/clis"
)

////////////////////////////////////////////////////////////////////////////
// Constant and …

======
cmd_search.go
		…h/bleve/v2/search/highlight/highlighter/ansi"

	"github.com/mkideal/cli"
	"github.com/mkideal/cli/clis"
)

const (
	MaxInt64 = 1<<63 - 1
)

////////////////////////////////////////////////////////////…

======
cmd_index.go
		…g"
	_ "github.com/leopku/bleve-gse-tokenizer"

	"github.com/mkideal/cli"
	"github.com/mkideal/cli/clis"
)

////////////////////////////////////////////////////////////////////////////
// index

func i…

======
doc-search_cliDef.go
		…ge main

import (
	//  	"fmt"
	//  	"os"

	"github.com/mkideal/cli"
	//  	"github.com/mkideal/cli/clis"
)

////////////////////////////////////////////////////////////////////////////
// Constant and …

```

Note the behavior of `bleve` is, it returns/highlights the first hit(s) from a file, but doesn't return the rest of hits. As per Marty Schoch, the creator of the full-text search and indexing engine,

> bleve has a unit of operation called a document. You add documents to the index, and when you search, you find out which documents matched. You're indexing files as the documents, so no matter how many times that file matches, bleve considers it one match.

So if you do want to highlights the rest of hits from a file, add the `-g/--grep` option to `search`, _"to search files for more hits than the first match"_:

```sh
$ DS_BASE=~/.ds ds search -G ds-index -q clis -g | head -12
Doc-search - Search the indexed doc archive

======
doc-search_main.go
	//  	"github.com/mkideal/cli/clis"

======
cmd_search.go
	"github.com/mkideal/cli/clis"
	clis.Setup(fmt.Sprintf("%s::%s", progname, ctx.Path()), rootArgv.Verbose.Value())
	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
	clis.AbortOn("Doing Search", err)
				clis.WarnOn("Run grep", err)

```


## Index, Chinese

```sh
rm -rf ~/.ds/test

$ ds index -B ~/.ds -G test -d test --cc
Doc-search - Index doc archives
Done
```

### Search, Chinese

```sh
$ DS_BASE=~/.ds DS_GROUP=test ds search -q 分词
Doc-search - Search the indexed doc archive
2 matches, showing 1 through 2, took 676.33µs

======
test/gse1.md
		…分词, 支持英文、中文、日文等


<a href="https://github.com/go-ego/gse/blob/master/dictionary.go">词典</a>用双数组 trie（Double-Array Trie）实现，
<a href="https://github.com/go-ego/gse/blob/master/segmenter.go">分词器</a>算法为基于词…

======
test/gse2.md
		…MM 分词, 使用 viterbi 算法.

支持普通、搜索引擎、全模式、精确模式和 HMM 模式多种分词模式，支持用户词典、词性标注，可运行<a href="https://github.com/go-ego/gse/blob/master/server/server.go"> JSON RPC 服务</a>。

分词速度<a href="https://github.com/go-ego/gs…

```

## How to get

### Debian package

Will be available...


### Install Source

To install the source code instead:

```
go get github.com/suntong/doc-search
```


## Credits

- https://github.com/blevesearch/bleve
- https://github.com/leopku/bleve-gse-tokenizer
- https://github.com/go-ego/gse
- https://github.com/ques0942/local-search
- https://github.com/yookoala/realpath/blob/master/realpath.go

## Author(s) & Contributor(s)

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

_Powered by_ [**WireFrame**](https://github.com/go-easygen/wireframe),  [![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-Y.svg)](http://godoc.org/github.com/go-easygen/wireframe), the _one-stop wire-framing solution_ for Go cli based projects, from start to deploy.

All patches welcome. 
