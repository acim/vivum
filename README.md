
# vivum

An extension of [go-zookeeper](https://github.com/samuel/go-zookeeper) package.

go-zookeeper provides methods ChildrenW, ExistsW and GetW which return just single event after which you have to call the corresponding method again and again. Though, this is also native behavior of Zookeeper's functions getChildren, exists and getData respectively.
Thanks to Go concurrency model, this project extends the package providing additional methods ChildrenC, ExistsC and GetC which will continuously send events so that you don't have to call the methods all over again. It also provides a way to get these methods called automatically after some threshold time in case you loose some events during the reattachment of listeners. But full functionality of go-zookeeper is kept so you can use this module as a replacement.

# Requirements

* Go 1.11 (just package management, otherwise comaptible with older versions)

# TODO

* Docker container providing environment for testing
* unit tests