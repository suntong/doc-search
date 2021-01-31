支持 HMM 分词, 使用 viterbi 算法.

支持普通、搜索引擎、全模式、精确模式和 HMM 模式多种分词模式，支持用户词典、词性标注，可运行<a href="https://github.com/go-ego/gse/blob/master/server/server.go"> JSON RPC 服务</a>。

分词速度<a href="https://github.com/go-ego/gse/blob/master/benchmark/benchmark.go">单线程</a> 9.2MB/s，<a href="https://github.com/go-ego/gse/blob/master/benchmark/goroutines/goroutines.go">goroutines 并发</a> 26.8MB/s. HMM 模式单线程分词速度 3.2MB/s.（ 双核 4 线程 Macbook Pro）。
