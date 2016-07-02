# Exercise 8.5
Take an existing CPU-bound sequential program, such as the Mandelbrot program of Section 3.3 or the 3-D surface computation of Section 3.2, and execute its main loop in parallel using channels for communication. How much faster does it run on a multiprocessor machine? What is the optional number of goroutines to use?

---
# 練習問題 8.5
3.3節のマンデブロのプログラムや3.2節の3-D面の計算といった既存のCPU制約の逐次プログラムを選び、通信用のチャネルを使って、それらのプログラムのメインループを並列に実行しなさい。マルチプロセッサのコンピュータ上ではどのくらい速く動作するでしょうか。使うゴルーチンの最適な数はいくつでしょうか。
