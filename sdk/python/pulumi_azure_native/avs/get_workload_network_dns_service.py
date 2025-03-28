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

__all__ = [
    'GetWorkloadNetworkDnsServiceResult',
    'AwaitableGetWorkloadNetworkDnsServiceResult',
    'get_workload_network_dns_service',
    'get_workload_network_dns_service_output',
]

@pulumi.output_type
class GetWorkloadNetworkDnsServiceResult:
    """
    NSX DNS Service
    """
    def __init__(__self__, default_dns_zone=None, display_name=None, dns_service_ip=None, fqdn_zones=None, id=None, log_level=None, name=None, provisioning_state=None, revision=None, status=None, type=None):
        if default_dns_zone and not isinstance(default_dns_zone, str):
            raise TypeError("Expected argument 'default_dns_zone' to be a str")
        pulumi.set(__self__, "default_dns_zone", default_dns_zone)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if dns_service_ip and not isinstance(dns_service_ip, str):
            raise TypeError("Expected argument 'dns_service_ip' to be a str")
        pulumi.set(__self__, "dns_service_ip", dns_service_ip)
        if fqdn_zones and not isinstance(fqdn_zones, list):
            raise TypeError("Expected argument 'fqdn_zones' to be a list")
        pulumi.set(__self__, "fqdn_zones", fqdn_zones)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if log_level and not isinstance(log_level, str):
            raise TypeError("Expected argument 'log_level' to be a str")
        pulumi.set(__self__, "log_level", log_level)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if revision and not isinstance(revision, float):
            raise TypeError("Expected argument 'revision' to be a float")
        pulumi.set(__self__, "revision", revision)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="defaultDnsZone")
    def default_dns_zone(self) -> Optional[str]:
        """
        Default DNS zone of the DNS Service.
        """
        return pulumi.get(self, "default_dns_zone")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        Display name of the DNS Service.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="dnsServiceIp")
    def dns_service_ip(self) -> Optional[str]:
        """
        DNS service IP of the DNS Service.
        """
        return pulumi.get(self, "dns_service_ip")

    @property
    @pulumi.getter(name="fqdnZones")
    def fqdn_zones(self) -> Optional[Sequence[str]]:
        """
        FQDN zones of the DNS Service.
        """
        return pulumi.get(self, "fqdn_zones")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="logLevel")
    def log_level(self) -> Optional[str]:
        """
        DNS Service log level.
        """
        return pulumi.get(self, "log_level")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def revision(self) -> Optional[float]:
        """
        NSX revision number.
        """
        return pulumi.get(self, "revision")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        DNS Service status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetWorkloadNetworkDnsServiceResult(GetWorkloadNetworkDnsServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkloadNetworkDnsServiceResult(
            default_dns_zone=self.default_dns_zone,
            display_name=self.display_name,
            dns_service_ip=self.dns_service_ip,
            fqdn_zones=self.fqdn_zones,
            id=self.id,
            log_level=self.log_level,
            name=self.name,
            provisioning_state=self.provisioning_state,
            revision=self.revision,
            status=self.status,
            type=self.type)


def get_workload_network_dns_service(dns_service_id: Optional[str] = None,
                                     private_cloud_name: Optional[str] = None,
                                     resource_group_name: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkloadNetworkDnsServiceResult:
    """
    NSX DNS Service

    Uses Azure REST API version 2022-05-01.

    Other available API versions: 2023-03-01, 2023-09-01.


    :param str dns_service_id: NSX DNS Service identifier. Generally the same as the DNS Service's display name
    :param str private_cloud_name: Name of the private cloud
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['dnsServiceId'] = dns_service_id
    __args__['privateCloudName'] = private_cloud_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:avs:getWorkloadNetworkDnsService', __args__, opts=opts, typ=GetWorkloadNetworkDnsServiceResult).value

    return AwaitableGetWorkloadNetworkDnsServiceResult(
        default_dns_zone=pulumi.get(__ret__, 'default_dns_zone'),
        display_name=pulumi.get(__ret__, 'display_name'),
        dns_service_ip=pulumi.get(__ret__, 'dns_service_ip'),
        fqdn_zones=pulumi.get(__ret__, 'fqdn_zones'),
        id=pulumi.get(__ret__, 'id'),
        log_level=pulumi.get(__ret__, 'log_level'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        revision=pulumi.get(__ret__, 'revision'),
        status=pulumi.get(__ret__, 'status'),
        type=pulumi.get(__ret__, 'type'))
def get_workload_network_dns_service_output(dns_service_id: Optional[pulumi.Input[str]] = None,
                                            private_cloud_name: Optional[pulumi.Input[str]] = None,
                                            resource_group_name: Optional[pulumi.Input[str]] = None,
                                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetWorkloadNetworkDnsServiceResult]:
    """
    NSX DNS Service

    Uses Azure REST API version 2022-05-01.

    Other available API versions: 2023-03-01, 2023-09-01.


    :param str dns_service_id: NSX DNS Service identifier. Generally the same as the DNS Service's display name
    :param str private_cloud_name: Name of the private cloud
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['dnsServiceId'] = dns_service_id
    __args__['privateCloudName'] = private_cloud_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:avs:getWorkloadNetworkDnsService', __args__, opts=opts, typ=GetWorkloadNetworkDnsServiceResult)
    return __ret__.apply(lambda __response__: GetWorkloadNetworkDnsServiceResult(
        default_dns_zone=pulumi.get(__response__, 'default_dns_zone'),
        display_name=pulumi.get(__response__, 'display_name'),
        dns_service_ip=pulumi.get(__response__, 'dns_service_ip'),
        fqdn_zones=pulumi.get(__response__, 'fqdn_zones'),
        id=pulumi.get(__response__, 'id'),
        log_level=pulumi.get(__response__, 'log_level'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        revision=pulumi.get(__response__, 'revision'),
        status=pulumi.get(__response__, 'status'),
        type=pulumi.get(__response__, 'type')))
