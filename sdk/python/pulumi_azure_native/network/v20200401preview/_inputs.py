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
from ._enums import *

__all__ = [
    'InboundEndpointIPConfigurationArgs',
    'InboundEndpointIPConfigurationArgsDict',
    'SubResourceArgs',
    'SubResourceArgsDict',
]

MYPY = False

if not MYPY:
    class InboundEndpointIPConfigurationArgsDict(TypedDict):
        """
        IP configuration.
        """
        private_ip_address: NotRequired[pulumi.Input[str]]
        """
        Private IP address of the IP configuration.
        """
        private_ip_allocation_method: NotRequired[pulumi.Input[Union[str, 'IpAllocationMethod']]]
        """
        Private IP address allocation method.
        """
        subnet: NotRequired[pulumi.Input['SubResourceArgsDict']]
        """
        The reference to the subnet bound to the IP configuration.
        """
elif False:
    InboundEndpointIPConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class InboundEndpointIPConfigurationArgs:
    def __init__(__self__, *,
                 private_ip_address: Optional[pulumi.Input[str]] = None,
                 private_ip_allocation_method: Optional[pulumi.Input[Union[str, 'IpAllocationMethod']]] = None,
                 subnet: Optional[pulumi.Input['SubResourceArgs']] = None):
        """
        IP configuration.
        :param pulumi.Input[str] private_ip_address: Private IP address of the IP configuration.
        :param pulumi.Input[Union[str, 'IpAllocationMethod']] private_ip_allocation_method: Private IP address allocation method.
        :param pulumi.Input['SubResourceArgs'] subnet: The reference to the subnet bound to the IP configuration.
        """
        if private_ip_address is not None:
            pulumi.set(__self__, "private_ip_address", private_ip_address)
        if private_ip_allocation_method is None:
            private_ip_allocation_method = 'Dynamic'
        if private_ip_allocation_method is not None:
            pulumi.set(__self__, "private_ip_allocation_method", private_ip_allocation_method)
        if subnet is not None:
            pulumi.set(__self__, "subnet", subnet)

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        Private IP address of the IP configuration.
        """
        return pulumi.get(self, "private_ip_address")

    @private_ip_address.setter
    def private_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_ip_address", value)

    @property
    @pulumi.getter(name="privateIpAllocationMethod")
    def private_ip_allocation_method(self) -> Optional[pulumi.Input[Union[str, 'IpAllocationMethod']]]:
        """
        Private IP address allocation method.
        """
        return pulumi.get(self, "private_ip_allocation_method")

    @private_ip_allocation_method.setter
    def private_ip_allocation_method(self, value: Optional[pulumi.Input[Union[str, 'IpAllocationMethod']]]):
        pulumi.set(self, "private_ip_allocation_method", value)

    @property
    @pulumi.getter
    def subnet(self) -> Optional[pulumi.Input['SubResourceArgs']]:
        """
        The reference to the subnet bound to the IP configuration.
        """
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: Optional[pulumi.Input['SubResourceArgs']]):
        pulumi.set(self, "subnet", value)


if not MYPY:
    class SubResourceArgsDict(TypedDict):
        """
        Reference to another ARM resource.
        """
        id: NotRequired[pulumi.Input[str]]
        """
        Sub-resource ID. Both absolute resource ID and a relative resource ID are accepted.
        An absolute ID starts with /subscriptions/ and contains the entire ID of the parent resource and the ID of the sub-resource in the end.
        A relative ID replaces the ID of the parent resource with a token '$self', followed by the sub-resource ID itself.
        Example of a relative ID: $self/frontEndConfigurations/my-frontend.
        """
elif False:
    SubResourceArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SubResourceArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None):
        """
        Reference to another ARM resource.
        :param pulumi.Input[str] id: Sub-resource ID. Both absolute resource ID and a relative resource ID are accepted.
               An absolute ID starts with /subscriptions/ and contains the entire ID of the parent resource and the ID of the sub-resource in the end.
               A relative ID replaces the ID of the parent resource with a token '$self', followed by the sub-resource ID itself.
               Example of a relative ID: $self/frontEndConfigurations/my-frontend.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        Sub-resource ID. Both absolute resource ID and a relative resource ID are accepted.
        An absolute ID starts with /subscriptions/ and contains the entire ID of the parent resource and the ID of the sub-resource in the end.
        A relative ID replaces the ID of the parent resource with a token '$self', followed by the sub-resource ID itself.
        Example of a relative ID: $self/frontEndConfigurations/my-frontend.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)


