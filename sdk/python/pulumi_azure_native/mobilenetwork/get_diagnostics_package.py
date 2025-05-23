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
    'GetDiagnosticsPackageResult',
    'AwaitableGetDiagnosticsPackageResult',
    'get_diagnostics_package',
    'get_diagnostics_package_output',
]

@pulumi.output_type
class GetDiagnosticsPackageResult:
    """
    Diagnostics package resource.
    """
    def __init__(__self__, azure_api_version=None, id=None, name=None, provisioning_state=None, reason=None, status=None, system_data=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if reason and not isinstance(reason, str):
            raise TypeError("Expected argument 'reason' to be a str")
        pulumi.set(__self__, "reason", reason)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
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
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

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
        The provisioning state of the diagnostics package resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def reason(self) -> builtins.str:
        """
        The reason for the current state of the diagnostics package collection.
        """
        return pulumi.get(self, "reason")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        The status of the diagnostics package collection.
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
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetDiagnosticsPackageResult(GetDiagnosticsPackageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDiagnosticsPackageResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            name=self.name,
            provisioning_state=self.provisioning_state,
            reason=self.reason,
            status=self.status,
            system_data=self.system_data,
            type=self.type)


def get_diagnostics_package(diagnostics_package_name: Optional[builtins.str] = None,
                            packet_core_control_plane_name: Optional[builtins.str] = None,
                            resource_group_name: Optional[builtins.str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDiagnosticsPackageResult:
    """
    Gets information about the specified diagnostics package.

    Uses Azure REST API version 2024-04-01.

    Other available API versions: 2023-06-01, 2023-09-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mobilenetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str diagnostics_package_name: The name of the diagnostics package.
    :param builtins.str packet_core_control_plane_name: The name of the packet core control plane.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['diagnosticsPackageName'] = diagnostics_package_name
    __args__['packetCoreControlPlaneName'] = packet_core_control_plane_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:mobilenetwork:getDiagnosticsPackage', __args__, opts=opts, typ=GetDiagnosticsPackageResult).value

    return AwaitableGetDiagnosticsPackageResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        reason=pulumi.get(__ret__, 'reason'),
        status=pulumi.get(__ret__, 'status'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_diagnostics_package_output(diagnostics_package_name: Optional[pulumi.Input[builtins.str]] = None,
                                   packet_core_control_plane_name: Optional[pulumi.Input[builtins.str]] = None,
                                   resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDiagnosticsPackageResult]:
    """
    Gets information about the specified diagnostics package.

    Uses Azure REST API version 2024-04-01.

    Other available API versions: 2023-06-01, 2023-09-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mobilenetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str diagnostics_package_name: The name of the diagnostics package.
    :param builtins.str packet_core_control_plane_name: The name of the packet core control plane.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['diagnosticsPackageName'] = diagnostics_package_name
    __args__['packetCoreControlPlaneName'] = packet_core_control_plane_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:mobilenetwork:getDiagnosticsPackage', __args__, opts=opts, typ=GetDiagnosticsPackageResult)
    return __ret__.apply(lambda __response__: GetDiagnosticsPackageResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        reason=pulumi.get(__response__, 'reason'),
        status=pulumi.get(__response__, 'status'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
