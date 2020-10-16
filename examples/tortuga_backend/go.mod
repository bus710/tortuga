module main

go 1.12

require github.com/bus710/tortuga v0.0.0
replace github.com/bus710/tortuga => ../../../tortuga

// require	github.com/bus710/tortuga/cmd/command v0.0.0
// replace github.com/bus710/tortuga/cmd/command  => ../../../tortuga/cmd/command
// require	github.com/bus710/tortuga/cmd/model v0.0.0
// replace github.com/bus710/tortuga/cmd/model => ../../../tortuga/cmd/model

require github.com/gorilla/websocket v1.4.1
require golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297
