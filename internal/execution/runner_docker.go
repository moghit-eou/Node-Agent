package execution

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

 type Result struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

 func RunCommand(command string) (*Result, error) {
	ctx := context.Background()

 	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to docker: %v", err)
	}
	defer cli.Close()

 	config := &container.Config{
		Image: "alpine",
		Cmd:   []string{"sh", "-c", command},
		Tty:   false, // to separate stdout and stderr streams
	}

 	resp, err := cli.ContainerCreate(ctx, config, nil, nil, nil, "")
	if err != nil {
		return nil, fmt.Errorf("container create failed: %v", err)
	}
	containerID := resp.ID

	defer func() {
		_ = cli.ContainerRemove(ctx, containerID, container.RemoveOptions{Force: true})
	}()

	if err := cli.ContainerStart(ctx, containerID, container.StartOptions{}); err != nil {
		return nil, fmt.Errorf("container start failed: %v", err)
	}

 	//  job 10 seconds to finish, otherwise we kill it.
	waitCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	statusCh, errCh := cli.ContainerWait(waitCtx, containerID, container.WaitConditionNotRunning)
	select {
		case err := <-errCh:
			if err != nil {
				return nil, fmt.Errorf("execution error: %v", err)
			}
		case <-statusCh:
		case <-waitCtx.Done():
			return nil, fmt.Errorf("execution timed out")
	}

 	out, err := cli.ContainerLogs(ctx, containerID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return nil, fmt.Errorf("log retrieval failed: %v", err)
	}

 	var stdout, stderr bytes.Buffer
	stdcopy.StdCopy(&stdout, &stderr, out)

	return &Result{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		ExitCode: 0,
	}, nil
}