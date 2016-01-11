# Exercise 1.7
The function call `io.Copy(dst, src)` reads from `src` and writes to `dst`. Use it instead of `ioutil.ReadAll` to copy the response body to `os.Stdout` without requiring a buffer large enough to hold the entire stream. Be sure to check the error result of `io.Copy`.

---
# 練習問題1.7
関数呼び出し`io.Copy(dst, src)`は`src`から読み込み`dst`へ書き込みます。レスポンスの内容を`os.Stdout`へコピーするために、ストリーム全体を保持するのに十分な大きさのバッファを必要としないようにするために、`ioutil.ReadAll`の代わりにその関数を使用しなさい。また、`io.Copy`のエラー結果を検査するようにしなさい。
