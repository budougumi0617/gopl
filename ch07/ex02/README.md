# Exercise 7.2
Write a function `CountingWriter` with the signature below that, given an `io.Writer`, returns a new `Writer` that wraps the original, and a pointer to an `int64` variable that at any moment contains the number of bytes written to the new `Writer`.


---
# 練習問題 7.2
下記のシグネチャを持つ関数`CounterWriter`を書きなさい。`io.Writer`が与えられたなら、それを包む新たな`Writer`と`int64`変数へのポインタを返します。その変数は新たな`Writer`に書き込まれたバイト数を常に保持しています。

````go
func CounterWriter(w io.Writer) (io.Writer, *int64)
````


# Result
```bash
$  go run counterwriter.go
test input
test input
Byte count 22
```
