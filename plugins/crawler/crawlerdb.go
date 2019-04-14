package plugincrawler

import (
	"context"

	"github.com/zhs007/jarviscore"

	"github.com/zhs007/ankadb"
	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"
)

// crawlerDB -
type crawlerDB struct {
	ankaDB ankadb.AnkaDB
}

// newCrawlerDB - new assistant db
func newCrawlerDB(dbpath string, httpAddr string, engine string) (*crawlerDB, error) {
	cfg := ankadb.NewConfig()

	cfg.AddrHTTP = httpAddr
	cfg.PathDBRoot = dbpath
	cfg.ListDB = append(cfg.ListDB, ankadb.DBConfig{
		Name:   CrawlerDBName,
		Engine: engine,
		PathDB: CrawlerDBName,
	})

	ankaDB, err := ankadb.NewAnkaDB(cfg, nil)
	if ankaDB == nil {
		jarvisbase.Error("newCrawlerDB", zap.Error(err))

		return nil, err
	}

	jarvisbase.Info("newCrawlerDB", zap.String("dbpath", dbpath),
		zap.String("httpAddr", httpAddr), zap.String("engine", engine))

	db := &crawlerDB{
		ankaDB: ankaDB,
	}

	return db, err
}

func (db *crawlerDB) addArticle(ctx context.Context, uid string, website string, url string) (bool, error) {
	has, err := db.getArticle(ctx, uid, website, url)
	if err != nil {
		return false, err
	}

	if has {
		return false, nil
	}

	err = db.ankaDB.Set(ctx, CrawlerDBName, jarviscore.AppendString(uid, ":", website, ":", url), []byte(url))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (db *crawlerDB) getArticle(ctx context.Context, uid string, website string, url string) (bool, error) {
	_, err := db.ankaDB.Get(ctx, CrawlerDBName, jarviscore.AppendString(uid, ":", website, ":", url))
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
