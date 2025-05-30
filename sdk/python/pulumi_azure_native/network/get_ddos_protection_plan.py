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
    'GetDdosProtectionPlanResult',
    'AwaitableGetDdosProtectionPlanResult',
    'get_ddos_protection_plan',
    'get_ddos_protection_plan_output',
]

@pulumi.output_type
class GetDdosProtectionPlanResult:
    """
    A DDoS protection plan in a resource group.
    """
    def __init__(__self__, azure_api_version=None, etag=None, id=None, location=None, name=None, provisioning_state=None, public_ip_addresses=None, resource_guid=None, tags=None, type=None, virtual_networks=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
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
        if public_ip_addresses and not isinstance(public_ip_addresses, list):
            raise TypeError("Expected argument 'public_ip_addresses' to be a list")
        pulumi.set(__self__, "public_ip_addresses", public_ip_addresses)
        if resource_guid and not isinstance(resource_guid, str):
            raise TypeError("Expected argument 'resource_guid' to be a str")
        pulumi.set(__self__, "resource_guid", resource_guid)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if virtual_networks and not isinstance(virtual_networks, list):
            raise TypeError("Expected argument 'virtual_networks' to be a list")
        pulumi.set(__self__, "virtual_networks", virtual_networks)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def etag(self) -> builtins.str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> Optional[builtins.str]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        The provisioning state of the DDoS protection plan resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicIPAddresses")
    def public_ip_addresses(self) -> Sequence['outputs.SubResourceResponse']:
        """
        The list of public IPs associated with the DDoS protection plan resource. This list is read-only.
        """
        return pulumi.get(self, "public_ip_addresses")

    @property
    @pulumi.getter(name="resourceGuid")
    def resource_guid(self) -> builtins.str:
        """
        The resource GUID property of the DDoS protection plan resource. It uniquely identifies the resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
        """
        return pulumi.get(self, "resource_guid")

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
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualNetworks")
    def virtual_networks(self) -> Sequence['outputs.SubResourceResponse']:
        """
        The list of virtual networks associated with the DDoS protection plan resource. This list is read-only.
        """
        return pulumi.get(self, "virtual_networks")


class AwaitableGetDdosProtectionPlanResult(GetDdosProtectionPlanResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDdosProtectionPlanResult(
            azure_api_version=self.azure_api_version,
            etag=self.etag,
            id=self.id,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            public_ip_addresses=self.public_ip_addresses,
            resource_guid=self.resource_guid,
            tags=self.tags,
            type=self.type,
            virtual_networks=self.virtual_networks)


def get_ddos_protection_plan(ddos_protection_plan_name: Optional[builtins.str] = None,
                             resource_group_name: Optional[builtins.str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDdosProtectionPlanResult:
    """
    Gets information about the specified DDoS protection plan.

    Uses Azure REST API version 2024-05-01.

    Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str ddos_protection_plan_name: The name of the DDoS protection plan.
    :param builtins.str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['ddosProtectionPlanName'] = ddos_protection_plan_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:network:getDdosProtectionPlan', __args__, opts=opts, typ=GetDdosProtectionPlanResult).value

    return AwaitableGetDdosProtectionPlanResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        etag=pulumi.get(__ret__, 'etag'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        public_ip_addresses=pulumi.get(__ret__, 'public_ip_addresses'),
        resource_guid=pulumi.get(__ret__, 'resource_guid'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        virtual_networks=pulumi.get(__ret__, 'virtual_networks'))
def get_ddos_protection_plan_output(ddos_protection_plan_name: Optional[pulumi.Input[builtins.str]] = None,
                                    resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDdosProtectionPlanResult]:
    """
    Gets information about the specified DDoS protection plan.

    Uses Azure REST API version 2024-05-01.

    Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str ddos_protection_plan_name: The name of the DDoS protection plan.
    :param builtins.str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['ddosProtectionPlanName'] = ddos_protection_plan_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:network:getDdosProtectionPlan', __args__, opts=opts, typ=GetDdosProtectionPlanResult)
    return __ret__.apply(lambda __response__: GetDdosProtectionPlanResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        etag=pulumi.get(__response__, 'etag'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        public_ip_addresses=pulumi.get(__response__, 'public_ip_addresses'),
        resource_guid=pulumi.get(__response__, 'resource_guid'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        virtual_networks=pulumi.get(__response__, 'virtual_networks')))
