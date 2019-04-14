package plugincrawler

import (
	"context"

	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"

	"github.com/zhs007/jarvistelebot/plugins/crawler/servproto"
	"google.golang.org/grpc"
)

type crawlerClient struct {
	cfg    *config
	conn   *grpc.ClientConn
	client jarviscrawlercore.JarvisCrawlerServiceClient
}

// newCrawlerClient - new crawlerClient
func newCrawlerClient(cfg *config) *crawlerClient {
	return &crawlerClient{
		cfg: cfg,
	}
}

// getArticles -
func (cc *crawlerClient) getArticles(ctx context.Context, url string, attachjquery bool) (*jarviscrawlercore.ArticleList, error) {
	if cc.cfg == nil {
		return nil, ErrNoConfig
	}

	if cc.conn == nil || cc.client == nil {
		conn, err := grpc.Dial(cc.cfg.CrawlerServAddr, grpc.WithInsecure())
		if err != nil {
			jarvisbase.Warn("crawlerClient.getArticles:grpc.Dial", zap.Error(err))

			return nil, err
		}

		cc.conn = conn

		cc.client = jarviscrawlercore.NewJarvisCrawlerServiceClient(conn)
	}

	reply, err := cc.client.GetArticles(ctx, &jarviscrawlercore.RequestArticles{
		Url:          url,
		AttachJQuery: attachjquery,
	})
	if err != nil {
		jarvisbase.Warn("crawlerClient.getArticles:GetArticles", zap.Error(err))

		// if error, close connect
		cc.conn.Close()

		cc.conn = nil
		cc.client = nil

		return nil, err
	}

	return reply.Articles, nil
}
