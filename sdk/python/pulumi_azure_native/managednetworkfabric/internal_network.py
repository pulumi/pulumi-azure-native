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

__all__ = ['InternalNetworkArgs', 'InternalNetwork']

@pulumi.input_type
class InternalNetworkArgs:
    def __init__(__self__, *,
                 l3_isolation_domain_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 vlan_id: pulumi.Input[builtins.int],
                 annotation: Optional[pulumi.Input[builtins.str]] = None,
                 bgp_configuration: Optional[pulumi.Input['InternalNetworkPropertiesBgpConfigurationArgs']] = None,
                 connected_i_pv4_subnets: Optional[pulumi.Input[Sequence[pulumi.Input['ConnectedSubnetArgs']]]] = None,
                 connected_i_pv6_subnets: Optional[pulumi.Input[Sequence[pulumi.Input['ConnectedSubnetArgs']]]] = None,
                 egress_acl_id: Optional[pulumi.Input[builtins.str]] = None,
                 export_route_policy: Optional[pulumi.Input['ExportRoutePolicyArgs']] = None,
                 export_route_policy_id: Optional[pulumi.Input[builtins.str]] = None,
                 extension: Optional[pulumi.Input[Union[builtins.str, 'Extension']]] = None,
                 import_route_policy: Optional[pulumi.Input['ImportRoutePolicyArgs']] = None,
                 import_route_policy_id: Optional[pulumi.Input[builtins.str]] = None,
                 ingress_acl_id: Optional[pulumi.Input[builtins.str]] = None,
                 internal_network_name: Optional[pulumi.Input[builtins.str]] = None,
                 is_monitoring_enabled: Optional[pulumi.Input[Union[builtins.str, 'IsMonitoringEnabled']]] = None,
                 mtu: Optional[pulumi.Input[builtins.int]] = None,
                 static_route_configuration: Optional[pulumi.Input['InternalNetworkPropertiesStaticRouteConfigurationArgs']] = None):
        """
        The set of arguments for constructing a InternalNetwork resource.
        :param pulumi.Input[builtins.str] l3_isolation_domain_name: Name of the L3 Isolation Domain.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.int] vlan_id: Vlan identifier. Example: 1001.
        :param pulumi.Input[builtins.str] annotation: Switch configuration description.
        :param pulumi.Input['InternalNetworkPropertiesBgpConfigurationArgs'] bgp_configuration: BGP configuration properties.
        :param pulumi.Input[Sequence[pulumi.Input['ConnectedSubnetArgs']]] connected_i_pv4_subnets: List of Connected IPv4 Subnets.
        :param pulumi.Input[Sequence[pulumi.Input['ConnectedSubnetArgs']]] connected_i_pv6_subnets: List of connected IPv6 Subnets.
        :param pulumi.Input[builtins.str] egress_acl_id: Egress Acl. ARM resource ID of Access Control Lists.
        :param pulumi.Input['ExportRoutePolicyArgs'] export_route_policy: Export Route Policy either IPv4 or IPv6.
        :param pulumi.Input[builtins.str] export_route_policy_id: ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
        :param pulumi.Input[Union[builtins.str, 'Extension']] extension: Extension. Example: NoExtension | NPB.
        :param pulumi.Input['ImportRoutePolicyArgs'] import_route_policy: Import Route Policy either IPv4 or IPv6.
        :param pulumi.Input[builtins.str] import_route_policy_id: ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
        :param pulumi.Input[builtins.str] ingress_acl_id: Ingress Acl. ARM resource ID of Access Control Lists.
        :param pulumi.Input[builtins.str] internal_network_name: Name of the Internal Network.
        :param pulumi.Input[Union[builtins.str, 'IsMonitoringEnabled']] is_monitoring_enabled: To check whether monitoring of internal network is enabled or not.
        :param pulumi.Input[builtins.int] mtu: Maximum transmission unit. Default value is 1500.
        :param pulumi.Input['InternalNetworkPropertiesStaticRouteConfigurationArgs'] static_route_configuration: Static Route Configuration properties.
        """
        pulumi.set(__self__, "l3_isolation_domain_name", l3_isolation_domain_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "vlan_id", vlan_id)
        if annotation is not None:
            pulumi.set(__self__, "annotation", annotation)
        if bgp_configuration is not None:
            pulumi.set(__self__, "bgp_configuration", bgp_configuration)
        if connected_i_pv4_subnets is not None:
            pulumi.set(__self__, "connected_i_pv4_subnets", connected_i_pv4_subnets)
        if connected_i_pv6_subnets is not None:
            pulumi.set(__self__, "connected_i_pv6_subnets", connected_i_pv6_subnets)
        if egress_acl_id is not None:
            pulumi.set(__self__, "egress_acl_id", egress_acl_id)
        if export_route_policy is not None:
            pulumi.set(__self__, "export_route_policy", export_route_policy)
        if export_route_policy_id is not None:
            pulumi.set(__self__, "export_route_policy_id", export_route_policy_id)
        if extension is None:
            extension = 'NoExtension'
        if extension is not None:
            pulumi.set(__self__, "extension", extension)
        if import_route_policy is not None:
            pulumi.set(__self__, "import_route_policy", import_route_policy)
        if import_route_policy_id is not None:
            pulumi.set(__self__, "import_route_policy_id", import_route_policy_id)
        if ingress_acl_id is not None:
            pulumi.set(__self__, "ingress_acl_id", ingress_acl_id)
        if internal_network_name is not None:
            pulumi.set(__self__, "internal_network_name", internal_network_name)
        if is_monitoring_enabled is None:
            is_monitoring_enabled = 'False'
        if is_monitoring_enabled is not None:
            pulumi.set(__self__, "is_monitoring_enabled", is_monitoring_enabled)
        if mtu is None:
            mtu = 1500
        if mtu is not None:
            pulumi.set(__self__, "mtu", mtu)
        if static_route_configuration is not None:
            pulumi.set(__self__, "static_route_configuration", static_route_configuration)

    @property
    @pulumi.getter(name="l3IsolationDomainName")
    def l3_isolation_domain_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the L3 Isolation Domain.
        """
        return pulumi.get(self, "l3_isolation_domain_name")

    @l3_isolation_domain_name.setter
    def l3_isolation_domain_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "l3_isolation_domain_name", value)

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
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> pulumi.Input[builtins.int]:
        """
        Vlan identifier. Example: 1001.
        """
        return pulumi.get(self, "vlan_id")

    @vlan_id.setter
    def vlan_id(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "vlan_id", value)

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
    @pulumi.getter(name="bgpConfiguration")
    def bgp_configuration(self) -> Optional[pulumi.Input['InternalNetworkPropertiesBgpConfigurationArgs']]:
        """
        BGP configuration properties.
        """
        return pulumi.get(self, "bgp_configuration")

    @bgp_configuration.setter
    def bgp_configuration(self, value: Optional[pulumi.Input['InternalNetworkPropertiesBgpConfigurationArgs']]):
        pulumi.set(self, "bgp_configuration", value)

    @property
    @pulumi.getter(name="connectedIPv4Subnets")
    def connected_i_pv4_subnets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ConnectedSubnetArgs']]]]:
        """
        List of Connected IPv4 Subnets.
        """
        return pulumi.get(self, "connected_i_pv4_subnets")

    @connected_i_pv4_subnets.setter
    def connected_i_pv4_subnets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ConnectedSubnetArgs']]]]):
        pulumi.set(self, "connected_i_pv4_subnets", value)

    @property
    @pulumi.getter(name="connectedIPv6Subnets")
    def connected_i_pv6_subnets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ConnectedSubnetArgs']]]]:
        """
        List of connected IPv6 Subnets.
        """
        return pulumi.get(self, "connected_i_pv6_subnets")

    @connected_i_pv6_subnets.setter
    def connected_i_pv6_subnets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ConnectedSubnetArgs']]]]):
        pulumi.set(self, "connected_i_pv6_subnets", value)

    @property
    @pulumi.getter(name="egressAclId")
    def egress_acl_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Egress Acl. ARM resource ID of Access Control Lists.
        """
        return pulumi.get(self, "egress_acl_id")

    @egress_acl_id.setter
    def egress_acl_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "egress_acl_id", value)

    @property
    @pulumi.getter(name="exportRoutePolicy")
    def export_route_policy(self) -> Optional[pulumi.Input['ExportRoutePolicyArgs']]:
        """
        Export Route Policy either IPv4 or IPv6.
        """
        return pulumi.get(self, "export_route_policy")

    @export_route_policy.setter
    def export_route_policy(self, value: Optional[pulumi.Input['ExportRoutePolicyArgs']]):
        pulumi.set(self, "export_route_policy", value)

    @property
    @pulumi.getter(name="exportRoutePolicyId")
    def export_route_policy_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
        """
        return pulumi.get(self, "export_route_policy_id")

    @export_route_policy_id.setter
    def export_route_policy_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "export_route_policy_id", value)

    @property
    @pulumi.getter
    def extension(self) -> Optional[pulumi.Input[Union[builtins.str, 'Extension']]]:
        """
        Extension. Example: NoExtension | NPB.
        """
        return pulumi.get(self, "extension")

    @extension.setter
    def extension(self, value: Optional[pulumi.Input[Union[builtins.str, 'Extension']]]):
        pulumi.set(self, "extension", value)

    @property
    @pulumi.getter(name="importRoutePolicy")
    def import_route_policy(self) -> Optional[pulumi.Input['ImportRoutePolicyArgs']]:
        """
        Import Route Policy either IPv4 or IPv6.
        """
        return pulumi.get(self, "import_route_policy")

    @import_route_policy.setter
    def import_route_policy(self, value: Optional[pulumi.Input['ImportRoutePolicyArgs']]):
        pulumi.set(self, "import_route_policy", value)

    @property
    @pulumi.getter(name="importRoutePolicyId")
    def import_route_policy_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
        """
        return pulumi.get(self, "import_route_policy_id")

    @import_route_policy_id.setter
    def import_route_policy_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "import_route_policy_id", value)

    @property
    @pulumi.getter(name="ingressAclId")
    def ingress_acl_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Ingress Acl. ARM resource ID of Access Control Lists.
        """
        return pulumi.get(self, "ingress_acl_id")

    @ingress_acl_id.setter
    def ingress_acl_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ingress_acl_id", value)

    @property
    @pulumi.getter(name="internalNetworkName")
    def internal_network_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the Internal Network.
        """
        return pulumi.get(self, "internal_network_name")

    @internal_network_name.setter
    def internal_network_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "internal_network_name", value)

    @property
    @pulumi.getter(name="isMonitoringEnabled")
    def is_monitoring_enabled(self) -> Optional[pulumi.Input[Union[builtins.str, 'IsMonitoringEnabled']]]:
        """
        To check whether monitoring of internal network is enabled or not.
        """
        return pulumi.get(self, "is_monitoring_enabled")

    @is_monitoring_enabled.setter
    def is_monitoring_enabled(self, value: Optional[pulumi.Input[Union[builtins.str, 'IsMonitoringEnabled']]]):
        pulumi.set(self, "is_monitoring_enabled", value)

    @property
    @pulumi.getter
    def mtu(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Maximum transmission unit. Default value is 1500.
        """
        return pulumi.get(self, "mtu")

    @mtu.setter
    def mtu(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "mtu", value)

    @property
    @pulumi.getter(name="staticRouteConfiguration")
    def static_route_configuration(self) -> Optional[pulumi.Input['InternalNetworkPropertiesStaticRouteConfigurationArgs']]:
        """
        Static Route Configuration properties.
        """
        return pulumi.get(self, "static_route_configuration")

    @static_route_configuration.setter
    def static_route_configuration(self, value: Optional[pulumi.Input['InternalNetworkPropertiesStaticRouteConfigurationArgs']]):
        pulumi.set(self, "static_route_configuration", value)


@pulumi.type_token("azure-native:managednetworkfabric:InternalNetwork")
class InternalNetwork(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 annotation: Optional[pulumi.Input[builtins.str]] = None,
                 bgp_configuration: Optional[pulumi.Input[Union['InternalNetworkPropertiesBgpConfigurationArgs', 'InternalNetworkPropertiesBgpConfigurationArgsDict']]] = None,
                 connected_i_pv4_subnets: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ConnectedSubnetArgs', 'ConnectedSubnetArgsDict']]]]] = None,
                 connected_i_pv6_subnets: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ConnectedSubnetArgs', 'ConnectedSubnetArgsDict']]]]] = None,
                 egress_acl_id: Optional[pulumi.Input[builtins.str]] = None,
                 export_route_policy: Optional[pulumi.Input[Union['ExportRoutePolicyArgs', 'ExportRoutePolicyArgsDict']]] = None,
                 export_route_policy_id: Optional[pulumi.Input[builtins.str]] = None,
                 extension: Optional[pulumi.Input[Union[builtins.str, 'Extension']]] = None,
                 import_route_policy: Optional[pulumi.Input[Union['ImportRoutePolicyArgs', 'ImportRoutePolicyArgsDict']]] = None,
                 import_route_policy_id: Optional[pulumi.Input[builtins.str]] = None,
                 ingress_acl_id: Optional[pulumi.Input[builtins.str]] = None,
                 internal_network_name: Optional[pulumi.Input[builtins.str]] = None,
                 is_monitoring_enabled: Optional[pulumi.Input[Union[builtins.str, 'IsMonitoringEnabled']]] = None,
                 l3_isolation_domain_name: Optional[pulumi.Input[builtins.str]] = None,
                 mtu: Optional[pulumi.Input[builtins.int]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 static_route_configuration: Optional[pulumi.Input[Union['InternalNetworkPropertiesStaticRouteConfigurationArgs', 'InternalNetworkPropertiesStaticRouteConfigurationArgsDict']]] = None,
                 vlan_id: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        """
        Defines the Internal Network resource.

        Uses Azure REST API version 2023-06-15. In version 2.x of the Azure Native provider, it used API version 2023-02-01-preview.

        Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] annotation: Switch configuration description.
        :param pulumi.Input[Union['InternalNetworkPropertiesBgpConfigurationArgs', 'InternalNetworkPropertiesBgpConfigurationArgsDict']] bgp_configuration: BGP configuration properties.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ConnectedSubnetArgs', 'ConnectedSubnetArgsDict']]]] connected_i_pv4_subnets: List of Connected IPv4 Subnets.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ConnectedSubnetArgs', 'ConnectedSubnetArgsDict']]]] connected_i_pv6_subnets: List of connected IPv6 Subnets.
        :param pulumi.Input[builtins.str] egress_acl_id: Egress Acl. ARM resource ID of Access Control Lists.
        :param pulumi.Input[Union['ExportRoutePolicyArgs', 'ExportRoutePolicyArgsDict']] export_route_policy: Export Route Policy either IPv4 or IPv6.
        :param pulumi.Input[builtins.str] export_route_policy_id: ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
        :param pulumi.Input[Union[builtins.str, 'Extension']] extension: Extension. Example: NoExtension | NPB.
        :param pulumi.Input[Union['ImportRoutePolicyArgs', 'ImportRoutePolicyArgsDict']] import_route_policy: Import Route Policy either IPv4 or IPv6.
        :param pulumi.Input[builtins.str] import_route_policy_id: ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
        :param pulumi.Input[builtins.str] ingress_acl_id: Ingress Acl. ARM resource ID of Access Control Lists.
        :param pulumi.Input[builtins.str] internal_network_name: Name of the Internal Network.
        :param pulumi.Input[Union[builtins.str, 'IsMonitoringEnabled']] is_monitoring_enabled: To check whether monitoring of internal network is enabled or not.
        :param pulumi.Input[builtins.str] l3_isolation_domain_name: Name of the L3 Isolation Domain.
        :param pulumi.Input[builtins.int] mtu: Maximum transmission unit. Default value is 1500.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Union['InternalNetworkPropertiesStaticRouteConfigurationArgs', 'InternalNetworkPropertiesStaticRouteConfigurationArgsDict']] static_route_configuration: Static Route Configuration properties.
        :param pulumi.Input[builtins.int] vlan_id: Vlan identifier. Example: 1001.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InternalNetworkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Defines the Internal Network resource.

        Uses Azure REST API version 2023-06-15. In version 2.x of the Azure Native provider, it used API version 2023-02-01-preview.

        Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param InternalNetworkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InternalNetworkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 annotation: Optional[pulumi.Input[builtins.str]] = None,
                 bgp_configuration: Optional[pulumi.Input[Union['InternalNetworkPropertiesBgpConfigurationArgs', 'InternalNetworkPropertiesBgpConfigurationArgsDict']]] = None,
                 connected_i_pv4_subnets: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ConnectedSubnetArgs', 'ConnectedSubnetArgsDict']]]]] = None,
                 connected_i_pv6_subnets: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ConnectedSubnetArgs', 'ConnectedSubnetArgsDict']]]]] = None,
                 egress_acl_id: Optional[pulumi.Input[builtins.str]] = None,
                 export_route_policy: Optional[pulumi.Input[Union['ExportRoutePolicyArgs', 'ExportRoutePolicyArgsDict']]] = None,
                 export_route_policy_id: Optional[pulumi.Input[builtins.str]] = None,
                 extension: Optional[pulumi.Input[Union[builtins.str, 'Extension']]] = None,
                 import_route_policy: Optional[pulumi.Input[Union['ImportRoutePolicyArgs', 'ImportRoutePolicyArgsDict']]] = None,
                 import_route_policy_id: Optional[pulumi.Input[builtins.str]] = None,
                 ingress_acl_id: Optional[pulumi.Input[builtins.str]] = None,
                 internal_network_name: Optional[pulumi.Input[builtins.str]] = None,
                 is_monitoring_enabled: Optional[pulumi.Input[Union[builtins.str, 'IsMonitoringEnabled']]] = None,
                 l3_isolation_domain_name: Optional[pulumi.Input[builtins.str]] = None,
                 mtu: Optional[pulumi.Input[builtins.int]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 static_route_configuration: Optional[pulumi.Input[Union['InternalNetworkPropertiesStaticRouteConfigurationArgs', 'InternalNetworkPropertiesStaticRouteConfigurationArgsDict']]] = None,
                 vlan_id: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InternalNetworkArgs.__new__(InternalNetworkArgs)

            __props__.__dict__["annotation"] = annotation
            __props__.__dict__["bgp_configuration"] = bgp_configuration
            __props__.__dict__["connected_i_pv4_subnets"] = connected_i_pv4_subnets
            __props__.__dict__["connected_i_pv6_subnets"] = connected_i_pv6_subnets
            __props__.__dict__["egress_acl_id"] = egress_acl_id
            __props__.__dict__["export_route_policy"] = export_route_policy
            __props__.__dict__["export_route_policy_id"] = export_route_policy_id
            if extension is None:
                extension = 'NoExtension'
            __props__.__dict__["extension"] = extension
            __props__.__dict__["import_route_policy"] = import_route_policy
            __props__.__dict__["import_route_policy_id"] = import_route_policy_id
            __props__.__dict__["ingress_acl_id"] = ingress_acl_id
            __props__.__dict__["internal_network_name"] = internal_network_name
            if is_monitoring_enabled is None:
                is_monitoring_enabled = 'False'
            __props__.__dict__["is_monitoring_enabled"] = is_monitoring_enabled
            if l3_isolation_domain_name is None and not opts.urn:
                raise TypeError("Missing required property 'l3_isolation_domain_name'")
            __props__.__dict__["l3_isolation_domain_name"] = l3_isolation_domain_name
            if mtu is None:
                mtu = 1500
            __props__.__dict__["mtu"] = mtu
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["static_route_configuration"] = static_route_configuration
            if vlan_id is None and not opts.urn:
                raise TypeError("Missing required property 'vlan_id'")
            __props__.__dict__["vlan_id"] = vlan_id
            __props__.__dict__["administrative_state"] = None
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["configuration_state"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:managednetworkfabric/v20230201preview:InternalNetwork"), pulumi.Alias(type_="azure-native:managednetworkfabric/v20230615:InternalNetwork")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(InternalNetwork, __self__).__init__(
            'azure-native:managednetworkfabric:InternalNetwork',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'InternalNetwork':
        """
        Get an existing InternalNetwork resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = InternalNetworkArgs.__new__(InternalNetworkArgs)

        __props__.__dict__["administrative_state"] = None
        __props__.__dict__["annotation"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["bgp_configuration"] = None
        __props__.__dict__["configuration_state"] = None
        __props__.__dict__["connected_i_pv4_subnets"] = None
        __props__.__dict__["connected_i_pv6_subnets"] = None
        __props__.__dict__["egress_acl_id"] = None
        __props__.__dict__["export_route_policy"] = None
        __props__.__dict__["export_route_policy_id"] = None
        __props__.__dict__["extension"] = None
        __props__.__dict__["import_route_policy"] = None
        __props__.__dict__["import_route_policy_id"] = None
        __props__.__dict__["ingress_acl_id"] = None
        __props__.__dict__["is_monitoring_enabled"] = None
        __props__.__dict__["mtu"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["static_route_configuration"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["vlan_id"] = None
        return InternalNetwork(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter(name="bgpConfiguration")
    def bgp_configuration(self) -> pulumi.Output[Optional['outputs.InternalNetworkPropertiesResponseBgpConfiguration']]:
        """
        BGP configuration properties.
        """
        return pulumi.get(self, "bgp_configuration")

    @property
    @pulumi.getter(name="configurationState")
    def configuration_state(self) -> pulumi.Output[builtins.str]:
        """
        Configuration state of the resource.
        """
        return pulumi.get(self, "configuration_state")

    @property
    @pulumi.getter(name="connectedIPv4Subnets")
    def connected_i_pv4_subnets(self) -> pulumi.Output[Optional[Sequence['outputs.ConnectedSubnetResponse']]]:
        """
        List of Connected IPv4 Subnets.
        """
        return pulumi.get(self, "connected_i_pv4_subnets")

    @property
    @pulumi.getter(name="connectedIPv6Subnets")
    def connected_i_pv6_subnets(self) -> pulumi.Output[Optional[Sequence['outputs.ConnectedSubnetResponse']]]:
        """
        List of connected IPv6 Subnets.
        """
        return pulumi.get(self, "connected_i_pv6_subnets")

    @property
    @pulumi.getter(name="egressAclId")
    def egress_acl_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Egress Acl. ARM resource ID of Access Control Lists.
        """
        return pulumi.get(self, "egress_acl_id")

    @property
    @pulumi.getter(name="exportRoutePolicy")
    def export_route_policy(self) -> pulumi.Output[Optional['outputs.ExportRoutePolicyResponse']]:
        """
        Export Route Policy either IPv4 or IPv6.
        """
        return pulumi.get(self, "export_route_policy")

    @property
    @pulumi.getter(name="exportRoutePolicyId")
    def export_route_policy_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
        """
        return pulumi.get(self, "export_route_policy_id")

    @property
    @pulumi.getter
    def extension(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Extension. Example: NoExtension | NPB.
        """
        return pulumi.get(self, "extension")

    @property
    @pulumi.getter(name="importRoutePolicy")
    def import_route_policy(self) -> pulumi.Output[Optional['outputs.ImportRoutePolicyResponse']]:
        """
        Import Route Policy either IPv4 or IPv6.
        """
        return pulumi.get(self, "import_route_policy")

    @property
    @pulumi.getter(name="importRoutePolicyId")
    def import_route_policy_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        ARM Resource ID of the RoutePolicy. This is used for the backward compatibility.
        """
        return pulumi.get(self, "import_route_policy_id")

    @property
    @pulumi.getter(name="ingressAclId")
    def ingress_acl_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Ingress Acl. ARM resource ID of Access Control Lists.
        """
        return pulumi.get(self, "ingress_acl_id")

    @property
    @pulumi.getter(name="isMonitoringEnabled")
    def is_monitoring_enabled(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        To check whether monitoring of internal network is enabled or not.
        """
        return pulumi.get(self, "is_monitoring_enabled")

    @property
    @pulumi.getter
    def mtu(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        Maximum transmission unit. Default value is 1500.
        """
        return pulumi.get(self, "mtu")

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
    @pulumi.getter(name="staticRouteConfiguration")
    def static_route_configuration(self) -> pulumi.Output[Optional['outputs.InternalNetworkPropertiesResponseStaticRouteConfiguration']]:
        """
        Static Route Configuration properties.
        """
        return pulumi.get(self, "static_route_configuration")

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

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> pulumi.Output[builtins.int]:
        """
        Vlan identifier. Example: 1001.
        """
        return pulumi.get(self, "vlan_id")

