# Exercise 7.11
Add additional handlers so that clients can create, read, update, and delete database entries. For example, a request of the form `/update?item=socks&price=6` will update the price of an item int the inventory and report an error if the item does not exist of the price is invalid. (Warning: this change introduces concurrent variable updates.)

---
# 練習問題 7.11
データベースのエントリをクライアントが作成、読み出し、更新、削除できるようにハンドラを追加しなさい。たとえば、フォーム`/update?item=socks&price=6`形式のリクエストは、商品全体の中の一つの商品の価格を更新し、その商品がないもしくは価格が不正であればエラーを報告します。（警告:この変更は、変数の並行な更新を発生させます。
