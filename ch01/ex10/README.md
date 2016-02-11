# Exercise 1.10
Find a web site that produces a large amount of data. Investigate caching by running `fetchall` twice in succession to see whether the reported time changes much. Do you get the same content each time? Modify `fetchall` to print its output to a file so it can be examined.

---
# 練習問題1.10
大量のデータを生成するウェブサイトを見つけなさい。報告される時間が大きく変化するかを調べるために`fetchall`を2回続けて実行して、キャッシュされているかを調査しなさい。毎回同じ内容を得ていますか。`fetchall`を修正して、その出力をファイルへ保存するようにして調べられるようにしなさい。


#Result
２度目のアクセスは速度向上しているので、キャッシュしていると思われる。

```sh
budougumi0617@~/git/gotraining/ch01/ex10 (issue/5@GoTraining)
$  fetchall http://www.bmw.co.jp/ja/index.html && fetchall http://www.bmw.co.jp/ja/index.html
1.69s  790623 http://www.bmw.co.jp/ja/index.html

1.69s elapsed
0.21s  790623 http://www.bmw.co.jp/ja/index.html

0.21s elapsed
budougumi0617@~/git/gotraining/ch01/ex10 (issue/5@GoTraining)
$  diff result result1
```
