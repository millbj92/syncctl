package management

import (
	//"fmt"
	//"os"

	//"io/fs"

	"github.com/shirou/gopsutil/v3/mem"

	"github.com/davecgh/go-spew/spew"
	//github.com/variantdev/chartify
	//https://github.com/variantdev/vals
	//go get gopkg.in/yaml.v3
	//"github.com/gosuri/uitable"
	//"github.com/logrusorgru/aurora"
	//"github.com/tatsushid/go-prettytable"
)

func GetVMem() (vmem *mem.VirtualMemoryStat, err error) {
	v, err := mem.VirtualMemory(); if err != nil {
		return vmem, err
	}


	spew.Dump(v)
	return v,nil
}
