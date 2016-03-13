# go_map_reduce

## Install go
https://golang.org/doc/install

Note:
go will look up the source code under $GOPATH/src/..

Add "export GOPATH=$HOME/Workspace" into startup script such as "~/.profile"

GOPATH should point to your Workspace since it will used by go to *locate* the source code

```
$GOPATH
	|_ hello_out_src.go
	|_ src
		|_ hello_in_src.go
		|_ hello
			|_ hello_in_pkg.go (package main)
			|_ mod_in_sth.go (package sth)
			|_ mod2_in_sth.go (package sth)
```

The source code should be put under *src* directory.

For compiled code, the *main* package will be under *bin* directory, other package will be under *pkg* directory

The function will be accessed from the package level instead of the file level.


```
go install hello # compile the code under dir $GOPATH/src/hello/

$GOPATH/bin/hello # execute complied code (main func in main package)

hello workd in src/hello
```
