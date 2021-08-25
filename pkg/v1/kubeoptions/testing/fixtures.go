package testing

import "github.com/selectel/mks-go/pkg/v1/kubeoptions"

var expectedFeatureGates = []*kubeoptions.View{
	{
		KubeVersion: "1.15",
		Names:       []string{"ProcMountType", "RemainingItemCount", "APIResponseCompression", "CSIMigrationOpenStack", "CSIMigrationAzureDisk", "TTLAfterFinished", "VolumePVCDataSource", "ServerSideApply", "BalanceAttachedNodeVolumes", "HyperVContainer", "WinOverlay", "QOSReserved", "SCTPSupport", "CSIInlineVolume", "CustomCPUCFSQuotaPeriod", "VolumeSnapshotDataSource", "WindowsGMSA", "WatchBookmark", "NonPreemptingPriority", "CSIMigrationAzureFile", "ExpandCSIVolumes", "CSIMigration", "ServiceNodeExclusion", "RequestManagement", "ResourceLimitsPriorityFunction", "DynamicAuditing", "ServiceLoadBalancerFinalizer", "CSIMigrationAWS", "LocalStorageCapacityIsolationFSQuotaMonitoring", "WinDSR", "CustomResourceDefaulting", "CSIMigrationGCE"},
	},
	{
		KubeVersion: "1.16",
		Names:       []string{"NonPreemptingPriority", "VolumeSnapshotDataSource", "CSIMigrationAzureFile", "EvenPodsSpread", "CSIMigration", "NodeDisruptionExclusion", "EndpointSlice", "ServiceNodeExclusion", "RequestManagement", "ResourceLimitsPriorityFunction", "CSIMigrationAWS", "LocalStorageCapacityIsolationFSQuotaMonitoring", "DynamicAuditing", "CSIMigrationGCE", "PodOverhead", "WinDSR", "WindowsRunAsUserName", "APIResponseCompression", "CSIMigrationOpenStack", "HPAScaleToZero", "ProcMountType", "RemainingItemCount", "CSIMigrationAzureDisk", "EphemeralContainers", "TTLAfterFinished", "IPv6DualStack", "StartupProbe", "TopologyManager", "BalanceAttachedNodeVolumes", "HyperVContainer", "WinOverlay", "CustomCPUCFSQuotaPeriod", "LegacyNodeRoleBehavior", "QOSReserved", "SCTPSupport"},
	},
	{
		KubeVersion: "1.17",
		Names:       []string{"BalanceAttachedNodeVolumes", "CSIMigrationAzureFileComplete", "HyperVContainer", "WinOverlay", "QOSReserved", "SCTPSupport", "CustomCPUCFSQuotaPeriod", "LegacyNodeRoleBehavior", "NonPreemptingPriority", "EvenPodsSpread", "CSIMigrationAzureFile", "CSIMigrationAzureDiskComplete", "NodeDisruptionExclusion", "ServiceTopology", "ServiceNodeExclusion", "ResourceLimitsPriorityFunction", "CSIMigrationGCEComplete", "CSIMigrationOpenStackComplete", "APIPriorityAndFairness", "DynamicAuditing", "CSIMigrationAWSComplete", "LocalStorageCapacityIsolationFSQuotaMonitoring", "WinDSR", "PodOverhead", "ProcMountType", "RemainingItemCount", "APIResponseCompression", "CSIMigrationOpenStack", "HPAScaleToZero", "CSIMigrationAzureDisk", "EphemeralContainers", "TTLAfterFinished", "IPv6DualStack", "StartupProbe", "TopologyManager"},
	},
	{
		KubeVersion: "1.18",
		Names:       []string{"AnyVolumeDataSource", "APIPriorityAndFairness", "APIResponseCompression", "BalanceAttachedNodeVolumes", "CSIMigrationAWSComplete", "CSIMigrationAzureDisk", "CSIMigrationAzureDiskComplete", "CSIMigrationAzureFile", "CSIMigrationAzureFileComplete", "CSIMigrationGCEComplete", "CSIMigrationOpenStack", "CSIMigrationOpenStackComplete", "ConfigurableFSGroupPolicy", "CustomCPUCFSQuotaPeriod", "DynamicAuditing", "EndpointSliceProxying", "EphemeralContainers", "HPAScaleToZero", "HugePageStorageMediumSize", "ImmutableEphemeralVolumes", "IPv6DualStack", "LegacyNodeRoleBehavior", "LocalStorageCapacityIsolationFSQuotaMonitoring", "NodeDisruptionExclusion", "NonPreemptingPriority", "PodOverhead", "ProcMountType", "QOSReserved", "RemainingItemCount", "ResourceLimitsPriorityFunction", "ServiceAccountIssuerDiscovery", "ServiceAppProtocol", "ServiceNodeExclusion", "ServiceTopology", "TTLAfterFinished", "TopologyManager"},
	},
}

var expectedAdmissionControllers = []*kubeoptions.View{
	{
		KubeVersion: "1.15",
		Names:       []string{"ExtendedResourceToleration", "ServiceAccount", "DefaultTolerationSeconds", "CertificateApproval", "PodSecurityPolicy", "AlwaysPullImages", "StorageObjectInUseProtection", "ImagePolicyWebhook", "LimitRanger", "NamespaceLifecycle", "PodNodeSelector", "Priority", "EventRateLimit", "PersistentVolumeClaimResize", "PodPreset", "SecurityContextDeny", "LimitPodHardAntiAffinityTopology", "CertificateSubjectRestrictions", "OwnerReferencesPermissionEnforcement", "ResourceQuota", "ValidatingAdmissionWebhook", "CertificateSigning", "NamespaceExists", "PodTolerationRestriction", "TaintNodesByCondition", "DefaultStorageClass", "NamespaceAutoProvision", "MutatingAdmissionWebhook"},
	},
	{
		KubeVersion: "1.16",
		Names:       []string{"AlwaysPullImages", "CertificateApproval", "PodSecurityPolicy", "StorageObjectInUseProtection", "Priority", "EventRateLimit", "ImagePolicyWebhook", "LimitRanger", "NamespaceLifecycle", "PodNodeSelector", "LimitPodHardAntiAffinityTopology", "PersistentVolumeClaimResize", "PodPreset", "SecurityContextDeny", "ValidatingAdmissionWebhook", "CertificateSubjectRestrictions", "OwnerReferencesPermissionEnforcement", "ResourceQuota", "RuntimeClass", "CertificateSigning", "NamespaceExists", "PodTolerationRestriction", "MutatingAdmissionWebhook", "TaintNodesByCondition", "DefaultStorageClass", "NamespaceAutoProvision", "DefaultTolerationSeconds", "ExtendedResourceToleration", "ServiceAccount"},
	},
	{
		KubeVersion: "1.17",
		Names:       []string{"StorageObjectInUseProtection", "TaintNodesByCondition", "LimitPodHardAntiAffinityTopology", "SecurityContextDeny", "LimitRanger", "PodSecurityPolicy", "MutatingAdmissionWebhook", "ValidatingAdmissionWebhook", "AlwaysPullImages", "CertificateSigning", "EventRateLimit", "DefaultStorageClass", "PodPreset", "ResourceQuota", "ExtendedResourceToleration", "NamespaceExists", "OwnerReferencesPermissionEnforcement", "ServiceAccount", "RuntimeClass", "CertificateSubjectRestrictions", "DefaultTolerationSeconds", "NamespaceAutoProvision", "ImagePolicyWebhook", "PodNodeSelector", "PodTolerationRestriction", "CertificateApproval", "NamespaceLifecycle", "PersistentVolumeClaimResize", "Priority"},
	},
	{
		KubeVersion: "1.18",
		Names:       []string{"StorageObjectInUseProtection", "TaintNodesByCondition", "LimitPodHardAntiAffinityTopology", "SecurityContextDeny", "LimitRanger", "PodSecurityPolicy", "MutatingAdmissionWebhook", "ValidatingAdmissionWebhook", "AlwaysPullImages", "CertificateSigning", "EventRateLimit", "DefaultStorageClass", "PodPreset", "ResourceQuota", "ExtendedResourceToleration", "NamespaceExists", "OwnerReferencesPermissionEnforcement", "ServiceAccount", "RuntimeClass", "CertificateSubjectRestrictions", "DefaultTolerationSeconds", "NamespaceAutoProvision", "ImagePolicyWebhook", "PodNodeSelector", "PodTolerationRestriction", "CertificateApproval", "NamespaceLifecycle", "PersistentVolumeClaimResize", "Priority"},
	},
}

var testListFeatureGatesResponseRaw = `{
    "feature_gates": [
        {
            "KubeVersionMinor": "1.15",
            "Names": [
                "ProcMountType",
                "RemainingItemCount",
                "APIResponseCompression",
                "CSIMigrationOpenStack",
                "CSIMigrationAzureDisk",
                "TTLAfterFinished",
                "VolumePVCDataSource",
                "ServerSideApply",
                "BalanceAttachedNodeVolumes",
                "HyperVContainer",
                "WinOverlay",
                "QOSReserved",
                "SCTPSupport",
                "CSIInlineVolume",
                "CustomCPUCFSQuotaPeriod",
                "VolumeSnapshotDataSource",
                "WindowsGMSA",
                "WatchBookmark",
                "NonPreemptingPriority",
                "CSIMigrationAzureFile",
                "ExpandCSIVolumes",
                "CSIMigration",
                "ServiceNodeExclusion",
                "RequestManagement",
                "ResourceLimitsPriorityFunction",
                "DynamicAuditing",
                "ServiceLoadBalancerFinalizer",
                "CSIMigrationAWS",
                "LocalStorageCapacityIsolationFSQuotaMonitoring",
                "WinDSR",
                "CustomResourceDefaulting",
                "CSIMigrationGCE"
            ]
        },
        {
            "KubeVersionMinor": "1.16",
            "Names": [
                "NonPreemptingPriority",
                "VolumeSnapshotDataSource",
                "CSIMigrationAzureFile",
                "EvenPodsSpread",
                "CSIMigration",
                "NodeDisruptionExclusion",
                "EndpointSlice",
                "ServiceNodeExclusion",
                "RequestManagement",
                "ResourceLimitsPriorityFunction",
                "CSIMigrationAWS",
                "LocalStorageCapacityIsolationFSQuotaMonitoring",
                "DynamicAuditing",
                "CSIMigrationGCE",
                "PodOverhead",
                "WinDSR",
                "WindowsRunAsUserName",
                "APIResponseCompression",
                "CSIMigrationOpenStack",
                "HPAScaleToZero",
                "ProcMountType",
                "RemainingItemCount",
                "CSIMigrationAzureDisk",
                "EphemeralContainers",
                "TTLAfterFinished",
                "IPv6DualStack",
                "StartupProbe",
                "TopologyManager",
                "BalanceAttachedNodeVolumes",
                "HyperVContainer",
                "WinOverlay",
                "CustomCPUCFSQuotaPeriod",
                "LegacyNodeRoleBehavior",
                "QOSReserved",
                "SCTPSupport"
            ]
        },
        {
            "KubeVersionMinor": "1.17",
            "Names": [
                "BalanceAttachedNodeVolumes",
                "CSIMigrationAzureFileComplete",
                "HyperVContainer",
                "WinOverlay",
                "QOSReserved",
                "SCTPSupport",
                "CustomCPUCFSQuotaPeriod",
                "LegacyNodeRoleBehavior",
                "NonPreemptingPriority",
                "EvenPodsSpread",
                "CSIMigrationAzureFile",
                "CSIMigrationAzureDiskComplete",
                "NodeDisruptionExclusion",
                "ServiceTopology",
                "ServiceNodeExclusion",
                "ResourceLimitsPriorityFunction",
                "CSIMigrationGCEComplete",
                "CSIMigrationOpenStackComplete",
                "APIPriorityAndFairness",
                "DynamicAuditing",
                "CSIMigrationAWSComplete",
                "LocalStorageCapacityIsolationFSQuotaMonitoring",
                "WinDSR",
                "PodOverhead",
                "ProcMountType",
                "RemainingItemCount",
                "APIResponseCompression",
                "CSIMigrationOpenStack",
                "HPAScaleToZero",
                "CSIMigrationAzureDisk",
                "EphemeralContainers",
                "TTLAfterFinished",
                "IPv6DualStack",
                "StartupProbe",
                "TopologyManager"
            ]
        },
        {
            "KubeVersionMinor": "1.18",
            "Names": [
                "AnyVolumeDataSource",
                "APIPriorityAndFairness",
                "APIResponseCompression",
                "BalanceAttachedNodeVolumes",
                "CSIMigrationAWSComplete",
                "CSIMigrationAzureDisk",
                "CSIMigrationAzureDiskComplete",
                "CSIMigrationAzureFile",
                "CSIMigrationAzureFileComplete",
                "CSIMigrationGCEComplete",
                "CSIMigrationOpenStack",
                "CSIMigrationOpenStackComplete",
                "ConfigurableFSGroupPolicy",
                "CustomCPUCFSQuotaPeriod",
                "DynamicAuditing",
                "EndpointSliceProxying",
                "EphemeralContainers",
                "HPAScaleToZero",
                "HugePageStorageMediumSize",
                "ImmutableEphemeralVolumes",
                "IPv6DualStack",
                "LegacyNodeRoleBehavior",
                "LocalStorageCapacityIsolationFSQuotaMonitoring",
                "NodeDisruptionExclusion",
                "NonPreemptingPriority",
                "PodOverhead",
                "ProcMountType",
                "QOSReserved",
                "RemainingItemCount",
                "ResourceLimitsPriorityFunction",
                "ServiceAccountIssuerDiscovery",
                "ServiceAppProtocol",
                "ServiceNodeExclusion",
                "ServiceTopology",
                "TTLAfterFinished",
                "TopologyManager"
            ]
        }
    ]
}`

var testFeatureGatesAsRawList = `[
		{
            "KubeVersionMinor": "1.15",
            "Names": [
                "ProcMountType",
                "RemainingItemCount",
                "APIResponseCompression",
            ]
        },
        {
            "KubeVersionMinor": "1.16",
            "Names": [
                "NonPreemptingPriority",
                "VolumeSnapshotDataSource",
                "CSIMigrationAzureFile",
            ]
        },
        {
            "KubeVersionMinor": "1.17",
            "Names": [
                "BalanceAttachedNodeVolumes",
                "CSIMigrationAzureFileComplete",
                "HyperVContainer",
            ]
        },
        {
            "KubeVersionMinor": "1.18",
            "Names": [
                "AnyVolumeDataSource",
                "APIPriorityAndFairness",
                "APIResponseCompression",
            ]
        }
    ]`

var testListAdmissionControllersResponseRaw = `{
    "admission_controllers": [
        {
            "KubeVersionMinor": "1.15",
            "Names": [
                "ExtendedResourceToleration",
                "ServiceAccount",
                "DefaultTolerationSeconds",
                "CertificateApproval",
                "PodSecurityPolicy",
                "AlwaysPullImages",
                "StorageObjectInUseProtection",
                "ImagePolicyWebhook",
                "LimitRanger",
                "NamespaceLifecycle",
                "PodNodeSelector",
                "Priority",
                "EventRateLimit",
                "PersistentVolumeClaimResize",
                "PodPreset",
                "SecurityContextDeny",
                "LimitPodHardAntiAffinityTopology",
                "CertificateSubjectRestrictions",
                "OwnerReferencesPermissionEnforcement",
                "ResourceQuota",
                "ValidatingAdmissionWebhook",
                "CertificateSigning",
                "NamespaceExists",
                "PodTolerationRestriction",
                "TaintNodesByCondition",
                "DefaultStorageClass",
                "NamespaceAutoProvision",
                "MutatingAdmissionWebhook"
            ]
        },
        {
            "KubeVersionMinor": "1.16",
            "Names": [
                "AlwaysPullImages",
                "CertificateApproval",
                "PodSecurityPolicy",
                "StorageObjectInUseProtection",
                "Priority",
                "EventRateLimit",
                "ImagePolicyWebhook",
                "LimitRanger",
                "NamespaceLifecycle",
                "PodNodeSelector",
                "LimitPodHardAntiAffinityTopology",
                "PersistentVolumeClaimResize",
                "PodPreset",
                "SecurityContextDeny",
                "ValidatingAdmissionWebhook",
                "CertificateSubjectRestrictions",
                "OwnerReferencesPermissionEnforcement",
                "ResourceQuota",
                "RuntimeClass",
                "CertificateSigning",
                "NamespaceExists",
                "PodTolerationRestriction",
                "MutatingAdmissionWebhook",
                "TaintNodesByCondition",
                "DefaultStorageClass",
                "NamespaceAutoProvision",
                "DefaultTolerationSeconds",
                "ExtendedResourceToleration",
                "ServiceAccount"
            ]
        },
        {
            "KubeVersionMinor": "1.17",
            "Names": [
                "StorageObjectInUseProtection",
                "TaintNodesByCondition",
                "LimitPodHardAntiAffinityTopology",
                "SecurityContextDeny",
                "LimitRanger",
                "PodSecurityPolicy",
                "MutatingAdmissionWebhook",
                "ValidatingAdmissionWebhook",
                "AlwaysPullImages",
                "CertificateSigning",
                "EventRateLimit",
                "DefaultStorageClass",
                "PodPreset",
                "ResourceQuota",
                "ExtendedResourceToleration",
                "NamespaceExists",
                "OwnerReferencesPermissionEnforcement",
                "ServiceAccount",
                "RuntimeClass",
                "CertificateSubjectRestrictions",
                "DefaultTolerationSeconds",
                "NamespaceAutoProvision",
                "ImagePolicyWebhook",
                "PodNodeSelector",
                "PodTolerationRestriction",
                "CertificateApproval",
                "NamespaceLifecycle",
                "PersistentVolumeClaimResize",
                "Priority"
            ]
        },
        {
            "KubeVersionMinor": "1.18",
            "Names": [
                "StorageObjectInUseProtection",
                "TaintNodesByCondition",
                "LimitPodHardAntiAffinityTopology",
                "SecurityContextDeny",
                "LimitRanger",
                "PodSecurityPolicy",
                "MutatingAdmissionWebhook",
                "ValidatingAdmissionWebhook",
                "AlwaysPullImages",
                "CertificateSigning",
                "EventRateLimit",
                "DefaultStorageClass",
                "PodPreset",
                "ResourceQuota",
                "ExtendedResourceToleration",
                "NamespaceExists",
                "OwnerReferencesPermissionEnforcement",
                "ServiceAccount",
                "RuntimeClass",
                "CertificateSubjectRestrictions",
                "DefaultTolerationSeconds",
                "NamespaceAutoProvision",
                "ImagePolicyWebhook",
                "PodNodeSelector",
                "PodTolerationRestriction",
                "CertificateApproval",
                "NamespaceLifecycle",
                "PersistentVolumeClaimResize",
                "Priority"
            ]
        }
    ]
}`

var testAdmissionControllersAsRawList = `[
        {
            "KubeVersionMinor": "1.15",
            "Names": [
                "ExtendedResourceToleration",
                "ServiceAccount",
                "DefaultTolerationSeconds",
            ]
        },
        {
            "KubeVersionMinor": "1.16",
            "Names": [
                "AlwaysPullImages",
                "CertificateApproval",
                "PodSecurityPolicy",
            ]
        },
        {
            "KubeVersionMinor": "1.17",
            "Names": [
                "StorageObjectInUseProtection",
                "TaintNodesByCondition",
                "LimitPodHardAntiAffinityTopology",
            ]
        },
        {
            "KubeVersionMinor": "1.18",
            "Names": [
                "StorageObjectInUseProtection",
                "TaintNodesByCondition",
                "LimitPodHardAntiAffinityTopology",
            ]
        }
    ]`
