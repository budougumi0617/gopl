# Exercise 8.10
HTTP requests may be cancelled by closing the optional `Cancel` channel in the `http.Request` struct. Modify the web crawler of Section 8.6 to support cancellation.
Hint: the `http.Get` convenience function does not give you an opportunity to customize a `Request`. instead, create the request using `http.NewRequest`, set its `Cancel` field, than perform the request by calling `http.DefaultClient.Do(req)`.

---
# 練習問題 8.10
HTTPリクエストは、`http.Request`構造体のオプションの`Cancel`チャネルを閉じることでキャンセルできます。キャンセルをサポートするように8.6節のウェブクローラを修正しなさい。

ヒント: `http.Get`は便利な関数ですが、`Request`をカスタマイズする機会を与えません。代わりに`http.NewRequest`を使ってリクエストを作成し、その`Cancel`フィールドを設定します。それから`http.DefaultClient.Do(req)`を呼び出してリクエストを実行しなさい。



````shel
budougumi0617@~/git/gotraining/ch08/ex10 (remainingwork@GoTraining)
$  go run findlinks.go http://google.com http://ssagesa http://qiita.com http://www.yahoo.co.jp
2016/08/20 02:02:15 Get http://ssagesa: dial tcp: lookup ssagesa: no such host
2016/08/20 02:02:15 Got link in http://ssagesa
2016/08/20 02:02:15 Got link in http://www.yahoo.co.jp
2016/08/20 02:02:15 Got link in http://news.yahoo.co.jp/pickup/6211681
2016/08/20 02:02:15 Got link in http://news.yahoo.co.jp/pickup/6211686
2016/08/20 02:02:15 Got link in http://news.yahoo.co.jp/pickup/6211673
2016/08/20 02:02:15 Got link in http://news.yahoo.co.jp/pickup/6211663
2016/08/20 02:02:15 Got link in http://news.yahoo.co.jp/pickup/6211683
2016/08/20 02:02:15 Got link in http://news.yahoo.co.jp/pickup/6211688
2016/08/20 02:02:15 Got link in http://news.yahoo.co.jp/pickup/6211679
2016/08/20 02:02:15 Got link in http://news.yahoo.co.jp/list/?d=20160820&mc=f&mp=f
2016/08/20 02:02:15 Got link in http://news.yahoo.co.jp/pickup/6211687
2016/08/20 02:02:15 Got link in http://news.yahoo.co.jp/fc
2016/08/20 02:02:15 Got link in http://promo.search.yahoo.co.jp/kcoupon/?push=3
2016/08/20 02:02:15 Got link in http://auctions.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://rdsig.yahoo.co.jp/adpromo/ramen2016/toplink/RV=1/RU=aHR0cDovL3JhbWVuLnlhaG9vLmNvLmpwLw--
2016/08/20 02:02:15 Got link in http://yahoo-mbga.jp/?_ref=aff%3Dysm001
2016/08/20 02:02:15 Got link in http://qiita.com
2016/08/20 02:02:15 Got link in http://tabelog.com/
2016/08/20 02:02:15 Got link in http://news.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://google.com
2016/08/20 02:02:15 Got link in http://realestate.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://rio.headlines.yahoo.co.jp/rio/gallery/detail?a=20160820-00010000-yrioolym-spo.view-000
2016/08/20 02:02:15 Got link in http://map.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://transit.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://loco.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://textream.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://weather.yahoo.co.jp/weather/
2016/08/20 02:02:15 Got link in http://m.baseball.yahoo.co.jp/hsb/
2016/08/20 02:02:15 Got link in http://finance.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://travel.yahoo.co.jp/?sc_e=ytc
2016/08/20 02:02:15 Got link in http://tv.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://services.yahoo.co.jp/?mode=pc
2016/08/20 02:02:15 Got link in http://shopping.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://bookstore.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://help.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://beauty.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://promo.mail.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://indival.yahoo.co.jp
2016/08/20 02:02:15 Got link in http://sports.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://blogs.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://calendar.yahoo.co.jp/info/guide/z/
2016/08/20 02:02:15 Got link in http://carview.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://movies.yahoo.co.jp/
2016/08/20 02:02:15 Got link in http://box.yahoo.co.jp
2016/08/20 02:02:15 Got link in http://ir.yahoo.co.jp/
2016/08/20 02:02:16 Got link in http://csr.yahoo.co.jp/
2016/08/20 02:02:16 Got link in http://docs.yahoo.co.jp/info/charter/
2016/08/20 02:02:16 Got link in https://login.yahoo.co.jp/config/login?.src=www&.done=http://www.yahoo.co.jp
2016/08/20 02:02:16 Got link in http://partner.yahoo.co.jp/
2016/08/20 02:02:16



---------Canceled



2016/08/20 02:02:16 Get http://gyao.yahoo.co.jp/: net/http: request canceled
2016/08/20 02:02:16 Get http://news.search.yahoo.co.jp/advanced: net/http: request canceled while waiting for connection
2016/08/20 02:02:16 Got link in http://news.search.yahoo.co.jp/advanced
---------------Canceled in crawler 0
2016/08/20 02:02:16 Get http://docs.yahoo.co.jp/docs/info/terms/: net/http: request canceled
2016/08/20 02:02:16 Got link in http://docs.yahoo.co.jp/docs/info/terms/
---------------Canceled in crawler 19
2016/08/20 02:02:16 Get http://security.yahoo.co.jp/: net/http: request canceled
2016/08/20 02:02:16 Got link in http://security.yahoo.co.jp/
---------------Canceled in crawler 12
2016/08/20 02:02:16 Get http://hr.yahoo.co.jp/: net/http: request canceled
2016/08/20 02:02:16 Got link in http://hr.yahoo.co.jp/
2016/08/20 02:02:16 Get http://news.yahoo.co.jp/pickup/6211681#yjContentsStart: net/http: request canceled while waiting for connection
2016/08/20 02:02:16 Got link in http://news.yahoo.co.jp/pickup/6211681#yjContentsStart
---------------Canceled in crawler 18
2016/08/20 02:02:16 Get https://account.edit.yahoo.co.jp/registration?.src=www&.done=http%3A%2f%2fwww.yahoo.co.jp: net/http: request canceled
2016/08/20 02:02:16 Got link in https://account.edit.yahoo.co.jp/registration?.src=www&.done=http%3A%2f%2fwww.yahoo.co.jp
---------------Canceled in crawler 9
2016/08/20 02:02:16 Get http://points.yahoo.co.jp/: net/http: request canceled while waiting for connection
2016/08/20 02:02:16 Got link in http://points.yahoo.co.jp/
---------------Canceled in crawler 5
2016/08/20 02:02:16 Got link in http://gyao.yahoo.co.jp/
---------------Canceled in crawler 1
2016/08/20 02:02:16 Get https://login.yahoo.co.jp/config/login_verify2?.src=www&.done=https://lh.login.yahoo.co.jp/%3F.done%3Dhttp%3A%2F%2Fwww.yahoo.co.jp%2F: net/http: request canceled while waiting for connection
2016/08/20 02:02:16 Got link in https://lh.login.yahoo.co.jp/
---------------Canceled in crawler 2
2016/08/20 02:02:16 Get https://login.yahoo.co.jp/config/login?.src=www&.done=https%3A%2F%2Faccount.edit.yahoo.co.jp%2Feval_profile%3F.src%3Dwww%26.done%3Dhttp%253A%252F%252Fwww.yahoo.co.jp%252F&t_cushion=1: net/http: request canceled while waiting for connection
2016/08/20 02:02:16 Got link in https://account.edit.yahoo.co.jp/eval_profile?.src=www&.done=http%3A%2f%2fwww.yahoo.co.jp%2F
---------------Canceled in crawler 15
2016/08/20 02:02:16 Get http://docs.yahoo.co.jp/info/mediastatement/: net/http: request canceled
2016/08/20 02:02:16 Got link in http://docs.yahoo.co.jp/info/mediastatement/
---------------Canceled in crawler 8
2016/08/20 02:02:16 Get http://docs.yahoo.co.jp/docs/info/terms/chapter1.html#cf2nd: net/http: request canceled
2016/08/20 02:02:16 Got link in http://docs.yahoo.co.jp/docs/info/terms/chapter1.html#cf2nd
---------------Canceled in crawler 17
2016/08/20 02:02:16 Get https://login.yahoo.co.jp/config/login?.src=ym&.done=http%3A%2F%2Fmail.yahoo.co.jp%2F: net/http: request canceled while waiting for connection
2016/08/20 02:02:16 Got link in http://mail.yahoo.co.jp/
---------------Canceled in crawler 16
2016/08/20 02:02:16 Get http://docs.yahoo.co.jp/info/: net/http: request canceled
2016/08/20 02:02:16 Got link in http://docs.yahoo.co.jp/info/
---------------Canceled in crawler 14
2016/08/20 02:02:16 Get https://login.yahoo.co.jp/config/login?auth_lv=pw&.src=yc&.done=http%3A%2F%2Fcalendar.yahoo.co.jp%2F: net/http: request canceled
2016/08/20 02:02:16 Got link in http://calendar.yahoo.co.jp/
---------------Canceled in crawler 7
2016/08/20 02:02:16 Get http://marketing.yahoo.co.jp/: net/http: request canceled while waiting for connection
2016/08/20 02:02:16 Got link in http://advertising.yahoo.co.jp/
---------------Canceled in crawler 13
2016/08/20 02:02:16 Get http://docs.yahoo.co.jp/docs/pr/disclaimer.html: net/http: request canceled
2016/08/20 02:02:16 Got link in http://docs.yahoo.co.jp/docs/pr/disclaimer.html
---------------Canceled in crawler 4
2016/08/20 02:02:16 Get http://fortune.yahoo.co.jp/: net/http: request canceled
2016/08/20 02:02:16 Got link in http://fortune.yahoo.co.jp/fortune/
---------------Canceled in crawler 6
2016/08/20 02:02:16 Get http://www.yahoo-help.jp/app/answers/detail/p/533/a_id/43883: net/http: request canceled
2016/08/20 02:02:16 Got link in http://www.yahoo-help.jp/app/answers/detail/p/533/a_id/43883
---------------Canceled in crawler 11
2016/08/20 02:02:16 parsing http://games.yahoo.co.jp/ as HTML: net/http: request canceled
2016/08/20 02:02:16 Got link in http://games.yahoo.co.jp/
---------------Canceled in crawler 10
2016/08/20 02:02:16 Get http://person.news.yahoo.co.jp/u/login: net/http: request canceled while waiting for connection
2016/08/20 02:02:16 Got link in http://person.news.yahoo.co.jp/u/login
---------------Canceled in crawler 3
budougumi0617@~/git/gotraining/ch08/ex10 (remainingwork@GoTraining)
````
