# accounts-and-transactions
This is a solution for an interview question.

### What is this repository for? ###
accounts-and-transactions keeps track of account details, and financial transactions linked to the accounts.

### How do I set it up? ###
Clone the repo, and run this command to sync dependencies.
```
go mod tidy -v
```

### How to run tests? ###
Unit and Functional tests can be run independently in the repo.

To run unit tests, use the command:
```
make test-unit
```

To run functinal tests, use the command:
```
make test-functional
```

### Local Docker Setup ###
```
make build-and-run
```

### Features ###
* Create new Account
* Get Account Details
* Register Transactions for the account

#### Technologies Used ###
* Golang
* MySql
* GORM
* gin

### Project Structure ###

```
|-internal
    |-entities - domain entities (account and transaction)
    |-account - account service and use cases
    |-transaction - transaction service and use cases
    |-repository - data access layer
    |-server/http - http handlers, routes & middlewares
    |-validator_service - validation pipeline for transactions
|- cmd
    |- main - entrypoint
|- tests - functional tests        
```

Each pacakge has similar structure:
* types: contains all interfaces and other types used/exposed by the package
* mocks: contains mocks for types defined in the package

### Assumptions taken for the project ###
* Account balance check is required for each debit transaction. A minimum balance has to be maintained.
