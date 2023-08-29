package main

import (
	"fmt"

	"github.com/pspiagicw/goreland"
)

func main() {
    fmt.Println()
    goreland.LogSuccess("The operation was successful")
    fmt.Println()
    goreland.LogError("Something went wrong")
    fmt.Println()
    goreland.LogInfo("Information to be noted")
    fmt.Println()
    goreland.LogBullet("Bullets!")
    fmt.Println()
    goreland.LogExec("Executed something")
    fmt.Println()
    goreland.LogExecSimple("Less fancy execution")
    fmt.Println()
    fmt.Println()
}
