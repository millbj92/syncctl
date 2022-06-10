package memory

import (
	"github.com/shirou/gopsutil/v3/mem"
)

type MemoryStats struct {
	VirtualMemory *mem.VirtualMemoryStat
	SwapMemory    *mem.SwapMemoryStat
	SwapDevices   []*mem.SwapDevice
}
