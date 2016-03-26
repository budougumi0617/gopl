# Exercise 5.
Write a function `expand(s string, f func(string) string) string` that replaces each substring "`$foo`" within `s` by the text returned by `f("foo")`.
---
# 練習問題 5.9
文字列`s`内のそれぞれの部部分文字列`"$foo"`を`f("foo")`が返すテキスト(`$`で始まる任意の単語を探して、`$`以降の文字列で関数`f`を呼び出した結果のテキスト)で置換する関数`expand(s string, f func(string) string) string`を書きなさい。
