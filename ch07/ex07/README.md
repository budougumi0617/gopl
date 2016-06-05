# Exercise 7.7
Explain why the help message contains °C when the default value of 20.0 does not.

---
# 練習問題 7.7
`20.0`のデフォルト値は°Cを含んでいないのに、ヘルプメッセージが°Cを含んでいる理由を説明しなさい。

# Answer
`String()` of `Celsius` add `°C`, and help message shows `Celsius` value. As a result, help message also contains `°C`.

````go
func (c Celsius) String() string { return fmt.Sprintf("%g°C added by Celsius.String()", c) }
````


````sh
budougumi0617@~/git/gotraining/ch07/ex07 (ch07@GoTraining)
$  go test
20°C added by Celsius.String()
PASS
ok  	github.com/budougumi0617/GoTraining/ch07/ex07	0.013s
````
