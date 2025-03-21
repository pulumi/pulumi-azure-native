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

__all__ = [
    'AzureBareMetalStorageInstanceIdentityResponse',
    'StorageBillingPropertiesResponse',
    'StoragePropertiesResponse',
    'SystemDataResponse',
]

@pulumi.output_type
class AzureBareMetalStorageInstanceIdentityResponse(dict):
    """
    Identity for Azure Bare Metal Storage Instance.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "principalId":
            suggest = "principal_id"
        elif key == "tenantId":
            suggest = "tenant_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AzureBareMetalStorageInstanceIdentityResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AzureBareMetalStorageInstanceIdentityResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AzureBareMetalStorageInstanceIdentityResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 principal_id: str,
                 tenant_id: str,
                 type: Optional[str] = None):
        """
        Identity for Azure Bare Metal Storage Instance.
        :param str principal_id: The principal ID of Azure Bare Metal Storage Instance identity. This property will only be provided for a system assigned identity.
        :param str tenant_id: The tenant ID associated with the Azure Bare Metal Storage Instance. This property will only be provided for a system assigned identity.
        :param str type: The type of identity used for the Azure Bare Metal Storage Instance. The type 'SystemAssigned' refers to an implicitly created identity. The type 'None' will remove any identities from the Azure Bare Metal Storage Instance.
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The principal ID of Azure Bare Metal Storage Instance identity. This property will only be provided for a system assigned identity.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The tenant ID associated with the Azure Bare Metal Storage Instance. This property will only be provided for a system assigned identity.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of identity used for the Azure Bare Metal Storage Instance. The type 'SystemAssigned' refers to an implicitly created identity. The type 'None' will remove any identities from the Azure Bare Metal Storage Instance.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class StorageBillingPropertiesResponse(dict):
    """
    Describes the billing related details of the AzureBareMetalStorageInstance.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "azureBareMetalStorageInstanceSize":
            suggest = "azure_bare_metal_storage_instance_size"
        elif key == "billingMode":
            suggest = "billing_mode"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StorageBillingPropertiesResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StorageBillingPropertiesResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StorageBillingPropertiesResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 azure_bare_metal_storage_instance_size: Optional[str] = None,
                 billing_mode: Optional[str] = None):
        """
        Describes the billing related details of the AzureBareMetalStorageInstance.
        :param str azure_bare_metal_storage_instance_size: the SKU type that is provisioned
        :param str billing_mode: the billing mode for the storage instance
        """
        if azure_bare_metal_storage_instance_size is not None:
            pulumi.set(__self__, "azure_bare_metal_storage_instance_size", azure_bare_metal_storage_instance_size)
        if billing_mode is not None:
            pulumi.set(__self__, "billing_mode", billing_mode)

    @property
    @pulumi.getter(name="azureBareMetalStorageInstanceSize")
    def azure_bare_metal_storage_instance_size(self) -> Optional[str]:
        """
        the SKU type that is provisioned
        """
        return pulumi.get(self, "azure_bare_metal_storage_instance_size")

    @property
    @pulumi.getter(name="billingMode")
    def billing_mode(self) -> Optional[str]:
        """
        the billing mode for the storage instance
        """
        return pulumi.get(self, "billing_mode")


@pulumi.output_type
class StoragePropertiesResponse(dict):
    """
    described the storage properties of the azure bare metal storage instance
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "hardwareType":
            suggest = "hardware_type"
        elif key == "offeringType":
            suggest = "offering_type"
        elif key == "provisioningState":
            suggest = "provisioning_state"
        elif key == "storageBillingProperties":
            suggest = "storage_billing_properties"
        elif key == "storageType":
            suggest = "storage_type"
        elif key == "workloadType":
            suggest = "workload_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StoragePropertiesResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StoragePropertiesResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StoragePropertiesResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 generation: Optional[str] = None,
                 hardware_type: Optional[str] = None,
                 offering_type: Optional[str] = None,
                 provisioning_state: Optional[str] = None,
                 storage_billing_properties: Optional['outputs.StorageBillingPropertiesResponse'] = None,
                 storage_type: Optional[str] = None,
                 workload_type: Optional[str] = None):
        """
        described the storage properties of the azure bare metal storage instance
        :param str generation: the kind of storage instance
        :param str hardware_type: the hardware type of the storage instance
        :param str offering_type: the offering type for which the resource is getting provisioned
        :param str provisioning_state: State of provisioning of the AzureBareMetalStorageInstance
        :param 'StorageBillingPropertiesResponse' storage_billing_properties: the billing related information for the resource
        :param str storage_type: the storage protocol for which the resource is getting provisioned
        :param str workload_type: the workload for which the resource is getting provisioned
        """
        if generation is not None:
            pulumi.set(__self__, "generation", generation)
        if hardware_type is not None:
            pulumi.set(__self__, "hardware_type", hardware_type)
        if offering_type is not None:
            pulumi.set(__self__, "offering_type", offering_type)
        if provisioning_state is not None:
            pulumi.set(__self__, "provisioning_state", provisioning_state)
        if storage_billing_properties is not None:
            pulumi.set(__self__, "storage_billing_properties", storage_billing_properties)
        if storage_type is not None:
            pulumi.set(__self__, "storage_type", storage_type)
        if workload_type is not None:
            pulumi.set(__self__, "workload_type", workload_type)

    @property
    @pulumi.getter
    def generation(self) -> Optional[str]:
        """
        the kind of storage instance
        """
        return pulumi.get(self, "generation")

    @property
    @pulumi.getter(name="hardwareType")
    def hardware_type(self) -> Optional[str]:
        """
        the hardware type of the storage instance
        """
        return pulumi.get(self, "hardware_type")

    @property
    @pulumi.getter(name="offeringType")
    def offering_type(self) -> Optional[str]:
        """
        the offering type for which the resource is getting provisioned
        """
        return pulumi.get(self, "offering_type")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        State of provisioning of the AzureBareMetalStorageInstance
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="storageBillingProperties")
    def storage_billing_properties(self) -> Optional['outputs.StorageBillingPropertiesResponse']:
        """
        the billing related information for the resource
        """
        return pulumi.get(self, "storage_billing_properties")

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> Optional[str]:
        """
        the storage protocol for which the resource is getting provisioned
        """
        return pulumi.get(self, "storage_type")

    @property
    @pulumi.getter(name="workloadType")
    def workload_type(self) -> Optional[str]:
        """
        the workload for which the resource is getting provisioned
        """
        return pulumi.get(self, "workload_type")


@pulumi.output_type
class SystemDataResponse(dict):
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
            pulumi.log.warn(f"Key '{key}' not found in SystemDataResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SystemDataResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SystemDataResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: Optional[str] = None,
                 created_by: Optional[str] = None,
                 created_by_type: Optional[str] = None,
                 last_modified_at: Optional[str] = None,
                 last_modified_by: Optional[str] = None,
                 last_modified_by_type: Optional[str] = None):
        """
        Metadata pertaining to creation and last modification of the resource.
        :param str created_at: The timestamp of resource creation (UTC).
        :param str created_by: The identity that created the resource.
        :param str created_by_type: The type of identity that created the resource.
        :param str last_modified_at: The timestamp of resource last modification (UTC)
        :param str last_modified_by: The identity that last modified the resource.
        :param str last_modified_by_type: The type of identity that last modified the resource.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if created_by is not None:
            pulumi.set(__self__, "created_by", created_by)
        if created_by_type is not None:
            pulumi.set(__self__, "created_by_type", created_by_type)
        if last_modified_at is not None:
            pulumi.set(__self__, "last_modified_at", last_modified_at)
        if last_modified_by is not None:
            pulumi.set(__self__, "last_modified_by", last_modified_by)
        if last_modified_by_type is not None:
            pulumi.set(__self__, "last_modified_by_type", last_modified_by_type)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        The timestamp of resource creation (UTC).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> Optional[str]:
        """
        The identity that created the resource.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="createdByType")
    def created_by_type(self) -> Optional[str]:
        """
        The type of identity that created the resource.
        """
        return pulumi.get(self, "created_by_type")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> Optional[str]:
        """
        The timestamp of resource last modification (UTC)
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter(name="lastModifiedBy")
    def last_modified_by(self) -> Optional[str]:
        """
        The identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by")

    @property
    @pulumi.getter(name="lastModifiedByType")
    def last_modified_by_type(self) -> Optional[str]:
        """
        The type of identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by_type")


