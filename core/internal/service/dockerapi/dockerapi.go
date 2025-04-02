package docker

import (
	v1 "billionmail-core/api/dockerapi/v1"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

// DockerAPI struct
type DockerAPI struct {
	client     *client.Client
	containers map[string]bool
}

// NewDockerAPI creates a new DockerAPI instance
func NewDockerAPI() (*DockerAPI, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker client: %v", err)
	}

	return &DockerAPI{
		client:     cli,
		containers: make(map[string]bool),
	}, nil
}

// StartContainer starts a container
func (d *DockerAPI) StartContainer(ctx context.Context, containerID string) error {
	return d.client.ContainerStart(ctx, containerID, container.StartOptions{})
}

// StopContainer stops a container
func (d *DockerAPI) StopContainer(ctx context.Context, containerID string) error {
	timeout := gconv.Int(10 * time.Second)
	return d.client.ContainerStop(ctx, containerID, container.StopOptions{
		Timeout: &timeout,
	})
}

// RestartContainer restarts a container
func (d *DockerAPI) RestartContainer(ctx context.Context, containerID string) error {
	timeout := gconv.Int(10 * time.Second)
	return d.client.ContainerRestart(ctx, containerID, container.StopOptions{
		Timeout: &timeout,
	})
}

// RestartContainerByName restarts a container by its name
func (d *DockerAPI) RestartContainerByName(ctx context.Context, name string) error {
	ctnr, err := d.GetContainerByName(ctx, name)
	if err != nil {
		return err
	}

	return d.RestartContainer(ctx, ctnr.ID)
}

// GetContainerStats gets container status
func (d *DockerAPI) GetContainerStats(ctx context.Context, containerID string) (*v1.ContainerStats, error) {
	stats, err := d.client.ContainerStats(ctx, containerID, false)
	if err != nil {
		return nil, err
	}
	defer stats.Body.Close()

	var statsJSON v1.DockerStatsJSON
	if err := json.NewDecoder(stats.Body).Decode(&statsJSON); err != nil {
		return nil, err
	}

	// Calculate CPU usage
	cpuPercent := calculateCPUPercentUnix(statsJSON)

	// Calculate memory usage
	memoryPercent := float64(statsJSON.MemoryStats.Usage) / float64(statsJSON.MemoryStats.Limit) * 100.0
	memoryMB := float64(statsJSON.MemoryStats.Usage) / 1024 / 1024

	// Calculate network IO
	var netIn, netOut float64
	for _, v := range statsJSON.Networks {
		netIn += float64(v.RxBytes)
		netOut += float64(v.TxBytes)
	}

	// Calculate disk IO
	var blockIn, blockOut float64
	for _, bioEntry := range statsJSON.BlkioStats.IoServiceBytesRecursive {
		if bioEntry.Op == "Read" {
			blockIn += float64(bioEntry.Value)
		}
		if bioEntry.Op == "Write" {
			blockOut += float64(bioEntry.Value)
		}
	}

	return &v1.ContainerStats{
		CPU:      cpuPercent,
		Memory:   memoryPercent,
		MemoryMB: memoryMB,
		NetIO: struct {
			In  float64 `json:"in"`
			Out float64 `json:"out"`
		}{
			In:  netIn,
			Out: netOut,
		},
		BlockIO: struct {
			In  float64 `json:"in"`
			Out float64 `json:"out"`
		}{
			In:  blockIn,
			Out: blockOut,
		},
	}, nil
}

// Calculate CPU usage percentage
func calculateCPUPercentUnix(stats v1.DockerStatsJSON) float64 {
	cpuPercent := 0.0
	// Calculate CPU usage
	cpuDelta := float64(stats.CPUStats.CPUUsage.TotalUsage) - float64(stats.PreCPUStats.CPUUsage.TotalUsage)
	systemDelta := float64(stats.CPUStats.SystemUsage) - float64(stats.PreCPUStats.SystemUsage)

	if systemDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = (cpuDelta / systemDelta) * float64(len(stats.CPUStats.CPUUsage.PercpuUsage)) * 100.0
	}
	return cpuPercent
}

// ExecCommand executes a command in a container
func (d *DockerAPI) ExecCommand(ctx context.Context, containerID string, cmd []string, user string) (*v1.ExecResult, error) {
	execConfig := container.ExecOptions{
		Cmd:          cmd,
		User:         user,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          false,
	}

	execID, err := d.client.ContainerExecCreate(ctx, containerID, execConfig)
	if err != nil {
		return nil, err
	}

	resp, err := d.client.ContainerExecAttach(ctx, execID.ID, container.ExecAttachOptions{})
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	// Use stdcopy to copy output
	var outBuf, errBuf strings.Builder
	_, err = stdcopy.StdCopy(&outBuf, &errBuf, resp.Reader)
	if err != nil {
		return nil, err
	}

	// Get exit code
	inspect, err := d.client.ContainerExecInspect(ctx, execID.ID)
	if err != nil {
		return nil, err
	}

	output := outBuf.String()
	if errBuf.String() != "" {
		if output != "" {
			output += "\n"
		}
		output += errBuf.String()
	}

	return &v1.ExecResult{
		ExitCode: inspect.ExitCode,
		Output:   output,
	}, nil
}

// ExecCommandByName executes a command in a container by its name
func (d *DockerAPI) ExecCommandByName(ctx context.Context, name string, cmd []string, user string) (*v1.ExecResult, error) {
	ctnr, err := d.GetContainerByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return d.ExecCommand(ctx, ctnr.ID, cmd, user)
}

// ListContainers lists all containers
func (d *DockerAPI) ListContainers(ctx context.Context) ([]container.Summary, error) {
	return d.client.ContainerList(ctx, container.ListOptions{All: true})
}

// GetContainerByName gets a container by its name
func (d *DockerAPI) GetContainerByName(ctx context.Context, name string) (*container.Summary, error) {
	args := filters.NewArgs()
	args.Add("name", name)

	containers, err := d.client.ContainerList(ctx, container.ListOptions{
		All:     true,
		Filters: args,
	})

	if err != nil {
		return nil, err
	}

	if len(containers) == 0 {
		return nil, fmt.Errorf("container with name %s not found", name)
	}

	return &containers[0], nil
}

// GetContainer gets container details
func (d *DockerAPI) GetContainer(ctx context.Context, containerID string) (container.InspectResponse, error) {
	return d.client.ContainerInspect(ctx, containerID)
}

// Close closes the Docker client connection
func (d *DockerAPI) Close() error {
	return d.client.Close()
}
