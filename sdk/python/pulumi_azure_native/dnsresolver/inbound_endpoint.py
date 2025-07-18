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

__all__ = ['InboundEndpointArgs', 'InboundEndpoint']

@pulumi.input_type
class InboundEndpointArgs:
    def __init__(__self__, *,
                 dns_resolver_name: pulumi.Input[builtins.str],
                 ip_configurations: pulumi.Input[Sequence[pulumi.Input['IpConfigurationArgs']]],
                 resource_group_name: pulumi.Input[builtins.str],
                 inbound_endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a InboundEndpoint resource.
        :param pulumi.Input[builtins.str] dns_resolver_name: The name of the DNS resolver.
        :param pulumi.Input[Sequence[pulumi.Input['IpConfigurationArgs']]] ip_configurations: IP configurations for the inbound endpoint.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] inbound_endpoint_name: The name of the inbound endpoint for the DNS resolver.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        pulumi.set(__self__, "dns_resolver_name", dns_resolver_name)
        pulumi.set(__self__, "ip_configurations", ip_configurations)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if inbound_endpoint_name is not None:
            pulumi.set(__self__, "inbound_endpoint_name", inbound_endpoint_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="dnsResolverName")
    def dns_resolver_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the DNS resolver.
        """
        return pulumi.get(self, "dns_resolver_name")

    @dns_resolver_name.setter
    def dns_resolver_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "dns_resolver_name", value)

    @property
    @pulumi.getter(name="ipConfigurations")
    def ip_configurations(self) -> pulumi.Input[Sequence[pulumi.Input['IpConfigurationArgs']]]:
        """
        IP configurations for the inbound endpoint.
        """
        return pulumi.get(self, "ip_configurations")

    @ip_configurations.setter
    def ip_configurations(self, value: pulumi.Input[Sequence[pulumi.Input['IpConfigurationArgs']]]):
        pulumi.set(self, "ip_configurations", value)

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
    @pulumi.getter(name="inboundEndpointName")
    def inbound_endpoint_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the inbound endpoint for the DNS resolver.
        """
        return pulumi.get(self, "inbound_endpoint_name")

    @inbound_endpoint_name.setter
    def inbound_endpoint_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "inbound_endpoint_name", value)

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


@pulumi.type_token("azure-native:dnsresolver:InboundEndpoint")
class InboundEndpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dns_resolver_name: Optional[pulumi.Input[builtins.str]] = None,
                 inbound_endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                 ip_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IpConfigurationArgs', 'IpConfigurationArgsDict']]]]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Describes an inbound endpoint for a DNS resolver.

        Uses Azure REST API version 2023-07-01-preview.

        Other available API versions: 2020-04-01-preview, 2022-07-01, 2025-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dnsresolver [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] dns_resolver_name: The name of the DNS resolver.
        :param pulumi.Input[builtins.str] inbound_endpoint_name: The name of the inbound endpoint for the DNS resolver.
        :param pulumi.Input[Sequence[pulumi.Input[Union['IpConfigurationArgs', 'IpConfigurationArgsDict']]]] ip_configurations: IP configurations for the inbound endpoint.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InboundEndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Describes an inbound endpoint for a DNS resolver.

        Uses Azure REST API version 2023-07-01-preview.

        Other available API versions: 2020-04-01-preview, 2022-07-01, 2025-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dnsresolver [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param InboundEndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InboundEndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dns_resolver_name: Optional[pulumi.Input[builtins.str]] = None,
                 inbound_endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                 ip_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IpConfigurationArgs', 'IpConfigurationArgsDict']]]]] = None,
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
            __props__ = InboundEndpointArgs.__new__(InboundEndpointArgs)

            if dns_resolver_name is None and not opts.urn:
                raise TypeError("Missing required property 'dns_resolver_name'")
            __props__.__dict__["dns_resolver_name"] = dns_resolver_name
            __props__.__dict__["inbound_endpoint_name"] = inbound_endpoint_name
            if ip_configurations is None and not opts.urn:
                raise TypeError("Missing required property 'ip_configurations'")
            __props__.__dict__["ip_configurations"] = ip_configurations
            __props__.__dict__["location"] = location
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["resource_guid"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:dnsresolver/v20200401preview:InboundEndpoint"), pulumi.Alias(type_="azure-native:dnsresolver/v20220701:InboundEndpoint"), pulumi.Alias(type_="azure-native:dnsresolver/v20230701preview:InboundEndpoint"), pulumi.Alias(type_="azure-native:dnsresolver/v20250501:InboundEndpoint"), pulumi.Alias(type_="azure-native:network/v20200401preview:InboundEndpoint"), pulumi.Alias(type_="azure-native:network/v20220701:InboundEndpoint"), pulumi.Alias(type_="azure-native:network/v20230701preview:InboundEndpoint"), pulumi.Alias(type_="azure-native:network:InboundEndpoint")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(InboundEndpoint, __self__).__init__(
            'azure-native:dnsresolver:InboundEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'InboundEndpoint':
        """
        Get an existing InboundEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = InboundEndpointArgs.__new__(InboundEndpointArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["ip_configurations"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["resource_guid"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return InboundEndpoint(resource_name, opts=opts, __props__=__props__)

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
        ETag of the inbound endpoint.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="ipConfigurations")
    def ip_configurations(self) -> pulumi.Output[Sequence['outputs.IpConfigurationResponse']]:
        """
        IP configurations for the inbound endpoint.
        """
        return pulumi.get(self, "ip_configurations")

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
        The current provisioning state of the inbound endpoint. This is a read-only property and any attempt to set this value will be ignored.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceGuid")
    def resource_guid(self) -> pulumi.Output[builtins.str]:
        """
        The resourceGuid property of the inbound endpoint resource.
        """
        return pulumi.get(self, "resource_guid")

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

