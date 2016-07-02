# Exercise 8.10
HTTP requests may be cancelled by closing the optional `Cancel` channel in the `http.Request` struct. Modify the web crawler of Section 8.6 to support cancellation.
Hint: the `http.Get` convenience function does not give you an opportunity to customize a `Request`. instead, create the request using `http.NewRequest`, set its `Cancel` field, than perform the request by calling `http.DefaultClient.Do(req)`.

---
# 練習問題 8.10
HTTPリクエストは、`http.Request`構造体のオプションの`Cancel`チャネルを閉じることでキャンセルできます。キャンセルをサポートするように8.6節のウェブクローラを修正しなさい。

ヒント: `http.Get`は便利な関数ですが、`Request`をカスタマイズする機会を与えません。代わりに`http.NewRequest`を使ってリクエストを作成し、その`Cancel`フィールドを設定します。それから`http.DefaultClient.Do(req)`を呼び出してリクエストを実行しなさい。
