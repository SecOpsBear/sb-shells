# Shells generator

A simple cli tool to generate reverse shell in various languages and environments written in GO.


## Demo

![sb-shells demo](https://secopsbear.com/assets/images/sb-shells-025a7c81680820f8f8d5d3cea89db1da.gif)

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

More information on demo and usage is available at [https:/secopsbear.com](https://secopsbear.com/tools/bear_tools/sb-shells)



## Find a bug?

If you found an issue or would like to submit an improvement to this project, please submit an issue using the issues tab above.