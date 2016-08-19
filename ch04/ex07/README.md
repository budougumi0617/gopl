# Exercise 4.7
Modify `reverse` to reverse the characters of a `[]byte` slice that represents a UTF-8-encoded string, in place. Can you do it without allocating new memory?

---
# 練習問題 4.7
UTF-8でエンコードされた文字列を表す`[]byte`スライスの文字を、そのスライス内で逆順にするように`reverse`を修正しなさい。新たなメモリを割り当てることなく行うことはできるでしょうか。


# 解答
できない。UTF-8エンコーディングで占められる1文字あたりのバイト数は一定ではないため、単純な一文字ずつの置換ができず、テンポラリメモリが必要になる。
