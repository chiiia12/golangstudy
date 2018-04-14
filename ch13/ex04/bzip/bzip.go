package bzip

import (
	"io"
	"sync"
	"os/exec"
)

type writer struct {
	sync.Mutex
	cmd *exec.Cmd
	w   io.WriteCloser
	wg  sync.WaitGroup
}

func NewWriter(out io.Writer) (io.WriteCloser, error) {
	var w writer
	w.cmd = exec.Command("bzip2")
	stdout, err := w.cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stdin, err := w.cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	w.w = stdin
	if err := w.cmd.Start(); err != nil {
		return nil, err
	}
	w.wg.Add(1)
	go func() {
		io.Copy(out, stdout)
		w.wg.Done()
	}()
	return &w, nil
}
func (w *writer) Write(data []byte) (int, error) {
	w.Lock()
	defer w.Unlock()
	var total int

	for len(data) > 0 {
		n, err := w.w.Write(data)
		if err != nil {
			return total + n, err
		}
		total += n
		data = data[total:]
	}
	return total, nil
}

func (w *writer) Close() error {
	w.Lock()
	defer w.Unlock()

	w.w.Close()
	w.wg.Wait()
	if err := w.cmd.Wait(); err != nil {
		return err
	}
	return nil
}
