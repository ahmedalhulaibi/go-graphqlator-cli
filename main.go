package main

import (
	"fmt"
	"os"

	"github.com/ahmedalhulaibi/graphqlator/cmd"

	_ "github.com/ahmedalhulaibi/substance/providers/mysqlsubstance"
	_ "github.com/ahmedalhulaibi/substance/providers/pgsqlsubstance"
	_ "github.com/ahmedalhulaibi/substance/providers/testsubstance"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
