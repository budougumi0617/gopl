# Exercise 7.5
The `LimitReader` function in the `io` package accepts an `io.Reader r` and a number of bytes `n`, and returns another `Reader` that reads from `r` but reports an end-of-file condition after `n` bytes. Implement it.

---
# 練習問題 7.5
`io`パッケージの`LimitReader`関数は`io.Reader`である`r`とバイト数`n`を受け取り、`r`から読み出す別の`Reader`を返しますが、`n`バイトを読み出した後にファイルの終わりの状態を報告します。その関数を実装しなさい。

````go
func LimitReader(r io.Reader, n int64) io.Reader
````
