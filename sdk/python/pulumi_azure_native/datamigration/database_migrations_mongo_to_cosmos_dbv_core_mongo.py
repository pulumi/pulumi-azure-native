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

__all__ = ['DatabaseMigrationsMongoToCosmosDbvCoreMongoArgs', 'DatabaseMigrationsMongoToCosmosDbvCoreMongo']

@pulumi.input_type
class DatabaseMigrationsMongoToCosmosDbvCoreMongoArgs:
    def __init__(__self__, *,
                 kind: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 target_resource_name: pulumi.Input[builtins.str],
                 collection_list: Optional[pulumi.Input[Sequence[pulumi.Input['MongoMigrationCollectionArgs']]]] = None,
                 migration_name: Optional[pulumi.Input[builtins.str]] = None,
                 migration_operation_id: Optional[pulumi.Input[builtins.str]] = None,
                 migration_service: Optional[pulumi.Input[builtins.str]] = None,
                 provisioning_error: Optional[pulumi.Input[builtins.str]] = None,
                 scope: Optional[pulumi.Input[builtins.str]] = None,
                 source_mongo_connection: Optional[pulumi.Input['MongoConnectionInformationArgs']] = None,
                 target_mongo_connection: Optional[pulumi.Input['MongoConnectionInformationArgs']] = None):
        """
        The set of arguments for constructing a DatabaseMigrationsMongoToCosmosDbvCoreMongo resource.
        :param pulumi.Input[builtins.str] kind: 
               Expected value is 'MongoToCosmosDbMongo'.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[builtins.str] target_resource_name: The name of the target resource/account.
        :param pulumi.Input[Sequence[pulumi.Input['MongoMigrationCollectionArgs']]] collection_list: List of Mongo Collections to be migrated.
        :param pulumi.Input[builtins.str] migration_name: Name of the migration.
        :param pulumi.Input[builtins.str] migration_operation_id: ID for current migration operation.
        :param pulumi.Input[builtins.str] migration_service: Resource Id of the Migration Service.
        :param pulumi.Input[builtins.str] provisioning_error: Error message for migration provisioning failure, if any.
        :param pulumi.Input[builtins.str] scope: Resource Id of the target resource.
        :param pulumi.Input['MongoConnectionInformationArgs'] source_mongo_connection: Source Mongo connection details.
        :param pulumi.Input['MongoConnectionInformationArgs'] target_mongo_connection: Target Cosmos DB Mongo connection details.
        """
        pulumi.set(__self__, "kind", 'MongoToCosmosDbMongo')
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "target_resource_name", target_resource_name)
        if collection_list is not None:
            pulumi.set(__self__, "collection_list", collection_list)
        if migration_name is not None:
            pulumi.set(__self__, "migration_name", migration_name)
        if migration_operation_id is not None:
            pulumi.set(__self__, "migration_operation_id", migration_operation_id)
        if migration_service is not None:
            pulumi.set(__self__, "migration_service", migration_service)
        if provisioning_error is not None:
            pulumi.set(__self__, "provisioning_error", provisioning_error)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if source_mongo_connection is not None:
            pulumi.set(__self__, "source_mongo_connection", source_mongo_connection)
        if target_mongo_connection is not None:
            pulumi.set(__self__, "target_mongo_connection", target_mongo_connection)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[builtins.str]:
        """

        Expected value is 'MongoToCosmosDbMongo'.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="targetResourceName")
    def target_resource_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the target resource/account.
        """
        return pulumi.get(self, "target_resource_name")

    @target_resource_name.setter
    def target_resource_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "target_resource_name", value)

    @property
    @pulumi.getter(name="collectionList")
    def collection_list(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MongoMigrationCollectionArgs']]]]:
        """
        List of Mongo Collections to be migrated.
        """
        return pulumi.get(self, "collection_list")

    @collection_list.setter
    def collection_list(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MongoMigrationCollectionArgs']]]]):
        pulumi.set(self, "collection_list", value)

    @property
    @pulumi.getter(name="migrationName")
    def migration_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the migration.
        """
        return pulumi.get(self, "migration_name")

    @migration_name.setter
    def migration_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "migration_name", value)

    @property
    @pulumi.getter(name="migrationOperationId")
    def migration_operation_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID for current migration operation.
        """
        return pulumi.get(self, "migration_operation_id")

    @migration_operation_id.setter
    def migration_operation_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "migration_operation_id", value)

    @property
    @pulumi.getter(name="migrationService")
    def migration_service(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Resource Id of the Migration Service.
        """
        return pulumi.get(self, "migration_service")

    @migration_service.setter
    def migration_service(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "migration_service", value)

    @property
    @pulumi.getter(name="provisioningError")
    def provisioning_error(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Error message for migration provisioning failure, if any.
        """
        return pulumi.get(self, "provisioning_error")

    @provisioning_error.setter
    def provisioning_error(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "provisioning_error", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Resource Id of the target resource.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter(name="sourceMongoConnection")
    def source_mongo_connection(self) -> Optional[pulumi.Input['MongoConnectionInformationArgs']]:
        """
        Source Mongo connection details.
        """
        return pulumi.get(self, "source_mongo_connection")

    @source_mongo_connection.setter
    def source_mongo_connection(self, value: Optional[pulumi.Input['MongoConnectionInformationArgs']]):
        pulumi.set(self, "source_mongo_connection", value)

    @property
    @pulumi.getter(name="targetMongoConnection")
    def target_mongo_connection(self) -> Optional[pulumi.Input['MongoConnectionInformationArgs']]:
        """
        Target Cosmos DB Mongo connection details.
        """
        return pulumi.get(self, "target_mongo_connection")

    @target_mongo_connection.setter
    def target_mongo_connection(self, value: Optional[pulumi.Input['MongoConnectionInformationArgs']]):
        pulumi.set(self, "target_mongo_connection", value)


@pulumi.type_token("azure-native:datamigration:DatabaseMigrationsMongoToCosmosDbvCoreMongo")
class DatabaseMigrationsMongoToCosmosDbvCoreMongo(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collection_list: Optional[pulumi.Input[Sequence[pulumi.Input[Union['MongoMigrationCollectionArgs', 'MongoMigrationCollectionArgsDict']]]]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 migration_name: Optional[pulumi.Input[builtins.str]] = None,
                 migration_operation_id: Optional[pulumi.Input[builtins.str]] = None,
                 migration_service: Optional[pulumi.Input[builtins.str]] = None,
                 provisioning_error: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 scope: Optional[pulumi.Input[builtins.str]] = None,
                 source_mongo_connection: Optional[pulumi.Input[Union['MongoConnectionInformationArgs', 'MongoConnectionInformationArgsDict']]] = None,
                 target_mongo_connection: Optional[pulumi.Input[Union['MongoConnectionInformationArgs', 'MongoConnectionInformationArgsDict']]] = None,
                 target_resource_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Database Migration Resource for Mongo to CosmosDb.

        Uses Azure REST API version 2023-07-15-preview. In version 2.x of the Azure Native provider, it used API version 2023-07-15-preview.

        Other available API versions: 2025-03-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native datamigration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['MongoMigrationCollectionArgs', 'MongoMigrationCollectionArgsDict']]]] collection_list: List of Mongo Collections to be migrated.
        :param pulumi.Input[builtins.str] kind: 
               Expected value is 'MongoToCosmosDbMongo'.
        :param pulumi.Input[builtins.str] migration_name: Name of the migration.
        :param pulumi.Input[builtins.str] migration_operation_id: ID for current migration operation.
        :param pulumi.Input[builtins.str] migration_service: Resource Id of the Migration Service.
        :param pulumi.Input[builtins.str] provisioning_error: Error message for migration provisioning failure, if any.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[builtins.str] scope: Resource Id of the target resource.
        :param pulumi.Input[Union['MongoConnectionInformationArgs', 'MongoConnectionInformationArgsDict']] source_mongo_connection: Source Mongo connection details.
        :param pulumi.Input[Union['MongoConnectionInformationArgs', 'MongoConnectionInformationArgsDict']] target_mongo_connection: Target Cosmos DB Mongo connection details.
        :param pulumi.Input[builtins.str] target_resource_name: The name of the target resource/account.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabaseMigrationsMongoToCosmosDbvCoreMongoArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Database Migration Resource for Mongo to CosmosDb.

        Uses Azure REST API version 2023-07-15-preview. In version 2.x of the Azure Native provider, it used API version 2023-07-15-preview.

        Other available API versions: 2025-03-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native datamigration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param DatabaseMigrationsMongoToCosmosDbvCoreMongoArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabaseMigrationsMongoToCosmosDbvCoreMongoArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collection_list: Optional[pulumi.Input[Sequence[pulumi.Input[Union['MongoMigrationCollectionArgs', 'MongoMigrationCollectionArgsDict']]]]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 migration_name: Optional[pulumi.Input[builtins.str]] = None,
                 migration_operation_id: Optional[pulumi.Input[builtins.str]] = None,
                 migration_service: Optional[pulumi.Input[builtins.str]] = None,
                 provisioning_error: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 scope: Optional[pulumi.Input[builtins.str]] = None,
                 source_mongo_connection: Optional[pulumi.Input[Union['MongoConnectionInformationArgs', 'MongoConnectionInformationArgsDict']]] = None,
                 target_mongo_connection: Optional[pulumi.Input[Union['MongoConnectionInformationArgs', 'MongoConnectionInformationArgsDict']]] = None,
                 target_resource_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabaseMigrationsMongoToCosmosDbvCoreMongoArgs.__new__(DatabaseMigrationsMongoToCosmosDbvCoreMongoArgs)

            __props__.__dict__["collection_list"] = collection_list
            if kind is None and not opts.urn:
                raise TypeError("Missing required property 'kind'")
            __props__.__dict__["kind"] = 'MongoToCosmosDbMongo'
            __props__.__dict__["migration_name"] = migration_name
            __props__.__dict__["migration_operation_id"] = migration_operation_id
            __props__.__dict__["migration_service"] = migration_service
            __props__.__dict__["provisioning_error"] = provisioning_error
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["scope"] = scope
            __props__.__dict__["source_mongo_connection"] = source_mongo_connection
            __props__.__dict__["target_mongo_connection"] = target_mongo_connection
            if target_resource_name is None and not opts.urn:
                raise TypeError("Missing required property 'target_resource_name'")
            __props__.__dict__["target_resource_name"] = target_resource_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["ended_on"] = None
            __props__.__dict__["migration_failure_error"] = None
            __props__.__dict__["migration_status"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["started_on"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:datamigration/v20230715preview:DatabaseMigrationsMongoToCosmosDbvCoreMongo"), pulumi.Alias(type_="azure-native:datamigration/v20250315preview:DatabaseMigrationsMongoToCosmosDbvCoreMongo")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DatabaseMigrationsMongoToCosmosDbvCoreMongo, __self__).__init__(
            'azure-native:datamigration:DatabaseMigrationsMongoToCosmosDbvCoreMongo',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DatabaseMigrationsMongoToCosmosDbvCoreMongo':
        """
        Get an existing DatabaseMigrationsMongoToCosmosDbvCoreMongo resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DatabaseMigrationsMongoToCosmosDbvCoreMongoArgs.__new__(DatabaseMigrationsMongoToCosmosDbvCoreMongoArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["collection_list"] = None
        __props__.__dict__["ended_on"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["migration_failure_error"] = None
        __props__.__dict__["migration_operation_id"] = None
        __props__.__dict__["migration_service"] = None
        __props__.__dict__["migration_status"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_error"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["scope"] = None
        __props__.__dict__["source_mongo_connection"] = None
        __props__.__dict__["started_on"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["target_mongo_connection"] = None
        __props__.__dict__["type"] = None
        return DatabaseMigrationsMongoToCosmosDbvCoreMongo(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="collectionList")
    def collection_list(self) -> pulumi.Output[Optional[Sequence['outputs.MongoMigrationCollectionResponse']]]:
        """
        List of Mongo Collections to be migrated.
        """
        return pulumi.get(self, "collection_list")

    @property
    @pulumi.getter(name="endedOn")
    def ended_on(self) -> pulumi.Output[builtins.str]:
        """
        Database migration end time.
        """
        return pulumi.get(self, "ended_on")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[builtins.str]:
        """

        Expected value is 'MongoToCosmosDbMongo'.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="migrationFailureError")
    def migration_failure_error(self) -> pulumi.Output['outputs.ErrorInfoResponse']:
        """
        Error details in case of migration failure.
        """
        return pulumi.get(self, "migration_failure_error")

    @property
    @pulumi.getter(name="migrationOperationId")
    def migration_operation_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        ID for current migration operation.
        """
        return pulumi.get(self, "migration_operation_id")

    @property
    @pulumi.getter(name="migrationService")
    def migration_service(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Resource Id of the Migration Service.
        """
        return pulumi.get(self, "migration_service")

    @property
    @pulumi.getter(name="migrationStatus")
    def migration_status(self) -> pulumi.Output[builtins.str]:
        """
        Migration status.
        """
        return pulumi.get(self, "migration_status")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningError")
    def provisioning_error(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Error message for migration provisioning failure, if any.
        """
        return pulumi.get(self, "provisioning_error")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Provisioning State of migration. ProvisioningState as Succeeded implies that validations have been performed and migration has started.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Resource Id of the target resource.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter(name="sourceMongoConnection")
    def source_mongo_connection(self) -> pulumi.Output[Optional['outputs.MongoConnectionInformationResponse']]:
        """
        Source Mongo connection details.
        """
        return pulumi.get(self, "source_mongo_connection")

    @property
    @pulumi.getter(name="startedOn")
    def started_on(self) -> pulumi.Output[builtins.str]:
        """
        Database migration start time.
        """
        return pulumi.get(self, "started_on")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter(name="targetMongoConnection")
    def target_mongo_connection(self) -> pulumi.Output[Optional['outputs.MongoConnectionInformationResponse']]:
        """
        Target Cosmos DB Mongo connection details.
        """
        return pulumi.get(self, "target_mongo_connection")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

