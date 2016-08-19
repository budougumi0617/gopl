# Exercise 3.4
Following the approach of the Lissajous example in Section 1.7, construct a web server that computes surfaces and writes and SVG data to the client. The server must set the `Content-Type` header like this:

`w.Header().Set("Content-Type", "image/svg+xml")`

(This step was not required in the Lissajous example because the server uses standard heuristics to recognize common formats like PNG from the first 512 bytes of the response, and generates the proper header.) Allow the client to specify values like height, width, and color as HTTP request parameters.

---
# 練習問題 3.4
1.7節のリサジューの例の方法に従って、面を計算して面を計算してSVGデータをクライアントに対して書き出すウェブサーバを作成しなさい。サーバーは次のように`Content-Type`を設定しなければなりません。

`w.Header().Set("Content-Type", "image/svg+xml")`

（このステップはリサジューの例では必要ありませんでした。それは、サーバが標準的なヒューリスティクスを使用してレスポンスの最初の512バイトからPNGなどの標準形式を認識し、適切なヘッダーを生成しているからです。）HTTPリクエストのパラメータとして、高さ、幅、色などの値をクライアントが指定できるようにしなさい。
