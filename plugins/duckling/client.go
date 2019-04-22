package pluginduckling

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/zhs007/jarviscore"
)

type ducklingClient struct {
	cfg *config
}

// newDucklingClient - new ducklingClient
func newDucklingClient(cfg *config) *ducklingClient {
	return &ducklingClient{
		cfg: cfg,
	}
}

// request -
func (dc *ducklingClient) request(ctx context.Context, lang string, text string) (string, error) {
	if dc.cfg == nil {
		return "", ErrNoDucklingConfig
	}

	resp, err := http.Post(dc.cfg.DucklingServAddr,
		"application/x-www-form-urlencoded",
		strings.NewReader(jarviscore.AppendString("local=", lang, "&text=", text)))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
