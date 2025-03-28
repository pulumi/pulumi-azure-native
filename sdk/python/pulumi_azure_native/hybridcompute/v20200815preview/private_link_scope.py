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
from ._enums import *

__all__ = ['PrivateLinkScopeArgs', 'PrivateLinkScope']

@pulumi.input_type
class PrivateLinkScopeArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[str],
                 location: Optional[pulumi.Input[str]] = None,
                 public_network_access: Optional[pulumi.Input[Union[str, 'PublicNetworkAccessType']]] = None,
                 scope_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a PrivateLinkScope resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[Union[str, 'PublicNetworkAccessType']] public_network_access: Indicates whether machines associated with the private link scope can also use public Azure Arc service endpoints.
        :param pulumi.Input[str] scope_name: The name of the Azure Arc PrivateLinkScope resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if public_network_access is not None:
            pulumi.set(__self__, "public_network_access", public_network_access)
        if scope_name is not None:
            pulumi.set(__self__, "scope_name", scope_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> Optional[pulumi.Input[Union[str, 'PublicNetworkAccessType']]]:
        """
        Indicates whether machines associated with the private link scope can also use public Azure Arc service endpoints.
        """
        return pulumi.get(self, "public_network_access")

    @public_network_access.setter
    def public_network_access(self, value: Optional[pulumi.Input[Union[str, 'PublicNetworkAccessType']]]):
        pulumi.set(self, "public_network_access", value)

    @property
    @pulumi.getter(name="scopeName")
    def scope_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Azure Arc PrivateLinkScope resource.
        """
        return pulumi.get(self, "scope_name")

    @scope_name.setter
    def scope_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class PrivateLinkScope(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 public_network_access: Optional[pulumi.Input[Union[str, 'PublicNetworkAccessType']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 scope_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        An Azure Arc PrivateLinkScope definition.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[Union[str, 'PublicNetworkAccessType']] public_network_access: Indicates whether machines associated with the private link scope can also use public Azure Arc service endpoints.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] scope_name: The name of the Azure Arc PrivateLinkScope resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateLinkScopeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An Azure Arc PrivateLinkScope definition.

        :param str resource_name: The name of the resource.
        :param PrivateLinkScopeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateLinkScopeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 public_network_access: Optional[pulumi.Input[Union[str, 'PublicNetworkAccessType']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 scope_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateLinkScopeArgs.__new__(PrivateLinkScopeArgs)

            __props__.__dict__["location"] = location
            __props__.__dict__["public_network_access"] = public_network_access
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["scope_name"] = scope_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["name"] = None
            __props__.__dict__["private_endpoint_connections"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:hybridcompute/v20210128preview:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20210325preview:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20210422preview:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20210517preview:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20210520:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20210610preview:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20211210preview:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20220310:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20220510preview:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20220811preview:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20221110:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20221227:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20221227preview:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20230315preview:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20230620preview:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20231003preview:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20240331preview:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20240520preview:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20240710:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20240731preview:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20240910preview:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20241110preview:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute/v20250113:PrivateLinkScope"), pulumi.Alias(type_="azure-native:hybridcompute:PrivateLinkScope")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(PrivateLinkScope, __self__).__init__(
            'azure-native:hybridcompute/v20200815preview:PrivateLinkScope',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PrivateLinkScope':
        """
        Get an existing PrivateLinkScope resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PrivateLinkScopeArgs.__new__(PrivateLinkScopeArgs)

        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["private_endpoint_connections"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["public_network_access"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return PrivateLinkScope(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Azure resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateEndpointConnections")
    def private_endpoint_connections(self) -> pulumi.Output[Sequence['outputs.PrivateEndpointConnectionResponse']]:
        """
        List of private endpoint connections.
        """
        return pulumi.get(self, "private_endpoint_connections")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Current state of this PrivateLinkScope: whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Provisioning ,Succeeded, Canceled and Failed.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> pulumi.Output[Optional[str]]:
        """
        Indicates whether machines associated with the private link scope can also use public Azure Arc service endpoints.
        """
        return pulumi.get(self, "public_network_access")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Azure resource type
        """
        return pulumi.get(self, "type")

