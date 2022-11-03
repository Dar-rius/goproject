
Do you sometimes forget where you stored your project ?


You search in several folders where you think it might be in and in the end you find it but you have wasted a lot of time and energy to find it, with
GoProject you can store and access your projects in a command line

## How to use 
This CLI has 3 commands (add, go and ls)


The add command allows you to add a project in the application by specifying the name of the project and its path, this will be saved in the application

```bash
$ goproject add name_project path_project
```

The ls command allows you to list the projects to store

```bash 
$ goproject ls 
```

The go command will allow you to move directly into your project by specifying the project you want to access

```bash
$ goproject go name_project
```


## How install

To install the CLI copy the following command:

```bash
$ go install github.com/Dar-rius/GoProject@latest
```


The stable version is not available, if you are under linux or mac os you will probably have errors on the use of the CLI
