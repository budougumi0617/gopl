# Exercise 1.8
Modify `fetch` to add the prefix `http://` to each argument URL if it is missing. You might want to use `strings.HasPrefix`.

---
# 練習問題1.8
`fecth`を修正して、指定されていない場合には`http://`の接頭辞を追加するようにしなさい。`strings.HasPrefix`を使用したくなるかもしれません。


# Result

```sh
$  go run fetch.go http://google.com/foo
<html><body><h1>404 Not Found</h1></body></html>

$  go run fetch.go google.com/foo
<html><body><h1>404 Not Found</h1></body></html>
```
