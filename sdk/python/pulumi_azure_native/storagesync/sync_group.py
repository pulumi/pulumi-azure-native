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

__all__ = ['SyncGroupArgs', 'SyncGroup']

@pulumi.input_type
class SyncGroupArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 storage_sync_service_name: pulumi.Input[builtins.str],
                 sync_group_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a SyncGroup resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] storage_sync_service_name: Name of Storage Sync Service resource.
        :param pulumi.Input[builtins.str] sync_group_name: Name of Sync Group resource.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "storage_sync_service_name", storage_sync_service_name)
        if sync_group_name is not None:
            pulumi.set(__self__, "sync_group_name", sync_group_name)

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
    @pulumi.getter(name="storageSyncServiceName")
    def storage_sync_service_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of Storage Sync Service resource.
        """
        return pulumi.get(self, "storage_sync_service_name")

    @storage_sync_service_name.setter
    def storage_sync_service_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "storage_sync_service_name", value)

    @property
    @pulumi.getter(name="syncGroupName")
    def sync_group_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of Sync Group resource.
        """
        return pulumi.get(self, "sync_group_name")

    @sync_group_name.setter
    def sync_group_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "sync_group_name", value)


@pulumi.type_token("azure-native:storagesync:SyncGroup")
class SyncGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 storage_sync_service_name: Optional[pulumi.Input[builtins.str]] = None,
                 sync_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Sync Group object.

        Uses Azure REST API version 2022-09-01. In version 2.x of the Azure Native provider, it used API version 2022-06-01.

        Other available API versions: 2022-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storagesync [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] storage_sync_service_name: Name of Storage Sync Service resource.
        :param pulumi.Input[builtins.str] sync_group_name: Name of Sync Group resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SyncGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Sync Group object.

        Uses Azure REST API version 2022-09-01. In version 2.x of the Azure Native provider, it used API version 2022-06-01.

        Other available API versions: 2022-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storagesync [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param SyncGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SyncGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 storage_sync_service_name: Optional[pulumi.Input[builtins.str]] = None,
                 sync_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SyncGroupArgs.__new__(SyncGroupArgs)

            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if storage_sync_service_name is None and not opts.urn:
                raise TypeError("Missing required property 'storage_sync_service_name'")
            __props__.__dict__["storage_sync_service_name"] = storage_sync_service_name
            __props__.__dict__["sync_group_name"] = sync_group_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["sync_group_status"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["unique_id"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:storagesync/v20170605preview:SyncGroup"), pulumi.Alias(type_="azure-native:storagesync/v20180402:SyncGroup"), pulumi.Alias(type_="azure-native:storagesync/v20180701:SyncGroup"), pulumi.Alias(type_="azure-native:storagesync/v20181001:SyncGroup"), pulumi.Alias(type_="azure-native:storagesync/v20190201:SyncGroup"), pulumi.Alias(type_="azure-native:storagesync/v20190301:SyncGroup"), pulumi.Alias(type_="azure-native:storagesync/v20190601:SyncGroup"), pulumi.Alias(type_="azure-native:storagesync/v20191001:SyncGroup"), pulumi.Alias(type_="azure-native:storagesync/v20200301:SyncGroup"), pulumi.Alias(type_="azure-native:storagesync/v20200901:SyncGroup"), pulumi.Alias(type_="azure-native:storagesync/v20220601:SyncGroup"), pulumi.Alias(type_="azure-native:storagesync/v20220901:SyncGroup")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(SyncGroup, __self__).__init__(
            'azure-native:storagesync:SyncGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SyncGroup':
        """
        Get an existing SyncGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SyncGroupArgs.__new__(SyncGroupArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["sync_group_status"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["unique_id"] = None
        return SyncGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="syncGroupStatus")
    def sync_group_status(self) -> pulumi.Output[builtins.str]:
        """
        Sync group status
        """
        return pulumi.get(self, "sync_group_status")

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
    @pulumi.getter(name="uniqueId")
    def unique_id(self) -> pulumi.Output[builtins.str]:
        """
        Unique Id
        """
        return pulumi.get(self, "unique_id")

