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
    'GetMaintenanceConfigurationResult',
    'AwaitableGetMaintenanceConfigurationResult',
    'get_maintenance_configuration',
    'get_maintenance_configuration_output',
]

@pulumi.output_type
class GetMaintenanceConfigurationResult:
    """
    Maintenance configuration record type
    """
    def __init__(__self__, azure_api_version=None, duration=None, expiration_date_time=None, extension_properties=None, id=None, install_patches=None, location=None, maintenance_scope=None, name=None, namespace=None, recur_every=None, start_date_time=None, system_data=None, tags=None, time_zone=None, type=None, visibility=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if duration and not isinstance(duration, str):
            raise TypeError("Expected argument 'duration' to be a str")
        pulumi.set(__self__, "duration", duration)
        if expiration_date_time and not isinstance(expiration_date_time, str):
            raise TypeError("Expected argument 'expiration_date_time' to be a str")
        pulumi.set(__self__, "expiration_date_time", expiration_date_time)
        if extension_properties and not isinstance(extension_properties, dict):
            raise TypeError("Expected argument 'extension_properties' to be a dict")
        pulumi.set(__self__, "extension_properties", extension_properties)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if install_patches and not isinstance(install_patches, dict):
            raise TypeError("Expected argument 'install_patches' to be a dict")
        pulumi.set(__self__, "install_patches", install_patches)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if maintenance_scope and not isinstance(maintenance_scope, str):
            raise TypeError("Expected argument 'maintenance_scope' to be a str")
        pulumi.set(__self__, "maintenance_scope", maintenance_scope)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if recur_every and not isinstance(recur_every, str):
            raise TypeError("Expected argument 'recur_every' to be a str")
        pulumi.set(__self__, "recur_every", recur_every)
        if start_date_time and not isinstance(start_date_time, str):
            raise TypeError("Expected argument 'start_date_time' to be a str")
        pulumi.set(__self__, "start_date_time", start_date_time)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if time_zone and not isinstance(time_zone, str):
            raise TypeError("Expected argument 'time_zone' to be a str")
        pulumi.set(__self__, "time_zone", time_zone)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def duration(self) -> Optional[builtins.str]:
        """
        Duration of the maintenance window in HH:mm format. If not provided, default value will be used based on maintenance scope provided. Example: 05:00.
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter(name="expirationDateTime")
    def expiration_date_time(self) -> Optional[builtins.str]:
        """
        Effective expiration date of the maintenance window in YYYY-MM-DD hh:mm format. The window will be created in the time zone provided and adjusted to daylight savings according to that time zone. Expiration date must be set to a future date. If not provided, it will be set to the maximum datetime 9999-12-31 23:59:59.
        """
        return pulumi.get(self, "expiration_date_time")

    @property
    @pulumi.getter(name="extensionProperties")
    def extension_properties(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Gets or sets extensionProperties of the maintenanceConfiguration
        """
        return pulumi.get(self, "extension_properties")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="installPatches")
    def install_patches(self) -> Optional['outputs.InputPatchConfigurationResponse']:
        """
        The input parameters to be passed to the patch run operation.
        """
        return pulumi.get(self, "install_patches")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maintenanceScope")
    def maintenance_scope(self) -> Optional[builtins.str]:
        """
        Gets or sets maintenanceScope of the configuration
        """
        return pulumi.get(self, "maintenance_scope")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[builtins.str]:
        """
        Gets or sets namespace of the resource
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="recurEvery")
    def recur_every(self) -> Optional[builtins.str]:
        """
        Rate at which a Maintenance window is expected to recur. The rate can be expressed as daily, weekly, or monthly schedules. Daily schedule are formatted as recurEvery: [Frequency as integer]['Day(s)']. If no frequency is provided, the default frequency is 1. Daily schedule examples are recurEvery: Day, recurEvery: 3Days.  Weekly schedule are formatted as recurEvery: [Frequency as integer]['Week(s)'] [Optional comma separated list of weekdays Monday-Sunday]. Weekly schedule examples are recurEvery: 3Weeks, recurEvery: Week Saturday,Sunday. Monthly schedules are formatted as [Frequency as integer]['Month(s)'] [Comma separated list of month days] or [Frequency as integer]['Month(s)'] [Week of Month (First, Second, Third, Fourth, Last)] [Weekday Monday-Sunday] [Optional Offset(No. of days)]. Offset value must be between -6 to 6 inclusive. Monthly schedule examples are recurEvery: Month, recurEvery: 2Months, recurEvery: Month day23,day24, recurEvery: Month Last Sunday, recurEvery: Month Fourth Monday, recurEvery: Month Last Sunday Offset-3, recurEvery: Month Third Sunday Offset6.
        """
        return pulumi.get(self, "recur_every")

    @property
    @pulumi.getter(name="startDateTime")
    def start_date_time(self) -> Optional[builtins.str]:
        """
        Effective start date of the maintenance window in YYYY-MM-DD hh:mm format. The start date can be set to either the current date or future date. The window will be created in the time zone provided and adjusted to daylight savings according to that time zone.
        """
        return pulumi.get(self, "start_date_time")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> Optional[builtins.str]:
        """
        Name of the timezone. List of timezones can be obtained by executing [System.TimeZoneInfo]::GetSystemTimeZones() in PowerShell. Example: Pacific Standard Time, UTC, W. Europe Standard Time, Korea Standard Time, Cen. Australia Standard Time.
        """
        return pulumi.get(self, "time_zone")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def visibility(self) -> Optional[builtins.str]:
        """
        Gets or sets the visibility of the configuration. The default value is 'Custom'
        """
        return pulumi.get(self, "visibility")


class AwaitableGetMaintenanceConfigurationResult(GetMaintenanceConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMaintenanceConfigurationResult(
            azure_api_version=self.azure_api_version,
            duration=self.duration,
            expiration_date_time=self.expiration_date_time,
            extension_properties=self.extension_properties,
            id=self.id,
            install_patches=self.install_patches,
            location=self.location,
            maintenance_scope=self.maintenance_scope,
            name=self.name,
            namespace=self.namespace,
            recur_every=self.recur_every,
            start_date_time=self.start_date_time,
            system_data=self.system_data,
            tags=self.tags,
            time_zone=self.time_zone,
            type=self.type,
            visibility=self.visibility)


def get_maintenance_configuration(resource_group_name: Optional[builtins.str] = None,
                                  resource_name: Optional[builtins.str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMaintenanceConfigurationResult:
    """
    Get Configuration record

    Uses Azure REST API version 2023-10-01-preview.

    Other available API versions: 2022-11-01-preview, 2023-04-01, 2023-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native maintenance [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str resource_name: The name of the MaintenanceConfiguration
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:maintenance:getMaintenanceConfiguration', __args__, opts=opts, typ=GetMaintenanceConfigurationResult).value

    return AwaitableGetMaintenanceConfigurationResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        duration=pulumi.get(__ret__, 'duration'),
        expiration_date_time=pulumi.get(__ret__, 'expiration_date_time'),
        extension_properties=pulumi.get(__ret__, 'extension_properties'),
        id=pulumi.get(__ret__, 'id'),
        install_patches=pulumi.get(__ret__, 'install_patches'),
        location=pulumi.get(__ret__, 'location'),
        maintenance_scope=pulumi.get(__ret__, 'maintenance_scope'),
        name=pulumi.get(__ret__, 'name'),
        namespace=pulumi.get(__ret__, 'namespace'),
        recur_every=pulumi.get(__ret__, 'recur_every'),
        start_date_time=pulumi.get(__ret__, 'start_date_time'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        time_zone=pulumi.get(__ret__, 'time_zone'),
        type=pulumi.get(__ret__, 'type'),
        visibility=pulumi.get(__ret__, 'visibility'))
def get_maintenance_configuration_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                         resource_name: Optional[pulumi.Input[builtins.str]] = None,
                                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMaintenanceConfigurationResult]:
    """
    Get Configuration record

    Uses Azure REST API version 2023-10-01-preview.

    Other available API versions: 2022-11-01-preview, 2023-04-01, 2023-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native maintenance [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str resource_name: The name of the MaintenanceConfiguration
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:maintenance:getMaintenanceConfiguration', __args__, opts=opts, typ=GetMaintenanceConfigurationResult)
    return __ret__.apply(lambda __response__: GetMaintenanceConfigurationResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        duration=pulumi.get(__response__, 'duration'),
        expiration_date_time=pulumi.get(__response__, 'expiration_date_time'),
        extension_properties=pulumi.get(__response__, 'extension_properties'),
        id=pulumi.get(__response__, 'id'),
        install_patches=pulumi.get(__response__, 'install_patches'),
        location=pulumi.get(__response__, 'location'),
        maintenance_scope=pulumi.get(__response__, 'maintenance_scope'),
        name=pulumi.get(__response__, 'name'),
        namespace=pulumi.get(__response__, 'namespace'),
        recur_every=pulumi.get(__response__, 'recur_every'),
        start_date_time=pulumi.get(__response__, 'start_date_time'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        time_zone=pulumi.get(__response__, 'time_zone'),
        type=pulumi.get(__response__, 'type'),
        visibility=pulumi.get(__response__, 'visibility')))
