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
    'GetGlobalReachConnectionResult',
    'AwaitableGetGlobalReachConnectionResult',
    'get_global_reach_connection',
    'get_global_reach_connection_output',
]

@pulumi.output_type
class GetGlobalReachConnectionResult:
    """
    A global reach connection resource
    """
    def __init__(__self__, address_prefix=None, authorization_key=None, circuit_connection_status=None, express_route_id=None, id=None, name=None, peer_express_route_circuit=None, provisioning_state=None, type=None):
        if address_prefix and not isinstance(address_prefix, str):
            raise TypeError("Expected argument 'address_prefix' to be a str")
        pulumi.set(__self__, "address_prefix", address_prefix)
        if authorization_key and not isinstance(authorization_key, str):
            raise TypeError("Expected argument 'authorization_key' to be a str")
        pulumi.set(__self__, "authorization_key", authorization_key)
        if circuit_connection_status and not isinstance(circuit_connection_status, str):
            raise TypeError("Expected argument 'circuit_connection_status' to be a str")
        pulumi.set(__self__, "circuit_connection_status", circuit_connection_status)
        if express_route_id and not isinstance(express_route_id, str):
            raise TypeError("Expected argument 'express_route_id' to be a str")
        pulumi.set(__self__, "express_route_id", express_route_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if peer_express_route_circuit and not isinstance(peer_express_route_circuit, str):
            raise TypeError("Expected argument 'peer_express_route_circuit' to be a str")
        pulumi.set(__self__, "peer_express_route_circuit", peer_express_route_circuit)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="addressPrefix")
    def address_prefix(self) -> str:
        """
        The network used for global reach carved out from the original network block provided for the private cloud
        """
        return pulumi.get(self, "address_prefix")

    @property
    @pulumi.getter(name="authorizationKey")
    def authorization_key(self) -> Optional[str]:
        """
        Authorization key from the peer express route used for the global reach connection
        """
        return pulumi.get(self, "authorization_key")

    @property
    @pulumi.getter(name="circuitConnectionStatus")
    def circuit_connection_status(self) -> str:
        """
        The connection status of the global reach connection
        """
        return pulumi.get(self, "circuit_connection_status")

    @property
    @pulumi.getter(name="expressRouteId")
    def express_route_id(self) -> Optional[str]:
        """
        The ID of the Private Cloud's ExpressRoute Circuit that is participating in the global reach connection
        """
        return pulumi.get(self, "express_route_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peerExpressRouteCircuit")
    def peer_express_route_circuit(self) -> Optional[str]:
        """
        Identifier of the ExpressRoute Circuit to peer with in the global reach connection
        """
        return pulumi.get(self, "peer_express_route_circuit")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The state of the  ExpressRoute Circuit Authorization provisioning
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetGlobalReachConnectionResult(GetGlobalReachConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGlobalReachConnectionResult(
            address_prefix=self.address_prefix,
            authorization_key=self.authorization_key,
            circuit_connection_status=self.circuit_connection_status,
            express_route_id=self.express_route_id,
            id=self.id,
            name=self.name,
            peer_express_route_circuit=self.peer_express_route_circuit,
            provisioning_state=self.provisioning_state,
            type=self.type)


def get_global_reach_connection(global_reach_connection_name: Optional[str] = None,
                                private_cloud_name: Optional[str] = None,
                                resource_group_name: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGlobalReachConnectionResult:
    """
    A global reach connection resource

    Uses Azure REST API version 2022-05-01.

    Other available API versions: 2023-03-01, 2023-09-01.


    :param str global_reach_connection_name: Name of the global reach connection in the private cloud
    :param str private_cloud_name: Name of the private cloud
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['globalReachConnectionName'] = global_reach_connection_name
    __args__['privateCloudName'] = private_cloud_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:avs:getGlobalReachConnection', __args__, opts=opts, typ=GetGlobalReachConnectionResult).value

    return AwaitableGetGlobalReachConnectionResult(
        address_prefix=pulumi.get(__ret__, 'address_prefix'),
        authorization_key=pulumi.get(__ret__, 'authorization_key'),
        circuit_connection_status=pulumi.get(__ret__, 'circuit_connection_status'),
        express_route_id=pulumi.get(__ret__, 'express_route_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        peer_express_route_circuit=pulumi.get(__ret__, 'peer_express_route_circuit'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        type=pulumi.get(__ret__, 'type'))
def get_global_reach_connection_output(global_reach_connection_name: Optional[pulumi.Input[str]] = None,
                                       private_cloud_name: Optional[pulumi.Input[str]] = None,
                                       resource_group_name: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetGlobalReachConnectionResult]:
    """
    A global reach connection resource

    Uses Azure REST API version 2022-05-01.

    Other available API versions: 2023-03-01, 2023-09-01.


    :param str global_reach_connection_name: Name of the global reach connection in the private cloud
    :param str private_cloud_name: Name of the private cloud
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['globalReachConnectionName'] = global_reach_connection_name
    __args__['privateCloudName'] = private_cloud_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:avs:getGlobalReachConnection', __args__, opts=opts, typ=GetGlobalReachConnectionResult)
    return __ret__.apply(lambda __response__: GetGlobalReachConnectionResult(
        address_prefix=pulumi.get(__response__, 'address_prefix'),
        authorization_key=pulumi.get(__response__, 'authorization_key'),
        circuit_connection_status=pulumi.get(__response__, 'circuit_connection_status'),
        express_route_id=pulumi.get(__response__, 'express_route_id'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        peer_express_route_circuit=pulumi.get(__response__, 'peer_express_route_circuit'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        type=pulumi.get(__response__, 'type')))
