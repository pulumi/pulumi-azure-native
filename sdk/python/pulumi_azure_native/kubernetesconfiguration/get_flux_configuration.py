# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs

__all__ = [
    'GetFluxConfigurationResult',
    'AwaitableGetFluxConfigurationResult',
    'get_flux_configuration',
    'get_flux_configuration_output',
]

@pulumi.output_type
class GetFluxConfigurationResult:
    """
    The Flux Configuration object returned in Get & Put response.
    """
    def __init__(__self__, azure_api_version=None, azure_blob=None, bucket=None, compliance_state=None, configuration_protected_settings=None, error_message=None, git_repository=None, id=None, kustomizations=None, name=None, namespace=None, provisioning_state=None, reconciliation_wait_duration=None, repository_public_key=None, scope=None, source_kind=None, source_synced_commit_id=None, source_updated_at=None, status_updated_at=None, statuses=None, suspend=None, system_data=None, type=None, wait_for_reconciliation=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if azure_blob and not isinstance(azure_blob, dict):
            raise TypeError("Expected argument 'azure_blob' to be a dict")
        pulumi.set(__self__, "azure_blob", azure_blob)
        if bucket and not isinstance(bucket, dict):
            raise TypeError("Expected argument 'bucket' to be a dict")
        pulumi.set(__self__, "bucket", bucket)
        if compliance_state and not isinstance(compliance_state, str):
            raise TypeError("Expected argument 'compliance_state' to be a str")
        pulumi.set(__self__, "compliance_state", compliance_state)
        if configuration_protected_settings and not isinstance(configuration_protected_settings, dict):
            raise TypeError("Expected argument 'configuration_protected_settings' to be a dict")
        pulumi.set(__self__, "configuration_protected_settings", configuration_protected_settings)
        if error_message and not isinstance(error_message, str):
            raise TypeError("Expected argument 'error_message' to be a str")
        pulumi.set(__self__, "error_message", error_message)
        if git_repository and not isinstance(git_repository, dict):
            raise TypeError("Expected argument 'git_repository' to be a dict")
        pulumi.set(__self__, "git_repository", git_repository)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kustomizations and not isinstance(kustomizations, dict):
            raise TypeError("Expected argument 'kustomizations' to be a dict")
        pulumi.set(__self__, "kustomizations", kustomizations)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if reconciliation_wait_duration and not isinstance(reconciliation_wait_duration, str):
            raise TypeError("Expected argument 'reconciliation_wait_duration' to be a str")
        pulumi.set(__self__, "reconciliation_wait_duration", reconciliation_wait_duration)
        if repository_public_key and not isinstance(repository_public_key, str):
            raise TypeError("Expected argument 'repository_public_key' to be a str")
        pulumi.set(__self__, "repository_public_key", repository_public_key)
        if scope and not isinstance(scope, str):
            raise TypeError("Expected argument 'scope' to be a str")
        pulumi.set(__self__, "scope", scope)
        if source_kind and not isinstance(source_kind, str):
            raise TypeError("Expected argument 'source_kind' to be a str")
        pulumi.set(__self__, "source_kind", source_kind)
        if source_synced_commit_id and not isinstance(source_synced_commit_id, str):
            raise TypeError("Expected argument 'source_synced_commit_id' to be a str")
        pulumi.set(__self__, "source_synced_commit_id", source_synced_commit_id)
        if source_updated_at and not isinstance(source_updated_at, str):
            raise TypeError("Expected argument 'source_updated_at' to be a str")
        pulumi.set(__self__, "source_updated_at", source_updated_at)
        if status_updated_at and not isinstance(status_updated_at, str):
            raise TypeError("Expected argument 'status_updated_at' to be a str")
        pulumi.set(__self__, "status_updated_at", status_updated_at)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)
        if suspend and not isinstance(suspend, bool):
            raise TypeError("Expected argument 'suspend' to be a bool")
        pulumi.set(__self__, "suspend", suspend)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if wait_for_reconciliation and not isinstance(wait_for_reconciliation, bool):
            raise TypeError("Expected argument 'wait_for_reconciliation' to be a bool")
        pulumi.set(__self__, "wait_for_reconciliation", wait_for_reconciliation)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="azureBlob")
    def azure_blob(self) -> Optional['outputs.AzureBlobDefinitionResponse']:
        """
        Parameters to reconcile to the AzureBlob source kind type.
        """
        return pulumi.get(self, "azure_blob")

    @property
    @pulumi.getter
    def bucket(self) -> Optional['outputs.BucketDefinitionResponse']:
        """
        Parameters to reconcile to the Bucket source kind type.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="complianceState")
    def compliance_state(self) -> builtins.str:
        """
        Combined status of the Flux Kubernetes resources created by the fluxConfiguration or created by the managed objects.
        """
        return pulumi.get(self, "compliance_state")

    @property
    @pulumi.getter(name="configurationProtectedSettings")
    def configuration_protected_settings(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Key-value pairs of protected configuration settings for the configuration
        """
        return pulumi.get(self, "configuration_protected_settings")

    @property
    @pulumi.getter(name="errorMessage")
    def error_message(self) -> builtins.str:
        """
        Error message returned to the user in the case of provisioning failure.
        """
        return pulumi.get(self, "error_message")

    @property
    @pulumi.getter(name="gitRepository")
    def git_repository(self) -> Optional['outputs.GitRepositoryDefinitionResponse']:
        """
        Parameters to reconcile to the GitRepository source kind type.
        """
        return pulumi.get(self, "git_repository")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def kustomizations(self) -> Optional[Mapping[str, 'outputs.KustomizationDefinitionResponse']]:
        """
        Array of kustomizations used to reconcile the artifact pulled by the source type on the cluster.
        """
        return pulumi.get(self, "kustomizations")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[builtins.str]:
        """
        The namespace to which this configuration is installed to. Maximum of 253 lower case alphanumeric characters, hyphen and period only.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Status of the creation of the fluxConfiguration.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="reconciliationWaitDuration")
    def reconciliation_wait_duration(self) -> Optional[builtins.str]:
        """
        Maximum duration to wait for flux configuration reconciliation. E.g PT1H, PT5M, P1D
        """
        return pulumi.get(self, "reconciliation_wait_duration")

    @property
    @pulumi.getter(name="repositoryPublicKey")
    def repository_public_key(self) -> builtins.str:
        """
        Public Key associated with this fluxConfiguration (either generated within the cluster or provided by the user).
        """
        return pulumi.get(self, "repository_public_key")

    @property
    @pulumi.getter
    def scope(self) -> Optional[builtins.str]:
        """
        Scope at which the operator will be installed.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter(name="sourceKind")
    def source_kind(self) -> Optional[builtins.str]:
        """
        Source Kind to pull the configuration data from.
        """
        return pulumi.get(self, "source_kind")

    @property
    @pulumi.getter(name="sourceSyncedCommitId")
    def source_synced_commit_id(self) -> builtins.str:
        """
        Branch and/or SHA of the source commit synced with the cluster.
        """
        return pulumi.get(self, "source_synced_commit_id")

    @property
    @pulumi.getter(name="sourceUpdatedAt")
    def source_updated_at(self) -> builtins.str:
        """
        Datetime the fluxConfiguration synced its source on the cluster.
        """
        return pulumi.get(self, "source_updated_at")

    @property
    @pulumi.getter(name="statusUpdatedAt")
    def status_updated_at(self) -> builtins.str:
        """
        Datetime the fluxConfiguration synced its status on the cluster with Azure.
        """
        return pulumi.get(self, "status_updated_at")

    @property
    @pulumi.getter
    def statuses(self) -> Sequence['outputs.ObjectStatusDefinitionResponse']:
        """
        Statuses of the Flux Kubernetes resources created by the fluxConfiguration or created by the managed objects provisioned by the fluxConfiguration.
        """
        return pulumi.get(self, "statuses")

    @property
    @pulumi.getter
    def suspend(self) -> Optional[builtins.bool]:
        """
        Whether this configuration should suspend its reconciliation of its kustomizations and sources.
        """
        return pulumi.get(self, "suspend")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Top level metadata https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="waitForReconciliation")
    def wait_for_reconciliation(self) -> Optional[builtins.bool]:
        """
        Whether flux configuration deployment should wait for cluster to reconcile the kustomizations.
        """
        return pulumi.get(self, "wait_for_reconciliation")


class AwaitableGetFluxConfigurationResult(GetFluxConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFluxConfigurationResult(
            azure_api_version=self.azure_api_version,
            azure_blob=self.azure_blob,
            bucket=self.bucket,
            compliance_state=self.compliance_state,
            configuration_protected_settings=self.configuration_protected_settings,
            error_message=self.error_message,
            git_repository=self.git_repository,
            id=self.id,
            kustomizations=self.kustomizations,
            name=self.name,
            namespace=self.namespace,
            provisioning_state=self.provisioning_state,
            reconciliation_wait_duration=self.reconciliation_wait_duration,
            repository_public_key=self.repository_public_key,
            scope=self.scope,
            source_kind=self.source_kind,
            source_synced_commit_id=self.source_synced_commit_id,
            source_updated_at=self.source_updated_at,
            status_updated_at=self.status_updated_at,
            statuses=self.statuses,
            suspend=self.suspend,
            system_data=self.system_data,
            type=self.type,
            wait_for_reconciliation=self.wait_for_reconciliation)


def get_flux_configuration(cluster_name: Optional[builtins.str] = None,
                           cluster_resource_name: Optional[builtins.str] = None,
                           cluster_rp: Optional[builtins.str] = None,
                           flux_configuration_name: Optional[builtins.str] = None,
                           resource_group_name: Optional[builtins.str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFluxConfigurationResult:
    """
    Gets details of the Flux Configuration.

    Uses Azure REST API version 2023-05-01.

    Other available API versions: 2022-07-01, 2022-11-01, 2024-04-01-preview, 2024-11-01, 2025-04-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native kubernetesconfiguration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str cluster_name: The name of the kubernetes cluster.
    :param builtins.str cluster_resource_name: The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
    :param builtins.str cluster_rp: The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
    :param builtins.str flux_configuration_name: Name of the Flux Configuration.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['clusterResourceName'] = cluster_resource_name
    __args__['clusterRp'] = cluster_rp
    __args__['fluxConfigurationName'] = flux_configuration_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:kubernetesconfiguration:getFluxConfiguration', __args__, opts=opts, typ=GetFluxConfigurationResult).value

    return AwaitableGetFluxConfigurationResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        azure_blob=pulumi.get(__ret__, 'azure_blob'),
        bucket=pulumi.get(__ret__, 'bucket'),
        compliance_state=pulumi.get(__ret__, 'compliance_state'),
        configuration_protected_settings=pulumi.get(__ret__, 'configuration_protected_settings'),
        error_message=pulumi.get(__ret__, 'error_message'),
        git_repository=pulumi.get(__ret__, 'git_repository'),
        id=pulumi.get(__ret__, 'id'),
        kustomizations=pulumi.get(__ret__, 'kustomizations'),
        name=pulumi.get(__ret__, 'name'),
        namespace=pulumi.get(__ret__, 'namespace'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        reconciliation_wait_duration=pulumi.get(__ret__, 'reconciliation_wait_duration'),
        repository_public_key=pulumi.get(__ret__, 'repository_public_key'),
        scope=pulumi.get(__ret__, 'scope'),
        source_kind=pulumi.get(__ret__, 'source_kind'),
        source_synced_commit_id=pulumi.get(__ret__, 'source_synced_commit_id'),
        source_updated_at=pulumi.get(__ret__, 'source_updated_at'),
        status_updated_at=pulumi.get(__ret__, 'status_updated_at'),
        statuses=pulumi.get(__ret__, 'statuses'),
        suspend=pulumi.get(__ret__, 'suspend'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'),
        wait_for_reconciliation=pulumi.get(__ret__, 'wait_for_reconciliation'))
def get_flux_configuration_output(cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                                  cluster_resource_name: Optional[pulumi.Input[builtins.str]] = None,
                                  cluster_rp: Optional[pulumi.Input[builtins.str]] = None,
                                  flux_configuration_name: Optional[pulumi.Input[builtins.str]] = None,
                                  resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFluxConfigurationResult]:
    """
    Gets details of the Flux Configuration.

    Uses Azure REST API version 2023-05-01.

    Other available API versions: 2022-07-01, 2022-11-01, 2024-04-01-preview, 2024-11-01, 2025-04-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native kubernetesconfiguration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str cluster_name: The name of the kubernetes cluster.
    :param builtins.str cluster_resource_name: The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
    :param builtins.str cluster_rp: The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
    :param builtins.str flux_configuration_name: Name of the Flux Configuration.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['clusterResourceName'] = cluster_resource_name
    __args__['clusterRp'] = cluster_rp
    __args__['fluxConfigurationName'] = flux_configuration_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:kubernetesconfiguration:getFluxConfiguration', __args__, opts=opts, typ=GetFluxConfigurationResult)
    return __ret__.apply(lambda __response__: GetFluxConfigurationResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        azure_blob=pulumi.get(__response__, 'azure_blob'),
        bucket=pulumi.get(__response__, 'bucket'),
        compliance_state=pulumi.get(__response__, 'compliance_state'),
        configuration_protected_settings=pulumi.get(__response__, 'configuration_protected_settings'),
        error_message=pulumi.get(__response__, 'error_message'),
        git_repository=pulumi.get(__response__, 'git_repository'),
        id=pulumi.get(__response__, 'id'),
        kustomizations=pulumi.get(__response__, 'kustomizations'),
        name=pulumi.get(__response__, 'name'),
        namespace=pulumi.get(__response__, 'namespace'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        reconciliation_wait_duration=pulumi.get(__response__, 'reconciliation_wait_duration'),
        repository_public_key=pulumi.get(__response__, 'repository_public_key'),
        scope=pulumi.get(__response__, 'scope'),
        source_kind=pulumi.get(__response__, 'source_kind'),
        source_synced_commit_id=pulumi.get(__response__, 'source_synced_commit_id'),
        source_updated_at=pulumi.get(__response__, 'source_updated_at'),
        status_updated_at=pulumi.get(__response__, 'status_updated_at'),
        statuses=pulumi.get(__response__, 'statuses'),
        suspend=pulumi.get(__response__, 'suspend'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type'),
        wait_for_reconciliation=pulumi.get(__response__, 'wait_for_reconciliation')))
