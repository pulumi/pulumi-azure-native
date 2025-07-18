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

__all__ = ['ConnectionRaiPolicyArgs', 'ConnectionRaiPolicy']

@pulumi.input_type
class ConnectionRaiPolicyArgs:
    def __init__(__self__, *,
                 connection_name: pulumi.Input[builtins.str],
                 properties: pulumi.Input['RaiPolicyPropertiesArgs'],
                 resource_group_name: pulumi.Input[builtins.str],
                 workspace_name: pulumi.Input[builtins.str],
                 proxy_api_version: Optional[pulumi.Input[builtins.str]] = None,
                 rai_policy_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ConnectionRaiPolicy resource.
        :param pulumi.Input[builtins.str] connection_name: Friendly name of the workspace connection
        :param pulumi.Input['RaiPolicyPropertiesArgs'] properties: Azure OpenAI Content Filters properties.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] workspace_name: Azure Machine Learning Workspace Name
        :param pulumi.Input[builtins.str] proxy_api_version: Api version used by proxy call
        :param pulumi.Input[builtins.str] rai_policy_name: Name of the Rai Policy.
        """
        pulumi.set(__self__, "connection_name", connection_name)
        pulumi.set(__self__, "properties", properties)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "workspace_name", workspace_name)
        if proxy_api_version is not None:
            pulumi.set(__self__, "proxy_api_version", proxy_api_version)
        if rai_policy_name is not None:
            pulumi.set(__self__, "rai_policy_name", rai_policy_name)

    @property
    @pulumi.getter(name="connectionName")
    def connection_name(self) -> pulumi.Input[builtins.str]:
        """
        Friendly name of the workspace connection
        """
        return pulumi.get(self, "connection_name")

    @connection_name.setter
    def connection_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "connection_name", value)

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Input['RaiPolicyPropertiesArgs']:
        """
        Azure OpenAI Content Filters properties.
        """
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: pulumi.Input['RaiPolicyPropertiesArgs']):
        pulumi.set(self, "properties", value)

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
    @pulumi.getter(name="workspaceName")
    def workspace_name(self) -> pulumi.Input[builtins.str]:
        """
        Azure Machine Learning Workspace Name
        """
        return pulumi.get(self, "workspace_name")

    @workspace_name.setter
    def workspace_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "workspace_name", value)

    @property
    @pulumi.getter(name="proxyApiVersion")
    def proxy_api_version(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Api version used by proxy call
        """
        return pulumi.get(self, "proxy_api_version")

    @proxy_api_version.setter
    def proxy_api_version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "proxy_api_version", value)

    @property
    @pulumi.getter(name="raiPolicyName")
    def rai_policy_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the Rai Policy.
        """
        return pulumi.get(self, "rai_policy_name")

    @rai_policy_name.setter
    def rai_policy_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "rai_policy_name", value)


@pulumi.type_token("azure-native:machinelearningservices:ConnectionRaiPolicy")
class ConnectionRaiPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['RaiPolicyPropertiesArgs', 'RaiPolicyPropertiesArgsDict']]] = None,
                 proxy_api_version: Optional[pulumi.Input[builtins.str]] = None,
                 rai_policy_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Azure OpenAI Content Filters resource.

        Uses Azure REST API version 2025-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2024-04-01-preview.

        Other available API versions: 2024-04-01-preview, 2024-07-01-preview, 2024-10-01-preview, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] connection_name: Friendly name of the workspace connection
        :param pulumi.Input[Union['RaiPolicyPropertiesArgs', 'RaiPolicyPropertiesArgsDict']] properties: Azure OpenAI Content Filters properties.
        :param pulumi.Input[builtins.str] proxy_api_version: Api version used by proxy call
        :param pulumi.Input[builtins.str] rai_policy_name: Name of the Rai Policy.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] workspace_name: Azure Machine Learning Workspace Name
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConnectionRaiPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Azure OpenAI Content Filters resource.

        Uses Azure REST API version 2025-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2024-04-01-preview.

        Other available API versions: 2024-04-01-preview, 2024-07-01-preview, 2024-10-01-preview, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param ConnectionRaiPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConnectionRaiPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['RaiPolicyPropertiesArgs', 'RaiPolicyPropertiesArgsDict']]] = None,
                 proxy_api_version: Optional[pulumi.Input[builtins.str]] = None,
                 rai_policy_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConnectionRaiPolicyArgs.__new__(ConnectionRaiPolicyArgs)

            if connection_name is None and not opts.urn:
                raise TypeError("Missing required property 'connection_name'")
            __props__.__dict__["connection_name"] = connection_name
            if properties is None and not opts.urn:
                raise TypeError("Missing required property 'properties'")
            __props__.__dict__["properties"] = properties
            __props__.__dict__["proxy_api_version"] = proxy_api_version
            __props__.__dict__["rai_policy_name"] = rai_policy_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if workspace_name is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_name'")
            __props__.__dict__["workspace_name"] = workspace_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:machinelearningservices/v20240401preview:ConnectionRaiPolicy"), pulumi.Alias(type_="azure-native:machinelearningservices/v20240701preview:ConnectionRaiPolicy"), pulumi.Alias(type_="azure-native:machinelearningservices/v20241001preview:ConnectionRaiPolicy"), pulumi.Alias(type_="azure-native:machinelearningservices/v20250101preview:ConnectionRaiPolicy"), pulumi.Alias(type_="azure-native:machinelearningservices/v20250401preview:ConnectionRaiPolicy"), pulumi.Alias(type_="azure-native:machinelearningservices/v20250701preview:ConnectionRaiPolicy")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ConnectionRaiPolicy, __self__).__init__(
            'azure-native:machinelearningservices:ConnectionRaiPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ConnectionRaiPolicy':
        """
        Get an existing ConnectionRaiPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ConnectionRaiPolicyArgs.__new__(ConnectionRaiPolicyArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["properties"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return ConnectionRaiPolicy(resource_name, opts=opts, __props__=__props__)

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
    def properties(self) -> pulumi.Output['outputs.RaiPolicyPropertiesResponse']:
        """
        Azure OpenAI Content Filters properties.
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

