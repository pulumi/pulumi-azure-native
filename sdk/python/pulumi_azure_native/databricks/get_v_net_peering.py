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

__all__ = [
    'GetVNetPeeringResult',
    'AwaitableGetVNetPeeringResult',
    'get_v_net_peering',
    'get_v_net_peering_output',
]

@pulumi.output_type
class GetVNetPeeringResult:
    """
    Peerings in a VirtualNetwork resource
    """
    def __init__(__self__, allow_forwarded_traffic=None, allow_gateway_transit=None, allow_virtual_network_access=None, azure_api_version=None, databricks_address_space=None, databricks_virtual_network=None, id=None, name=None, peering_state=None, provisioning_state=None, remote_address_space=None, remote_virtual_network=None, type=None, use_remote_gateways=None):
        if allow_forwarded_traffic and not isinstance(allow_forwarded_traffic, bool):
            raise TypeError("Expected argument 'allow_forwarded_traffic' to be a bool")
        pulumi.set(__self__, "allow_forwarded_traffic", allow_forwarded_traffic)
        if allow_gateway_transit and not isinstance(allow_gateway_transit, bool):
            raise TypeError("Expected argument 'allow_gateway_transit' to be a bool")
        pulumi.set(__self__, "allow_gateway_transit", allow_gateway_transit)
        if allow_virtual_network_access and not isinstance(allow_virtual_network_access, bool):
            raise TypeError("Expected argument 'allow_virtual_network_access' to be a bool")
        pulumi.set(__self__, "allow_virtual_network_access", allow_virtual_network_access)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if databricks_address_space and not isinstance(databricks_address_space, dict):
            raise TypeError("Expected argument 'databricks_address_space' to be a dict")
        pulumi.set(__self__, "databricks_address_space", databricks_address_space)
        if databricks_virtual_network and not isinstance(databricks_virtual_network, dict):
            raise TypeError("Expected argument 'databricks_virtual_network' to be a dict")
        pulumi.set(__self__, "databricks_virtual_network", databricks_virtual_network)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if peering_state and not isinstance(peering_state, str):
            raise TypeError("Expected argument 'peering_state' to be a str")
        pulumi.set(__self__, "peering_state", peering_state)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if remote_address_space and not isinstance(remote_address_space, dict):
            raise TypeError("Expected argument 'remote_address_space' to be a dict")
        pulumi.set(__self__, "remote_address_space", remote_address_space)
        if remote_virtual_network and not isinstance(remote_virtual_network, dict):
            raise TypeError("Expected argument 'remote_virtual_network' to be a dict")
        pulumi.set(__self__, "remote_virtual_network", remote_virtual_network)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if use_remote_gateways and not isinstance(use_remote_gateways, bool):
            raise TypeError("Expected argument 'use_remote_gateways' to be a bool")
        pulumi.set(__self__, "use_remote_gateways", use_remote_gateways)

    @property
    @pulumi.getter(name="allowForwardedTraffic")
    def allow_forwarded_traffic(self) -> Optional[builtins.bool]:
        """
        Whether the forwarded traffic from the VMs in the local virtual network will be allowed/disallowed in remote virtual network.
        """
        return pulumi.get(self, "allow_forwarded_traffic")

    @property
    @pulumi.getter(name="allowGatewayTransit")
    def allow_gateway_transit(self) -> Optional[builtins.bool]:
        """
        If gateway links can be used in remote virtual networking to link to this virtual network.
        """
        return pulumi.get(self, "allow_gateway_transit")

    @property
    @pulumi.getter(name="allowVirtualNetworkAccess")
    def allow_virtual_network_access(self) -> Optional[builtins.bool]:
        """
        Whether the VMs in the local virtual network space would be able to access the VMs in remote virtual network space.
        """
        return pulumi.get(self, "allow_virtual_network_access")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="databricksAddressSpace")
    def databricks_address_space(self) -> Optional['outputs.AddressSpaceResponse']:
        """
        The reference to the databricks virtual network address space.
        """
        return pulumi.get(self, "databricks_address_space")

    @property
    @pulumi.getter(name="databricksVirtualNetwork")
    def databricks_virtual_network(self) -> Optional['outputs.VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork']:
        """
         The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
        """
        return pulumi.get(self, "databricks_virtual_network")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Name of the virtual network peering resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peeringState")
    def peering_state(self) -> builtins.str:
        """
        The status of the virtual network peering.
        """
        return pulumi.get(self, "peering_state")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        The provisioning state of the virtual network peering resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="remoteAddressSpace")
    def remote_address_space(self) -> Optional['outputs.AddressSpaceResponse']:
        """
        The reference to the remote virtual network address space.
        """
        return pulumi.get(self, "remote_address_space")

    @property
    @pulumi.getter(name="remoteVirtualNetwork")
    def remote_virtual_network(self) -> 'outputs.VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork':
        """
         The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
        """
        return pulumi.get(self, "remote_virtual_network")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        type of the virtual network peering resource
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="useRemoteGateways")
    def use_remote_gateways(self) -> Optional[builtins.bool]:
        """
        If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
        """
        return pulumi.get(self, "use_remote_gateways")


class AwaitableGetVNetPeeringResult(GetVNetPeeringResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVNetPeeringResult(
            allow_forwarded_traffic=self.allow_forwarded_traffic,
            allow_gateway_transit=self.allow_gateway_transit,
            allow_virtual_network_access=self.allow_virtual_network_access,
            azure_api_version=self.azure_api_version,
            databricks_address_space=self.databricks_address_space,
            databricks_virtual_network=self.databricks_virtual_network,
            id=self.id,
            name=self.name,
            peering_state=self.peering_state,
            provisioning_state=self.provisioning_state,
            remote_address_space=self.remote_address_space,
            remote_virtual_network=self.remote_virtual_network,
            type=self.type,
            use_remote_gateways=self.use_remote_gateways)


def get_v_net_peering(peering_name: Optional[builtins.str] = None,
                      resource_group_name: Optional[builtins.str] = None,
                      workspace_name: Optional[builtins.str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVNetPeeringResult:
    """
    Gets the workspace vNet Peering.

    Uses Azure REST API version 2024-05-01.

    Other available API versions: 2023-02-01, 2023-09-15-preview, 2024-09-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databricks [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str peering_name: The name of the workspace vNet peering.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['peeringName'] = peering_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:databricks:getVNetPeering', __args__, opts=opts, typ=GetVNetPeeringResult).value

    return AwaitableGetVNetPeeringResult(
        allow_forwarded_traffic=pulumi.get(__ret__, 'allow_forwarded_traffic'),
        allow_gateway_transit=pulumi.get(__ret__, 'allow_gateway_transit'),
        allow_virtual_network_access=pulumi.get(__ret__, 'allow_virtual_network_access'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        databricks_address_space=pulumi.get(__ret__, 'databricks_address_space'),
        databricks_virtual_network=pulumi.get(__ret__, 'databricks_virtual_network'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        peering_state=pulumi.get(__ret__, 'peering_state'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        remote_address_space=pulumi.get(__ret__, 'remote_address_space'),
        remote_virtual_network=pulumi.get(__ret__, 'remote_virtual_network'),
        type=pulumi.get(__ret__, 'type'),
        use_remote_gateways=pulumi.get(__ret__, 'use_remote_gateways'))
def get_v_net_peering_output(peering_name: Optional[pulumi.Input[builtins.str]] = None,
                             resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                             workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVNetPeeringResult]:
    """
    Gets the workspace vNet Peering.

    Uses Azure REST API version 2024-05-01.

    Other available API versions: 2023-02-01, 2023-09-15-preview, 2024-09-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databricks [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str peering_name: The name of the workspace vNet peering.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['peeringName'] = peering_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:databricks:getVNetPeering', __args__, opts=opts, typ=GetVNetPeeringResult)
    return __ret__.apply(lambda __response__: GetVNetPeeringResult(
        allow_forwarded_traffic=pulumi.get(__response__, 'allow_forwarded_traffic'),
        allow_gateway_transit=pulumi.get(__response__, 'allow_gateway_transit'),
        allow_virtual_network_access=pulumi.get(__response__, 'allow_virtual_network_access'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        databricks_address_space=pulumi.get(__response__, 'databricks_address_space'),
        databricks_virtual_network=pulumi.get(__response__, 'databricks_virtual_network'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        peering_state=pulumi.get(__response__, 'peering_state'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        remote_address_space=pulumi.get(__response__, 'remote_address_space'),
        remote_virtual_network=pulumi.get(__response__, 'remote_virtual_network'),
        type=pulumi.get(__response__, 'type'),
        use_remote_gateways=pulumi.get(__response__, 'use_remote_gateways')))
