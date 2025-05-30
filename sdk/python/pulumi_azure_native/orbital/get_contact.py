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
    'GetContactResult',
    'AwaitableGetContactResult',
    'get_contact',
    'get_contact_output',
]

@pulumi.output_type
class GetContactResult:
    """
    Customer creates a contact resource for a spacecraft resource.
    """
    def __init__(__self__, antenna_configuration=None, azure_api_version=None, contact_profile=None, end_azimuth_degrees=None, end_elevation_degrees=None, error_message=None, ground_station_name=None, id=None, maximum_elevation_degrees=None, name=None, reservation_end_time=None, reservation_start_time=None, rx_end_time=None, rx_start_time=None, start_azimuth_degrees=None, start_elevation_degrees=None, status=None, system_data=None, tx_end_time=None, tx_start_time=None, type=None):
        if antenna_configuration and not isinstance(antenna_configuration, dict):
            raise TypeError("Expected argument 'antenna_configuration' to be a dict")
        pulumi.set(__self__, "antenna_configuration", antenna_configuration)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if contact_profile and not isinstance(contact_profile, dict):
            raise TypeError("Expected argument 'contact_profile' to be a dict")
        pulumi.set(__self__, "contact_profile", contact_profile)
        if end_azimuth_degrees and not isinstance(end_azimuth_degrees, float):
            raise TypeError("Expected argument 'end_azimuth_degrees' to be a float")
        pulumi.set(__self__, "end_azimuth_degrees", end_azimuth_degrees)
        if end_elevation_degrees and not isinstance(end_elevation_degrees, float):
            raise TypeError("Expected argument 'end_elevation_degrees' to be a float")
        pulumi.set(__self__, "end_elevation_degrees", end_elevation_degrees)
        if error_message and not isinstance(error_message, str):
            raise TypeError("Expected argument 'error_message' to be a str")
        pulumi.set(__self__, "error_message", error_message)
        if ground_station_name and not isinstance(ground_station_name, str):
            raise TypeError("Expected argument 'ground_station_name' to be a str")
        pulumi.set(__self__, "ground_station_name", ground_station_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if maximum_elevation_degrees and not isinstance(maximum_elevation_degrees, float):
            raise TypeError("Expected argument 'maximum_elevation_degrees' to be a float")
        pulumi.set(__self__, "maximum_elevation_degrees", maximum_elevation_degrees)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if reservation_end_time and not isinstance(reservation_end_time, str):
            raise TypeError("Expected argument 'reservation_end_time' to be a str")
        pulumi.set(__self__, "reservation_end_time", reservation_end_time)
        if reservation_start_time and not isinstance(reservation_start_time, str):
            raise TypeError("Expected argument 'reservation_start_time' to be a str")
        pulumi.set(__self__, "reservation_start_time", reservation_start_time)
        if rx_end_time and not isinstance(rx_end_time, str):
            raise TypeError("Expected argument 'rx_end_time' to be a str")
        pulumi.set(__self__, "rx_end_time", rx_end_time)
        if rx_start_time and not isinstance(rx_start_time, str):
            raise TypeError("Expected argument 'rx_start_time' to be a str")
        pulumi.set(__self__, "rx_start_time", rx_start_time)
        if start_azimuth_degrees and not isinstance(start_azimuth_degrees, float):
            raise TypeError("Expected argument 'start_azimuth_degrees' to be a float")
        pulumi.set(__self__, "start_azimuth_degrees", start_azimuth_degrees)
        if start_elevation_degrees and not isinstance(start_elevation_degrees, float):
            raise TypeError("Expected argument 'start_elevation_degrees' to be a float")
        pulumi.set(__self__, "start_elevation_degrees", start_elevation_degrees)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tx_end_time and not isinstance(tx_end_time, str):
            raise TypeError("Expected argument 'tx_end_time' to be a str")
        pulumi.set(__self__, "tx_end_time", tx_end_time)
        if tx_start_time and not isinstance(tx_start_time, str):
            raise TypeError("Expected argument 'tx_start_time' to be a str")
        pulumi.set(__self__, "tx_start_time", tx_start_time)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="antennaConfiguration")
    def antenna_configuration(self) -> 'outputs.ContactsPropertiesResponseAntennaConfiguration':
        """
        The configuration associated with the allocated antenna.
        """
        return pulumi.get(self, "antenna_configuration")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="contactProfile")
    def contact_profile(self) -> 'outputs.ContactsPropertiesResponseContactProfile':
        """
        The reference to the contact profile resource.
        """
        return pulumi.get(self, "contact_profile")

    @property
    @pulumi.getter(name="endAzimuthDegrees")
    def end_azimuth_degrees(self) -> builtins.float:
        """
        Azimuth of the antenna at the end of the contact in decimal degrees.
        """
        return pulumi.get(self, "end_azimuth_degrees")

    @property
    @pulumi.getter(name="endElevationDegrees")
    def end_elevation_degrees(self) -> builtins.float:
        """
        Spacecraft elevation above the horizon at contact end.
        """
        return pulumi.get(self, "end_elevation_degrees")

    @property
    @pulumi.getter(name="errorMessage")
    def error_message(self) -> builtins.str:
        """
        Any error message while scheduling a contact.
        """
        return pulumi.get(self, "error_message")

    @property
    @pulumi.getter(name="groundStationName")
    def ground_station_name(self) -> builtins.str:
        """
        Azure Ground Station name.
        """
        return pulumi.get(self, "ground_station_name")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="maximumElevationDegrees")
    def maximum_elevation_degrees(self) -> builtins.float:
        """
        Maximum elevation of the antenna during the contact in decimal degrees.
        """
        return pulumi.get(self, "maximum_elevation_degrees")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="reservationEndTime")
    def reservation_end_time(self) -> builtins.str:
        """
        Reservation end time of a contact (ISO 8601 UTC standard).
        """
        return pulumi.get(self, "reservation_end_time")

    @property
    @pulumi.getter(name="reservationStartTime")
    def reservation_start_time(self) -> builtins.str:
        """
        Reservation start time of a contact (ISO 8601 UTC standard).
        """
        return pulumi.get(self, "reservation_start_time")

    @property
    @pulumi.getter(name="rxEndTime")
    def rx_end_time(self) -> builtins.str:
        """
        Receive end time of a contact (ISO 8601 UTC standard).
        """
        return pulumi.get(self, "rx_end_time")

    @property
    @pulumi.getter(name="rxStartTime")
    def rx_start_time(self) -> builtins.str:
        """
        Receive start time of a contact (ISO 8601 UTC standard).
        """
        return pulumi.get(self, "rx_start_time")

    @property
    @pulumi.getter(name="startAzimuthDegrees")
    def start_azimuth_degrees(self) -> builtins.float:
        """
        Azimuth of the antenna at the start of the contact in decimal degrees.
        """
        return pulumi.get(self, "start_azimuth_degrees")

    @property
    @pulumi.getter(name="startElevationDegrees")
    def start_elevation_degrees(self) -> builtins.float:
        """
        Spacecraft elevation above the horizon at contact start.
        """
        return pulumi.get(self, "start_elevation_degrees")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        Status of a contact.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter(name="txEndTime")
    def tx_end_time(self) -> builtins.str:
        """
        Transmit end time of a contact (ISO 8601 UTC standard).
        """
        return pulumi.get(self, "tx_end_time")

    @property
    @pulumi.getter(name="txStartTime")
    def tx_start_time(self) -> builtins.str:
        """
        Transmit start time of a contact (ISO 8601 UTC standard).
        """
        return pulumi.get(self, "tx_start_time")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetContactResult(GetContactResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContactResult(
            antenna_configuration=self.antenna_configuration,
            azure_api_version=self.azure_api_version,
            contact_profile=self.contact_profile,
            end_azimuth_degrees=self.end_azimuth_degrees,
            end_elevation_degrees=self.end_elevation_degrees,
            error_message=self.error_message,
            ground_station_name=self.ground_station_name,
            id=self.id,
            maximum_elevation_degrees=self.maximum_elevation_degrees,
            name=self.name,
            reservation_end_time=self.reservation_end_time,
            reservation_start_time=self.reservation_start_time,
            rx_end_time=self.rx_end_time,
            rx_start_time=self.rx_start_time,
            start_azimuth_degrees=self.start_azimuth_degrees,
            start_elevation_degrees=self.start_elevation_degrees,
            status=self.status,
            system_data=self.system_data,
            tx_end_time=self.tx_end_time,
            tx_start_time=self.tx_start_time,
            type=self.type)


def get_contact(contact_name: Optional[builtins.str] = None,
                resource_group_name: Optional[builtins.str] = None,
                spacecraft_name: Optional[builtins.str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContactResult:
    """
    Gets the specified contact in a specified resource group.

    Uses Azure REST API version 2022-11-01.


    :param builtins.str contact_name: Contact name.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str spacecraft_name: Spacecraft ID.
    """
    __args__ = dict()
    __args__['contactName'] = contact_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['spacecraftName'] = spacecraft_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:orbital:getContact', __args__, opts=opts, typ=GetContactResult).value

    return AwaitableGetContactResult(
        antenna_configuration=pulumi.get(__ret__, 'antenna_configuration'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        contact_profile=pulumi.get(__ret__, 'contact_profile'),
        end_azimuth_degrees=pulumi.get(__ret__, 'end_azimuth_degrees'),
        end_elevation_degrees=pulumi.get(__ret__, 'end_elevation_degrees'),
        error_message=pulumi.get(__ret__, 'error_message'),
        ground_station_name=pulumi.get(__ret__, 'ground_station_name'),
        id=pulumi.get(__ret__, 'id'),
        maximum_elevation_degrees=pulumi.get(__ret__, 'maximum_elevation_degrees'),
        name=pulumi.get(__ret__, 'name'),
        reservation_end_time=pulumi.get(__ret__, 'reservation_end_time'),
        reservation_start_time=pulumi.get(__ret__, 'reservation_start_time'),
        rx_end_time=pulumi.get(__ret__, 'rx_end_time'),
        rx_start_time=pulumi.get(__ret__, 'rx_start_time'),
        start_azimuth_degrees=pulumi.get(__ret__, 'start_azimuth_degrees'),
        start_elevation_degrees=pulumi.get(__ret__, 'start_elevation_degrees'),
        status=pulumi.get(__ret__, 'status'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tx_end_time=pulumi.get(__ret__, 'tx_end_time'),
        tx_start_time=pulumi.get(__ret__, 'tx_start_time'),
        type=pulumi.get(__ret__, 'type'))
def get_contact_output(contact_name: Optional[pulumi.Input[builtins.str]] = None,
                       resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                       spacecraft_name: Optional[pulumi.Input[builtins.str]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetContactResult]:
    """
    Gets the specified contact in a specified resource group.

    Uses Azure REST API version 2022-11-01.


    :param builtins.str contact_name: Contact name.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str spacecraft_name: Spacecraft ID.
    """
    __args__ = dict()
    __args__['contactName'] = contact_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['spacecraftName'] = spacecraft_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:orbital:getContact', __args__, opts=opts, typ=GetContactResult)
    return __ret__.apply(lambda __response__: GetContactResult(
        antenna_configuration=pulumi.get(__response__, 'antenna_configuration'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        contact_profile=pulumi.get(__response__, 'contact_profile'),
        end_azimuth_degrees=pulumi.get(__response__, 'end_azimuth_degrees'),
        end_elevation_degrees=pulumi.get(__response__, 'end_elevation_degrees'),
        error_message=pulumi.get(__response__, 'error_message'),
        ground_station_name=pulumi.get(__response__, 'ground_station_name'),
        id=pulumi.get(__response__, 'id'),
        maximum_elevation_degrees=pulumi.get(__response__, 'maximum_elevation_degrees'),
        name=pulumi.get(__response__, 'name'),
        reservation_end_time=pulumi.get(__response__, 'reservation_end_time'),
        reservation_start_time=pulumi.get(__response__, 'reservation_start_time'),
        rx_end_time=pulumi.get(__response__, 'rx_end_time'),
        rx_start_time=pulumi.get(__response__, 'rx_start_time'),
        start_azimuth_degrees=pulumi.get(__response__, 'start_azimuth_degrees'),
        start_elevation_degrees=pulumi.get(__response__, 'start_elevation_degrees'),
        status=pulumi.get(__response__, 'status'),
        system_data=pulumi.get(__response__, 'system_data'),
        tx_end_time=pulumi.get(__response__, 'tx_end_time'),
        tx_start_time=pulumi.get(__response__, 'tx_start_time'),
        type=pulumi.get(__response__, 'type')))
