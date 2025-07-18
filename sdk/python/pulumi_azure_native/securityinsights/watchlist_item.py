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
from ._inputs import *

__all__ = ['WatchlistItemArgs', 'WatchlistItem']

@pulumi.input_type
class WatchlistItemArgs:
    def __init__(__self__, *,
                 items_key_value: Any,
                 resource_group_name: pulumi.Input[builtins.str],
                 watchlist_alias: pulumi.Input[builtins.str],
                 workspace_name: pulumi.Input[builtins.str],
                 created: Optional[pulumi.Input[builtins.str]] = None,
                 created_by: Optional[pulumi.Input['WatchlistUserInfoArgs']] = None,
                 entity_mapping: Optional[Any] = None,
                 is_deleted: Optional[pulumi.Input[builtins.bool]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 updated: Optional[pulumi.Input[builtins.str]] = None,
                 updated_by: Optional[pulumi.Input['WatchlistUserInfoArgs']] = None,
                 watchlist_item_id: Optional[pulumi.Input[builtins.str]] = None,
                 watchlist_item_type: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a WatchlistItem resource.
        :param Any items_key_value: key-value pairs for a watchlist item
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] watchlist_alias: The watchlist alias
        :param pulumi.Input[builtins.str] workspace_name: The name of the workspace.
        :param pulumi.Input[builtins.str] created: The time the watchlist item was created
        :param pulumi.Input['WatchlistUserInfoArgs'] created_by: Describes a user that created the watchlist item
        :param Any entity_mapping: key-value pairs for a watchlist item entity mapping
        :param pulumi.Input[builtins.bool] is_deleted: A flag that indicates if the watchlist item is deleted or not
        :param pulumi.Input[builtins.str] tenant_id: The tenantId to which the watchlist item belongs to
        :param pulumi.Input[builtins.str] updated: The last time the watchlist item was updated
        :param pulumi.Input['WatchlistUserInfoArgs'] updated_by: Describes a user that updated the watchlist item
        :param pulumi.Input[builtins.str] watchlist_item_id: The id (a Guid) of the watchlist item
        :param pulumi.Input[builtins.str] watchlist_item_type: The type of the watchlist item
        """
        pulumi.set(__self__, "items_key_value", items_key_value)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "watchlist_alias", watchlist_alias)
        pulumi.set(__self__, "workspace_name", workspace_name)
        if created is not None:
            pulumi.set(__self__, "created", created)
        if created_by is not None:
            pulumi.set(__self__, "created_by", created_by)
        if entity_mapping is not None:
            pulumi.set(__self__, "entity_mapping", entity_mapping)
        if is_deleted is not None:
            pulumi.set(__self__, "is_deleted", is_deleted)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if updated is not None:
            pulumi.set(__self__, "updated", updated)
        if updated_by is not None:
            pulumi.set(__self__, "updated_by", updated_by)
        if watchlist_item_id is not None:
            pulumi.set(__self__, "watchlist_item_id", watchlist_item_id)
        if watchlist_item_type is not None:
            pulumi.set(__self__, "watchlist_item_type", watchlist_item_type)

    @property
    @pulumi.getter(name="itemsKeyValue")
    def items_key_value(self) -> Any:
        """
        key-value pairs for a watchlist item
        """
        return pulumi.get(self, "items_key_value")

    @items_key_value.setter
    def items_key_value(self, value: Any):
        pulumi.set(self, "items_key_value", value)

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
    @pulumi.getter(name="watchlistAlias")
    def watchlist_alias(self) -> pulumi.Input[builtins.str]:
        """
        The watchlist alias
        """
        return pulumi.get(self, "watchlist_alias")

    @watchlist_alias.setter
    def watchlist_alias(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "watchlist_alias", value)

    @property
    @pulumi.getter(name="workspaceName")
    def workspace_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the workspace.
        """
        return pulumi.get(self, "workspace_name")

    @workspace_name.setter
    def workspace_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "workspace_name", value)

    @property
    @pulumi.getter
    def created(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The time the watchlist item was created
        """
        return pulumi.get(self, "created")

    @created.setter
    def created(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created", value)

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> Optional[pulumi.Input['WatchlistUserInfoArgs']]:
        """
        Describes a user that created the watchlist item
        """
        return pulumi.get(self, "created_by")

    @created_by.setter
    def created_by(self, value: Optional[pulumi.Input['WatchlistUserInfoArgs']]):
        pulumi.set(self, "created_by", value)

    @property
    @pulumi.getter(name="entityMapping")
    def entity_mapping(self) -> Optional[Any]:
        """
        key-value pairs for a watchlist item entity mapping
        """
        return pulumi.get(self, "entity_mapping")

    @entity_mapping.setter
    def entity_mapping(self, value: Optional[Any]):
        pulumi.set(self, "entity_mapping", value)

    @property
    @pulumi.getter(name="isDeleted")
    def is_deleted(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        A flag that indicates if the watchlist item is deleted or not
        """
        return pulumi.get(self, "is_deleted")

    @is_deleted.setter
    def is_deleted(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "is_deleted", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The tenantId to which the watchlist item belongs to
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def updated(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The last time the watchlist item was updated
        """
        return pulumi.get(self, "updated")

    @updated.setter
    def updated(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "updated", value)

    @property
    @pulumi.getter(name="updatedBy")
    def updated_by(self) -> Optional[pulumi.Input['WatchlistUserInfoArgs']]:
        """
        Describes a user that updated the watchlist item
        """
        return pulumi.get(self, "updated_by")

    @updated_by.setter
    def updated_by(self, value: Optional[pulumi.Input['WatchlistUserInfoArgs']]):
        pulumi.set(self, "updated_by", value)

    @property
    @pulumi.getter(name="watchlistItemId")
    def watchlist_item_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The id (a Guid) of the watchlist item
        """
        return pulumi.get(self, "watchlist_item_id")

    @watchlist_item_id.setter
    def watchlist_item_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "watchlist_item_id", value)

    @property
    @pulumi.getter(name="watchlistItemType")
    def watchlist_item_type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The type of the watchlist item
        """
        return pulumi.get(self, "watchlist_item_type")

    @watchlist_item_type.setter
    def watchlist_item_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "watchlist_item_type", value)


@pulumi.type_token("azure-native:securityinsights:WatchlistItem")
class WatchlistItem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 created: Optional[pulumi.Input[builtins.str]] = None,
                 created_by: Optional[pulumi.Input[Union['WatchlistUserInfoArgs', 'WatchlistUserInfoArgsDict']]] = None,
                 entity_mapping: Optional[Any] = None,
                 is_deleted: Optional[pulumi.Input[builtins.bool]] = None,
                 items_key_value: Optional[Any] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 updated: Optional[pulumi.Input[builtins.str]] = None,
                 updated_by: Optional[pulumi.Input[Union['WatchlistUserInfoArgs', 'WatchlistUserInfoArgsDict']]] = None,
                 watchlist_alias: Optional[pulumi.Input[builtins.str]] = None,
                 watchlist_item_id: Optional[pulumi.Input[builtins.str]] = None,
                 watchlist_item_type: Optional[pulumi.Input[builtins.str]] = None,
                 workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Represents a Watchlist Item in Azure Security Insights.

        Uses Azure REST API version 2024-09-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.

        Other available API versions: 2023-02-01, 2023-03-01-preview, 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-11-01, 2023-12-01-preview, 2024-01-01-preview, 2024-03-01, 2024-04-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-03-01, 2025-04-01-preview, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] created: The time the watchlist item was created
        :param pulumi.Input[Union['WatchlistUserInfoArgs', 'WatchlistUserInfoArgsDict']] created_by: Describes a user that created the watchlist item
        :param Any entity_mapping: key-value pairs for a watchlist item entity mapping
        :param pulumi.Input[builtins.bool] is_deleted: A flag that indicates if the watchlist item is deleted or not
        :param Any items_key_value: key-value pairs for a watchlist item
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] tenant_id: The tenantId to which the watchlist item belongs to
        :param pulumi.Input[builtins.str] updated: The last time the watchlist item was updated
        :param pulumi.Input[Union['WatchlistUserInfoArgs', 'WatchlistUserInfoArgsDict']] updated_by: Describes a user that updated the watchlist item
        :param pulumi.Input[builtins.str] watchlist_alias: The watchlist alias
        :param pulumi.Input[builtins.str] watchlist_item_id: The id (a Guid) of the watchlist item
        :param pulumi.Input[builtins.str] watchlist_item_type: The type of the watchlist item
        :param pulumi.Input[builtins.str] workspace_name: The name of the workspace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WatchlistItemArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a Watchlist Item in Azure Security Insights.

        Uses Azure REST API version 2024-09-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.

        Other available API versions: 2023-02-01, 2023-03-01-preview, 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-11-01, 2023-12-01-preview, 2024-01-01-preview, 2024-03-01, 2024-04-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-03-01, 2025-04-01-preview, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param WatchlistItemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WatchlistItemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 created: Optional[pulumi.Input[builtins.str]] = None,
                 created_by: Optional[pulumi.Input[Union['WatchlistUserInfoArgs', 'WatchlistUserInfoArgsDict']]] = None,
                 entity_mapping: Optional[Any] = None,
                 is_deleted: Optional[pulumi.Input[builtins.bool]] = None,
                 items_key_value: Optional[Any] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 updated: Optional[pulumi.Input[builtins.str]] = None,
                 updated_by: Optional[pulumi.Input[Union['WatchlistUserInfoArgs', 'WatchlistUserInfoArgsDict']]] = None,
                 watchlist_alias: Optional[pulumi.Input[builtins.str]] = None,
                 watchlist_item_id: Optional[pulumi.Input[builtins.str]] = None,
                 watchlist_item_type: Optional[pulumi.Input[builtins.str]] = None,
                 workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WatchlistItemArgs.__new__(WatchlistItemArgs)

            __props__.__dict__["created"] = created
            __props__.__dict__["created_by"] = created_by
            __props__.__dict__["entity_mapping"] = entity_mapping
            __props__.__dict__["is_deleted"] = is_deleted
            if items_key_value is None and not opts.urn:
                raise TypeError("Missing required property 'items_key_value'")
            __props__.__dict__["items_key_value"] = items_key_value
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["updated"] = updated
            __props__.__dict__["updated_by"] = updated_by
            if watchlist_alias is None and not opts.urn:
                raise TypeError("Missing required property 'watchlist_alias'")
            __props__.__dict__["watchlist_alias"] = watchlist_alias
            __props__.__dict__["watchlist_item_id"] = watchlist_item_id
            __props__.__dict__["watchlist_item_type"] = watchlist_item_type
            if workspace_name is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_name'")
            __props__.__dict__["workspace_name"] = workspace_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:securityinsights/v20190101preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20210301preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20210401:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20210901preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20211001:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20211001preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20220101preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20220401preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20220501preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20220601preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20220701preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20220801:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20220801preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20220901preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20221001preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20221101:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20221101preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20221201preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20230201:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20230201preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20230301preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20230401preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20230501preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20230601preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20230701preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20230801preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20230901preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20231001preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20231101:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20231201preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20240101preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20240301:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20240401preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20240901:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20241001preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20250101preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20250301:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20250401preview:WatchlistItem"), pulumi.Alias(type_="azure-native:securityinsights/v20250601:WatchlistItem")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(WatchlistItem, __self__).__init__(
            'azure-native:securityinsights:WatchlistItem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WatchlistItem':
        """
        Get an existing WatchlistItem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WatchlistItemArgs.__new__(WatchlistItemArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["created"] = None
        __props__.__dict__["created_by"] = None
        __props__.__dict__["entity_mapping"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["is_deleted"] = None
        __props__.__dict__["items_key_value"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tenant_id"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["updated"] = None
        __props__.__dict__["updated_by"] = None
        __props__.__dict__["watchlist_item_id"] = None
        __props__.__dict__["watchlist_item_type"] = None
        return WatchlistItem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The time the watchlist item was created
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> pulumi.Output[Optional['outputs.WatchlistUserInfoResponse']]:
        """
        Describes a user that created the watchlist item
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="entityMapping")
    def entity_mapping(self) -> pulumi.Output[Optional[Any]]:
        """
        key-value pairs for a watchlist item entity mapping
        """
        return pulumi.get(self, "entity_mapping")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Etag of the azure resource
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="isDeleted")
    def is_deleted(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        A flag that indicates if the watchlist item is deleted or not
        """
        return pulumi.get(self, "is_deleted")

    @property
    @pulumi.getter(name="itemsKeyValue")
    def items_key_value(self) -> pulumi.Output[Any]:
        """
        key-value pairs for a watchlist item
        """
        return pulumi.get(self, "items_key_value")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The tenantId to which the watchlist item belongs to
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def updated(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The last time the watchlist item was updated
        """
        return pulumi.get(self, "updated")

    @property
    @pulumi.getter(name="updatedBy")
    def updated_by(self) -> pulumi.Output[Optional['outputs.WatchlistUserInfoResponse']]:
        """
        Describes a user that updated the watchlist item
        """
        return pulumi.get(self, "updated_by")

    @property
    @pulumi.getter(name="watchlistItemId")
    def watchlist_item_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The id (a Guid) of the watchlist item
        """
        return pulumi.get(self, "watchlist_item_id")

    @property
    @pulumi.getter(name="watchlistItemType")
    def watchlist_item_type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The type of the watchlist item
        """
        return pulumi.get(self, "watchlist_item_type")

