package main

import (
	"bytes"
	"context"
	"fmt"
 	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
)

type Result struct 
{
	Stdout string 
	Stderr string 
	ExitCode int 
}

func (d *DockerClient) Run(ctx context.Context, command string) (*Result, error) {
	
	// 1. CONFIG: Use Alpine Linux
	config := &container.Config{
		Image: "alpine",
		Cmd:   []string{"sh", "-c", command},
	}

	// 2. CREATE the container
	resp, err := d.cli.ContainerCreate(ctx, config, nil, nil, nil, "")
	if err != nil {
		return nil, fmt.Errorf("create failed: %w", err)
	}
	containerID := resp.ID

	// 3. CLEANUP: Always remove the container, even if it crashes
	defer func() {
		_ = d.cli.ContainerRemove(ctx, containerID, container.RemoveOptions{Force: true})
	}()

	// 4. START the container
	if err := d.cli.ContainerStart(ctx, containerID, container.StartOptions{}); err != nil {
		return nil, fmt.Errorf("start failed: %w", err)
	}

	// 5. WAIT for it to finish
	statusCh, errCh := d.cli.ContainerWait(ctx, containerID, container.WaitConditionNotRunning)
	select {
		case err := <-errCh:
			if err != nil {
				return nil, fmt.Errorf("wait error: %w", err)
			}
		case <-statusCh:
			// Success
		case <-ctx.Done():
			return nil, fmt.Errorf("timeout reached")
	}

	// 6. LOGS: Fetch and clean the output
		out, err := d.cli.ContainerLogs(ctx, containerID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
		if err != nil {
			return nil, fmt.Errorf("logs failed: %w", err)
		}

	// Docker logs have headers. We use stdcopy to strip them.
	var stdout, stderr bytes.Buffer
	stdcopy.StdCopy(&stdout, &stderr, out)

	return &Result{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		ExitCode: 0, // Simplified for now
	}, nil
}