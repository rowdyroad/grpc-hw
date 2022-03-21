package handler

import (
	proto "github.com/rowdyroad/grpc-hw/internal/grpc"
	"github.com/rowdyroad/grpc-hw/internal/storage"
	"google.golang.org/protobuf/types/known/timestamppb"
)
import "context"

type Handler struct {
	proto.StorageServer
	storage storage.IStorage
}

func NewHandler(storage storage.IStorage) *Handler {
	return &Handler{storage: storage}
}

func (s *Handler) GetTotalCount(ctx context.Context, req *proto.TotalCountRequest) (*proto.TotalCountResponse,error) {
	count, err := s.storage.GetTotalCount(req.From.AsTime(),req.To.AsTime(), req.Low, req.High)
	return &proto.TotalCountResponse{
		Count: uint64(count),
	}, err
}

func (s *Handler) GetList(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	list, err := s.storage.GetList(req.From.AsTime(), req.To.AsTime(), req.Low, req.High, req.Offset, req.Limit)

	ret := make([]*proto.Record, 0, len(list))
	for _, item := range list {
		ret = append(ret, &proto.Record{
			Time: timestamppb.New(item.Time),
			Value: item.Value,
		})
	}
	return &proto.ListResponse{
		Records: ret,
	}, err
}
func (s *Handler) GetValue(ctx context.Context, req *proto.ValueRequest) (*proto.ValueResponse, error) {
	value, err := s.storage.GetValue(req.Time.AsTime())
	return &proto.ValueResponse{
		Value: value,
	}, err
}


func (s *Handler) GetDailyStats(ctx context.Context, req *proto.DailyStatsRequest) (*proto.DailyStatsResponse, error) {
	stats, err := s.storage.GetDailyStats(req.From.AsTime(), req.To.AsTime())
	resp := proto.DailyStatsResponse{
		Stats: map[int64]*proto.Stat{},
	}
	for t, stat := range stats {
		resp.Stats[t.Unix()] = &proto.Stat{
			Count: uint64(stat.Count),
			Average: stat.Average,
			Min: stat.Min,
			Max: stat.Max,
		}
	}
	return &resp, err
}