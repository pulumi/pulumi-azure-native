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
    'GetSchemaResult',
    'AwaitableGetSchemaResult',
    'get_schema',
    'get_schema_output',
]

@pulumi.output_type
class GetSchemaResult:
    """
    Schema definition.
    """
    def __init__(__self__, azure_api_version=None, description=None, display_name=None, format=None, id=None, name=None, provisioning_state=None, schema_type=None, system_data=None, tags=None, type=None, uuid=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if format and not isinstance(format, str):
            raise TypeError("Expected argument 'format' to be a str")
        pulumi.set(__self__, "format", format)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if schema_type and not isinstance(schema_type, str):
            raise TypeError("Expected argument 'schema_type' to be a str")
        pulumi.set(__self__, "schema_type", schema_type)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if uuid and not isinstance(uuid, str):
            raise TypeError("Expected argument 'uuid' to be a str")
        pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        Human-readable description of the schema.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[builtins.str]:
        """
        Human-readable display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def format(self) -> builtins.str:
        """
        Format of the schema.
        """
        return pulumi.get(self, "format")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

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
        Provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="schemaType")
    def schema_type(self) -> builtins.str:
        """
        Type of the schema.
        """
        return pulumi.get(self, "schema_type")

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
        Schema tags.
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
    @pulumi.getter
    def uuid(self) -> builtins.str:
        """
        Globally unique, immutable, non-reusable id.
        """
        return pulumi.get(self, "uuid")


class AwaitableGetSchemaResult(GetSchemaResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSchemaResult(
            azure_api_version=self.azure_api_version,
            description=self.description,
            display_name=self.display_name,
            format=self.format,
            id=self.id,
            name=self.name,
            provisioning_state=self.provisioning_state,
            schema_type=self.schema_type,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type,
            uuid=self.uuid)


def get_schema(resource_group_name: Optional[builtins.str] = None,
               schema_name: Optional[builtins.str] = None,
               schema_registry_name: Optional[builtins.str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSchemaResult:
    """
    Get a Schema

    Uses Azure REST API version 2024-09-01-preview.

    Other available API versions: 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native deviceregistry [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str schema_name: Schema name parameter.
    :param builtins.str schema_registry_name: Schema registry name parameter.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['schemaName'] = schema_name
    __args__['schemaRegistryName'] = schema_registry_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:deviceregistry:getSchema', __args__, opts=opts, typ=GetSchemaResult).value

    return AwaitableGetSchemaResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        description=pulumi.get(__ret__, 'description'),
        display_name=pulumi.get(__ret__, 'display_name'),
        format=pulumi.get(__ret__, 'format'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        schema_type=pulumi.get(__ret__, 'schema_type'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        uuid=pulumi.get(__ret__, 'uuid'))
def get_schema_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                      schema_name: Optional[pulumi.Input[builtins.str]] = None,
                      schema_registry_name: Optional[pulumi.Input[builtins.str]] = None,
                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSchemaResult]:
    """
    Get a Schema

    Uses Azure REST API version 2024-09-01-preview.

    Other available API versions: 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native deviceregistry [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str schema_name: Schema name parameter.
    :param builtins.str schema_registry_name: Schema registry name parameter.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['schemaName'] = schema_name
    __args__['schemaRegistryName'] = schema_registry_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:deviceregistry:getSchema', __args__, opts=opts, typ=GetSchemaResult)
    return __ret__.apply(lambda __response__: GetSchemaResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        description=pulumi.get(__response__, 'description'),
        display_name=pulumi.get(__response__, 'display_name'),
        format=pulumi.get(__response__, 'format'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        schema_type=pulumi.get(__response__, 'schema_type'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        uuid=pulumi.get(__response__, 'uuid')))
