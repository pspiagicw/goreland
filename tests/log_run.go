package main

import (
	"fmt"

	"github.com/pspiagicw/goreland"
)

func main() {
	goreland.LogSuccess("The operation was successful")
	goreland.LogError("Something went wrong")
	goreland.LogInfo("Information to be noted")
	goreland.LogBullet("Bullets!")
	goreland.LogExec("Executed something")
	goreland.LogExecSimple("Less fancy execution")
	fmt.Println()
	fmt.Println()

	goreland.LogTable([]string{"Packge Name", "Package Location"}, [][]string{
		{"something", "some path"},
		{"something else", "some other path"},
	})
}
