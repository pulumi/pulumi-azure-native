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
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['NetworkVirtualApplianceConnectionArgs', 'NetworkVirtualApplianceConnection']

@pulumi.input_type
class NetworkVirtualApplianceConnectionArgs:
    def __init__(__self__, *,
                 network_virtual_appliance_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 asn: Optional[pulumi.Input[float]] = None,
                 bgp_peer_address: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 connection_name: Optional[pulumi.Input[str]] = None,
                 enable_internet_security: Optional[pulumi.Input[bool]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 routing_configuration: Optional[pulumi.Input['RoutingConfigurationArgs']] = None,
                 tunnel_identifier: Optional[pulumi.Input[float]] = None):
        """
        The set of arguments for constructing a NetworkVirtualApplianceConnection resource.
        :param pulumi.Input[str] network_virtual_appliance_name: The name of the Network Virtual Appliance.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[float] asn: Network Virtual Appliance ASN.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] bgp_peer_address: List of bgpPeerAddresses for the NVA instances
        :param pulumi.Input[str] connection_name: The name of the NVA connection.
        :param pulumi.Input[bool] enable_internet_security: Enable internet security.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: The name of the resource.
        :param pulumi.Input['RoutingConfigurationArgs'] routing_configuration: The Routing Configuration indicating the associated and propagated route tables on this connection.
        :param pulumi.Input[float] tunnel_identifier: Unique identifier for the connection.
        """
        pulumi.set(__self__, "network_virtual_appliance_name", network_virtual_appliance_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if asn is not None:
            pulumi.set(__self__, "asn", asn)
        if bgp_peer_address is not None:
            pulumi.set(__self__, "bgp_peer_address", bgp_peer_address)
        if connection_name is not None:
            pulumi.set(__self__, "connection_name", connection_name)
        if enable_internet_security is not None:
            pulumi.set(__self__, "enable_internet_security", enable_internet_security)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if routing_configuration is not None:
            pulumi.set(__self__, "routing_configuration", routing_configuration)
        if tunnel_identifier is not None:
            pulumi.set(__self__, "tunnel_identifier", tunnel_identifier)

    @property
    @pulumi.getter(name="networkVirtualApplianceName")
    def network_virtual_appliance_name(self) -> pulumi.Input[str]:
        """
        The name of the Network Virtual Appliance.
        """
        return pulumi.get(self, "network_virtual_appliance_name")

    @network_virtual_appliance_name.setter
    def network_virtual_appliance_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_virtual_appliance_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def asn(self) -> Optional[pulumi.Input[float]]:
        """
        Network Virtual Appliance ASN.
        """
        return pulumi.get(self, "asn")

    @asn.setter
    def asn(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "asn", value)

    @property
    @pulumi.getter(name="bgpPeerAddress")
    def bgp_peer_address(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of bgpPeerAddresses for the NVA instances
        """
        return pulumi.get(self, "bgp_peer_address")

    @bgp_peer_address.setter
    def bgp_peer_address(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "bgp_peer_address", value)

    @property
    @pulumi.getter(name="connectionName")
    def connection_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the NVA connection.
        """
        return pulumi.get(self, "connection_name")

    @connection_name.setter
    def connection_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_name", value)

    @property
    @pulumi.getter(name="enableInternetSecurity")
    def enable_internet_security(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable internet security.
        """
        return pulumi.get(self, "enable_internet_security")

    @enable_internet_security.setter
    def enable_internet_security(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_internet_security", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="routingConfiguration")
    def routing_configuration(self) -> Optional[pulumi.Input['RoutingConfigurationArgs']]:
        """
        The Routing Configuration indicating the associated and propagated route tables on this connection.
        """
        return pulumi.get(self, "routing_configuration")

    @routing_configuration.setter
    def routing_configuration(self, value: Optional[pulumi.Input['RoutingConfigurationArgs']]):
        pulumi.set(self, "routing_configuration", value)

    @property
    @pulumi.getter(name="tunnelIdentifier")
    def tunnel_identifier(self) -> Optional[pulumi.Input[float]]:
        """
        Unique identifier for the connection.
        """
        return pulumi.get(self, "tunnel_identifier")

    @tunnel_identifier.setter
    def tunnel_identifier(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "tunnel_identifier", value)


class NetworkVirtualApplianceConnection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asn: Optional[pulumi.Input[float]] = None,
                 bgp_peer_address: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 connection_name: Optional[pulumi.Input[str]] = None,
                 enable_internet_security: Optional[pulumi.Input[bool]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_virtual_appliance_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 routing_configuration: Optional[pulumi.Input[Union['RoutingConfigurationArgs', 'RoutingConfigurationArgsDict']]] = None,
                 tunnel_identifier: Optional[pulumi.Input[float]] = None,
                 __props__=None):
        """
        NetworkVirtualApplianceConnection resource.

        Uses Azure REST API version 2023-06-01.

        Other available API versions: 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] asn: Network Virtual Appliance ASN.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] bgp_peer_address: List of bgpPeerAddresses for the NVA instances
        :param pulumi.Input[str] connection_name: The name of the NVA connection.
        :param pulumi.Input[bool] enable_internet_security: Enable internet security.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: The name of the resource.
        :param pulumi.Input[str] network_virtual_appliance_name: The name of the Network Virtual Appliance.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Union['RoutingConfigurationArgs', 'RoutingConfigurationArgsDict']] routing_configuration: The Routing Configuration indicating the associated and propagated route tables on this connection.
        :param pulumi.Input[float] tunnel_identifier: Unique identifier for the connection.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetworkVirtualApplianceConnectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        NetworkVirtualApplianceConnection resource.

        Uses Azure REST API version 2023-06-01.

        Other available API versions: 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.

        :param str resource_name: The name of the resource.
        :param NetworkVirtualApplianceConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkVirtualApplianceConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asn: Optional[pulumi.Input[float]] = None,
                 bgp_peer_address: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 connection_name: Optional[pulumi.Input[str]] = None,
                 enable_internet_security: Optional[pulumi.Input[bool]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_virtual_appliance_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 routing_configuration: Optional[pulumi.Input[Union['RoutingConfigurationArgs', 'RoutingConfigurationArgsDict']]] = None,
                 tunnel_identifier: Optional[pulumi.Input[float]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkVirtualApplianceConnectionArgs.__new__(NetworkVirtualApplianceConnectionArgs)

            __props__.__dict__["asn"] = asn
            __props__.__dict__["bgp_peer_address"] = bgp_peer_address
            __props__.__dict__["connection_name"] = connection_name
            __props__.__dict__["enable_internet_security"] = enable_internet_security
            __props__.__dict__["id"] = id
            __props__.__dict__["name"] = name
            if network_virtual_appliance_name is None and not opts.urn:
                raise TypeError("Missing required property 'network_virtual_appliance_name'")
            __props__.__dict__["network_virtual_appliance_name"] = network_virtual_appliance_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["routing_configuration"] = routing_configuration
            __props__.__dict__["tunnel_identifier"] = tunnel_identifier
            __props__.__dict__["provisioning_state"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:network/v20230601:NetworkVirtualApplianceConnection"), pulumi.Alias(type_="azure-native:network/v20230901:NetworkVirtualApplianceConnection"), pulumi.Alias(type_="azure-native:network/v20231101:NetworkVirtualApplianceConnection"), pulumi.Alias(type_="azure-native:network/v20240101:NetworkVirtualApplianceConnection"), pulumi.Alias(type_="azure-native:network/v20240301:NetworkVirtualApplianceConnection"), pulumi.Alias(type_="azure-native:network/v20240501:NetworkVirtualApplianceConnection")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(NetworkVirtualApplianceConnection, __self__).__init__(
            'azure-native:network:NetworkVirtualApplianceConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NetworkVirtualApplianceConnection':
        """
        Get an existing NetworkVirtualApplianceConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NetworkVirtualApplianceConnectionArgs.__new__(NetworkVirtualApplianceConnectionArgs)

        __props__.__dict__["asn"] = None
        __props__.__dict__["bgp_peer_address"] = None
        __props__.__dict__["enable_internet_security"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["routing_configuration"] = None
        __props__.__dict__["tunnel_identifier"] = None
        return NetworkVirtualApplianceConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def asn(self) -> pulumi.Output[Optional[float]]:
        """
        Network Virtual Appliance ASN.
        """
        return pulumi.get(self, "asn")

    @property
    @pulumi.getter(name="bgpPeerAddress")
    def bgp_peer_address(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of bgpPeerAddresses for the NVA instances
        """
        return pulumi.get(self, "bgp_peer_address")

    @property
    @pulumi.getter(name="enableInternetSecurity")
    def enable_internet_security(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable internet security.
        """
        return pulumi.get(self, "enable_internet_security")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the NetworkVirtualApplianceConnection resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="routingConfiguration")
    def routing_configuration(self) -> pulumi.Output[Optional['outputs.RoutingConfigurationResponse']]:
        """
        The Routing Configuration indicating the associated and propagated route tables on this connection.
        """
        return pulumi.get(self, "routing_configuration")

    @property
    @pulumi.getter(name="tunnelIdentifier")
    def tunnel_identifier(self) -> pulumi.Output[Optional[float]]:
        """
        Unique identifier for the connection.
        """
        return pulumi.get(self, "tunnel_identifier")

