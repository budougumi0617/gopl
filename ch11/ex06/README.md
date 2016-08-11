# Exercise 11.6
Write benchmarks to compare the `PopCount` implementation in Section2.6.2 with your solutions to Exercise 2.4 and Exercise 2.5. At what point does the table-based approach break even?

---
# 練習問題 11.6
2.6.2節の`PopCount`の実装方法と練習問題2.4と2.5のみなさんの解答を比較するためのベンチマークをかきなさい。表に基づく方法のほうがよくなるのはどの時点からですか。


# Result


About 5,000.  
**BenchmarkPopCountByTable5000-4       	   50000	     35597 ns/op**  
BenchmarkPopCountByClear5000-4       	   50000	     37239 ns/op  
BenchmarkPopCountByShifting5000-4    	    2000	    613033 ns/op  

````bash
budougumi0617@~/go/src/github.com/budougumi0617/GoTraining/ch11/ex06 (ch10ch11@GoTraining)
$  go test -bench .
BenchmarkPopCountByTable10-4         	 3000000	       486 ns/op
BenchmarkPopCountByClear10-4         	50000000	        41.7 ns/op
BenchmarkPopCountByShifting10-4      	 2000000	       713 ns/op
BenchmarkPopCountByTable100-4        	 1000000	      1106 ns/op
BenchmarkPopCountByClear100-4        	 3000000	       540 ns/op
BenchmarkPopCountByShifting100-4     	  200000	      9130 ns/op
BenchmarkPopCountByTable1000-4       	  200000	      7364 ns/op
BenchmarkPopCountByClear1000-4       	  200000	      7225 ns/op
BenchmarkPopCountByShifting1000-4    	   10000	    120477 ns/op
BenchmarkPopCountByTable5000-4       	   50000	     35597 ns/op
BenchmarkPopCountByClear5000-4       	   50000	     37239 ns/op
BenchmarkPopCountByShifting5000-4    	    2000	    613033 ns/op
BenchmarkPopCountByTable10000-4      	   20000	     71680 ns/op
BenchmarkPopCountByClear10000-4      	   20000	     77377 ns/op
BenchmarkPopCountByShifting10000-4   	    1000	   1248871 ns/op
BenchmarkPopCountByTable20000-4      	   10000	    143470 ns/op
BenchmarkPopCountByClear20000-4      	   10000	    167136 ns/op
BenchmarkPopCountByShifting20000-4   	     500	   2499140 ns/op
PASS
ok  	github.com/budougumi0617/GoTraining/ch11/ex06	32.159s
````
