package main

//go:generate bash -c "mkdir -p $GOPATH/src/github.com/vasili-v/grpc-unary-test/gut-server/unary && protoc -I $GOPATH/src/github.com/vasili-v/grpc-unary-test/gut-server/ $GOPATH/src/github.com/vasili-v/grpc-unary-test/gut-server/unary.proto --go_out=plugins=grpc:$GOPATH/src/github.com/vasili-v/grpc-unary-test/gut-server/unary && ls $GOPATH/src/github.com/vasili-v/grpc-unary-test/gut-server/unary"

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/vasili-v/grpc-unary-test/gut-server/unary"
)

type pair struct {
	req *pb.Request

	sent time.Time
	recv *time.Time
}

func main() {
	pairs := newPairs(total, size)

	c, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Errorf("dialing error: %s", err))
	}
	defer c.Close()

	s := pb.NewStreamClient(c)
	if limit > 1 {
		var wg sync.WaitGroup
		th := make(chan int, limit)
		for _, p := range pairs {
			p.sent = time.Now()
			th <- 0
			wg.Add(1)
			go func(p *pair) {
				defer func() {
					<-th
					wg.Done()
				}()

				if _, err := s.Test(context.Background(), p.req); err == nil {
					t := time.Now()
					p.recv = &t
				}
			}(p)
		}

		wg.Wait()
	} else {
		for _, p := range pairs {
			p.sent = time.Now()
			if _, err := s.Test(context.Background(), p.req); err == nil {
				t := time.Now()
				p.recv = &t
			}
		}
	}

	dump(pairs, "")
}

func newPairs(n, size int) []*pair {
	out := make([]*pair, n)

	if size > 0 {
		fmt.Fprintf(os.Stderr, "making messages to send:\n")
	}

	for i := range out {
		if size > 0 {
			buf := make([]byte, size)
			for j := range buf {
				buf[j] = byte(rand.Intn(256))
			}

			if i < 3 {
				fmt.Fprintf(os.Stderr, "\t%d: % x\n", i, buf)
			} else if i == 3 {
				fmt.Fprintf(os.Stderr, "\t%d: ...\n", i)
			}

			out[i] = &pair{
				req: &pb.Request{
					Id:      uint32(i),
					Payload: buf,
				},
			}
		} else {
			out[i] = &pair{}
		}
	}

	return out
}
