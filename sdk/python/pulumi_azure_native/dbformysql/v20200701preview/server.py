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
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['ServerArgs', 'Server']

@pulumi.input_type
class ServerArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[str],
                 administrator_login: Optional[pulumi.Input[str]] = None,
                 administrator_login_password: Optional[pulumi.Input[str]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 create_mode: Optional[pulumi.Input[Union[str, 'CreateMode']]] = None,
                 delegated_subnet_arguments: Optional[pulumi.Input['DelegatedSubnetArgumentsArgs']] = None,
                 ha_enabled: Optional[pulumi.Input[Union[str, 'HaEnabledEnum']]] = None,
                 identity: Optional[pulumi.Input['IdentityArgs']] = None,
                 infrastructure_encryption: Optional[pulumi.Input[Union[str, 'InfrastructureEncryptionEnum']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 maintenance_window: Optional[pulumi.Input['MaintenanceWindowArgs']] = None,
                 private_dns_zone_arguments: Optional[pulumi.Input['PrivateDnsZoneArgumentsArgs']] = None,
                 replication_role: Optional[pulumi.Input[str]] = None,
                 restore_point_in_time: Optional[pulumi.Input[str]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input['SkuArgs']] = None,
                 source_server_id: Optional[pulumi.Input[str]] = None,
                 ssl_enforcement: Optional[pulumi.Input[Union[str, 'SslEnforcementEnum']]] = None,
                 storage_profile: Optional[pulumi.Input['StorageProfileArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 version: Optional[pulumi.Input[Union[str, 'ServerVersion']]] = None):
        """
        The set of arguments for constructing a Server resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] administrator_login: The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
        :param pulumi.Input[str] administrator_login_password: The password of the administrator login (required for server creation).
        :param pulumi.Input[str] availability_zone: availability Zone information of the server.
        :param pulumi.Input[Union[str, 'CreateMode']] create_mode: The mode to create a new MySQL server.
        :param pulumi.Input['DelegatedSubnetArgumentsArgs'] delegated_subnet_arguments: Delegated subnet arguments.
        :param pulumi.Input[Union[str, 'HaEnabledEnum']] ha_enabled: Enable HA or not for a server.
        :param pulumi.Input['IdentityArgs'] identity: The Azure Active Directory identity of the server.
        :param pulumi.Input[Union[str, 'InfrastructureEncryptionEnum']] infrastructure_encryption: Status showing whether the server enabled infrastructure encryption.
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input['MaintenanceWindowArgs'] maintenance_window: Maintenance window of a server.
        :param pulumi.Input['PrivateDnsZoneArgumentsArgs'] private_dns_zone_arguments: private dns zone arguments.
        :param pulumi.Input[str] replication_role: The replication role.
        :param pulumi.Input[str] restore_point_in_time: Restore point creation time (ISO8601 format), specifying the time to restore from.
        :param pulumi.Input[str] server_name: The name of the server.
        :param pulumi.Input['SkuArgs'] sku: The SKU (pricing tier) of the server.
        :param pulumi.Input[str] source_server_id: The source MySQL server id.
        :param pulumi.Input[Union[str, 'SslEnforcementEnum']] ssl_enforcement: Enable ssl enforcement or not when connect to server.
        :param pulumi.Input['StorageProfileArgs'] storage_profile: Storage profile of a server.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[Union[str, 'ServerVersion']] version: Server version.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if administrator_login is not None:
            pulumi.set(__self__, "administrator_login", administrator_login)
        if administrator_login_password is not None:
            pulumi.set(__self__, "administrator_login_password", administrator_login_password)
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if create_mode is not None:
            pulumi.set(__self__, "create_mode", create_mode)
        if delegated_subnet_arguments is not None:
            pulumi.set(__self__, "delegated_subnet_arguments", delegated_subnet_arguments)
        if ha_enabled is not None:
            pulumi.set(__self__, "ha_enabled", ha_enabled)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if infrastructure_encryption is not None:
            pulumi.set(__self__, "infrastructure_encryption", infrastructure_encryption)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if maintenance_window is not None:
            pulumi.set(__self__, "maintenance_window", maintenance_window)
        if private_dns_zone_arguments is not None:
            pulumi.set(__self__, "private_dns_zone_arguments", private_dns_zone_arguments)
        if replication_role is not None:
            pulumi.set(__self__, "replication_role", replication_role)
        if restore_point_in_time is not None:
            pulumi.set(__self__, "restore_point_in_time", restore_point_in_time)
        if server_name is not None:
            pulumi.set(__self__, "server_name", server_name)
        if sku is not None:
            pulumi.set(__self__, "sku", sku)
        if source_server_id is not None:
            pulumi.set(__self__, "source_server_id", source_server_id)
        if ssl_enforcement is not None:
            pulumi.set(__self__, "ssl_enforcement", ssl_enforcement)
        if storage_profile is not None:
            pulumi.set(__self__, "storage_profile", storage_profile)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group. The name is case insensitive.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="administratorLogin")
    def administrator_login(self) -> Optional[pulumi.Input[str]]:
        """
        The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
        """
        return pulumi.get(self, "administrator_login")

    @administrator_login.setter
    def administrator_login(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "administrator_login", value)

    @property
    @pulumi.getter(name="administratorLoginPassword")
    def administrator_login_password(self) -> Optional[pulumi.Input[str]]:
        """
        The password of the administrator login (required for server creation).
        """
        return pulumi.get(self, "administrator_login_password")

    @administrator_login_password.setter
    def administrator_login_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "administrator_login_password", value)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        """
        availability Zone information of the server.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="createMode")
    def create_mode(self) -> Optional[pulumi.Input[Union[str, 'CreateMode']]]:
        """
        The mode to create a new MySQL server.
        """
        return pulumi.get(self, "create_mode")

    @create_mode.setter
    def create_mode(self, value: Optional[pulumi.Input[Union[str, 'CreateMode']]]):
        pulumi.set(self, "create_mode", value)

    @property
    @pulumi.getter(name="delegatedSubnetArguments")
    def delegated_subnet_arguments(self) -> Optional[pulumi.Input['DelegatedSubnetArgumentsArgs']]:
        """
        Delegated subnet arguments.
        """
        return pulumi.get(self, "delegated_subnet_arguments")

    @delegated_subnet_arguments.setter
    def delegated_subnet_arguments(self, value: Optional[pulumi.Input['DelegatedSubnetArgumentsArgs']]):
        pulumi.set(self, "delegated_subnet_arguments", value)

    @property
    @pulumi.getter(name="haEnabled")
    def ha_enabled(self) -> Optional[pulumi.Input[Union[str, 'HaEnabledEnum']]]:
        """
        Enable HA or not for a server.
        """
        return pulumi.get(self, "ha_enabled")

    @ha_enabled.setter
    def ha_enabled(self, value: Optional[pulumi.Input[Union[str, 'HaEnabledEnum']]]):
        pulumi.set(self, "ha_enabled", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['IdentityArgs']]:
        """
        The Azure Active Directory identity of the server.
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['IdentityArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter(name="infrastructureEncryption")
    def infrastructure_encryption(self) -> Optional[pulumi.Input[Union[str, 'InfrastructureEncryptionEnum']]]:
        """
        Status showing whether the server enabled infrastructure encryption.
        """
        return pulumi.get(self, "infrastructure_encryption")

    @infrastructure_encryption.setter
    def infrastructure_encryption(self, value: Optional[pulumi.Input[Union[str, 'InfrastructureEncryptionEnum']]]):
        pulumi.set(self, "infrastructure_encryption", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> Optional[pulumi.Input['MaintenanceWindowArgs']]:
        """
        Maintenance window of a server.
        """
        return pulumi.get(self, "maintenance_window")

    @maintenance_window.setter
    def maintenance_window(self, value: Optional[pulumi.Input['MaintenanceWindowArgs']]):
        pulumi.set(self, "maintenance_window", value)

    @property
    @pulumi.getter(name="privateDnsZoneArguments")
    def private_dns_zone_arguments(self) -> Optional[pulumi.Input['PrivateDnsZoneArgumentsArgs']]:
        """
        private dns zone arguments.
        """
        return pulumi.get(self, "private_dns_zone_arguments")

    @private_dns_zone_arguments.setter
    def private_dns_zone_arguments(self, value: Optional[pulumi.Input['PrivateDnsZoneArgumentsArgs']]):
        pulumi.set(self, "private_dns_zone_arguments", value)

    @property
    @pulumi.getter(name="replicationRole")
    def replication_role(self) -> Optional[pulumi.Input[str]]:
        """
        The replication role.
        """
        return pulumi.get(self, "replication_role")

    @replication_role.setter
    def replication_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replication_role", value)

    @property
    @pulumi.getter(name="restorePointInTime")
    def restore_point_in_time(self) -> Optional[pulumi.Input[str]]:
        """
        Restore point creation time (ISO8601 format), specifying the time to restore from.
        """
        return pulumi.get(self, "restore_point_in_time")

    @restore_point_in_time.setter
    def restore_point_in_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "restore_point_in_time", value)

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the server.
        """
        return pulumi.get(self, "server_name")

    @server_name.setter
    def server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_name", value)

    @property
    @pulumi.getter
    def sku(self) -> Optional[pulumi.Input['SkuArgs']]:
        """
        The SKU (pricing tier) of the server.
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: Optional[pulumi.Input['SkuArgs']]):
        pulumi.set(self, "sku", value)

    @property
    @pulumi.getter(name="sourceServerId")
    def source_server_id(self) -> Optional[pulumi.Input[str]]:
        """
        The source MySQL server id.
        """
        return pulumi.get(self, "source_server_id")

    @source_server_id.setter
    def source_server_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_server_id", value)

    @property
    @pulumi.getter(name="sslEnforcement")
    def ssl_enforcement(self) -> Optional[pulumi.Input[Union[str, 'SslEnforcementEnum']]]:
        """
        Enable ssl enforcement or not when connect to server.
        """
        return pulumi.get(self, "ssl_enforcement")

    @ssl_enforcement.setter
    def ssl_enforcement(self, value: Optional[pulumi.Input[Union[str, 'SslEnforcementEnum']]]):
        pulumi.set(self, "ssl_enforcement", value)

    @property
    @pulumi.getter(name="storageProfile")
    def storage_profile(self) -> Optional[pulumi.Input['StorageProfileArgs']]:
        """
        Storage profile of a server.
        """
        return pulumi.get(self, "storage_profile")

    @storage_profile.setter
    def storage_profile(self, value: Optional[pulumi.Input['StorageProfileArgs']]):
        pulumi.set(self, "storage_profile", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[Union[str, 'ServerVersion']]]:
        """
        Server version.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[Union[str, 'ServerVersion']]]):
        pulumi.set(self, "version", value)


class Server(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 administrator_login: Optional[pulumi.Input[str]] = None,
                 administrator_login_password: Optional[pulumi.Input[str]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 create_mode: Optional[pulumi.Input[Union[str, 'CreateMode']]] = None,
                 delegated_subnet_arguments: Optional[pulumi.Input[Union['DelegatedSubnetArgumentsArgs', 'DelegatedSubnetArgumentsArgsDict']]] = None,
                 ha_enabled: Optional[pulumi.Input[Union[str, 'HaEnabledEnum']]] = None,
                 identity: Optional[pulumi.Input[Union['IdentityArgs', 'IdentityArgsDict']]] = None,
                 infrastructure_encryption: Optional[pulumi.Input[Union[str, 'InfrastructureEncryptionEnum']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 maintenance_window: Optional[pulumi.Input[Union['MaintenanceWindowArgs', 'MaintenanceWindowArgsDict']]] = None,
                 private_dns_zone_arguments: Optional[pulumi.Input[Union['PrivateDnsZoneArgumentsArgs', 'PrivateDnsZoneArgumentsArgsDict']]] = None,
                 replication_role: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 restore_point_in_time: Optional[pulumi.Input[str]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[Union['SkuArgs', 'SkuArgsDict']]] = None,
                 source_server_id: Optional[pulumi.Input[str]] = None,
                 ssl_enforcement: Optional[pulumi.Input[Union[str, 'SslEnforcementEnum']]] = None,
                 storage_profile: Optional[pulumi.Input[Union['StorageProfileArgs', 'StorageProfileArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 version: Optional[pulumi.Input[Union[str, 'ServerVersion']]] = None,
                 __props__=None):
        """
        Represents a server.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] administrator_login: The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
        :param pulumi.Input[str] administrator_login_password: The password of the administrator login (required for server creation).
        :param pulumi.Input[str] availability_zone: availability Zone information of the server.
        :param pulumi.Input[Union[str, 'CreateMode']] create_mode: The mode to create a new MySQL server.
        :param pulumi.Input[Union['DelegatedSubnetArgumentsArgs', 'DelegatedSubnetArgumentsArgsDict']] delegated_subnet_arguments: Delegated subnet arguments.
        :param pulumi.Input[Union[str, 'HaEnabledEnum']] ha_enabled: Enable HA or not for a server.
        :param pulumi.Input[Union['IdentityArgs', 'IdentityArgsDict']] identity: The Azure Active Directory identity of the server.
        :param pulumi.Input[Union[str, 'InfrastructureEncryptionEnum']] infrastructure_encryption: Status showing whether the server enabled infrastructure encryption.
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[Union['MaintenanceWindowArgs', 'MaintenanceWindowArgsDict']] maintenance_window: Maintenance window of a server.
        :param pulumi.Input[Union['PrivateDnsZoneArgumentsArgs', 'PrivateDnsZoneArgumentsArgsDict']] private_dns_zone_arguments: private dns zone arguments.
        :param pulumi.Input[str] replication_role: The replication role.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] restore_point_in_time: Restore point creation time (ISO8601 format), specifying the time to restore from.
        :param pulumi.Input[str] server_name: The name of the server.
        :param pulumi.Input[Union['SkuArgs', 'SkuArgsDict']] sku: The SKU (pricing tier) of the server.
        :param pulumi.Input[str] source_server_id: The source MySQL server id.
        :param pulumi.Input[Union[str, 'SslEnforcementEnum']] ssl_enforcement: Enable ssl enforcement or not when connect to server.
        :param pulumi.Input[Union['StorageProfileArgs', 'StorageProfileArgsDict']] storage_profile: Storage profile of a server.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[Union[str, 'ServerVersion']] version: Server version.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a server.

        :param str resource_name: The name of the resource.
        :param ServerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 administrator_login: Optional[pulumi.Input[str]] = None,
                 administrator_login_password: Optional[pulumi.Input[str]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 create_mode: Optional[pulumi.Input[Union[str, 'CreateMode']]] = None,
                 delegated_subnet_arguments: Optional[pulumi.Input[Union['DelegatedSubnetArgumentsArgs', 'DelegatedSubnetArgumentsArgsDict']]] = None,
                 ha_enabled: Optional[pulumi.Input[Union[str, 'HaEnabledEnum']]] = None,
                 identity: Optional[pulumi.Input[Union['IdentityArgs', 'IdentityArgsDict']]] = None,
                 infrastructure_encryption: Optional[pulumi.Input[Union[str, 'InfrastructureEncryptionEnum']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 maintenance_window: Optional[pulumi.Input[Union['MaintenanceWindowArgs', 'MaintenanceWindowArgsDict']]] = None,
                 private_dns_zone_arguments: Optional[pulumi.Input[Union['PrivateDnsZoneArgumentsArgs', 'PrivateDnsZoneArgumentsArgsDict']]] = None,
                 replication_role: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 restore_point_in_time: Optional[pulumi.Input[str]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[Union['SkuArgs', 'SkuArgsDict']]] = None,
                 source_server_id: Optional[pulumi.Input[str]] = None,
                 ssl_enforcement: Optional[pulumi.Input[Union[str, 'SslEnforcementEnum']]] = None,
                 storage_profile: Optional[pulumi.Input[Union['StorageProfileArgs', 'StorageProfileArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 version: Optional[pulumi.Input[Union[str, 'ServerVersion']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServerArgs.__new__(ServerArgs)

            __props__.__dict__["administrator_login"] = administrator_login
            __props__.__dict__["administrator_login_password"] = administrator_login_password
            __props__.__dict__["availability_zone"] = availability_zone
            __props__.__dict__["create_mode"] = create_mode
            __props__.__dict__["delegated_subnet_arguments"] = delegated_subnet_arguments
            __props__.__dict__["ha_enabled"] = ha_enabled
            __props__.__dict__["identity"] = identity
            __props__.__dict__["infrastructure_encryption"] = infrastructure_encryption
            __props__.__dict__["location"] = location
            __props__.__dict__["maintenance_window"] = maintenance_window
            __props__.__dict__["private_dns_zone_arguments"] = private_dns_zone_arguments
            __props__.__dict__["replication_role"] = replication_role
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["restore_point_in_time"] = restore_point_in_time
            __props__.__dict__["server_name"] = server_name
            __props__.__dict__["sku"] = sku
            __props__.__dict__["source_server_id"] = source_server_id
            __props__.__dict__["ssl_enforcement"] = ssl_enforcement
            __props__.__dict__["storage_profile"] = storage_profile
            __props__.__dict__["tags"] = tags
            __props__.__dict__["version"] = version
            __props__.__dict__["byok_enforcement"] = None
            __props__.__dict__["earliest_restore_date"] = None
            __props__.__dict__["fully_qualified_domain_name"] = None
            __props__.__dict__["ha_state"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["public_network_access"] = None
            __props__.__dict__["replica_capacity"] = None
            __props__.__dict__["standby_availability_zone"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:dbformysql/v20200701privatepreview:Server"), pulumi.Alias(type_="azure-native:dbformysql/v20210501:Server"), pulumi.Alias(type_="azure-native:dbformysql/v20210501preview:Server"), pulumi.Alias(type_="azure-native:dbformysql/v20211201preview:Server"), pulumi.Alias(type_="azure-native:dbformysql/v20220101:Server"), pulumi.Alias(type_="azure-native:dbformysql/v20220930preview:Server"), pulumi.Alias(type_="azure-native:dbformysql/v20230601preview:Server"), pulumi.Alias(type_="azure-native:dbformysql/v20230630:Server"), pulumi.Alias(type_="azure-native:dbformysql/v20231001preview:Server"), pulumi.Alias(type_="azure-native:dbformysql/v20231201preview:Server"), pulumi.Alias(type_="azure-native:dbformysql/v20231230:Server"), pulumi.Alias(type_="azure-native:dbformysql/v20240201preview:Server"), pulumi.Alias(type_="azure-native:dbformysql/v20240601preview:Server"), pulumi.Alias(type_="azure-native:dbformysql/v20241001preview:Server"), pulumi.Alias(type_="azure-native:dbformysql:Server")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Server, __self__).__init__(
            'azure-native:dbformysql/v20200701preview:Server',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Server':
        """
        Get an existing Server resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ServerArgs.__new__(ServerArgs)

        __props__.__dict__["administrator_login"] = None
        __props__.__dict__["availability_zone"] = None
        __props__.__dict__["byok_enforcement"] = None
        __props__.__dict__["delegated_subnet_arguments"] = None
        __props__.__dict__["earliest_restore_date"] = None
        __props__.__dict__["fully_qualified_domain_name"] = None
        __props__.__dict__["ha_enabled"] = None
        __props__.__dict__["ha_state"] = None
        __props__.__dict__["identity"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["maintenance_window"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["private_dns_zone_arguments"] = None
        __props__.__dict__["public_network_access"] = None
        __props__.__dict__["replica_capacity"] = None
        __props__.__dict__["replication_role"] = None
        __props__.__dict__["sku"] = None
        __props__.__dict__["source_server_id"] = None
        __props__.__dict__["ssl_enforcement"] = None
        __props__.__dict__["standby_availability_zone"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["storage_profile"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["version"] = None
        return Server(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="administratorLogin")
    def administrator_login(self) -> pulumi.Output[Optional[str]]:
        """
        The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
        """
        return pulumi.get(self, "administrator_login")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[Optional[str]]:
        """
        availability Zone information of the server.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="byokEnforcement")
    def byok_enforcement(self) -> pulumi.Output[str]:
        """
        Status showing whether the data encryption is enabled with customer-managed keys.
        """
        return pulumi.get(self, "byok_enforcement")

    @property
    @pulumi.getter(name="delegatedSubnetArguments")
    def delegated_subnet_arguments(self) -> pulumi.Output[Optional['outputs.DelegatedSubnetArgumentsResponse']]:
        """
        Delegated subnet arguments.
        """
        return pulumi.get(self, "delegated_subnet_arguments")

    @property
    @pulumi.getter(name="earliestRestoreDate")
    def earliest_restore_date(self) -> pulumi.Output[str]:
        """
        Earliest restore point creation time (ISO8601 format)
        """
        return pulumi.get(self, "earliest_restore_date")

    @property
    @pulumi.getter(name="fullyQualifiedDomainName")
    def fully_qualified_domain_name(self) -> pulumi.Output[str]:
        """
        The fully qualified domain name of a server.
        """
        return pulumi.get(self, "fully_qualified_domain_name")

    @property
    @pulumi.getter(name="haEnabled")
    def ha_enabled(self) -> pulumi.Output[Optional[str]]:
        """
        Enable HA or not for a server.
        """
        return pulumi.get(self, "ha_enabled")

    @property
    @pulumi.getter(name="haState")
    def ha_state(self) -> pulumi.Output[str]:
        """
        The state of a HA server.
        """
        return pulumi.get(self, "ha_state")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.IdentityResponse']]:
        """
        The Azure Active Directory identity of the server.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> pulumi.Output[Optional['outputs.MaintenanceWindowResponse']]:
        """
        Maintenance window of a server.
        """
        return pulumi.get(self, "maintenance_window")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateDnsZoneArguments")
    def private_dns_zone_arguments(self) -> pulumi.Output[Optional['outputs.PrivateDnsZoneArgumentsResponse']]:
        """
        private dns zone arguments.
        """
        return pulumi.get(self, "private_dns_zone_arguments")

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> pulumi.Output[str]:
        """
        Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
        """
        return pulumi.get(self, "public_network_access")

    @property
    @pulumi.getter(name="replicaCapacity")
    def replica_capacity(self) -> pulumi.Output[int]:
        """
        The maximum number of replicas that a primary server can have.
        """
        return pulumi.get(self, "replica_capacity")

    @property
    @pulumi.getter(name="replicationRole")
    def replication_role(self) -> pulumi.Output[Optional[str]]:
        """
        The replication role.
        """
        return pulumi.get(self, "replication_role")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.SkuResponse']]:
        """
        The SKU (pricing tier) of the server.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="sourceServerId")
    def source_server_id(self) -> pulumi.Output[Optional[str]]:
        """
        The source MySQL server id.
        """
        return pulumi.get(self, "source_server_id")

    @property
    @pulumi.getter(name="sslEnforcement")
    def ssl_enforcement(self) -> pulumi.Output[Optional[str]]:
        """
        Enable ssl enforcement or not when connect to server.
        """
        return pulumi.get(self, "ssl_enforcement")

    @property
    @pulumi.getter(name="standbyAvailabilityZone")
    def standby_availability_zone(self) -> pulumi.Output[str]:
        """
        availability Zone information of the server.
        """
        return pulumi.get(self, "standby_availability_zone")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The state of a server.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="storageProfile")
    def storage_profile(self) -> pulumi.Output[Optional['outputs.StorageProfileResponse']]:
        """
        Storage profile of a server.
        """
        return pulumi.get(self, "storage_profile")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[Optional[str]]:
        """
        Server version.
        """
        return pulumi.get(self, "version")

