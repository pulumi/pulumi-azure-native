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

__all__ = [
    'GetSapVirtualInstanceInvokeAvailabilityZoneDetailsResult',
    'AwaitableGetSapVirtualInstanceInvokeAvailabilityZoneDetailsResult',
    'get_sap_virtual_instance_invoke_availability_zone_details',
    'get_sap_virtual_instance_invoke_availability_zone_details_output',
]

@pulumi.output_type
class GetSapVirtualInstanceInvokeAvailabilityZoneDetailsResult:
    """
    The list of supported availability zone pairs which are part of SAP HA deployment.
    """
    def __init__(__self__, availability_zone_pairs=None):
        if availability_zone_pairs and not isinstance(availability_zone_pairs, list):
            raise TypeError("Expected argument 'availability_zone_pairs' to be a list")
        pulumi.set(__self__, "availability_zone_pairs", availability_zone_pairs)

    @property
    @pulumi.getter(name="availabilityZonePairs")
    def availability_zone_pairs(self) -> Optional[Sequence['outputs.SAPAvailabilityZonePairResponse']]:
        """
        Gets the list of availability zone pairs.
        """
        return pulumi.get(self, "availability_zone_pairs")


class AwaitableGetSapVirtualInstanceInvokeAvailabilityZoneDetailsResult(GetSapVirtualInstanceInvokeAvailabilityZoneDetailsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSapVirtualInstanceInvokeAvailabilityZoneDetailsResult(
            availability_zone_pairs=self.availability_zone_pairs)


def get_sap_virtual_instance_invoke_availability_zone_details(app_location: Optional[builtins.str] = None,
                                                              database_type: Optional[Union[builtins.str, 'SAPDatabaseType']] = None,
                                                              location: Optional[builtins.str] = None,
                                                              sap_product: Optional[Union[builtins.str, 'SAPProductType']] = None,
                                                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSapVirtualInstanceInvokeAvailabilityZoneDetailsResult:
    """
    Get the recommended SAP Availability Zone Pair Details for your region.

    Uses Azure REST API version 2024-09-01.


    :param builtins.str app_location: The geo-location where the SAP resources will be created.
    :param Union[builtins.str, 'SAPDatabaseType'] database_type: The database type. Eg: HANA, DB2, etc
    :param builtins.str location: The name of the Azure region.
    :param Union[builtins.str, 'SAPProductType'] sap_product: Defines the SAP Product type.
    """
    __args__ = dict()
    __args__['appLocation'] = app_location
    __args__['databaseType'] = database_type
    __args__['location'] = location
    __args__['sapProduct'] = sap_product
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:workloads:getSapVirtualInstanceInvokeAvailabilityZoneDetails', __args__, opts=opts, typ=GetSapVirtualInstanceInvokeAvailabilityZoneDetailsResult).value

    return AwaitableGetSapVirtualInstanceInvokeAvailabilityZoneDetailsResult(
        availability_zone_pairs=pulumi.get(__ret__, 'availability_zone_pairs'))
def get_sap_virtual_instance_invoke_availability_zone_details_output(app_location: Optional[pulumi.Input[builtins.str]] = None,
                                                                     database_type: Optional[pulumi.Input[Union[builtins.str, 'SAPDatabaseType']]] = None,
                                                                     location: Optional[pulumi.Input[builtins.str]] = None,
                                                                     sap_product: Optional[pulumi.Input[Union[builtins.str, 'SAPProductType']]] = None,
                                                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSapVirtualInstanceInvokeAvailabilityZoneDetailsResult]:
    """
    Get the recommended SAP Availability Zone Pair Details for your region.

    Uses Azure REST API version 2024-09-01.


    :param builtins.str app_location: The geo-location where the SAP resources will be created.
    :param Union[builtins.str, 'SAPDatabaseType'] database_type: The database type. Eg: HANA, DB2, etc
    :param builtins.str location: The name of the Azure region.
    :param Union[builtins.str, 'SAPProductType'] sap_product: Defines the SAP Product type.
    """
    __args__ = dict()
    __args__['appLocation'] = app_location
    __args__['databaseType'] = database_type
    __args__['location'] = location
    __args__['sapProduct'] = sap_product
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:workloads:getSapVirtualInstanceInvokeAvailabilityZoneDetails', __args__, opts=opts, typ=GetSapVirtualInstanceInvokeAvailabilityZoneDetailsResult)
    return __ret__.apply(lambda __response__: GetSapVirtualInstanceInvokeAvailabilityZoneDetailsResult(
        availability_zone_pairs=pulumi.get(__response__, 'availability_zone_pairs')))
