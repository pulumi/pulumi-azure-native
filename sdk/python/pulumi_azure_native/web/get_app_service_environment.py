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
    'GetAppServiceEnvironmentResult',
    'AwaitableGetAppServiceEnvironmentResult',
    'get_app_service_environment',
    'get_app_service_environment_output',
]

@pulumi.output_type
class GetAppServiceEnvironmentResult:
    """
    App Service Environment ARM resource.
    """
    def __init__(__self__, azure_api_version=None, cluster_settings=None, custom_dns_suffix_configuration=None, dedicated_host_count=None, dns_suffix=None, front_end_scale_factor=None, has_linux_workers=None, id=None, internal_load_balancing_mode=None, ipssl_address_count=None, kind=None, location=None, maximum_number_of_machines=None, multi_role_count=None, multi_size=None, name=None, networking_configuration=None, provisioning_state=None, status=None, suspended=None, tags=None, type=None, upgrade_availability=None, upgrade_preference=None, user_whitelisted_ip_ranges=None, virtual_network=None, zone_redundant=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if cluster_settings and not isinstance(cluster_settings, list):
            raise TypeError("Expected argument 'cluster_settings' to be a list")
        pulumi.set(__self__, "cluster_settings", cluster_settings)
        if custom_dns_suffix_configuration and not isinstance(custom_dns_suffix_configuration, dict):
            raise TypeError("Expected argument 'custom_dns_suffix_configuration' to be a dict")
        pulumi.set(__self__, "custom_dns_suffix_configuration", custom_dns_suffix_configuration)
        if dedicated_host_count and not isinstance(dedicated_host_count, int):
            raise TypeError("Expected argument 'dedicated_host_count' to be a int")
        pulumi.set(__self__, "dedicated_host_count", dedicated_host_count)
        if dns_suffix and not isinstance(dns_suffix, str):
            raise TypeError("Expected argument 'dns_suffix' to be a str")
        pulumi.set(__self__, "dns_suffix", dns_suffix)
        if front_end_scale_factor and not isinstance(front_end_scale_factor, int):
            raise TypeError("Expected argument 'front_end_scale_factor' to be a int")
        pulumi.set(__self__, "front_end_scale_factor", front_end_scale_factor)
        if has_linux_workers and not isinstance(has_linux_workers, bool):
            raise TypeError("Expected argument 'has_linux_workers' to be a bool")
        pulumi.set(__self__, "has_linux_workers", has_linux_workers)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if internal_load_balancing_mode and not isinstance(internal_load_balancing_mode, str):
            raise TypeError("Expected argument 'internal_load_balancing_mode' to be a str")
        pulumi.set(__self__, "internal_load_balancing_mode", internal_load_balancing_mode)
        if ipssl_address_count and not isinstance(ipssl_address_count, int):
            raise TypeError("Expected argument 'ipssl_address_count' to be a int")
        pulumi.set(__self__, "ipssl_address_count", ipssl_address_count)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if maximum_number_of_machines and not isinstance(maximum_number_of_machines, int):
            raise TypeError("Expected argument 'maximum_number_of_machines' to be a int")
        pulumi.set(__self__, "maximum_number_of_machines", maximum_number_of_machines)
        if multi_role_count and not isinstance(multi_role_count, int):
            raise TypeError("Expected argument 'multi_role_count' to be a int")
        pulumi.set(__self__, "multi_role_count", multi_role_count)
        if multi_size and not isinstance(multi_size, str):
            raise TypeError("Expected argument 'multi_size' to be a str")
        pulumi.set(__self__, "multi_size", multi_size)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if networking_configuration and not isinstance(networking_configuration, dict):
            raise TypeError("Expected argument 'networking_configuration' to be a dict")
        pulumi.set(__self__, "networking_configuration", networking_configuration)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if suspended and not isinstance(suspended, bool):
            raise TypeError("Expected argument 'suspended' to be a bool")
        pulumi.set(__self__, "suspended", suspended)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if upgrade_availability and not isinstance(upgrade_availability, str):
            raise TypeError("Expected argument 'upgrade_availability' to be a str")
        pulumi.set(__self__, "upgrade_availability", upgrade_availability)
        if upgrade_preference and not isinstance(upgrade_preference, str):
            raise TypeError("Expected argument 'upgrade_preference' to be a str")
        pulumi.set(__self__, "upgrade_preference", upgrade_preference)
        if user_whitelisted_ip_ranges and not isinstance(user_whitelisted_ip_ranges, list):
            raise TypeError("Expected argument 'user_whitelisted_ip_ranges' to be a list")
        pulumi.set(__self__, "user_whitelisted_ip_ranges", user_whitelisted_ip_ranges)
        if virtual_network and not isinstance(virtual_network, dict):
            raise TypeError("Expected argument 'virtual_network' to be a dict")
        pulumi.set(__self__, "virtual_network", virtual_network)
        if zone_redundant and not isinstance(zone_redundant, bool):
            raise TypeError("Expected argument 'zone_redundant' to be a bool")
        pulumi.set(__self__, "zone_redundant", zone_redundant)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="clusterSettings")
    def cluster_settings(self) -> Optional[Sequence['outputs.NameValuePairResponse']]:
        """
        Custom settings for changing the behavior of the App Service Environment.
        """
        return pulumi.get(self, "cluster_settings")

    @property
    @pulumi.getter(name="customDnsSuffixConfiguration")
    def custom_dns_suffix_configuration(self) -> Optional['outputs.CustomDnsSuffixConfigurationResponse']:
        """
        Full view of the custom domain suffix configuration for ASEv3.
        """
        return pulumi.get(self, "custom_dns_suffix_configuration")

    @property
    @pulumi.getter(name="dedicatedHostCount")
    def dedicated_host_count(self) -> Optional[builtins.int]:
        """
        Dedicated Host Count
        """
        return pulumi.get(self, "dedicated_host_count")

    @property
    @pulumi.getter(name="dnsSuffix")
    def dns_suffix(self) -> Optional[builtins.str]:
        """
        DNS suffix of the App Service Environment.
        """
        return pulumi.get(self, "dns_suffix")

    @property
    @pulumi.getter(name="frontEndScaleFactor")
    def front_end_scale_factor(self) -> Optional[builtins.int]:
        """
        Scale factor for front-ends.
        """
        return pulumi.get(self, "front_end_scale_factor")

    @property
    @pulumi.getter(name="hasLinuxWorkers")
    def has_linux_workers(self) -> builtins.bool:
        """
        Flag that displays whether an ASE has linux workers or not
        """
        return pulumi.get(self, "has_linux_workers")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Resource Id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="internalLoadBalancingMode")
    def internal_load_balancing_mode(self) -> Optional[builtins.str]:
        """
        Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment.
        """
        return pulumi.get(self, "internal_load_balancing_mode")

    @property
    @pulumi.getter(name="ipsslAddressCount")
    def ipssl_address_count(self) -> Optional[builtins.int]:
        """
        Number of IP SSL addresses reserved for the App Service Environment.
        """
        return pulumi.get(self, "ipssl_address_count")

    @property
    @pulumi.getter
    def kind(self) -> Optional[builtins.str]:
        """
        Kind of resource. If the resource is an app, you can refer to https://github.com/Azure/app-service-linux-docs/blob/master/Things_You_Should_Know/kind_property.md#app-service-resource-kind-reference for details supported values for kind.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        Resource Location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maximumNumberOfMachines")
    def maximum_number_of_machines(self) -> builtins.int:
        """
        Maximum number of VMs in the App Service Environment.
        """
        return pulumi.get(self, "maximum_number_of_machines")

    @property
    @pulumi.getter(name="multiRoleCount")
    def multi_role_count(self) -> builtins.int:
        """
        Number of front-end instances.
        """
        return pulumi.get(self, "multi_role_count")

    @property
    @pulumi.getter(name="multiSize")
    def multi_size(self) -> Optional[builtins.str]:
        """
        Front-end VM size, e.g. "Medium", "Large".
        """
        return pulumi.get(self, "multi_size")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Resource Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkingConfiguration")
    def networking_configuration(self) -> Optional['outputs.AseV3NetworkingConfigurationResponse']:
        """
        Full view of networking configuration for an ASE.
        """
        return pulumi.get(self, "networking_configuration")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Provisioning state of the App Service Environment.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        Current status of the App Service Environment.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def suspended(self) -> builtins.bool:
        """
        <code>true</code> if the App Service Environment is suspended; otherwise, <code>false</code>. The environment can be suspended, e.g. when the management endpoint is no longer available
         (most likely because NSG blocked the incoming traffic).
        """
        return pulumi.get(self, "suspended")

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
    @pulumi.getter(name="upgradeAvailability")
    def upgrade_availability(self) -> builtins.str:
        """
        Whether an upgrade is available for this App Service Environment.
        """
        return pulumi.get(self, "upgrade_availability")

    @property
    @pulumi.getter(name="upgradePreference")
    def upgrade_preference(self) -> Optional[builtins.str]:
        """
        Upgrade Preference
        """
        return pulumi.get(self, "upgrade_preference")

    @property
    @pulumi.getter(name="userWhitelistedIpRanges")
    def user_whitelisted_ip_ranges(self) -> Optional[Sequence[builtins.str]]:
        """
        User added ip ranges to whitelist on ASE db
        """
        return pulumi.get(self, "user_whitelisted_ip_ranges")

    @property
    @pulumi.getter(name="virtualNetwork")
    def virtual_network(self) -> 'outputs.VirtualNetworkProfileResponse':
        """
        Description of the Virtual Network.
        """
        return pulumi.get(self, "virtual_network")

    @property
    @pulumi.getter(name="zoneRedundant")
    def zone_redundant(self) -> Optional[builtins.bool]:
        """
        Whether or not this App Service Environment is zone-redundant.
        """
        return pulumi.get(self, "zone_redundant")


class AwaitableGetAppServiceEnvironmentResult(GetAppServiceEnvironmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppServiceEnvironmentResult(
            azure_api_version=self.azure_api_version,
            cluster_settings=self.cluster_settings,
            custom_dns_suffix_configuration=self.custom_dns_suffix_configuration,
            dedicated_host_count=self.dedicated_host_count,
            dns_suffix=self.dns_suffix,
            front_end_scale_factor=self.front_end_scale_factor,
            has_linux_workers=self.has_linux_workers,
            id=self.id,
            internal_load_balancing_mode=self.internal_load_balancing_mode,
            ipssl_address_count=self.ipssl_address_count,
            kind=self.kind,
            location=self.location,
            maximum_number_of_machines=self.maximum_number_of_machines,
            multi_role_count=self.multi_role_count,
            multi_size=self.multi_size,
            name=self.name,
            networking_configuration=self.networking_configuration,
            provisioning_state=self.provisioning_state,
            status=self.status,
            suspended=self.suspended,
            tags=self.tags,
            type=self.type,
            upgrade_availability=self.upgrade_availability,
            upgrade_preference=self.upgrade_preference,
            user_whitelisted_ip_ranges=self.user_whitelisted_ip_ranges,
            virtual_network=self.virtual_network,
            zone_redundant=self.zone_redundant)


def get_app_service_environment(name: Optional[builtins.str] = None,
                                resource_group_name: Optional[builtins.str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppServiceEnvironmentResult:
    """
    Description for Get the properties of an App Service Environment.

    Uses Azure REST API version 2024-04-01.

    Other available API versions: 2016-09-01, 2018-02-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str name: Name of the App Service Environment.
    :param builtins.str resource_group_name: Name of the resource group to which the resource belongs.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:web:getAppServiceEnvironment', __args__, opts=opts, typ=GetAppServiceEnvironmentResult).value

    return AwaitableGetAppServiceEnvironmentResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        cluster_settings=pulumi.get(__ret__, 'cluster_settings'),
        custom_dns_suffix_configuration=pulumi.get(__ret__, 'custom_dns_suffix_configuration'),
        dedicated_host_count=pulumi.get(__ret__, 'dedicated_host_count'),
        dns_suffix=pulumi.get(__ret__, 'dns_suffix'),
        front_end_scale_factor=pulumi.get(__ret__, 'front_end_scale_factor'),
        has_linux_workers=pulumi.get(__ret__, 'has_linux_workers'),
        id=pulumi.get(__ret__, 'id'),
        internal_load_balancing_mode=pulumi.get(__ret__, 'internal_load_balancing_mode'),
        ipssl_address_count=pulumi.get(__ret__, 'ipssl_address_count'),
        kind=pulumi.get(__ret__, 'kind'),
        location=pulumi.get(__ret__, 'location'),
        maximum_number_of_machines=pulumi.get(__ret__, 'maximum_number_of_machines'),
        multi_role_count=pulumi.get(__ret__, 'multi_role_count'),
        multi_size=pulumi.get(__ret__, 'multi_size'),
        name=pulumi.get(__ret__, 'name'),
        networking_configuration=pulumi.get(__ret__, 'networking_configuration'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        status=pulumi.get(__ret__, 'status'),
        suspended=pulumi.get(__ret__, 'suspended'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        upgrade_availability=pulumi.get(__ret__, 'upgrade_availability'),
        upgrade_preference=pulumi.get(__ret__, 'upgrade_preference'),
        user_whitelisted_ip_ranges=pulumi.get(__ret__, 'user_whitelisted_ip_ranges'),
        virtual_network=pulumi.get(__ret__, 'virtual_network'),
        zone_redundant=pulumi.get(__ret__, 'zone_redundant'))
def get_app_service_environment_output(name: Optional[pulumi.Input[builtins.str]] = None,
                                       resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAppServiceEnvironmentResult]:
    """
    Description for Get the properties of an App Service Environment.

    Uses Azure REST API version 2024-04-01.

    Other available API versions: 2016-09-01, 2018-02-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str name: Name of the App Service Environment.
    :param builtins.str resource_group_name: Name of the resource group to which the resource belongs.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:web:getAppServiceEnvironment', __args__, opts=opts, typ=GetAppServiceEnvironmentResult)
    return __ret__.apply(lambda __response__: GetAppServiceEnvironmentResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        cluster_settings=pulumi.get(__response__, 'cluster_settings'),
        custom_dns_suffix_configuration=pulumi.get(__response__, 'custom_dns_suffix_configuration'),
        dedicated_host_count=pulumi.get(__response__, 'dedicated_host_count'),
        dns_suffix=pulumi.get(__response__, 'dns_suffix'),
        front_end_scale_factor=pulumi.get(__response__, 'front_end_scale_factor'),
        has_linux_workers=pulumi.get(__response__, 'has_linux_workers'),
        id=pulumi.get(__response__, 'id'),
        internal_load_balancing_mode=pulumi.get(__response__, 'internal_load_balancing_mode'),
        ipssl_address_count=pulumi.get(__response__, 'ipssl_address_count'),
        kind=pulumi.get(__response__, 'kind'),
        location=pulumi.get(__response__, 'location'),
        maximum_number_of_machines=pulumi.get(__response__, 'maximum_number_of_machines'),
        multi_role_count=pulumi.get(__response__, 'multi_role_count'),
        multi_size=pulumi.get(__response__, 'multi_size'),
        name=pulumi.get(__response__, 'name'),
        networking_configuration=pulumi.get(__response__, 'networking_configuration'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        status=pulumi.get(__response__, 'status'),
        suspended=pulumi.get(__response__, 'suspended'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        upgrade_availability=pulumi.get(__response__, 'upgrade_availability'),
        upgrade_preference=pulumi.get(__response__, 'upgrade_preference'),
        user_whitelisted_ip_ranges=pulumi.get(__response__, 'user_whitelisted_ip_ranges'),
        virtual_network=pulumi.get(__response__, 'virtual_network'),
        zone_redundant=pulumi.get(__response__, 'zone_redundant')))
