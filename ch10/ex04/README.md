# Exercise 10.4
Construct a tool that reports the set of all packages in the workspace that transitively depend on the packages specified by the arguments. Hint: you will need to run `go list` twice, once for the initial packages and once for all packages. You may want to parse its JSON output using the `encoding/json` package(Sec 4.5).

---
# 練習問題 10.4
引数で指定されたパッケージに推移的に依存している、ワークスペース内のすべてのパッケージの集まりを報告するツールを作成しなさい。ヒント: `go list`を二回実行する必要があります。最初のパッケージに対して一回、すべてのパッケージに対して一回です。`encoding/json`パッケージ(4.5節)を使ってJSONの出力を解析するとよいでしょう。


# Result

````bash
budougumi0617@~/go/src/github.com/budougumi0617/GoTraining/ch10/ex04 (ch10ch11@GoTraining)
$  go run search_dependpackage.go github.com/budougumi0617/GoTraining/ch11/ex07
[bytes errors fmt internal/race io math os reflect runtime runtime/internal/atomic runtime/internal/sys sort strconv sync sync/atomic syscall time unicode unicode/utf8 unsafe]
````
