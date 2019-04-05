package plugintranslate

import (
	"context"

	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"

	"github.com/zhs007/jarvistelebot/plugins/translate/servproto"
	"google.golang.org/grpc"
)

type translateClient struct {
	cfg    *config
	conn   *grpc.ClientConn
	client jarviscrawlercore.JarvisCrawlerServiceClient
}

// newTranslateClient - new translateClient
func newTranslateClient(cfg *config) *translateClient {
	return &translateClient{
		cfg: cfg,
	}
}

// translate -
func (tc *translateClient) translate(ctx context.Context, text string, srclang string, destlang string) (string, error) {
	if tc.cfg == nil {
		return "", ErrNoConfig
	}

	if tc.conn == nil || tc.client == nil {
		conn, err := grpc.Dial(tc.cfg.TranslateServAddr, grpc.WithInsecure())
		if err != nil {
			jarvisbase.Warn("translateClient.translate:grpc.Dial", zap.Error(err))

			return "", err
		}

		tc.conn = conn

		tc.client = jarviscrawlercore.NewJarvisCrawlerServiceClient(conn)
	}

	reply, err := tc.client.Translate(ctx, &jarviscrawlercore.RequestTranslate{
		Text:     text,
		Platform: "google",
		SrcLang:  srclang,
		DestLang: destlang,
	})
	if err != nil {
		jarvisbase.Warn("translateClient.translate:Translate", zap.Error(err))

		return "", err
	}

	return reply.Text, nil
}
