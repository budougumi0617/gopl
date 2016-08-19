# Exercise 8.1
Modify `clock2` to accept a port number, and write a program, `clockwall`, that acts as a client of server clock servers at once, reading the times from each one and displaying the results in a table, akin to the wall of clocks seen in some business offices. If you have access to geographically distributed computers, run instances remotely; otherwise run local instances on different with ports with fake time zones.

---
# 練習問題 8.1
`clock2`を修正してポート番号を受け付けるようにしなさい。そして、一度に複数の時計サーバのクライアントとして振る舞い、それぞれのサーバから時刻を読みだし、ビジネスオフィスで見かける壁にかかった複数の時計に似せた表で結果を表示するプログラム`clockwall`を書きなさい。みなさんが地理的に分散したコンピュータへアクセスできるなら、サーバをリモートで実行しなさい。そうでなければ、次のように擬似的なタイムゾーンを用いて異なるポートでローカルにサーバを実行しなさい。

````shell
$ TZ=US/Eastern  go run clock2/clock2.go -port 8010 &
$ TZ=Asia/Tokyo  go run clock2/clock2.go -port 8020 &
$ TZ=Europe/London  go run clock2/clock2.go -port 8030 &
$ go run clockwall.go NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
````


# Result

````shell
budougumi0617@~/git/gotraining/ch08/ex01 (remainingwork@GoTraining)
$  TZ=US/Eastern  go run clock2/clock2.go -port 8010 &
[1] 25801
budougumi0617@~/git/gotraining/ch08/ex01 (remainingwork@GoTraining)
$  ls
README.md    clock2       clockwall.go
budougumi0617@~/git/gotraining/ch08/ex01 (remainingwork@GoTraining)
$  TZ=Asia/Tokyo  go run clock2/clock2.go -port 8020 &
[2] 25906
budougumi0617@~/git/gotraining/ch08/ex01 (remainingwork@GoTraining)
$  TZ=Europe/London  go run clock2/clock2.go -port 8030 &
[3] 25953
budougumi0617@~/git/gotraining/ch08/ex01 (remainingwork@GoTraining)
$  go run clockwall.go NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
NewYork	Tokyo	London
00:05:39	13:05:39	05:05:39
NewYork	Tokyo	London
00:05:40	13:05:40	05:05:40
NewYork	Tokyo	London
00:05:41	13:05:41	05:05:41
NewYork	Tokyo	London
00:05:42	13:05:42	05:05:42
NewYork	Tokyo	London
00:05:43	13:05:43	05:05:43
NewYork	Tokyo	London
00:05:44	13:05:44	05:05:44
````
