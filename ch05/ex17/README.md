# Exercise 5.17
Write a variadic function `ElementByTagName` that given an HTML node tree and zero or moe names, returns all the elements that match one of those names. Here are two example calls:

---
# 練習問題 5.17
HTMLノードツリーとゼロ個以上の名前が与えられたら、それらの名前のひとつと一致する要素をすべてを返す可変個引数関数`ElementByTagName`を書きなさい。二つの呼び出し例を次に示します。

````go
func ElementByTagName(doc *html.Node, name ...string) []*html.Node

images := ElementByTagName(doc, "img")
headings := ElementByTagName(doc, "h1", "h2", "h3", "h4")

````


# Result

````
$  go run elementbytagname.go http://github.com
img
No.0 is img
	"src" element, value is "https://assets-cdn.github.com/images/modules/site/home-ill-build.png?sn"
	"alt" element, value is ""
	"class" element, value is "img-responsive featurette-illo-sm mb-4 mt-4"
No.1 is img
	"src" element, value is "https://assets-cdn.github.com/images/modules/site/home-ill-work.png?sn"
	"alt" element, value is ""
	"class" element, value is "img-responsive featurette-illo-sm mb-4 mt-4"
No.2 is img
	"src" element, value is "https://assets-cdn.github.com/images/modules/site/home-ill-projects.png?sn"
	"alt" element, value is ""
	"class" element, value is "img-responsive featurette-illo-sm mb-4 mt-4"
No.3 is img
	"src" element, value is "https://assets-cdn.github.com/images/modules/site/home-ill-platform.png?sn"
	"alt" element, value is ""
	"class" element, value is "img-responsive featurette-illo-sm mb-4 mt-4"
No.4 is img
	"src" element, value is "https://assets-cdn.github.com/images/modules/site/org_example_nasa.png?sn"
	"class" element, value is "img-responsive org-example-drop-shadow"
	"alt" element, value is "NASA is on GitHub"
headings
No.0 is h1
	"class" element, value is "display-heading-1 jumbotron-title"
No.1 is h2
	"class" element, value is "featurette-heading display-heading-2 text-center"
No.2 is h4
	"class" element, value is "display-heading-4"
No.3 is h4
	"class" element, value is "display-heading-4"
No.4 is h4
	"class" element, value is "display-heading-4"
No.5 is h4
	"class" element, value is "display-heading-4"
No.6 is h2
	"class" element, value is "display-heading-2"
No.7 is h3
	"class" element, value is "display-heading-3 mt-2"
No.8 is h3
	"class" element, value is "display-heading-3 mt-2"
No.9 is h3
	"class" element, value is "display-heading-3 mt-2"
````
