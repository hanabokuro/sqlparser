package main

import (
	"encoding/json"
	"fmt"
	"github.com/hanabokuro/sqlparser"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	body, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err) // @@@
	}

	var trees = make([]sqlparser.Statement, 0, 100)
	for _, aSql := range strings.Split(string(body), ";") {
		aSql = strings.TrimSpace(aSql)
		if aSql == "" {
			continue
		}
		tree, err := sqlparser.ParseStrictDDL(aSql)
		if err != nil {
			panic(fmt.Sprintf("\n%v\n\n\nbad\n=====\n%s\n=====\n", err, aSql))
		}
		trees = append(trees, tree)
	}

	ret, _ := json.MarshalIndent(trees, "", " ")
	fmt.Println(string(ret))
}
