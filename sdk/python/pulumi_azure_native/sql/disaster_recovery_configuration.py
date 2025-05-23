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

__all__ = ['DisasterRecoveryConfigurationArgs', 'DisasterRecoveryConfiguration']

@pulumi.input_type
class DisasterRecoveryConfigurationArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 server_name: pulumi.Input[builtins.str],
                 disaster_recovery_configuration_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a DisasterRecoveryConfiguration resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[builtins.str] server_name: The name of the server.
        :param pulumi.Input[builtins.str] disaster_recovery_configuration_name: The name of the disaster recovery configuration to be created/updated.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "server_name", server_name)
        if disaster_recovery_configuration_name is not None:
            pulumi.set(__self__, "disaster_recovery_configuration_name", disaster_recovery_configuration_name)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the server.
        """
        return pulumi.get(self, "server_name")

    @server_name.setter
    def server_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "server_name", value)

    @property
    @pulumi.getter(name="disasterRecoveryConfigurationName")
    def disaster_recovery_configuration_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the disaster recovery configuration to be created/updated.
        """
        return pulumi.get(self, "disaster_recovery_configuration_name")

    @disaster_recovery_configuration_name.setter
    def disaster_recovery_configuration_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "disaster_recovery_configuration_name", value)


@pulumi.type_token("azure-native:sql:DisasterRecoveryConfiguration")
class DisasterRecoveryConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disaster_recovery_configuration_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 server_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Represents a disaster recovery configuration.

        Uses Azure REST API version 2014-04-01. In version 2.x of the Azure Native provider, it used API version 2014-04-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] disaster_recovery_configuration_name: The name of the disaster recovery configuration to be created/updated.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[builtins.str] server_name: The name of the server.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DisasterRecoveryConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a disaster recovery configuration.

        Uses Azure REST API version 2014-04-01. In version 2.x of the Azure Native provider, it used API version 2014-04-01.

        :param str resource_name: The name of the resource.
        :param DisasterRecoveryConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DisasterRecoveryConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disaster_recovery_configuration_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 server_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DisasterRecoveryConfigurationArgs.__new__(DisasterRecoveryConfigurationArgs)

            __props__.__dict__["disaster_recovery_configuration_name"] = disaster_recovery_configuration_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if server_name is None and not opts.urn:
                raise TypeError("Missing required property 'server_name'")
            __props__.__dict__["server_name"] = server_name
            __props__.__dict__["auto_failover"] = None
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["failover_policy"] = None
            __props__.__dict__["location"] = None
            __props__.__dict__["logical_server_name"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["partner_logical_server_name"] = None
            __props__.__dict__["partner_server_id"] = None
            __props__.__dict__["role"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:sql/v20140401:DisasterRecoveryConfiguration")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DisasterRecoveryConfiguration, __self__).__init__(
            'azure-native:sql:DisasterRecoveryConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DisasterRecoveryConfiguration':
        """
        Get an existing DisasterRecoveryConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DisasterRecoveryConfigurationArgs.__new__(DisasterRecoveryConfigurationArgs)

        __props__.__dict__["auto_failover"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["failover_policy"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["logical_server_name"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["partner_logical_server_name"] = None
        __props__.__dict__["partner_server_id"] = None
        __props__.__dict__["role"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["type"] = None
        return DisasterRecoveryConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoFailover")
    def auto_failover(self) -> pulumi.Output[builtins.str]:
        """
        Whether or not failover can be done automatically.
        """
        return pulumi.get(self, "auto_failover")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="failoverPolicy")
    def failover_policy(self) -> pulumi.Output[builtins.str]:
        """
        How aggressive the automatic failover should be.
        """
        return pulumi.get(self, "failover_policy")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        Location of the server that contains this disaster recovery configuration.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="logicalServerName")
    def logical_server_name(self) -> pulumi.Output[builtins.str]:
        """
        Logical name of the server.
        """
        return pulumi.get(self, "logical_server_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="partnerLogicalServerName")
    def partner_logical_server_name(self) -> pulumi.Output[builtins.str]:
        """
        Logical name of the partner server.
        """
        return pulumi.get(self, "partner_logical_server_name")

    @property
    @pulumi.getter(name="partnerServerId")
    def partner_server_id(self) -> pulumi.Output[builtins.str]:
        """
        Id of the partner server.
        """
        return pulumi.get(self, "partner_server_id")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[builtins.str]:
        """
        The role of the current server in the disaster recovery configuration.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        The status of the disaster recovery configuration.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

