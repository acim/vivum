
# vivum

An extension of go-zookeeper package.

go-zookeeper provides methods ChildrenW, ExistsW and GetW which return just single event after which you have the corresponding method again.
This is native behavior of Zookeeper's functions getChildren, exists and getData respectively.
This project provides additional methods ChildrenC, ExistsC and GetC which will continuously send events so that you don't have to call the methods all over again. It also provides a way to get these methods called automatically after some threshold duration in case we loose some events during reattaching the listeners.

Original project is available at [here](https://github.com/samuel/go-zookeeper).

# Requirements

* Go 1.11