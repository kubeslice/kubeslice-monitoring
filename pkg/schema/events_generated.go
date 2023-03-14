package schema

// Autogenerated file. DO NOT MODIFY DIRECTLY!
/*
 *  Copyright (c) 2022 Avesha, Inc. All rights reserved.
 *
 *  SPDX-License-Identifier: Apache-2.0
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

var eventsMap = map[string]*EventSchema{
	"LicenseSecretNotFound": {
		Name:                "LicenseSecretNotFound",
		Reason:              "LicenseNotFound",
		Action:              "LicenseValidation",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Secret with license config not found. Please ensure to create license secret if this is airgapped cluster.",
	},
	"MachineFileNotFound": {
		Name:                "MachineFileNotFound",
		Reason:              "MachineFileNotFound",
		Action:              "FetchMachineFileFromSecret",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Machine File not present in license secret. Please contact kubeslice support team.",
	},
	"MachineFileInvalid": {
		Name:                "MachineFileInvalid",
		Reason:              "MachineFileInvalid",
		Action:              "VerifyMachineFile",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Machine File is invalid in license secret. Please contact kubeslice support team.",
	},
	"LicenseKeyInvalid": {
		Name:                "LicenseKeyInvalid",
		Reason:              "LicenseKeyInvalid",
		Action:              "DecryptMachineFile",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "License Key is invalid in license secret. Please contact kubeslice support team.",
	},
	"LicenseExpired": {
		Name:                "LicenseExpired",
		Reason:              "LicenseExpired",
		Action:              "ValidateLicense",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "License Expired, please contact kubeslice support team to get it renewed!",
	},
	"LicenseExpiredGracePeriodOn": {
		Name:                "LicenseExpiredGracePeriodOn",
		Reason:              "LicenseExpiredGracePeriodOn",
		Action:              "ValidateLicense",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "License Expired, grace period is on, please contact kubeslice support team to get it renewed!",
	},
	"MachineFingerPrintErr": {
		Name:                "MachineFingerPrintErr",
		Reason:              "MachineFingerPrintErr",
		Action:              "GetMachineFingerPrint",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Unable to fetch MachineFingerPrint. Please look at the kubeslice-controller logs for more details.",
	},
	"GotMachineFingerPrint": {
		Name:                "GotMachineFingerPrint",
		Reason:              "GotMachineFingerPrint",
		Action:              "GetMachineFingerPrint",
		Type:                EventTypeNormal,
		ReportingController: "",
		Message:             "Got MachineFingerPrint successfully",
	},
	"ConfigMapErr": {
		Name:                "ConfigMapErr",
		Reason:              "ConfigMapErr",
		Action:              "GetConfigMap",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Unable to validate license keys from configmap. Please contact kubeslice support team",
	},
	"GotConfigMap": {
		Name:                "GotConfigMap",
		Reason:              "GotConfigMap",
		Action:              "GetConfigMap",
		Type:                EventTypeNormal,
		ReportingController: "",
		Message:             "Got ConfigMap and kubeslice-license public keys successfully.",
	},
	"LicenseProxyUnreachable": {
		Name:                "LicenseProxyUnreachable",
		Reason:              "LicenseProxyUnreachable",
		Action:              "LicenseProxy",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Failed to get response from license proxy for automatic license creation. Please contact Avesha to obtain licernse manually.",
	},
	"LicenseDeployError": {
		Name:                "LicenseDeployError",
		Reason:              "LicenseDeployError",
		Action:              "LicenseDeploy",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Failed to deploy license secret.",
	},
	"LicenseDeploySuccess": {
		Name:                "LicenseDeploySuccess",
		Reason:              "LicenseDeploySuccess",
		Action:              "LicenseDeploy",
		Type:                EventTypeNormal,
		ReportingController: "",
		Message:             "Successfully fetched license data & deployed license secret.",
	},
	"ClusterMetadataCollectionFailed": {
		Name:                "ClusterMetadataCollectionFailed",
		Reason:              "ClusterMetadataCollectionFailed",
		Action:              "CollectClusterMetadata",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Failed to collect cluster metadata for cluster creation.",
	},
	"ClusterMetadataCollectionSuccess": {
		Name:                "ClusterMetadataCollectionSuccess",
		Reason:              "ClusterMetadataCollectionSuccess",
		Action:              "CollectClusterMetadata",
		Type:                EventTypeNormal,
		ReportingController: "",
		Message:             "Successfully collected cluster metadata for license creation.",
	},
	"LicenseDataFetchError": {
		Name:                "LicenseDataFetchError",
		Reason:              "LicenseDataFetchError",
		Action:              "LicenseDataFetch",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Failed to fetch license data from license proxy.",
	},
	"LicenseDataFetchSuccess": {
		Name:                "LicenseDataFetchSuccess",
		Reason:              "LicenseDataFetchSuccess",
		Action:              "LicenseDataFetch",
		Type:                EventTypeNormal,
		ReportingController: "",
		Message:             "Successfully fetched license data from license proxy.",
	},
	"LicenseSecretCreationFailed": {
		Name:                "LicenseSecretCreationFailed",
		Reason:              "LicenseSecretCreationFailed",
		Action:              "LicenseSecretCreation",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Failed to create license secret.",
	},
	"LicenseSecretCreationSuccess": {
		Name:                "LicenseSecretCreationSuccess",
		Reason:              "LicenseSecretCreationSuccess",
		Action:              "LicenseSecretCreation",
		Type:                EventTypeNormal,
		ReportingController: "",
		Message:             "Successfully created license secret.",
	},
	"ProjectDeleted": {
		Name:                "ProjectDeleted",
		Reason:              "ProjectDeleted",
		Action:              "DeleteProject",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Project got deleted.",
	},
	"ProjectDeletionFailed": {
		Name:                "ProjectDeletionFailed",
		Reason:              "ProjectDeletionFailed",
		Action:              "DeleteProject",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Project deletion failed.",
	},
	"ClusterDeleted": {
		Name:                "ClusterDeleted",
		Reason:              "ClusterDeleted",
		Action:              "DeleteCluster",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Cluster got deleted.",
	},
	"ClusterDeletionFailed": {
		Name:                "ClusterDeletionFailed",
		Reason:              "ClusterDeletionFailed",
		Action:              "DeleteCluster",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Cluster deletion failed.",
	},
	"SliceConfigDeleted": {
		Name:                "SliceConfigDeleted",
		Reason:              "SliceConfigDeleted",
		Action:              "DeleteSliceConfig",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Slice config got deleted.",
	},
	"SliceConfigDeletionFailed": {
		Name:                "SliceConfigDeletionFailed",
		Reason:              "SliceConfigDeletionFailed",
		Action:              "DeleteSliceConfig",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Slice config deletion failed.",
	},
	"ServiceExportConfigDeleted": {
		Name:                "ServiceExportConfigDeleted",
		Reason:              "ServiceExportConfigDeleted",
		Action:              "DeleteServiceExportConfig",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Service export config got deleted.",
	},
	"ServiceExportConfigDeletionFailed": {
		Name:                "ServiceExportConfigDeletionFailed",
		Reason:              "ServiceExportConfigDeletionFailed",
		Action:              "DeleteServiceExportConfig",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Service export config deletion failed.",
	},
	"SliceQoSConfigDeleted": {
		Name:                "SliceQoSConfigDeleted",
		Reason:              "SliceQoSConfigDeleted",
		Action:              "DeleteSliceQoSConfig",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Slice QoS config got deleted.",
	},
	"SliceQoSConfigDeletionFailed": {
		Name:                "SliceQoSConfigDeletionFailed",
		Reason:              "SliceQoSConfigDeletionFailed",
		Action:              "DeleteSliceQoSConfig",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Slice QoS config deletion failed.",
	},
	"SecretDeleted": {
		Name:                "SecretDeleted",
		Reason:              "SecretDeleted",
		Action:              "DeleteSecret",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Secret got deleted.",
	},
	"SecretDeletionFailed": {
		Name:                "SecretDeletionFailed",
		Reason:              "SecretDeletionFailed",
		Action:              "DeleteSecret",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Secret deletion failed.",
	},
	"NamespaceCreated": {
		Name:                "NamespaceCreated",
		Reason:              "NamespaceCreated",
		Action:              "CreateNamespace",
		Type:                EventTypeNormal,
		ReportingController: "",
		Message:             "Namespace got created.",
	},
	"NamespaceCreationFailed": {
		Name:                "NamespaceCreationFailed",
		Reason:              "NamespaceCreationFailed",
		Action:              "CreateNamespace",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Namespace creation failed.",
	},
	"NamespaceDeleted": {
		Name:                "NamespaceDeleted",
		Reason:              "NamespaceDeleted",
		Action:              "DeleteNamespace",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Namespace got deleted.",
	},
	"NamespaceDeletionFailed": {
		Name:                "NamespaceDeletionFailed",
		Reason:              "NamespaceDeletionFailed",
		Action:              "DeleteNamespace",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Namespace deletion failed.",
	},
	"WorkerClusterRoleCreated": {
		Name:                "WorkerClusterRoleCreated",
		Reason:              "WorkerClusterRoleCreated",
		Action:              "CreateWorkerClusterRole",
		Type:                EventTypeNormal,
		ReportingController: "",
		Message:             "Worker cluster role got created.",
	},
	"WorkerClusterRoleCreationFailed": {
		Name:                "WorkerClusterRoleCreationFailed",
		Reason:              "WorkerClusterRoleCreationFailed",
		Action:              "CreateWorkerClusterRole",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Worker cluster role creation failed.",
	},
	"WorkerClusterRoleUpdated": {
		Name:                "WorkerClusterRoleUpdated",
		Reason:              "WorkerClusterRoleUpdated",
		Action:              "UpdateWorkerClusterRole",
		Type:                EventTypeNormal,
		ReportingController: "",
		Message:             "Worker cluster role got updated.",
	},
	"WorkerClusterRoleUpdateFailed": {
		Name:                "WorkerClusterRoleUpdateFailed",
		Reason:              "WorkerClusterRoleUpdateFailed",
		Action:              "UpdateWorkerClusterRole",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Worker cluster role update failed.",
	},
	"ReadOnlyRoleCreated": {
		Name:                "ReadOnlyRoleCreated",
		Reason:              "ReadOnlyRoleCreated",
		Action:              "CreateReadOnlyRole",
		Type:                EventTypeNormal,
		ReportingController: "",
		Message:             "Read only role got created.",
	},
	"ReadOnlyRoleCreationFailed": {
		Name:                "ReadOnlyRoleCreationFailed",
		Reason:              "ReadOnlyRoleCreationFailed",
		Action:              "CreateReadOnlyRole",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Read only role creation failed.",
	},
	"ReadOnlyRoleUpdated": {
		Name:                "ReadOnlyRoleUpdated",
		Reason:              "ReadOnlyRoleUpdated",
		Action:              "UpdateReadOnlyRole",
		Type:                EventTypeNormal,
		ReportingController: "",
		Message:             "Read only role got updated.",
	},
	"ReadOnlyRoleUpdateFailed": {
		Name:                "ReadOnlyRoleUpdateFailed",
		Reason:              "ReadOnlyRoleUpdateFailed",
		Action:              "UpdateReadOnlyRole",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Read only role update failed.",
	},
	"ReadWriteRoleCreated": {
		Name:                "ReadWriteRoleCreated",
		Reason:              "ReadWriteRoleCreated",
		Action:              "CreateReadWriteRole",
		Type:                EventTypeNormal,
		ReportingController: "",
		Message:             "Read write role got created.",
	},
	"ReadWriteRoleCreationFailed": {
		Name:                "ReadWriteRoleCreationFailed",
		Reason:              "ReadWriteRoleCreationFailed",
		Action:              "CreateReadWriteRole",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Read write role creation failed.",
	},
	"ReadWriteRoleUpdated": {
		Name:                "ReadWriteRoleUpdated",
		Reason:              "ReadWriteRoleUpdated",
		Action:              "UpdateReadWriteRole",
		Type:                EventTypeNormal,
		ReportingController: "",
		Message:             "Read write role got updated.",
	},
	"ReadWriteRoleUpdateFailed": {
		Name:                "ReadWriteRoleUpdateFailed",
		Reason:              "ReadWriteRoleUpdateFailed",
		Action:              "UpdateReadWriteRole",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Read write role update failed.",
	},
	"ServiceAccountCreated": {
		Name:                "ServiceAccountCreated",
		Reason:              "ServiceAccountCreated",
		Action:              "CreateServiceAccount",
		Type:                EventTypeNormal,
		ReportingController: "",
		Message:             "Service account got created.",
	},
	"ServiceAccountCreationFailed": {
		Name:                "ServiceAccountCreationFailed",
		Reason:              "ServiceAccountCreationFailed",
		Action:              "CreateServiceAccount",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Service account creation failed.",
	},
	"ServiceAccountSecretCreated": {
		Name:                "ServiceAccountSecretCreated",
		Reason:              "ServiceAccountSecretCreated",
		Action:              "CreateServiceAccountSecret",
		Type:                EventTypeNormal,
		ReportingController: "",
		Message:             "Service account secret got created.",
	},
	"ServiceAccountSecretCreationFailed": {
		Name:                "ServiceAccountSecretCreationFailed",
		Reason:              "ServiceAccountSecretCreationFailed",
		Action:              "CreateServiceAccountSecret",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Service account secret creation failed.",
	},
	"DefaultRoleBindingCreated": {
		Name:                "DefaultRoleBindingCreated",
		Reason:              "DefaultRoleBindingCreated",
		Action:              "CreateDefaultRoleBinding",
		Type:                EventTypeNormal,
		ReportingController: "",
		Message:             "Default role binding got created.",
	},
	"DefaultRoleBindingCreationFailed": {
		Name:                "DefaultRoleBindingCreationFailed",
		Reason:              "DefaultRoleBindingCreationFailed",
		Action:              "CreateDefaultRoleBinding",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Default role binding creation failed.",
	},
	"DefaultRoleBindingUpdated": {
		Name:                "DefaultRoleBindingUpdated",
		Reason:              "DefaultRoleBindingUpdated",
		Action:              "UpdateDefaultRoleBinding",
		Type:                EventTypeNormal,
		ReportingController: "",
		Message:             "Default role binding got updated.",
	},
	"DefaultRoleBindingUpdateFailed": {
		Name:                "DefaultRoleBindingUpdateFailed",
		Reason:              "DefaultRoleBindingUpdateFailed",
		Action:              "UpdateDefaultRoleBinding",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Default role binding update failed.",
	},
	"DefaultRoleBindingDeleted": {
		Name:                "DefaultRoleBindingDeleted",
		Reason:              "DefaultRoleBindingDeleted",
		Action:              "DeleteDefaultRoleBinding",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Default role binding got deleted.",
	},
	"DefaultRoleBindingDeletionFailed": {
		Name:                "DefaultRoleBindingDeletionFailed",
		Reason:              "DefaultRoleBindingDeletionFailed",
		Action:              "DeleteDefaultRoleBinding",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Default role binding deletion failed.",
	},
	"InactiveRoleBindingDeleted": {
		Name:                "InactiveRoleBindingDeleted",
		Reason:              "InactiveRoleBindingDeleted",
		Action:              "DeleteInactiveRoleBinding",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Inactive role binding got deleted.",
	},
	"InactiveRoleBindingDeletionFailed": {
		Name:                "InactiveRoleBindingDeletionFailed",
		Reason:              "InactiveRoleBindingDeletionFailed",
		Action:              "DeleteInactiveRoleBinding",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Inactive role binding deletion failed.",
	},
	"InactiveServiceAccountDeleted": {
		Name:                "InactiveServiceAccountDeleted",
		Reason:              "InactiveServiceAccountDeleted",
		Action:              "DeleteInactiveServiceAccount",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Inactive service account got deleted.",
	},
	"InactiveServiceAccountDeletionFailed": {
		Name:                "InactiveServiceAccountDeletionFailed",
		Reason:              "InactiveServiceAccountDeletionFailed",
		Action:              "DeleteInactiveServiceAccount",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Inactive service account deletion failed.",
	},
	"ServiceAccountDeleted": {
		Name:                "ServiceAccountDeleted",
		Reason:              "ServiceAccountDeleted",
		Action:              "DeleteServiceAccount",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Service account got deleted.",
	},
	"ServiceAccountDeletionFailed": {
		Name:                "ServiceAccountDeletionFailed",
		Reason:              "ServiceAccountDeletionFailed",
		Action:              "DeleteServiceAccount",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Service account deletion failed.",
	},
	"NetPolViolation": {
		Name:                "NetPolViolation",
		Reason:              "PolicyViolation",
		Action:              "PolicyMonitoring",
		Type:                EventTypeWarning,
		ReportingController: "",
		Message:             "Network policy violation - please ask admin to check the network policy configuration on the worker cluster. <link to tech doc event-list>",
	},
}

var (
	EventLicenseSecretNotFound                = "LicenseSecretNotFound"
	EventMachineFileNotFound                  = "MachineFileNotFound"
	EventMachineFileInvalid                   = "MachineFileInvalid"
	EventLicenseKeyInvalid                    = "LicenseKeyInvalid"
	EventLicenseExpired                       = "LicenseExpired"
	EventLicenseExpiredGracePeriodOn          = "LicenseExpiredGracePeriodOn"
	EventMachineFingerPrintErr                = "MachineFingerPrintErr"
	EventGotMachineFingerPrint                = "GotMachineFingerPrint"
	EventConfigMapErr                         = "ConfigMapErr"
	EventGotConfigMap                         = "GotConfigMap"
	EventLicenseProxyUnreachable              = "LicenseProxyUnreachable"
	EventLicenseDeployError                   = "LicenseDeployError"
	EventLicenseDeploySuccess                 = "LicenseDeploySuccess"
	EventClusterMetadataCollectionFailed      = "ClusterMetadataCollectionFailed"
	EventClusterMetadataCollectionSuccess     = "ClusterMetadataCollectionSuccess"
	EventLicenseDataFetchError                = "LicenseDataFetchError"
	EventLicenseDataFetchSuccess              = "LicenseDataFetchSuccess"
	EventLicenseSecretCreationFailed          = "LicenseSecretCreationFailed"
	EventLicenseSecretCreationSuccess         = "LicenseSecretCreationSuccess"
	EventProjectDeleted                       = "ProjectDeleted"
	EventProjectDeletionFailed                = "ProjectDeletionFailed"
	EventClusterDeleted                       = "ClusterDeleted"
	EventClusterDeletionFailed                = "ClusterDeletionFailed"
	EventSliceConfigDeleted                   = "SliceConfigDeleted"
	EventSliceConfigDeletionFailed            = "SliceConfigDeletionFailed"
	EventServiceExportConfigDeleted           = "ServiceExportConfigDeleted"
	EventServiceExportConfigDeletionFailed    = "ServiceExportConfigDeletionFailed"
	EventSliceQoSConfigDeleted                = "SliceQoSConfigDeleted"
	EventSliceQoSConfigDeletionFailed         = "SliceQoSConfigDeletionFailed"
	EventSecretDeleted                        = "SecretDeleted"
	EventSecretDeletionFailed                 = "SecretDeletionFailed"
	EventNamespaceCreated                     = "NamespaceCreated"
	EventNamespaceCreationFailed              = "NamespaceCreationFailed"
	EventNamespaceDeleted                     = "NamespaceDeleted"
	EventNamespaceDeletionFailed              = "NamespaceDeletionFailed"
	EventWorkerClusterRoleCreated             = "WorkerClusterRoleCreated"
	EventWorkerClusterRoleCreationFailed      = "WorkerClusterRoleCreationFailed"
	EventWorkerClusterRoleUpdated             = "WorkerClusterRoleUpdated"
	EventWorkerClusterRoleUpdateFailed        = "WorkerClusterRoleUpdateFailed"
	EventReadOnlyRoleCreated                  = "ReadOnlyRoleCreated"
	EventReadOnlyRoleCreationFailed           = "ReadOnlyRoleCreationFailed"
	EventReadOnlyRoleUpdated                  = "ReadOnlyRoleUpdated"
	EventReadOnlyRoleUpdateFailed             = "ReadOnlyRoleUpdateFailed"
	EventReadWriteRoleCreated                 = "ReadWriteRoleCreated"
	EventReadWriteRoleCreationFailed          = "ReadWriteRoleCreationFailed"
	EventReadWriteRoleUpdated                 = "ReadWriteRoleUpdated"
	EventReadWriteRoleUpdateFailed            = "ReadWriteRoleUpdateFailed"
	EventServiceAccountCreated                = "ServiceAccountCreated"
	EventServiceAccountCreationFailed         = "ServiceAccountCreationFailed"
	EventServiceAccountSecretCreated          = "ServiceAccountSecretCreated"
	EventServiceAccountSecretCreationFailed   = "ServiceAccountSecretCreationFailed"
	EventDefaultRoleBindingCreated            = "DefaultRoleBindingCreated"
	EventDefaultRoleBindingCreationFailed     = "DefaultRoleBindingCreationFailed"
	EventDefaultRoleBindingUpdated            = "DefaultRoleBindingUpdated"
	EventDefaultRoleBindingUpdateFailed       = "DefaultRoleBindingUpdateFailed"
	EventDefaultRoleBindingDeleted            = "DefaultRoleBindingDeleted"
	EventDefaultRoleBindingDeletionFailed     = "DefaultRoleBindingDeletionFailed"
	EventInactiveRoleBindingDeleted           = "InactiveRoleBindingDeleted"
	EventInactiveRoleBindingDeletionFailed    = "InactiveRoleBindingDeletionFailed"
	EventInactiveServiceAccountDeleted        = "InactiveServiceAccountDeleted"
	EventInactiveServiceAccountDeletionFailed = "InactiveServiceAccountDeletionFailed"
	EventServiceAccountDeleted                = "ServiceAccountDeleted"
	EventServiceAccountDeletionFailed         = "ServiceAccountDeletionFailed"
	EventNetPolViolation                      = "NetPolViolation"
)
