# Exercise 5.11
The instructor of the linear algebra course decides that calculus is now a prerequisite. Extend the `topoSort` function to report cycles.

---
# 練習問題 5.11
線形代数（linear algebra）の講座の教官は、微積分学（calculus）をこれからは事前条件にすると決めました。循環を報告するように`topoSort`関数を拡張しなさい。


# Result

````
$  go run toposort.go
1:	intro to programming
2:	discrete math
3:	data structures
4:	algorithms
5:	Found Cycles between "linear algebra" and "calculus"
6:	Found Cycles between "calculus" and "linear algebra"
7:	linear algebra
8:	calculus
9:	formal languages
10:	computer organization
11:	compilers
12:	databases
13:	operating systems
14:	networks
15:	programming languages
````
