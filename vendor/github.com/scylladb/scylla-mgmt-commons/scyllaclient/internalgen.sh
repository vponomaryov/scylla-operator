#!/usr/bin/env bash
#
# Copyright (C) 2017 ScyllaDB
#

set -eu -o pipefail

echo "Swagger $(swagger version)"

rm -rf scylla/client scylla/models
swagger generate client -A scylla -T templates -f scylla.json -t ./scylla

rm -rf scylla_v2/client scylla_v2/models
swagger generate client -A scylla2 -T templates -f scylla_v2.json -t ./scylla_v2

rm -rf agent/client agent/models
swagger generate client -A agent -T agent/templates -f agent.json -t ./agent