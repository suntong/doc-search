# [gse](https://github.com/go-ego/gse)

Go 语言高效分词, 支持英文、中文、日文等


<a href="https://github.com/go-ego/gse/blob/master/dictionary.go">词典</a>用双数组 trie（Double-Array Trie）实现，
<a href="https://github.com/go-ego/gse/blob/master/segmenter.go">分词器</a>算法为基于词频的最短路径加动态规划, 以及 DAG 和 HMM 算法分词.
