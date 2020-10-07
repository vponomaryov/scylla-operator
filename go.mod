module github.com/scylladb/scylla-operator

go 1.13

require (
	github.com/blang/semver v3.5.0+incompatible
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/runtime v0.19.22
	github.com/go-openapi/strfmt v0.19.5
	github.com/gocql/gocql v0.0.0-20200926162733-393f0c961220
	github.com/google/go-cmp v0.4.0
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed
	github.com/hashicorp/go-version v1.2.1
	github.com/magiconair/properties v1.8.0
	github.com/mitchellh/mapstructure v1.3.2
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/pkg/errors v0.9.1
	github.com/scylladb/go-log v0.0.4
	github.com/scylladb/go-set v1.0.2
	github.com/scylladb/scylla-mgmt-commons/format v0.0.0-20201007140813-8ae21b32e0d5
	github.com/scylladb/scylla-mgmt-commons/managerclient v0.0.0-20201007140813-8ae21b32e0d5
	github.com/scylladb/scylla-mgmt-commons/middleware v0.0.0-20201007141033-a2b414e94e66
	github.com/scylladb/scylla-mgmt-commons/scyllaclient v0.0.0-20201007140813-8ae21b32e0d5
	github.com/scylladb/scylla-mgmt-commons/uuid v0.0.0-20201007140813-8ae21b32e0d5
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.6.1
	go.uber.org/config v1.4.0
	go.uber.org/multierr v1.4.0
	go.uber.org/zap v1.14.0
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.18.4
	k8s.io/apimachinery v0.18.4
	k8s.io/client-go v0.18.4
	k8s.io/utils v0.0.0-20200603063816-c1c6865ac451
	sigs.k8s.io/controller-runtime v0.6.1
)
