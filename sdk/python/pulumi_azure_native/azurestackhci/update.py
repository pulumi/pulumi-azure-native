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
from ._inputs import *

__all__ = ['UpdateArgs', 'Update']

@pulumi.input_type
class UpdateArgs:
    def __init__(__self__, *,
                 cluster_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 additional_properties: Optional[pulumi.Input[builtins.str]] = None,
                 availability_type: Optional[pulumi.Input[Union[builtins.str, 'AvailabilityType']]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 health_check_date: Optional[pulumi.Input[builtins.str]] = None,
                 installed_date: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 min_sbe_version_required: Optional[pulumi.Input[builtins.str]] = None,
                 notify_message: Optional[pulumi.Input[builtins.str]] = None,
                 package_path: Optional[pulumi.Input[builtins.str]] = None,
                 package_size_in_mb: Optional[pulumi.Input[builtins.float]] = None,
                 package_type: Optional[pulumi.Input[builtins.str]] = None,
                 prerequisites: Optional[pulumi.Input[Sequence[pulumi.Input['UpdatePrerequisiteArgs']]]] = None,
                 progress_percentage: Optional[pulumi.Input[builtins.float]] = None,
                 publisher: Optional[pulumi.Input[builtins.str]] = None,
                 release_link: Optional[pulumi.Input[builtins.str]] = None,
                 state: Optional[pulumi.Input[Union[builtins.str, 'State']]] = None,
                 update_name: Optional[pulumi.Input[builtins.str]] = None,
                 version: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Update resource.
        :param pulumi.Input[builtins.str] cluster_name: The name of the cluster.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] additional_properties: Extensible KV pairs serialized as a string. This is currently used to report the stamp OEM family and hardware model information when an update is flagged as Invalid for the stamp based on OEM type.
        :param pulumi.Input[Union[builtins.str, 'AvailabilityType']] availability_type: Indicates the way the update content can be downloaded.
        :param pulumi.Input[builtins.str] description: Description of the update.
        :param pulumi.Input[builtins.str] display_name: Display name of the Update
        :param pulumi.Input[builtins.str] health_check_date: Last time the package-specific checks were run.
        :param pulumi.Input[builtins.str] installed_date: Date that the update was installed.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] min_sbe_version_required: Minimum Sbe Version of the update.
        :param pulumi.Input[builtins.str] notify_message: Brief message with instructions for updates of AvailabilityType Notify.
        :param pulumi.Input[builtins.str] package_path: Path where the update package is available.
        :param pulumi.Input[builtins.float] package_size_in_mb: Size of the package. This value is a combination of the size from update metadata and size of the payload that results from the live scan operation for OS update content.
        :param pulumi.Input[builtins.str] package_type: Customer-visible type of the update.
        :param pulumi.Input[Sequence[pulumi.Input['UpdatePrerequisiteArgs']]] prerequisites: If update State is HasPrerequisite, this property contains an array of objects describing prerequisite updates before installing this update. Otherwise, it is empty.
        :param pulumi.Input[builtins.float] progress_percentage: Progress percentage of ongoing operation. Currently this property is only valid when the update is in the Downloading state, where it maps to how much of the update content has been downloaded.
        :param pulumi.Input[builtins.str] publisher: Publisher of the update package.
        :param pulumi.Input[builtins.str] release_link: Link to release notes for the update.
        :param pulumi.Input[Union[builtins.str, 'State']] state: State of the update as it relates to this stamp.
        :param pulumi.Input[builtins.str] update_name: The name of the Update
        :param pulumi.Input[builtins.str] version: Version of the update.
        """
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if additional_properties is not None:
            pulumi.set(__self__, "additional_properties", additional_properties)
        if availability_type is not None:
            pulumi.set(__self__, "availability_type", availability_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if health_check_date is not None:
            pulumi.set(__self__, "health_check_date", health_check_date)
        if installed_date is not None:
            pulumi.set(__self__, "installed_date", installed_date)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if min_sbe_version_required is not None:
            pulumi.set(__self__, "min_sbe_version_required", min_sbe_version_required)
        if notify_message is not None:
            pulumi.set(__self__, "notify_message", notify_message)
        if package_path is not None:
            pulumi.set(__self__, "package_path", package_path)
        if package_size_in_mb is not None:
            pulumi.set(__self__, "package_size_in_mb", package_size_in_mb)
        if package_type is not None:
            pulumi.set(__self__, "package_type", package_type)
        if prerequisites is not None:
            pulumi.set(__self__, "prerequisites", prerequisites)
        if progress_percentage is not None:
            pulumi.set(__self__, "progress_percentage", progress_percentage)
        if publisher is not None:
            pulumi.set(__self__, "publisher", publisher)
        if release_link is not None:
            pulumi.set(__self__, "release_link", release_link)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if update_name is not None:
            pulumi.set(__self__, "update_name", update_name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the cluster.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group. The name is case insensitive.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="additionalProperties")
    def additional_properties(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Extensible KV pairs serialized as a string. This is currently used to report the stamp OEM family and hardware model information when an update is flagged as Invalid for the stamp based on OEM type.
        """
        return pulumi.get(self, "additional_properties")

    @additional_properties.setter
    def additional_properties(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "additional_properties", value)

    @property
    @pulumi.getter(name="availabilityType")
    def availability_type(self) -> Optional[pulumi.Input[Union[builtins.str, 'AvailabilityType']]]:
        """
        Indicates the way the update content can be downloaded.
        """
        return pulumi.get(self, "availability_type")

    @availability_type.setter
    def availability_type(self, value: Optional[pulumi.Input[Union[builtins.str, 'AvailabilityType']]]):
        pulumi.set(self, "availability_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description of the update.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Display name of the Update
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="healthCheckDate")
    def health_check_date(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Last time the package-specific checks were run.
        """
        return pulumi.get(self, "health_check_date")

    @health_check_date.setter
    def health_check_date(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "health_check_date", value)

    @property
    @pulumi.getter(name="installedDate")
    def installed_date(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Date that the update was installed.
        """
        return pulumi.get(self, "installed_date")

    @installed_date.setter
    def installed_date(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "installed_date", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="minSbeVersionRequired")
    def min_sbe_version_required(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Minimum Sbe Version of the update.
        """
        return pulumi.get(self, "min_sbe_version_required")

    @min_sbe_version_required.setter
    def min_sbe_version_required(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "min_sbe_version_required", value)

    @property
    @pulumi.getter(name="notifyMessage")
    def notify_message(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Brief message with instructions for updates of AvailabilityType Notify.
        """
        return pulumi.get(self, "notify_message")

    @notify_message.setter
    def notify_message(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "notify_message", value)

    @property
    @pulumi.getter(name="packagePath")
    def package_path(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Path where the update package is available.
        """
        return pulumi.get(self, "package_path")

    @package_path.setter
    def package_path(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "package_path", value)

    @property
    @pulumi.getter(name="packageSizeInMb")
    def package_size_in_mb(self) -> Optional[pulumi.Input[builtins.float]]:
        """
        Size of the package. This value is a combination of the size from update metadata and size of the payload that results from the live scan operation for OS update content.
        """
        return pulumi.get(self, "package_size_in_mb")

    @package_size_in_mb.setter
    def package_size_in_mb(self, value: Optional[pulumi.Input[builtins.float]]):
        pulumi.set(self, "package_size_in_mb", value)

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Customer-visible type of the update.
        """
        return pulumi.get(self, "package_type")

    @package_type.setter
    def package_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "package_type", value)

    @property
    @pulumi.getter
    def prerequisites(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UpdatePrerequisiteArgs']]]]:
        """
        If update State is HasPrerequisite, this property contains an array of objects describing prerequisite updates before installing this update. Otherwise, it is empty.
        """
        return pulumi.get(self, "prerequisites")

    @prerequisites.setter
    def prerequisites(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UpdatePrerequisiteArgs']]]]):
        pulumi.set(self, "prerequisites", value)

    @property
    @pulumi.getter(name="progressPercentage")
    def progress_percentage(self) -> Optional[pulumi.Input[builtins.float]]:
        """
        Progress percentage of ongoing operation. Currently this property is only valid when the update is in the Downloading state, where it maps to how much of the update content has been downloaded.
        """
        return pulumi.get(self, "progress_percentage")

    @progress_percentage.setter
    def progress_percentage(self, value: Optional[pulumi.Input[builtins.float]]):
        pulumi.set(self, "progress_percentage", value)

    @property
    @pulumi.getter
    def publisher(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Publisher of the update package.
        """
        return pulumi.get(self, "publisher")

    @publisher.setter
    def publisher(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "publisher", value)

    @property
    @pulumi.getter(name="releaseLink")
    def release_link(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Link to release notes for the update.
        """
        return pulumi.get(self, "release_link")

    @release_link.setter
    def release_link(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "release_link", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[Union[builtins.str, 'State']]]:
        """
        State of the update as it relates to this stamp.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[Union[builtins.str, 'State']]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="updateName")
    def update_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the Update
        """
        return pulumi.get(self, "update_name")

    @update_name.setter
    def update_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "update_name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Version of the update.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "version", value)


@pulumi.type_token("azure-native:azurestackhci:Update")
class Update(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_properties: Optional[pulumi.Input[builtins.str]] = None,
                 availability_type: Optional[pulumi.Input[Union[builtins.str, 'AvailabilityType']]] = None,
                 cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 health_check_date: Optional[pulumi.Input[builtins.str]] = None,
                 installed_date: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 min_sbe_version_required: Optional[pulumi.Input[builtins.str]] = None,
                 notify_message: Optional[pulumi.Input[builtins.str]] = None,
                 package_path: Optional[pulumi.Input[builtins.str]] = None,
                 package_size_in_mb: Optional[pulumi.Input[builtins.float]] = None,
                 package_type: Optional[pulumi.Input[builtins.str]] = None,
                 prerequisites: Optional[pulumi.Input[Sequence[pulumi.Input[Union['UpdatePrerequisiteArgs', 'UpdatePrerequisiteArgsDict']]]]] = None,
                 progress_percentage: Optional[pulumi.Input[builtins.float]] = None,
                 publisher: Optional[pulumi.Input[builtins.str]] = None,
                 release_link: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 state: Optional[pulumi.Input[Union[builtins.str, 'State']]] = None,
                 update_name: Optional[pulumi.Input[builtins.str]] = None,
                 version: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Update details

        Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01.

        Other available API versions: 2022-12-15-preview, 2023-02-01, 2023-03-01, 2023-06-01, 2023-08-01, 2023-08-01-preview, 2023-11-01-preview, 2024-01-01, 2024-02-15-preview, 2024-09-01-preview, 2024-12-01-preview, 2025-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestackhci [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] additional_properties: Extensible KV pairs serialized as a string. This is currently used to report the stamp OEM family and hardware model information when an update is flagged as Invalid for the stamp based on OEM type.
        :param pulumi.Input[Union[builtins.str, 'AvailabilityType']] availability_type: Indicates the way the update content can be downloaded.
        :param pulumi.Input[builtins.str] cluster_name: The name of the cluster.
        :param pulumi.Input[builtins.str] description: Description of the update.
        :param pulumi.Input[builtins.str] display_name: Display name of the Update
        :param pulumi.Input[builtins.str] health_check_date: Last time the package-specific checks were run.
        :param pulumi.Input[builtins.str] installed_date: Date that the update was installed.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] min_sbe_version_required: Minimum Sbe Version of the update.
        :param pulumi.Input[builtins.str] notify_message: Brief message with instructions for updates of AvailabilityType Notify.
        :param pulumi.Input[builtins.str] package_path: Path where the update package is available.
        :param pulumi.Input[builtins.float] package_size_in_mb: Size of the package. This value is a combination of the size from update metadata and size of the payload that results from the live scan operation for OS update content.
        :param pulumi.Input[builtins.str] package_type: Customer-visible type of the update.
        :param pulumi.Input[Sequence[pulumi.Input[Union['UpdatePrerequisiteArgs', 'UpdatePrerequisiteArgsDict']]]] prerequisites: If update State is HasPrerequisite, this property contains an array of objects describing prerequisite updates before installing this update. Otherwise, it is empty.
        :param pulumi.Input[builtins.float] progress_percentage: Progress percentage of ongoing operation. Currently this property is only valid when the update is in the Downloading state, where it maps to how much of the update content has been downloaded.
        :param pulumi.Input[builtins.str] publisher: Publisher of the update package.
        :param pulumi.Input[builtins.str] release_link: Link to release notes for the update.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Union[builtins.str, 'State']] state: State of the update as it relates to this stamp.
        :param pulumi.Input[builtins.str] update_name: The name of the Update
        :param pulumi.Input[builtins.str] version: Version of the update.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UpdateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Update details

        Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01.

        Other available API versions: 2022-12-15-preview, 2023-02-01, 2023-03-01, 2023-06-01, 2023-08-01, 2023-08-01-preview, 2023-11-01-preview, 2024-01-01, 2024-02-15-preview, 2024-09-01-preview, 2024-12-01-preview, 2025-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestackhci [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param UpdateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UpdateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_properties: Optional[pulumi.Input[builtins.str]] = None,
                 availability_type: Optional[pulumi.Input[Union[builtins.str, 'AvailabilityType']]] = None,
                 cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 health_check_date: Optional[pulumi.Input[builtins.str]] = None,
                 installed_date: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 min_sbe_version_required: Optional[pulumi.Input[builtins.str]] = None,
                 notify_message: Optional[pulumi.Input[builtins.str]] = None,
                 package_path: Optional[pulumi.Input[builtins.str]] = None,
                 package_size_in_mb: Optional[pulumi.Input[builtins.float]] = None,
                 package_type: Optional[pulumi.Input[builtins.str]] = None,
                 prerequisites: Optional[pulumi.Input[Sequence[pulumi.Input[Union['UpdatePrerequisiteArgs', 'UpdatePrerequisiteArgsDict']]]]] = None,
                 progress_percentage: Optional[pulumi.Input[builtins.float]] = None,
                 publisher: Optional[pulumi.Input[builtins.str]] = None,
                 release_link: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 state: Optional[pulumi.Input[Union[builtins.str, 'State']]] = None,
                 update_name: Optional[pulumi.Input[builtins.str]] = None,
                 version: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UpdateArgs.__new__(UpdateArgs)

            __props__.__dict__["additional_properties"] = additional_properties
            __props__.__dict__["availability_type"] = availability_type
            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["health_check_date"] = health_check_date
            __props__.__dict__["installed_date"] = installed_date
            __props__.__dict__["location"] = location
            __props__.__dict__["min_sbe_version_required"] = min_sbe_version_required
            __props__.__dict__["notify_message"] = notify_message
            __props__.__dict__["package_path"] = package_path
            __props__.__dict__["package_size_in_mb"] = package_size_in_mb
            __props__.__dict__["package_type"] = package_type
            __props__.__dict__["prerequisites"] = prerequisites
            __props__.__dict__["progress_percentage"] = progress_percentage
            __props__.__dict__["publisher"] = publisher
            __props__.__dict__["release_link"] = release_link
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["state"] = state
            __props__.__dict__["update_name"] = update_name
            __props__.__dict__["version"] = version
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:azurestackhci/v20221201:Update"), pulumi.Alias(type_="azure-native:azurestackhci/v20221215preview:Update"), pulumi.Alias(type_="azure-native:azurestackhci/v20230201:Update"), pulumi.Alias(type_="azure-native:azurestackhci/v20230301:Update"), pulumi.Alias(type_="azure-native:azurestackhci/v20230601:Update"), pulumi.Alias(type_="azure-native:azurestackhci/v20230801:Update"), pulumi.Alias(type_="azure-native:azurestackhci/v20230801preview:Update"), pulumi.Alias(type_="azure-native:azurestackhci/v20231101preview:Update"), pulumi.Alias(type_="azure-native:azurestackhci/v20240101:Update"), pulumi.Alias(type_="azure-native:azurestackhci/v20240215preview:Update"), pulumi.Alias(type_="azure-native:azurestackhci/v20240401:Update"), pulumi.Alias(type_="azure-native:azurestackhci/v20240901preview:Update"), pulumi.Alias(type_="azure-native:azurestackhci/v20241201preview:Update"), pulumi.Alias(type_="azure-native:azurestackhci/v20250201preview:Update")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Update, __self__).__init__(
            'azure-native:azurestackhci:Update',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Update':
        """
        Get an existing Update resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = UpdateArgs.__new__(UpdateArgs)

        __props__.__dict__["additional_properties"] = None
        __props__.__dict__["availability_type"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["health_check_date"] = None
        __props__.__dict__["installed_date"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["min_sbe_version_required"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["notify_message"] = None
        __props__.__dict__["package_path"] = None
        __props__.__dict__["package_size_in_mb"] = None
        __props__.__dict__["package_type"] = None
        __props__.__dict__["prerequisites"] = None
        __props__.__dict__["progress_percentage"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["publisher"] = None
        __props__.__dict__["release_link"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["version"] = None
        return Update(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalProperties")
    def additional_properties(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Extensible KV pairs serialized as a string. This is currently used to report the stamp OEM family and hardware model information when an update is flagged as Invalid for the stamp based on OEM type.
        """
        return pulumi.get(self, "additional_properties")

    @property
    @pulumi.getter(name="availabilityType")
    def availability_type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Indicates the way the update content can be downloaded.
        """
        return pulumi.get(self, "availability_type")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Description of the update.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Display name of the Update
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="healthCheckDate")
    def health_check_date(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Last time the package-specific checks were run.
        """
        return pulumi.get(self, "health_check_date")

    @property
    @pulumi.getter(name="installedDate")
    def installed_date(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Date that the update was installed.
        """
        return pulumi.get(self, "installed_date")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="minSbeVersionRequired")
    def min_sbe_version_required(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Minimum Sbe Version of the update.
        """
        return pulumi.get(self, "min_sbe_version_required")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notifyMessage")
    def notify_message(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Brief message with instructions for updates of AvailabilityType Notify.
        """
        return pulumi.get(self, "notify_message")

    @property
    @pulumi.getter(name="packagePath")
    def package_path(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Path where the update package is available.
        """
        return pulumi.get(self, "package_path")

    @property
    @pulumi.getter(name="packageSizeInMb")
    def package_size_in_mb(self) -> pulumi.Output[Optional[builtins.float]]:
        """
        Size of the package. This value is a combination of the size from update metadata and size of the payload that results from the live scan operation for OS update content.
        """
        return pulumi.get(self, "package_size_in_mb")

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Customer-visible type of the update.
        """
        return pulumi.get(self, "package_type")

    @property
    @pulumi.getter
    def prerequisites(self) -> pulumi.Output[Optional[Sequence['outputs.UpdatePrerequisiteResponse']]]:
        """
        If update State is HasPrerequisite, this property contains an array of objects describing prerequisite updates before installing this update. Otherwise, it is empty.
        """
        return pulumi.get(self, "prerequisites")

    @property
    @pulumi.getter(name="progressPercentage")
    def progress_percentage(self) -> pulumi.Output[Optional[builtins.float]]:
        """
        Progress percentage of ongoing operation. Currently this property is only valid when the update is in the Downloading state, where it maps to how much of the update content has been downloaded.
        """
        return pulumi.get(self, "progress_percentage")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Provisioning state of the Updates proxy resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def publisher(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Publisher of the update package.
        """
        return pulumi.get(self, "publisher")

    @property
    @pulumi.getter(name="releaseLink")
    def release_link(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Link to release notes for the update.
        """
        return pulumi.get(self, "release_link")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        State of the update as it relates to this stamp.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Version of the update.
        """
        return pulumi.get(self, "version")

