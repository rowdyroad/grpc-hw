package client

import "google.golang.org/grpc"

type Client struct {
	client *grpc.ClientConn
}

func NewClient(addr string) (*Client, error) {
	cc, err := grpc.Dial(addr)
	if err != nil {
		return nil, err
	}

	return &Client{cc}, nil
}

func (c *)