package plugintranslate

import "testing"

func TestLoadConfig(t *testing.T) {
	cfg := loadConfig("../../test/translate.yaml")

	err := checkConfig(cfg)
	if err != nil {
		t.Fatalf("TestLoadConfig checkConfig %v", err)

		return
	}

	if cfg.TranslateServAddr != "192.168.0.1:7051" {
		t.Fatalf("TestLoadConfig invalid TranslateServAddr %v", cfg.TranslateServAddr)

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
			err: ErrConfigNoTranslateServAddr,
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
