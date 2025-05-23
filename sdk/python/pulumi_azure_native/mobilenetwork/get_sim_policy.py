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
    'GetSimPolicyResult',
    'AwaitableGetSimPolicyResult',
    'get_sim_policy',
    'get_sim_policy_output',
]

@pulumi.output_type
class GetSimPolicyResult:
    """
    SIM policy resource.
    """
    def __init__(__self__, azure_api_version=None, default_slice=None, id=None, location=None, name=None, provisioning_state=None, registration_timer=None, rfsp_index=None, site_provisioning_state=None, slice_configurations=None, system_data=None, tags=None, type=None, ue_ambr=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if default_slice and not isinstance(default_slice, dict):
            raise TypeError("Expected argument 'default_slice' to be a dict")
        pulumi.set(__self__, "default_slice", default_slice)
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
        if registration_timer and not isinstance(registration_timer, int):
            raise TypeError("Expected argument 'registration_timer' to be a int")
        pulumi.set(__self__, "registration_timer", registration_timer)
        if rfsp_index and not isinstance(rfsp_index, int):
            raise TypeError("Expected argument 'rfsp_index' to be a int")
        pulumi.set(__self__, "rfsp_index", rfsp_index)
        if site_provisioning_state and not isinstance(site_provisioning_state, dict):
            raise TypeError("Expected argument 'site_provisioning_state' to be a dict")
        pulumi.set(__self__, "site_provisioning_state", site_provisioning_state)
        if slice_configurations and not isinstance(slice_configurations, list):
            raise TypeError("Expected argument 'slice_configurations' to be a list")
        pulumi.set(__self__, "slice_configurations", slice_configurations)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if ue_ambr and not isinstance(ue_ambr, dict):
            raise TypeError("Expected argument 'ue_ambr' to be a dict")
        pulumi.set(__self__, "ue_ambr", ue_ambr)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="defaultSlice")
    def default_slice(self) -> 'outputs.SliceResourceIdResponse':
        """
        The default slice to use if the UE does not explicitly specify it. This slice must exist in the `sliceConfigurations` map. The slice must be in the same location as the SIM policy.
        """
        return pulumi.get(self, "default_slice")

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
        The provisioning state of the SIM policy resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="registrationTimer")
    def registration_timer(self) -> Optional[builtins.int]:
        """
        UE periodic registration update timer (5G) or UE periodic tracking area update timer (4G), in seconds.
        """
        return pulumi.get(self, "registration_timer")

    @property
    @pulumi.getter(name="rfspIndex")
    def rfsp_index(self) -> Optional[builtins.int]:
        """
        RAT/Frequency Selection Priority Index, defined in 3GPP TS 36.413. This is an optional setting and by default is unspecified.
        """
        return pulumi.get(self, "rfsp_index")

    @property
    @pulumi.getter(name="siteProvisioningState")
    def site_provisioning_state(self) -> Mapping[str, builtins.str]:
        """
        A dictionary of sites to the provisioning state of this SIM policy on that site.
        """
        return pulumi.get(self, "site_provisioning_state")

    @property
    @pulumi.getter(name="sliceConfigurations")
    def slice_configurations(self) -> Sequence['outputs.SliceConfigurationResponse']:
        """
        The allowed slices and the settings to use for them. The list must not contain duplicate items and must contain at least one item.
        """
        return pulumi.get(self, "slice_configurations")

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

    @property
    @pulumi.getter(name="ueAmbr")
    def ue_ambr(self) -> 'outputs.AmbrResponse':
        """
        Aggregate maximum bit rate across all non-GBR QoS flows of all PDU sessions of a given UE. See 3GPP TS23.501 section 5.7.2.6 for a full description of the UE-AMBR.
        """
        return pulumi.get(self, "ue_ambr")


class AwaitableGetSimPolicyResult(GetSimPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSimPolicyResult(
            azure_api_version=self.azure_api_version,
            default_slice=self.default_slice,
            id=self.id,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            registration_timer=self.registration_timer,
            rfsp_index=self.rfsp_index,
            site_provisioning_state=self.site_provisioning_state,
            slice_configurations=self.slice_configurations,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type,
            ue_ambr=self.ue_ambr)


def get_sim_policy(mobile_network_name: Optional[builtins.str] = None,
                   resource_group_name: Optional[builtins.str] = None,
                   sim_policy_name: Optional[builtins.str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSimPolicyResult:
    """
    Gets information about the specified SIM policy.

    Uses Azure REST API version 2024-04-01.

    Other available API versions: 2022-04-01-preview, 2022-11-01, 2023-06-01, 2023-09-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mobilenetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str mobile_network_name: The name of the mobile network.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str sim_policy_name: The name of the SIM policy.
    """
    __args__ = dict()
    __args__['mobileNetworkName'] = mobile_network_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['simPolicyName'] = sim_policy_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:mobilenetwork:getSimPolicy', __args__, opts=opts, typ=GetSimPolicyResult).value

    return AwaitableGetSimPolicyResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        default_slice=pulumi.get(__ret__, 'default_slice'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        registration_timer=pulumi.get(__ret__, 'registration_timer'),
        rfsp_index=pulumi.get(__ret__, 'rfsp_index'),
        site_provisioning_state=pulumi.get(__ret__, 'site_provisioning_state'),
        slice_configurations=pulumi.get(__ret__, 'slice_configurations'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        ue_ambr=pulumi.get(__ret__, 'ue_ambr'))
def get_sim_policy_output(mobile_network_name: Optional[pulumi.Input[builtins.str]] = None,
                          resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                          sim_policy_name: Optional[pulumi.Input[builtins.str]] = None,
                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSimPolicyResult]:
    """
    Gets information about the specified SIM policy.

    Uses Azure REST API version 2024-04-01.

    Other available API versions: 2022-04-01-preview, 2022-11-01, 2023-06-01, 2023-09-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mobilenetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str mobile_network_name: The name of the mobile network.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str sim_policy_name: The name of the SIM policy.
    """
    __args__ = dict()
    __args__['mobileNetworkName'] = mobile_network_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['simPolicyName'] = sim_policy_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:mobilenetwork:getSimPolicy', __args__, opts=opts, typ=GetSimPolicyResult)
    return __ret__.apply(lambda __response__: GetSimPolicyResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        default_slice=pulumi.get(__response__, 'default_slice'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        registration_timer=pulumi.get(__response__, 'registration_timer'),
        rfsp_index=pulumi.get(__response__, 'rfsp_index'),
        site_provisioning_state=pulumi.get(__response__, 'site_provisioning_state'),
        slice_configurations=pulumi.get(__response__, 'slice_configurations'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        ue_ambr=pulumi.get(__response__, 'ue_ambr')))
