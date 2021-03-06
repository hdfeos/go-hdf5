go-hdf5
=======
[![Build Status](https://drone.io/github.com/sbinet/go-hdf5/status.png)](https://drone.io/github.com/sbinet/go-hdf5/latest)
[![GoDoc](https://godoc.org/github.com/sbinet/go-hdf5?status.svg)](https://godoc.org/github.com/sbinet/go-hdf5)

Naive ``cgo`` bindings for the ``C-API`` of ``hdf5``.

Documentation
-------------

http://godoc.org/github.com/sbinet/go-hdf5

Example
-------

- Hello world example: https://github.com/sbinet/go-hdf5/blob/master/cmd/test-go-hdf5/main.go

- Writing/reading an ``hdf5`` with compound data: https://github.com/sbinet/go-hdf5/blob/master/cmd/test-go-cpxcmpd/main.go

Note
----

- *Only* version *1.8.x* of ``HDF5`` is supported.
- In order to use ``HDF5`` functions in more than one goroutine simultaneously, you must build the HDF5 library with threading support. Many binary distributions (RHEL/centos/Fedora packages, etc.) do not have this enabled. Therefore, you must build HDF5 yourself on these systems.


Known problems
--------------

- the ``h5pt`` packet table interface is broken.
- support for structs with slices and strings as fields is broken
