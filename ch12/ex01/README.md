# Exercise 12.1
Extend `Display` so that it can display maps whose keys are structs or arrays.

---
# 練習問題 12.1
`Display`を拡張して、キーが構造体か配列であるマップを表示出来るようにしなさい。


# Result

````shell
budougumi0617@~/git/gotraining/ch12/ex01 (ch12ch13_remainingwork@GoTraining)
$  go run display.go
Display TestMap (map[main.Key]int):
TestMap[main.Key.i = 1
main.Key.structs.b = true
main.Key.arrays[0] = "foo"
main.Key.arrays[1] = "bar"
] = 10
Display TestMapArrays (map[[2]string]int):
TestMapArrays[[2]string[0] = "foo"
[2]string[1] = "bar"
] = 10
````
