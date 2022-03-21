package client

import (
	"context"
	"github.com/rowdyroad/grpc-hw/internal/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)
import proto "github.com/rowdyroad/grpc-hw/internal/grpc"
type Client struct {
	client *grpc.ClientConn
	cl proto.StorageClient
}

func NewClient(addr string) (*Client, error) {

	cc, err := grpc.Dial(addr,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := proto.NewStorageClient(cc)
	log.Println(client)
	return &Client{cc, client}, nil
}

func (c *Client) GetTotalCount(from,to time.Time, low, high float64) (int,error) {
	req, err := c.cl.GetTotalCount(context.TODO(), &proto.TotalCountRequest{
		From: timestamppb.New(from),
		To: timestamppb.New(to),
		Low: low,
		High: high,
	})
	if err != nil {
		return 0, err
	}
	return int(req.Count), nil
}

func (c *Client) GetList(from ,to time.Time, low, high float64, offset,limit int) ([]storage.Record,error) {
	req, err := c.cl.GetList(context.TODO(), &proto.ListRequest{
		From: timestamppb.New(from),
		To: timestamppb.New(to),
		Low: low,
		High: high,
		Offset: uint64(offset),
		Limit: uint64(limit),
	})
	if err != nil {
		return nil, err
	}
	ret := make([]storage.Record, 0, len(req.Records))
	for _,r := range req.Records {
		ret = append(ret, storage.Record{
			Time: r.Time.AsTime(),
			Value: r.Value,
		})
	}

	return ret, nil
}

func (c *Client) GetValue(time time.Time) (float64,error) {
	req, err := c.cl.GetValue(context.TODO(), &proto.ValueRequest{
		Time: timestamppb.New(time),
	})
	if err != nil {
		return 0, err
	}
	return req.Value, nil
}

func (c *Client) GetDailyStats(from, to time.Time) (map[time.Time]storage.Stat, error) {
	req, err := c.cl.GetDailyStats(context.TODO(), &proto.DailyStatsRequest{
		From: timestamppb.New(from),
		To: timestamppb.New(to),
	})
	if err != nil {
		return nil, err
	}
	ret := make(map[time.Time]storage.Stat, len(req.Stats))
	for t, r := range req.Stats {
		ret[time.Unix(int64(t),0)] = storage.Stat{
			Count: r.Count,
			Average: r.Average,
			Min: r.Min,
			Max: r.Max,
		}
	}
	return ret, nil
}