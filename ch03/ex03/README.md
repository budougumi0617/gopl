# Exercise 3.3
Color each polygon based on its height, so that the peaks are colored red(#ff0000) and the valleys blue(#0000ff).

---
# 練習問題 3.3
高さに基づいて個々のポリゴンに色付けし、頂点が赤(#ff0000)となり谷が青(#0000ff)になるようにしなさい。

# result

````sh
$  go run surface.go > test.svg
$  qlmanage -t -s 600 -o ./ ./test.svg
````

![surface with color](./test.svg.png)
