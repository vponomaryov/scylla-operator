module github.com/scylladb/scylla-mgmt-commons/format

go 1.14

require (
	github.com/pkg/errors v0.9.1
	github.com/scylladb/scylla-mgmt-commons/timeutc v0.0.0-20201007115835-7e4a89cd16ab
)

replace github.com/scylladb/scylla-mgmt-commons/timeutc => ../timeutc
