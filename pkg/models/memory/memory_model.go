package memory

type AllMemoryResponse struct {
	Error bool         `json:"error"`
	Msg   string       `json:"msg"`
	Data  *MemoryStats `json:"data"`
}

type MemoryStats struct {
	VirtualMemory *VirtualMemory
	SwapMemory    *SwapMemory
	SwapDevices   []*SwapDevice
}
type MemoryStat struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

type SwapMemory struct {
	MemoryStat
	Free uint64 `json:"free"`
}

type VirtualMemory struct {
	MemoryStat
	Available uint64 `json:"available"`
}

type SwapDevice struct {
	Name      string `json:"name"`
	UsedBytes uint64 `json:"usedBytes"`
	FreeBytes uint64 `json:"freeBytes"`
}
