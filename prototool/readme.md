# Prereq
- [asdf](https://asdf-vm.com/)
- [asdf-golang](https://github.com/asdf-community/asdf-golang)

# Steps to test
1. install protobuf: 
    `brew install protobuf`
2. install go runtime for protobuf: 
    `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
3. reshim asdf to ensure that `protoc-gen-go` is in path
    `asdf reshim golang`
4. Generate protobuf
    `protoc --go_out=paths=source_relative:. prototool/payload.proto`
5. Run code to generate `vulnerabilities.pb`