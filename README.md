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
   cocilova, coci     coci extension
   costarella, costi  costi extension
   help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

You can try the __coci extension__ with the following commands:
```bash 
go run to-do-list/to-do-list.go cocilova test
go run to-do-list/to-do-list.go cocilova curiosity
```

You can try the __costi extension__ with the following commands:
```bash 
go run to-do-list/to-do-list.go costarella test
go run to-do-list/to-do-list.go costarella curiosity
```