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

__all__ = ['PrivateLinkScopedResourceArgs', 'PrivateLinkScopedResource']

@pulumi.input_type
class PrivateLinkScopedResourceArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 scope_name: pulumi.Input[builtins.str],
                 kind: Optional[pulumi.Input[Union[builtins.str, 'ScopedResourceKind']]] = None,
                 linked_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 subscription_location: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a PrivateLinkScopedResource resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] scope_name: The name of the Azure Monitor PrivateLinkScope resource.
        :param pulumi.Input[Union[builtins.str, 'ScopedResourceKind']] kind: The kind of scoped Azure monitor resource.
        :param pulumi.Input[builtins.str] linked_resource_id: The resource id of the scoped Azure monitor resource.
        :param pulumi.Input[builtins.str] name: The name of the scoped resource object.
        :param pulumi.Input[builtins.str] subscription_location: The location of a scoped subscription. Only needs to be specified for metric dataplane subscriptions.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "scope_name", scope_name)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if linked_resource_id is not None:
            pulumi.set(__self__, "linked_resource_id", linked_resource_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if subscription_location is not None:
            pulumi.set(__self__, "subscription_location", subscription_location)

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
    @pulumi.getter(name="scopeName")
    def scope_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the Azure Monitor PrivateLinkScope resource.
        """
        return pulumi.get(self, "scope_name")

    @scope_name.setter
    def scope_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "scope_name", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[Union[builtins.str, 'ScopedResourceKind']]]:
        """
        The kind of scoped Azure monitor resource.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[Union[builtins.str, 'ScopedResourceKind']]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="linkedResourceId")
    def linked_resource_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The resource id of the scoped Azure monitor resource.
        """
        return pulumi.get(self, "linked_resource_id")

    @linked_resource_id.setter
    def linked_resource_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "linked_resource_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the scoped resource object.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="subscriptionLocation")
    def subscription_location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The location of a scoped subscription. Only needs to be specified for metric dataplane subscriptions.
        """
        return pulumi.get(self, "subscription_location")

    @subscription_location.setter
    def subscription_location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "subscription_location", value)


@pulumi.type_token("azure-native:monitor:PrivateLinkScopedResource")
class PrivateLinkScopedResource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 kind: Optional[pulumi.Input[Union[builtins.str, 'ScopedResourceKind']]] = None,
                 linked_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 scope_name: Optional[pulumi.Input[builtins.str]] = None,
                 subscription_location: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        A private link scoped resource

        Uses Azure REST API version 2023-06-01-preview.

        Other available API versions: 2021-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native monitor [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union[builtins.str, 'ScopedResourceKind']] kind: The kind of scoped Azure monitor resource.
        :param pulumi.Input[builtins.str] linked_resource_id: The resource id of the scoped Azure monitor resource.
        :param pulumi.Input[builtins.str] name: The name of the scoped resource object.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] scope_name: The name of the Azure Monitor PrivateLinkScope resource.
        :param pulumi.Input[builtins.str] subscription_location: The location of a scoped subscription. Only needs to be specified for metric dataplane subscriptions.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateLinkScopedResourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A private link scoped resource

        Uses Azure REST API version 2023-06-01-preview.

        Other available API versions: 2021-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native monitor [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param PrivateLinkScopedResourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateLinkScopedResourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 kind: Optional[pulumi.Input[Union[builtins.str, 'ScopedResourceKind']]] = None,
                 linked_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 scope_name: Optional[pulumi.Input[builtins.str]] = None,
                 subscription_location: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateLinkScopedResourceArgs.__new__(PrivateLinkScopedResourceArgs)

            __props__.__dict__["kind"] = kind
            __props__.__dict__["linked_resource_id"] = linked_resource_id
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if scope_name is None and not opts.urn:
                raise TypeError("Missing required property 'scope_name'")
            __props__.__dict__["scope_name"] = scope_name
            __props__.__dict__["subscription_location"] = subscription_location
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:insights/v20210701preview:PrivateLinkScopedResource"), pulumi.Alias(type_="azure-native:insights/v20210901:PrivateLinkScopedResource"), pulumi.Alias(type_="azure-native:insights/v20230601preview:PrivateLinkScopedResource"), pulumi.Alias(type_="azure-native:insights:PrivateLinkScopedResource"), pulumi.Alias(type_="azure-native:monitor/v20191017preview:PrivateLinkScopedResource"), pulumi.Alias(type_="azure-native:monitor/v20210701preview:PrivateLinkScopedResource"), pulumi.Alias(type_="azure-native:monitor/v20210901:PrivateLinkScopedResource"), pulumi.Alias(type_="azure-native:monitor/v20230601preview:PrivateLinkScopedResource")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(PrivateLinkScopedResource, __self__).__init__(
            'azure-native:monitor:PrivateLinkScopedResource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PrivateLinkScopedResource':
        """
        Get an existing PrivateLinkScopedResource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PrivateLinkScopedResourceArgs.__new__(PrivateLinkScopedResourceArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["linked_resource_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["subscription_location"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return PrivateLinkScopedResource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The kind of scoped Azure monitor resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="linkedResourceId")
    def linked_resource_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The resource id of the scoped Azure monitor resource.
        """
        return pulumi.get(self, "linked_resource_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        State of the Azure monitor resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="subscriptionLocation")
    def subscription_location(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The location of a scoped subscription. Only needs to be specified for metric dataplane subscriptions.
        """
        return pulumi.get(self, "subscription_location")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        System data
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

