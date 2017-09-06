dweller  
---
[![Build Status](https://travis-ci.org/cloudflavor/dweller.svg?branch=master)](https://travis-ci.org/cloudflavor/dweller)
[![codecov](https://codecov.io/gh/cloudflavor/dweller/branch/master/graph/badge.svg)](https://codecov.io/gh/cloudflavor/dweller)

Setups up the `Cloudflavor` infrastructure on different providers. Currently
`libvirt` with `qemu` is supported and `xen` to be implemented in the future.

The default provisioning is 1xMaster and 2xWorkers but it can be overridden.

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
