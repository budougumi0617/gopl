# Exercise 5.4
Extend the `visit` function so that it extracts other kinds of links from the document, such as images, scripts, and style sheets.

---
# 練習問題 5.4
`visit`関数を拡張して、画像、スクリプト、スタイルシートなどのほかの種類のリンクをドキュメントから抽出するようにしなさい。


# Result

````
$  go run ../../ch01/ex07/fetch.go https://golang.org | go run findlinks.go
link href:/lib/godoc/style.css
link href:/opensearch.xml
link href:/lib/godoc/jquery.treeview.css
a href:/
a href:/
a href:#
a href:/doc/
a href:/pkg/
a href:/project/
a href:/help/
a href:/blog/
a href:http://play.golang.org/
a href:#
a href:#
a href://tour.golang.org/
a href:https://golang.org/dl/
a href://blog.golang.org/
a href:https://developers.google.com/site-policies#restrictions
a href:/LICENSE
a href:/doc/tos.html
a href:http://www.google.com/intl/en/policies/privacy/
script src:https://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js
script src:/lib/godoc/jquery.treeview.js
script src:/lib/godoc/jquery.treeview.edit.js
script src:/lib/godoc/playground.js
script src:/lib/godoc/godocs.js
````
