dweller  
---
[![Build Status](https://travis-ci.org/cloudflavor/dweller.svg?branch=master)](https://travis-ci.org/cloudflavor/dweller)
[![Codecov](https://codecov.io/gh/cloudflavor/dweller/branch/master/graph/badge.svg)](https://codecov.io/gh/cloudflavor/dweller)
[![Documentation](https://godoc.org/github.com/cloudflavor/dweller?status.svg)](http://godoc.org/github.com/cloudflavor/dweller/)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudflavor/dweller)](https://goreportcard.com/report/github.com/cloudflavor/dweller)

Sets up the `Cloudflavor` infrastructure on different providers. Currently
`libvirt` with `qemu` is supported and `xen` to be implemented in the future.

`dweller` manages opinionated virtualized infrastructure and is only designed to be fast in manipulating it.

* no DSL - configuration is one *simple* `yaml` file that contains the `master/worker` spec and the `qemu` connections pool. (see the [example config file](example-infra.yaml))
* manages multiple `qemu/xen` nodes to scale the infrastructure horizontally

The default provisioning is 1xMaster and 2xWorkers but the default number of
workers can be overridden. The infrastructure scales horizontally across
multiple machines that have `qemu or xen` installed.  


#### Building

Running `make` in the root directory will create the binary in `_output/bin/dw`.


#### Architecture

![dweller arch](assets/dweller-cloudflavor.png)
