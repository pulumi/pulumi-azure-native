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
    'GetClusterPoolClusterResult',
    'AwaitableGetClusterPoolClusterResult',
    'get_cluster_pool_cluster',
    'get_cluster_pool_cluster_output',
]

@pulumi.output_type
class GetClusterPoolClusterResult:
    """
    The cluster.
    """
    def __init__(__self__, azure_api_version=None, cluster_profile=None, cluster_type=None, compute_profile=None, deployment_id=None, id=None, location=None, name=None, provisioning_state=None, status=None, system_data=None, tags=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if cluster_profile and not isinstance(cluster_profile, dict):
            raise TypeError("Expected argument 'cluster_profile' to be a dict")
        pulumi.set(__self__, "cluster_profile", cluster_profile)
        if cluster_type and not isinstance(cluster_type, str):
            raise TypeError("Expected argument 'cluster_type' to be a str")
        pulumi.set(__self__, "cluster_type", cluster_type)
        if compute_profile and not isinstance(compute_profile, dict):
            raise TypeError("Expected argument 'compute_profile' to be a dict")
        pulumi.set(__self__, "compute_profile", compute_profile)
        if deployment_id and not isinstance(deployment_id, str):
            raise TypeError("Expected argument 'deployment_id' to be a str")
        pulumi.set(__self__, "deployment_id", deployment_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="clusterProfile")
    def cluster_profile(self) -> 'outputs.ClusterProfileResponse':
        """
        Cluster profile.
        """
        return pulumi.get(self, "cluster_profile")

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> builtins.str:
        """
        The type of cluster.
        """
        return pulumi.get(self, "cluster_type")

    @property
    @pulumi.getter(name="computeProfile")
    def compute_profile(self) -> 'outputs.ClusterPoolComputeProfileResponse':
        """
        The compute profile.
        """
        return pulumi.get(self, "compute_profile")

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> builtins.str:
        """
        A unique id generated by the RP to identify the resource.
        """
        return pulumi.get(self, "deployment_id")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        Business status of the resource.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetClusterPoolClusterResult(GetClusterPoolClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterPoolClusterResult(
            azure_api_version=self.azure_api_version,
            cluster_profile=self.cluster_profile,
            cluster_type=self.cluster_type,
            compute_profile=self.compute_profile,
            deployment_id=self.deployment_id,
            id=self.id,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            status=self.status,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type)


def get_cluster_pool_cluster(cluster_name: Optional[builtins.str] = None,
                             cluster_pool_name: Optional[builtins.str] = None,
                             resource_group_name: Optional[builtins.str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterPoolClusterResult:
    """
    Gets a HDInsight cluster.

    Uses Azure REST API version 2024-05-01-preview.

    Other available API versions: 2023-06-01-preview, 2023-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hdinsight [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str cluster_name: The name of the HDInsight cluster.
    :param builtins.str cluster_pool_name: The name of the cluster pool.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['clusterPoolName'] = cluster_pool_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:hdinsight:getClusterPoolCluster', __args__, opts=opts, typ=GetClusterPoolClusterResult).value

    return AwaitableGetClusterPoolClusterResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        cluster_profile=pulumi.get(__ret__, 'cluster_profile'),
        cluster_type=pulumi.get(__ret__, 'cluster_type'),
        compute_profile=pulumi.get(__ret__, 'compute_profile'),
        deployment_id=pulumi.get(__ret__, 'deployment_id'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        status=pulumi.get(__ret__, 'status'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_cluster_pool_cluster_output(cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                                    cluster_pool_name: Optional[pulumi.Input[builtins.str]] = None,
                                    resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetClusterPoolClusterResult]:
    """
    Gets a HDInsight cluster.

    Uses Azure REST API version 2024-05-01-preview.

    Other available API versions: 2023-06-01-preview, 2023-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hdinsight [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str cluster_name: The name of the HDInsight cluster.
    :param builtins.str cluster_pool_name: The name of the cluster pool.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['clusterPoolName'] = cluster_pool_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:hdinsight:getClusterPoolCluster', __args__, opts=opts, typ=GetClusterPoolClusterResult)
    return __ret__.apply(lambda __response__: GetClusterPoolClusterResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        cluster_profile=pulumi.get(__response__, 'cluster_profile'),
        cluster_type=pulumi.get(__response__, 'cluster_type'),
        compute_profile=pulumi.get(__response__, 'compute_profile'),
        deployment_id=pulumi.get(__response__, 'deployment_id'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        status=pulumi.get(__response__, 'status'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
