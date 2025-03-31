package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/docker/docker/api/types/container"
	"github.com/gogf/gf/v2/frame/g"
)

// ContainerStats container status struct
type ContainerStats struct {
	CPU      float64 `json:"cpu_percent"`
	Memory   float64 `json:"memory_percent"`
	MemoryMB float64 `json:"memory_mb"`
	NetIO    struct {
		In  float64 `json:"in"`
		Out float64 `json:"out"`
	} `json:"net_io"`
	BlockIO struct {
		In  float64 `json:"in"`
		Out float64 `json:"out"`
	} `json:"block_io"`
}

// DockerStatsJSON container statistics JSON struct
type DockerStatsJSON struct {
	CPUStats struct {
		CPUUsage struct {
			TotalUsage        uint64   `json:"total_usage"`
			PercpuUsage       []uint64 `json:"percpu_usage"`
			UsageInKernelmode uint64   `json:"usage_in_kernelmode"`
			UsageInUsermode   uint64   `json:"usage_in_usermode"`
		} `json:"cpu_usage"`
		SystemUsage    uint64 `json:"system_usage"`
		OnlineCPUs     uint32 `json:"online_cpus"`
		ThrottlingData struct {
			Periods          uint64 `json:"periods"`
			ThrottledPeriods uint64 `json:"throttled_periods"`
			ThrottledTime    uint64 `json:"throttled_time"`
		} `json:"throttling_data"`
	} `json:"cpu_stats"`
	PreCPUStats struct {
		CPUUsage struct {
			TotalUsage        uint64   `json:"total_usage"`
			PercpuUsage       []uint64 `json:"percpu_usage"`
			UsageInKernelmode uint64   `json:"usage_in_kernelmode"`
			UsageInUsermode   uint64   `json:"usage_in_usermode"`
		} `json:"cpu_usage"`
		SystemUsage    uint64 `json:"system_usage"`
		OnlineCPUs     uint32 `json:"online_cpus"`
		ThrottlingData struct {
			Periods          uint64 `json:"periods"`
			ThrottledPeriods uint64 `json:"throttled_periods"`
			ThrottledTime    uint64 `json:"throttled_time"`
		} `json:"throttling_data"`
	} `json:"precpu_stats"`
	MemoryStats struct {
		Usage    uint64 `json:"usage"`
		MaxUsage uint64 `json:"max_usage"`
		Limit    uint64 `json:"limit"`
	} `json:"memory_stats"`
	Networks map[string]struct {
		RxBytes   uint64 `json:"rx_bytes"`
		RxPackets uint64 `json:"rx_packets"`
		RxErrors  uint64 `json:"rx_errors"`
		RxDropped uint64 `json:"rx_dropped"`
		TxBytes   uint64 `json:"tx_bytes"`
		TxPackets uint64 `json:"tx_packets"`
		TxErrors  uint64 `json:"tx_errors"`
		TxDropped uint64 `json:"tx_dropped"`
	} `json:"networks"`
	BlkioStats struct {
		IoServiceBytesRecursive []struct {
			Major uint64 `json:"major"`
			Minor uint64 `json:"minor"`
			Op    string `json:"op"`
			Value uint64 `json:"value"`
		} `json:"io_service_bytes_recursive"`
	} `json:"blkio_stats"`
}

// ExecResult execution result
type ExecResult struct {
	ExitCode int    `json:"exit_code"`
	Output   string `json:"output"`
}

type ListContainerReq struct {
	g.Meta        `path:"/docker_api/list" tags:"DockerApi" method:"get" summary:"list containers" in:"query"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
}

type ListContainerRes struct {
	api_v1.StandardRes
	Data []container.Summary `json:"data"`
}

type ContainerStateReq struct {
	g.Meta        `path:"/docker_api/state" tags:"DockerApi" method:"get" summary:"get container status" in:"query"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	ContainerID   string `json:"container_id" v:"required|min-length:1" dc:"Container ID" dc:"Container ID"`
}

type ContainerStateRes struct {
	api_v1.StandardRes
	Data *ContainerStats `json:"data"`
}

type RestartContainerReq struct {
	g.Meta        `path:"/docker_api/restart" tags:"DockerApi" method:"post" summary:"restart container" in:"query"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	ContainerID   string `json:"container_id" v:"required|min-length:1" dc:"Container ID" dc:"Container ID"`
}

type RestartContainerRes struct {
	api_v1.StandardRes
}
