# Exercise 9.2
Rewrite the `PopCount` example from Section 2.6.2 so that it initializes the lookup table using `sync.Once` the first time it is needed. (Realistically, the cost of synchronization would be prohibitive for a small and highly optimized function like `PopCount`.)

---
# 練習問題 9.2
2.6.2節の`PopCount`のコード例を書き直し、`sync.Once`を使ってルックアップテーブルが必要になった最初のときに初期化しなさい。（現実的には、`PopCount`のような小さく高度に最適化された関数では同期のコストは非常に高いでしょう。）
