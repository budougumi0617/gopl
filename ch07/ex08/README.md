# Exercise 7.8
Many GUIs provide a table widget with a stateful multi-tier sort: the primary sort key is the most recently clicked column head, the secondary sort key is the second-most recently clicked column head, and so on. Define an implementation of `sort.Interface` for use by such a table. Compare that approach with repeated sorting using `sort.Stable`.


---
# 練習問題 7.8
多くのGUIは、状態を持つ多段ソートの表ウィジェットを提供しています。一次ソートキーは最も直近にクリックされた列の見出し、二次ソートキーは二番目に近くクリックされた列の見出しといった具合になります。このような表が使用する`sort.Interface`の実装を定義しなさい。その実装を`sort.Stable`を使用する繰り返しソートと比較しなさい。



# Result

````bash
Sort by Title, Year, and Length
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go          Moby            Moby               1992  3m37s
Go          Delilah         From the Roots Up  2012  3m38s
Go Ahead    Alicia Keys     As I Am            2007  4m36s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s

Sort by Title, Year, and Length using sort.Stable
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go          Moby            Moby               1992  3m37s
Go          Delilah         From the Roots Up  2012  3m38s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
Go Ahead    Alicia Keys     As I Am            2007  4m36s
````
