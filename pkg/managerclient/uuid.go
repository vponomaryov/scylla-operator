// Copyright (C) 2017 ScyllaDB

package managerclient

import (
	"net/url"
	"path"

	"github.com/scylladb/scylla-mgmt-commons/uuid"
)

// uuidFromLocation returns a UUID extracted from Location header.
func uuidFromLocation(location string) (uuid.UUID, error) {
	l, err := url.Parse(location)
	if err != nil {
		return uuid.Nil, err
	}
	_, id := path.Split(l.Path)

	return uuid.Parse(id)
}
