# Exercise 12.2
Make `Display` safe to use on cyclic data structures by bounding the number of steps it takes before abandoning the recursion. (In Section 13.3, we'll see another way to detect cycles.)

---
# 練習問題 12.2
再帰呼び出しの中断に先立って`display`が行うステップを制限することで、循環したデータ構造に対して、`display`を安全に使えるようにしなさい。(13.3節では、循環を検出する別の方法を説明します。)


# Result

````shell
budougumi0617@~/git/gotraining/ch12/ex02 (ch12ch13_remainingwork@GoTraining)
$  go run display.go
Display TestCycle (main.Cycle):
TestCycle.Value = 42
(*TestCycle.Tail).Value = 42
(*(*TestCycle.Tail).Tail).Value = 42
(*(*(*TestCycle.Tail).Tail).Tail).Value = 42
(*(*(*(*TestCycle.Tail).Tail).Tail).Tail).Value = 42
(*(*(*(*(*TestCycle.Tail).Tail).Tail).Tail).Tail).Value = 42
cyclic limit 5
````
