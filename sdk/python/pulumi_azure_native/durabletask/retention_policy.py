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

__all__ = ['RetentionPolicyArgs', 'RetentionPolicy']

@pulumi.input_type
class RetentionPolicyArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 scheduler_name: pulumi.Input[builtins.str],
                 properties: Optional[pulumi.Input['RetentionPolicyPropertiesArgs']] = None):
        """
        The set of arguments for constructing a RetentionPolicy resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] scheduler_name: The name of the Scheduler
        :param pulumi.Input['RetentionPolicyPropertiesArgs'] properties: The resource-specific properties for this resource.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "scheduler_name", scheduler_name)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)

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
    @pulumi.getter(name="schedulerName")
    def scheduler_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the Scheduler
        """
        return pulumi.get(self, "scheduler_name")

    @scheduler_name.setter
    def scheduler_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "scheduler_name", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input['RetentionPolicyPropertiesArgs']]:
        """
        The resource-specific properties for this resource.
        """
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input['RetentionPolicyPropertiesArgs']]):
        pulumi.set(self, "properties", value)


@pulumi.type_token("azure-native:durabletask:RetentionPolicy")
class RetentionPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 properties: Optional[pulumi.Input[Union['RetentionPolicyPropertiesArgs', 'RetentionPolicyPropertiesArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 scheduler_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        A retention policy resource belonging to the scheduler

        Uses Azure REST API version 2025-04-01-preview.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['RetentionPolicyPropertiesArgs', 'RetentionPolicyPropertiesArgsDict']] properties: The resource-specific properties for this resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] scheduler_name: The name of the Scheduler
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RetentionPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A retention policy resource belonging to the scheduler

        Uses Azure REST API version 2025-04-01-preview.

        :param str resource_name: The name of the resource.
        :param RetentionPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RetentionPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 properties: Optional[pulumi.Input[Union['RetentionPolicyPropertiesArgs', 'RetentionPolicyPropertiesArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 scheduler_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RetentionPolicyArgs.__new__(RetentionPolicyArgs)

            __props__.__dict__["properties"] = properties
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if scheduler_name is None and not opts.urn:
                raise TypeError("Missing required property 'scheduler_name'")
            __props__.__dict__["scheduler_name"] = scheduler_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:durabletask/v20250401preview:RetentionPolicy")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(RetentionPolicy, __self__).__init__(
            'azure-native:durabletask:RetentionPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RetentionPolicy':
        """
        Get an existing RetentionPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RetentionPolicyArgs.__new__(RetentionPolicyArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["properties"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return RetentionPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output['outputs.RetentionPolicyPropertiesResponse']:
        """
        The resource-specific properties for this resource.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

