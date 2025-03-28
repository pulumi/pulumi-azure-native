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
    'ClusterDesiredPropertiesArgs',
    'ClusterDesiredPropertiesArgsDict',
    'SoftwareAssurancePropertiesArgs',
    'SoftwareAssurancePropertiesArgsDict',
]

MYPY = False

if not MYPY:
    class ClusterDesiredPropertiesArgsDict(TypedDict):
        """
        Desired properties of the cluster.
        """
        diagnostic_level: NotRequired[pulumi.Input[Union[str, 'DiagnosticLevel']]]
        """
        Desired level of diagnostic data emitted by the cluster.
        """
        windows_server_subscription: NotRequired[pulumi.Input[Union[str, 'WindowsServerSubscription']]]
        """
        Desired state of Windows Server Subscription.
        """
elif False:
    ClusterDesiredPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ClusterDesiredPropertiesArgs:
    def __init__(__self__, *,
                 diagnostic_level: Optional[pulumi.Input[Union[str, 'DiagnosticLevel']]] = None,
                 windows_server_subscription: Optional[pulumi.Input[Union[str, 'WindowsServerSubscription']]] = None):
        """
        Desired properties of the cluster.
        :param pulumi.Input[Union[str, 'DiagnosticLevel']] diagnostic_level: Desired level of diagnostic data emitted by the cluster.
        :param pulumi.Input[Union[str, 'WindowsServerSubscription']] windows_server_subscription: Desired state of Windows Server Subscription.
        """
        if diagnostic_level is not None:
            pulumi.set(__self__, "diagnostic_level", diagnostic_level)
        if windows_server_subscription is not None:
            pulumi.set(__self__, "windows_server_subscription", windows_server_subscription)

    @property
    @pulumi.getter(name="diagnosticLevel")
    def diagnostic_level(self) -> Optional[pulumi.Input[Union[str, 'DiagnosticLevel']]]:
        """
        Desired level of diagnostic data emitted by the cluster.
        """
        return pulumi.get(self, "diagnostic_level")

    @diagnostic_level.setter
    def diagnostic_level(self, value: Optional[pulumi.Input[Union[str, 'DiagnosticLevel']]]):
        pulumi.set(self, "diagnostic_level", value)

    @property
    @pulumi.getter(name="windowsServerSubscription")
    def windows_server_subscription(self) -> Optional[pulumi.Input[Union[str, 'WindowsServerSubscription']]]:
        """
        Desired state of Windows Server Subscription.
        """
        return pulumi.get(self, "windows_server_subscription")

    @windows_server_subscription.setter
    def windows_server_subscription(self, value: Optional[pulumi.Input[Union[str, 'WindowsServerSubscription']]]):
        pulumi.set(self, "windows_server_subscription", value)


if not MYPY:
    class SoftwareAssurancePropertiesArgsDict(TypedDict):
        """
        Software Assurance properties of the cluster.
        """
        software_assurance_intent: NotRequired[pulumi.Input[Union[str, 'SoftwareAssuranceIntent']]]
        """
        Customer Intent for Software Assurance Benefit.
        """
        software_assurance_status: NotRequired[pulumi.Input[Union[str, 'SoftwareAssuranceStatus']]]
        """
        Status of the Software Assurance for the cluster.
        """
elif False:
    SoftwareAssurancePropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SoftwareAssurancePropertiesArgs:
    def __init__(__self__, *,
                 software_assurance_intent: Optional[pulumi.Input[Union[str, 'SoftwareAssuranceIntent']]] = None,
                 software_assurance_status: Optional[pulumi.Input[Union[str, 'SoftwareAssuranceStatus']]] = None):
        """
        Software Assurance properties of the cluster.
        :param pulumi.Input[Union[str, 'SoftwareAssuranceIntent']] software_assurance_intent: Customer Intent for Software Assurance Benefit.
        :param pulumi.Input[Union[str, 'SoftwareAssuranceStatus']] software_assurance_status: Status of the Software Assurance for the cluster.
        """
        if software_assurance_intent is not None:
            pulumi.set(__self__, "software_assurance_intent", software_assurance_intent)
        if software_assurance_status is not None:
            pulumi.set(__self__, "software_assurance_status", software_assurance_status)

    @property
    @pulumi.getter(name="softwareAssuranceIntent")
    def software_assurance_intent(self) -> Optional[pulumi.Input[Union[str, 'SoftwareAssuranceIntent']]]:
        """
        Customer Intent for Software Assurance Benefit.
        """
        return pulumi.get(self, "software_assurance_intent")

    @software_assurance_intent.setter
    def software_assurance_intent(self, value: Optional[pulumi.Input[Union[str, 'SoftwareAssuranceIntent']]]):
        pulumi.set(self, "software_assurance_intent", value)

    @property
    @pulumi.getter(name="softwareAssuranceStatus")
    def software_assurance_status(self) -> Optional[pulumi.Input[Union[str, 'SoftwareAssuranceStatus']]]:
        """
        Status of the Software Assurance for the cluster.
        """
        return pulumi.get(self, "software_assurance_status")

    @software_assurance_status.setter
    def software_assurance_status(self, value: Optional[pulumi.Input[Union[str, 'SoftwareAssuranceStatus']]]):
        pulumi.set(self, "software_assurance_status", value)


