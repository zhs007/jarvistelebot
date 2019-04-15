package plugincrawler

import (
	"context"

	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"

	"github.com/zhs007/jarvistelebot/jarviscrawlercore"
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
func (cc *crawlerClient) getArticles(ctx context.Context, website string) (*jarviscrawlercore.ArticleList, error) {
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
		Website: website,
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

// translate -
func (cc *crawlerClient) translate(ctx context.Context, text string, srclang string, destlang string) (string, error) {
	if cc.cfg == nil {
		return "", ErrNoConfig
	}

	if cc.conn == nil || cc.client == nil {
		conn, err := grpc.Dial(cc.cfg.CrawlerServAddr, grpc.WithInsecure())
		if err != nil {
			jarvisbase.Warn("crawlerClient.translate:grpc.Dial", zap.Error(err))

			return "", err
		}

		cc.conn = conn

		cc.client = jarviscrawlercore.NewJarvisCrawlerServiceClient(conn)
	}

	reply, err := cc.client.Translate(ctx, &jarviscrawlercore.RequestTranslate{
		Text:     text,
		Platform: "google",
		SrcLang:  srclang,
		DestLang: destlang,
	})
	if err != nil {
		jarvisbase.Warn("crawlerClient.translate:Translate", zap.Error(err))

		// if error, close connect
		cc.conn.Close()

		cc.conn = nil
		cc.client = nil

		return "", err
	}

	return reply.Text, nil
}
