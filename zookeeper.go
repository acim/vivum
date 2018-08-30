package zkv

import (
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

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
func (c *Conn) GetC(path string) ([]byte, *zk.Stat, <-chan zk.Event, error) {
	return c.Conn.GetW(path)
}
