# Exercise 9.1
Add a function `Withdraw(amount int) bool` to the `gopl.io/ch9/bank1` program. The result should indicate whether the transaction succeeded of failed due to insufficient funds. The message sent to the monitor goroutine must contain both the amount to withdraw and a new channel over which the monitor goroutine can send the boolean result back to `Withdraw`.

---
# 練習問題 9.1
関数`Withdraw(amount int) bool`を`gopl.io/ch9/bank1`プログラムに追加しなさい。結果としては、取引が成功したか、残高不足で失敗したかを示すべきです。モニターゴルーチンへ送信されるメッセージには、引き出し金額およびモニターゴルーチンがブーリアンの結果を`Withdraw`へ送り返すことができる新たなチャネルが含まれていなければなりません。
