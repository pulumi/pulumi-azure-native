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
from ._enums import *
from ._inputs import *

__all__ = ['ProfileArgs', 'Profile']

@pulumi.input_type
class ProfileArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 allowed_endpoint_record_types: Optional[pulumi.Input[Sequence[pulumi.Input[Union[builtins.str, 'AllowedEndpointRecordType']]]]] = None,
                 dns_config: Optional[pulumi.Input['DnsConfigArgs']] = None,
                 endpoints: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointArgs']]]] = None,
                 id: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 max_return: Optional[pulumi.Input[builtins.float]] = None,
                 monitor_config: Optional[pulumi.Input['MonitorConfigArgs']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 profile_name: Optional[pulumi.Input[builtins.str]] = None,
                 profile_status: Optional[pulumi.Input[Union[builtins.str, 'ProfileStatus']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 traffic_routing_method: Optional[pulumi.Input[Union[builtins.str, 'TrafficRoutingMethod']]] = None,
                 traffic_view_enrollment_status: Optional[pulumi.Input[Union[builtins.str, 'TrafficViewEnrollmentStatus']]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Profile resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Sequence[pulumi.Input[Union[builtins.str, 'AllowedEndpointRecordType']]]] allowed_endpoint_record_types: The list of allowed endpoint record types.
        :param pulumi.Input['DnsConfigArgs'] dns_config: The DNS settings of the Traffic Manager profile.
        :param pulumi.Input[Sequence[pulumi.Input['EndpointArgs']]] endpoints: The list of endpoints in the Traffic Manager profile.
               These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
        :param pulumi.Input[builtins.str] id: Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
        :param pulumi.Input[builtins.str] location: The Azure Region where the resource lives
        :param pulumi.Input[builtins.float] max_return: Maximum number of endpoints to be returned for MultiValue routing type.
        :param pulumi.Input['MonitorConfigArgs'] monitor_config: The endpoint monitoring settings of the Traffic Manager profile.
        :param pulumi.Input[builtins.str] name: The name of the resource
        :param pulumi.Input[builtins.str] profile_name: The name of the Traffic Manager profile.
        :param pulumi.Input[Union[builtins.str, 'ProfileStatus']] profile_status: The status of the Traffic Manager profile.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[Union[builtins.str, 'TrafficRoutingMethod']] traffic_routing_method: The traffic routing method of the Traffic Manager profile.
        :param pulumi.Input[Union[builtins.str, 'TrafficViewEnrollmentStatus']] traffic_view_enrollment_status: Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
        :param pulumi.Input[builtins.str] type: The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if allowed_endpoint_record_types is not None:
            pulumi.set(__self__, "allowed_endpoint_record_types", allowed_endpoint_record_types)
        if dns_config is not None:
            pulumi.set(__self__, "dns_config", dns_config)
        if endpoints is not None:
            pulumi.set(__self__, "endpoints", endpoints)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if max_return is not None:
            pulumi.set(__self__, "max_return", max_return)
        if monitor_config is not None:
            pulumi.set(__self__, "monitor_config", monitor_config)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if profile_name is not None:
            pulumi.set(__self__, "profile_name", profile_name)
        if profile_status is not None:
            pulumi.set(__self__, "profile_status", profile_status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if traffic_routing_method is not None:
            pulumi.set(__self__, "traffic_routing_method", traffic_routing_method)
        if traffic_view_enrollment_status is not None:
            pulumi.set(__self__, "traffic_view_enrollment_status", traffic_view_enrollment_status)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group. The name is case insensitive.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="allowedEndpointRecordTypes")
    def allowed_endpoint_record_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[Union[builtins.str, 'AllowedEndpointRecordType']]]]]:
        """
        The list of allowed endpoint record types.
        """
        return pulumi.get(self, "allowed_endpoint_record_types")

    @allowed_endpoint_record_types.setter
    def allowed_endpoint_record_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[Union[builtins.str, 'AllowedEndpointRecordType']]]]]):
        pulumi.set(self, "allowed_endpoint_record_types", value)

    @property
    @pulumi.getter(name="dnsConfig")
    def dns_config(self) -> Optional[pulumi.Input['DnsConfigArgs']]:
        """
        The DNS settings of the Traffic Manager profile.
        """
        return pulumi.get(self, "dns_config")

    @dns_config.setter
    def dns_config(self, value: Optional[pulumi.Input['DnsConfigArgs']]):
        pulumi.set(self, "dns_config", value)

    @property
    @pulumi.getter
    def endpoints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EndpointArgs']]]]:
        """
        The list of endpoints in the Traffic Manager profile.
        These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
        """
        return pulumi.get(self, "endpoints")

    @endpoints.setter
    def endpoints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointArgs']]]]):
        pulumi.set(self, "endpoints", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Azure Region where the resource lives
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="maxReturn")
    def max_return(self) -> Optional[pulumi.Input[builtins.float]]:
        """
        Maximum number of endpoints to be returned for MultiValue routing type.
        """
        return pulumi.get(self, "max_return")

    @max_return.setter
    def max_return(self, value: Optional[pulumi.Input[builtins.float]]):
        pulumi.set(self, "max_return", value)

    @property
    @pulumi.getter(name="monitorConfig")
    def monitor_config(self) -> Optional[pulumi.Input['MonitorConfigArgs']]:
        """
        The endpoint monitoring settings of the Traffic Manager profile.
        """
        return pulumi.get(self, "monitor_config")

    @monitor_config.setter
    def monitor_config(self, value: Optional[pulumi.Input['MonitorConfigArgs']]):
        pulumi.set(self, "monitor_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="profileName")
    def profile_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the Traffic Manager profile.
        """
        return pulumi.get(self, "profile_name")

    @profile_name.setter
    def profile_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "profile_name", value)

    @property
    @pulumi.getter(name="profileStatus")
    def profile_status(self) -> Optional[pulumi.Input[Union[builtins.str, 'ProfileStatus']]]:
        """
        The status of the Traffic Manager profile.
        """
        return pulumi.get(self, "profile_status")

    @profile_status.setter
    def profile_status(self, value: Optional[pulumi.Input[Union[builtins.str, 'ProfileStatus']]]):
        pulumi.set(self, "profile_status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="trafficRoutingMethod")
    def traffic_routing_method(self) -> Optional[pulumi.Input[Union[builtins.str, 'TrafficRoutingMethod']]]:
        """
        The traffic routing method of the Traffic Manager profile.
        """
        return pulumi.get(self, "traffic_routing_method")

    @traffic_routing_method.setter
    def traffic_routing_method(self, value: Optional[pulumi.Input[Union[builtins.str, 'TrafficRoutingMethod']]]):
        pulumi.set(self, "traffic_routing_method", value)

    @property
    @pulumi.getter(name="trafficViewEnrollmentStatus")
    def traffic_view_enrollment_status(self) -> Optional[pulumi.Input[Union[builtins.str, 'TrafficViewEnrollmentStatus']]]:
        """
        Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
        """
        return pulumi.get(self, "traffic_view_enrollment_status")

    @traffic_view_enrollment_status.setter
    def traffic_view_enrollment_status(self, value: Optional[pulumi.Input[Union[builtins.str, 'TrafficViewEnrollmentStatus']]]):
        pulumi.set(self, "traffic_view_enrollment_status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)


@pulumi.type_token("azure-native:trafficmanager:Profile")
class Profile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_endpoint_record_types: Optional[pulumi.Input[Sequence[pulumi.Input[Union[builtins.str, 'AllowedEndpointRecordType']]]]] = None,
                 dns_config: Optional[pulumi.Input[Union['DnsConfigArgs', 'DnsConfigArgsDict']]] = None,
                 endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EndpointArgs', 'EndpointArgsDict']]]]] = None,
                 id: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 max_return: Optional[pulumi.Input[builtins.float]] = None,
                 monitor_config: Optional[pulumi.Input[Union['MonitorConfigArgs', 'MonitorConfigArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 profile_name: Optional[pulumi.Input[builtins.str]] = None,
                 profile_status: Optional[pulumi.Input[Union[builtins.str, 'ProfileStatus']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 traffic_routing_method: Optional[pulumi.Input[Union[builtins.str, 'TrafficRoutingMethod']]] = None,
                 traffic_view_enrollment_status: Optional[pulumi.Input[Union[builtins.str, 'TrafficViewEnrollmentStatus']]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Class representing a Traffic Manager profile.

        Uses Azure REST API version 2022-04-01.

        Other available API versions: 2015-11-01, 2017-03-01, 2017-05-01, 2018-02-01, 2018-03-01, 2018-04-01, 2018-08-01, 2022-04-01-preview, 2024-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native trafficmanager [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union[builtins.str, 'AllowedEndpointRecordType']]]] allowed_endpoint_record_types: The list of allowed endpoint record types.
        :param pulumi.Input[Union['DnsConfigArgs', 'DnsConfigArgsDict']] dns_config: The DNS settings of the Traffic Manager profile.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EndpointArgs', 'EndpointArgsDict']]]] endpoints: The list of endpoints in the Traffic Manager profile.
               These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
        :param pulumi.Input[builtins.str] id: Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
        :param pulumi.Input[builtins.str] location: The Azure Region where the resource lives
        :param pulumi.Input[builtins.float] max_return: Maximum number of endpoints to be returned for MultiValue routing type.
        :param pulumi.Input[Union['MonitorConfigArgs', 'MonitorConfigArgsDict']] monitor_config: The endpoint monitoring settings of the Traffic Manager profile.
        :param pulumi.Input[builtins.str] name: The name of the resource
        :param pulumi.Input[builtins.str] profile_name: The name of the Traffic Manager profile.
        :param pulumi.Input[Union[builtins.str, 'ProfileStatus']] profile_status: The status of the Traffic Manager profile.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[Union[builtins.str, 'TrafficRoutingMethod']] traffic_routing_method: The traffic routing method of the Traffic Manager profile.
        :param pulumi.Input[Union[builtins.str, 'TrafficViewEnrollmentStatus']] traffic_view_enrollment_status: Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
        :param pulumi.Input[builtins.str] type: The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Class representing a Traffic Manager profile.

        Uses Azure REST API version 2022-04-01.

        Other available API versions: 2015-11-01, 2017-03-01, 2017-05-01, 2018-02-01, 2018-03-01, 2018-04-01, 2018-08-01, 2022-04-01-preview, 2024-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native trafficmanager [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param ProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_endpoint_record_types: Optional[pulumi.Input[Sequence[pulumi.Input[Union[builtins.str, 'AllowedEndpointRecordType']]]]] = None,
                 dns_config: Optional[pulumi.Input[Union['DnsConfigArgs', 'DnsConfigArgsDict']]] = None,
                 endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EndpointArgs', 'EndpointArgsDict']]]]] = None,
                 id: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 max_return: Optional[pulumi.Input[builtins.float]] = None,
                 monitor_config: Optional[pulumi.Input[Union['MonitorConfigArgs', 'MonitorConfigArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 profile_name: Optional[pulumi.Input[builtins.str]] = None,
                 profile_status: Optional[pulumi.Input[Union[builtins.str, 'ProfileStatus']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 traffic_routing_method: Optional[pulumi.Input[Union[builtins.str, 'TrafficRoutingMethod']]] = None,
                 traffic_view_enrollment_status: Optional[pulumi.Input[Union[builtins.str, 'TrafficViewEnrollmentStatus']]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProfileArgs.__new__(ProfileArgs)

            __props__.__dict__["allowed_endpoint_record_types"] = allowed_endpoint_record_types
            __props__.__dict__["dns_config"] = dns_config
            __props__.__dict__["endpoints"] = endpoints
            __props__.__dict__["id"] = id
            __props__.__dict__["location"] = location
            __props__.__dict__["max_return"] = max_return
            __props__.__dict__["monitor_config"] = monitor_config
            __props__.__dict__["name"] = name
            __props__.__dict__["profile_name"] = profile_name
            __props__.__dict__["profile_status"] = profile_status
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["traffic_routing_method"] = traffic_routing_method
            __props__.__dict__["traffic_view_enrollment_status"] = traffic_view_enrollment_status
            __props__.__dict__["type"] = type
            __props__.__dict__["azure_api_version"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:network/v20220401:Profile"), pulumi.Alias(type_="azure-native:network/v20220401preview:Profile"), pulumi.Alias(type_="azure-native:network:Profile"), pulumi.Alias(type_="azure-native:trafficmanager/v20151101:Profile"), pulumi.Alias(type_="azure-native:trafficmanager/v20170301:Profile"), pulumi.Alias(type_="azure-native:trafficmanager/v20170501:Profile"), pulumi.Alias(type_="azure-native:trafficmanager/v20180201:Profile"), pulumi.Alias(type_="azure-native:trafficmanager/v20180301:Profile"), pulumi.Alias(type_="azure-native:trafficmanager/v20180401:Profile"), pulumi.Alias(type_="azure-native:trafficmanager/v20180801:Profile"), pulumi.Alias(type_="azure-native:trafficmanager/v20220401:Profile"), pulumi.Alias(type_="azure-native:trafficmanager/v20220401preview:Profile"), pulumi.Alias(type_="azure-native:trafficmanager/v20240401preview:Profile")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Profile, __self__).__init__(
            'azure-native:trafficmanager:Profile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Profile':
        """
        Get an existing Profile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ProfileArgs.__new__(ProfileArgs)

        __props__.__dict__["allowed_endpoint_record_types"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["dns_config"] = None
        __props__.__dict__["endpoints"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["max_return"] = None
        __props__.__dict__["monitor_config"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["profile_status"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["traffic_routing_method"] = None
        __props__.__dict__["traffic_view_enrollment_status"] = None
        __props__.__dict__["type"] = None
        return Profile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedEndpointRecordTypes")
    def allowed_endpoint_record_types(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        The list of allowed endpoint record types.
        """
        return pulumi.get(self, "allowed_endpoint_record_types")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="dnsConfig")
    def dns_config(self) -> pulumi.Output[Optional['outputs.DnsConfigResponse']]:
        """
        The DNS settings of the Traffic Manager profile.
        """
        return pulumi.get(self, "dns_config")

    @property
    @pulumi.getter
    def endpoints(self) -> pulumi.Output[Optional[Sequence['outputs.EndpointResponse']]]:
        """
        The list of endpoints in the Traffic Manager profile.
        """
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The Azure Region where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maxReturn")
    def max_return(self) -> pulumi.Output[Optional[builtins.float]]:
        """
        Maximum number of endpoints to be returned for MultiValue routing type.
        """
        return pulumi.get(self, "max_return")

    @property
    @pulumi.getter(name="monitorConfig")
    def monitor_config(self) -> pulumi.Output[Optional['outputs.MonitorConfigResponse']]:
        """
        The endpoint monitoring settings of the Traffic Manager profile.
        """
        return pulumi.get(self, "monitor_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="profileStatus")
    def profile_status(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The status of the Traffic Manager profile.
        """
        return pulumi.get(self, "profile_status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trafficRoutingMethod")
    def traffic_routing_method(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The traffic routing method of the Traffic Manager profile.
        """
        return pulumi.get(self, "traffic_routing_method")

    @property
    @pulumi.getter(name="trafficViewEnrollmentStatus")
    def traffic_view_enrollment_status(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
        """
        return pulumi.get(self, "traffic_view_enrollment_status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
        """
        return pulumi.get(self, "type")

