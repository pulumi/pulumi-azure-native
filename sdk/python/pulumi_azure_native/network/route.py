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
from ._enums import *

__all__ = ['RouteInitArgs', 'Route']

@pulumi.input_type
class RouteInitArgs:
    def __init__(__self__, *,
                 next_hop_type: pulumi.Input[Union[builtins.str, 'RouteNextHopType']],
                 resource_group_name: pulumi.Input[builtins.str],
                 route_table_name: pulumi.Input[builtins.str],
                 address_prefix: Optional[pulumi.Input[builtins.str]] = None,
                 id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 next_hop_ip_address: Optional[pulumi.Input[builtins.str]] = None,
                 route_name: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Route resource.
        :param pulumi.Input[Union[builtins.str, 'RouteNextHopType']] next_hop_type: The type of Azure hop the packet should be sent to.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group.
        :param pulumi.Input[builtins.str] route_table_name: The name of the route table.
        :param pulumi.Input[builtins.str] address_prefix: The destination CIDR to which the route applies.
        :param pulumi.Input[builtins.str] id: Resource ID.
        :param pulumi.Input[builtins.str] name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
        :param pulumi.Input[builtins.str] next_hop_ip_address: The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
        :param pulumi.Input[builtins.str] route_name: The name of the route.
        :param pulumi.Input[builtins.str] type: The type of the resource.
        """
        pulumi.set(__self__, "next_hop_type", next_hop_type)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "route_table_name", route_table_name)
        if address_prefix is not None:
            pulumi.set(__self__, "address_prefix", address_prefix)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if next_hop_ip_address is not None:
            pulumi.set(__self__, "next_hop_ip_address", next_hop_ip_address)
        if route_name is not None:
            pulumi.set(__self__, "route_name", route_name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="nextHopType")
    def next_hop_type(self) -> pulumi.Input[Union[builtins.str, 'RouteNextHopType']]:
        """
        The type of Azure hop the packet should be sent to.
        """
        return pulumi.get(self, "next_hop_type")

    @next_hop_type.setter
    def next_hop_type(self, value: pulumi.Input[Union[builtins.str, 'RouteNextHopType']]):
        pulumi.set(self, "next_hop_type", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="routeTableName")
    def route_table_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the route table.
        """
        return pulumi.get(self, "route_table_name")

    @route_table_name.setter
    def route_table_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "route_table_name", value)

    @property
    @pulumi.getter(name="addressPrefix")
    def address_prefix(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The destination CIDR to which the route applies.
        """
        return pulumi.get(self, "address_prefix")

    @address_prefix.setter
    def address_prefix(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "address_prefix", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nextHopIpAddress")
    def next_hop_ip_address(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
        """
        return pulumi.get(self, "next_hop_ip_address")

    @next_hop_ip_address.setter
    def next_hop_ip_address(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "next_hop_ip_address", value)

    @property
    @pulumi.getter(name="routeName")
    def route_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the route.
        """
        return pulumi.get(self, "route_name")

    @route_name.setter
    def route_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "route_name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)


@pulumi.type_token("azure-native:network:Route")
class Route(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_prefix: Optional[pulumi.Input[builtins.str]] = None,
                 id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 next_hop_ip_address: Optional[pulumi.Input[builtins.str]] = None,
                 next_hop_type: Optional[pulumi.Input[Union[builtins.str, 'RouteNextHopType']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 route_name: Optional[pulumi.Input[builtins.str]] = None,
                 route_table_name: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Route resource.

        Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.

        Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] address_prefix: The destination CIDR to which the route applies.
        :param pulumi.Input[builtins.str] id: Resource ID.
        :param pulumi.Input[builtins.str] name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
        :param pulumi.Input[builtins.str] next_hop_ip_address: The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
        :param pulumi.Input[Union[builtins.str, 'RouteNextHopType']] next_hop_type: The type of Azure hop the packet should be sent to.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group.
        :param pulumi.Input[builtins.str] route_name: The name of the route.
        :param pulumi.Input[builtins.str] route_table_name: The name of the route table.
        :param pulumi.Input[builtins.str] type: The type of the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RouteInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Route resource.

        Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.

        Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param RouteInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouteInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_prefix: Optional[pulumi.Input[builtins.str]] = None,
                 id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 next_hop_ip_address: Optional[pulumi.Input[builtins.str]] = None,
                 next_hop_type: Optional[pulumi.Input[Union[builtins.str, 'RouteNextHopType']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 route_name: Optional[pulumi.Input[builtins.str]] = None,
                 route_table_name: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RouteInitArgs.__new__(RouteInitArgs)

            __props__.__dict__["address_prefix"] = address_prefix
            __props__.__dict__["id"] = id
            __props__.__dict__["name"] = name
            __props__.__dict__["next_hop_ip_address"] = next_hop_ip_address
            if next_hop_type is None and not opts.urn:
                raise TypeError("Missing required property 'next_hop_type'")
            __props__.__dict__["next_hop_type"] = next_hop_type
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["route_name"] = route_name
            if route_table_name is None and not opts.urn:
                raise TypeError("Missing required property 'route_table_name'")
            __props__.__dict__["route_table_name"] = route_table_name
            __props__.__dict__["type"] = type
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["has_bgp_override"] = None
            __props__.__dict__["provisioning_state"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:network/v20150501preview:Route"), pulumi.Alias(type_="azure-native:network/v20150615:Route"), pulumi.Alias(type_="azure-native:network/v20160330:Route"), pulumi.Alias(type_="azure-native:network/v20160601:Route"), pulumi.Alias(type_="azure-native:network/v20160901:Route"), pulumi.Alias(type_="azure-native:network/v20161201:Route"), pulumi.Alias(type_="azure-native:network/v20170301:Route"), pulumi.Alias(type_="azure-native:network/v20170601:Route"), pulumi.Alias(type_="azure-native:network/v20170801:Route"), pulumi.Alias(type_="azure-native:network/v20170901:Route"), pulumi.Alias(type_="azure-native:network/v20171001:Route"), pulumi.Alias(type_="azure-native:network/v20171101:Route"), pulumi.Alias(type_="azure-native:network/v20180101:Route"), pulumi.Alias(type_="azure-native:network/v20180201:Route"), pulumi.Alias(type_="azure-native:network/v20180401:Route"), pulumi.Alias(type_="azure-native:network/v20180601:Route"), pulumi.Alias(type_="azure-native:network/v20180701:Route"), pulumi.Alias(type_="azure-native:network/v20180801:Route"), pulumi.Alias(type_="azure-native:network/v20181001:Route"), pulumi.Alias(type_="azure-native:network/v20181101:Route"), pulumi.Alias(type_="azure-native:network/v20181201:Route"), pulumi.Alias(type_="azure-native:network/v20190201:Route"), pulumi.Alias(type_="azure-native:network/v20190401:Route"), pulumi.Alias(type_="azure-native:network/v20190601:Route"), pulumi.Alias(type_="azure-native:network/v20190701:Route"), pulumi.Alias(type_="azure-native:network/v20190801:Route"), pulumi.Alias(type_="azure-native:network/v20190901:Route"), pulumi.Alias(type_="azure-native:network/v20191101:Route"), pulumi.Alias(type_="azure-native:network/v20191201:Route"), pulumi.Alias(type_="azure-native:network/v20200301:Route"), pulumi.Alias(type_="azure-native:network/v20200401:Route"), pulumi.Alias(type_="azure-native:network/v20200501:Route"), pulumi.Alias(type_="azure-native:network/v20200601:Route"), pulumi.Alias(type_="azure-native:network/v20200701:Route"), pulumi.Alias(type_="azure-native:network/v20200801:Route"), pulumi.Alias(type_="azure-native:network/v20201101:Route"), pulumi.Alias(type_="azure-native:network/v20210201:Route"), pulumi.Alias(type_="azure-native:network/v20210301:Route"), pulumi.Alias(type_="azure-native:network/v20210501:Route"), pulumi.Alias(type_="azure-native:network/v20210801:Route"), pulumi.Alias(type_="azure-native:network/v20220101:Route"), pulumi.Alias(type_="azure-native:network/v20220501:Route"), pulumi.Alias(type_="azure-native:network/v20220701:Route"), pulumi.Alias(type_="azure-native:network/v20220901:Route"), pulumi.Alias(type_="azure-native:network/v20221101:Route"), pulumi.Alias(type_="azure-native:network/v20230201:Route"), pulumi.Alias(type_="azure-native:network/v20230401:Route"), pulumi.Alias(type_="azure-native:network/v20230501:Route"), pulumi.Alias(type_="azure-native:network/v20230601:Route"), pulumi.Alias(type_="azure-native:network/v20230901:Route"), pulumi.Alias(type_="azure-native:network/v20231101:Route"), pulumi.Alias(type_="azure-native:network/v20240101:Route"), pulumi.Alias(type_="azure-native:network/v20240301:Route"), pulumi.Alias(type_="azure-native:network/v20240501:Route"), pulumi.Alias(type_="azure-native:network/v20240701:Route")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Route, __self__).__init__(
            'azure-native:network:Route',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Route':
        """
        Get an existing Route resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RouteInitArgs.__new__(RouteInitArgs)

        __props__.__dict__["address_prefix"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["has_bgp_override"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["next_hop_ip_address"] = None
        __props__.__dict__["next_hop_type"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["type"] = None
        return Route(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressPrefix")
    def address_prefix(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The destination CIDR to which the route applies.
        """
        return pulumi.get(self, "address_prefix")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[builtins.str]:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="hasBgpOverride")
    def has_bgp_override(self) -> pulumi.Output[builtins.bool]:
        """
        A value indicating whether this route overrides overlapping BGP routes regardless of LPM.
        """
        return pulumi.get(self, "has_bgp_override")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nextHopIpAddress")
    def next_hop_ip_address(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
        """
        return pulumi.get(self, "next_hop_ip_address")

    @property
    @pulumi.getter(name="nextHopType")
    def next_hop_type(self) -> pulumi.Output[builtins.str]:
        """
        The type of Azure hop the packet should be sent to.
        """
        return pulumi.get(self, "next_hop_type")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        The provisioning state of the route resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

