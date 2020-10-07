// Copyright (C) 2017 ScyllaDB

package managerclient

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"
	"sync"

	api "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/scylladb/scylla-mgmt-commons/managerclient/client/operations"
	"github.com/scylladb/scylla-mgmt-commons/managerclient/models"
	"github.com/scylladb/scylla-mgmt-commons/uuid"
)

var disableOpenAPIDebugOnce sync.Once

// Client provides means to interact with Mermaid.
type Client struct {
	operations *operations.Client
}

// DefaultTLSConfig specifies default TLS configuration used when creating a new
// client.
var DefaultTLSConfig = func() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
	}
}

func NewClient(rawURL string, transport http.RoundTripper) (Client, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return Client{}, err
	}

	disableOpenAPIDebugOnce.Do(func() {
		middleware.Debug = false
	})

	if transport == nil {
		transport = &http.Transport{
			TLSClientConfig: DefaultTLSConfig(),
		}
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	r := api.NewWithClient(u.Host, u.Path, []string{u.Scheme}, httpClient)
	return Client{operations: operations.New(r, strfmt.Default)}, nil
}

// CreateCluster creates a new cluster.
func (c Client) CreateCluster(ctx context.Context, cluster *models.Cluster) (string, error) {
	resp, err := c.operations.PostClusters(&operations.PostClustersParams{
		Context: ctx,
		Cluster: cluster,
	})
	if err != nil {
		return "", err
	}

	clusterID, err := uuidFromLocation(resp.Location)
	if err != nil {
		return "", errors.Wrap(err, "cannot parse response")
	}

	return clusterID.String(), nil
}

// GetCluster returns a cluster for a given ID.
func (c Client) GetCluster(ctx context.Context, clusterID string) (*models.Cluster, error) {
	resp, err := c.operations.GetClusterClusterID(&operations.GetClusterClusterIDParams{
		Context:   ctx,
		ClusterID: clusterID,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

// UpdateCluster updates cluster.
func (c Client) UpdateCluster(ctx context.Context, cluster *models.Cluster) error {
	_, err := c.operations.PutClusterClusterID(&operations.PutClusterClusterIDParams{ // nolint: errcheck
		Context:   ctx,
		ClusterID: cluster.ID,
		Cluster:   cluster,
	})
	return err
}

// DeleteCluster removes cluster.
func (c Client) DeleteCluster(ctx context.Context, clusterID string) error {
	_, err := c.operations.DeleteClusterClusterID(&operations.DeleteClusterClusterIDParams{ // nolint: errcheck
		Context:   ctx,
		ClusterID: clusterID,
	})
	return err
}

// ListClusters returns clusters.
func (c Client) ListClusters(ctx context.Context) ([]*models.Cluster, error) {
	resp, err := c.operations.GetClusters(&operations.GetClustersParams{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

// CreateTask creates a new task.
func (c *Client) CreateTask(ctx context.Context, clusterID string, t *models.Task) (uuid.UUID, error) {
	params := &operations.PostClusterClusterIDTasksParams{
		Context:   ctx,
		ClusterID: clusterID,
		TaskFields: &models.TaskUpdate{
			Type:       t.Type,
			Enabled:    t.Enabled,
			Name:       t.Name,
			Schedule:   t.Schedule,
			Tags:       t.Tags,
			Properties: t.Properties,
		},
	}
	resp, err := c.operations.PostClusterClusterIDTasks(params)
	if err != nil {
		return uuid.Nil, err
	}

	taskID, err := uuidFromLocation(resp.Location)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "cannot parse response")
	}

	return taskID, nil
}

// DeleteTask stops executing a task.
func (c *Client) DeleteTask(ctx context.Context, clusterID, taskType string, taskID uuid.UUID) error {
	_, err := c.operations.DeleteClusterClusterIDTaskTaskTypeTaskID(&operations.DeleteClusterClusterIDTaskTaskTypeTaskIDParams{ // nolint: errcheck
		Context:   ctx,
		ClusterID: clusterID,
		TaskType:  taskType,
		TaskID:    taskID.String(),
	})

	return err
}

// UpdateTask updates an existing task unit.
func (c *Client) UpdateTask(ctx context.Context, clusterID string, t *models.Task) error {
	_, err := c.operations.PutClusterClusterIDTaskTaskTypeTaskID(&operations.PutClusterClusterIDTaskTaskTypeTaskIDParams{ // nolint: errcheck
		Context:   ctx,
		ClusterID: clusterID,
		TaskType:  t.Type,
		TaskID:    t.ID,
		TaskFields: &models.TaskUpdate{
			Enabled:    t.Enabled,
			Name:       t.Name,
			Schedule:   t.Schedule,
			Tags:       t.Tags,
			Properties: t.Properties,
		},
	})
	return err
}

// ListTasks returns scheduled tasks within a clusterID, optionally filtered by task type tp.
func (c *Client) ListTasks(ctx context.Context, clusterID, taskType string, all bool, status string) ([]*models.ExtendedTask, error) {
	resp, err := c.operations.GetClusterClusterIDTasks(&operations.GetClusterClusterIDTasksParams{
		Context:   ctx,
		ClusterID: clusterID,
		Type:      &taskType,
		All:       &all,
		Status:    &status,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

// Version returns server version.
func (c Client) Version(ctx context.Context) (*models.Version, error) {
	resp, err := c.operations.GetVersion(&operations.GetVersionParams{
		Context: ctx,
	})
	if err != nil {
		return &models.Version{}, err
	}

	return resp.Payload, nil
}
