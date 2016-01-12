package config

import (
	"time"

	proto "github.com/micro/config-srv/proto/config"
)

type watcher struct {
	w proto.Config_WatchClient
}

func (w *watcher) Next() (*ChangeSet, error) {
	c, err := w.w.Recv()
	if err != nil {
		return nil, err
	}
	return &ChangeSet{
		Timestamp: time.Unix(c.ChangeSet.Timestamp, 0),
		Data:      []byte(c.ChangeSet.Data),
		Checksum:  c.ChangeSet.Checksum,
		Source:    c.ChangeSet.Source,
	}, nil
}

func (w *watcher) Stop() error {
	return w.w.Close()
}
