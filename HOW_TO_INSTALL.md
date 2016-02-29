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
			|_hello_in_pkg.go
```

The source code should be put under *src* directory, the compiled code will be under *bin* directory

```
go install hello

$GOPATH/bin/hello

hello workd in src/hello
```

go compiled the file under $GOPATH/src/hello/