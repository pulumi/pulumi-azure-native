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
    'GetStorageTaskResult',
    'AwaitableGetStorageTaskResult',
    'get_storage_task',
    'get_storage_task_output',
]

@pulumi.output_type
class GetStorageTaskResult:
    """
    Represents Storage Task.
    """
    def __init__(__self__, action=None, azure_api_version=None, creation_time_in_utc=None, description=None, enabled=None, id=None, identity=None, location=None, name=None, provisioning_state=None, system_data=None, tags=None, task_version=None, type=None):
        if action and not isinstance(action, dict):
            raise TypeError("Expected argument 'action' to be a dict")
        pulumi.set(__self__, "action", action)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if creation_time_in_utc and not isinstance(creation_time_in_utc, str):
            raise TypeError("Expected argument 'creation_time_in_utc' to be a str")
        pulumi.set(__self__, "creation_time_in_utc", creation_time_in_utc)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if task_version and not isinstance(task_version, float):
            raise TypeError("Expected argument 'task_version' to be a float")
        pulumi.set(__self__, "task_version", task_version)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def action(self) -> 'outputs.StorageTaskActionResponse':
        """
        The storage task action that is executed
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="creationTimeInUtc")
    def creation_time_in_utc(self) -> builtins.str:
        """
        The creation date and time of the storage task in UTC.
        """
        return pulumi.get(self, "creation_time_in_utc")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        Text that describes the purpose of the storage task
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> builtins.bool:
        """
        Storage Task is enabled when set to true and disabled when set to false
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identity(self) -> 'outputs.ManagedServiceIdentityResponse':
        """
        The managed service identity of the resource.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Represents the provisioning state of the storage task.
        """
        return pulumi.get(self, "provisioning_state")

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
    @pulumi.getter(name="taskVersion")
    def task_version(self) -> builtins.float:
        """
        Storage task version.
        """
        return pulumi.get(self, "task_version")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetStorageTaskResult(GetStorageTaskResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStorageTaskResult(
            action=self.action,
            azure_api_version=self.azure_api_version,
            creation_time_in_utc=self.creation_time_in_utc,
            description=self.description,
            enabled=self.enabled,
            id=self.id,
            identity=self.identity,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            system_data=self.system_data,
            tags=self.tags,
            task_version=self.task_version,
            type=self.type)


def get_storage_task(resource_group_name: Optional[builtins.str] = None,
                     storage_task_name: Optional[builtins.str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStorageTaskResult:
    """
    Get the storage task properties

    Uses Azure REST API version 2023-01-01.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str storage_task_name: The name of the storage task within the specified resource group. Storage task names must be between 3 and 18 characters in length and use numbers and lower-case letters only.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['storageTaskName'] = storage_task_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:storageactions:getStorageTask', __args__, opts=opts, typ=GetStorageTaskResult).value

    return AwaitableGetStorageTaskResult(
        action=pulumi.get(__ret__, 'action'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        creation_time_in_utc=pulumi.get(__ret__, 'creation_time_in_utc'),
        description=pulumi.get(__ret__, 'description'),
        enabled=pulumi.get(__ret__, 'enabled'),
        id=pulumi.get(__ret__, 'id'),
        identity=pulumi.get(__ret__, 'identity'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        task_version=pulumi.get(__ret__, 'task_version'),
        type=pulumi.get(__ret__, 'type'))
def get_storage_task_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                            storage_task_name: Optional[pulumi.Input[builtins.str]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetStorageTaskResult]:
    """
    Get the storage task properties

    Uses Azure REST API version 2023-01-01.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str storage_task_name: The name of the storage task within the specified resource group. Storage task names must be between 3 and 18 characters in length and use numbers and lower-case letters only.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['storageTaskName'] = storage_task_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:storageactions:getStorageTask', __args__, opts=opts, typ=GetStorageTaskResult)
    return __ret__.apply(lambda __response__: GetStorageTaskResult(
        action=pulumi.get(__response__, 'action'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        creation_time_in_utc=pulumi.get(__response__, 'creation_time_in_utc'),
        description=pulumi.get(__response__, 'description'),
        enabled=pulumi.get(__response__, 'enabled'),
        id=pulumi.get(__response__, 'id'),
        identity=pulumi.get(__response__, 'identity'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        task_version=pulumi.get(__response__, 'task_version'),
        type=pulumi.get(__response__, 'type')))
