# Exercise 8.4
Modify the `reverb2` server to use a `sync`. `WaitGroup` per connection to count the number of active `echo` goroutines. When it falls to zero, close the write half of the TCP connection as in Exercise 8.3. Verify that your modified `netcat3` client from that exercise waits for the final echoes of multiple concurrent shouts, even after the standard input has been closed.

---
# 練習問題 8.4
活動している`echo`ゴルーチンを数えるために接続ごとに`sync.WaitGroup`を使うように`reverb2`サーバを修正しなさい。ゼロになったら、練習問題8.3で説明されているようにTCP接続の書き込み側を閉じなさい。その練習問題で修正した`netcat3`クライアントは、標準入力が閉じられた後でも、複数の平行な叫びの最後のエコーを持つことを検証しなさい。

# Result

````shell
budougumi0617@~/git/gotraining/ch08/ex04 (remainingwork@GoTraining)
$  go run reverb2.go &
[1] 31998
budougumi0617@~/git/gotraining/ch08/ex04 (remainingwork@GoTraining)
$  go run ../ex03/netchat.go
foo
	 FOO
bar
	 BAR
	 foo
foo
	 FOO
	 bar
b	 foo
ar
	 BAR
	 foo
	 bar
^D       bar
	 foo
	 bar
````
