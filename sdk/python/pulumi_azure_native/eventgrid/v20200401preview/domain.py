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
from ._inputs import *

__all__ = ['DomainArgs', 'Domain']

@pulumi.input_type
class DomainArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[str],
                 domain_name: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input['IdentityInfoArgs']] = None,
                 inbound_ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input['InboundIpRuleArgs']]]] = None,
                 input_schema: Optional[pulumi.Input[Union[str, 'InputSchema']]] = None,
                 input_schema_mapping: Optional[pulumi.Input['JsonInputSchemaMappingArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 private_endpoint_connections: Optional[pulumi.Input[Sequence[pulumi.Input['PrivateEndpointConnectionArgs']]]] = None,
                 public_network_access: Optional[pulumi.Input[Union[str, 'PublicNetworkAccess']]] = None,
                 sku: Optional[pulumi.Input['ResourceSkuArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Domain resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the user's subscription.
        :param pulumi.Input[str] domain_name: Name of the domain.
        :param pulumi.Input['IdentityInfoArgs'] identity: Identity information for the resource.
        :param pulumi.Input[Sequence[pulumi.Input['InboundIpRuleArgs']]] inbound_ip_rules: This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
        :param pulumi.Input[Union[str, 'InputSchema']] input_schema: This determines the format that Event Grid should expect for incoming events published to the domain.
        :param pulumi.Input['JsonInputSchemaMappingArgs'] input_schema_mapping: Information about the InputSchemaMapping which specified the info about mapping event payload.
        :param pulumi.Input[str] location: Location of the resource.
        :param pulumi.Input[Sequence[pulumi.Input['PrivateEndpointConnectionArgs']]] private_endpoint_connections: List of private endpoint connections.
        :param pulumi.Input[Union[str, 'PublicNetworkAccess']] public_network_access: This determines if traffic is allowed over public network. By default it is enabled. 
               You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
        :param pulumi.Input['ResourceSkuArgs'] sku: The Sku pricing tier for the domain.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Tags of the resource.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if inbound_ip_rules is not None:
            pulumi.set(__self__, "inbound_ip_rules", inbound_ip_rules)
        if input_schema is None:
            input_schema = 'EventGridSchema'
        if input_schema is not None:
            pulumi.set(__self__, "input_schema", input_schema)
        if input_schema_mapping is not None:
            pulumi.set(__self__, "input_schema_mapping", input_schema_mapping)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if private_endpoint_connections is not None:
            pulumi.set(__self__, "private_endpoint_connections", private_endpoint_connections)
        if public_network_access is None:
            public_network_access = 'Enabled'
        if public_network_access is not None:
            pulumi.set(__self__, "public_network_access", public_network_access)
        if sku is not None:
            pulumi.set(__self__, "sku", sku)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group within the user's subscription.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the domain.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['IdentityInfoArgs']]:
        """
        Identity information for the resource.
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['IdentityInfoArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter(name="inboundIpRules")
    def inbound_ip_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InboundIpRuleArgs']]]]:
        """
        This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
        """
        return pulumi.get(self, "inbound_ip_rules")

    @inbound_ip_rules.setter
    def inbound_ip_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InboundIpRuleArgs']]]]):
        pulumi.set(self, "inbound_ip_rules", value)

    @property
    @pulumi.getter(name="inputSchema")
    def input_schema(self) -> Optional[pulumi.Input[Union[str, 'InputSchema']]]:
        """
        This determines the format that Event Grid should expect for incoming events published to the domain.
        """
        return pulumi.get(self, "input_schema")

    @input_schema.setter
    def input_schema(self, value: Optional[pulumi.Input[Union[str, 'InputSchema']]]):
        pulumi.set(self, "input_schema", value)

    @property
    @pulumi.getter(name="inputSchemaMapping")
    def input_schema_mapping(self) -> Optional[pulumi.Input['JsonInputSchemaMappingArgs']]:
        """
        Information about the InputSchemaMapping which specified the info about mapping event payload.
        """
        return pulumi.get(self, "input_schema_mapping")

    @input_schema_mapping.setter
    def input_schema_mapping(self, value: Optional[pulumi.Input['JsonInputSchemaMappingArgs']]):
        pulumi.set(self, "input_schema_mapping", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Location of the resource.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="privateEndpointConnections")
    def private_endpoint_connections(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PrivateEndpointConnectionArgs']]]]:
        """
        List of private endpoint connections.
        """
        return pulumi.get(self, "private_endpoint_connections")

    @private_endpoint_connections.setter
    def private_endpoint_connections(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PrivateEndpointConnectionArgs']]]]):
        pulumi.set(self, "private_endpoint_connections", value)

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> Optional[pulumi.Input[Union[str, 'PublicNetworkAccess']]]:
        """
        This determines if traffic is allowed over public network. By default it is enabled. 
        You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
        """
        return pulumi.get(self, "public_network_access")

    @public_network_access.setter
    def public_network_access(self, value: Optional[pulumi.Input[Union[str, 'PublicNetworkAccess']]]):
        pulumi.set(self, "public_network_access", value)

    @property
    @pulumi.getter
    def sku(self) -> Optional[pulumi.Input['ResourceSkuArgs']]:
        """
        The Sku pricing tier for the domain.
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: Optional[pulumi.Input['ResourceSkuArgs']]):
        pulumi.set(self, "sku", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Tags of the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class Domain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input[Union['IdentityInfoArgs', 'IdentityInfoArgsDict']]] = None,
                 inbound_ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['InboundIpRuleArgs', 'InboundIpRuleArgsDict']]]]] = None,
                 input_schema: Optional[pulumi.Input[Union[str, 'InputSchema']]] = None,
                 input_schema_mapping: Optional[pulumi.Input[Union['JsonInputSchemaMappingArgs', 'JsonInputSchemaMappingArgsDict']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 private_endpoint_connections: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PrivateEndpointConnectionArgs', 'PrivateEndpointConnectionArgsDict']]]]] = None,
                 public_network_access: Optional[pulumi.Input[Union[str, 'PublicNetworkAccess']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[Union['ResourceSkuArgs', 'ResourceSkuArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        EventGrid Domain.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name: Name of the domain.
        :param pulumi.Input[Union['IdentityInfoArgs', 'IdentityInfoArgsDict']] identity: Identity information for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['InboundIpRuleArgs', 'InboundIpRuleArgsDict']]]] inbound_ip_rules: This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
        :param pulumi.Input[Union[str, 'InputSchema']] input_schema: This determines the format that Event Grid should expect for incoming events published to the domain.
        :param pulumi.Input[Union['JsonInputSchemaMappingArgs', 'JsonInputSchemaMappingArgsDict']] input_schema_mapping: Information about the InputSchemaMapping which specified the info about mapping event payload.
        :param pulumi.Input[str] location: Location of the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['PrivateEndpointConnectionArgs', 'PrivateEndpointConnectionArgsDict']]]] private_endpoint_connections: List of private endpoint connections.
        :param pulumi.Input[Union[str, 'PublicNetworkAccess']] public_network_access: This determines if traffic is allowed over public network. By default it is enabled. 
               You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the user's subscription.
        :param pulumi.Input[Union['ResourceSkuArgs', 'ResourceSkuArgsDict']] sku: The Sku pricing tier for the domain.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Tags of the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        EventGrid Domain.

        :param str resource_name: The name of the resource.
        :param DomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input[Union['IdentityInfoArgs', 'IdentityInfoArgsDict']]] = None,
                 inbound_ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['InboundIpRuleArgs', 'InboundIpRuleArgsDict']]]]] = None,
                 input_schema: Optional[pulumi.Input[Union[str, 'InputSchema']]] = None,
                 input_schema_mapping: Optional[pulumi.Input[Union['JsonInputSchemaMappingArgs', 'JsonInputSchemaMappingArgsDict']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 private_endpoint_connections: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PrivateEndpointConnectionArgs', 'PrivateEndpointConnectionArgsDict']]]]] = None,
                 public_network_access: Optional[pulumi.Input[Union[str, 'PublicNetworkAccess']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[Union['ResourceSkuArgs', 'ResourceSkuArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DomainArgs.__new__(DomainArgs)

            __props__.__dict__["domain_name"] = domain_name
            __props__.__dict__["identity"] = identity
            __props__.__dict__["inbound_ip_rules"] = inbound_ip_rules
            if input_schema is None:
                input_schema = 'EventGridSchema'
            __props__.__dict__["input_schema"] = input_schema
            __props__.__dict__["input_schema_mapping"] = input_schema_mapping
            __props__.__dict__["location"] = location
            __props__.__dict__["private_endpoint_connections"] = private_endpoint_connections
            if public_network_access is None:
                public_network_access = 'Enabled'
            __props__.__dict__["public_network_access"] = public_network_access
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["sku"] = sku
            __props__.__dict__["tags"] = tags
            __props__.__dict__["endpoint"] = None
            __props__.__dict__["metric_resource_id"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:eventgrid/v20180915preview:Domain"), pulumi.Alias(type_="azure-native:eventgrid/v20190201preview:Domain"), pulumi.Alias(type_="azure-native:eventgrid/v20190601:Domain"), pulumi.Alias(type_="azure-native:eventgrid/v20200101preview:Domain"), pulumi.Alias(type_="azure-native:eventgrid/v20200601:Domain"), pulumi.Alias(type_="azure-native:eventgrid/v20201015preview:Domain"), pulumi.Alias(type_="azure-native:eventgrid/v20210601preview:Domain"), pulumi.Alias(type_="azure-native:eventgrid/v20211015preview:Domain"), pulumi.Alias(type_="azure-native:eventgrid/v20211201:Domain"), pulumi.Alias(type_="azure-native:eventgrid/v20220615:Domain"), pulumi.Alias(type_="azure-native:eventgrid/v20230601preview:Domain"), pulumi.Alias(type_="azure-native:eventgrid/v20231215preview:Domain"), pulumi.Alias(type_="azure-native:eventgrid/v20240601preview:Domain"), pulumi.Alias(type_="azure-native:eventgrid/v20241215preview:Domain"), pulumi.Alias(type_="azure-native:eventgrid/v20250215:Domain"), pulumi.Alias(type_="azure-native:eventgrid:Domain")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Domain, __self__).__init__(
            'azure-native:eventgrid/v20200401preview:Domain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Domain':
        """
        Get an existing Domain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DomainArgs.__new__(DomainArgs)

        __props__.__dict__["endpoint"] = None
        __props__.__dict__["identity"] = None
        __props__.__dict__["inbound_ip_rules"] = None
        __props__.__dict__["input_schema"] = None
        __props__.__dict__["input_schema_mapping"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["metric_resource_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["private_endpoint_connections"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["public_network_access"] = None
        __props__.__dict__["sku"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return Domain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        Endpoint for the domain.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.IdentityInfoResponse']]:
        """
        Identity information for the resource.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="inboundIpRules")
    def inbound_ip_rules(self) -> pulumi.Output[Optional[Sequence['outputs.InboundIpRuleResponse']]]:
        """
        This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
        """
        return pulumi.get(self, "inbound_ip_rules")

    @property
    @pulumi.getter(name="inputSchema")
    def input_schema(self) -> pulumi.Output[Optional[str]]:
        """
        This determines the format that Event Grid should expect for incoming events published to the domain.
        """
        return pulumi.get(self, "input_schema")

    @property
    @pulumi.getter(name="inputSchemaMapping")
    def input_schema_mapping(self) -> pulumi.Output[Optional['outputs.JsonInputSchemaMappingResponse']]:
        """
        Information about the InputSchemaMapping which specified the info about mapping event payload.
        """
        return pulumi.get(self, "input_schema_mapping")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="metricResourceId")
    def metric_resource_id(self) -> pulumi.Output[str]:
        """
        Metric resource id for the domain.
        """
        return pulumi.get(self, "metric_resource_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateEndpointConnections")
    def private_endpoint_connections(self) -> pulumi.Output[Optional[Sequence['outputs.PrivateEndpointConnectionResponse']]]:
        """
        List of private endpoint connections.
        """
        return pulumi.get(self, "private_endpoint_connections")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Provisioning state of the domain.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> pulumi.Output[Optional[str]]:
        """
        This determines if traffic is allowed over public network. By default it is enabled. 
        You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
        """
        return pulumi.get(self, "public_network_access")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.ResourceSkuResponse']]:
        """
        The Sku pricing tier for the domain.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of the resource.
        """
        return pulumi.get(self, "type")

