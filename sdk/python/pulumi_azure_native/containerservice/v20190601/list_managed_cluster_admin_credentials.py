# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from ... import _utilities
from . import outputs

__all__ = [
    'ListManagedClusterAdminCredentialsResult',
    'AwaitableListManagedClusterAdminCredentialsResult',
    'list_managed_cluster_admin_credentials',
    'list_managed_cluster_admin_credentials_output',
]

@pulumi.output_type
class ListManagedClusterAdminCredentialsResult:
    """
    The list of credential result response.
    """
    def __init__(__self__, kubeconfigs=None):
        if kubeconfigs and not isinstance(kubeconfigs, list):
            raise TypeError("Expected argument 'kubeconfigs' to be a list")
        pulumi.set(__self__, "kubeconfigs", kubeconfigs)

    @property
    @pulumi.getter
    def kubeconfigs(self) -> Sequence['outputs.CredentialResultResponse']:
        """
        Base64-encoded Kubernetes configuration file.
        """
        return pulumi.get(self, "kubeconfigs")


class AwaitableListManagedClusterAdminCredentialsResult(ListManagedClusterAdminCredentialsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListManagedClusterAdminCredentialsResult(
            kubeconfigs=self.kubeconfigs)


def list_managed_cluster_admin_credentials(resource_group_name: Optional[str] = None,
                                           resource_name: Optional[str] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListManagedClusterAdminCredentialsResult:
    """
    Gets cluster admin credential of the managed cluster with a specified resource group and name.


    :param str resource_group_name: The name of the resource group.
    :param str resource_name: The name of the managed cluster resource.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:containerservice/v20190601:listManagedClusterAdminCredentials', __args__, opts=opts, typ=ListManagedClusterAdminCredentialsResult).value

    return AwaitableListManagedClusterAdminCredentialsResult(
        kubeconfigs=pulumi.get(__ret__, 'kubeconfigs'))
def list_managed_cluster_admin_credentials_output(resource_group_name: Optional[pulumi.Input[str]] = None,
                                                  resource_name: Optional[pulumi.Input[str]] = None,
                                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListManagedClusterAdminCredentialsResult]:
    """
    Gets cluster admin credential of the managed cluster with a specified resource group and name.


    :param str resource_group_name: The name of the resource group.
    :param str resource_name: The name of the managed cluster resource.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:containerservice/v20190601:listManagedClusterAdminCredentials', __args__, opts=opts, typ=ListManagedClusterAdminCredentialsResult)
    return __ret__.apply(lambda __response__: ListManagedClusterAdminCredentialsResult(
        kubeconfigs=pulumi.get(__response__, 'kubeconfigs')))
