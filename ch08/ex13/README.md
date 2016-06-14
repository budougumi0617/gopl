# Exercise 8.13
Make the chat server disconnect idle clients, such as those that have sent no messages in the last five minutes. Hint: calling `conn.Close()` in another goroutine unblocks active `Read` calls such as the one done by `input.Scan()`.

---
# 練習問題 8.
