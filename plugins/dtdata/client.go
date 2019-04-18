package plugindtdata

import (
	"context"

	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"

	"github.com/zhs007/jarvistelebot/jarviscrawlercore"
	"google.golang.org/grpc"
)

type dtdataClient struct {
	cfg    *config
	conn   *grpc.ClientConn
	client jarviscrawlercore.JarvisCrawlerServiceClient
}

// newDTDataClient - new dtdataClient
func newDTDataClient(cfg *config) *dtdataClient {
	return &dtdataClient{
		cfg: cfg,
	}
}

// getDTData -
func (tc *dtdataClient) getDTData(ctx context.Context, mode string, startTime string, endTime string) (*jarviscrawlercore.ReplyDTData, error) {
	if tc.cfg == nil {
		return nil, ErrNoConfig
	}

	if tc.conn == nil || tc.client == nil {
		conn, err := grpc.Dial(tc.cfg.DTDataServAddr, grpc.WithInsecure())
		if err != nil {
			jarvisbase.Warn("dtdataClient.getDTData:grpc.Dial", zap.Error(err))

			return nil, err
		}

		tc.conn = conn

		tc.client = jarviscrawlercore.NewJarvisCrawlerServiceClient(conn)
	}

	reply, err := tc.client.GetDTData(ctx, &jarviscrawlercore.RequestDTData{
		Mode:      mode,
		StartTime: startTime,
		EndTime:   endTime,
	})
	if err != nil {
		jarvisbase.Warn("dtdataClient.getDTData:GetDTData", zap.Error(err))

		// if error, close connect
		tc.conn.Close()

		tc.conn = nil
		tc.client = nil

		return nil, err
	}

	return reply, nil
}
