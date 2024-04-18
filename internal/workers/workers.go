package workers

import (
	"context"
)

type Daemon interface {
	Do(ctx context.Context) error
}

type DaemonServer struct {
	ctx     context.Context
	daemons []Daemon
}

func New(ctx context.Context) *DaemonServer {
	return &DaemonServer{
		ctx: ctx,
	}
}

func (w *DaemonServer) Add(daemons ...Daemon) {
	w.daemons = append(w.daemons, daemons...)
}

// Start starts the DaemonServer by running the daemons in separate goroutines.
func (s *DaemonServer) Start() {
	for _, daemon := range s.daemons {
		daemon := daemon
		go func() {
			daemon.Do(s.ctx)
		}()
	}
}

// Stop stops the DaemonServer.
func (s *DaemonServer) Stop() {
	s.ctx.Done()
}
