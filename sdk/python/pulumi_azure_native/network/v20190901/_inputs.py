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

__all__ = [
    'ConnectionMonitorDestinationArgs',
    'ConnectionMonitorDestinationArgsDict',
    'ConnectionMonitorSourceArgs',
    'ConnectionMonitorSourceArgsDict',
]

MYPY = False

if not MYPY:
    class ConnectionMonitorDestinationArgsDict(TypedDict):
        """
        Describes the destination of connection monitor.
        """
        address: NotRequired[pulumi.Input[str]]
        """
        Address of the connection monitor destination (IP or domain name).
        """
        port: NotRequired[pulumi.Input[int]]
        """
        The destination port used by connection monitor.
        """
        resource_id: NotRequired[pulumi.Input[str]]
        """
        The ID of the resource used as the destination by connection monitor.
        """
elif False:
    ConnectionMonitorDestinationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ConnectionMonitorDestinationArgs:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None):
        """
        Describes the destination of connection monitor.
        :param pulumi.Input[str] address: Address of the connection monitor destination (IP or domain name).
        :param pulumi.Input[int] port: The destination port used by connection monitor.
        :param pulumi.Input[str] resource_id: The ID of the resource used as the destination by connection monitor.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        Address of the connection monitor destination (IP or domain name).
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The destination port used by connection monitor.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the resource used as the destination by connection monitor.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)


if not MYPY:
    class ConnectionMonitorSourceArgsDict(TypedDict):
        """
        Describes the source of connection monitor.
        """
        resource_id: pulumi.Input[str]
        """
        The ID of the resource used as the source by connection monitor.
        """
        port: NotRequired[pulumi.Input[int]]
        """
        The source port used by connection monitor.
        """
elif False:
    ConnectionMonitorSourceArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ConnectionMonitorSourceArgs:
    def __init__(__self__, *,
                 resource_id: pulumi.Input[str],
                 port: Optional[pulumi.Input[int]] = None):
        """
        Describes the source of connection monitor.
        :param pulumi.Input[str] resource_id: The ID of the resource used as the source by connection monitor.
        :param pulumi.Input[int] port: The source port used by connection monitor.
        """
        pulumi.set(__self__, "resource_id", resource_id)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        The ID of the resource used as the source by connection monitor.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The source port used by connection monitor.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)


