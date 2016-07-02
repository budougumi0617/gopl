# Exercise 9.3
Extend the `Func` type and the `(*Memo).Get` method so that callers may provide an optional `done` channel through which they can cancel the operation (Section 8.9). The results of a cancelled `Func` call should not be cached.

---
# 練習問題 9.3
`Func`型と`(*Memo).Get`メソッドを拡張して、呼び出しもとがオプションの`done`チャネルを渡して、そのチャネルを介して操作をキャンセルできるようにしなさい（8.9節）。キャンセルされた`Func`呼び出しの結果はキャッシュされるべきではありません。
