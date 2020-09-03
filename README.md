# Hermes: Message Distributor

Hermes is a message distribution system in Go. It receives messages of a specific user and subject, and distributes the message to the users folowers or groups.

# Install depencies

``` sh
go mod vendor
```

# Run main.go file

```sh
go run main.go feeds(or whatever name you want)
```

# Run compiled version

## Setup

Firstly, the progam should be compiled with the following command:

```` sh
go build ./main.go
````
After that you should take the `.env.example` and rename it `.env` to setup your environmental variables. Inside that file you should be able to set your local MySql database credentials.

## Running the migration file

To create all the necessary tables for Hermes to work, the migrations must be run with the following command:

```` sh
go run migration.go
````

## Usage

To use simply execute the program passing it the name of the queue you want to work with.

```` sh
./feeds example-queue
````
