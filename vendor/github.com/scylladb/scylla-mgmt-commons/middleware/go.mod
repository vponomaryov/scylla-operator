module github.com/scylladb/scylla-mgmt-commons/middleware

go 1.14

require (
	github.com/bitly/go-hostpool v0.1.0 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/go-openapi/runtime v0.19.22
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed
	github.com/pkg/errors v0.9.1
	github.com/scylladb/go-log v0.0.4
	github.com/scylladb/scylla-mgmt-commons/httpx v0.0.0-20201007133817-3a5a9f249639
	github.com/scylladb/scylla-mgmt-commons/retry v0.0.0-20201007133817-3a5a9f249639
	github.com/scylladb/scylla-mgmt-commons/timeutc v0.0.0-20201007133817-3a5a9f249639
)

replace (
	github.com/scylladb/scylla-mgmt-commons/httpx => ../httpx
	github.com/scylladb/scylla-mgmt-commons/retry => ../retry
	github.com/scylladb/scylla-mgmt-commons/timeutc => ../timeutc
)
