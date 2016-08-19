# Exercise 8.6
Add depth-limiting too the concurrent crawler. That is, if the user sets `-depth=3`, then only URLs reachable by at most three links will be fetched.

---
# 練習問題 8.6
平行なクローラに深さ制限を追加しなさい。すなわち、ユーザーが`-depth=3`と指定したら、たかだか三つのリンクをたどって到達可能なURLだけを取得します。

# Result

````shel
budougumi0617@~/git/gotraining/ch08/ex06 (remainingwork@GoTraining)
$  go run findlinks.go links.go -depth 1 http://gopl.io
depth 0, url http://www.informit.com/store/go-programming-language-9780134190440
depth 0, url http://www.amazon.com/dp/0134190440
depth 0, url http://www.informit.com/store/go-programming-language-9780134190440
depth 0, url http://www.barnesandnoble.com/w/1121601944
depth 0, url http://www.gopl.io/ch1.pdf
depth 0, url http://www.gopl.io/ch1.pdf
depth 0, url http://www.gopl.io/ch1.pdf
depth 0, url http://www.gopl.io/ch1.pdf
depth 0, url https://github.com/adonovan/gopl.io/
depth 0, url http://www.gopl.io/reviews.html
depth 0, url http://www.gopl.io/translations.html
depth 0, url http://www.gopl.io/errata.html
depth 0, url http://golang.org/s/oracle-user-manual
depth 0, url http://golang.org/lib/godoc/analysis/help.html
depth 0, url https://github.com/golang/tools/blob/master/refactor/eg/eg.go
depth 0, url https://github.com/golang/tools/blob/master/refactor/rename/rename.go
depth 0, url http://www.amazon.com/dp/0131103628?tracking_id=disfordig-20
depth 0, url http://www.amazon.com/dp/020161586X?tracking_id=disfordig-20
````
