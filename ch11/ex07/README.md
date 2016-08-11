# Exercise 11.7
Write benchmarks for `Add`, `UnionWith`, and other methods for of `*IntSet`(Sec6.5) using large pseudo-random inputs. How fast can you make these methods run? How does the choice of words of word size affect performance? How fast is `IntSet` compared to a set implementation based on the build.

---
# 練習問題 11.7
大きな擬似乱数の入力を使って、`*IntSet`(6.5節)の`Add`、`UnionWith`、および他のメソッドに対するベンチマークを書きなさい。みなさんは、これらのメソッドをどのくらい速くできますか。ワードの大きさの選択は性能にどのように影響するでしょうか。組み込みのマップ型に基づくセットの実装と比較して、`IntSet`はどの程度速いでしょうか。

# Result

````bash
budougumi0617@~/go/src/github.com/budougumi0617/GoTraining/ch11/ex07 (ch10ch11@GoTraining)
$  go test -bench .
BenchmarkMapIntSetHas10-4                     	20000000	        87.0 ns/op
BenchmarkMyIntSetHas10-4                      	20000000	        63.6 ns/op
BenchmarkMapIntSetHas100-4                    	20000000	       103 ns/op
BenchmarkMyIntSetHas100-4                     	20000000	        66.5 ns/op
BenchmarkMapIntSetHas1000-4                   	20000000	       120 ns/op
BenchmarkMyIntSetHas1000-4                    	20000000	        76.8 ns/op
BenchmarkMapIntSetAdd10-4                     	  500000	      2385 ns/op
BenchmarkMyIntSetAdd10-4                      	 2000000	       705 ns/op
BenchmarkMapIntSetAdd100-4                    	  100000	     24698 ns/op
BenchmarkMyIntSetAdd100-4                     	  200000	      7986 ns/op
BenchmarkMapIntSetAdd1000-4                   	    5000	    378305 ns/op
BenchmarkMyIntSetAdd1000-4                    	   10000	    106171 ns/op
BenchmarkMapIntSetAddAll10-4                  	  500000	      2059 ns/op
BenchmarkMyIntSetAddAll10-4                   	10000000	       143 ns/op
BenchmarkMapIntSetAddAll100-4                 	  100000	     17390 ns/op
BenchmarkMyIntSetAddAll100-4                  	 2000000	       928 ns/op
BenchmarkMapIntSetAddAll1000-4                	    5000	    235860 ns/op
BenchmarkMyIntSetAddAll1000-4                 	  200000	      8897 ns/op
BenchmarkMapIntSetUnionWith10-4               	 1000000	      1290 ns/op
BenchmarkMyIntSetUnionWith10-4                	 3000000	       552 ns/op
BenchmarkMapIntSetUnionWith100-4              	  100000	     20204 ns/op
BenchmarkMyIntSetUnionWith100-4               	 2000000	       650 ns/op
BenchmarkMapIntSetUnionWith1000-4             	    5000	    247778 ns/op
BenchmarkMyIntSetUnionWith1000-4              	 2000000	       634 ns/op
BenchmarkMapIntSetIntersectWith10-4           	 5000000	       235 ns/op
BenchmarkMyIntSetIntersectWith10-4            	 3000000	       394 ns/op
BenchmarkMapIntSetIntersectWith100-4          	 2000000	       615 ns/op
BenchmarkMyIntSetIntersectWith100-4           	 3000000	      1032 ns/op
BenchmarkMapIntSetIntersectWith1000-4         	  100000	     14294 ns/op
BenchmarkMyIntSetIntersectWith1000-4          	  200000	      9251 ns/op
BenchmarkMapIntSetDifferenceWith10-4          	  500000	      2668 ns/op
BenchmarkMyIntSetDifferenceWith10-4           	 1000000	      2314 ns/op
BenchmarkMapIntSetDifferenceWith100-4         	   50000	     31547 ns/op
BenchmarkMyIntSetDifferenceWith100-4          	  100000	     16991 ns/op
BenchmarkMapIntSetDifferenceWith1000-4        	    5000	    388077 ns/op
BenchmarkMyIntSetDifferenceWith1000-4         	   20000	     73368 ns/op
BenchmarkMapIntSetSymmetricDifference10-4     	 1000000	      1463 ns/op
BenchmarkMyIntSetSymmetricDifference10-4      	 1000000	      2268 ns/op
BenchmarkMapIntSetSymmetricDifference100-4    	   50000	     23352 ns/op
BenchmarkMyIntSetSymmetricDifference100-4     	  100000	     17800 ns/op
BenchmarkMapIntSetSymmetricDifference1000-4   	    5000	    280146 ns/op
BenchmarkMyIntSetSymmetricDifference1000-4    	   20000	     72553 ns/op
PASS
ok  	github.com/budougumi0617/GoTraining/ch11/ex07	79.377s
````
