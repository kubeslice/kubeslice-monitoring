package schema

var Events = map[string]EventSchema {
  EventSliceDeletionFailed: {
    Reason: "SliceDeletionFailed",
    Action: "SliceDeletion",
    Type: "Warning",
    ReportingController: "controller",
    Message: "Slice deletion failed - please ask admin to check the slice configuration on the worker cluster. <link to tech doc event-list>",
  },
  EventLicenseSecretNotFound: {
    Reason: "LicenseNotFound",
    Action: "LicenseValidation",
    Type: "Warning",
    ReportingController: "controller",
    Message: "Secret with license config not found. Please ensure to create license secret if this is airgapped cluster.",
  },
  EventMachineFileNotFound: {
    Reason: "MachineFileNotFound",
    Action: "FetchMachineFileFromSecret",
    Type: "Warning",
    ReportingController: "controller",
    Message: "Machine File not present in license secret.Please contact kubeslice support team.",
  },
  EventMachineFileInvalid: {
    Reason: "MachineFileInvalid",
    Action: "VerifyMachineFile",
    Type: "Warning",
    ReportingController: "controller",
    Message: "Machine File is invalid in license secret.Please contact kubeslice support team.",
  },
  EventLicenseKeyInvalid: {
    Reason: "LicenseKeyInvalid",
    Action: "DecryptMachineFile",
    Type: "Warning",
    ReportingController: "controller",
    Message: "License Key is invalid in license secret.Please contact kubeslice support team.",
  },
  EventLicenseExpired: {
    Reason: "LicenseExpired",
    Action: "ValidateLicense",
    Type: "Warning",
    ReportingController: "controller",
    Message: "License Expired , please contact kubeslice support team to get it renewed!",
  },
  EventLicenseExpiredGracePeriodOn: {
    Reason: "LicenseExpiredGracePeriodOn",
    Action: "ValidateLicense",
    Type: "Warning",
    ReportingController: "controller",
    Message: "License Expired , grace period is on ,please contact kubeslice support team to get it renewed!",
  },
  EventProjectDeleted: {
    Reason: "ProjectDeleted",
    Action: "DeleteProject",
    Type: "Warning",
    ReportingController: "controller",
    Message: "Project got deleted, please contact kubeslice support team to get it restored!",
  },
  EventProjectCreated: {
    Reason: "ProjectCreated",
    Action: "CreateProject",
    Type: "Normal",
    ReportingController: "controller",
    Message: "Project got created!",
  },
  EventProjectUpdated: {
    Reason: "ProjectUpdated",
    Action: "UpdateProject",
    Type: "Normal",
    ReportingController: "controller",
    Message: "Project got updated!",
  },
  EventUserCreated: {
    Reason: "UserCreated",
    Action: "CreateUser",
    Type: "Normal",
    ReportingController: "controller",
    Message: "User got created!",
  },
  EventUserRemoved: {
    Reason: "UserRemoved",
    Action: "RemoveUser",
    Type: "Normal",
    ReportingController: "controller",
    Message: "User got removed, please contact kubeslice support team to get it restored!",
  },
  EventClusterDeleted: {
    Reason: "ClusterDeleted",
    Action: "DeleteCluster",
    Type: "Warning",
    ReportingController: "controller",
    Message: "Cluster got deleted, please contact kubeslice support team to get it restored!",
  },
  EventClusterCreated: {
    Reason: "ClusterCreated",
    Action: "CreateCluster",
    Type: "Normal",
    ReportingController: "controller",
    Message: "Cluster got created!",
  },
  EventClusterUpdated: {
    Reason: "ClusterUpdated",
    Action: "UpdateCluster",
    Type: "Normal",
    ReportingController: "controller",
    Message: "Cluster got updated!",
  },
  EventSliceConfigDeleted: {
    Reason: "SliceConfigDeleted",
    Action: "DeleteSliceConfig",
    Type: "Warning",
    ReportingController: "controller",
    Message: "Slice config got deleted, please contact kubeslice support team to get it restored!",
  },
  EventSliceConfigCreated: {
    Reason: "SliceConfigCreated",
    Action: "CreateSliceConfig",
    Type: "Normal",
    ReportingController: "controller",
    Message: "Slice config got created!",
  },
  EventSliceConfigUpdated: {
    Reason: "SliceConfigUpdated",
    Action: "UpdateSliceConfig",
    Type: "Normal",
    ReportingController: "controller",
    Message: "Slice config got updated!",
  },
  EventServiceExportConfigDeleted: {
    Reason: "ServiceExportConfigDeleted",
    Action: "DeleteServiceExportConfig",
    Type: "Warning",
    ReportingController: "controller",
    Message: "Service export config got deleted, please contact kubeslice support team to get it restored!",
  },
  EventServiceExportConfigCreated: {
    Reason: "ServiceExportConfigCreated",
    Action: "CreateServiceExportConfig",
    Type: "Normal",
    ReportingController: "controller",
    Message: "Service export config got created!",
  },
  EventServiceExportConfigUpdated: {
    Reason: "ServiceExportConfigUpdated",
    Action: "UpdateServiceExportConfig",
    Type: "Normal",
    ReportingController: "controller",
    Message: "Service export config got updated!",
  },
  EventSliceQoSConfigDeleted: {
    Reason: "SliceQoSConfigDeleted",
    Action: "DeleteSliceQoSConfig",
    Type: "Warning",
    ReportingController: "controller",
    Message: "Slice QoS config got deleted, please contact kubeslice support team to get it restored!",
  },
  EventSliceQoSConfigCreated: {
    Reason: "SliceQoSConfigCreated",
    Action: "CreateSliceQoSConfig",
    Type: "Normal",
    ReportingController: "controller",
    Message: "Slice QoS config got created!",
  },
  EventSliceQoSConfigUpdated: {
    Reason: "SliceQoSConfigUpdated",
    Action: "UpdateSliceQoSConfig",
    Type: "Normal",
    ReportingController: "controller",
    Message: "Slice QoS config got updated!",
  },
}