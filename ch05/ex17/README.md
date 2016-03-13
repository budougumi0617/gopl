# Exercise 5.

---
# 練習問題 5.17
HTMLノードツリーとゼロ個以上の名前が与えられたら、それらの名前のひとつと一致する要素をすべてを返す可変個引数関数`ElementByTagName`を書きなさい。二つの呼び出し例を次に示します。

````go
func ElementByTagName(doc *html.Node, name ...string) []*html.Node

images := ElementByTagName(doc, "img")
headings := ElementByTagName(doc, "h1", "h2", "h3", "h4")

````
