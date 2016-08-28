# Exercise 12.3
Implement the missing cases of the `encode` function. Encode booleans as `t` and `nil`, floating-point numbers using Go's notation, and complex numbers like `1+2i` as `#C(1.0 2.0)`. Interfaces can be encoded as a pair of a type name and a value, for instance `("[]int" (1 2 3))`, but beware that this notation is ambiguous: the `reflect.Type.String` method may return the same string for different types.

---
# 練習問題 12.3
`encode`関数の不足している場合を実装しなさい。ブーリアンを`t`と`nil`として、浮動小数点数に`Go`の表記を使い、`1+2i`などの複素数は`#C(1.0, 2.0)`としてエンコードしなさい。インターフェースは、たとえば`("[]int" (1 2 3))`といったように型の名前と値の組としてエンコードできますが、この表記は曖昧であることに注意しなさい。`reflect.Type.String`メソッドは異なる方に対して同じ文字列を返すかもしれません。
