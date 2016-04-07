# Exercise 6.5
The type of each word used by `IntSet` is `uint64`, but 64-bit arithmetic may be inefficient on a 32-bit platform. Modify the program to use the `uint` type, which is the most efficient unsigned integer type for the platform. Instead of dividing by 64, define a constant holding the effective size of `uint` in bits, 32 or 64. You can use the perhaps too-claver expression `32 << (^uint(0) >> 63)` for this purpose.

---
# 練習問題 6.5
`IntSet`で使用されている個々のワードの型は`uint64`ですが、64ビット演算は32ビットプラットフォームでは非効率かもしれません。プラットフォームに対して最も効率的な符号なし整数である`uint`を使用するようプログラムを修正しなさい。64で割る代わりに、`uint`の実質的サイズのビット数である32あるいは64を保持する定数を定義しなさい。そのためには、おそらくかなり賢い式`32 << (^uint(0) >> 63)`を使用できます。
