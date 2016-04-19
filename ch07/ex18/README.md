# Exercise 7.18
Using the token-based decoder API, write a program that will read an arbitrary XML document and construct a tree of generic nodes that represents it. Nodes are of two kinds: `CharData` nodes represent text strings, and `Element` nodes represent named elements and their attributes. Each element node has a slice of child nodes.
You may find the following declarations helpful.


---
# 練習問題 7.18
トークンに基づくデコーダのAPIを使用して、任意のXMLのドキュメントを読み込んで、そのドキュメントを表す総称的なノードのツリーを構築するプログラムを書きなさい。ノードには二種類あり、`CharData`ノードはテキスト文字列を表し、`Element`ノードは名前付き要素とその属性を表します。それぞれの要素のノードは子ノードのスライスを持ちます。

````go
import "encoding/xml"

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
  Type     xml.Name
  Attr     []xml.Attr
  Children []Node
}
````
