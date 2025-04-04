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
from ._enums import *

__all__ = [
    'AccountMergeInfoResponse',
    'AccountPropertiesResponseAccountStatus',
    'AccountPropertiesResponseEndpoints',
    'AccountPropertiesResponseManagedResources',
    'AccountResponseSku',
    'AccountStatusResponseErrorDetails',
    'CloudConnectorsResponse',
    'CredentialsResponse',
    'ErrorModelResponse',
    'IdentityResponse',
    'IngestionStorageResponse',
    'PrivateEndpointConnectionResponse',
    'PrivateEndpointResponse',
    'PrivateLinkServiceConnectionStateResponse',
    'ProxyResourceResponseSystemData',
    'TrackedResourceResponseSystemData',
    'UserAssignedIdentityResponse',
]

@pulumi.output_type
class AccountMergeInfoResponse(dict):
    """
    The public Account Merge Info model.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accountLocation":
            suggest = "account_location"
        elif key == "accountName":
            suggest = "account_name"
        elif key == "accountResourceGroupName":
            suggest = "account_resource_group_name"
        elif key == "accountSubscriptionId":
            suggest = "account_subscription_id"
        elif key == "mergeStatus":
            suggest = "merge_status"
        elif key == "typeOfAccount":
            suggest = "type_of_account"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AccountMergeInfoResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AccountMergeInfoResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AccountMergeInfoResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 account_location: str,
                 account_name: str,
                 account_resource_group_name: str,
                 account_subscription_id: str,
                 deprovisioned: bool,
                 merge_status: str,
                 type_of_account: str):
        """
        The public Account Merge Info model.
        :param str account_location: The account location of the *other* account in the merge operation.
        :param str account_name: The account name of the *other* account in the merge operation.
        :param str account_resource_group_name: The resource group name of the *other* account in the merge operation.
        :param str account_subscription_id: The subscription id of the *other* account in the merge operation.
        :param bool deprovisioned: The deprovisioned status of the account.
               Only applicable for the secondary account.
        :param str merge_status: The status of the merge operation.
        :param str type_of_account: The account's type for the merge operation.
        """
        pulumi.set(__self__, "account_location", account_location)
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "account_resource_group_name", account_resource_group_name)
        pulumi.set(__self__, "account_subscription_id", account_subscription_id)
        pulumi.set(__self__, "deprovisioned", deprovisioned)
        pulumi.set(__self__, "merge_status", merge_status)
        pulumi.set(__self__, "type_of_account", type_of_account)

    @property
    @pulumi.getter(name="accountLocation")
    def account_location(self) -> str:
        """
        The account location of the *other* account in the merge operation.
        """
        return pulumi.get(self, "account_location")

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> str:
        """
        The account name of the *other* account in the merge operation.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="accountResourceGroupName")
    def account_resource_group_name(self) -> str:
        """
        The resource group name of the *other* account in the merge operation.
        """
        return pulumi.get(self, "account_resource_group_name")

    @property
    @pulumi.getter(name="accountSubscriptionId")
    def account_subscription_id(self) -> str:
        """
        The subscription id of the *other* account in the merge operation.
        """
        return pulumi.get(self, "account_subscription_id")

    @property
    @pulumi.getter
    def deprovisioned(self) -> bool:
        """
        The deprovisioned status of the account.
        Only applicable for the secondary account.
        """
        return pulumi.get(self, "deprovisioned")

    @property
    @pulumi.getter(name="mergeStatus")
    def merge_status(self) -> str:
        """
        The status of the merge operation.
        """
        return pulumi.get(self, "merge_status")

    @property
    @pulumi.getter(name="typeOfAccount")
    def type_of_account(self) -> str:
        """
        The account's type for the merge operation.
        """
        return pulumi.get(self, "type_of_account")


@pulumi.output_type
class AccountPropertiesResponseAccountStatus(dict):
    """
    Gets or sets the status of the account.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accountProvisioningState":
            suggest = "account_provisioning_state"
        elif key == "errorDetails":
            suggest = "error_details"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AccountPropertiesResponseAccountStatus. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AccountPropertiesResponseAccountStatus.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AccountPropertiesResponseAccountStatus.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 account_provisioning_state: str,
                 error_details: 'outputs.AccountStatusResponseErrorDetails'):
        """
        Gets or sets the status of the account.
        :param str account_provisioning_state: Gets the account status code.
        :param 'AccountStatusResponseErrorDetails' error_details: Gets the account error details.
        """
        pulumi.set(__self__, "account_provisioning_state", account_provisioning_state)
        pulumi.set(__self__, "error_details", error_details)

    @property
    @pulumi.getter(name="accountProvisioningState")
    def account_provisioning_state(self) -> str:
        """
        Gets the account status code.
        """
        return pulumi.get(self, "account_provisioning_state")

    @property
    @pulumi.getter(name="errorDetails")
    def error_details(self) -> 'outputs.AccountStatusResponseErrorDetails':
        """
        Gets the account error details.
        """
        return pulumi.get(self, "error_details")


@pulumi.output_type
class AccountPropertiesResponseEndpoints(dict):
    """
    The URIs that are the public endpoints of the account.
    """
    def __init__(__self__, *,
                 catalog: str,
                 scan: str):
        """
        The URIs that are the public endpoints of the account.
        :param str catalog: Gets the catalog endpoint.
        :param str scan: Gets the scan endpoint.
        """
        pulumi.set(__self__, "catalog", catalog)
        pulumi.set(__self__, "scan", scan)

    @property
    @pulumi.getter
    def catalog(self) -> str:
        """
        Gets the catalog endpoint.
        """
        return pulumi.get(self, "catalog")

    @property
    @pulumi.getter
    def scan(self) -> str:
        """
        Gets the scan endpoint.
        """
        return pulumi.get(self, "scan")


@pulumi.output_type
class AccountPropertiesResponseManagedResources(dict):
    """
    Gets the resource identifiers of the managed resources.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "eventHubNamespace":
            suggest = "event_hub_namespace"
        elif key == "resourceGroup":
            suggest = "resource_group"
        elif key == "storageAccount":
            suggest = "storage_account"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AccountPropertiesResponseManagedResources. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AccountPropertiesResponseManagedResources.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AccountPropertiesResponseManagedResources.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 event_hub_namespace: str,
                 resource_group: str,
                 storage_account: str):
        """
        Gets the resource identifiers of the managed resources.
        :param str event_hub_namespace: Gets the managed event hub namespace resource identifier.
        :param str resource_group: Gets the managed resource group resource identifier. This resource group will host resource dependencies for the account.
        :param str storage_account: Gets the managed storage account resource identifier.
        """
        pulumi.set(__self__, "event_hub_namespace", event_hub_namespace)
        pulumi.set(__self__, "resource_group", resource_group)
        pulumi.set(__self__, "storage_account", storage_account)

    @property
    @pulumi.getter(name="eventHubNamespace")
    def event_hub_namespace(self) -> str:
        """
        Gets the managed event hub namespace resource identifier.
        """
        return pulumi.get(self, "event_hub_namespace")

    @property
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> str:
        """
        Gets the managed resource group resource identifier. This resource group will host resource dependencies for the account.
        """
        return pulumi.get(self, "resource_group")

    @property
    @pulumi.getter(name="storageAccount")
    def storage_account(self) -> str:
        """
        Gets the managed storage account resource identifier.
        """
        return pulumi.get(self, "storage_account")


@pulumi.output_type
class AccountResponseSku(dict):
    """
    Gets or sets the Sku.
    """
    def __init__(__self__, *,
                 capacity: Optional[int] = None,
                 name: Optional[str] = None):
        """
        Gets or sets the Sku.
        :param int capacity: Gets or sets the sku capacity.
        :param str name: Gets or sets the sku name.
        """
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def capacity(self) -> Optional[int]:
        """
        Gets or sets the sku capacity.
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Gets or sets the sku name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class AccountStatusResponseErrorDetails(dict):
    """
    Gets the account error details.
    """
    def __init__(__self__, *,
                 code: str,
                 details: Sequence['outputs.ErrorModelResponse'],
                 message: str,
                 target: str):
        """
        Gets the account error details.
        :param str code: Gets or sets the code.
        :param Sequence['ErrorModelResponse'] details: Gets or sets the details.
        :param str message: Gets or sets the messages.
        :param str target: Gets or sets the target.
        """
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "details", details)
        pulumi.set(__self__, "message", message)
        pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        Gets or sets the code.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def details(self) -> Sequence['outputs.ErrorModelResponse']:
        """
        Gets or sets the details.
        """
        return pulumi.get(self, "details")

    @property
    @pulumi.getter
    def message(self) -> str:
        """
        Gets or sets the messages.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def target(self) -> str:
        """
        Gets or sets the target.
        """
        return pulumi.get(self, "target")


@pulumi.output_type
class CloudConnectorsResponse(dict):
    """
    External Cloud Service connectors
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "awsExternalId":
            suggest = "aws_external_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CloudConnectorsResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CloudConnectorsResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CloudConnectorsResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 aws_external_id: str):
        """
        External Cloud Service connectors
        :param str aws_external_id: AWS external identifier.
               Configured in AWS to allow use of the role arn used for scanning
        """
        pulumi.set(__self__, "aws_external_id", aws_external_id)

    @property
    @pulumi.getter(name="awsExternalId")
    def aws_external_id(self) -> str:
        """
        AWS external identifier.
        Configured in AWS to allow use of the role arn used for scanning
        """
        return pulumi.get(self, "aws_external_id")


@pulumi.output_type
class CredentialsResponse(dict):
    """
    Credentials to access the event streaming service attached to the purview account.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "identityId":
            suggest = "identity_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CredentialsResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CredentialsResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CredentialsResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 identity_id: Optional[str] = None,
                 type: Optional[str] = None):
        """
        Credentials to access the event streaming service attached to the purview account.
        :param str identity_id: Identity identifier for UserAssign type.
        :param str type: Identity Type.
        """
        if identity_id is not None:
            pulumi.set(__self__, "identity_id", identity_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="identityId")
    def identity_id(self) -> Optional[str]:
        """
        Identity identifier for UserAssign type.
        """
        return pulumi.get(self, "identity_id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Identity Type.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class ErrorModelResponse(dict):
    """
    Default error model
    """
    def __init__(__self__, *,
                 code: str,
                 details: Sequence['outputs.ErrorModelResponse'],
                 message: str,
                 target: str):
        """
        Default error model
        :param str code: Gets or sets the code.
        :param Sequence['ErrorModelResponse'] details: Gets or sets the details.
        :param str message: Gets or sets the messages.
        :param str target: Gets or sets the target.
        """
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "details", details)
        pulumi.set(__self__, "message", message)
        pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        Gets or sets the code.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def details(self) -> Sequence['outputs.ErrorModelResponse']:
        """
        Gets or sets the details.
        """
        return pulumi.get(self, "details")

    @property
    @pulumi.getter
    def message(self) -> str:
        """
        Gets or sets the messages.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def target(self) -> str:
        """
        Gets or sets the target.
        """
        return pulumi.get(self, "target")


@pulumi.output_type
class IdentityResponse(dict):
    """
    The Managed Identity of the resource
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "principalId":
            suggest = "principal_id"
        elif key == "tenantId":
            suggest = "tenant_id"
        elif key == "userAssignedIdentities":
            suggest = "user_assigned_identities"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IdentityResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IdentityResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IdentityResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 principal_id: str,
                 tenant_id: str,
                 type: Optional[str] = None,
                 user_assigned_identities: Optional[Mapping[str, 'outputs.UserAssignedIdentityResponse']] = None):
        """
        The Managed Identity of the resource
        :param str principal_id: Service principal object Id
        :param str tenant_id: Tenant Id
        :param str type: Identity Type
        :param Mapping[str, 'UserAssignedIdentityResponse'] user_assigned_identities: User Assigned Identities
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user_assigned_identities is not None:
            pulumi.set(__self__, "user_assigned_identities", user_assigned_identities)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        Service principal object Id
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        Tenant Id
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Identity Type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userAssignedIdentities")
    def user_assigned_identities(self) -> Optional[Mapping[str, 'outputs.UserAssignedIdentityResponse']]:
        """
        User Assigned Identities
        """
        return pulumi.get(self, "user_assigned_identities")


@pulumi.output_type
class IngestionStorageResponse(dict):
    """
    Ingestion Storage Account Info
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "primaryEndpoint":
            suggest = "primary_endpoint"
        elif key == "publicNetworkAccess":
            suggest = "public_network_access"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IngestionStorageResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IngestionStorageResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IngestionStorageResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 id: str,
                 primary_endpoint: str,
                 public_network_access: Optional[str] = None):
        """
        Ingestion Storage Account Info
        :param str id: Gets or sets the Id.
        :param str primary_endpoint: Gets or sets the primary endpoint.
        :param str public_network_access: Gets or sets the public network access setting
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "primary_endpoint", primary_endpoint)
        if public_network_access is not None:
            pulumi.set(__self__, "public_network_access", public_network_access)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Gets or sets the Id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="primaryEndpoint")
    def primary_endpoint(self) -> str:
        """
        Gets or sets the primary endpoint.
        """
        return pulumi.get(self, "primary_endpoint")

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> Optional[str]:
        """
        Gets or sets the public network access setting
        """
        return pulumi.get(self, "public_network_access")


@pulumi.output_type
class PrivateEndpointConnectionResponse(dict):
    """
    A private endpoint connection class.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "provisioningState":
            suggest = "provisioning_state"
        elif key == "systemData":
            suggest = "system_data"
        elif key == "privateEndpoint":
            suggest = "private_endpoint"
        elif key == "privateLinkServiceConnectionState":
            suggest = "private_link_service_connection_state"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PrivateEndpointConnectionResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PrivateEndpointConnectionResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PrivateEndpointConnectionResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 id: str,
                 name: str,
                 provisioning_state: str,
                 system_data: 'outputs.ProxyResourceResponseSystemData',
                 type: str,
                 private_endpoint: Optional['outputs.PrivateEndpointResponse'] = None,
                 private_link_service_connection_state: Optional['outputs.PrivateLinkServiceConnectionStateResponse'] = None):
        """
        A private endpoint connection class.
        :param str id: Gets or sets the identifier.
        :param str name: Gets or sets the name.
        :param str provisioning_state: The provisioning state.
        :param 'ProxyResourceResponseSystemData' system_data: Metadata pertaining to creation and last modification of the resource.
        :param str type: Gets or sets the type.
        :param 'PrivateEndpointResponse' private_endpoint: The private endpoint information.
        :param 'PrivateLinkServiceConnectionStateResponse' private_link_service_connection_state: The private link service connection state.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        pulumi.set(__self__, "system_data", system_data)
        pulumi.set(__self__, "type", type)
        if private_endpoint is not None:
            pulumi.set(__self__, "private_endpoint", private_endpoint)
        if private_link_service_connection_state is not None:
            pulumi.set(__self__, "private_link_service_connection_state", private_link_service_connection_state)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Gets or sets the identifier.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Gets or sets the name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.ProxyResourceResponseSystemData':
        """
        Metadata pertaining to creation and last modification of the resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Gets or sets the type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="privateEndpoint")
    def private_endpoint(self) -> Optional['outputs.PrivateEndpointResponse']:
        """
        The private endpoint information.
        """
        return pulumi.get(self, "private_endpoint")

    @property
    @pulumi.getter(name="privateLinkServiceConnectionState")
    def private_link_service_connection_state(self) -> Optional['outputs.PrivateLinkServiceConnectionStateResponse']:
        """
        The private link service connection state.
        """
        return pulumi.get(self, "private_link_service_connection_state")


@pulumi.output_type
class PrivateEndpointResponse(dict):
    """
    A private endpoint class.
    """
    def __init__(__self__, *,
                 id: Optional[str] = None):
        """
        A private endpoint class.
        :param str id: The private endpoint identifier.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The private endpoint identifier.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class PrivateLinkServiceConnectionStateResponse(dict):
    """
    The private link service connection state.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "actionsRequired":
            suggest = "actions_required"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PrivateLinkServiceConnectionStateResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PrivateLinkServiceConnectionStateResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PrivateLinkServiceConnectionStateResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 actions_required: Optional[str] = None,
                 description: Optional[str] = None,
                 status: Optional[str] = None):
        """
        The private link service connection state.
        :param str actions_required: The required actions.
        :param str description: The description.
        :param str status: The status.
        """
        if actions_required is not None:
            pulumi.set(__self__, "actions_required", actions_required)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="actionsRequired")
    def actions_required(self) -> Optional[str]:
        """
        The required actions.
        """
        return pulumi.get(self, "actions_required")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class ProxyResourceResponseSystemData(dict):
    """
    Metadata pertaining to creation and last modification of the resource.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "createdBy":
            suggest = "created_by"
        elif key == "createdByType":
            suggest = "created_by_type"
        elif key == "lastModifiedAt":
            suggest = "last_modified_at"
        elif key == "lastModifiedBy":
            suggest = "last_modified_by"
        elif key == "lastModifiedByType":
            suggest = "last_modified_by_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProxyResourceResponseSystemData. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProxyResourceResponseSystemData.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProxyResourceResponseSystemData.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: str,
                 created_by: str,
                 created_by_type: str,
                 last_modified_at: str,
                 last_modified_by: str,
                 last_modified_by_type: str):
        """
        Metadata pertaining to creation and last modification of the resource.
        :param str created_at: The timestamp of resource creation (UTC).
        :param str created_by: The identity that created the resource.
        :param str created_by_type: The type of identity that created the resource.
        :param str last_modified_at: The timestamp of the last modification the resource (UTC).
        :param str last_modified_by: The identity that last modified the resource.
        :param str last_modified_by_type: The type of identity that last modified the resource.
        """
        pulumi.set(__self__, "created_at", created_at)
        pulumi.set(__self__, "created_by", created_by)
        pulumi.set(__self__, "created_by_type", created_by_type)
        pulumi.set(__self__, "last_modified_at", last_modified_at)
        pulumi.set(__self__, "last_modified_by", last_modified_by)
        pulumi.set(__self__, "last_modified_by_type", last_modified_by_type)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The timestamp of resource creation (UTC).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> str:
        """
        The identity that created the resource.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="createdByType")
    def created_by_type(self) -> str:
        """
        The type of identity that created the resource.
        """
        return pulumi.get(self, "created_by_type")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> str:
        """
        The timestamp of the last modification the resource (UTC).
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter(name="lastModifiedBy")
    def last_modified_by(self) -> str:
        """
        The identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by")

    @property
    @pulumi.getter(name="lastModifiedByType")
    def last_modified_by_type(self) -> str:
        """
        The type of identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by_type")


@pulumi.output_type
class TrackedResourceResponseSystemData(dict):
    """
    Metadata pertaining to creation and last modification of the resource.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "createdBy":
            suggest = "created_by"
        elif key == "createdByType":
            suggest = "created_by_type"
        elif key == "lastModifiedAt":
            suggest = "last_modified_at"
        elif key == "lastModifiedBy":
            suggest = "last_modified_by"
        elif key == "lastModifiedByType":
            suggest = "last_modified_by_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TrackedResourceResponseSystemData. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TrackedResourceResponseSystemData.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TrackedResourceResponseSystemData.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: str,
                 created_by: str,
                 created_by_type: str,
                 last_modified_at: str,
                 last_modified_by: str,
                 last_modified_by_type: str):
        """
        Metadata pertaining to creation and last modification of the resource.
        :param str created_at: The timestamp of resource creation (UTC).
        :param str created_by: The identity that created the resource.
        :param str created_by_type: The type of identity that created the resource.
        :param str last_modified_at: The timestamp of the last modification the resource (UTC).
        :param str last_modified_by: The identity that last modified the resource.
        :param str last_modified_by_type: The type of identity that last modified the resource.
        """
        pulumi.set(__self__, "created_at", created_at)
        pulumi.set(__self__, "created_by", created_by)
        pulumi.set(__self__, "created_by_type", created_by_type)
        pulumi.set(__self__, "last_modified_at", last_modified_at)
        pulumi.set(__self__, "last_modified_by", last_modified_by)
        pulumi.set(__self__, "last_modified_by_type", last_modified_by_type)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The timestamp of resource creation (UTC).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> str:
        """
        The identity that created the resource.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="createdByType")
    def created_by_type(self) -> str:
        """
        The type of identity that created the resource.
        """
        return pulumi.get(self, "created_by_type")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> str:
        """
        The timestamp of the last modification the resource (UTC).
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter(name="lastModifiedBy")
    def last_modified_by(self) -> str:
        """
        The identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by")

    @property
    @pulumi.getter(name="lastModifiedByType")
    def last_modified_by_type(self) -> str:
        """
        The type of identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by_type")


@pulumi.output_type
class UserAssignedIdentityResponse(dict):
    """
    Uses client ID and Principal ID
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clientId":
            suggest = "client_id"
        elif key == "principalId":
            suggest = "principal_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in UserAssignedIdentityResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        UserAssignedIdentityResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        UserAssignedIdentityResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 client_id: str,
                 principal_id: str):
        """
        Uses client ID and Principal ID
        :param str client_id: Gets or Sets Client ID
        :param str principal_id: Gets or Sets Principal ID
        """
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "principal_id", principal_id)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        Gets or Sets Client ID
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        Gets or Sets Principal ID
        """
        return pulumi.get(self, "principal_id")


