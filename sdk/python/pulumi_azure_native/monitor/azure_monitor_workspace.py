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

__all__ = ['AzureMonitorWorkspaceArgs', 'AzureMonitorWorkspace']

@pulumi.input_type
class AzureMonitorWorkspaceArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 azure_monitor_workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 public_network_access: Optional[pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a AzureMonitorWorkspace resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] azure_monitor_workspace_name: The name of the Azure Monitor Workspace. The name is case insensitive
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']] public_network_access: Gets or sets allow or disallow public network access to Azure Monitor Workspace
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if azure_monitor_workspace_name is not None:
            pulumi.set(__self__, "azure_monitor_workspace_name", azure_monitor_workspace_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if public_network_access is not None:
            pulumi.set(__self__, "public_network_access", public_network_access)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

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
    @pulumi.getter(name="azureMonitorWorkspaceName")
    def azure_monitor_workspace_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the Azure Monitor Workspace. The name is case insensitive
        """
        return pulumi.get(self, "azure_monitor_workspace_name")

    @azure_monitor_workspace_name.setter
    def azure_monitor_workspace_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "azure_monitor_workspace_name", value)

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
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> Optional[pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']]]:
        """
        Gets or sets allow or disallow public network access to Azure Monitor Workspace
        """
        return pulumi.get(self, "public_network_access")

    @public_network_access.setter
    def public_network_access(self, value: Optional[pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']]]):
        pulumi.set(self, "public_network_access", value)

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


@pulumi.type_token("azure-native:monitor:AzureMonitorWorkspace")
class AzureMonitorWorkspace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 azure_monitor_workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 public_network_access: Optional[pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        An Azure Monitor Workspace definition

        Uses Azure REST API version 2023-04-03. In version 2.x of the Azure Native provider, it used API version 2023-04-03.

        Other available API versions: 2023-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native monitor [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] azure_monitor_workspace_name: The name of the Azure Monitor Workspace. The name is case insensitive
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']] public_network_access: Gets or sets allow or disallow public network access to Azure Monitor Workspace
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AzureMonitorWorkspaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An Azure Monitor Workspace definition

        Uses Azure REST API version 2023-04-03. In version 2.x of the Azure Native provider, it used API version 2023-04-03.

        Other available API versions: 2023-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native monitor [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param AzureMonitorWorkspaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AzureMonitorWorkspaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 azure_monitor_workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 public_network_access: Optional[pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AzureMonitorWorkspaceArgs.__new__(AzureMonitorWorkspaceArgs)

            __props__.__dict__["azure_monitor_workspace_name"] = azure_monitor_workspace_name
            __props__.__dict__["location"] = location
            __props__.__dict__["public_network_access"] = public_network_access
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["account_id"] = None
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["default_ingestion_settings"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["metrics"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["private_endpoint_connections"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:monitor/v20210603preview:AzureMonitorWorkspace"), pulumi.Alias(type_="azure-native:monitor/v20230403:AzureMonitorWorkspace"), pulumi.Alias(type_="azure-native:monitor/v20231001preview:AzureMonitorWorkspace")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AzureMonitorWorkspace, __self__).__init__(
            'azure-native:monitor:AzureMonitorWorkspace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AzureMonitorWorkspace':
        """
        Get an existing AzureMonitorWorkspace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AzureMonitorWorkspaceArgs.__new__(AzureMonitorWorkspaceArgs)

        __props__.__dict__["account_id"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["default_ingestion_settings"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["metrics"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["private_endpoint_connections"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["public_network_access"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return AzureMonitorWorkspace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[builtins.str]:
        """
        The immutable Id of the Azure Monitor Workspace. This property is read-only.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="defaultIngestionSettings")
    def default_ingestion_settings(self) -> pulumi.Output['outputs.AzureMonitorWorkspaceResponseDefaultIngestionSettings']:
        """
        The Data Collection Rule and Endpoint used for ingestion by default.
        """
        return pulumi.get(self, "default_ingestion_settings")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[builtins.str]:
        """
        Resource entity tag (ETag)
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def metrics(self) -> pulumi.Output['outputs.AzureMonitorWorkspaceResponseMetrics']:
        """
        Properties related to the metrics container in the Azure Monitor Workspace
        """
        return pulumi.get(self, "metrics")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateEndpointConnections")
    def private_endpoint_connections(self) -> pulumi.Output[Sequence['outputs.PrivateEndpointConnectionResponse']]:
        """
        List of private endpoint connections
        """
        return pulumi.get(self, "private_endpoint_connections")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        The provisioning state of the Azure Monitor Workspace. Set to Succeeded if everything is healthy.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Gets or sets allow or disallow public network access to Azure Monitor Workspace
        """
        return pulumi.get(self, "public_network_access")

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

