# Exercise 9.6
Measure how the performance of a compute-bound parallel program (See Exercise 8.5) varies with `GOMAXPROCS`. What is the optimal value on your computer? How many CPUs does your computer have?

---
# 練習問題 9.6
計算が主体の並列プログラム（練習問題8.5を参照）の性能が`GOMAXPROCS`でどのように変化するか測定しなさい。みなさんのコンピュータ上での最適な値はいくつでしょうか。みなさんのコンピュータはCPUを何個もっているのでしょうか。


# Result
Try the Mandelbrot program, width = 5,000, and  height = 5,000.
ch08ex05 result was below, but it program does not be tuned. 


````shel
budougumi0617@~/git/gotraining/ch08/ex05 (remainingwork@GoTraining)
$  go run main.go > test.png
2016/08/19 12:23:09
MAX CPU NUM 4
USE CPU NUM 1
, Calculated time : 13.915584437s
2016/08/19 12:23:17
MAX CPU NUM 4
USE CPU NUM 2
, Calculated time : 8.586978251s
2016/08/19 12:23:26
MAX CPU NUM 4
USE CPU NUM 3
, Calculated time : 8.90179669s
2016/08/19 12:23:35
MAX CPU NUM 4
USE CPU NUM 4
, Calculated time : 9.18917844s
