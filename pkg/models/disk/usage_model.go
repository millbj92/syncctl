package disk

type DiskUsage struct {
	Path        string  `json:"path"`
	Fstype      string  `json:"fstype"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

type DiskResponse struct {
	Error bool       `json:"error"`
	Msg   string     `json:"msg"`
	Data  *DiskUsage `json:"data"`
}
