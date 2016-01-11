# Exercise 1.12
Modify the lissajous server to read parameter values from the URL like `http://localhost:800/?cycles=20` sets the number of cycles to 20 instead of the default 5. Use the `strconv.Atoi` function to convert the string parameter into an integer. You can see its documentation with `go doc strconv.Atoi`.

---
# 練習問題1.12
リサジュー図形のサーバを修正して、URLからパラメータ値を読み取るようにしなさい。たとえば、`http://localhost:800/?cycles=20`などのURLによって、サイクル数をデフォルトの5ではなく20に設定するようにしなさい。文字列パラメータを整数へ変換するために`strconv.Atoi`関数を使用しなさい。その変換関数のドキュメントは`go doc strconv.Atoi`で見ることができます。
