# Exercise 5.14
Use the `breadFirst` function to explore a different structure. For example, you could use the course dependencies from the `topoSort` example (a directed graph), the file system hierarchy downloaded from your city government's web site (an undirected graph).

---
# 練習問題 5.14
異なる構造を調べるために`breadFirst`関数を使用しなさい。たとえば、`topoSort`の例（有向グラフ）の講座の依存関係、コンピュータ上のファイルシステムの階層、公共機関のウェブサイトからダウンロードしたバスや地下鉄の経路（無向グラフ）のリストなどを利用できます。


# Result

````
$  go run breadthfirst.go
---------------------
"algorithms" is requested to already study below corses,
"data structures",
"discrete math",
"intro to programming"
---------------------
"calculus" is requested to already study below corses,
"linear algebra"
---------------------
"networks" is requested to already study below corses,
"operating systems",
"data structures",
"computer organization",
"discrete math",
"intro to programming"
---------------------
"operating systems" is requested to already study below corses,
"data structures",
"computer organization",
"discrete math",
"intro to programming"
---------------------
"programming languages" is requested to already study below corses,
"data structures",
"computer organization",
"discrete math",
"intro to programming"
---------------------
"compilers" is requested to already study below corses,
"data structures",
"formal languages",
"computer organization",
"discrete math",
"intro to programming"
---------------------
"data structures" is requested to already study below corses,
"discrete math",
"intro to programming"
---------------------
"databases" is requested to already study below corses,
"data structures",
"discrete math",
"intro to programming"
---------------------
"discrete math" is requested to already study below corses,
"intro to programming"
---------------------
"formal languages" is requested to already study below corses,
"discrete math",
"intro to programming"
````
