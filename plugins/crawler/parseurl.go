package plugincrawler

import (
	"sync"

	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"
)

// URLResult - URL result
type URLResult struct {
	URLType string `json:"URLType"`
	URL     string `json:"URL"`
	PDF     string `json:"PDF"`
}

// isSame - is same
func isSame(src *URLResult, dest *URLResult) bool {
	if src == nil && dest == nil {
		return true
	}

	if src == nil || dest == nil {
		return false
	}

	return src.PDF == dest.PDF && src.URL == dest.URL && src.URLType == dest.URLType
}

// FuncParseURL - func parseurl
type FuncParseURL func(url string) *URLResult

// URLParser - URL parser
type URLParser struct {
	mapFuncParseURL sync.Map
}

// NewURLParser - new NewURLParser
func NewURLParser() *URLParser {
	p := &URLParser{}

	p.Reg(articleSMZDM, parseArticleSMZDM)
	p.Reg(article36kr, parseArticle36kr)
	p.Reg(articlebaijingapp, parseArticlebaijingapp)
	p.Reg(articlehuxiu, parseArticlehuxiu)
	p.Reg(articlegeekpark, parseArticlegeekpark)
	p.Reg(articlelieyunwang, parseArticlelieyunwang)
	p.Reg(articletmtpost, parseArticletmtpost)

	return p
}

// Reg - parse url
func (p *URLParser) Reg(name string, funcParse FuncParseURL) error {
	oldf := p.get(name)
	if oldf != nil {
		jarvisbase.Warn("URLParser.Reg", zap.Error(ErrDuplicateURLParserType))

		return ErrDuplicateURLParserType
	}

	p.mapFuncParseURL.Store(name, funcParse)

	return nil
}

// get - get
func (p *URLParser) get(name string) FuncParseURL {
	v, ok := p.mapFuncParseURL.Load(name)
	if ok {
		f, typeok := v.(FuncParseURL)
		if typeok {
			return f
		}

		jarvisbase.Warn("URLParser.get", zap.Error(ErrInvalidURLParserType))

		return nil
	}

	return nil
}

// ParseURL - parse url
func (p *URLParser) ParseURL(url string) *URLResult {
	var firstret *URLResult

	// jarvisbase.Info("URLParser.ParseURL", zap.String("url", url))

	p.mapFuncParseURL.Range(func(key, val interface{}) bool {
		f, typeok := val.(FuncParseURL)
		if typeok {
			ret := f(url)
			if ret != nil {
				firstret = ret

				// break
				return false
			}
		}

		return true
	})

	// if firstret != nil {
	// 	jarvisbase.Info("URLParser.ParseURL", jarvisbase.JSON("ret", firstret))
	// }

	return firstret
}
