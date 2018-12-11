package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/golang/protobuf/proto"
	desc "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/genproto/googleapis/api/annotations"
)

func main() {
	f := &desc.FileDescriptorProto{
		Syntax:  proto.String("proto3"),
		Name:    proto.String("break_java"),
		Package: proto.String("example"),
	}

	// Create a dummy message that will be used by the method
	// as input and output type.
	f.MessageType = []*desc.DescriptorProto{
		{
			Name: proto.String("example_message"),
		},
	}

	// Build up the google api annotations.
	rule := &annotations.HttpRule{
		Pattern: &annotations.HttpRule_Get{Get: "/v1/example"},
	}
	o := &desc.MethodOptions{}
	// Attach the annotations to the MessageOptions
	if err := proto.SetExtension(o, annotations.E_Http, rule); err != nil {
		log.Fatal(err)
	}
	// Create a dummy method to attach to the service. This will contain
	// the google api annotations.
	method := &desc.MethodDescriptorProto{
		Name:       proto.String("example_method"),
		InputType:  proto.String(".example.example_message"),
		OutputType: proto.String(".example.example_message"),
		Options:    o,
	}

	// Create a dummy service to attach to the file.
	f.Service = []*desc.ServiceDescriptorProto{
		{
			Name:   proto.String("example_service"),
			Method: []*desc.MethodDescriptorProto{method},
		},
	}

	req := &desc.FileDescriptorSet{
		File: []*desc.FileDescriptorProto{f},
	}

	// Languages to try creating this for. All but 'java' should
	// work.
	languages := []string{
		"python",
		"csharp",
		"java",
		"php",
	}
	// Marshal the FileDescriptorSet to the on-the-wire format.
	bs, err := proto.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}
	// Loop through the languages.
	for _, lang := range languages {
		// Call out the protoc
		cmd := exec.Command("protoc", "--"+lang+"_out=./generated", "--descriptor_set_in=/dev/stdin", "break_java")
		// Feed the marshalled FileDescriptorSet to the commands stdin,
		// which 'protoc' is looking for data from.
		cmd.Stdin = bytes.NewReader(bs)
		out, err := cmd.CombinedOutput()
		fmt.Printf("%s -> stdout=%q err=%v\n", cmd.Args, out, err)
	}
}
