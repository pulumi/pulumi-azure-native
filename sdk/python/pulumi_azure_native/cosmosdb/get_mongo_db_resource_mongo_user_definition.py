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
    'GetMongoDBResourceMongoUserDefinitionResult',
    'AwaitableGetMongoDBResourceMongoUserDefinitionResult',
    'get_mongo_db_resource_mongo_user_definition',
    'get_mongo_db_resource_mongo_user_definition_output',
]

@pulumi.output_type
class GetMongoDBResourceMongoUserDefinitionResult:
    """
    An Azure Cosmos DB User Definition
    """
    def __init__(__self__, azure_api_version=None, custom_data=None, database_name=None, id=None, mechanisms=None, name=None, password=None, roles=None, type=None, user_name=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if custom_data and not isinstance(custom_data, str):
            raise TypeError("Expected argument 'custom_data' to be a str")
        pulumi.set(__self__, "custom_data", custom_data)
        if database_name and not isinstance(database_name, str):
            raise TypeError("Expected argument 'database_name' to be a str")
        pulumi.set(__self__, "database_name", database_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mechanisms and not isinstance(mechanisms, str):
            raise TypeError("Expected argument 'mechanisms' to be a str")
        pulumi.set(__self__, "mechanisms", mechanisms)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="customData")
    def custom_data(self) -> Optional[builtins.str]:
        """
        A custom definition for the USer Definition.
        """
        return pulumi.get(self, "custom_data")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[builtins.str]:
        """
        The database name for which access is being granted for this User Definition.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The unique resource identifier of the database account.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def mechanisms(self) -> Optional[builtins.str]:
        """
        The Mongo Auth mechanism. For now, we only support auth mechanism SCRAM-SHA-256.
        """
        return pulumi.get(self, "mechanisms")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the database account.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> Optional[builtins.str]:
        """
        The password for User Definition. Response does not contain user password.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def roles(self) -> Optional[Sequence['outputs.RoleResponse']]:
        """
        The set of roles inherited by the User Definition.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of Azure resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[builtins.str]:
        """
        The user name for User Definition.
        """
        return pulumi.get(self, "user_name")


class AwaitableGetMongoDBResourceMongoUserDefinitionResult(GetMongoDBResourceMongoUserDefinitionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMongoDBResourceMongoUserDefinitionResult(
            azure_api_version=self.azure_api_version,
            custom_data=self.custom_data,
            database_name=self.database_name,
            id=self.id,
            mechanisms=self.mechanisms,
            name=self.name,
            password=self.password,
            roles=self.roles,
            type=self.type,
            user_name=self.user_name)


def get_mongo_db_resource_mongo_user_definition(account_name: Optional[builtins.str] = None,
                                                mongo_user_definition_id: Optional[builtins.str] = None,
                                                resource_group_name: Optional[builtins.str] = None,
                                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMongoDBResourceMongoUserDefinitionResult:
    """
    Retrieves the properties of an existing Azure Cosmos DB Mongo User Definition with the given Id.

    Uses Azure REST API version 2024-11-15.

    Other available API versions: 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15-preview, 2022-08-15, 2022-08-15-preview, 2022-11-15, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15, 2023-03-15-preview, 2023-04-15, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-12-01-preview, 2025-04-15, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str account_name: Cosmos DB database account name.
    :param builtins.str mongo_user_definition_id: The ID for the User Definition {dbName.userName}.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['mongoUserDefinitionId'] = mongo_user_definition_id
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:cosmosdb:getMongoDBResourceMongoUserDefinition', __args__, opts=opts, typ=GetMongoDBResourceMongoUserDefinitionResult).value

    return AwaitableGetMongoDBResourceMongoUserDefinitionResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        custom_data=pulumi.get(__ret__, 'custom_data'),
        database_name=pulumi.get(__ret__, 'database_name'),
        id=pulumi.get(__ret__, 'id'),
        mechanisms=pulumi.get(__ret__, 'mechanisms'),
        name=pulumi.get(__ret__, 'name'),
        password=pulumi.get(__ret__, 'password'),
        roles=pulumi.get(__ret__, 'roles'),
        type=pulumi.get(__ret__, 'type'),
        user_name=pulumi.get(__ret__, 'user_name'))
def get_mongo_db_resource_mongo_user_definition_output(account_name: Optional[pulumi.Input[builtins.str]] = None,
                                                       mongo_user_definition_id: Optional[pulumi.Input[builtins.str]] = None,
                                                       resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMongoDBResourceMongoUserDefinitionResult]:
    """
    Retrieves the properties of an existing Azure Cosmos DB Mongo User Definition with the given Id.

    Uses Azure REST API version 2024-11-15.

    Other available API versions: 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15-preview, 2022-08-15, 2022-08-15-preview, 2022-11-15, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15, 2023-03-15-preview, 2023-04-15, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-12-01-preview, 2025-04-15, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str account_name: Cosmos DB database account name.
    :param builtins.str mongo_user_definition_id: The ID for the User Definition {dbName.userName}.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['mongoUserDefinitionId'] = mongo_user_definition_id
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:cosmosdb:getMongoDBResourceMongoUserDefinition', __args__, opts=opts, typ=GetMongoDBResourceMongoUserDefinitionResult)
    return __ret__.apply(lambda __response__: GetMongoDBResourceMongoUserDefinitionResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        custom_data=pulumi.get(__response__, 'custom_data'),
        database_name=pulumi.get(__response__, 'database_name'),
        id=pulumi.get(__response__, 'id'),
        mechanisms=pulumi.get(__response__, 'mechanisms'),
        name=pulumi.get(__response__, 'name'),
        password=pulumi.get(__response__, 'password'),
        roles=pulumi.get(__response__, 'roles'),
        type=pulumi.get(__response__, 'type'),
        user_name=pulumi.get(__response__, 'user_name')))
