package service

import (
	"context"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/full"
	"github.com/xh-polaris/openapi-charge/biz/infrastructure/mapper/log"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/charge"
	"time"
)

type ILogService interface {
	CreateLog(ctx context.Context, req *charge.CreateLogReq) (*charge.CreateLogResp, error)
	GetLog(ctx context.Context, req *charge.GetLogReq) (*charge.GetLogResp, error)
}

type LogService struct {
	LogMongoMapper  *log.MongoMapper
	FullMongoMapper *full.MongoMapper
}

var LogServiceSet = wire.NewSet(
	wire.Struct(new(LogService), "*"),
	wire.Bind(new(ILogService), new(*LogService)),
)

func (s *LogService) CreateLog(ctx context.Context, req *charge.CreateLogReq) (res *charge.CreateLogResp, err error) {
	inf, err := s.FullMongoMapper.FindOne(ctx, req.FullInterfaceId)
	if err != nil {
		return &charge.CreateLogResp{
			Done: false,
			Msg:  "完整接口不存在或被删除",
		}, err
	}
	l := &log.Log{
		FullInterfaceId: req.FullInterfaceId,
		UserId:          req.UserId,
		KeyId:           req.KeyId,
		Status:          int64(req.Status),
		Info:            req.Info,
		Count:           req.Count,
		Value:           inf.Price * req.Count,
		Timestamp:       time.Unix(req.Timestamp, 0),
		CreateTime:      time.Now(),
	}
	err = s.LogMongoMapper.Insert(ctx, l)
	if err != nil {
		return &charge.CreateLogResp{
			Done: false,
			Msg:  "创建调用记录失败",
		}, err
	}
	return &charge.CreateLogResp{
		Done: true,
		Msg:  "创建调用记录成功",
	}, nil
}

func (s *LogService) GetLog(ctx context.Context, req *charge.GetLogReq) (res *charge.GetLogResp, err error) {
	data, total, err := s.LogMongoMapper.FindAndCountByInfId(ctx, req.FullInterfaceId, req.PaginationOptions)
	if err != nil {
		return nil, err
	}
	logs := make([]*charge.Log, 0)
	for _, val := range data {
		l := &charge.Log{}
		err = copier.Copy(l, val)
		if err != nil {
			return nil, err
		}
		l.Id = val.ID.Hex()
		l.Status = charge.LogStatus(val.Status)
		l.Timestamp = val.Timestamp.Unix()
		l.CreateTime = val.CreateTime.Unix()
		logs = append(logs, l)
	}
	return &charge.GetLogResp{
		Logs:  logs,
		Total: total,
	}, nil
}
