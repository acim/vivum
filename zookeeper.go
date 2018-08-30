package zkc

import (
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// Conn extends go-zookeeper's functionality.
type Conn struct {
	*zk.Conn
}

// New creates new *zkc.Conn instance.
func New(servers []string, sessionTimeout time.Duration) (*Conn, error) {
	conn, _, err := zk.Connect(servers, sessionTimeout)
	if err != nil {
		return nil, err
	}

	return &Conn{Conn: conn}, nil
}

// ChildrenC implements Zookeeper's getChildren method with watcher continuosly sending events.
func (c *Conn) ChildrenC(path string, refreshThreshold time.Duration) <-chan ChildrenEvent {
	events := make(chan ChildrenEvent)
	timer := time.NewTimer(refreshThreshold)

	go func(path string) {
		for {
			children, stat, event, err := c.Conn.ChildrenW(path)
			events <- ChildrenEvent{
				Children: children,
				Stat:     stat,
				Err:      err,
			}
			select {
			case e := <-event:
				events <- ChildrenEvent{
					Children: children,
					Stat:     stat,
					Evt:      &e,
					Err:      err,
				}
			case <-timer.C:
			}
		}
	}(path)

	return events
}

// GetC implements Zookeeper's getData method with watcher continuosly sending events.
func (c *Conn) GetC(path string, refreshThreshold time.Duration) <-chan DataEvent {
	events := make(chan DataEvent)
	timer := time.NewTimer(refreshThreshold)

	go func(path string) {
		for {
			data, stat, event, err := c.Conn.GetW(path)
			events <- DataEvent{
				Data: data,
				Stat: stat,
				Err:  err,
			}
			select {
			case e := <-event:
				events <- DataEvent{
					Data: data,
					Stat: stat,
					Evt:  &e,
					Err:  err,
				}
			case <-timer.C:
			}
		}
	}(path)

	return events
}

// DataEvent contains aggregated data from zk.GetW.
type DataEvent struct {
	Data []byte
	Stat *zk.Stat
	Evt  *zk.Event
	Err  error
}

// ChildrenEvent contains aggregated data from zk.ChildrenW.
type ChildrenEvent struct {
	Children []string
	Stat     *zk.Stat
	Evt      *zk.Event
	Err      error
}
