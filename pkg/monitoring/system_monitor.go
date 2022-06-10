package monitoring

import (
	"fmt"
	"time"

	"github.com/millbj92/synctl/pkg/models/system"
	"github.com/shirou/gopsutil/cpu"
)

func GetSystemInfo() (info *system.SystemStats, err error) {
systemInfo := &system.SystemStats{}
physicalCnt, err := cpu.Counts(false); if err != nil {
	return nil, err
}
systemInfo.PhysicalCnt = physicalCnt
logicalCnt, err := cpu.Counts(true);  if err != nil {
	return nil, err
}
systemInfo.LogicalCnt = logicalCnt
fmt.Printf("physical count:%d logical count:%d\n", physicalCnt, logicalCnt)

// Obtain the total CPU utilization rate and the respective utilization rate of each CPU within 3s
totalPercent, err := cpu.Percent(3*time.Second, false); if err != nil {
	return nil, err
}
systemInfo.TotalPercent = totalPercent
perPercents, _ := cpu.Percent(3*time.Second, true); if err != nil {
	return nil, err
}
systemInfo.PerPercents = perPercents
fmt.Printf("total percent:%v per percents:%v", totalPercent, perPercents)


return systemInfo, nil
	// // Obtain the total memory utilization rate and the respective utilization rate of each memory within 3s
	// totalPercent, _ = mem.Percent(false)	// Total memory utilization
	// systemInfo.TotalMemoryUsage = totalPercent

}
