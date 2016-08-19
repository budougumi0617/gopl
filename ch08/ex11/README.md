# Exercise 8.11
Following the approach of `mirroredQuery` in Section 8.4.4. implement a variant of `fetch` that requests several URLs concurrently. As soon as the first response arrives, cancel the other requests.

---
# 練習問題 8.11
8.4.4節の`mirroredQuery`のやり方に従って、複数のURLへ平行してリクエストを行う`fetch`の変形を実装しなさい。最初のレスポンスが到着したら、すぐに他のリクエストをキャンセルしなさい。


# Result

````shel
budougumi0617@~/git/gotraining/ch08/ex11 (remainingwork@GoTraining)
$  go run fetch.go http://google.com http://ssagesa http://qiita.com http://www.yahoo.co.jp
Failed fetch Get http://ssagesa: dial tcp: lookup ssagesa: no such host
Got http://qiita.com
Failed fetch Get http://www.google.co.jp/?gfe_rd=cr&ei=N-m2V86IL67f8Afkw6RA: net/http: request canceled while waiting for connection
Save file index.html, Size 17229
````
