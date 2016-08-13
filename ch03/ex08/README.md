# Exercise 3.8
Rendering fractals at high zoom levels demands great arithmetic precision. Implement the same fractal using four different representations of numbers: `complex64`, `complex128`, `big.Float`, and `big.Rat`. (The latter to types are found in the `math/big` package. `Float` uses arbitrary but bounded-precision floating-point; `Rat` uses unbounded-precision rational numbers.) How do the compare in performance and memory usage? At what zoom levels do rendering artifacts become visible?

---
# 練習問題 3.8
高倍率の水準でフラクタルをレンダリングするには高い算術精度が求められます。`complex64`,`complex128`,`big.Float`,`big.Rat`の四つの異なる数値の表現を使用して同じフラクタル画像を実装しなさい。（最後の2つの型は`math/big`パッケージにあります。`Float`は任意精度ですが、有界精度の浮動小数点を使用しています。`Rat`は非有界精度の有理数を使用しています。）性能とメモリ使用量に関してどのような比較結果になりますか。どの倍率の水準になるとレンダリングの結果が視覚的に分かるようになりますか。
