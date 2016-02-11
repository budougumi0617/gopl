# Exercise 1.9
Modify `fetch` to also print the HTTP status code, found in `resp.Status`.

---
# 練習問題1.9
`fetch`を修正して、`resp.Status`に設定されているHTTPステータスコードも表示するようにしなさい。


# Result

```sh
$  go run fetch.go http://google.com
[http://google.com]resp.status 200 OK
<!doctype html><html itemscope="" itemtype="http://schema.org/WebPage" lang="ja"><head><meta content="���E������������邽�߂̃c�[����񋟂��Ă��܂��B���܂��܂Ȍ���@�\��p���āA���T��������Ă��������B" name=...

$  go run fetch.go http://google.com/foo
[http://google.com/foo]resp.status 404 Not Found
<html><body><h1>404 Not Found</h1></body></html>
```
