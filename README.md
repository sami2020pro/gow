# Gow - Work and todo manager CLI tool

Gow is a cli tool helps you to manage add and done jobs and todos easier right when you are doing it

`Still in development`

# How to run it (Development)
```
$ go run main.go --help
```

# How to run tests
```
$ go test test/gow_test.go
```

# Commands

### Works

Use works command to get all of your works and todos or add new work or todo.

#### Sub-commands
[add] [string name] Add new work or todo

[check] [string name or id] Mark work as done

[remove] [string name or id] Remove work

#### Flags

[-d] Use this to add description to work or todo
[-f] Use this to get just doned works
[-r] Use this to get remaining works

### Init

This command generate a TODO.md file for your given path if not exists that you can add todos and check them without writing directly in TODO.md file

#### Sub-commands

[add] [int index] Add new todo

[check] [int index] check as done

[remove] [int index] Remove work

#### Flags

[-p] Path of your project root default is .(Means current folder)
[-d] For adding description only available in init command

> This is temporary readme untill project receive stable releases
