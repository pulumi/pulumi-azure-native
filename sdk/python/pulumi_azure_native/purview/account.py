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
from ._enums import *
from ._inputs import *

__all__ = ['AccountArgs', 'Account']

@pulumi.input_type
class AccountArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 account_name: Optional[pulumi.Input[builtins.str]] = None,
                 identity: Optional[pulumi.Input['IdentityArgs']] = None,
                 ingestion_storage: Optional[pulumi.Input['IngestionStorageArgs']] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 managed_event_hub_state: Optional[pulumi.Input[Union[builtins.str, 'ManagedEventHubState']]] = None,
                 managed_resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 managed_resources_public_network_access: Optional[pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']]] = None,
                 public_network_access: Optional[pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']]] = None,
                 sku: Optional[pulumi.Input['AccountSkuArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 tenant_endpoint_state: Optional[pulumi.Input[Union[builtins.str, 'TenantEndpointState']]] = None):
        """
        The set of arguments for constructing a Account resource.
        :param pulumi.Input[builtins.str] resource_group_name: The resource group name.
        :param pulumi.Input[builtins.str] account_name: The name of the account.
        :param pulumi.Input['IdentityArgs'] identity: The Managed Identity of the resource
        :param pulumi.Input['IngestionStorageArgs'] ingestion_storage: Ingestion Storage Account Info
        :param pulumi.Input[builtins.str] location: Gets or sets the location.
        :param pulumi.Input[Union[builtins.str, 'ManagedEventHubState']] managed_event_hub_state: Gets or sets the state of managed eventhub. If enabled managed eventhub will be created, if disabled the managed eventhub will be removed.
        :param pulumi.Input[builtins.str] managed_resource_group_name: Gets or sets the managed resource group name
        :param pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']] managed_resources_public_network_access: Gets or sets the public network access for managed resources.
        :param pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']] public_network_access: Gets or sets the public network access.
        :param pulumi.Input['AccountSkuArgs'] sku: Gets or sets the Sku.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Tags on the azure resource.
        :param pulumi.Input[Union[builtins.str, 'TenantEndpointState']] tenant_endpoint_state: Gets or sets the state of tenant endpoint.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if account_name is not None:
            pulumi.set(__self__, "account_name", account_name)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if ingestion_storage is not None:
            pulumi.set(__self__, "ingestion_storage", ingestion_storage)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if managed_event_hub_state is None:
            managed_event_hub_state = 'NotSpecified'
        if managed_event_hub_state is not None:
            pulumi.set(__self__, "managed_event_hub_state", managed_event_hub_state)
        if managed_resource_group_name is not None:
            pulumi.set(__self__, "managed_resource_group_name", managed_resource_group_name)
        if managed_resources_public_network_access is None:
            managed_resources_public_network_access = 'NotSpecified'
        if managed_resources_public_network_access is not None:
            pulumi.set(__self__, "managed_resources_public_network_access", managed_resources_public_network_access)
        if public_network_access is None:
            public_network_access = 'Enabled'
        if public_network_access is not None:
            pulumi.set(__self__, "public_network_access", public_network_access)
        if sku is not None:
            pulumi.set(__self__, "sku", sku)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tenant_endpoint_state is not None:
            pulumi.set(__self__, "tenant_endpoint_state", tenant_endpoint_state)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The resource group name.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the account.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['IdentityArgs']]:
        """
        The Managed Identity of the resource
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['IdentityArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter(name="ingestionStorage")
    def ingestion_storage(self) -> Optional[pulumi.Input['IngestionStorageArgs']]:
        """
        Ingestion Storage Account Info
        """
        return pulumi.get(self, "ingestion_storage")

    @ingestion_storage.setter
    def ingestion_storage(self, value: Optional[pulumi.Input['IngestionStorageArgs']]):
        pulumi.set(self, "ingestion_storage", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Gets or sets the location.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="managedEventHubState")
    def managed_event_hub_state(self) -> Optional[pulumi.Input[Union[builtins.str, 'ManagedEventHubState']]]:
        """
        Gets or sets the state of managed eventhub. If enabled managed eventhub will be created, if disabled the managed eventhub will be removed.
        """
        return pulumi.get(self, "managed_event_hub_state")

    @managed_event_hub_state.setter
    def managed_event_hub_state(self, value: Optional[pulumi.Input[Union[builtins.str, 'ManagedEventHubState']]]):
        pulumi.set(self, "managed_event_hub_state", value)

    @property
    @pulumi.getter(name="managedResourceGroupName")
    def managed_resource_group_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Gets or sets the managed resource group name
        """
        return pulumi.get(self, "managed_resource_group_name")

    @managed_resource_group_name.setter
    def managed_resource_group_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "managed_resource_group_name", value)

    @property
    @pulumi.getter(name="managedResourcesPublicNetworkAccess")
    def managed_resources_public_network_access(self) -> Optional[pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']]]:
        """
        Gets or sets the public network access for managed resources.
        """
        return pulumi.get(self, "managed_resources_public_network_access")

    @managed_resources_public_network_access.setter
    def managed_resources_public_network_access(self, value: Optional[pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']]]):
        pulumi.set(self, "managed_resources_public_network_access", value)

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> Optional[pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']]]:
        """
        Gets or sets the public network access.
        """
        return pulumi.get(self, "public_network_access")

    @public_network_access.setter
    def public_network_access(self, value: Optional[pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']]]):
        pulumi.set(self, "public_network_access", value)

    @property
    @pulumi.getter
    def sku(self) -> Optional[pulumi.Input['AccountSkuArgs']]:
        """
        Gets or sets the Sku.
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: Optional[pulumi.Input['AccountSkuArgs']]):
        pulumi.set(self, "sku", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Tags on the azure resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tenantEndpointState")
    def tenant_endpoint_state(self) -> Optional[pulumi.Input[Union[builtins.str, 'TenantEndpointState']]]:
        """
        Gets or sets the state of tenant endpoint.
        """
        return pulumi.get(self, "tenant_endpoint_state")

    @tenant_endpoint_state.setter
    def tenant_endpoint_state(self, value: Optional[pulumi.Input[Union[builtins.str, 'TenantEndpointState']]]):
        pulumi.set(self, "tenant_endpoint_state", value)


@pulumi.type_token("azure-native:purview:Account")
class Account(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[builtins.str]] = None,
                 identity: Optional[pulumi.Input[Union['IdentityArgs', 'IdentityArgsDict']]] = None,
                 ingestion_storage: Optional[pulumi.Input[Union['IngestionStorageArgs', 'IngestionStorageArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 managed_event_hub_state: Optional[pulumi.Input[Union[builtins.str, 'ManagedEventHubState']]] = None,
                 managed_resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 managed_resources_public_network_access: Optional[pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']]] = None,
                 public_network_access: Optional[pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input[Union['AccountSkuArgs', 'AccountSkuArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 tenant_endpoint_state: Optional[pulumi.Input[Union[builtins.str, 'TenantEndpointState']]] = None,
                 __props__=None):
        """
        Account resource

        Uses Azure REST API version 2024-04-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-12-01.

        Other available API versions: 2021-12-01, 2023-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native purview [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] account_name: The name of the account.
        :param pulumi.Input[Union['IdentityArgs', 'IdentityArgsDict']] identity: The Managed Identity of the resource
        :param pulumi.Input[Union['IngestionStorageArgs', 'IngestionStorageArgsDict']] ingestion_storage: Ingestion Storage Account Info
        :param pulumi.Input[builtins.str] location: Gets or sets the location.
        :param pulumi.Input[Union[builtins.str, 'ManagedEventHubState']] managed_event_hub_state: Gets or sets the state of managed eventhub. If enabled managed eventhub will be created, if disabled the managed eventhub will be removed.
        :param pulumi.Input[builtins.str] managed_resource_group_name: Gets or sets the managed resource group name
        :param pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']] managed_resources_public_network_access: Gets or sets the public network access for managed resources.
        :param pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']] public_network_access: Gets or sets the public network access.
        :param pulumi.Input[builtins.str] resource_group_name: The resource group name.
        :param pulumi.Input[Union['AccountSkuArgs', 'AccountSkuArgsDict']] sku: Gets or sets the Sku.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Tags on the azure resource.
        :param pulumi.Input[Union[builtins.str, 'TenantEndpointState']] tenant_endpoint_state: Gets or sets the state of tenant endpoint.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Account resource

        Uses Azure REST API version 2024-04-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-12-01.

        Other available API versions: 2021-12-01, 2023-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native purview [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param AccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[builtins.str]] = None,
                 identity: Optional[pulumi.Input[Union['IdentityArgs', 'IdentityArgsDict']]] = None,
                 ingestion_storage: Optional[pulumi.Input[Union['IngestionStorageArgs', 'IngestionStorageArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 managed_event_hub_state: Optional[pulumi.Input[Union[builtins.str, 'ManagedEventHubState']]] = None,
                 managed_resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 managed_resources_public_network_access: Optional[pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']]] = None,
                 public_network_access: Optional[pulumi.Input[Union[builtins.str, 'PublicNetworkAccess']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input[Union['AccountSkuArgs', 'AccountSkuArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 tenant_endpoint_state: Optional[pulumi.Input[Union[builtins.str, 'TenantEndpointState']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccountArgs.__new__(AccountArgs)

            __props__.__dict__["account_name"] = account_name
            __props__.__dict__["identity"] = identity
            __props__.__dict__["ingestion_storage"] = ingestion_storage
            __props__.__dict__["location"] = location
            if managed_event_hub_state is None:
                managed_event_hub_state = 'NotSpecified'
            __props__.__dict__["managed_event_hub_state"] = managed_event_hub_state
            __props__.__dict__["managed_resource_group_name"] = managed_resource_group_name
            if managed_resources_public_network_access is None:
                managed_resources_public_network_access = 'NotSpecified'
            __props__.__dict__["managed_resources_public_network_access"] = managed_resources_public_network_access
            if public_network_access is None:
                public_network_access = 'Enabled'
            __props__.__dict__["public_network_access"] = public_network_access
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["sku"] = sku
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tenant_endpoint_state"] = tenant_endpoint_state
            __props__.__dict__["account_status"] = None
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["cloud_connectors"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["created_by"] = None
            __props__.__dict__["created_by_object_id"] = None
            __props__.__dict__["default_domain"] = None
            __props__.__dict__["endpoints"] = None
            __props__.__dict__["friendly_name"] = None
            __props__.__dict__["managed_resources"] = None
            __props__.__dict__["merge_info"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["private_endpoint_connections"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:purview/v20201201preview:Account"), pulumi.Alias(type_="azure-native:purview/v20210701:Account"), pulumi.Alias(type_="azure-native:purview/v20211201:Account"), pulumi.Alias(type_="azure-native:purview/v20230501preview:Account"), pulumi.Alias(type_="azure-native:purview/v20240401preview:Account")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Account, __self__).__init__(
            'azure-native:purview:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Account':
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AccountArgs.__new__(AccountArgs)

        __props__.__dict__["account_status"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["cloud_connectors"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["created_by"] = None
        __props__.__dict__["created_by_object_id"] = None
        __props__.__dict__["default_domain"] = None
        __props__.__dict__["endpoints"] = None
        __props__.__dict__["friendly_name"] = None
        __props__.__dict__["identity"] = None
        __props__.__dict__["ingestion_storage"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["managed_event_hub_state"] = None
        __props__.__dict__["managed_resource_group_name"] = None
        __props__.__dict__["managed_resources"] = None
        __props__.__dict__["managed_resources_public_network_access"] = None
        __props__.__dict__["merge_info"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["private_endpoint_connections"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["public_network_access"] = None
        __props__.__dict__["sku"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["tenant_endpoint_state"] = None
        __props__.__dict__["type"] = None
        return Account(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountStatus")
    def account_status(self) -> pulumi.Output['outputs.AccountPropertiesResponseAccountStatus']:
        """
        Gets or sets the status of the account.
        """
        return pulumi.get(self, "account_status")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="cloudConnectors")
    def cloud_connectors(self) -> pulumi.Output[Optional['outputs.CloudConnectorsResponse']]:
        """
        External Cloud Service connectors
        """
        return pulumi.get(self, "cloud_connectors")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        Gets the time at which the entity was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> pulumi.Output[builtins.str]:
        """
        Gets the creator of the entity.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="createdByObjectId")
    def created_by_object_id(self) -> pulumi.Output[builtins.str]:
        """
        Gets the creators of the entity's object id.
        """
        return pulumi.get(self, "created_by_object_id")

    @property
    @pulumi.getter(name="defaultDomain")
    def default_domain(self) -> pulumi.Output[builtins.str]:
        """
        Gets the default domain in the account.
        """
        return pulumi.get(self, "default_domain")

    @property
    @pulumi.getter
    def endpoints(self) -> pulumi.Output['outputs.AccountPropertiesResponseEndpoints']:
        """
        The URIs that are the public endpoints of the account.
        """
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> pulumi.Output[builtins.str]:
        """
        Gets or sets the friendly name.
        """
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.IdentityResponse']]:
        """
        The Managed Identity of the resource
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="ingestionStorage")
    def ingestion_storage(self) -> pulumi.Output[Optional['outputs.IngestionStorageResponse']]:
        """
        Ingestion Storage Account Info
        """
        return pulumi.get(self, "ingestion_storage")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Gets or sets the location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedEventHubState")
    def managed_event_hub_state(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Gets or sets the state of managed eventhub. If enabled managed eventhub will be created, if disabled the managed eventhub will be removed.
        """
        return pulumi.get(self, "managed_event_hub_state")

    @property
    @pulumi.getter(name="managedResourceGroupName")
    def managed_resource_group_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Gets or sets the managed resource group name
        """
        return pulumi.get(self, "managed_resource_group_name")

    @property
    @pulumi.getter(name="managedResources")
    def managed_resources(self) -> pulumi.Output['outputs.AccountPropertiesResponseManagedResources']:
        """
        Gets the resource identifiers of the managed resources.
        """
        return pulumi.get(self, "managed_resources")

    @property
    @pulumi.getter(name="managedResourcesPublicNetworkAccess")
    def managed_resources_public_network_access(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Gets or sets the public network access for managed resources.
        """
        return pulumi.get(self, "managed_resources_public_network_access")

    @property
    @pulumi.getter(name="mergeInfo")
    def merge_info(self) -> pulumi.Output[Optional['outputs.AccountMergeInfoResponse']]:
        """
        Gets or sets the Merge Info.
        """
        return pulumi.get(self, "merge_info")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Gets or sets the name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateEndpointConnections")
    def private_endpoint_connections(self) -> pulumi.Output[Sequence['outputs.PrivateEndpointConnectionResponse']]:
        """
        Gets the private endpoint connections information.
        """
        return pulumi.get(self, "private_endpoint_connections")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Gets or sets the state of the provisioning.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Gets or sets the public network access.
        """
        return pulumi.get(self, "public_network_access")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.AccountResponseSku']]:
        """
        Gets or sets the Sku.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.TrackedResourceResponseSystemData']:
        """
        Metadata pertaining to creation and last modification of the resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Tags on the azure resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantEndpointState")
    def tenant_endpoint_state(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Gets or sets the state of tenant endpoint.
        """
        return pulumi.get(self, "tenant_endpoint_state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Gets or sets the type.
        """
        return pulumi.get(self, "type")

