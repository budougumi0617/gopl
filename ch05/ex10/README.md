# Exercise 5.10
Rewrite `topoSort` to use maps instead of slices and eliminate the initial sort. Verify that the results, though nondeterministic, are valid topologcal order orderings.

---
# 練習問題 5.10
スライスの代わりにマップを使用するように`topoSort`を書き直して、最初のソートを削除しなさい。結果は非決定的ですが、結果が有効なトポロジカル順序になっていることを検証しなさい。



# Result

````
$  go test
----------------------------
Try count 1
1:	intro to programming
2:	discrete math
3:	data structures
4:	databases
5:	algorithms
6:	computer organization
7:	programming languages
8:	operating systems
9:	linear algebra
10:	calculus
11:	formal languages
12:	compilers
13:	networks

Check answer
"discrete math" requests "intro to programming"
"data structures" requests "discrete math"
"databases" requests "data structures"
"algorithms" requests "data structures"
"programming languages" requests "data structures"
"programming languages" requests "computer organization"
"operating systems" requests "data structures"
"operating systems" requests "computer organization"
"calculus" requests "linear algebra"
"formal languages" requests "discrete math"
"compilers" requests "data structures"
"compilers" requests "formal languages"
"compilers" requests "computer organization"
"networks" requests "operating systems"
----------------------------
Try count 2
1:	intro to programming
2:	discrete math
3:	data structures
4:	computer organization
5:	programming languages
6:	algorithms
7:	linear algebra
8:	calculus
9:	databases
10:	formal languages
11:	operating systems
12:	networks
13:	compilers

Check answer
"discrete math" requests "intro to programming"
"data structures" requests "discrete math"
"programming languages" requests "data structures"
"programming languages" requests "computer organization"
"algorithms" requests "data structures"
"calculus" requests "linear algebra"
"databases" requests "data structures"
"formal languages" requests "discrete math"
"operating systems" requests "data structures"
"operating systems" requests "computer organization"
"networks" requests "operating systems"
"compilers" requests "computer organization"
"compilers" requests "data structures"
"compilers" requests "formal languages"
----------------------------
Try count 3
1:	intro to programming
2:	discrete math
3:	data structures
4:	algorithms
5:	computer organization
6:	operating systems
7:	formal languages
8:	compilers
9:	networks
10:	linear algebra
11:	calculus
12:	programming languages
13:	databases

Check answer
"discrete math" requests "intro to programming"
"data structures" requests "discrete math"
"algorithms" requests "data structures"
"operating systems" requests "data structures"
"operating systems" requests "computer organization"
"formal languages" requests "discrete math"
"compilers" requests "data structures"
"compilers" requests "formal languages"
"compilers" requests "computer organization"
"networks" requests "operating systems"
"calculus" requests "linear algebra"
"programming languages" requests "data structures"
"programming languages" requests "computer organization"
"databases" requests "data structures"
----------------------------
Try count 4
1:	intro to programming
2:	discrete math
3:	formal languages
4:	data structures
5:	computer organization
6:	operating systems
7:	networks
8:	programming languages
9:	algorithms
10:	linear algebra
11:	calculus
12:	compilers
13:	databases

Check answer
"discrete math" requests "intro to programming"
"formal languages" requests "discrete math"
"data structures" requests "discrete math"
"operating systems" requests "data structures"
"operating systems" requests "computer organization"
"networks" requests "operating systems"
"programming languages" requests "computer organization"
"programming languages" requests "data structures"
"algorithms" requests "data structures"
"calculus" requests "linear algebra"
"compilers" requests "data structures"
"compilers" requests "formal languages"
"compilers" requests "computer organization"
"databases" requests "data structures"
----------------------------
Try count 5
1:	intro to programming
2:	discrete math
3:	data structures
4:	algorithms
5:	formal languages
6:	computer organization
7:	compilers
8:	databases
9:	operating systems
10:	networks
11:	programming languages
12:	linear algebra
13:	calculus

Check answer
"discrete math" requests "intro to programming"
"data structures" requests "discrete math"
"algorithms" requests "data structures"
"formal languages" requests "discrete math"
"compilers" requests "data structures"
"compilers" requests "formal languages"
"compilers" requests "computer organization"
"databases" requests "data structures"
"operating systems" requests "data structures"
"operating systems" requests "computer organization"
"networks" requests "operating systems"
"programming languages" requests "data structures"
"programming languages" requests "computer organization"
"calculus" requests "linear algebra"
----------------------------
Try count 6
1:	intro to programming
2:	discrete math
3:	data structures
4:	formal languages
5:	computer organization
6:	compilers
7:	operating systems
8:	networks
9:	programming languages
10:	databases
11:	linear algebra
12:	calculus
13:	algorithms

Check answer
"discrete math" requests "intro to programming"
"data structures" requests "discrete math"
"formal languages" requests "discrete math"
"compilers" requests "data structures"
"compilers" requests "formal languages"
"compilers" requests "computer organization"
"operating systems" requests "data structures"
"operating systems" requests "computer organization"
"networks" requests "operating systems"
"programming languages" requests "computer organization"
"programming languages" requests "data structures"
"databases" requests "data structures"
"calculus" requests "linear algebra"
"algorithms" requests "data structures"
----------------------------
Try count 7
1:	intro to programming
2:	discrete math
3:	data structures
4:	formal languages
5:	computer organization
6:	programming languages
7:	compilers
8:	databases
9:	operating systems
10:	networks
11:	algorithms
12:	linear algebra
13:	calculus

Check answer
"discrete math" requests "intro to programming"
"data structures" requests "discrete math"
"formal languages" requests "discrete math"
"programming languages" requests "data structures"
"programming languages" requests "computer organization"
"compilers" requests "data structures"
"compilers" requests "formal languages"
"compilers" requests "computer organization"
"databases" requests "data structures"
"operating systems" requests "data structures"
"operating systems" requests "computer organization"
"networks" requests "operating systems"
"algorithms" requests "data structures"
"calculus" requests "linear algebra"
----------------------------
Try count 8
1:	intro to programming
2:	discrete math
3:	formal languages
4:	data structures
5:	computer organization
6:	operating systems
7:	networks
8:	programming languages
9:	databases
10:	linear algebra
11:	calculus
12:	compilers
13:	algorithms

Check answer
"discrete math" requests "intro to programming"
"formal languages" requests "discrete math"
"data structures" requests "discrete math"
"operating systems" requests "data structures"
"operating systems" requests "computer organization"
"networks" requests "operating systems"
"programming languages" requests "data structures"
"programming languages" requests "computer organization"
"databases" requests "data structures"
"calculus" requests "linear algebra"
"compilers" requests "data structures"
"compilers" requests "formal languages"
"compilers" requests "computer organization"
"algorithms" requests "data structures"
----------------------------
Try count 9
1:	intro to programming
2:	discrete math
3:	data structures
4:	computer organization
5:	operating systems
6:	formal languages
7:	networks
8:	programming languages
9:	compilers
10:	databases
11:	algorithms
12:	linear algebra
13:	calculus

Check answer
"discrete math" requests "intro to programming"
"data structures" requests "discrete math"
"operating systems" requests "data structures"
"operating systems" requests "computer organization"
"formal languages" requests "discrete math"
"networks" requests "operating systems"
"programming languages" requests "data structures"
"programming languages" requests "computer organization"
"compilers" requests "formal languages"
"compilers" requests "computer organization"
"compilers" requests "data structures"
"databases" requests "data structures"
"algorithms" requests "data structures"
"calculus" requests "linear algebra"
----------------------------
Try count 10
1:	intro to programming
2:	discrete math
3:	data structures
4:	databases
5:	algorithms
6:	formal languages
7:	computer organization
8:	compilers
9:	operating systems
10:	linear algebra
11:	calculus
12:	networks
13:	programming languages

Check answer
"discrete math" requests "intro to programming"
"data structures" requests "discrete math"
"databases" requests "data structures"
"algorithms" requests "data structures"
"formal languages" requests "discrete math"
"compilers" requests "data structures"
"compilers" requests "formal languages"
"compilers" requests "computer organization"
"operating systems" requests "data structures"
"operating systems" requests "computer organization"
"calculus" requests "linear algebra"
"networks" requests "operating systems"
"programming languages" requests "data structures"
"programming languages" requests "computer organization"
PASS
ok  	github.com/budougumi0617/GoTraining/ch05/ex10	0.018s
````
