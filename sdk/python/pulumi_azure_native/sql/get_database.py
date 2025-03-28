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
from .. import _utilities
from . import outputs

__all__ = [
    'GetDatabaseResult',
    'AwaitableGetDatabaseResult',
    'get_database',
    'get_database_output',
]

@pulumi.output_type
class GetDatabaseResult:
    """
    A database resource.
    """
    def __init__(__self__, auto_pause_delay=None, catalog_collation=None, collation=None, creation_date=None, current_backup_storage_redundancy=None, current_service_objective_name=None, current_sku=None, database_id=None, default_secondary_location=None, earliest_restore_date=None, elastic_pool_id=None, failover_group_id=None, federated_client_id=None, high_availability_replica_count=None, id=None, identity=None, is_infra_encryption_enabled=None, is_ledger_on=None, kind=None, license_type=None, location=None, maintenance_configuration_id=None, managed_by=None, max_log_size_bytes=None, max_size_bytes=None, min_capacity=None, name=None, paused_date=None, read_scale=None, requested_backup_storage_redundancy=None, requested_service_objective_name=None, resumed_date=None, secondary_type=None, sku=None, status=None, tags=None, type=None, zone_redundant=None):
        if auto_pause_delay and not isinstance(auto_pause_delay, int):
            raise TypeError("Expected argument 'auto_pause_delay' to be a int")
        pulumi.set(__self__, "auto_pause_delay", auto_pause_delay)
        if catalog_collation and not isinstance(catalog_collation, str):
            raise TypeError("Expected argument 'catalog_collation' to be a str")
        pulumi.set(__self__, "catalog_collation", catalog_collation)
        if collation and not isinstance(collation, str):
            raise TypeError("Expected argument 'collation' to be a str")
        pulumi.set(__self__, "collation", collation)
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if current_backup_storage_redundancy and not isinstance(current_backup_storage_redundancy, str):
            raise TypeError("Expected argument 'current_backup_storage_redundancy' to be a str")
        pulumi.set(__self__, "current_backup_storage_redundancy", current_backup_storage_redundancy)
        if current_service_objective_name and not isinstance(current_service_objective_name, str):
            raise TypeError("Expected argument 'current_service_objective_name' to be a str")
        pulumi.set(__self__, "current_service_objective_name", current_service_objective_name)
        if current_sku and not isinstance(current_sku, dict):
            raise TypeError("Expected argument 'current_sku' to be a dict")
        pulumi.set(__self__, "current_sku", current_sku)
        if database_id and not isinstance(database_id, str):
            raise TypeError("Expected argument 'database_id' to be a str")
        pulumi.set(__self__, "database_id", database_id)
        if default_secondary_location and not isinstance(default_secondary_location, str):
            raise TypeError("Expected argument 'default_secondary_location' to be a str")
        pulumi.set(__self__, "default_secondary_location", default_secondary_location)
        if earliest_restore_date and not isinstance(earliest_restore_date, str):
            raise TypeError("Expected argument 'earliest_restore_date' to be a str")
        pulumi.set(__self__, "earliest_restore_date", earliest_restore_date)
        if elastic_pool_id and not isinstance(elastic_pool_id, str):
            raise TypeError("Expected argument 'elastic_pool_id' to be a str")
        pulumi.set(__self__, "elastic_pool_id", elastic_pool_id)
        if failover_group_id and not isinstance(failover_group_id, str):
            raise TypeError("Expected argument 'failover_group_id' to be a str")
        pulumi.set(__self__, "failover_group_id", failover_group_id)
        if federated_client_id and not isinstance(federated_client_id, str):
            raise TypeError("Expected argument 'federated_client_id' to be a str")
        pulumi.set(__self__, "federated_client_id", federated_client_id)
        if high_availability_replica_count and not isinstance(high_availability_replica_count, int):
            raise TypeError("Expected argument 'high_availability_replica_count' to be a int")
        pulumi.set(__self__, "high_availability_replica_count", high_availability_replica_count)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if is_infra_encryption_enabled and not isinstance(is_infra_encryption_enabled, bool):
            raise TypeError("Expected argument 'is_infra_encryption_enabled' to be a bool")
        pulumi.set(__self__, "is_infra_encryption_enabled", is_infra_encryption_enabled)
        if is_ledger_on and not isinstance(is_ledger_on, bool):
            raise TypeError("Expected argument 'is_ledger_on' to be a bool")
        pulumi.set(__self__, "is_ledger_on", is_ledger_on)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if license_type and not isinstance(license_type, str):
            raise TypeError("Expected argument 'license_type' to be a str")
        pulumi.set(__self__, "license_type", license_type)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if maintenance_configuration_id and not isinstance(maintenance_configuration_id, str):
            raise TypeError("Expected argument 'maintenance_configuration_id' to be a str")
        pulumi.set(__self__, "maintenance_configuration_id", maintenance_configuration_id)
        if managed_by and not isinstance(managed_by, str):
            raise TypeError("Expected argument 'managed_by' to be a str")
        pulumi.set(__self__, "managed_by", managed_by)
        if max_log_size_bytes and not isinstance(max_log_size_bytes, float):
            raise TypeError("Expected argument 'max_log_size_bytes' to be a float")
        pulumi.set(__self__, "max_log_size_bytes", max_log_size_bytes)
        if max_size_bytes and not isinstance(max_size_bytes, float):
            raise TypeError("Expected argument 'max_size_bytes' to be a float")
        pulumi.set(__self__, "max_size_bytes", max_size_bytes)
        if min_capacity and not isinstance(min_capacity, float):
            raise TypeError("Expected argument 'min_capacity' to be a float")
        pulumi.set(__self__, "min_capacity", min_capacity)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if paused_date and not isinstance(paused_date, str):
            raise TypeError("Expected argument 'paused_date' to be a str")
        pulumi.set(__self__, "paused_date", paused_date)
        if read_scale and not isinstance(read_scale, str):
            raise TypeError("Expected argument 'read_scale' to be a str")
        pulumi.set(__self__, "read_scale", read_scale)
        if requested_backup_storage_redundancy and not isinstance(requested_backup_storage_redundancy, str):
            raise TypeError("Expected argument 'requested_backup_storage_redundancy' to be a str")
        pulumi.set(__self__, "requested_backup_storage_redundancy", requested_backup_storage_redundancy)
        if requested_service_objective_name and not isinstance(requested_service_objective_name, str):
            raise TypeError("Expected argument 'requested_service_objective_name' to be a str")
        pulumi.set(__self__, "requested_service_objective_name", requested_service_objective_name)
        if resumed_date and not isinstance(resumed_date, str):
            raise TypeError("Expected argument 'resumed_date' to be a str")
        pulumi.set(__self__, "resumed_date", resumed_date)
        if secondary_type and not isinstance(secondary_type, str):
            raise TypeError("Expected argument 'secondary_type' to be a str")
        pulumi.set(__self__, "secondary_type", secondary_type)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if zone_redundant and not isinstance(zone_redundant, bool):
            raise TypeError("Expected argument 'zone_redundant' to be a bool")
        pulumi.set(__self__, "zone_redundant", zone_redundant)

    @property
    @pulumi.getter(name="autoPauseDelay")
    def auto_pause_delay(self) -> Optional[int]:
        """
        Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
        """
        return pulumi.get(self, "auto_pause_delay")

    @property
    @pulumi.getter(name="catalogCollation")
    def catalog_collation(self) -> Optional[str]:
        """
        Collation of the metadata catalog.
        """
        return pulumi.get(self, "catalog_collation")

    @property
    @pulumi.getter
    def collation(self) -> Optional[str]:
        """
        The collation of the database.
        """
        return pulumi.get(self, "collation")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        """
        The creation date of the database (ISO8601 format).
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="currentBackupStorageRedundancy")
    def current_backup_storage_redundancy(self) -> str:
        """
        The storage account type used to store backups for this database.
        """
        return pulumi.get(self, "current_backup_storage_redundancy")

    @property
    @pulumi.getter(name="currentServiceObjectiveName")
    def current_service_objective_name(self) -> str:
        """
        The current service level objective name of the database.
        """
        return pulumi.get(self, "current_service_objective_name")

    @property
    @pulumi.getter(name="currentSku")
    def current_sku(self) -> 'outputs.SkuResponse':
        """
        The name and tier of the SKU.
        """
        return pulumi.get(self, "current_sku")

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> str:
        """
        The ID of the database.
        """
        return pulumi.get(self, "database_id")

    @property
    @pulumi.getter(name="defaultSecondaryLocation")
    def default_secondary_location(self) -> str:
        """
        The default secondary region for this database.
        """
        return pulumi.get(self, "default_secondary_location")

    @property
    @pulumi.getter(name="earliestRestoreDate")
    def earliest_restore_date(self) -> str:
        """
        This records the earliest start date and time that restore is available for this database (ISO8601 format).
        """
        return pulumi.get(self, "earliest_restore_date")

    @property
    @pulumi.getter(name="elasticPoolId")
    def elastic_pool_id(self) -> Optional[str]:
        """
        The resource identifier of the elastic pool containing this database.
        """
        return pulumi.get(self, "elastic_pool_id")

    @property
    @pulumi.getter(name="failoverGroupId")
    def failover_group_id(self) -> str:
        """
        Failover Group resource identifier that this database belongs to.
        """
        return pulumi.get(self, "failover_group_id")

    @property
    @pulumi.getter(name="federatedClientId")
    def federated_client_id(self) -> Optional[str]:
        """
        The Client id used for cross tenant per database CMK scenario
        """
        return pulumi.get(self, "federated_client_id")

    @property
    @pulumi.getter(name="highAvailabilityReplicaCount")
    def high_availability_replica_count(self) -> Optional[int]:
        """
        The number of secondary replicas associated with the database that are used to provide high availability. Not applicable to a Hyperscale database within an elastic pool.
        """
        return pulumi.get(self, "high_availability_replica_count")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.DatabaseIdentityResponse']:
        """
        The Azure Active Directory identity of the database.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="isInfraEncryptionEnabled")
    def is_infra_encryption_enabled(self) -> bool:
        """
        Infra encryption is enabled for this database.
        """
        return pulumi.get(self, "is_infra_encryption_enabled")

    @property
    @pulumi.getter(name="isLedgerOn")
    def is_ledger_on(self) -> Optional[bool]:
        """
        Whether or not this database is a ledger database, which means all tables in the database are ledger tables. Note: the value of this property cannot be changed after the database has been created.
        """
        return pulumi.get(self, "is_ledger_on")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Kind of database. This is metadata used for the Azure portal experience.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="licenseType")
    def license_type(self) -> Optional[str]:
        """
        The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
        """
        return pulumi.get(self, "license_type")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maintenanceConfigurationId")
    def maintenance_configuration_id(self) -> Optional[str]:
        """
        Maintenance configuration id assigned to the database. This configuration defines the period when the maintenance updates will occur.
        """
        return pulumi.get(self, "maintenance_configuration_id")

    @property
    @pulumi.getter(name="managedBy")
    def managed_by(self) -> str:
        """
        Resource that manages the database.
        """
        return pulumi.get(self, "managed_by")

    @property
    @pulumi.getter(name="maxLogSizeBytes")
    def max_log_size_bytes(self) -> float:
        """
        The max log size for this database.
        """
        return pulumi.get(self, "max_log_size_bytes")

    @property
    @pulumi.getter(name="maxSizeBytes")
    def max_size_bytes(self) -> Optional[float]:
        """
        The max size of the database expressed in bytes.
        """
        return pulumi.get(self, "max_size_bytes")

    @property
    @pulumi.getter(name="minCapacity")
    def min_capacity(self) -> Optional[float]:
        """
        Minimal capacity that database will always have allocated, if not paused
        """
        return pulumi.get(self, "min_capacity")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pausedDate")
    def paused_date(self) -> str:
        """
        The date when database was paused by user configuration or action(ISO8601 format). Null if the database is ready.
        """
        return pulumi.get(self, "paused_date")

    @property
    @pulumi.getter(name="readScale")
    def read_scale(self) -> Optional[str]:
        """
        The state of read-only routing. If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region. Not applicable to a Hyperscale database within an elastic pool.
        """
        return pulumi.get(self, "read_scale")

    @property
    @pulumi.getter(name="requestedBackupStorageRedundancy")
    def requested_backup_storage_redundancy(self) -> Optional[str]:
        """
        The storage account type to be used to store backups for this database.
        """
        return pulumi.get(self, "requested_backup_storage_redundancy")

    @property
    @pulumi.getter(name="requestedServiceObjectiveName")
    def requested_service_objective_name(self) -> str:
        """
        The requested service level objective name of the database.
        """
        return pulumi.get(self, "requested_service_objective_name")

    @property
    @pulumi.getter(name="resumedDate")
    def resumed_date(self) -> str:
        """
        The date when database was resumed by user action or database login (ISO8601 format). Null if the database is paused.
        """
        return pulumi.get(self, "resumed_date")

    @property
    @pulumi.getter(name="secondaryType")
    def secondary_type(self) -> Optional[str]:
        """
        The secondary type of the database if it is a secondary.  Valid values are Geo and Named.
        """
        return pulumi.get(self, "secondary_type")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.SkuResponse']:
        """
        The database SKU.
        
        The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
        
        ```azurecli
        az sql db list-editions -l <location> -o table
        ````
        
        ```powershell
        Get-AzSqlServerServiceObjective -Location <location>
        ````
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the database.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="zoneRedundant")
    def zone_redundant(self) -> Optional[bool]:
        """
        Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
        """
        return pulumi.get(self, "zone_redundant")


class AwaitableGetDatabaseResult(GetDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseResult(
            auto_pause_delay=self.auto_pause_delay,
            catalog_collation=self.catalog_collation,
            collation=self.collation,
            creation_date=self.creation_date,
            current_backup_storage_redundancy=self.current_backup_storage_redundancy,
            current_service_objective_name=self.current_service_objective_name,
            current_sku=self.current_sku,
            database_id=self.database_id,
            default_secondary_location=self.default_secondary_location,
            earliest_restore_date=self.earliest_restore_date,
            elastic_pool_id=self.elastic_pool_id,
            failover_group_id=self.failover_group_id,
            federated_client_id=self.federated_client_id,
            high_availability_replica_count=self.high_availability_replica_count,
            id=self.id,
            identity=self.identity,
            is_infra_encryption_enabled=self.is_infra_encryption_enabled,
            is_ledger_on=self.is_ledger_on,
            kind=self.kind,
            license_type=self.license_type,
            location=self.location,
            maintenance_configuration_id=self.maintenance_configuration_id,
            managed_by=self.managed_by,
            max_log_size_bytes=self.max_log_size_bytes,
            max_size_bytes=self.max_size_bytes,
            min_capacity=self.min_capacity,
            name=self.name,
            paused_date=self.paused_date,
            read_scale=self.read_scale,
            requested_backup_storage_redundancy=self.requested_backup_storage_redundancy,
            requested_service_objective_name=self.requested_service_objective_name,
            resumed_date=self.resumed_date,
            secondary_type=self.secondary_type,
            sku=self.sku,
            status=self.status,
            tags=self.tags,
            type=self.type,
            zone_redundant=self.zone_redundant)


def get_database(database_name: Optional[str] = None,
                 resource_group_name: Optional[str] = None,
                 server_name: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseResult:
    """
    Gets a database.

    Uses Azure REST API version 2021-11-01.

    Other available API versions: 2014-04-01, 2019-06-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01, 2023-08-01-preview, 2024-05-01-preview.


    :param str database_name: The name of the database.
    :param str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:sql:getDatabase', __args__, opts=opts, typ=GetDatabaseResult).value

    return AwaitableGetDatabaseResult(
        auto_pause_delay=pulumi.get(__ret__, 'auto_pause_delay'),
        catalog_collation=pulumi.get(__ret__, 'catalog_collation'),
        collation=pulumi.get(__ret__, 'collation'),
        creation_date=pulumi.get(__ret__, 'creation_date'),
        current_backup_storage_redundancy=pulumi.get(__ret__, 'current_backup_storage_redundancy'),
        current_service_objective_name=pulumi.get(__ret__, 'current_service_objective_name'),
        current_sku=pulumi.get(__ret__, 'current_sku'),
        database_id=pulumi.get(__ret__, 'database_id'),
        default_secondary_location=pulumi.get(__ret__, 'default_secondary_location'),
        earliest_restore_date=pulumi.get(__ret__, 'earliest_restore_date'),
        elastic_pool_id=pulumi.get(__ret__, 'elastic_pool_id'),
        failover_group_id=pulumi.get(__ret__, 'failover_group_id'),
        federated_client_id=pulumi.get(__ret__, 'federated_client_id'),
        high_availability_replica_count=pulumi.get(__ret__, 'high_availability_replica_count'),
        id=pulumi.get(__ret__, 'id'),
        identity=pulumi.get(__ret__, 'identity'),
        is_infra_encryption_enabled=pulumi.get(__ret__, 'is_infra_encryption_enabled'),
        is_ledger_on=pulumi.get(__ret__, 'is_ledger_on'),
        kind=pulumi.get(__ret__, 'kind'),
        license_type=pulumi.get(__ret__, 'license_type'),
        location=pulumi.get(__ret__, 'location'),
        maintenance_configuration_id=pulumi.get(__ret__, 'maintenance_configuration_id'),
        managed_by=pulumi.get(__ret__, 'managed_by'),
        max_log_size_bytes=pulumi.get(__ret__, 'max_log_size_bytes'),
        max_size_bytes=pulumi.get(__ret__, 'max_size_bytes'),
        min_capacity=pulumi.get(__ret__, 'min_capacity'),
        name=pulumi.get(__ret__, 'name'),
        paused_date=pulumi.get(__ret__, 'paused_date'),
        read_scale=pulumi.get(__ret__, 'read_scale'),
        requested_backup_storage_redundancy=pulumi.get(__ret__, 'requested_backup_storage_redundancy'),
        requested_service_objective_name=pulumi.get(__ret__, 'requested_service_objective_name'),
        resumed_date=pulumi.get(__ret__, 'resumed_date'),
        secondary_type=pulumi.get(__ret__, 'secondary_type'),
        sku=pulumi.get(__ret__, 'sku'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        zone_redundant=pulumi.get(__ret__, 'zone_redundant'))
def get_database_output(database_name: Optional[pulumi.Input[str]] = None,
                        resource_group_name: Optional[pulumi.Input[str]] = None,
                        server_name: Optional[pulumi.Input[str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDatabaseResult]:
    """
    Gets a database.

    Uses Azure REST API version 2021-11-01.

    Other available API versions: 2014-04-01, 2019-06-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01, 2023-08-01-preview, 2024-05-01-preview.


    :param str database_name: The name of the database.
    :param str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:sql:getDatabase', __args__, opts=opts, typ=GetDatabaseResult)
    return __ret__.apply(lambda __response__: GetDatabaseResult(
        auto_pause_delay=pulumi.get(__response__, 'auto_pause_delay'),
        catalog_collation=pulumi.get(__response__, 'catalog_collation'),
        collation=pulumi.get(__response__, 'collation'),
        creation_date=pulumi.get(__response__, 'creation_date'),
        current_backup_storage_redundancy=pulumi.get(__response__, 'current_backup_storage_redundancy'),
        current_service_objective_name=pulumi.get(__response__, 'current_service_objective_name'),
        current_sku=pulumi.get(__response__, 'current_sku'),
        database_id=pulumi.get(__response__, 'database_id'),
        default_secondary_location=pulumi.get(__response__, 'default_secondary_location'),
        earliest_restore_date=pulumi.get(__response__, 'earliest_restore_date'),
        elastic_pool_id=pulumi.get(__response__, 'elastic_pool_id'),
        failover_group_id=pulumi.get(__response__, 'failover_group_id'),
        federated_client_id=pulumi.get(__response__, 'federated_client_id'),
        high_availability_replica_count=pulumi.get(__response__, 'high_availability_replica_count'),
        id=pulumi.get(__response__, 'id'),
        identity=pulumi.get(__response__, 'identity'),
        is_infra_encryption_enabled=pulumi.get(__response__, 'is_infra_encryption_enabled'),
        is_ledger_on=pulumi.get(__response__, 'is_ledger_on'),
        kind=pulumi.get(__response__, 'kind'),
        license_type=pulumi.get(__response__, 'license_type'),
        location=pulumi.get(__response__, 'location'),
        maintenance_configuration_id=pulumi.get(__response__, 'maintenance_configuration_id'),
        managed_by=pulumi.get(__response__, 'managed_by'),
        max_log_size_bytes=pulumi.get(__response__, 'max_log_size_bytes'),
        max_size_bytes=pulumi.get(__response__, 'max_size_bytes'),
        min_capacity=pulumi.get(__response__, 'min_capacity'),
        name=pulumi.get(__response__, 'name'),
        paused_date=pulumi.get(__response__, 'paused_date'),
        read_scale=pulumi.get(__response__, 'read_scale'),
        requested_backup_storage_redundancy=pulumi.get(__response__, 'requested_backup_storage_redundancy'),
        requested_service_objective_name=pulumi.get(__response__, 'requested_service_objective_name'),
        resumed_date=pulumi.get(__response__, 'resumed_date'),
        secondary_type=pulumi.get(__response__, 'secondary_type'),
        sku=pulumi.get(__response__, 'sku'),
        status=pulumi.get(__response__, 'status'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        zone_redundant=pulumi.get(__response__, 'zone_redundant')))
