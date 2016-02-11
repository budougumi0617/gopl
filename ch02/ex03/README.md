# Exercise 2.3
Rewrite `PopCount` to use a loop instead of a single expression. Compare the performance of the two versions. (Section 11.4 shows how to compare the performance of different implementations systematically.)

---
# 練習問題 2.3
単一の式の代わりにループを使用するように`PopCount`を書き直しなさい。2つのバージョンの性能を比較しなさい（11.4節で異なる実装の性能をシステム的に比較する方法を説明しています）。

---
# Result of benchmark test

```sh
BenchmarkPopCount-4      	200000000	         7.02 ns/op
BenchmarkPopCountByLoop-4	100000000	        18.3 ns/op
ok  	github.com/budougumi0617/GoTraining/ch02/ex03	4.015s
```
