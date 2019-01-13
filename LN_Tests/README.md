# Network and Multimedia Final Project - Lighting Network Lite

This program runs in a Go envieronment. Click [here](https://golang.org/doc/code.html) for more information about environment set up.

This project focuses on implementing and testing the features of Lightning Networks. Files related to Lighting Network implementation and testings (listed below) are original works. The Blockchain prototype is modified from [here](https://github.com/Jeiwan/blockchain_go/tree/part_6).

Original works:
main.go, cli_newfund.go, cli_litsend.go, cli_execcommit.go, cli_endfund.go, lit_wallets.go, lit_wallet.go, testOne.go, testTwo.go, testThree.go, testFour.go, sign_and_verify.go, transaction_commitment.go, transaction_funding.go, transaction_settlement.go, timelock.go

## Tests
Test 1 - Blockchain: execute 30 transaction on blockchain and print out execution time

Test 2 - Lightning Network: normal scenario, execute 30 lightning transaction and print out execution time

Test 3 - Lightning Network: Broadcast a commitment prematurely and spend the output before time lock expires (The last transaction should be invalid)

Test 4 - Lightning Network: Broadcast a commitment prematurely and spend the output after time lock expires (Should success)


## Usage

Specify a test when you run the program.

Example:
```bash
./LN_Tests test1
```

REMINDER: Remember to delete all previous data files (.db, .dat) before starting a new test
```bash
rm *.dat *.db
```


