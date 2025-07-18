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
    'GetApplicationGroupResult',
    'AwaitableGetApplicationGroupResult',
    'get_application_group',
    'get_application_group_output',
]

@pulumi.output_type
class GetApplicationGroupResult:
    """
    Represents a ApplicationGroup definition.
    """
    def __init__(__self__, application_group_type=None, azure_api_version=None, cloud_pc_resource=None, description=None, etag=None, friendly_name=None, host_pool_arm_path=None, id=None, identity=None, kind=None, location=None, managed_by=None, name=None, object_id=None, plan=None, show_in_feed=None, sku=None, system_data=None, tags=None, type=None, workspace_arm_path=None):
        if application_group_type and not isinstance(application_group_type, str):
            raise TypeError("Expected argument 'application_group_type' to be a str")
        pulumi.set(__self__, "application_group_type", application_group_type)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if cloud_pc_resource and not isinstance(cloud_pc_resource, bool):
            raise TypeError("Expected argument 'cloud_pc_resource' to be a bool")
        pulumi.set(__self__, "cloud_pc_resource", cloud_pc_resource)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if friendly_name and not isinstance(friendly_name, str):
            raise TypeError("Expected argument 'friendly_name' to be a str")
        pulumi.set(__self__, "friendly_name", friendly_name)
        if host_pool_arm_path and not isinstance(host_pool_arm_path, str):
            raise TypeError("Expected argument 'host_pool_arm_path' to be a str")
        pulumi.set(__self__, "host_pool_arm_path", host_pool_arm_path)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if managed_by and not isinstance(managed_by, str):
            raise TypeError("Expected argument 'managed_by' to be a str")
        pulumi.set(__self__, "managed_by", managed_by)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if object_id and not isinstance(object_id, str):
            raise TypeError("Expected argument 'object_id' to be a str")
        pulumi.set(__self__, "object_id", object_id)
        if plan and not isinstance(plan, dict):
            raise TypeError("Expected argument 'plan' to be a dict")
        pulumi.set(__self__, "plan", plan)
        if show_in_feed and not isinstance(show_in_feed, bool):
            raise TypeError("Expected argument 'show_in_feed' to be a bool")
        pulumi.set(__self__, "show_in_feed", show_in_feed)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if workspace_arm_path and not isinstance(workspace_arm_path, str):
            raise TypeError("Expected argument 'workspace_arm_path' to be a str")
        pulumi.set(__self__, "workspace_arm_path", workspace_arm_path)

    @property
    @pulumi.getter(name="applicationGroupType")
    def application_group_type(self) -> builtins.str:
        """
        Resource Type of ApplicationGroup.
        """
        return pulumi.get(self, "application_group_type")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="cloudPcResource")
    def cloud_pc_resource(self) -> builtins.bool:
        """
        Is cloud pc resource.
        """
        return pulumi.get(self, "cloud_pc_resource")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        Description of ApplicationGroup.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> builtins.str:
        """
        The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields. 
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[builtins.str]:
        """
        Friendly name of ApplicationGroup.
        """
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter(name="hostPoolArmPath")
    def host_pool_arm_path(self) -> builtins.str:
        """
        HostPool arm path of ApplicationGroup.
        """
        return pulumi.get(self, "host_pool_arm_path")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.ResourceModelWithAllowedPropertySetResponseIdentity']:
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def kind(self) -> Optional[builtins.str]:
        """
        Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type. E.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedBy")
    def managed_by(self) -> Optional[builtins.str]:
        """
        The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
        """
        return pulumi.get(self, "managed_by")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> builtins.str:
        """
        ObjectId of ApplicationGroup. (internal use)
        """
        return pulumi.get(self, "object_id")

    @property
    @pulumi.getter
    def plan(self) -> Optional['outputs.ResourceModelWithAllowedPropertySetResponsePlan']:
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="showInFeed")
    def show_in_feed(self) -> Optional[builtins.bool]:
        """
        Boolean representing whether the applicationGroup is show in the feed.
        """
        return pulumi.get(self, "show_in_feed")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.ResourceModelWithAllowedPropertySetResponseSku']:
        return pulumi.get(self, "sku")

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
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="workspaceArmPath")
    def workspace_arm_path(self) -> builtins.str:
        """
        Workspace arm path of ApplicationGroup.
        """
        return pulumi.get(self, "workspace_arm_path")


class AwaitableGetApplicationGroupResult(GetApplicationGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationGroupResult(
            application_group_type=self.application_group_type,
            azure_api_version=self.azure_api_version,
            cloud_pc_resource=self.cloud_pc_resource,
            description=self.description,
            etag=self.etag,
            friendly_name=self.friendly_name,
            host_pool_arm_path=self.host_pool_arm_path,
            id=self.id,
            identity=self.identity,
            kind=self.kind,
            location=self.location,
            managed_by=self.managed_by,
            name=self.name,
            object_id=self.object_id,
            plan=self.plan,
            show_in_feed=self.show_in_feed,
            sku=self.sku,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type,
            workspace_arm_path=self.workspace_arm_path)


def get_application_group(application_group_name: Optional[builtins.str] = None,
                          resource_group_name: Optional[builtins.str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationGroupResult:
    """
    Get an application group.

    Uses Azure REST API version 2024-04-03.

    Other available API versions: 2022-09-09, 2022-10-14-preview, 2023-09-05, 2023-10-04-preview, 2023-11-01-preview, 2024-01-16-preview, 2024-03-06-preview, 2024-04-08-preview, 2024-08-08-preview, 2024-11-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native desktopvirtualization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str application_group_name: The name of the application group
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['applicationGroupName'] = application_group_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:desktopvirtualization:getApplicationGroup', __args__, opts=opts, typ=GetApplicationGroupResult).value

    return AwaitableGetApplicationGroupResult(
        application_group_type=pulumi.get(__ret__, 'application_group_type'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        cloud_pc_resource=pulumi.get(__ret__, 'cloud_pc_resource'),
        description=pulumi.get(__ret__, 'description'),
        etag=pulumi.get(__ret__, 'etag'),
        friendly_name=pulumi.get(__ret__, 'friendly_name'),
        host_pool_arm_path=pulumi.get(__ret__, 'host_pool_arm_path'),
        id=pulumi.get(__ret__, 'id'),
        identity=pulumi.get(__ret__, 'identity'),
        kind=pulumi.get(__ret__, 'kind'),
        location=pulumi.get(__ret__, 'location'),
        managed_by=pulumi.get(__ret__, 'managed_by'),
        name=pulumi.get(__ret__, 'name'),
        object_id=pulumi.get(__ret__, 'object_id'),
        plan=pulumi.get(__ret__, 'plan'),
        show_in_feed=pulumi.get(__ret__, 'show_in_feed'),
        sku=pulumi.get(__ret__, 'sku'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        workspace_arm_path=pulumi.get(__ret__, 'workspace_arm_path'))
def get_application_group_output(application_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetApplicationGroupResult]:
    """
    Get an application group.

    Uses Azure REST API version 2024-04-03.

    Other available API versions: 2022-09-09, 2022-10-14-preview, 2023-09-05, 2023-10-04-preview, 2023-11-01-preview, 2024-01-16-preview, 2024-03-06-preview, 2024-04-08-preview, 2024-08-08-preview, 2024-11-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native desktopvirtualization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str application_group_name: The name of the application group
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['applicationGroupName'] = application_group_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:desktopvirtualization:getApplicationGroup', __args__, opts=opts, typ=GetApplicationGroupResult)
    return __ret__.apply(lambda __response__: GetApplicationGroupResult(
        application_group_type=pulumi.get(__response__, 'application_group_type'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        cloud_pc_resource=pulumi.get(__response__, 'cloud_pc_resource'),
        description=pulumi.get(__response__, 'description'),
        etag=pulumi.get(__response__, 'etag'),
        friendly_name=pulumi.get(__response__, 'friendly_name'),
        host_pool_arm_path=pulumi.get(__response__, 'host_pool_arm_path'),
        id=pulumi.get(__response__, 'id'),
        identity=pulumi.get(__response__, 'identity'),
        kind=pulumi.get(__response__, 'kind'),
        location=pulumi.get(__response__, 'location'),
        managed_by=pulumi.get(__response__, 'managed_by'),
        name=pulumi.get(__response__, 'name'),
        object_id=pulumi.get(__response__, 'object_id'),
        plan=pulumi.get(__response__, 'plan'),
        show_in_feed=pulumi.get(__response__, 'show_in_feed'),
        sku=pulumi.get(__response__, 'sku'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        workspace_arm_path=pulumi.get(__response__, 'workspace_arm_path')))
