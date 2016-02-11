# Exercise 1.11
Try `fetchall` with longer arugment lists, such as samples from the top million web sites available at `alexa.com`. How does the program behave if a web site just doesn't respond? (Section 8.9 describes mechanisms for coping in such cases.)

---
# 練習問題1.11
`alexa.com`にあるトップウェブサイトを参考に、より長い引数リストで`fetchall`を試しなさい。あるウェブサイトが応答しない場合には、プログラムはどのように振る舞うでしょうか（そのような場合に対処するための機構は8.9節で説明されています）。


# Result
応答しない場合を再現できませんでした。

```sh
$  go run fetchall.go http://google.com http://facebook.com http://youtube.com http://baidu.com http://amazon.com http://wikipedia.org http://qq.com http://google.co.in http://twitter.com http://live.com http://taobao.com http://sina.com.cn http://msn.com http://yahoo.co.jp http://google.co.jp http://linkedin.com http://vk.com http://bing.com http://yandex.ru http://hao123.com http://ebay.com http://instagram.com http://amazon.co.jp http://google.de http://mail.ru http://360.cn http://netflix.com http://t.co http://google.co.uk http://pinterest.com http://tmail.com http://google.ru http://google.com.br http://raddit.com
0.16s    19198  http://yahoo.co.jp
0.18s    19298  http://google.com
Get http://tmail.com: dial tcp: lookup tmail.com: no such host
0.20s     1559  http://raddit.com
0.22s    19274  http://google.co.jp
0.23s    19551  http://google.com.br
0.26s     3154  http://t.co
0.45s   111149  http://360.cn
0.58s    19045  http://google.co.uk
0.60s    19041  http://google.de
0.61s    20763  http://google.co.in
0.90s   529455  http://sina.com.cn
0.91s    19680  http://google.ru
0.92s    79262  http://bing.com
0.99s   210939  http://pinterest.com
1.01s   251663  http://twitter.com
1.12s   182852  http://ebay.com
1.15s    41530  http://linkedin.com
1.18s   334990  http://youtube.com
1.29s     6064  http://vk.com
1.34s    44357  http://msn.com
1.58s    54537  http://wikipedia.org
1.67s    14219  http://instagram.com
1.79s    70220  http://facebook.com
1.81s   566833  http://amazon.co.jp
1.81s   691717  http://hao123.com
1.89s   376779  http://amazon.com
2.14s    97668  http://netflix.com
2.17s     9649  http://live.com
3.40s    52855  http://yandex.ru
3.75s   255286  http://mail.ru
4.57s    70260  http://taobao.com
8.11s       81  http://baidu.com
15.64s   624571  http://qq.com
15.64s elapsed
```
