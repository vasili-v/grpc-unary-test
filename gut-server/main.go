package main

//go:generate bash -c "mkdir -p $GOPATH/src/github.com/vasili-v/grpc-unary-test/gut-server/unary && protoc -I $GOPATH/src/github.com/vasili-v/grpc-unary-test/gut-server/ $GOPATH/src/github.com/vasili-v/grpc-unary-test/gut-server/unary.proto --go_out=plugins=grpc:$GOPATH/src/github.com/vasili-v/grpc-unary-test/gut-server/unary && ls $GOPATH/src/github.com/vasili-v/grpc-unary-test/gut-server/unary"

import (
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/vasili-v/grpc-unary-test/gut-server/unary"
)

type server struct{}

func (s *server) Test(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	out := pb.Response{
		Id:      in.Id,
		Payload: make([]byte, len(in.Payload)),
	}

	for i := range out.Payload {
		out.Payload[i] ^= 0x55
	}

	return &out, nil
}

func main() {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	p := grpc.NewServer()
	pb.RegisterStreamServer(p, &server{})
	err = p.Serve(ln)
	if err != nil {
		panic(err)
	}
}
