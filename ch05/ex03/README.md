# Exercise 5.3
Write a function to print the contents of all text nodes in an HTML document tree. Do not descend int `<script>` or `<style>` elements, since their contents are not visible in a web browser.

---
# 練習問題 5.3
HTMLドキュメントツリー内のすべてのテキストノードの内容を表示する関数を書きなさい。ウェブブラウザでは内容が表示されない`<script>`要素と`<style>`要素の中は調べないようにしなさい。


# Result

````
$  go run ../../ch01/ex07/fetch.go https://golang.org | go run showtextnode.go

The Go Programming Language
...
The Go Programming Language
Go
▽
Documents
Packages
The Project
Help
Blog
Play

package main
import "fmt"
func main() {
	fmt.Println("Hello, 世界")
}



Run

Format


Share


Pop-out
Try Go
// You can edit this code!
// Click here and start typing.
package main
import "fmt"
func main() {
	fmt.Println("Hello, 世界")
}
Hello, 世界
Run
Share
Tour

Hello, World!

Conway's Game of Life

Fibonacci Closure

Peano Integers

Concurrent pi

Concurrent Prime Sieve

Peg Solitaire Solver

Tree Comparison
Go is an open source programming language that makes it easy to build
simple, reliable, and efficient software.
Download Go
Binary distributions available for
Linux, Mac OS X, Windows, and more.
Featured video
Featured articles
Read more
Build version go1.6.
Except as
noted
,
the content of this page is licensed under the
Creative Commons Attribution 3.0 License,
and code is licensed under a
BSD license
.
Terms of Service
 |
Privacy Policy
````
