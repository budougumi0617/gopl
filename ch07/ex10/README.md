# Exercise 7.10
The `sort.Interface` type can be adapted to other uses. Write a function `IsPalindrome(s sort.Interface) bool` that reports whether the sequence `s` is palindrome, in other words, reversing the sequence would not change it. Assume that the elements at indices `i` and `j` are equal if `!s.Less(i,j) && s.Less(j,i)`.

---
# 練習問題 7.10
`sort.Interface`型は、他の利用にも適応できます。列`s`が回分(palindrome)であるか、つまり列を逆順にしても変わらないかを報告する関数`IsPalindrome(s sort.Interface) bool`を書きなさい。インデックス`i`と`j`の要素は`!s.Less(i,j) && s.Less(j,i)`であれば等しいとみなしなさい。
