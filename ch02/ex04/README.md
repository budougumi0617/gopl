# Exercise 2.4
Write a version of `PopCount` that counts bits by shifting its arguments its argument through 64 bit positions, testing the rightmost bit each time. Compare its performance to the table-lookup version.

---
# 練習問題 2.4
引数をビットシフトしながら最下位ビットの検査を64回繰り返すことでビット数を数える`PopCount`を作成しなさい。テーブル参照を行うバージョンと性能を比較しなさい。

---
# Result of benchmark test

```sh
BenchmarkPopCount-4          	200000000	         7.11 ns/op
BenchmarkPopCountByShifting-4	20000000	        78.3 ns/op
ok  	github.com/budougumi0617/GoTraining/ch02/ex04	3.818s
```
