openfoxwq
===

This repository contains my attempt to reverse engineer the [Fox Go Server](https://foxwq.com/) protocol and build a client that is usable in other platforms (at least Linux and OSX). The only native client available at the moment is for Windows and the mobile clients (Android and iOS), and they are not entirely localized although this is a lesser problem for most people.

Build Instructions
===

Even though this repository contains only the server component, most people want to build the UI component as well so these instructions cover both.

Dependencies:
- [Go](https://go.dev/doc/install)
- [protobuf](https://protobuf.dev/downloads/) compiler with [Go plugin](https://protobuf.dev/getting-started/gotutorial/#compiling-protocol-buffers)
- [Qt](https://www.qt.io/offline-installers) (I use the 5.12.X offline installer)

Note: make sure that the Go binary path is in your `PATH` variable (e.g. `export PATH="$PATH:$HOME/go/bin"`). Otherwise the protobuf plugin for Go might not be found.

1. Clone both repositories into the same parent directory:
```
$ git clone https://github.com/openfoxwq/openfoxwq.git
$ git clone https://github.com/openfoxwq/openfoxwq_qtclient.git
```

2. Move to the server directory and compile the protos:
```
$ cd openfoxwq
$ ./compile_protos.sh
```

3. Build the server component:
```
$ go build -o openfoxwq_wsserver cmd/wsserver/main.go
```

This should create the server binary named `openfoxwq_wsserver` in the current directory.

4. Move to the UI directory and build the client:
```
$ cd ../openfoxwq_qtclient
$ qmake
$ make
```

This should create the client binary named `openfoxwq_qtclient` in the current directory.
