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

__all__ = ['NamespaceIpFilterRuleArgs', 'NamespaceIpFilterRule']

@pulumi.input_type
class NamespaceIpFilterRuleArgs:
    def __init__(__self__, *,
                 namespace_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 action: Optional[pulumi.Input[Union[builtins.str, 'IPAction']]] = None,
                 filter_name: Optional[pulumi.Input[builtins.str]] = None,
                 ip_filter_rule_name: Optional[pulumi.Input[builtins.str]] = None,
                 ip_mask: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a NamespaceIpFilterRule resource.
        :param pulumi.Input[builtins.str] namespace_name: The namespace name
        :param pulumi.Input[builtins.str] resource_group_name: Name of the Resource group within the Azure subscription.
        :param pulumi.Input[Union[builtins.str, 'IPAction']] action: The IP Filter Action
        :param pulumi.Input[builtins.str] filter_name: IP Filter name
        :param pulumi.Input[builtins.str] ip_filter_rule_name: The IP Filter Rule name.
        :param pulumi.Input[builtins.str] ip_mask: IP Mask
        """
        pulumi.set(__self__, "namespace_name", namespace_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if filter_name is not None:
            pulumi.set(__self__, "filter_name", filter_name)
        if ip_filter_rule_name is not None:
            pulumi.set(__self__, "ip_filter_rule_name", ip_filter_rule_name)
        if ip_mask is not None:
            pulumi.set(__self__, "ip_mask", ip_mask)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Input[builtins.str]:
        """
        The namespace name
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the Resource group within the Azure subscription.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[Union[builtins.str, 'IPAction']]]:
        """
        The IP Filter Action
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[Union[builtins.str, 'IPAction']]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="filterName")
    def filter_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        IP Filter name
        """
        return pulumi.get(self, "filter_name")

    @filter_name.setter
    def filter_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "filter_name", value)

    @property
    @pulumi.getter(name="ipFilterRuleName")
    def ip_filter_rule_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The IP Filter Rule name.
        """
        return pulumi.get(self, "ip_filter_rule_name")

    @ip_filter_rule_name.setter
    def ip_filter_rule_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ip_filter_rule_name", value)

    @property
    @pulumi.getter(name="ipMask")
    def ip_mask(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        IP Mask
        """
        return pulumi.get(self, "ip_mask")

    @ip_mask.setter
    def ip_mask(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ip_mask", value)


@pulumi.type_token("azure-native:servicebus:NamespaceIpFilterRule")
class NamespaceIpFilterRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[Union[builtins.str, 'IPAction']]] = None,
                 filter_name: Optional[pulumi.Input[builtins.str]] = None,
                 ip_filter_rule_name: Optional[pulumi.Input[builtins.str]] = None,
                 ip_mask: Optional[pulumi.Input[builtins.str]] = None,
                 namespace_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Single item in a List or Get IpFilterRules operation

        Uses Azure REST API version 2018-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2018-01-01-preview.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union[builtins.str, 'IPAction']] action: The IP Filter Action
        :param pulumi.Input[builtins.str] filter_name: IP Filter name
        :param pulumi.Input[builtins.str] ip_filter_rule_name: The IP Filter Rule name.
        :param pulumi.Input[builtins.str] ip_mask: IP Mask
        :param pulumi.Input[builtins.str] namespace_name: The namespace name
        :param pulumi.Input[builtins.str] resource_group_name: Name of the Resource group within the Azure subscription.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NamespaceIpFilterRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Single item in a List or Get IpFilterRules operation

        Uses Azure REST API version 2018-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2018-01-01-preview.

        :param str resource_name: The name of the resource.
        :param NamespaceIpFilterRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NamespaceIpFilterRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[Union[builtins.str, 'IPAction']]] = None,
                 filter_name: Optional[pulumi.Input[builtins.str]] = None,
                 ip_filter_rule_name: Optional[pulumi.Input[builtins.str]] = None,
                 ip_mask: Optional[pulumi.Input[builtins.str]] = None,
                 namespace_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NamespaceIpFilterRuleArgs.__new__(NamespaceIpFilterRuleArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["filter_name"] = filter_name
            __props__.__dict__["ip_filter_rule_name"] = ip_filter_rule_name
            __props__.__dict__["ip_mask"] = ip_mask
            if namespace_name is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_name'")
            __props__.__dict__["namespace_name"] = namespace_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:servicebus/v20180101preview:NamespaceIpFilterRule")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(NamespaceIpFilterRule, __self__).__init__(
            'azure-native:servicebus:NamespaceIpFilterRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NamespaceIpFilterRule':
        """
        Get an existing NamespaceIpFilterRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NamespaceIpFilterRuleArgs.__new__(NamespaceIpFilterRuleArgs)

        __props__.__dict__["action"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["filter_name"] = None
        __props__.__dict__["ip_mask"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["type"] = None
        return NamespaceIpFilterRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The IP Filter Action
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="filterName")
    def filter_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        IP Filter name
        """
        return pulumi.get(self, "filter_name")

    @property
    @pulumi.getter(name="ipMask")
    def ip_mask(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        IP Mask
        """
        return pulumi.get(self, "ip_mask")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

