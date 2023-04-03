# Shells generator

This is a simple cli tool to store commands in a sqlite database which is written in GO. 

## A simple cli tool to store and retrieve data such as example commands for a user 



## How to build   

### Install in linux sb-shells

```console
go install github.com/secopsbear/sb-shells@latest
```

### Build for linux

```console
go build -o sb-shells
```

 > Add **`-ldflags "-s -w"`** flags to reduce the file size by deleting the debug links and info in the binary 

### Reduce the file size

```console
upx --ultra-brute sb-shells
```

### Build for window

Generate an executable **`sb-shells.exe`** for windows environment.   

```console
env GOOS=windows GOARCH=amd64 go build -o sb-shells.exe -ldflags "-s -w"
```

## Basic commands

```console
$ sb-shells -h                                                         
Generate shell codes for various environments Powershell, bash, php..    
                                                                         
Usage:                                                                   
  sb-shells [command]                                                    
                                                                         
Available Commands:                                                      
  bash        Generate bash reverse shell scripts                        
  completion  Generate the autocompletion script for the specified shell 
  groovy      Generate Groovy reverse shell script                       
  help        Help about any command                                     
  php         Generate php reverse shell                                 
  ps          Generate powershell reverse shell scripts                  
                                                                         
Flags:                                                                   
  -h, --help                 help for sb-shells                          
  -i, --ipAddrLocal string   Local IP address                            
  -p, --portLocal string     Local Port                                  
                                                                         
Use "sb-shells [command] --help" for more information about a command.   
```





## Find a bug?

If you found an issue or would like to submit an improvement to this project, please submit an issue using the issues tab above.