package test

import (
	"context"
	"fmt"
	"net/url"
	"testing"
	"time"

	pds "github.com/portworx/pds-api-go-client/pds/v1alpha1"
	api "github.com/portworx/pds-functional-test/pkg/api"
	. "github.com/portworx/pds-functional-test/pkg/common"
	logger "github.com/portworx/pds-functional-test/pkg/logger"
	"github.com/stretchr/testify/suite"
)

var (
	accountId, tenantId, projectId string

	log                      = logger.Log
	accountName              = "Portworx"
	S3BackupTarget           = "agaurav-aws-target"
	S3CompatibleBackupTarget = "agaurav-s3-compatible-target"
	BLOBBackuptarget         = "agaurav-azure-target"
	serviceType              = "LoadBalancer"
	pdsNamespaces            []string
	//serviceTypeOnprem        = "ClusterIP"
	supportedDataServices = map[string]string{"cas": "Cassandra", "zk": "Zookeeper", "rmq": "Rabbitmq", "pg": "Postgresql"}

	// supportedDataServices = map[string]string{"cas": "Cassandra", "zk": "Zookeeper", "kf": "Kafka", "rmq": "Rabbitmq", "pg": "Postgresql"}
	backupSupportedDataService = map[string]string{"cas": "Cassandra", "pg": "Postgresql"}
	//futureSupportedDataService = map[string]string{"mdb": "Mongodb", "red": "Redis", "con": "Consul", "cbs": "Couchbase", "dse": "DatastaxEnterprise", "ess": "Elasticsearch"}
)

type PDSTestSuite struct {
	suite.Suite
	TargetCluster   *TargetCluster
	TestEnvironment *Environment
	ctx             context.Context
	apiClient       *pds.APIClient
	components      *api.Components
	env             Environment
}

const (
	duration  = 600
	sleepTime = 10

	defaultNumPods = 3
	dnsZone        = "portworx.pds-dns-dev.io"

	// FIX-ME
	// Create the template manually for all the data serices with below name (i.e Default)
	storageTemplateName   = "Default"
	resourceTemplateName  = "Default"
	appConfigTemplateName = "Default"
	deploymentName        = "automation"
	templateName          = "Default"
)

var (
	dataServiceDefaultResourceTemplateIdMap = make(map[string]string)
	dataServiceNameIdMap                    = make(map[string]string)
	dataServiceNameVersionMap               = make(map[string][]string)
	dataServiceNameImagesMap                = make(map[string]string)
	dataServiceNameDefaultAppConfigMap      = make(map[string]string)
	deployementIdNameMap                    = make(map[string]string)
	namespaceNameIdMap                      = make(map[string]string)
	backuppolicyIdNameMap                   = make(map[string]string)
	backupTargetNameIdMap                   = make(map[string]string)
	deployementIdnameWithSchBkpMap          = make(map[string]string)
	deployementIdnameWithAdhocBkpMap        = make(map[string]string)
)

func (suite *PDSTestSuite) SetupSuite() {

	// Perform basic setup with sanity checks.
	log.Info("Check for environmental variable.")
	suite.env = MustHaveEnvVariables()

	log.Info("Get target cluster.")
	suite.TargetCluster = NewTargetCluster(suite.env.TARGET_KUBECONFIG)

	endpointUrl, err := url.Parse(suite.env.CONTROL_PLANE_URL)
	if err != nil {
		log.Panicf("Unable to access the URL: %s", suite.env.CONTROL_PLANE_URL)
	}
	apiConf := pds.NewConfiguration()
	apiConf.Host = endpointUrl.Host
	apiConf.Scheme = endpointUrl.Scheme

	// Use Configuration or context with WithValue (above)
	suite.ctx = context.WithValue(context.Background(), pds.ContextAPIKeys, map[string]pds.APIKey{"ApiKeyAuth": {Key: GetBearerToken(true), Prefix: "Bearer"}})
	suite.apiClient = pds.NewAPIClient(apiConf)
	suite.components = api.NewComponents(suite.ctx, suite.apiClient)

}

func (s *PDSTestSuite) BeforeTest(suiteName, testName string) {
	acc := s.components.Account
	accounts, _ := acc.GetAccountsList()

	for i := 0; i < len(accounts); i++ {
		log.Infof("Account Name: %v", accounts[i].GetName())
		if accounts[i].GetName() == accountName {
			accountId = accounts[i].GetId()
		}
	}
	log.Infof("Account Detail- Name: %s, UUID: %s ", accountName, accountId)
	tnts := s.components.Tenant
	tenants, _ := tnts.GetTenantsList(accountId)
	tenantId = tenants[0].GetId()
	tenantName := tenants[0].GetName()
	log.Infof("Tenant Details- Name: %s, UUID: %s ", tenantName, tenantId)
	projcts := s.components.Project
	projects, _ := projcts.GetprojectsList(tenantId)
	projectId = projects[0].GetId()
	projectName := projects[0].GetName()
	log.Infof("Project Details- Name: %s, UUID: %s ", projectName, projectId)

	log.Info("Get helm version")
	version, _ := s.components.ApiVersion.GetHelmChartVersion()
	log.Infof("Helm chart Version : %s ", version)

	clusterId, err := GetClusterId(s.env.TARGET_KUBECONFIG)
	if err != nil {
		log.Panicf("Unable to fetch the cluster Id")
	}

	log.Infof("Register cluster %s to control plane %v ", clusterId, s.env.CONTROL_PLANE_URL)
	err = s.TargetCluster.RegisterToControlPlane(s.env.CONTROL_PLANE_URL, version, GetBearerToken(true), tenantId)
	if err != nil {
		log.Panicf("Unable to register the target cluster to control plane %v", s.env.CONTROL_PLANE_URL)
	}
	log.Info("Creating namespaces for data service deployment")
	pdsNamespaces = []string{"automation-1", "automation-2", "automation-3"}
	for _, ns := range pdsNamespaces {
		log.Infof("Namespace name - %s", ns)
		s.TargetCluster.CreatePDSNamespace(ns)
	}
}

func (suite *PDSTestSuite) AfterTest(suiteName, testName string) {
	log.Warn("Cleaning all the deployment created as part of this test run")
	log.Info("Sleep for 10 minutes.")
	time.Sleep(10 * time.Minute)
	for id := range deployementIdNameMap {
		log.Infof("Deleting the deployment: %v", id)
		suite.components.DataServiceDeployment.DeleteDeployment(id)
		time.Sleep(sleepTime)
	}
	log.Info("Sleep for 10 minutes.")
	time.Sleep(10 * time.Minute)
	for id := range deployementIdnameWithSchBkpMap {
		backups, _ := suite.components.Backup.ListBackup(id)
		for _, backup := range backups {
			backupId := backup.GetId()
			log.Infof("Delete back up having Id - %v", backupId)
			response, err := suite.components.Backup.DeleteBackup(backupId)
			// Success is indicated with 2xx status codes:
			statusOK := response.StatusCode >= 200 && response.StatusCode < 300
			if !statusOK {
				fmt.Println("Non-OK HTTP status:", response.StatusCode)
				// You may read / inspect response body
				log.Error(err)
			}
			time.Sleep(10 * time.Second)
		}
		log.Infof("Deleting the deployment: %v", id)
		suite.components.DataServiceDeployment.DeleteDeployment(id)
		time.Sleep(15 * time.Second)
	}
	log.Info("Sleep for 10 minutes.")
	time.Sleep(10 * time.Minute)
	for id := range deployementIdnameWithAdhocBkpMap {
		backups, _ := suite.components.Backup.ListBackup(id)
		for _, backup := range backups {
			backupId := backup.GetId()
			log.Infof("Delete back up having Id - %v", backupId)
			response, err := suite.components.Backup.DeleteBackup(backupId)
			// Success is indicated with 2xx status codes:
			statusOK := response.StatusCode >= 200 && response.StatusCode < 300
			if !statusOK {
				fmt.Println("Non-OK HTTP status:", response.StatusCode)
				// You may read / inspect response body
				log.Error(err)
			}
			time.Sleep(10 * time.Second)
		}
		log.Infof("Deleting the deployment: %v", id)
		suite.components.DataServiceDeployment.DeleteDeployment(id)
		time.Sleep(25 * time.Second)
	}
	if suite.T().Failed() {
		log.Errorf(fmt.Sprintf("Failed test %s:", testName))
	}
}

func TestPDSSuite(t *testing.T) {
	suite.Run(t, new(PDSTestSuite))
}
