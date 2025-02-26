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
    'ApplicationLogsConfigArgs',
    'ApplicationLogsConfigArgsDict',
    'AzureBlobStorageApplicationLogsConfigArgs',
    'AzureBlobStorageApplicationLogsConfigArgsDict',
    'AzureBlobStorageHttpLogsConfigArgs',
    'AzureBlobStorageHttpLogsConfigArgsDict',
    'AzureTableStorageApplicationLogsConfigArgs',
    'AzureTableStorageApplicationLogsConfigArgsDict',
    'EnabledConfigArgs',
    'EnabledConfigArgsDict',
    'FileSystemApplicationLogsConfigArgs',
    'FileSystemApplicationLogsConfigArgsDict',
    'FileSystemHttpLogsConfigArgs',
    'FileSystemHttpLogsConfigArgsDict',
    'HttpLogsConfigArgs',
    'HttpLogsConfigArgsDict',
    'NameValuePairArgs',
    'NameValuePairArgsDict',
    'VirtualNetworkProfileArgs',
    'VirtualNetworkProfileArgsDict',
]

MYPY = False

if not MYPY:
    class ApplicationLogsConfigArgsDict(TypedDict):
        """
        Application logs configuration.
        """
        azure_blob_storage: NotRequired[pulumi.Input['AzureBlobStorageApplicationLogsConfigArgsDict']]
        """
        Application logs to blob storage configuration.
        """
        azure_table_storage: NotRequired[pulumi.Input['AzureTableStorageApplicationLogsConfigArgsDict']]
        """
        Application logs to azure table storage configuration.
        """
        file_system: NotRequired[pulumi.Input['FileSystemApplicationLogsConfigArgsDict']]
        """
        Application logs to file system configuration.
        """
elif False:
    ApplicationLogsConfigArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ApplicationLogsConfigArgs:
    def __init__(__self__, *,
                 azure_blob_storage: Optional[pulumi.Input['AzureBlobStorageApplicationLogsConfigArgs']] = None,
                 azure_table_storage: Optional[pulumi.Input['AzureTableStorageApplicationLogsConfigArgs']] = None,
                 file_system: Optional[pulumi.Input['FileSystemApplicationLogsConfigArgs']] = None):
        """
        Application logs configuration.
        :param pulumi.Input['AzureBlobStorageApplicationLogsConfigArgs'] azure_blob_storage: Application logs to blob storage configuration.
        :param pulumi.Input['AzureTableStorageApplicationLogsConfigArgs'] azure_table_storage: Application logs to azure table storage configuration.
        :param pulumi.Input['FileSystemApplicationLogsConfigArgs'] file_system: Application logs to file system configuration.
        """
        if azure_blob_storage is not None:
            pulumi.set(__self__, "azure_blob_storage", azure_blob_storage)
        if azure_table_storage is not None:
            pulumi.set(__self__, "azure_table_storage", azure_table_storage)
        if file_system is not None:
            pulumi.set(__self__, "file_system", file_system)

    @property
    @pulumi.getter(name="azureBlobStorage")
    def azure_blob_storage(self) -> Optional[pulumi.Input['AzureBlobStorageApplicationLogsConfigArgs']]:
        """
        Application logs to blob storage configuration.
        """
        return pulumi.get(self, "azure_blob_storage")

    @azure_blob_storage.setter
    def azure_blob_storage(self, value: Optional[pulumi.Input['AzureBlobStorageApplicationLogsConfigArgs']]):
        pulumi.set(self, "azure_blob_storage", value)

    @property
    @pulumi.getter(name="azureTableStorage")
    def azure_table_storage(self) -> Optional[pulumi.Input['AzureTableStorageApplicationLogsConfigArgs']]:
        """
        Application logs to azure table storage configuration.
        """
        return pulumi.get(self, "azure_table_storage")

    @azure_table_storage.setter
    def azure_table_storage(self, value: Optional[pulumi.Input['AzureTableStorageApplicationLogsConfigArgs']]):
        pulumi.set(self, "azure_table_storage", value)

    @property
    @pulumi.getter(name="fileSystem")
    def file_system(self) -> Optional[pulumi.Input['FileSystemApplicationLogsConfigArgs']]:
        """
        Application logs to file system configuration.
        """
        return pulumi.get(self, "file_system")

    @file_system.setter
    def file_system(self, value: Optional[pulumi.Input['FileSystemApplicationLogsConfigArgs']]):
        pulumi.set(self, "file_system", value)


if not MYPY:
    class AzureBlobStorageApplicationLogsConfigArgsDict(TypedDict):
        """
        Application logs azure blob storage configuration.
        """
        level: NotRequired[pulumi.Input['LogLevel']]
        """
        Log level.
        """
        retention_in_days: NotRequired[pulumi.Input[int]]
        """
        Retention in days.
        Remove blobs older than X days.
        0 or lower means no retention.
        """
        sas_url: NotRequired[pulumi.Input[str]]
        """
        SAS url to a azure blob container with read/write/list/delete permissions.
        """
elif False:
    AzureBlobStorageApplicationLogsConfigArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AzureBlobStorageApplicationLogsConfigArgs:
    def __init__(__self__, *,
                 level: Optional[pulumi.Input['LogLevel']] = None,
                 retention_in_days: Optional[pulumi.Input[int]] = None,
                 sas_url: Optional[pulumi.Input[str]] = None):
        """
        Application logs azure blob storage configuration.
        :param pulumi.Input['LogLevel'] level: Log level.
        :param pulumi.Input[int] retention_in_days: Retention in days.
               Remove blobs older than X days.
               0 or lower means no retention.
        :param pulumi.Input[str] sas_url: SAS url to a azure blob container with read/write/list/delete permissions.
        """
        if level is not None:
            pulumi.set(__self__, "level", level)
        if retention_in_days is not None:
            pulumi.set(__self__, "retention_in_days", retention_in_days)
        if sas_url is not None:
            pulumi.set(__self__, "sas_url", sas_url)

    @property
    @pulumi.getter
    def level(self) -> Optional[pulumi.Input['LogLevel']]:
        """
        Log level.
        """
        return pulumi.get(self, "level")

    @level.setter
    def level(self, value: Optional[pulumi.Input['LogLevel']]):
        pulumi.set(self, "level", value)

    @property
    @pulumi.getter(name="retentionInDays")
    def retention_in_days(self) -> Optional[pulumi.Input[int]]:
        """
        Retention in days.
        Remove blobs older than X days.
        0 or lower means no retention.
        """
        return pulumi.get(self, "retention_in_days")

    @retention_in_days.setter
    def retention_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_in_days", value)

    @property
    @pulumi.getter(name="sasUrl")
    def sas_url(self) -> Optional[pulumi.Input[str]]:
        """
        SAS url to a azure blob container with read/write/list/delete permissions.
        """
        return pulumi.get(self, "sas_url")

    @sas_url.setter
    def sas_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sas_url", value)


if not MYPY:
    class AzureBlobStorageHttpLogsConfigArgsDict(TypedDict):
        """
        Http logs to azure blob storage configuration.
        """
        enabled: NotRequired[pulumi.Input[bool]]
        """
        True if configuration is enabled, false if it is disabled and null if configuration is not set.
        """
        retention_in_days: NotRequired[pulumi.Input[int]]
        """
        Retention in days.
        Remove blobs older than X days.
        0 or lower means no retention.
        """
        sas_url: NotRequired[pulumi.Input[str]]
        """
        SAS url to a azure blob container with read/write/list/delete permissions.
        """
elif False:
    AzureBlobStorageHttpLogsConfigArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AzureBlobStorageHttpLogsConfigArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 retention_in_days: Optional[pulumi.Input[int]] = None,
                 sas_url: Optional[pulumi.Input[str]] = None):
        """
        Http logs to azure blob storage configuration.
        :param pulumi.Input[bool] enabled: True if configuration is enabled, false if it is disabled and null if configuration is not set.
        :param pulumi.Input[int] retention_in_days: Retention in days.
               Remove blobs older than X days.
               0 or lower means no retention.
        :param pulumi.Input[str] sas_url: SAS url to a azure blob container with read/write/list/delete permissions.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if retention_in_days is not None:
            pulumi.set(__self__, "retention_in_days", retention_in_days)
        if sas_url is not None:
            pulumi.set(__self__, "sas_url", sas_url)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        True if configuration is enabled, false if it is disabled and null if configuration is not set.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="retentionInDays")
    def retention_in_days(self) -> Optional[pulumi.Input[int]]:
        """
        Retention in days.
        Remove blobs older than X days.
        0 or lower means no retention.
        """
        return pulumi.get(self, "retention_in_days")

    @retention_in_days.setter
    def retention_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_in_days", value)

    @property
    @pulumi.getter(name="sasUrl")
    def sas_url(self) -> Optional[pulumi.Input[str]]:
        """
        SAS url to a azure blob container with read/write/list/delete permissions.
        """
        return pulumi.get(self, "sas_url")

    @sas_url.setter
    def sas_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sas_url", value)


if not MYPY:
    class AzureTableStorageApplicationLogsConfigArgsDict(TypedDict):
        """
        Application logs to Azure table storage configuration.
        """
        sas_url: pulumi.Input[str]
        """
        SAS URL to an Azure table with add/query/delete permissions.
        """
        level: NotRequired[pulumi.Input['LogLevel']]
        """
        Log level.
        """
elif False:
    AzureTableStorageApplicationLogsConfigArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AzureTableStorageApplicationLogsConfigArgs:
    def __init__(__self__, *,
                 sas_url: pulumi.Input[str],
                 level: Optional[pulumi.Input['LogLevel']] = None):
        """
        Application logs to Azure table storage configuration.
        :param pulumi.Input[str] sas_url: SAS URL to an Azure table with add/query/delete permissions.
        :param pulumi.Input['LogLevel'] level: Log level.
        """
        pulumi.set(__self__, "sas_url", sas_url)
        if level is not None:
            pulumi.set(__self__, "level", level)

    @property
    @pulumi.getter(name="sasUrl")
    def sas_url(self) -> pulumi.Input[str]:
        """
        SAS URL to an Azure table with add/query/delete permissions.
        """
        return pulumi.get(self, "sas_url")

    @sas_url.setter
    def sas_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "sas_url", value)

    @property
    @pulumi.getter
    def level(self) -> Optional[pulumi.Input['LogLevel']]:
        """
        Log level.
        """
        return pulumi.get(self, "level")

    @level.setter
    def level(self, value: Optional[pulumi.Input['LogLevel']]):
        pulumi.set(self, "level", value)


if not MYPY:
    class EnabledConfigArgsDict(TypedDict):
        """
        Enabled configuration.
        """
        enabled: NotRequired[pulumi.Input[bool]]
        """
        True if configuration is enabled, false if it is disabled and null if configuration is not set.
        """
elif False:
    EnabledConfigArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class EnabledConfigArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        Enabled configuration.
        :param pulumi.Input[bool] enabled: True if configuration is enabled, false if it is disabled and null if configuration is not set.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        True if configuration is enabled, false if it is disabled and null if configuration is not set.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


if not MYPY:
    class FileSystemApplicationLogsConfigArgsDict(TypedDict):
        """
        Application logs to file system configuration.
        """
        level: NotRequired[pulumi.Input['LogLevel']]
        """
        Log level.
        """
elif False:
    FileSystemApplicationLogsConfigArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class FileSystemApplicationLogsConfigArgs:
    def __init__(__self__, *,
                 level: Optional[pulumi.Input['LogLevel']] = None):
        """
        Application logs to file system configuration.
        :param pulumi.Input['LogLevel'] level: Log level.
        """
        if level is None:
            level = 'Off'
        if level is not None:
            pulumi.set(__self__, "level", level)

    @property
    @pulumi.getter
    def level(self) -> Optional[pulumi.Input['LogLevel']]:
        """
        Log level.
        """
        return pulumi.get(self, "level")

    @level.setter
    def level(self, value: Optional[pulumi.Input['LogLevel']]):
        pulumi.set(self, "level", value)


if not MYPY:
    class FileSystemHttpLogsConfigArgsDict(TypedDict):
        """
        Http logs to file system configuration.
        """
        enabled: NotRequired[pulumi.Input[bool]]
        """
        True if configuration is enabled, false if it is disabled and null if configuration is not set.
        """
        retention_in_days: NotRequired[pulumi.Input[int]]
        """
        Retention in days.
        Remove files older than X days.
        0 or lower means no retention.
        """
        retention_in_mb: NotRequired[pulumi.Input[int]]
        """
        Maximum size in megabytes that http log files can use.
        When reached old log files will be removed to make space for new ones.
        Value can range between 25 and 100.
        """
elif False:
    FileSystemHttpLogsConfigArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class FileSystemHttpLogsConfigArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 retention_in_days: Optional[pulumi.Input[int]] = None,
                 retention_in_mb: Optional[pulumi.Input[int]] = None):
        """
        Http logs to file system configuration.
        :param pulumi.Input[bool] enabled: True if configuration is enabled, false if it is disabled and null if configuration is not set.
        :param pulumi.Input[int] retention_in_days: Retention in days.
               Remove files older than X days.
               0 or lower means no retention.
        :param pulumi.Input[int] retention_in_mb: Maximum size in megabytes that http log files can use.
               When reached old log files will be removed to make space for new ones.
               Value can range between 25 and 100.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if retention_in_days is not None:
            pulumi.set(__self__, "retention_in_days", retention_in_days)
        if retention_in_mb is not None:
            pulumi.set(__self__, "retention_in_mb", retention_in_mb)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        True if configuration is enabled, false if it is disabled and null if configuration is not set.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="retentionInDays")
    def retention_in_days(self) -> Optional[pulumi.Input[int]]:
        """
        Retention in days.
        Remove files older than X days.
        0 or lower means no retention.
        """
        return pulumi.get(self, "retention_in_days")

    @retention_in_days.setter
    def retention_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_in_days", value)

    @property
    @pulumi.getter(name="retentionInMb")
    def retention_in_mb(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum size in megabytes that http log files can use.
        When reached old log files will be removed to make space for new ones.
        Value can range between 25 and 100.
        """
        return pulumi.get(self, "retention_in_mb")

    @retention_in_mb.setter
    def retention_in_mb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_in_mb", value)


if not MYPY:
    class HttpLogsConfigArgsDict(TypedDict):
        """
        Http logs configuration.
        """
        azure_blob_storage: NotRequired[pulumi.Input['AzureBlobStorageHttpLogsConfigArgsDict']]
        """
        Http logs to azure blob storage configuration.
        """
        file_system: NotRequired[pulumi.Input['FileSystemHttpLogsConfigArgsDict']]
        """
        Http logs to file system configuration.
        """
elif False:
    HttpLogsConfigArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class HttpLogsConfigArgs:
    def __init__(__self__, *,
                 azure_blob_storage: Optional[pulumi.Input['AzureBlobStorageHttpLogsConfigArgs']] = None,
                 file_system: Optional[pulumi.Input['FileSystemHttpLogsConfigArgs']] = None):
        """
        Http logs configuration.
        :param pulumi.Input['AzureBlobStorageHttpLogsConfigArgs'] azure_blob_storage: Http logs to azure blob storage configuration.
        :param pulumi.Input['FileSystemHttpLogsConfigArgs'] file_system: Http logs to file system configuration.
        """
        if azure_blob_storage is not None:
            pulumi.set(__self__, "azure_blob_storage", azure_blob_storage)
        if file_system is not None:
            pulumi.set(__self__, "file_system", file_system)

    @property
    @pulumi.getter(name="azureBlobStorage")
    def azure_blob_storage(self) -> Optional[pulumi.Input['AzureBlobStorageHttpLogsConfigArgs']]:
        """
        Http logs to azure blob storage configuration.
        """
        return pulumi.get(self, "azure_blob_storage")

    @azure_blob_storage.setter
    def azure_blob_storage(self, value: Optional[pulumi.Input['AzureBlobStorageHttpLogsConfigArgs']]):
        pulumi.set(self, "azure_blob_storage", value)

    @property
    @pulumi.getter(name="fileSystem")
    def file_system(self) -> Optional[pulumi.Input['FileSystemHttpLogsConfigArgs']]:
        """
        Http logs to file system configuration.
        """
        return pulumi.get(self, "file_system")

    @file_system.setter
    def file_system(self, value: Optional[pulumi.Input['FileSystemHttpLogsConfigArgs']]):
        pulumi.set(self, "file_system", value)


if not MYPY:
    class NameValuePairArgsDict(TypedDict):
        """
        Name value pair.
        """
        name: NotRequired[pulumi.Input[str]]
        """
        Pair name.
        """
        value: NotRequired[pulumi.Input[str]]
        """
        Pair value.
        """
elif False:
    NameValuePairArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class NameValuePairArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Name value pair.
        :param pulumi.Input[str] name: Pair name.
        :param pulumi.Input[str] value: Pair value.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Pair name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        Pair value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


if not MYPY:
    class VirtualNetworkProfileArgsDict(TypedDict):
        """
        Specification for using a Virtual Network.
        """
        id: pulumi.Input[str]
        """
        Resource id of the Virtual Network.
        """
        subnet: NotRequired[pulumi.Input[str]]
        """
        Subnet within the Virtual Network.
        """
elif False:
    VirtualNetworkProfileArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class VirtualNetworkProfileArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 subnet: Optional[pulumi.Input[str]] = None):
        """
        Specification for using a Virtual Network.
        :param pulumi.Input[str] id: Resource id of the Virtual Network.
        :param pulumi.Input[str] subnet: Subnet within the Virtual Network.
        """
        pulumi.set(__self__, "id", id)
        if subnet is not None:
            pulumi.set(__self__, "subnet", subnet)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        Resource id of the Virtual Network.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def subnet(self) -> Optional[pulumi.Input[str]]:
        """
        Subnet within the Virtual Network.
        """
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet", value)


