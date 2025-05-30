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
from ._enums import *

__all__ = [
    'ManagedServiceIdentityArgs',
    'ManagedServiceIdentityArgsDict',
    'ProviderArgs',
    'ProviderArgsDict',
    'WorkspaceResourcePropertiesArgs',
    'WorkspaceResourcePropertiesArgsDict',
]

MYPY = False

if not MYPY:
    class ManagedServiceIdentityArgsDict(TypedDict):
        """
        Managed service identity (system assigned and/or user assigned identities)
        """
        type: pulumi.Input[Union[builtins.str, 'ManagedServiceIdentityType']]
        """
        Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
        """
        user_assigned_identities: NotRequired[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]
        """
        The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
        """
elif False:
    ManagedServiceIdentityArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ManagedServiceIdentityArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[Union[builtins.str, 'ManagedServiceIdentityType']],
                 user_assigned_identities: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        Managed service identity (system assigned and/or user assigned identities)
        :param pulumi.Input[Union[builtins.str, 'ManagedServiceIdentityType']] type: Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] user_assigned_identities: The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
        """
        pulumi.set(__self__, "type", type)
        if user_assigned_identities is not None:
            pulumi.set(__self__, "user_assigned_identities", user_assigned_identities)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[Union[builtins.str, 'ManagedServiceIdentityType']]:
        """
        Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[Union[builtins.str, 'ManagedServiceIdentityType']]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userAssignedIdentities")
    def user_assigned_identities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
        """
        return pulumi.get(self, "user_assigned_identities")

    @user_assigned_identities.setter
    def user_assigned_identities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "user_assigned_identities", value)


if not MYPY:
    class ProviderArgsDict(TypedDict):
        """
        Information about a Provider. A Provider is an entity that offers Targets to run Azure Quantum Jobs.
        """
        application_name: NotRequired[pulumi.Input[builtins.str]]
        """
        The provider's marketplace application display name.
        """
        instance_uri: NotRequired[pulumi.Input[builtins.str]]
        """
        A Uri identifying the specific instance of this provider.
        """
        provider_id: NotRequired[pulumi.Input[builtins.str]]
        """
        Unique id of this provider.
        """
        provider_sku: NotRequired[pulumi.Input[builtins.str]]
        """
        The sku associated with pricing information for this provider.
        """
        provisioning_state: NotRequired[pulumi.Input[Union[builtins.str, 'ProviderStatus']]]
        """
        Provisioning status field
        """
        resource_usage_id: NotRequired[pulumi.Input[builtins.str]]
        """
        Id to track resource usage for the provider.
        """
elif False:
    ProviderArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 application_name: Optional[pulumi.Input[builtins.str]] = None,
                 instance_uri: Optional[pulumi.Input[builtins.str]] = None,
                 provider_id: Optional[pulumi.Input[builtins.str]] = None,
                 provider_sku: Optional[pulumi.Input[builtins.str]] = None,
                 provisioning_state: Optional[pulumi.Input[Union[builtins.str, 'ProviderStatus']]] = None,
                 resource_usage_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Information about a Provider. A Provider is an entity that offers Targets to run Azure Quantum Jobs.
        :param pulumi.Input[builtins.str] application_name: The provider's marketplace application display name.
        :param pulumi.Input[builtins.str] instance_uri: A Uri identifying the specific instance of this provider.
        :param pulumi.Input[builtins.str] provider_id: Unique id of this provider.
        :param pulumi.Input[builtins.str] provider_sku: The sku associated with pricing information for this provider.
        :param pulumi.Input[Union[builtins.str, 'ProviderStatus']] provisioning_state: Provisioning status field
        :param pulumi.Input[builtins.str] resource_usage_id: Id to track resource usage for the provider.
        """
        if application_name is not None:
            pulumi.set(__self__, "application_name", application_name)
        if instance_uri is not None:
            pulumi.set(__self__, "instance_uri", instance_uri)
        if provider_id is not None:
            pulumi.set(__self__, "provider_id", provider_id)
        if provider_sku is not None:
            pulumi.set(__self__, "provider_sku", provider_sku)
        if provisioning_state is not None:
            pulumi.set(__self__, "provisioning_state", provisioning_state)
        if resource_usage_id is not None:
            pulumi.set(__self__, "resource_usage_id", resource_usage_id)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The provider's marketplace application display name.
        """
        return pulumi.get(self, "application_name")

    @application_name.setter
    def application_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "application_name", value)

    @property
    @pulumi.getter(name="instanceUri")
    def instance_uri(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A Uri identifying the specific instance of this provider.
        """
        return pulumi.get(self, "instance_uri")

    @instance_uri.setter
    def instance_uri(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "instance_uri", value)

    @property
    @pulumi.getter(name="providerId")
    def provider_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Unique id of this provider.
        """
        return pulumi.get(self, "provider_id")

    @provider_id.setter
    def provider_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "provider_id", value)

    @property
    @pulumi.getter(name="providerSku")
    def provider_sku(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The sku associated with pricing information for this provider.
        """
        return pulumi.get(self, "provider_sku")

    @provider_sku.setter
    def provider_sku(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "provider_sku", value)

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[pulumi.Input[Union[builtins.str, 'ProviderStatus']]]:
        """
        Provisioning status field
        """
        return pulumi.get(self, "provisioning_state")

    @provisioning_state.setter
    def provisioning_state(self, value: Optional[pulumi.Input[Union[builtins.str, 'ProviderStatus']]]):
        pulumi.set(self, "provisioning_state", value)

    @property
    @pulumi.getter(name="resourceUsageId")
    def resource_usage_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Id to track resource usage for the provider.
        """
        return pulumi.get(self, "resource_usage_id")

    @resource_usage_id.setter
    def resource_usage_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "resource_usage_id", value)


if not MYPY:
    class WorkspaceResourcePropertiesArgsDict(TypedDict):
        """
        Properties of a Workspace
        """
        api_key_enabled: NotRequired[pulumi.Input[builtins.bool]]
        """
        Indicator of enablement of the Quantum workspace Api keys.
        """
        providers: NotRequired[pulumi.Input[Sequence[pulumi.Input['ProviderArgsDict']]]]
        """
        List of Providers selected for this Workspace
        """
        storage_account: NotRequired[pulumi.Input[builtins.str]]
        """
        ARM Resource Id of the storage account associated with this workspace.
        """
elif False:
    WorkspaceResourcePropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class WorkspaceResourcePropertiesArgs:
    def __init__(__self__, *,
                 api_key_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 providers: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderArgs']]]] = None,
                 storage_account: Optional[pulumi.Input[builtins.str]] = None):
        """
        Properties of a Workspace
        :param pulumi.Input[builtins.bool] api_key_enabled: Indicator of enablement of the Quantum workspace Api keys.
        :param pulumi.Input[Sequence[pulumi.Input['ProviderArgs']]] providers: List of Providers selected for this Workspace
        :param pulumi.Input[builtins.str] storage_account: ARM Resource Id of the storage account associated with this workspace.
        """
        if api_key_enabled is not None:
            pulumi.set(__self__, "api_key_enabled", api_key_enabled)
        if providers is not None:
            pulumi.set(__self__, "providers", providers)
        if storage_account is not None:
            pulumi.set(__self__, "storage_account", storage_account)

    @property
    @pulumi.getter(name="apiKeyEnabled")
    def api_key_enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Indicator of enablement of the Quantum workspace Api keys.
        """
        return pulumi.get(self, "api_key_enabled")

    @api_key_enabled.setter
    def api_key_enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "api_key_enabled", value)

    @property
    @pulumi.getter
    def providers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProviderArgs']]]]:
        """
        List of Providers selected for this Workspace
        """
        return pulumi.get(self, "providers")

    @providers.setter
    def providers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderArgs']]]]):
        pulumi.set(self, "providers", value)

    @property
    @pulumi.getter(name="storageAccount")
    def storage_account(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ARM Resource Id of the storage account associated with this workspace.
        """
        return pulumi.get(self, "storage_account")

    @storage_account.setter
    def storage_account(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "storage_account", value)


