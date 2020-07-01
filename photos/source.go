package photos

import (
	"bytes"
	"context"
	"os/exec"
	"path"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	_cached  = make(map[string]struct{}, 10)
	_data    chan string
	t        *time.Ticker
	running  bool
	interval = time.Second
)

func Listen(dir string) (<-chan string, error) {
	if running {
		return nil, errors.New("already started")
	}
	running = true
	_data = make(chan string, 10)
	t = time.NewTicker(interval)
	if _, err := scan(dir); err != nil {
		zap.L().Error("Listen", zap.Error(err))
	}
	go func() {
		for range t.C {
			files, err := scan(dir)
			if err != nil {
				zap.L().Error("Listen", zap.Error(err))
				continue
			}
			for _, f := range files {
				_data <- f
			}
		}
	}()
	return _data, nil
}

func Stop() error {
	close(_data)
	running = false
	t.Stop()
	return nil
}

func scan(dir string) ([]string, error) {
	l := zap.L()
	result := make([]string, 0, 10)
	ctx, _ := context.WithTimeout(context.Background(), interval)
	cmd := exec.CommandContext(ctx, "adb", "shell", "ls -1t "+dir+" | head")
	stdErr, stdOut := bytes.NewBuffer(nil), bytes.NewBuffer(nil)
	cmd.Stderr = stdErr
	cmd.Stdout = stdOut
	if err := cmd.Run(); err != nil {
		return nil, errors.Wrap(err, "scan "+cmd.String()+":\n"+stdErr.String())
	}
	if stdErr.Len() > 0 {
		l.Error(cmd.String() + ": " + stdErr.String())
	}
	rows := bytes.Split(stdOut.Bytes(), []byte("\n"))
	for _, row := range rows {
		s := string(bytes.TrimSpace(row))
		if s == "" {
			continue
		}
		if _, ok := _cached[s]; ok {
			break
		}
		_cached[s] = struct{}{}
		result = append(result, path.Join(dir, s))
	}
	return result, nil
}
