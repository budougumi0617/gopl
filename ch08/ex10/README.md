# Exercise 8.10
HTTP requests may be cancelled by closing the optional `Cancel` channel in the `http.Request` struct. Modify the web crawler of Section 8.6 to support cancellation.
Hint: the `http.Get` convenience function does not give you an opportunity to customize a `Request`. instead, create the request using `http.NewRequest`, set its `Cancel` field, than perform the request by calling `http.DefaultClient.Do(req)`.

---
# 練習問題 8.
