# Exercise 12.4
Modify `encode` to pretty-print the S-expression in the style shown above.

---
# 練習問題 12.4
前述のスタイルでS式をプリティプリントするように`encode`を修正しなさい。


# Result

````shel
budougumi0617@~/git/gotraining/ch12/ex04/sexpr (ch12ch13_remainingwork@GoTraining)
$  go test -v
=== RUN   Test
--- PASS: Test (0.00s)
	encode_test.go:55: Marshal() =
		((Title "Dr. Strangelove")
		 (Subtitle "How I Learned to Stop Worrying and Love the Bomb")
		 (Year 1964)
		 (Actor (("Gen. Buck Turgidson" "George C. Scott")
		         ("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
		         ("Maj. T.J. \"King\" Kong" "Slim Pickens")
		         ("Dr. Strangelove" "Peter Sellers")
		         ("Grp. Capt. Lionel Mandrake" "Peter Sellers")
		         ("Pres. Merkin Muffley" "Peter Sellers")))
		 (Oscars ("Best Actor (Nomin.)"
		          "Best Adapted Screenplay (Nomin.)"
		          "Best Director (Nomin.)"
		          "Best Picture (Nomin.)"))
		 (Sequel nil))
	encode_test.go:62: Unmarshal() = {Title:Dr. Strangelove Subtitle:How I Learned to Stop Worrying and Love the Bomb Year:1964 Actor:map[Dr. Strangelove:Peter Sellers Grp. Capt. Lionel Mandrake:Peter Sellers Pres. Merkin Muffley:Peter Sellers Gen. Buck Turgidson:George C. Scott Brig. Gen. Jack D. Ripper:Sterling Hayden Maj. T.J. "King" Kong:Slim Pickens] Oscars:[Best Actor (Nomin.) Best Adapted Screenplay (Nomin.) Best Director (Nomin.) Best Picture (Nomin.)] Sequel:<nil>}
=== RUN   TestFloat
--- PASS: TestFloat (0.00s)
=== RUN   TestComplex
--- PASS: TestComplex (0.00s)
=== RUN   TestInterface
--- PASS: TestInterface (0.00s)
PASS
ok  	github.com/budougumi0617/GoTraining/ch12/ex04/sexpr	0.015s
````
