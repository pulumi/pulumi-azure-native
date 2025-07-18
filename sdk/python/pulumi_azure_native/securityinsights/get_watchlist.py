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
    'GetWatchlistResult',
    'AwaitableGetWatchlistResult',
    'get_watchlist',
    'get_watchlist_output',
]

@pulumi.output_type
class GetWatchlistResult:
    """
    Represents a Watchlist in Azure Security Insights.
    """
    def __init__(__self__, azure_api_version=None, content_type=None, created=None, created_by=None, default_duration=None, description=None, display_name=None, etag=None, id=None, is_deleted=None, items_search_key=None, labels=None, name=None, number_of_lines_to_skip=None, provider=None, provisioning_state=None, raw_content=None, source=None, source_type=None, system_data=None, tenant_id=None, type=None, updated=None, updated_by=None, upload_status=None, watchlist_alias=None, watchlist_id=None, watchlist_type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if content_type and not isinstance(content_type, str):
            raise TypeError("Expected argument 'content_type' to be a str")
        pulumi.set(__self__, "content_type", content_type)
        if created and not isinstance(created, str):
            raise TypeError("Expected argument 'created' to be a str")
        pulumi.set(__self__, "created", created)
        if created_by and not isinstance(created_by, dict):
            raise TypeError("Expected argument 'created_by' to be a dict")
        pulumi.set(__self__, "created_by", created_by)
        if default_duration and not isinstance(default_duration, str):
            raise TypeError("Expected argument 'default_duration' to be a str")
        pulumi.set(__self__, "default_duration", default_duration)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_deleted and not isinstance(is_deleted, bool):
            raise TypeError("Expected argument 'is_deleted' to be a bool")
        pulumi.set(__self__, "is_deleted", is_deleted)
        if items_search_key and not isinstance(items_search_key, str):
            raise TypeError("Expected argument 'items_search_key' to be a str")
        pulumi.set(__self__, "items_search_key", items_search_key)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if number_of_lines_to_skip and not isinstance(number_of_lines_to_skip, int):
            raise TypeError("Expected argument 'number_of_lines_to_skip' to be a int")
        pulumi.set(__self__, "number_of_lines_to_skip", number_of_lines_to_skip)
        if provider and not isinstance(provider, str):
            raise TypeError("Expected argument 'provider' to be a str")
        pulumi.set(__self__, "provider", provider)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if raw_content and not isinstance(raw_content, str):
            raise TypeError("Expected argument 'raw_content' to be a str")
        pulumi.set(__self__, "raw_content", raw_content)
        if source and not isinstance(source, str):
            raise TypeError("Expected argument 'source' to be a str")
        pulumi.set(__self__, "source", source)
        if source_type and not isinstance(source_type, str):
            raise TypeError("Expected argument 'source_type' to be a str")
        pulumi.set(__self__, "source_type", source_type)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if updated and not isinstance(updated, str):
            raise TypeError("Expected argument 'updated' to be a str")
        pulumi.set(__self__, "updated", updated)
        if updated_by and not isinstance(updated_by, dict):
            raise TypeError("Expected argument 'updated_by' to be a dict")
        pulumi.set(__self__, "updated_by", updated_by)
        if upload_status and not isinstance(upload_status, str):
            raise TypeError("Expected argument 'upload_status' to be a str")
        pulumi.set(__self__, "upload_status", upload_status)
        if watchlist_alias and not isinstance(watchlist_alias, str):
            raise TypeError("Expected argument 'watchlist_alias' to be a str")
        pulumi.set(__self__, "watchlist_alias", watchlist_alias)
        if watchlist_id and not isinstance(watchlist_id, str):
            raise TypeError("Expected argument 'watchlist_id' to be a str")
        pulumi.set(__self__, "watchlist_id", watchlist_id)
        if watchlist_type and not isinstance(watchlist_type, str):
            raise TypeError("Expected argument 'watchlist_type' to be a str")
        pulumi.set(__self__, "watchlist_type", watchlist_type)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[builtins.str]:
        """
        The content type of the raw content. Example : text/csv or text/tsv
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter
    def created(self) -> Optional[builtins.str]:
        """
        The time the watchlist was created
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> Optional['outputs.WatchlistUserInfoResponse']:
        """
        Describes a user that created the watchlist
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="defaultDuration")
    def default_duration(self) -> Optional[builtins.str]:
        """
        The default duration of a watchlist (in ISO 8601 duration format)
        """
        return pulumi.get(self, "default_duration")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        A description of the watchlist
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> builtins.str:
        """
        The display name of the watchlist
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def etag(self) -> Optional[builtins.str]:
        """
        Etag of the azure resource
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isDeleted")
    def is_deleted(self) -> Optional[builtins.bool]:
        """
        A flag that indicates if the watchlist is deleted or not
        """
        return pulumi.get(self, "is_deleted")

    @property
    @pulumi.getter(name="itemsSearchKey")
    def items_search_key(self) -> builtins.str:
        """
        The search key is used to optimize query performance when using watchlists for joins with other data. For example, enable a column with IP addresses to be the designated SearchKey field, then use this field as the key field when joining to other event data by IP address.
        """
        return pulumi.get(self, "items_search_key")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Sequence[builtins.str]]:
        """
        List of labels relevant to this watchlist
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numberOfLinesToSkip")
    def number_of_lines_to_skip(self) -> Optional[builtins.int]:
        """
        The number of lines in a csv/tsv content to skip before the header
        """
        return pulumi.get(self, "number_of_lines_to_skip")

    @property
    @pulumi.getter
    def provider(self) -> builtins.str:
        """
        The provider of the watchlist
        """
        return pulumi.get(self, "provider")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Describes provisioning state
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="rawContent")
    def raw_content(self) -> Optional[builtins.str]:
        """
        The raw content that represents to watchlist items to create. In case of csv/tsv content type, it's the content of the file that will parsed by the endpoint
        """
        return pulumi.get(self, "raw_content")

    @property
    @pulumi.getter
    def source(self) -> Optional[builtins.str]:
        """
        The filename of the watchlist, called 'source'
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> Optional[builtins.str]:
        """
        The sourceType of the watchlist
        """
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[builtins.str]:
        """
        The tenantId where the watchlist belongs to
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def updated(self) -> Optional[builtins.str]:
        """
        The last time the watchlist was updated
        """
        return pulumi.get(self, "updated")

    @property
    @pulumi.getter(name="updatedBy")
    def updated_by(self) -> Optional['outputs.WatchlistUserInfoResponse']:
        """
        Describes a user that updated the watchlist
        """
        return pulumi.get(self, "updated_by")

    @property
    @pulumi.getter(name="uploadStatus")
    def upload_status(self) -> Optional[builtins.str]:
        """
        The status of the Watchlist upload : New, InProgress or Complete. **Note** : When a Watchlist upload status is InProgress, the Watchlist cannot be deleted
        """
        return pulumi.get(self, "upload_status")

    @property
    @pulumi.getter(name="watchlistAlias")
    def watchlist_alias(self) -> Optional[builtins.str]:
        """
        The alias of the watchlist
        """
        return pulumi.get(self, "watchlist_alias")

    @property
    @pulumi.getter(name="watchlistId")
    def watchlist_id(self) -> Optional[builtins.str]:
        """
        The id (a Guid) of the watchlist
        """
        return pulumi.get(self, "watchlist_id")

    @property
    @pulumi.getter(name="watchlistType")
    def watchlist_type(self) -> Optional[builtins.str]:
        """
        The type of the watchlist
        """
        return pulumi.get(self, "watchlist_type")


class AwaitableGetWatchlistResult(GetWatchlistResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWatchlistResult(
            azure_api_version=self.azure_api_version,
            content_type=self.content_type,
            created=self.created,
            created_by=self.created_by,
            default_duration=self.default_duration,
            description=self.description,
            display_name=self.display_name,
            etag=self.etag,
            id=self.id,
            is_deleted=self.is_deleted,
            items_search_key=self.items_search_key,
            labels=self.labels,
            name=self.name,
            number_of_lines_to_skip=self.number_of_lines_to_skip,
            provider=self.provider,
            provisioning_state=self.provisioning_state,
            raw_content=self.raw_content,
            source=self.source,
            source_type=self.source_type,
            system_data=self.system_data,
            tenant_id=self.tenant_id,
            type=self.type,
            updated=self.updated,
            updated_by=self.updated_by,
            upload_status=self.upload_status,
            watchlist_alias=self.watchlist_alias,
            watchlist_id=self.watchlist_id,
            watchlist_type=self.watchlist_type)


def get_watchlist(resource_group_name: Optional[builtins.str] = None,
                  watchlist_alias: Optional[builtins.str] = None,
                  workspace_name: Optional[builtins.str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWatchlistResult:
    """
    Get a watchlist, without its watchlist items.

    Uses Azure REST API version 2024-09-01.

    Other available API versions: 2023-02-01, 2023-03-01-preview, 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-11-01, 2023-12-01-preview, 2024-01-01-preview, 2024-03-01, 2024-04-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-03-01, 2025-04-01-preview, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str watchlist_alias: The watchlist alias
    :param builtins.str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['watchlistAlias'] = watchlist_alias
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:securityinsights:getWatchlist', __args__, opts=opts, typ=GetWatchlistResult).value

    return AwaitableGetWatchlistResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        content_type=pulumi.get(__ret__, 'content_type'),
        created=pulumi.get(__ret__, 'created'),
        created_by=pulumi.get(__ret__, 'created_by'),
        default_duration=pulumi.get(__ret__, 'default_duration'),
        description=pulumi.get(__ret__, 'description'),
        display_name=pulumi.get(__ret__, 'display_name'),
        etag=pulumi.get(__ret__, 'etag'),
        id=pulumi.get(__ret__, 'id'),
        is_deleted=pulumi.get(__ret__, 'is_deleted'),
        items_search_key=pulumi.get(__ret__, 'items_search_key'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        number_of_lines_to_skip=pulumi.get(__ret__, 'number_of_lines_to_skip'),
        provider=pulumi.get(__ret__, 'provider'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        raw_content=pulumi.get(__ret__, 'raw_content'),
        source=pulumi.get(__ret__, 'source'),
        source_type=pulumi.get(__ret__, 'source_type'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'),
        type=pulumi.get(__ret__, 'type'),
        updated=pulumi.get(__ret__, 'updated'),
        updated_by=pulumi.get(__ret__, 'updated_by'),
        upload_status=pulumi.get(__ret__, 'upload_status'),
        watchlist_alias=pulumi.get(__ret__, 'watchlist_alias'),
        watchlist_id=pulumi.get(__ret__, 'watchlist_id'),
        watchlist_type=pulumi.get(__ret__, 'watchlist_type'))
def get_watchlist_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                         watchlist_alias: Optional[pulumi.Input[builtins.str]] = None,
                         workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetWatchlistResult]:
    """
    Get a watchlist, without its watchlist items.

    Uses Azure REST API version 2024-09-01.

    Other available API versions: 2023-02-01, 2023-03-01-preview, 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-11-01, 2023-12-01-preview, 2024-01-01-preview, 2024-03-01, 2024-04-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-03-01, 2025-04-01-preview, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str watchlist_alias: The watchlist alias
    :param builtins.str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['watchlistAlias'] = watchlist_alias
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:securityinsights:getWatchlist', __args__, opts=opts, typ=GetWatchlistResult)
    return __ret__.apply(lambda __response__: GetWatchlistResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        content_type=pulumi.get(__response__, 'content_type'),
        created=pulumi.get(__response__, 'created'),
        created_by=pulumi.get(__response__, 'created_by'),
        default_duration=pulumi.get(__response__, 'default_duration'),
        description=pulumi.get(__response__, 'description'),
        display_name=pulumi.get(__response__, 'display_name'),
        etag=pulumi.get(__response__, 'etag'),
        id=pulumi.get(__response__, 'id'),
        is_deleted=pulumi.get(__response__, 'is_deleted'),
        items_search_key=pulumi.get(__response__, 'items_search_key'),
        labels=pulumi.get(__response__, 'labels'),
        name=pulumi.get(__response__, 'name'),
        number_of_lines_to_skip=pulumi.get(__response__, 'number_of_lines_to_skip'),
        provider=pulumi.get(__response__, 'provider'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        raw_content=pulumi.get(__response__, 'raw_content'),
        source=pulumi.get(__response__, 'source'),
        source_type=pulumi.get(__response__, 'source_type'),
        system_data=pulumi.get(__response__, 'system_data'),
        tenant_id=pulumi.get(__response__, 'tenant_id'),
        type=pulumi.get(__response__, 'type'),
        updated=pulumi.get(__response__, 'updated'),
        updated_by=pulumi.get(__response__, 'updated_by'),
        upload_status=pulumi.get(__response__, 'upload_status'),
        watchlist_alias=pulumi.get(__response__, 'watchlist_alias'),
        watchlist_id=pulumi.get(__response__, 'watchlist_id'),
        watchlist_type=pulumi.get(__response__, 'watchlist_type')))
