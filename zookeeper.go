package zkv

import (
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// Conn extends go-zookeeper functionality.
type Conn struct {
	*zk.Conn
}

// New creates new *Conn instance.
func New(servers []string, sessionTimeout time.Duration) (*Conn, error) {
	conn, _, err := zk.Connect(servers, sessionTimeout)
	if err != nil {
		return nil, err
	}

	return &Conn{Conn: conn}, nil
}

// ChildrenC implements Zookeeper's getChildren method with watcher continuosly sending events.
func (c *Conn) ChildrenC(path string) ([]string, *zk.Stat, <-chan zk.Event, error) {
	return c.Conn.ChildrenW(path)
}

// GetC implements Zookeeper's getData method with watcher continuosly sending events.
func (c *Conn) GetC(path string, refreshThreshold time.Duration) <-chan Event {
	events := make(chan Event)
	timer := time.NewTimer(refreshThreshold)

	go func(path string) {
		for {
			data, stat, event, err := c.Conn.GetW(path)
			events <- Event{
				Data: data,
				Stat: stat,
				Err:  err,
			}
			select {
			case e := <-event:
				events <- Event{
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

// Event contains aggregate data from zk.ChildrenW.
type Event struct {
	Data []byte
	Stat *zk.Stat
	Evt  *zk.Event
	Err  error
}
