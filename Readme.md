# usage
In order to use this package, the protobuf stubs must be genetated. To generate file run
```sh
make create_proto_definitions
```
#### issues
- There is a known issue with the context package being used in the generated ```.pb.go```. To avoid errors, replace:
```context "context"``` in definitions.pb.go file with ```context "golang.org/x/net/context"```.
# dependencies
Both protoc and protoc-gen-go are needed for this project
##### protoc
for Mac
```sh
brew install protobuf
```
for arch linux 
```sh
sudo pacman -S protobuf
```
or visit https://github.com/protocolbuffers/protobuf
##### protoc-gen-go
see	https://developers.google.com/protocol-buffers/
```sh
go get -u github.com/golang/protobuf/protoc-gen-go
```
# testing
To test, have the main file running then start the test file.