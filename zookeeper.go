package zkc

import (
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// Conn extends go-zookeeper's functionality.
type Conn struct {
	*zk.Conn
	rt time.Duration
}

// New creates new *zkc.Conn instance.
func New(servers []string, sessionTimeout time.Duration, refreshThreshold time.Duration) (*Conn, error) {
	conn, _, err := zk.Connect(servers, sessionTimeout)
	if err != nil {
		return nil, err
	}

	return &Conn{Conn: conn, rt: refreshThreshold}, nil
}

// ChildrenC implements Zookeeper's getChildren method with watcher continuosly sending events.
func (c *Conn) ChildrenC(path string) <-chan ChildrenEvent {
	events := make(chan ChildrenEvent)
	timer := time.NewTimer(c.rt)

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
func (c *Conn) GetC(path string) <-chan DataEvent {
	events := make(chan DataEvent)
	timer := time.NewTimer(c.rt)

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

// ExistsC implements Zookeeper's exists method with watcher continuosly sending events.
func (c *Conn) ExistsC(path string) <-chan ExistsEvent {
	events := make(chan ExistsEvent)
	timer := time.NewTimer(c.rt)

	go func(path string) {
		for {
			exists, stat, event, err := c.Conn.ExistsW(path)
			events <- ExistsEvent{
				Exists: exists,
				Stat:   stat,
				Err:    err,
			}
			select {
			case e := <-event:
				events <- ExistsEvent{
					Exists: exists,
					Stat:   stat,
					Evt:    &e,
					Err:    err,
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

// ExistsEvent contains aggregated data from zk.ExistsW.
type ExistsEvent struct {
	Exists bool
	Stat   *zk.Stat
	Evt    *zk.Event
	Err    error
}
