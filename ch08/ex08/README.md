# Exercise 8.8
Using a select statement, add a timeout to the echo server from Section 8.3 so that it disconnects any client that shouts nothing within 10 seconds.

---
# 練習問題 8.8
`select`文を使って、8.3節のエコーサーバにタイムアウトを追加し、10秒以内に何も叫ばないクライアントとの接続を切るようにしなさい。


````shell
budougumi0617@~/git/gotraining/ch08/ex08 (remainingwork@GoTraining)
$  go run reverb2.go &
[1] 38925
budougumi0617@~/git/gotraining/ch08/ex08 (remainingwork@GoTraining)
$  go run ../ex03/netchat.go
2016/08/17 15:51:29 Time out
2016/08/17 15:51:29 done
budougumi0617@~/git/gotraining/ch08/ex08 (remainingwork@GoTraining)
````
