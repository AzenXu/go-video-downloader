package extractor

import (
	rnt "github.com/tenta-browser/go-video-downloader/runtime"
)

type BildIE struct {
	*rnt.Context
	_VALID_URL string
	IE_DESC    string
}

func NewBildIE(ctx *rnt.Context) rnt.InfoExtractor {
	ret := &BildIE{}
	ret.Context = ctx
	ret._VALID_URL = `https?://(?:www\.)?bild\.de/(?:[^/]+/)+(?P<display_id>[^/]+)-(?P<id>\d+)(?:,auto=true)?\.bild\.html`
	ret.IE_DESC = `Bild.de`
	return ret
}

func (self *BildIE) Ctx() *rnt.Context {
	return self.Context
}

func (self *BildIE) Key() string {
	return "Bild"
}

func (self *BildIE) ValidUrl() string {
	return self._VALID_URL
}

func (self *BildIE) Name() string {
	return `Bild extractor` + " (" + self.IE_DESC + ")"
}

func (self *BildIE) Tests() []map[string]interface{} {
	return []map[string]interface{}{}
}

func (self *BildIE) _real_extract(url string) map[string]interface{} {
	video_id := rnt.MatchID(self, url)
	video_data := rnt.DownloadJSON(self, ((rnt.StrSplit(url, `.bild.html`, -(1)))[0] + `,view=json.bild.html`), video_id, rnt.AsOptString(`Downloading JSON metadata`), rnt.AsOptString(`Unable to download JSON metadata`), 0, true, rnt.OptString{}, rnt.OptString{}, map[string]interface{}{}, map[string]interface{}{})
	return map[string]interface{}{`id`: video_id,
		`title`:       rnt.StrStrip(rnt.UnescapeHTML(rnt.CastToOptString((video_data)[`title`])).Get(), ``),
		`description`: rnt.UnescapeHTML(rnt.CastToOptString(rnt.DictGet(video_data, `description`, nil))),
		`url`:         rnt.UnsafeSubscript(rnt.UnsafeSubscript(rnt.UnsafeSubscript(rnt.UnsafeSubscript((video_data)[`clipList`], 0), `srces`), 0), `src`),
		`thumbnail`:   rnt.DictGet(video_data, `poster`, nil),
		`duration`:    rnt.IntOrNone(rnt.DictGet(video_data, `durationSec`, nil), 1, rnt.OptInt{}, 1)}
}
func (self *BildIE) Extract(url string) (*rnt.VideoResult, error) {
	return rnt.RunExtractor(url, self._real_extract)
}

func init() {
	registerFactory(`Bild`, NewBildIE)
}
