package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        printUsage()
        os.Exit(1)
	}


    switch os.Args[1] {
    case "test1":
        testOne()

	case "test2":
        testTwo()

    case "test3":
        testThree()
        
    case "test4":
        testFour()
        
    default:
		printUsage()
		os.Exit(1)
	}


}

func printUsage() {
    fmt.Println("Please select a test: test1/test2/test3/test4")
    fmt.Println("Test 1 - Blockchain: execute 30 transaction on blockchain and print out execution time")
    fmt.Println("Test 2 - Lightning Network: normal scenario, execute 30 lightning transaction and print out execution time")
    fmt.Println("Test 3 - Lightning Network: Broadcast a commitment prematurely and spend the output before time lock expires")
    fmt.Println("Test 4 - Lightning Network: Broadcast a commitment prematurely and spend the output after time lock expires")
    fmt.Println("REMINDER: Remember to delete all previous data files (.db, .dat) before starting a new test")
}


