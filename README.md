# Kontrax

## Setup New Machine

### Download and install protobuf compiler, protoc

Download the appropriate release here: https://github.com/google/protobuf/releases

Unzip the folder
Enter the folder and run ./autogen.sh && ./configure && make

If you run into this error: autoreconf: failed to run aclocal: No such file or directory, 
run brew install autoconf && brew install automake. And run the command from step 3 again.

Then run these other commands. They should run without issues
$ make check
$ sudo make install
$ which protoc
$ protoc --version

### Download and install gRPC and protoc plugin for Go 

go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
