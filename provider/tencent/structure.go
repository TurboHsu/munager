package tencent

type SongSearchResult struct {
	Code int `json:"code"`
	Data struct {
		Keyword  string `json:"keyword"`
		Priority int    `json:"priority"`
		Qc       []any  `json:"qc"`
		Semantic struct {
			Curnum   int   `json:"curnum"`
			Curpage  int   `json:"curpage"`
			List     []any `json:"list"`
			Totalnum int   `json:"totalnum"`
		} `json:"semantic"`
		Song struct {
			Curnum  int `json:"curnum"`
			Curpage int `json:"curpage"`
			List    []struct {
				Albumid               int    `json:"albumid"`
				Albummid              string `json:"albummid"`
				Albumname             string `json:"albumname"`
				AlbumnameHilight      string `json:"albumname_hilight"`
				Albumtransname        string `json:"albumtransname,omitempty"`
				AlbumtransnameHilight string `json:"albumtransname_hilight,omitempty"`
				Alertid               int    `json:"alertid"`
				Chinesesinger         int    `json:"chinesesinger"`
				Docid                 string `json:"docid"`
				Grp                   []any  `json:"grp"`
				Interval              int    `json:"interval"`
				Isonly                int    `json:"isonly"`
				Lyric                 string `json:"lyric"`
				LyricHilight          string `json:"lyric_hilight"`
				MediaMid              string `json:"media_mid"`
				Msgid                 int    `json:"msgid"`
				Nt                    int    `json:"nt"`
				Pay                   struct {
					Payalbum      int `json:"payalbum"`
					Payalbumprice int `json:"payalbumprice"`
					Paydownload   int `json:"paydownload"`
					Payinfo       int `json:"payinfo"`
					Payplay       int `json:"payplay"`
					Paytrackmouth int `json:"paytrackmouth"`
					Paytrackprice int `json:"paytrackprice"`
				} `json:"pay"`
				Preview struct {
					Trybegin int `json:"trybegin"`
					Tryend   int `json:"tryend"`
					Trysize  int `json:"trysize"`
				} `json:"preview"`
				Pubtime int `json:"pubtime"`
				Pure    int `json:"pure"`
				Singer  []struct {
					ID          int    `json:"id"`
					Mid         string `json:"mid"`
					Name        string `json:"name"`
					NameHilight string `json:"name_hilight"`
				} `json:"singer"`
				Size128         int    `json:"size128"`
				Size320         int    `json:"size320"`
				Sizeape         int    `json:"sizeape"`
				Sizeflac        int    `json:"sizeflac"`
				Sizeogg         int    `json:"sizeogg"`
				Songid          int    `json:"songid"`
				Songmid         string `json:"songmid"`
				Songname        string `json:"songname"`
				SongnameHilight string `json:"songname_hilight"`
				StrMediaMid     string `json:"strMediaMid"`
				Stream          int    `json:"stream"`
				Switch          int    `json:"switch"`
				T               int    `json:"t"`
				Tag             int    `json:"tag"`
				Type            int    `json:"type"`
				Ver             int    `json:"ver"`
				Vid             string `json:"vid"`
			} `json:"list"`
			Totalnum int `json:"totalnum"`
		} `json:"song"`
		Totaltime int `json:"totaltime"`
		Zhida     struct {
			Chinesesinger int `json:"chinesesinger"`
			Type          int `json:"type"`
		} `json:"zhida"`
	} `json:"data"`
	Message string `json:"message"`
	Notice  string `json:"notice"`
	Subcode int    `json:"subcode"`
	Time    int    `json:"time"`
	Tips    string `json:"tips"`
}

type LyricFetchResult struct {
	Retcode int    `json:"retcode"`
	Code    int    `json:"code"`
	Subcode int    `json:"subcode"`
	Type    int    `json:"type"`
	Songt   int    `json:"songt"`
	Lyric   string `json:"lyric"`
	Trans	string `json:"trans"`
}