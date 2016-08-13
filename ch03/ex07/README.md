# Exercise 3.7
Another simple fractal uses Newton's method to find complex solution to a function such as `z^4-1=0`. Shade each starting point by the number of iterations required to get close to one of the four roots. Color each point by the root it approaches.

---
# 練習問題 3.7
別の単純なフラクタルは`z^4-1=0`などの方程式に対する複素素数解を求めるためにニュートン法を使用します。四つの根の1つの根に近づくのに必要な繰り返し回数で書く開始点にグラデーションをつけなさい。それが近づいている根ごとに各点に色付けしなさい。

# Result

````sh
$  go run newton.go > newton.png
````

![newton with color](./newton.png)
