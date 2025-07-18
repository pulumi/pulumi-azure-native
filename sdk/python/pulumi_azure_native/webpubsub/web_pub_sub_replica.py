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

__all__ = ['WebPubSubReplicaArgs', 'WebPubSubReplica']

@pulumi.input_type
class WebPubSubReplicaArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 resource_name: pulumi.Input[builtins.str],
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 region_endpoint_enabled: Optional[pulumi.Input[builtins.str]] = None,
                 replica_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_stopped: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input['ResourceSkuArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a WebPubSubReplica resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] resource_name: The name of the resource.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] region_endpoint_enabled: Enable or disable the regional endpoint. Default to "Enabled".
               When it's Disabled, new connections will not be routed to this endpoint, however existing connections will not be affected.
        :param pulumi.Input[builtins.str] replica_name: The name of the replica.
        :param pulumi.Input[builtins.str] resource_stopped: Stop or start the resource.  Default to "false".
               When it's true, the data plane of the resource is shutdown.
               When it's false, the data plane of the resource is started.
        :param pulumi.Input['ResourceSkuArgs'] sku: The billing information of the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "resource_name", resource_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if region_endpoint_enabled is None:
            region_endpoint_enabled = 'Enabled'
        if region_endpoint_enabled is not None:
            pulumi.set(__self__, "region_endpoint_enabled", region_endpoint_enabled)
        if replica_name is not None:
            pulumi.set(__self__, "replica_name", replica_name)
        if resource_stopped is None:
            resource_stopped = 'false'
        if resource_stopped is not None:
            pulumi.set(__self__, "resource_stopped", resource_stopped)
        if sku is not None:
            pulumi.set(__self__, "sku", sku)
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
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "resource_name")

    @resource_name.setter
    def resource_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_name", value)

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
    @pulumi.getter(name="regionEndpointEnabled")
    def region_endpoint_enabled(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Enable or disable the regional endpoint. Default to "Enabled".
        When it's Disabled, new connections will not be routed to this endpoint, however existing connections will not be affected.
        """
        return pulumi.get(self, "region_endpoint_enabled")

    @region_endpoint_enabled.setter
    def region_endpoint_enabled(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region_endpoint_enabled", value)

    @property
    @pulumi.getter(name="replicaName")
    def replica_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the replica.
        """
        return pulumi.get(self, "replica_name")

    @replica_name.setter
    def replica_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "replica_name", value)

    @property
    @pulumi.getter(name="resourceStopped")
    def resource_stopped(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Stop or start the resource.  Default to "false".
        When it's true, the data plane of the resource is shutdown.
        When it's false, the data plane of the resource is started.
        """
        return pulumi.get(self, "resource_stopped")

    @resource_stopped.setter
    def resource_stopped(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "resource_stopped", value)

    @property
    @pulumi.getter
    def sku(self) -> Optional[pulumi.Input['ResourceSkuArgs']]:
        """
        The billing information of the resource.
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: Optional[pulumi.Input['ResourceSkuArgs']]):
        pulumi.set(self, "sku", value)

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


@pulumi.type_token("azure-native:webpubsub:WebPubSubReplica")
class WebPubSubReplica(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 region_endpoint_enabled: Optional[pulumi.Input[builtins.str]] = None,
                 replica_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_name_: Optional[pulumi.Input[builtins.str]] = None,
                 resource_stopped: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input[Union['ResourceSkuArgs', 'ResourceSkuArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        A class represent a replica resource.

        Uses Azure REST API version 2024-03-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01-preview.

        Other available API versions: 2023-03-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native webpubsub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] region_endpoint_enabled: Enable or disable the regional endpoint. Default to "Enabled".
               When it's Disabled, new connections will not be routed to this endpoint, however existing connections will not be affected.
        :param pulumi.Input[builtins.str] replica_name: The name of the replica.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] resource_name_: The name of the resource.
        :param pulumi.Input[builtins.str] resource_stopped: Stop or start the resource.  Default to "false".
               When it's true, the data plane of the resource is shutdown.
               When it's false, the data plane of the resource is started.
        :param pulumi.Input[Union['ResourceSkuArgs', 'ResourceSkuArgsDict']] sku: The billing information of the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebPubSubReplicaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A class represent a replica resource.

        Uses Azure REST API version 2024-03-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01-preview.

        Other available API versions: 2023-03-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native webpubsub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param WebPubSubReplicaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebPubSubReplicaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 region_endpoint_enabled: Optional[pulumi.Input[builtins.str]] = None,
                 replica_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_name_: Optional[pulumi.Input[builtins.str]] = None,
                 resource_stopped: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input[Union['ResourceSkuArgs', 'ResourceSkuArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebPubSubReplicaArgs.__new__(WebPubSubReplicaArgs)

            __props__.__dict__["location"] = location
            if region_endpoint_enabled is None:
                region_endpoint_enabled = 'Enabled'
            __props__.__dict__["region_endpoint_enabled"] = region_endpoint_enabled
            __props__.__dict__["replica_name"] = replica_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if resource_name_ is None and not opts.urn:
                raise TypeError("Missing required property 'resource_name_'")
            __props__.__dict__["resource_name"] = resource_name_
            if resource_stopped is None:
                resource_stopped = 'false'
            __props__.__dict__["resource_stopped"] = resource_stopped
            __props__.__dict__["sku"] = sku
            __props__.__dict__["tags"] = tags
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:webpubsub/v20230301preview:WebPubSubReplica"), pulumi.Alias(type_="azure-native:webpubsub/v20230601preview:WebPubSubReplica"), pulumi.Alias(type_="azure-native:webpubsub/v20230801preview:WebPubSubReplica"), pulumi.Alias(type_="azure-native:webpubsub/v20240101preview:WebPubSubReplica"), pulumi.Alias(type_="azure-native:webpubsub/v20240301:WebPubSubReplica"), pulumi.Alias(type_="azure-native:webpubsub/v20240401preview:WebPubSubReplica"), pulumi.Alias(type_="azure-native:webpubsub/v20240801preview:WebPubSubReplica"), pulumi.Alias(type_="azure-native:webpubsub/v20241001preview:WebPubSubReplica"), pulumi.Alias(type_="azure-native:webpubsub/v20250101preview:WebPubSubReplica")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(WebPubSubReplica, __self__).__init__(
            'azure-native:webpubsub:WebPubSubReplica',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WebPubSubReplica':
        """
        Get an existing WebPubSubReplica resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WebPubSubReplicaArgs.__new__(WebPubSubReplicaArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["region_endpoint_enabled"] = None
        __props__.__dict__["resource_stopped"] = None
        __props__.__dict__["sku"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return WebPubSubReplica(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

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
    @pulumi.getter(name="regionEndpointEnabled")
    def region_endpoint_enabled(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Enable or disable the regional endpoint. Default to "Enabled".
        When it's Disabled, new connections will not be routed to this endpoint, however existing connections will not be affected.
        """
        return pulumi.get(self, "region_endpoint_enabled")

    @property
    @pulumi.getter(name="resourceStopped")
    def resource_stopped(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Stop or start the resource.  Default to "false".
        When it's true, the data plane of the resource is shutdown.
        When it's false, the data plane of the resource is started.
        """
        return pulumi.get(self, "resource_stopped")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.ResourceSkuResponse']]:
        """
        The billing information of the resource.
        """
        return pulumi.get(self, "sku")

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

