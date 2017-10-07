package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"strconv"

	"github.com/budougumi0617/gopl/ch07/ex15/eval"
)

func main() {
	Query()
}

func Query() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input expression: ")
	expr, _ := reader.ReadString('\n')
	exp, err := eval.Parse(expr)

	if err != nil {
		log.Fatalln(err)
	}

	var vars = make(map[eval.Var]bool)
	if err := exp.Check(vars); err != nil {
		log.Fatalln(err)
	}
	env := eval.Env{}
	for key := range vars {
		fmt.Printf("key %s:", key)
		value, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		value = value[:len(value)-1]
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatalln(err)

		}
		env[key] = f
	}

	fmt.Printf("Result %+v\n", exp.Eval(env))
}
