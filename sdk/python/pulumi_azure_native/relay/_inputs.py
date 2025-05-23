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
from ._enums import *

__all__ = [
    'ConnectionStateArgs',
    'ConnectionStateArgsDict',
    'PrivateEndpointConnectionArgs',
    'PrivateEndpointConnectionArgsDict',
    'PrivateEndpointArgs',
    'PrivateEndpointArgsDict',
    'SkuArgs',
    'SkuArgsDict',
]

MYPY = False

if not MYPY:
    class ConnectionStateArgsDict(TypedDict):
        """
        ConnectionState information.
        """
        description: NotRequired[pulumi.Input[builtins.str]]
        """
        Description of the connection state.
        """
        status: NotRequired[pulumi.Input[Union[builtins.str, 'PrivateLinkConnectionStatus']]]
        """
        Status of the connection.
        """
elif False:
    ConnectionStateArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ConnectionStateArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[Union[builtins.str, 'PrivateLinkConnectionStatus']]] = None):
        """
        ConnectionState information.
        :param pulumi.Input[builtins.str] description: Description of the connection state.
        :param pulumi.Input[Union[builtins.str, 'PrivateLinkConnectionStatus']] status: Status of the connection.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description of the connection state.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[Union[builtins.str, 'PrivateLinkConnectionStatus']]]:
        """
        Status of the connection.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[Union[builtins.str, 'PrivateLinkConnectionStatus']]]):
        pulumi.set(self, "status", value)


if not MYPY:
    class PrivateEndpointConnectionArgsDict(TypedDict):
        """
        Properties of the PrivateEndpointConnection.
        """
        private_endpoint: NotRequired[pulumi.Input['PrivateEndpointArgsDict']]
        """
        The Private Endpoint resource for this Connection.
        """
        private_link_service_connection_state: NotRequired[pulumi.Input['ConnectionStateArgsDict']]
        """
        Details about the state of the connection.
        """
        provisioning_state: NotRequired[pulumi.Input[Union[builtins.str, 'EndPointProvisioningState']]]
        """
        Provisioning state of the Private Endpoint Connection.
        """
elif False:
    PrivateEndpointConnectionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class PrivateEndpointConnectionArgs:
    def __init__(__self__, *,
                 private_endpoint: Optional[pulumi.Input['PrivateEndpointArgs']] = None,
                 private_link_service_connection_state: Optional[pulumi.Input['ConnectionStateArgs']] = None,
                 provisioning_state: Optional[pulumi.Input[Union[builtins.str, 'EndPointProvisioningState']]] = None):
        """
        Properties of the PrivateEndpointConnection.
        :param pulumi.Input['PrivateEndpointArgs'] private_endpoint: The Private Endpoint resource for this Connection.
        :param pulumi.Input['ConnectionStateArgs'] private_link_service_connection_state: Details about the state of the connection.
        :param pulumi.Input[Union[builtins.str, 'EndPointProvisioningState']] provisioning_state: Provisioning state of the Private Endpoint Connection.
        """
        if private_endpoint is not None:
            pulumi.set(__self__, "private_endpoint", private_endpoint)
        if private_link_service_connection_state is not None:
            pulumi.set(__self__, "private_link_service_connection_state", private_link_service_connection_state)
        if provisioning_state is not None:
            pulumi.set(__self__, "provisioning_state", provisioning_state)

    @property
    @pulumi.getter(name="privateEndpoint")
    def private_endpoint(self) -> Optional[pulumi.Input['PrivateEndpointArgs']]:
        """
        The Private Endpoint resource for this Connection.
        """
        return pulumi.get(self, "private_endpoint")

    @private_endpoint.setter
    def private_endpoint(self, value: Optional[pulumi.Input['PrivateEndpointArgs']]):
        pulumi.set(self, "private_endpoint", value)

    @property
    @pulumi.getter(name="privateLinkServiceConnectionState")
    def private_link_service_connection_state(self) -> Optional[pulumi.Input['ConnectionStateArgs']]:
        """
        Details about the state of the connection.
        """
        return pulumi.get(self, "private_link_service_connection_state")

    @private_link_service_connection_state.setter
    def private_link_service_connection_state(self, value: Optional[pulumi.Input['ConnectionStateArgs']]):
        pulumi.set(self, "private_link_service_connection_state", value)

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[pulumi.Input[Union[builtins.str, 'EndPointProvisioningState']]]:
        """
        Provisioning state of the Private Endpoint Connection.
        """
        return pulumi.get(self, "provisioning_state")

    @provisioning_state.setter
    def provisioning_state(self, value: Optional[pulumi.Input[Union[builtins.str, 'EndPointProvisioningState']]]):
        pulumi.set(self, "provisioning_state", value)


if not MYPY:
    class PrivateEndpointArgsDict(TypedDict):
        """
        PrivateEndpoint information.
        """
        id: NotRequired[pulumi.Input[builtins.str]]
        """
        The ARM identifier for Private Endpoint.
        """
elif False:
    PrivateEndpointArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class PrivateEndpointArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[builtins.str]] = None):
        """
        PrivateEndpoint information.
        :param pulumi.Input[builtins.str] id: The ARM identifier for Private Endpoint.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ARM identifier for Private Endpoint.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "id", value)


if not MYPY:
    class SkuArgsDict(TypedDict):
        """
        SKU of the namespace.
        """
        name: pulumi.Input[Union[builtins.str, 'SkuName']]
        """
        Name of this SKU.
        """
        tier: NotRequired[pulumi.Input[Union[builtins.str, 'SkuTier']]]
        """
        The tier of this SKU.
        """
elif False:
    SkuArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SkuArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[Union[builtins.str, 'SkuName']],
                 tier: Optional[pulumi.Input[Union[builtins.str, 'SkuTier']]] = None):
        """
        SKU of the namespace.
        :param pulumi.Input[Union[builtins.str, 'SkuName']] name: Name of this SKU.
        :param pulumi.Input[Union[builtins.str, 'SkuTier']] tier: The tier of this SKU.
        """
        pulumi.set(__self__, "name", name)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[Union[builtins.str, 'SkuName']]:
        """
        Name of this SKU.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[Union[builtins.str, 'SkuName']]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tier(self) -> Optional[pulumi.Input[Union[builtins.str, 'SkuTier']]]:
        """
        The tier of this SKU.
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: Optional[pulumi.Input[Union[builtins.str, 'SkuTier']]]):
        pulumi.set(self, "tier", value)


