# Exercise 12.8
The `sexpr.Unmarshal` function, like `json.Marshal`, requires the complete input in a byte slice before it can begin decoding. Define a `sexpr.Decoder` type that, like `json.Decoder`, allows a sequence of values to be decoded from an `io.Reader`. Change `sexpr.Unmarshal` to use this new type.

---
# 練習問題 12.8
`sexpr.Unmarshal`関数は、`json.Marshal`のようにデコードを開始する前にバイトスライスの形で完全な入力を必要とします。`json.Decoder`のように、`io.Reader`からデコードされる値の列を許す`sexpr.Decoder`型を定義しなさい。その新たな型を使うように`sexpr.Unmarshal`を変更しなさい。
