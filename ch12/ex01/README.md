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
main.Key.structs.arrays[0] = "foo"
main.Key.structs.arrays[1] = "bar"
] = 10
TestMap[main.Key.i = 2
main.Key.structs.arrays[0] = "bar"
main.Key.structs.arrays[1] = "foo"
] = 20
````
