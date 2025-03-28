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

__all__ = [
    'ApplicationTypeVersionsCleanupPolicyResponse',
    'AzureActiveDirectoryResponse',
    'ClientCertificateResponse',
    'IPTagResponse',
    'LoadBalancingRuleResponse',
    'NetworkSecurityRuleResponse',
    'ServiceEndpointResponse',
    'SettingsParameterDescriptionResponse',
    'SettingsSectionDescriptionResponse',
    'SkuResponse',
    'SubnetResponse',
    'SystemDataResponse',
]

@pulumi.output_type
class ApplicationTypeVersionsCleanupPolicyResponse(dict):
    """
    The policy used to clean up unused versions. When the policy is not specified explicitly, the default unused application versions to keep will be 3.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maxUnusedVersionsToKeep":
            suggest = "max_unused_versions_to_keep"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApplicationTypeVersionsCleanupPolicyResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApplicationTypeVersionsCleanupPolicyResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApplicationTypeVersionsCleanupPolicyResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 max_unused_versions_to_keep: int):
        """
        The policy used to clean up unused versions. When the policy is not specified explicitly, the default unused application versions to keep will be 3.
        :param int max_unused_versions_to_keep: Number of unused versions per application type to keep.
        """
        pulumi.set(__self__, "max_unused_versions_to_keep", max_unused_versions_to_keep)

    @property
    @pulumi.getter(name="maxUnusedVersionsToKeep")
    def max_unused_versions_to_keep(self) -> int:
        """
        Number of unused versions per application type to keep.
        """
        return pulumi.get(self, "max_unused_versions_to_keep")


@pulumi.output_type
class AzureActiveDirectoryResponse(dict):
    """
    The settings to enable AAD authentication on the cluster.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clientApplication":
            suggest = "client_application"
        elif key == "clusterApplication":
            suggest = "cluster_application"
        elif key == "tenantId":
            suggest = "tenant_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AzureActiveDirectoryResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AzureActiveDirectoryResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AzureActiveDirectoryResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 client_application: Optional[str] = None,
                 cluster_application: Optional[str] = None,
                 tenant_id: Optional[str] = None):
        """
        The settings to enable AAD authentication on the cluster.
        :param str client_application: Azure active directory client application id.
        :param str cluster_application: Azure active directory cluster application id.
        :param str tenant_id: Azure active directory tenant id.
        """
        if client_application is not None:
            pulumi.set(__self__, "client_application", client_application)
        if cluster_application is not None:
            pulumi.set(__self__, "cluster_application", cluster_application)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="clientApplication")
    def client_application(self) -> Optional[str]:
        """
        Azure active directory client application id.
        """
        return pulumi.get(self, "client_application")

    @property
    @pulumi.getter(name="clusterApplication")
    def cluster_application(self) -> Optional[str]:
        """
        Azure active directory cluster application id.
        """
        return pulumi.get(self, "cluster_application")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        """
        Azure active directory tenant id.
        """
        return pulumi.get(self, "tenant_id")


@pulumi.output_type
class ClientCertificateResponse(dict):
    """
    Client certificate definition.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "isAdmin":
            suggest = "is_admin"
        elif key == "commonName":
            suggest = "common_name"
        elif key == "issuerThumbprint":
            suggest = "issuer_thumbprint"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClientCertificateResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClientCertificateResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClientCertificateResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 is_admin: bool,
                 common_name: Optional[str] = None,
                 issuer_thumbprint: Optional[str] = None,
                 thumbprint: Optional[str] = None):
        """
        Client certificate definition.
        :param bool is_admin: Indicates if the client certificate has admin access to the cluster. Non admin clients can perform only read only operations on the cluster.
        :param str common_name: Certificate common name.
        :param str issuer_thumbprint: Issuer thumbprint for the certificate. Only used together with CommonName.
        :param str thumbprint: Certificate thumbprint.
        """
        pulumi.set(__self__, "is_admin", is_admin)
        if common_name is not None:
            pulumi.set(__self__, "common_name", common_name)
        if issuer_thumbprint is not None:
            pulumi.set(__self__, "issuer_thumbprint", issuer_thumbprint)
        if thumbprint is not None:
            pulumi.set(__self__, "thumbprint", thumbprint)

    @property
    @pulumi.getter(name="isAdmin")
    def is_admin(self) -> bool:
        """
        Indicates if the client certificate has admin access to the cluster. Non admin clients can perform only read only operations on the cluster.
        """
        return pulumi.get(self, "is_admin")

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> Optional[str]:
        """
        Certificate common name.
        """
        return pulumi.get(self, "common_name")

    @property
    @pulumi.getter(name="issuerThumbprint")
    def issuer_thumbprint(self) -> Optional[str]:
        """
        Issuer thumbprint for the certificate. Only used together with CommonName.
        """
        return pulumi.get(self, "issuer_thumbprint")

    @property
    @pulumi.getter
    def thumbprint(self) -> Optional[str]:
        """
        Certificate thumbprint.
        """
        return pulumi.get(self, "thumbprint")


@pulumi.output_type
class IPTagResponse(dict):
    """
    IPTag associated with the object.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ipTagType":
            suggest = "ip_tag_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IPTagResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IPTagResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IPTagResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ip_tag_type: str,
                 tag: str):
        """
        IPTag associated with the object.
        :param str ip_tag_type: The IP tag type.
        :param str tag: The value of the IP tag.
        """
        pulumi.set(__self__, "ip_tag_type", ip_tag_type)
        pulumi.set(__self__, "tag", tag)

    @property
    @pulumi.getter(name="ipTagType")
    def ip_tag_type(self) -> str:
        """
        The IP tag type.
        """
        return pulumi.get(self, "ip_tag_type")

    @property
    @pulumi.getter
    def tag(self) -> str:
        """
        The value of the IP tag.
        """
        return pulumi.get(self, "tag")


@pulumi.output_type
class LoadBalancingRuleResponse(dict):
    """
    Describes a load balancing rule.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "backendPort":
            suggest = "backend_port"
        elif key == "frontendPort":
            suggest = "frontend_port"
        elif key == "probeProtocol":
            suggest = "probe_protocol"
        elif key == "loadDistribution":
            suggest = "load_distribution"
        elif key == "probePort":
            suggest = "probe_port"
        elif key == "probeRequestPath":
            suggest = "probe_request_path"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in LoadBalancingRuleResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        LoadBalancingRuleResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        LoadBalancingRuleResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 backend_port: int,
                 frontend_port: int,
                 probe_protocol: str,
                 protocol: str,
                 load_distribution: Optional[str] = None,
                 probe_port: Optional[int] = None,
                 probe_request_path: Optional[str] = None):
        """
        Describes a load balancing rule.
        :param int backend_port: The port used for internal connections on the endpoint. Acceptable values are between 1 and 65535.
        :param int frontend_port: The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values are between 1 and 65534.
        :param str probe_protocol: the reference to the load balancer probe used by the load balancing rule.
        :param str protocol: The reference to the transport protocol used by the load balancing rule.
        :param str load_distribution: The load distribution policy for this rule.
        :param int probe_port: The prob port used by the load balancing rule. Acceptable values are between 1 and 65535.
        :param str probe_request_path: The probe request path. Only supported for HTTP/HTTPS probes.
        """
        pulumi.set(__self__, "backend_port", backend_port)
        pulumi.set(__self__, "frontend_port", frontend_port)
        pulumi.set(__self__, "probe_protocol", probe_protocol)
        pulumi.set(__self__, "protocol", protocol)
        if load_distribution is not None:
            pulumi.set(__self__, "load_distribution", load_distribution)
        if probe_port is not None:
            pulumi.set(__self__, "probe_port", probe_port)
        if probe_request_path is not None:
            pulumi.set(__self__, "probe_request_path", probe_request_path)

    @property
    @pulumi.getter(name="backendPort")
    def backend_port(self) -> int:
        """
        The port used for internal connections on the endpoint. Acceptable values are between 1 and 65535.
        """
        return pulumi.get(self, "backend_port")

    @property
    @pulumi.getter(name="frontendPort")
    def frontend_port(self) -> int:
        """
        The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values are between 1 and 65534.
        """
        return pulumi.get(self, "frontend_port")

    @property
    @pulumi.getter(name="probeProtocol")
    def probe_protocol(self) -> str:
        """
        the reference to the load balancer probe used by the load balancing rule.
        """
        return pulumi.get(self, "probe_protocol")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        The reference to the transport protocol used by the load balancing rule.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="loadDistribution")
    def load_distribution(self) -> Optional[str]:
        """
        The load distribution policy for this rule.
        """
        return pulumi.get(self, "load_distribution")

    @property
    @pulumi.getter(name="probePort")
    def probe_port(self) -> Optional[int]:
        """
        The prob port used by the load balancing rule. Acceptable values are between 1 and 65535.
        """
        return pulumi.get(self, "probe_port")

    @property
    @pulumi.getter(name="probeRequestPath")
    def probe_request_path(self) -> Optional[str]:
        """
        The probe request path. Only supported for HTTP/HTTPS probes.
        """
        return pulumi.get(self, "probe_request_path")


@pulumi.output_type
class NetworkSecurityRuleResponse(dict):
    """
    Describes a network security rule.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "destinationAddressPrefix":
            suggest = "destination_address_prefix"
        elif key == "destinationAddressPrefixes":
            suggest = "destination_address_prefixes"
        elif key == "destinationPortRange":
            suggest = "destination_port_range"
        elif key == "destinationPortRanges":
            suggest = "destination_port_ranges"
        elif key == "sourceAddressPrefix":
            suggest = "source_address_prefix"
        elif key == "sourceAddressPrefixes":
            suggest = "source_address_prefixes"
        elif key == "sourcePortRange":
            suggest = "source_port_range"
        elif key == "sourcePortRanges":
            suggest = "source_port_ranges"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NetworkSecurityRuleResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NetworkSecurityRuleResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NetworkSecurityRuleResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 access: str,
                 direction: str,
                 name: str,
                 priority: int,
                 protocol: str,
                 description: Optional[str] = None,
                 destination_address_prefix: Optional[str] = None,
                 destination_address_prefixes: Optional[Sequence[str]] = None,
                 destination_port_range: Optional[str] = None,
                 destination_port_ranges: Optional[Sequence[str]] = None,
                 source_address_prefix: Optional[str] = None,
                 source_address_prefixes: Optional[Sequence[str]] = None,
                 source_port_range: Optional[str] = None,
                 source_port_ranges: Optional[Sequence[str]] = None):
        """
        Describes a network security rule.
        :param str access: The network traffic is allowed or denied.
        :param str direction: Network security rule direction.
        :param str name: Network security rule name.
        :param int priority: The priority of the rule. The value can be in the range 1000 to 3000. Values outside this range are reserved for Service Fabric ManagerCluster Resource Provider. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
        :param str protocol: Network protocol this rule applies to.
        :param str description: Network security rule description.
        :param str destination_address_prefix: The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
        :param Sequence[str] destination_address_prefixes: The destination address prefixes. CIDR or destination IP ranges.
        :param str destination_port_range: he destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
        :param Sequence[str] destination_port_ranges: The destination port ranges.
        :param str source_address_prefix: The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
        :param Sequence[str] source_address_prefixes: The CIDR or source IP ranges.
        :param str source_port_range: The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
        :param Sequence[str] source_port_ranges: The source port ranges.
        """
        pulumi.set(__self__, "access", access)
        pulumi.set(__self__, "direction", direction)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "priority", priority)
        pulumi.set(__self__, "protocol", protocol)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destination_address_prefix is not None:
            pulumi.set(__self__, "destination_address_prefix", destination_address_prefix)
        if destination_address_prefixes is not None:
            pulumi.set(__self__, "destination_address_prefixes", destination_address_prefixes)
        if destination_port_range is not None:
            pulumi.set(__self__, "destination_port_range", destination_port_range)
        if destination_port_ranges is not None:
            pulumi.set(__self__, "destination_port_ranges", destination_port_ranges)
        if source_address_prefix is not None:
            pulumi.set(__self__, "source_address_prefix", source_address_prefix)
        if source_address_prefixes is not None:
            pulumi.set(__self__, "source_address_prefixes", source_address_prefixes)
        if source_port_range is not None:
            pulumi.set(__self__, "source_port_range", source_port_range)
        if source_port_ranges is not None:
            pulumi.set(__self__, "source_port_ranges", source_port_ranges)

    @property
    @pulumi.getter
    def access(self) -> str:
        """
        The network traffic is allowed or denied.
        """
        return pulumi.get(self, "access")

    @property
    @pulumi.getter
    def direction(self) -> str:
        """
        Network security rule direction.
        """
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Network security rule name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> int:
        """
        The priority of the rule. The value can be in the range 1000 to 3000. Values outside this range are reserved for Service Fabric ManagerCluster Resource Provider. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        Network protocol this rule applies to.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Network security rule description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="destinationAddressPrefix")
    def destination_address_prefix(self) -> Optional[str]:
        """
        The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
        """
        return pulumi.get(self, "destination_address_prefix")

    @property
    @pulumi.getter(name="destinationAddressPrefixes")
    def destination_address_prefixes(self) -> Optional[Sequence[str]]:
        """
        The destination address prefixes. CIDR or destination IP ranges.
        """
        return pulumi.get(self, "destination_address_prefixes")

    @property
    @pulumi.getter(name="destinationPortRange")
    def destination_port_range(self) -> Optional[str]:
        """
        he destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
        """
        return pulumi.get(self, "destination_port_range")

    @property
    @pulumi.getter(name="destinationPortRanges")
    def destination_port_ranges(self) -> Optional[Sequence[str]]:
        """
        The destination port ranges.
        """
        return pulumi.get(self, "destination_port_ranges")

    @property
    @pulumi.getter(name="sourceAddressPrefix")
    def source_address_prefix(self) -> Optional[str]:
        """
        The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
        """
        return pulumi.get(self, "source_address_prefix")

    @property
    @pulumi.getter(name="sourceAddressPrefixes")
    def source_address_prefixes(self) -> Optional[Sequence[str]]:
        """
        The CIDR or source IP ranges.
        """
        return pulumi.get(self, "source_address_prefixes")

    @property
    @pulumi.getter(name="sourcePortRange")
    def source_port_range(self) -> Optional[str]:
        """
        The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
        """
        return pulumi.get(self, "source_port_range")

    @property
    @pulumi.getter(name="sourcePortRanges")
    def source_port_ranges(self) -> Optional[Sequence[str]]:
        """
        The source port ranges.
        """
        return pulumi.get(self, "source_port_ranges")


@pulumi.output_type
class ServiceEndpointResponse(dict):
    """
    The service endpoint properties.
    """
    def __init__(__self__, *,
                 service: str,
                 locations: Optional[Sequence[str]] = None):
        """
        The service endpoint properties.
        :param str service: The type of the endpoint service.
        :param Sequence[str] locations: A list of locations.
        """
        pulumi.set(__self__, "service", service)
        if locations is not None:
            pulumi.set(__self__, "locations", locations)

    @property
    @pulumi.getter
    def service(self) -> str:
        """
        The type of the endpoint service.
        """
        return pulumi.get(self, "service")

    @property
    @pulumi.getter
    def locations(self) -> Optional[Sequence[str]]:
        """
        A list of locations.
        """
        return pulumi.get(self, "locations")


@pulumi.output_type
class SettingsParameterDescriptionResponse(dict):
    """
    Describes a parameter in fabric settings of the cluster.
    """
    def __init__(__self__, *,
                 name: str,
                 value: str):
        """
        Describes a parameter in fabric settings of the cluster.
        :param str name: The parameter name of fabric setting.
        :param str value: The parameter value of fabric setting.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The parameter name of fabric setting.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The parameter value of fabric setting.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class SettingsSectionDescriptionResponse(dict):
    """
    Describes a section in the fabric settings of the cluster.
    """
    def __init__(__self__, *,
                 name: str,
                 parameters: Sequence['outputs.SettingsParameterDescriptionResponse']):
        """
        Describes a section in the fabric settings of the cluster.
        :param str name: The section name of the fabric settings.
        :param Sequence['SettingsParameterDescriptionResponse'] parameters: The collection of parameters in the section.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The section name of the fabric settings.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> Sequence['outputs.SettingsParameterDescriptionResponse']:
        """
        The collection of parameters in the section.
        """
        return pulumi.get(self, "parameters")


@pulumi.output_type
class SkuResponse(dict):
    """
    Service Fabric managed cluster Sku definition
    """
    def __init__(__self__, *,
                 name: str):
        """
        Service Fabric managed cluster Sku definition
        :param str name: Sku Name.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Sku Name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class SubnetResponse(dict):
    """
    Describes a Subnet.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "enableIpv6":
            suggest = "enable_ipv6"
        elif key == "networkSecurityGroupId":
            suggest = "network_security_group_id"
        elif key == "privateEndpointNetworkPolicies":
            suggest = "private_endpoint_network_policies"
        elif key == "privateLinkServiceNetworkPolicies":
            suggest = "private_link_service_network_policies"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SubnetResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SubnetResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SubnetResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 enable_ipv6: Optional[bool] = None,
                 network_security_group_id: Optional[str] = None,
                 private_endpoint_network_policies: Optional[str] = None,
                 private_link_service_network_policies: Optional[str] = None):
        """
        Describes a Subnet.
        :param str name: Subnet name.
        :param bool enable_ipv6: Indicates wether to enable Ipv6 or not. If not provided, it will take the same configuration as the cluster.
        :param str network_security_group_id: Full resource id for the network security group.
        :param str private_endpoint_network_policies: Enable or Disable apply network policies on private end point in the subnet.
        :param str private_link_service_network_policies: Enable or Disable apply network policies on private link service in the subnet.
        """
        pulumi.set(__self__, "name", name)
        if enable_ipv6 is not None:
            pulumi.set(__self__, "enable_ipv6", enable_ipv6)
        if network_security_group_id is not None:
            pulumi.set(__self__, "network_security_group_id", network_security_group_id)
        if private_endpoint_network_policies is not None:
            pulumi.set(__self__, "private_endpoint_network_policies", private_endpoint_network_policies)
        if private_link_service_network_policies is not None:
            pulumi.set(__self__, "private_link_service_network_policies", private_link_service_network_policies)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Subnet name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="enableIpv6")
    def enable_ipv6(self) -> Optional[bool]:
        """
        Indicates wether to enable Ipv6 or not. If not provided, it will take the same configuration as the cluster.
        """
        return pulumi.get(self, "enable_ipv6")

    @property
    @pulumi.getter(name="networkSecurityGroupId")
    def network_security_group_id(self) -> Optional[str]:
        """
        Full resource id for the network security group.
        """
        return pulumi.get(self, "network_security_group_id")

    @property
    @pulumi.getter(name="privateEndpointNetworkPolicies")
    def private_endpoint_network_policies(self) -> Optional[str]:
        """
        Enable or Disable apply network policies on private end point in the subnet.
        """
        return pulumi.get(self, "private_endpoint_network_policies")

    @property
    @pulumi.getter(name="privateLinkServiceNetworkPolicies")
    def private_link_service_network_policies(self) -> Optional[str]:
        """
        Enable or Disable apply network policies on private link service in the subnet.
        """
        return pulumi.get(self, "private_link_service_network_policies")


@pulumi.output_type
class SystemDataResponse(dict):
    """
    Metadata pertaining to creation and last modification of the resource.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "createdBy":
            suggest = "created_by"
        elif key == "createdByType":
            suggest = "created_by_type"
        elif key == "lastModifiedAt":
            suggest = "last_modified_at"
        elif key == "lastModifiedBy":
            suggest = "last_modified_by"
        elif key == "lastModifiedByType":
            suggest = "last_modified_by_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SystemDataResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SystemDataResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SystemDataResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: Optional[str] = None,
                 created_by: Optional[str] = None,
                 created_by_type: Optional[str] = None,
                 last_modified_at: Optional[str] = None,
                 last_modified_by: Optional[str] = None,
                 last_modified_by_type: Optional[str] = None):
        """
        Metadata pertaining to creation and last modification of the resource.
        :param str created_at: The timestamp of resource creation (UTC).
        :param str created_by: The identity that created the resource.
        :param str created_by_type: The type of identity that created the resource.
        :param str last_modified_at: The timestamp of resource last modification (UTC).
        :param str last_modified_by: The identity that last modified the resource.
        :param str last_modified_by_type: The type of identity that last modified the resource.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if created_by is not None:
            pulumi.set(__self__, "created_by", created_by)
        if created_by_type is not None:
            pulumi.set(__self__, "created_by_type", created_by_type)
        if last_modified_at is not None:
            pulumi.set(__self__, "last_modified_at", last_modified_at)
        if last_modified_by is not None:
            pulumi.set(__self__, "last_modified_by", last_modified_by)
        if last_modified_by_type is not None:
            pulumi.set(__self__, "last_modified_by_type", last_modified_by_type)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        The timestamp of resource creation (UTC).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> Optional[str]:
        """
        The identity that created the resource.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="createdByType")
    def created_by_type(self) -> Optional[str]:
        """
        The type of identity that created the resource.
        """
        return pulumi.get(self, "created_by_type")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> Optional[str]:
        """
        The timestamp of resource last modification (UTC).
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter(name="lastModifiedBy")
    def last_modified_by(self) -> Optional[str]:
        """
        The identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by")

    @property
    @pulumi.getter(name="lastModifiedByType")
    def last_modified_by_type(self) -> Optional[str]:
        """
        The type of identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by_type")


