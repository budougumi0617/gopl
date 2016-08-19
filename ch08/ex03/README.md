# Exercise 8.3
In `netcat3`, the interface value `conn` has the concrete type `*net.TCPConn`, which represent a TCP connection. A TCP connection consists of two halves that may be closed independently using `CloseRead` and `CloseWrite` methods. Modify the main goroutine of  of `netcat3` to close only the write half of the connection so that the program will continue to print the final echoes from the `reverb1` server even after the standard input has been closed.(Doing this for the `reverb2` server is harder; see Exercise 8.4)

---
# 練習問題 8.3
`netcat3`では、インターフェース値`conn`はTCP接続を表す具象型`*net.TCPConn`を持っています。TCP接続は、その`CloseRead`と`CloseWrite`のメソッドを使って、独立に閉じることができる送信部分と受信部分から構成されています。接続の送信部分だけを閉じるように`netcat3`のメインゴルーチンを修正して、標準入力が閉じられた後でもプログラムが`reverb1`サーバからの最後のエコーの表示を続けれられるようにしなさい。（`reverb2`サーバに対してこの修正を行うほうが難しいです。練習問題8.4を参照してください。）


# Result

````shell
budougumi0617@~/git/gotraining/ch08/ex03 (remainingwork@GoTraining)
$  go run ./reverb1/reverb1.go &
[1] 30024
budougumi0617@~/git/gotraining/ch08/ex03 (remainingwork@GoTraining)
$  go run netchat.go
foo
	 FOO
bar
	 foo
foo
	 foo
	 BAR
bar
	 bar
^D       bar
	 FOO
	 foo
	 foo
	 BAR
	 bar
	 bar
2016/08/17 14:23:54 done
````
