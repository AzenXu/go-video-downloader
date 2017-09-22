package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type WDRMobileIE struct {
	*rnt.Context
	_VALID_URL string
	IE_NAME    string
}

func NewWDRMobileIE(ctx *rnt.Context) rnt.InfoExtractor {
	ret := &WDRMobileIE{}
	ret.Context = ctx
	ret._VALID_URL = `(?x)
        https?://mobile-ondemand\.wdr\.de/
        .*?/fsk(?P<age_limit>[0-9]+)
        /[0-9]+/[0-9]+/
        (?P<id>[0-9]+)_(?P<title>[0-9]+)`
	ret.IE_NAME = `wdr:mobile`
	return ret
}

func (self *WDRMobileIE) Ctx() *rnt.Context {
	return self.Context
}

func (self *WDRMobileIE) Key() string {
	return "WDRMobile"
}

func (self *WDRMobileIE) ValidUrl() string {
	return self._VALID_URL
}

func (self *WDRMobileIE) Name() string {
	return self.IE_NAME
}

func (self *WDRMobileIE) Tests() []map[string]interface{} {
	return []map[string]interface{}{}
}

func (self *WDRMobileIE) _real_extract(url string) map[string]interface{} {
	mobj := rnt.ReMatch(rnt.Re, (self)._VALID_URL, url, 0)
	return map[string]interface{}{`id`: rnt.ReMatchGroupOne(mobj, `id`),
		`title`:        rnt.ReMatchGroupOne(mobj, `title`),
		`age_limit`:    rnt.ToInt(rnt.ReMatchGroupOne(mobj, `age_limit`)),
		`url`:          url,
		`http_headers`: map[string]interface{}{`User-Agent`: `mobile`}}
}
func (self *WDRMobileIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`WDRMobile`, NewWDRMobileIE)
}
