
# vivum

An extension of [go-zookeeper](https://github.com/samuel/go-zookeeper) package.

go-zookeeper provides methods ChildrenW, ExistsW and GetW which return just single event after which you have to call the corresponding method again and again. Though, this is also native behavior of Zookeeper's functions getChildren, exists and getData respectively.
This project extends the package providing additional methods ChildrenC, ExistsC and GetC which will continuously send events so that you don't have to call the methods all over again. It also provides a way to get these methods called automatically after some threshold duration in case we loose some events during reattaching the listeners. It also keeps full functionality of go-zookeeper so you can use this package as a replacement.

# Requirements

* Go 1.11