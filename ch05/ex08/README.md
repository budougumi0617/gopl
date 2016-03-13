# Exercise 5.

---
# 練習問題 5.8
走査を続けるか否かを示すブーリアンの結果を`pre`関数と`post`関数が返すようにして、それに対応するように`forEachNode`を修正しなさい。修正した`forEachNode`を使用して、指定された`id`属性を持つ最初のHTML要素を見つけるような下記のシグニチャの関数`ElementByID`を書きなさい。`ElementByID`関数は、一致が見つかったら走査を中止しなければなりません。


`func ElementByID(doc *html.Node, id string) *html.Node`
