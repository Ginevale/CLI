# CLI
CLI with urfave/cli

## urfave/cli

What does it support?
Custom commands with:
* Arguments
* Flags 
* Subcommands 
```bash 
go run to-do-list/to-do-list.go
> NAME:
   to-do-list - A new cli application

USAGE:
   to-do-list [global options] command [command options] [arguments...]

COMMANDS:
   cocilova, coci     coci extention
   complete, c        complete a task on the list
   costarella, costi  costi extension
   help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help

go run to-do-list/to-do-list.go cocilova test
```