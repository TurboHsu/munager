package structure

const (
	HandshakeMagicWord = "Magic smoke is not good"
)

type Handshake struct {
	MagicWord   string `json:"magic_word"`
	Fingerprint string `json:"fingerprint"`
}

type ListResponse struct {
	Files []FileInfo `json:"files"`
}

type FileServeRequest struct {
	Fingerprint string `json:"fingerprint"`
	PathBase    string `json:"path_base"`
	Extension   string `json:"extension"`
}

type BasicResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ListRequest struct {
	Files       []FileInfo `json:"files"`
	Fingerprint string     `json:"fingerprint"`
}

type FileInfo struct {
	PathBase  string `json:"path_base"`
	Extension string `json:"extension"`
}
