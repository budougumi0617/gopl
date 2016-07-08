# Exercise 8.13
Make the chat server disconnect idle clients, such as those that have sent no messages in the last five minutes. Hint: calling `conn.Close()` in another goroutine unblocks active `Read` calls such as the one done by `input.Scan()`.

---
# 練習問題 8.13
五分間、何もメッセージを送ってこない休眠状態のクライアントの接続をチャットサーバが切断するようにしなさい。ヒント:別のゴルーチンで`conn.Close()`を呼び出すことで、`input.Scan()`によって行われているような実行中の`Read`呼び出しの待ちを解除します。
