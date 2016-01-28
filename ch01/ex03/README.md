#Exercise 1.3
Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses `strings.Join`.(Section 1.6 illustrates part of the `time` package, and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)

---
# 練習問題 1.3
非効率な可能性のあるバージョンと`strings.Join`を使用したバージョンとで、実行時間の差を計測しなさい（1.6節は`time`パッケージの一部を説明いますし、11.4節ではシステム的に性能評価を行うためのベンチマークテストの書き方を説明しています）。

-------

# Result of benchmark test

```sh
budougumi0617@~/git/gotraining/ch01/ex03 (master@GoTraining)
$  go test -bench .
PASS
BenchmarkEcho1-4	    2000	    756040 ns/op
BenchmarkEcho2-4	    2000	    569847 ns/op
BenchmarkEcho3-4	   50000	     34543 ns/op
BenchmarkEcho4-4	   10000	    311670 ns/op
ok  	github.com/budougumi0617/GoTraining/ch01/ex03	8.115s
```
