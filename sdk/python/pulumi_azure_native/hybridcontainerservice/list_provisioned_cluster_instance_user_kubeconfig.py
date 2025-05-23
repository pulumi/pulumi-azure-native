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
    'ListProvisionedClusterInstanceUserKubeconfigResult',
    'AwaitableListProvisionedClusterInstanceUserKubeconfigResult',
    'list_provisioned_cluster_instance_user_kubeconfig',
    'list_provisioned_cluster_instance_user_kubeconfig_output',
]

@pulumi.output_type
class ListProvisionedClusterInstanceUserKubeconfigResult:
    """
    The list kubeconfig result response.
    """
    def __init__(__self__, error=None, id=None, name=None, properties=None, resource_id=None, status=None):
        if error and not isinstance(error, dict):
            raise TypeError("Expected argument 'error' to be a dict")
        pulumi.set(__self__, "error", error)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)
        if resource_id and not isinstance(resource_id, str):
            raise TypeError("Expected argument 'resource_id' to be a str")
        pulumi.set(__self__, "resource_id", resource_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def error(self) -> Optional['outputs.ListCredentialResponseResponseError']:
        return pulumi.get(self, "error")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Operation Id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Operation Name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.ListCredentialResponseResponseProperties':
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> builtins.str:
        """
        ARM Resource Id of the provisioned cluster instance
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        Provisioning state of the resource
        """
        return pulumi.get(self, "status")


class AwaitableListProvisionedClusterInstanceUserKubeconfigResult(ListProvisionedClusterInstanceUserKubeconfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListProvisionedClusterInstanceUserKubeconfigResult(
            error=self.error,
            id=self.id,
            name=self.name,
            properties=self.properties,
            resource_id=self.resource_id,
            status=self.status)


def list_provisioned_cluster_instance_user_kubeconfig(connected_cluster_resource_uri: Optional[builtins.str] = None,
                                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListProvisionedClusterInstanceUserKubeconfigResult:
    """
    Lists the user credentials of the provisioned cluster (can only be used within private network)

    Uses Azure REST API version 2024-01-01.

    Other available API versions: 2023-11-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridcontainerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str connected_cluster_resource_uri: The fully qualified Azure Resource Manager identifier of the connected cluster resource.
    """
    __args__ = dict()
    __args__['connectedClusterResourceUri'] = connected_cluster_resource_uri
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:hybridcontainerservice:listProvisionedClusterInstanceUserKubeconfig', __args__, opts=opts, typ=ListProvisionedClusterInstanceUserKubeconfigResult).value

    return AwaitableListProvisionedClusterInstanceUserKubeconfigResult(
        error=pulumi.get(__ret__, 'error'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        properties=pulumi.get(__ret__, 'properties'),
        resource_id=pulumi.get(__ret__, 'resource_id'),
        status=pulumi.get(__ret__, 'status'))
def list_provisioned_cluster_instance_user_kubeconfig_output(connected_cluster_resource_uri: Optional[pulumi.Input[builtins.str]] = None,
                                                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListProvisionedClusterInstanceUserKubeconfigResult]:
    """
    Lists the user credentials of the provisioned cluster (can only be used within private network)

    Uses Azure REST API version 2024-01-01.

    Other available API versions: 2023-11-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridcontainerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str connected_cluster_resource_uri: The fully qualified Azure Resource Manager identifier of the connected cluster resource.
    """
    __args__ = dict()
    __args__['connectedClusterResourceUri'] = connected_cluster_resource_uri
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:hybridcontainerservice:listProvisionedClusterInstanceUserKubeconfig', __args__, opts=opts, typ=ListProvisionedClusterInstanceUserKubeconfigResult)
    return __ret__.apply(lambda __response__: ListProvisionedClusterInstanceUserKubeconfigResult(
        error=pulumi.get(__response__, 'error'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        properties=pulumi.get(__response__, 'properties'),
        resource_id=pulumi.get(__response__, 'resource_id'),
        status=pulumi.get(__response__, 'status')))
