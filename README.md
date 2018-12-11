```
$ go run main.go
[protoc --python_out=./generated --descriptor_set_in=/dev/stdin break_java] -> stdout="" err=<nil>
[protoc --csharp_out=./generated --descriptor_set_in=/dev/stdin break_java] -> stdout="" err=<nil>
[protoc --java_out=./generated --descriptor_set_in=/dev/stdin break_java] -> stdout="[libprotobuf FATAL ../../../../../src/google/protobuf/compiler/java/java_file.cc:129] CHECK failed: file_proto_desc: Find unknown fields in FileDescriptorProto when building break_java. It's likely that those fields are custom options, however, descriptor.proto is not in the transitive dependencies. This normally should not happen. Please report a bug.\nterminate called after throwing an instance of 'google::protobuf::FatalException'\n  what():  CHECK failed: file_proto_desc: Find unknown fields in FileDescriptorProto when building break_java. It's likely that those fields are custom options, however, descriptor.proto is not in the transitive dependencies. This normally should not happen. Please report a bug.\n" err=signal: aborted (core dumped)
[protoc --php_out=./generated --descriptor_set_in=/dev/stdin break_java] -> stdout="" err=<nil>
```
