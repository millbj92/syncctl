package monitoring

import (
	//"fmt"
	//"os"

	//"io/fs"

	"github.com/millbj92/synctl/pkg/models/memory"
	"github.com/shirou/gopsutil/v3/mem"
	//github.com/variantdev/chartify
	//https://github.com/variantdev/vals
	//"gopkg.in/yaml.v3"
	//"github.com/gosuri/uitable"
	//"github.com/logrusorgru/aurora"
	//"github.com/tatsushid/go-prettytable"
)

func GetSwapDevices() (swapdvcs []*mem.SwapDevice, err error) {
	v, err := mem.SwapDevices(); if err != nil {
		return nil, err
	}
	return v,nil
}

func GetSwapUsage() (swap *mem.SwapMemoryStat, err error) {
	v, err := mem.SwapMemory(); if err != nil {
		return swap, err
}
	return v,nil
}

func GetMemoryUsage() (vmem *mem.VirtualMemoryStat, err error) {
	m, err := mem.VirtualMemory(); if err != nil {
		return vmem, err
	}

	return m,nil
}

func GetAllMemoryStats() (stats *memory.MemoryStats, err error) {
	stats = &memory.MemoryStats{}
	stats.VirtualMemory, err = GetMemoryUsage()
	if err != nil {
		return nil, err
	}
	stats.SwapMemory, err = GetSwapUsage()
	if err != nil {
		return nil, err
	}
	stats.SwapDevices, err = GetSwapDevices()
	if err != nil {
		return nil, err
	}

	return stats, nil
}










