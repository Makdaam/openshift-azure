package fakerp

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/sirupsen/logrus"

	"github.com/openshift/openshift-azure/pkg/cluster"
	"github.com/openshift/openshift-azure/pkg/util/cloudprovider"
	"github.com/openshift/openshift-azure/pkg/util/configblob"
)

// handleGetControlPlanePods handles admin requests for the list of control plane pods
func (s *Server) handleGetControlPlanePods(w http.ResponseWriter, req *http.Request) {
	cs := s.read()
	if cs == nil {
		s.internalError(w, "Failed to read the internal config")
		return
	}
	ctx, err := enrichContext(context.Background())
	if err != nil {
		s.internalError(w, fmt.Sprintf("Failed to enrich context: %v", err))
		return
	}
	pods, err := s.plugin.GetControlPlanePods(ctx, cs)
	if err != nil {
		s.internalError(w, fmt.Sprintf("Failed to fetch control plane pods: %v", err))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(pods)
	s.log.Info("fetched control plane pods")
}

// handleRestore handles admin requests to restore an etcd cluster from a backup
func (s *Server) handleRestore(w http.ResponseWriter, req *http.Request) {
	cs := s.read()
	if cs == nil {
		s.internalError(w, "Failed to read the internal config")
		return
	}

	cpc := &cloudprovider.Config{
		TenantID:        cs.Properties.AzProfile.TenantID,
		SubscriptionID:  cs.Properties.AzProfile.SubscriptionID,
		AadClientID:     cs.Properties.ServicePrincipalProfile.ClientID,
		AadClientSecret: cs.Properties.ServicePrincipalProfile.Secret,
		ResourceGroup:   cs.Properties.AzProfile.ResourceGroup,
	}

	blobName, err := readBlobName(req)
	if err != nil {
		s.badRequest(w, fmt.Sprintf("Cannot read blob name from request: %v", err))
		return
	}
	if len(blobName) == 0 {
		s.badRequest(w, "Blob name to restore from was not provided")
		return
	}

	ctx := context.Background()
	bsc, err := configblob.GetService(ctx, cpc)
	if err != nil {
		s.internalError(w, fmt.Sprintf("Failed to configure blob client: %v", err))
		return
	}
	etcdContainer := bsc.GetContainerReference(cluster.EtcdBackupContainerName)

	blob := etcdContainer.GetBlobReference(blobName)
	exists, err := blob.Exists()
	if err != nil {
		s.internalError(w, fmt.Sprintf("Cannot get blob ref for %s: %v", blobName, err))
		return
	}
	if !exists {
		resp, err := etcdContainer.ListBlobs(storage.ListBlobsParameters{})
		if err == nil {
			s.log.Infof("available blobs:")
			for _, blob := range resp.Blobs {
				s.log.Infof("  %s", blob.Name)
			}
		}
		s.badRequest(w, fmt.Sprintf("Blob %s does not exist", blobName))
		return
	}

	ctx, err = enrichContext(context.Background())
	if err != nil {
		s.internalError(w, fmt.Sprintf("Failed to enrich context: %v", err))
		return
	}
	deployer := GetDeployer(s.log, cs, s.pluginConfig)
	if err := s.plugin.RecoverEtcdCluster(ctx, cs, deployer, blobName); err != nil {
		s.internalError(w, fmt.Sprintf("Failed to recover cluster: %v", err))
		return
	}

	s.log.Info("recovered cluster")
}

// handleRotateSecrets handles admin requests for the rotation of cluster secrets
func (s *Server) handleRotateSecrets(w http.ResponseWriter, req *http.Request) {
	cs := s.read()
	if cs == nil {
		s.internalError(w, "Failed to read the internal config")
		return
	}
	ctx, err := enrichContext(context.Background())
	if err != nil {
		s.internalError(w, fmt.Sprintf("Failed to enrich context: %v", err))
		return
	}
	deployer := GetDeployer(s.log, cs, s.pluginConfig)
	if err := s.plugin.RotateClusterSecrets(ctx, cs, deployer, s.pluginTemplate); err != nil {
		s.internalError(w, fmt.Sprintf("Failed to rotate cluster secrets: %v", err))
		return
	}
	s.log.Info("rotated cluster secrets")
}

// handleGetLogLevel handles admin requests to get the plugin log level
func (s *Server) handleGetLogLevel(w http.ResponseWriter, req *http.Request) {
	cs := s.read()
	if cs == nil {
		s.internalError(w, "Failed to read the internal config")
		return
	}
	ctx, err := enrichContext(context.Background())
	if err != nil {
		s.internalError(w, fmt.Sprintf("Failed to enrich context: %v", err))
		return
	}
	level, err := s.plugin.GetLogLevel(ctx, cs)
	if err != nil {
		s.internalError(w, fmt.Sprintf("Failed to fetch plugin log level: %v", err))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(level)
	s.log.Info("fetched log level")
}

// handleChangeLogLevel handles admin requests to change the plugin log level
func (s *Server) handleChangeLogLevel(w http.ResponseWriter, req *http.Request) {
	type jsonParams struct {
		Level string `json:"level"`
	}
	var parameters jsonParams
	cs := s.read()
	if cs == nil {
		s.internalError(w, "Failed to read the internal config")
		return
	}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.internalError(w, fmt.Sprintf("failed to read request body: %v", err))
		return
	}
	err = json.Unmarshal(data, &parameters)
	if err != nil {
		s.internalError(w, fmt.Sprintf("failed to unmarshall JSON: %v", err))
		return
	}
	s.log.Info(parameters.Level)
	level, err := logrus.ParseLevel(parameters.Level)
	if err != nil {
		s.internalError(w, fmt.Sprintf("failed to parse LogLevel: %v", err))
		return
	}
	ctx, err := enrichContext(context.Background())
	if err != nil {
		s.internalError(w, fmt.Sprintf("Failed to enrich context: %v", err))
		return
	}
	s.plugin.ChangeLogLevel(ctx, cs, level)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	s.log.Info("changed log level")
}
