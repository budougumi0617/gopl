# Exercise 8.15
Failure of any client program to read data in a timely manner ultimately causes all clients to get stuck. Modify the `broadcaster` to skip a message rather than wait if a client writer is not ready to accept it. Alternatively, add buffering to each client's outgoing message channel so that most message are not dropped; the `broadcaster` should use a non-blocking send to this channel.

---
# 練習問題 8.15
タイミングよくクライアントのプログラムがデータの読み込みを行わないと、すべてのクライアントの動作が止まってしまいます。`clientWriter`がメッセージを受け付ける準備ができるのを待つのではなく、メッセージをスキップするように`broadcaster`を修正しなさい。あるいは、個々のクライアントの送信用メッセージチャネルにバッファを追加して、ほとんどのメッセージが失われないようにしなさい。`broadcaster`は、そのチャネルに対して待たされない送信を使うべきです。
