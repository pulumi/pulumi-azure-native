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

__all__ = ['IpPrefixArgs', 'IpPrefix']

@pulumi.input_type
class IpPrefixArgs:
    def __init__(__self__, *,
                 ip_prefix_rules: pulumi.Input[Sequence[pulumi.Input['IpPrefixRuleArgs']]],
                 resource_group_name: pulumi.Input[builtins.str],
                 annotation: Optional[pulumi.Input[builtins.str]] = None,
                 ip_prefix_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a IpPrefix resource.
        :param pulumi.Input[Sequence[pulumi.Input['IpPrefixRuleArgs']]] ip_prefix_rules: The list of IP Prefix Rules.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] annotation: Switch configuration description.
        :param pulumi.Input[builtins.str] ip_prefix_name: Name of the IP Prefix.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        pulumi.set(__self__, "ip_prefix_rules", ip_prefix_rules)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if annotation is not None:
            pulumi.set(__self__, "annotation", annotation)
        if ip_prefix_name is not None:
            pulumi.set(__self__, "ip_prefix_name", ip_prefix_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="ipPrefixRules")
    def ip_prefix_rules(self) -> pulumi.Input[Sequence[pulumi.Input['IpPrefixRuleArgs']]]:
        """
        The list of IP Prefix Rules.
        """
        return pulumi.get(self, "ip_prefix_rules")

    @ip_prefix_rules.setter
    def ip_prefix_rules(self, value: pulumi.Input[Sequence[pulumi.Input['IpPrefixRuleArgs']]]):
        pulumi.set(self, "ip_prefix_rules", value)

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
    @pulumi.getter
    def annotation(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Switch configuration description.
        """
        return pulumi.get(self, "annotation")

    @annotation.setter
    def annotation(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "annotation", value)

    @property
    @pulumi.getter(name="ipPrefixName")
    def ip_prefix_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the IP Prefix.
        """
        return pulumi.get(self, "ip_prefix_name")

    @ip_prefix_name.setter
    def ip_prefix_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ip_prefix_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

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


@pulumi.type_token("azure-native:managednetworkfabric:IpPrefix")
class IpPrefix(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 annotation: Optional[pulumi.Input[builtins.str]] = None,
                 ip_prefix_name: Optional[pulumi.Input[builtins.str]] = None,
                 ip_prefix_rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IpPrefixRuleArgs', 'IpPrefixRuleArgsDict']]]]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        The IP Prefix resource definition.

        Uses Azure REST API version 2023-06-15. In version 2.x of the Azure Native provider, it used API version 2023-02-01-preview.

        Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] annotation: Switch configuration description.
        :param pulumi.Input[builtins.str] ip_prefix_name: Name of the IP Prefix.
        :param pulumi.Input[Sequence[pulumi.Input[Union['IpPrefixRuleArgs', 'IpPrefixRuleArgsDict']]]] ip_prefix_rules: The list of IP Prefix Rules.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpPrefixArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The IP Prefix resource definition.

        Uses Azure REST API version 2023-06-15. In version 2.x of the Azure Native provider, it used API version 2023-02-01-preview.

        Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param IpPrefixArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpPrefixArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 annotation: Optional[pulumi.Input[builtins.str]] = None,
                 ip_prefix_name: Optional[pulumi.Input[builtins.str]] = None,
                 ip_prefix_rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IpPrefixRuleArgs', 'IpPrefixRuleArgsDict']]]]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpPrefixArgs.__new__(IpPrefixArgs)

            __props__.__dict__["annotation"] = annotation
            __props__.__dict__["ip_prefix_name"] = ip_prefix_name
            if ip_prefix_rules is None and not opts.urn:
                raise TypeError("Missing required property 'ip_prefix_rules'")
            __props__.__dict__["ip_prefix_rules"] = ip_prefix_rules
            __props__.__dict__["location"] = location
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["administrative_state"] = None
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["configuration_state"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:managednetworkfabric/v20230201preview:IpPrefix"), pulumi.Alias(type_="azure-native:managednetworkfabric/v20230615:IpPrefix")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(IpPrefix, __self__).__init__(
            'azure-native:managednetworkfabric:IpPrefix',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'IpPrefix':
        """
        Get an existing IpPrefix resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = IpPrefixArgs.__new__(IpPrefixArgs)

        __props__.__dict__["administrative_state"] = None
        __props__.__dict__["annotation"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["configuration_state"] = None
        __props__.__dict__["ip_prefix_rules"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return IpPrefix(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="administrativeState")
    def administrative_state(self) -> pulumi.Output[builtins.str]:
        """
        Administrative state of the resource.
        """
        return pulumi.get(self, "administrative_state")

    @property
    @pulumi.getter
    def annotation(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Switch configuration description.
        """
        return pulumi.get(self, "annotation")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="configurationState")
    def configuration_state(self) -> pulumi.Output[builtins.str]:
        """
        Configuration state of the resource.
        """
        return pulumi.get(self, "configuration_state")

    @property
    @pulumi.getter(name="ipPrefixRules")
    def ip_prefix_rules(self) -> pulumi.Output[Sequence['outputs.IpPrefixRuleResponse']]:
        """
        The list of IP Prefix Rules.
        """
        return pulumi.get(self, "ip_prefix_rules")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

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
        Provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

