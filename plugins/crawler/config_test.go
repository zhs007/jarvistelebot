package plugincrawler

import "testing"

func TestLoadConfig(t *testing.T) {
	cfg := loadConfig("../../test/crawler.yaml")

	err := checkConfig(cfg)
	if err != nil {
		t.Fatalf("TestLoadConfig checkConfig %v", err)

		return
	}

	if cfg.CrawlerNodeAddr != "12LyThj17Dj88EsHgVonn1eJffMSwjsXf4" {
		t.Fatalf("TestLoadConfig invalid crawlerNodeAddr %v", cfg.CrawlerNodeAddr)

		return
	}

	if cfg.CrawlerPath != "/mnt/jarviscrawlercore" {
		t.Fatalf("TestLoadConfig invalid crawlerPath %v", cfg.CrawlerPath)

		return
	}

	t.Log("TestLoadConfig OK")
}

func TestCheckConfig(t *testing.T) {
	type data struct {
		cfg *config
		err error
	}

	lst := []data{
		data{
			cfg: nil,
			err: ErrNoConfig,
		},
		data{
			cfg: &config{},
			err: ErrConfigNoCrawlerServAddr,
		},
		data{
			cfg: &config{
				CrawlerServAddr: "127.0.0.1:7081",
				CrawlerNodeAddr: "12LyThj17Dj88EsHgVonn1eJffMSwjsXf4",
			},
			err: ErrConfigNoCrawlerPath,
		},
		data{
			cfg: &config{
				CrawlerServAddr: "127.0.0.1:7081",
				CrawlerNodeAddr: "12LyThj17Dj88EsHgVonn1eJffMSwjsXf45",
			},
			err: ErrConfigInvalidCrawlerNodeAddr,
		},
		data{
			cfg: &config{
				CrawlerServAddr: "127.0.0.1:7081",
				CrawlerNodeAddr: "12LyThj17Dj88EsHgVonn1eJffMSwjsXf",
			},
			err: ErrConfigInvalidCrawlerNodeAddr,
		},
		data{
			cfg: &config{
				CrawlerServAddr: "127.0.0.1:7081",
				CrawlerNodeAddr: "12LyThj17Dj88EsHgVonn1eJffMSwjsXf4",
				CrawlerPath:     "/jarviscrawlercore",
			},
			err: ErrConfigNoUpdateScript,
		},
		data{
			cfg: &config{
				CrawlerServAddr: "127.0.0.1:7081",
				CrawlerNodeAddr: "12LyThj17Dj88EsHgVonn1eJffMSwjsXf4",
				CrawlerPath:     "/jarviscrawlercore",
				UpdateScript:    "update",
			},
			err: ErrConfigNoExpArticleScript,
		},
		data{
			cfg: &config{
				CrawlerServAddr:  "127.0.0.1:7081",
				CrawlerNodeAddr:  "12LyThj17Dj88EsHgVonn1eJffMSwjsXf4",
				CrawlerPath:      "/jarviscrawlercore",
				UpdateScript:     "update",
				ExpArticleScript: "exparticle",
			},
			err: nil,
		},
	}

	for i := 0; i < len(lst); i++ {
		curerr := checkConfig(lst[i].cfg)
		if curerr != lst[i].err {
			t.Fatalf("TestCheckConfig checkConfig %v - %v", lst[i], curerr)
		}
	}

	t.Log("TestCheckConfig OK")
}
