# Exercise 8.15
Failure of any client program to read data in a timely manner ultimately causes all clients to get stuck. Modify the broadcaster to skip a message rather than wait if a client writer is not ready to accept it. Alternatively, add buffering to each client's outgoing message channel so that most message are not dropped; the broadcaster should use a non-blocking send to this channel.

---
# 練習問題 8.
