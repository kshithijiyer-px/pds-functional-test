package test

import (
	"math/rand"
	"reflect"
	"time"

	. "github.com/portworx/pds-functional-test/pkg/common"
)

func (suite *PDSTestSuite) TestBasicBackupOperation() {

	var (
		deploymentTargetId, storageTemplateId string
		deploymentTargetComponent             = suite.components.DeploymentTarget
		nsComponent                           = suite.components.Namespace
		storagetemplateComponent              = suite.components.StorageSettingsTemplate
		resourceTemplateComponent             = suite.components.ResourceSettingsTemplate
		dataServiceComponent                  = suite.components.DataService
		versionComponent                      = suite.components.Version
		imageComponent                        = suite.components.Image
		appConfigComponent                    = suite.components.AppConfigTemplate
		dataServiceDeploymentComponent        = suite.components.DataServiceDeployment
	)

	clusterId, err := GetClusterId(suite.env.TARGET_KUBECONFIG)
	if err != nil {
		log.Panicf("Unable to fetch the cluster Id")
	}

	log.Info("Get the Target cluster details")
	targetClusters, _ := deploymentTargetComponent.ListDeploymentTargetsBelongsToTenant(tenantId)
	for i := 0; i < len(targetClusters); i++ {
		if targetClusters[i].GetClusterId() == clusterId {
			deploymentTargetId = targetClusters[i].GetId()
			log.Infof("Cluster ID: %v, Name: %v,Status: %v", targetClusters[i].GetClusterId(), targetClusters[i].GetName(), targetClusters[i].GetStatus())
		}
	}

	log.Infof("Get the available namespaces in the Cluster having Id: %v", clusterId)
	namespaces, _ := nsComponent.ListNamespaces(deploymentTargetId)
	for i := 0; i < len(namespaces); i++ {
		if namespaces[i].GetStatus() == "available" {
			namespaceNameIdMap[namespaces[i].GetName()] = namespaces[i].GetId()
			log.Infof("Available namespace - Name: %v , Id: %v , Status: %v", namespaces[i].GetName(), namespaces[i].GetId(), namespaces[i].GetStatus())
		}
	}
	log.Infof("Get the storage template")
	storageTemplates, _ := storagetemplateComponent.ListTemplates(tenantId)
	for i := 0; i < len(storageTemplates); i++ {
		if storageTemplates[i].GetName() == storageTemplateName {
			log.Infof("Storage template details -----> Name %v,Repl %v , Fg %v , Fs %v",
				storageTemplates[i].GetName(),
				storageTemplates[i].GetRepl(),
				storageTemplates[i].GetFg(),
				storageTemplates[i].GetFs())
			storageTemplateId = storageTemplates[i].GetId()
			log.Infof("Storage Id: %v", storageTemplateId)
		}
	}

	log.Infof("Get the resource template for each data services")
	resourceTemplates, _ := resourceTemplateComponent.ListTemplates(tenantId)
	for i := 0; i < len(resourceTemplates); i++ {
		if resourceTemplates[i].GetName() == resourceTemplateName {
			dataService, _ := dataServiceComponent.GetDataService(resourceTemplates[i].GetDataServiceId())
			log.Infof("Data service name: %v", dataService.GetName())
			log.Infof("Resource template details ---> Name %v, Id : %v ,DataServiceId %v , StorageReq %v , Memoryrequest %v",
				resourceTemplates[i].GetName(),
				resourceTemplates[i].GetId(),
				resourceTemplates[i].GetDataServiceId(),
				resourceTemplates[i].GetStorageRequest(),
				resourceTemplates[i].GetMemoryRequest())
			dataServiceDefaultResourceTemplateIdMap[dataService.GetName()] =
				resourceTemplates[i].GetId()
			dataServiceNameIdMap[dataService.GetName()] = dataService.GetId()
		}
	}

	log.Infof("Get the Versions.")
	for key := range dataServiceNameIdMap {
		versions, _ := versionComponent.ListDataServiceVersions(dataServiceNameIdMap[key])
		for i := 0; i < len(versions); i++ {
			dataServiceNameVersionMap[key] = append(dataServiceNameVersionMap[key], versions[i].GetId())
		}
	}

	log.Infof("Get the Versions.")
	for key := range dataServiceNameVersionMap {
		images, _ := imageComponent.ListImages(dataServiceNameVersionMap[key][0])
		for i := 0; i < len(images); i++ {
			dataServiceNameImagesMap[key] = images[i].GetId()
		}
	}

	log.Infof("Get the Application configuration template")
	appConfigs, _ := appConfigComponent.ListTemplates(tenantId)
	for i := 0; i < len(appConfigs); i++ {
		if appConfigs[i].GetName() == appConfigTemplateName {
			for key := range dataServiceNameIdMap {
				if dataServiceNameIdMap[key] == appConfigs[i].GetDataServiceId() {
					dataServiceNameDefaultAppConfigMap[key] = appConfigs[i].GetId()
				}
			}
		}

	}

	log.Infof("Get the backup details")
	log.Infof("Get the backup policy")
	backupPolicies, _ := suite.components.BackupPolicy.ListBackupPolicy(tenantId)
	for i := 0; i < len(backupPolicies); i++ {
		backuppolicyIdNameMap[backupPolicies[i].GetId()] = backupPolicies[i].GetName()
	}
	backupPolicyId := MapRandomKeyGet(backuppolicyIdNameMap).(string)

	log.Infof("Get the backup target")
	backupTargets, _ := suite.components.BackupTarget.ListBackupTarget(tenantId)
	for i := 0; i < len(backupTargets); i++ {
		name := backupTargets[i].GetName()
		if name == S3BackupTarget || name == S3CompatibleBackupTarget || name == BLOBBackuptarget {
			backupTargetNameIdMap[backupTargets[i].GetName()] = backupTargets[i].GetId()
		}

	}
	duration := 600
	deploymentNameSch := "agaurav-schbkp-1"
	deploymentNameAdhoc := "agaurav-adhocBkp-1"

	log.Info("Create dataservice with scheduled backup enabled")
	for i := range backupSupportedDataService {
		log.Infof("Key: %v, Value %v", backupSupportedDataService[i], dataServiceNameDefaultAppConfigMap[backupSupportedDataService[i]])
		n := rand.Int() % len(pdsNamespaces)
		namespace := pdsNamespaces[n]
		namespaceId := namespaceNameIdMap[namespace]
		log.Infof("Created %v deployment  in the namespace %v with no scheduled back up.", backupSupportedDataService[i], namespace)
		deployment, _ :=
			suite.components.DataServiceDeployment.CreateDeploymentWithScehduleBackup(projectId,
				deploymentTargetId,
				dnsZone,
				deploymentNameSch,
				namespaceId,
				dataServiceNameDefaultAppConfigMap[backupSupportedDataService[i]],
				dataServiceNameImagesMap[backupSupportedDataService[i]],
				3,
				serviceType,
				dataServiceDefaultResourceTemplateIdMap[backupSupportedDataService[i]],
				storageTemplateId,
				backupPolicyId,
				backupTargetNameIdMap[S3BackupTarget],
			)
		deployementIdnameWithSchBkpMap[deployment.GetId()] = deployment.GetName()

		status, _ := suite.components.DataServiceDeployment.GetDeploymentSatus(deployment.GetId())
		sleeptime := 0
		for status.GetHealth() != "Healthy" && sleeptime < duration {
			time.Sleep(15 * time.Second)
			sleeptime += 15
			status, _ = suite.components.DataServiceDeployment.GetDeploymentSatus(deployment.GetId())
			log.Infof("Health status -  %v", status.GetHealth())
		}
		log.Infof("Deployment state - %v,Health status -  %v,Replicas - %v, Ready replicas - %v", deployment.GetState(), status.GetHealth(), status.GetReplicas(), status.GetReadyReplicas())

	}
	log.Info("Create dataservice to perform adhoc backups.")
	for i := range backupSupportedDataService {
		log.Infof("Key: %v, Value %v", backupSupportedDataService[i], dataServiceNameDefaultAppConfigMap[backupSupportedDataService[i]])
		n := rand.Int() % len(pdsNamespaces)
		namespace := pdsNamespaces[n]
		namespaceId := namespaceNameIdMap[namespace]
		log.Infof("Created %v deployment  in the namespace %v with no scheduled back up.", backupSupportedDataService[i], namespace)

		deployment, _ :=
			dataServiceDeploymentComponent.CreateDeployment(projectId,
				deploymentTargetId,
				dnsZone,
				deploymentNameAdhoc,
				namespaceId,
				dataServiceNameDefaultAppConfigMap[backupSupportedDataService[i]],
				dataServiceNameImagesMap[backupSupportedDataService[i]],
				defaultNumPods,
				serviceType,
				dataServiceDefaultResourceTemplateIdMap[backupSupportedDataService[i]],
				storageTemplateId,
			)
		deployementIdnameWithAdhocBkpMap[deployment.GetId()] = deployment.GetName()
		status, _ := suite.components.DataServiceDeployment.GetDeploymentSatus(deployment.GetId())
		sleeptime := 0
		for status.GetHealth() != "Healthy" && sleeptime < duration {
			time.Sleep(10 * time.Second)
			sleeptime += 10
			status, _ = suite.components.DataServiceDeployment.GetDeploymentSatus(deployment.GetId())
			log.Infof("Health status -  %v", status.GetHealth())
		}
		log.Infof("Deployment state - %v,Health status -  %v,Replicas - %v, Ready replicas - %v", deployment.GetState(), status.GetHealth(), status.GetReplicas(), status.GetReadyReplicas())

	}

	log.Infof("Sleeping for %v minutes", 5*time.Minute)
	time.Sleep(5 * time.Minute)
	log.Info("Take Adhoc backups for dataservices")
	for id := range deployementIdnameWithAdhocBkpMap {
		for backupTarget := range backupTargetNameIdMap {
			log.Infof("Creating ADHOC backup for deployment -  %v to backup target - %v", deployementIdnameWithAdhocBkpMap[id], backupTarget)
			backup, _ := suite.components.Backup.CreateBackup(id, backupTargetNameIdMap[backupTarget], 30, true)
			log.Info(backup.GetState())
			jobs, _ := suite.components.BackupJob.ListBackupJobs(backup.GetId())
			sleeptime := 0
			for _, job := range jobs {
				for job.GetStatus() == "Succeeded" && sleeptime < duration {
					time.Sleep(15 * time.Second)
					sleeptime += 15
					log.Infof("Backup Job Name - %v,  Job status %v", job.GetName(), job.GetStatus())
				}
			}
		}

	}

	log.Infof("Sleeping for %v minutes", 5*time.Minute)
	time.Sleep(5 * time.Minute)
	log.Info("Take Adhoc backups for dataservices created for scheduled backup")
	for id := range deployementIdnameWithSchBkpMap {
		for backupTarget := range backupTargetNameIdMap {
			log.Infof("Creating ADHOC backup for deployment -  %v to backup target - %v", deployementIdnameWithSchBkpMap[id], backupTarget)
			backup, _ := suite.components.Backup.CreateBackup(id, backupTargetNameIdMap[backupTarget], 30, true)
			log.Info(backup.GetState())
			jobs, _ := suite.components.BackupJob.ListBackupJobs(backup.GetId())
			sleeptime := 0
			for _, job := range jobs {
				for job.GetStatus() == "Succeeded" && sleeptime < duration {
					time.Sleep(15 * time.Second)
					sleeptime += 15
					log.Infof("Backup Job Name - %v,  Job status %v", job.GetName(), job.GetStatus())
				}
			}
		}

	}

}

func MapRandomKeyGet(mapI interface{}) interface{} {
	keys := reflect.ValueOf(mapI).MapKeys()

	return keys[rand.Intn(len(keys))].Interface()
}
