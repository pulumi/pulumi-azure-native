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

__all__ = [
    'ListPrivateCloudAdminCredentialsResult',
    'AwaitableListPrivateCloudAdminCredentialsResult',
    'list_private_cloud_admin_credentials',
    'list_private_cloud_admin_credentials_output',
]

@pulumi.output_type
class ListPrivateCloudAdminCredentialsResult:
    """
    Administrative credentials for accessing vCenter and NSX-T
    """
    def __init__(__self__, nsxt_password=None, nsxt_username=None, vcenter_password=None, vcenter_username=None):
        if nsxt_password and not isinstance(nsxt_password, str):
            raise TypeError("Expected argument 'nsxt_password' to be a str")
        pulumi.set(__self__, "nsxt_password", nsxt_password)
        if nsxt_username and not isinstance(nsxt_username, str):
            raise TypeError("Expected argument 'nsxt_username' to be a str")
        pulumi.set(__self__, "nsxt_username", nsxt_username)
        if vcenter_password and not isinstance(vcenter_password, str):
            raise TypeError("Expected argument 'vcenter_password' to be a str")
        pulumi.set(__self__, "vcenter_password", vcenter_password)
        if vcenter_username and not isinstance(vcenter_username, str):
            raise TypeError("Expected argument 'vcenter_username' to be a str")
        pulumi.set(__self__, "vcenter_username", vcenter_username)

    @property
    @pulumi.getter(name="nsxtPassword")
    def nsxt_password(self) -> builtins.str:
        """
        NSX-T Manager password
        """
        return pulumi.get(self, "nsxt_password")

    @property
    @pulumi.getter(name="nsxtUsername")
    def nsxt_username(self) -> builtins.str:
        """
        NSX-T Manager username
        """
        return pulumi.get(self, "nsxt_username")

    @property
    @pulumi.getter(name="vcenterPassword")
    def vcenter_password(self) -> builtins.str:
        """
        vCenter admin password
        """
        return pulumi.get(self, "vcenter_password")

    @property
    @pulumi.getter(name="vcenterUsername")
    def vcenter_username(self) -> builtins.str:
        """
        vCenter admin username
        """
        return pulumi.get(self, "vcenter_username")


class AwaitableListPrivateCloudAdminCredentialsResult(ListPrivateCloudAdminCredentialsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListPrivateCloudAdminCredentialsResult(
            nsxt_password=self.nsxt_password,
            nsxt_username=self.nsxt_username,
            vcenter_password=self.vcenter_password,
            vcenter_username=self.vcenter_username)


def list_private_cloud_admin_credentials(private_cloud_name: Optional[builtins.str] = None,
                                         resource_group_name: Optional[builtins.str] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListPrivateCloudAdminCredentialsResult:
    """
    List the admin credentials for the private cloud

    Uses Azure REST API version 2023-09-01.

    Other available API versions: 2022-05-01, 2023-03-01, 2024-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native avs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str private_cloud_name: Name of the private cloud
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['privateCloudName'] = private_cloud_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:avs:listPrivateCloudAdminCredentials', __args__, opts=opts, typ=ListPrivateCloudAdminCredentialsResult).value

    return AwaitableListPrivateCloudAdminCredentialsResult(
        nsxt_password=pulumi.get(__ret__, 'nsxt_password'),
        nsxt_username=pulumi.get(__ret__, 'nsxt_username'),
        vcenter_password=pulumi.get(__ret__, 'vcenter_password'),
        vcenter_username=pulumi.get(__ret__, 'vcenter_username'))
def list_private_cloud_admin_credentials_output(private_cloud_name: Optional[pulumi.Input[builtins.str]] = None,
                                                resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListPrivateCloudAdminCredentialsResult]:
    """
    List the admin credentials for the private cloud

    Uses Azure REST API version 2023-09-01.

    Other available API versions: 2022-05-01, 2023-03-01, 2024-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native avs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str private_cloud_name: Name of the private cloud
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['privateCloudName'] = private_cloud_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:avs:listPrivateCloudAdminCredentials', __args__, opts=opts, typ=ListPrivateCloudAdminCredentialsResult)
    return __ret__.apply(lambda __response__: ListPrivateCloudAdminCredentialsResult(
        nsxt_password=pulumi.get(__response__, 'nsxt_password'),
        nsxt_username=pulumi.get(__response__, 'nsxt_username'),
        vcenter_password=pulumi.get(__response__, 'vcenter_password'),
        vcenter_username=pulumi.get(__response__, 'vcenter_username')))
