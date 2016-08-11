# Exercise 10.3
Using `fetch http://gopl.io/ch1/helloworld?go-get=1`, find out which service hosts the code samples for this book. (HTTP requests from `go get` include the `go-get` parameter so that servers can distinguish them from ordinary browser requests.)

---
# 練習問題 10.3
`fetch http://gopl.io/ch1/helloworld?go-get=1`を使って、この本のサンプルコードをホストしているサービスを調べなさい。(`go get`からのHTTPリクエストは`go-get`パラメータを含んでいるので、サーバは通常のブラウザのリクエストと区別することができます。)


# Result

````bash
budougumi0617@~/go/src/github.com/budougumi0617/GoTraining/ch10/ex03 (ch10ch11@GoTraining)
$  go build  gopl.io/ch1/fetch
budougumi0617@~/go/src/github.com/budougumi0617/GoTraining/ch10/ex03 (ch10ch11@GoTraining)
$  ./fetch  http://gopl.io/ch1/helloworld\?go-get\=1 | grep go-import
<meta name="go-import" content="gopl.io git https://github.com/adonovan/gopl.io">
````
