dweller  
---
[![Build Status](https://travis-ci.org/cloudflavor/dweller.svg?branch=master)](https://travis-ci.org/cloudflavor/dweller)
[![Codecov](https://codecov.io/gh/cloudflavor/dweller/branch/master/graph/badge.svg)](https://codecov.io/gh/cloudflavor/dweller)
[![Documentation](https://godoc.org/github.com/cloudflavor/dweller?status.svg)](http://godoc.org/github.com/cloudflavor/dweller/)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudflavor/dweller)](https://goreportcard.com/report/github.com/cloudflavor/dweller)

Sets up the `Cloudflavor` infrastructure on different providers. Currently
`libvirt` with `qemu` is supported and `xen` to be implemented in the future.

dweller manages opinionated virtualized infrastructure and is only designed to be fast in bringing it up.
It's not as flexible as `terraform`.

Why dweller?

* is fast to work with
* ships one binary
* no DSL - configuration is one *simple* `yaml` file that contains the `master/worker` spec and the `qemu` connections pool. (see the [example config file](example-infra.yaml))
* manages multiple `qemu` nodes to scale the infrastructure horizontally

The default provisioning is 1xMaster and 2xWorkers but the default number of
workers can be overridden. The infrastructure scales horizontally across
multiple machines that have libvirt installed. To setup the bare metal machines
to be able to support this, you will need to run [Cloudflavor's
ansible-infra](https://github.com/cloudflavor/ansible-infra).

```
dw - A CLI for provisioning a new Cloudflavor infrastructure

Usage:
  dw [flags]
  dw [command]

Available Commands:
  delete      Delete a specific instance.
  halt        Halts the currently running infrastructure
  help        Help about any command
  new         Add a new instance to an already running cloudflavor infrastructure
  up          Bring up a simple Cloudflavor infrastructure

Flags:
  -h, --help   help for dw

Use "dw [command] --help" for more information about a command.
```
