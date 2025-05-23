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
    'BlobStorageClassTypePropertiesResponse',
    'NativeStorageClassTypePropertiesResponse',
    'NfsStorageClassTypePropertiesResponse',
    'RwxStorageClassTypePropertiesResponse',
    'SmbStorageClassTypePropertiesResponse',
    'SystemDataResponse',
]

@pulumi.output_type
class BlobStorageClassTypePropertiesResponse(dict):
    """
    The properties of Blob StorageClass
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "azureStorageAccountKey":
            suggest = "azure_storage_account_key"
        elif key == "azureStorageAccountName":
            suggest = "azure_storage_account_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BlobStorageClassTypePropertiesResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BlobStorageClassTypePropertiesResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BlobStorageClassTypePropertiesResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 azure_storage_account_key: builtins.str,
                 azure_storage_account_name: builtins.str,
                 type: builtins.str):
        """
        The properties of Blob StorageClass
        :param builtins.str azure_storage_account_key: Azure Storage Account Key
        :param builtins.str azure_storage_account_name: Azure Storage Account Name
        :param builtins.str type: Type of a storage class
               Expected value is 'Blob'.
        """
        pulumi.set(__self__, "azure_storage_account_key", azure_storage_account_key)
        pulumi.set(__self__, "azure_storage_account_name", azure_storage_account_name)
        pulumi.set(__self__, "type", 'Blob')

    @property
    @pulumi.getter(name="azureStorageAccountKey")
    def azure_storage_account_key(self) -> builtins.str:
        """
        Azure Storage Account Key
        """
        return pulumi.get(self, "azure_storage_account_key")

    @property
    @pulumi.getter(name="azureStorageAccountName")
    def azure_storage_account_name(self) -> builtins.str:
        """
        Azure Storage Account Name
        """
        return pulumi.get(self, "azure_storage_account_name")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Type of a storage class
        Expected value is 'Blob'.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class NativeStorageClassTypePropertiesResponse(dict):
    """
    The properties of Native StorageClass
    """
    def __init__(__self__, *,
                 type: builtins.str):
        """
        The properties of Native StorageClass
        :param builtins.str type: Type of a storage class
               Expected value is 'Native'.
        """
        pulumi.set(__self__, "type", 'Native')

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Type of a storage class
        Expected value is 'Native'.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class NfsStorageClassTypePropertiesResponse(dict):
    """
    The properties of NFS StorageClass
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "mountPermissions":
            suggest = "mount_permissions"
        elif key == "onDelete":
            suggest = "on_delete"
        elif key == "subDir":
            suggest = "sub_dir"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NfsStorageClassTypePropertiesResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NfsStorageClassTypePropertiesResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NfsStorageClassTypePropertiesResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 server: builtins.str,
                 share: builtins.str,
                 type: builtins.str,
                 mount_permissions: Optional[builtins.str] = None,
                 on_delete: Optional[builtins.str] = None,
                 sub_dir: Optional[builtins.str] = None):
        """
        The properties of NFS StorageClass
        :param builtins.str server: NFS Server
        :param builtins.str share: NFS share
        :param builtins.str type: Type of a storage class
               Expected value is 'NFS'.
        :param builtins.str mount_permissions: Mounted folder permissions. Default is 0. If set as non-zero, driver will perform `chmod` after mount
        :param builtins.str on_delete: The action to take when a NFS volume is deleted. Default is Delete
        :param builtins.str sub_dir: Sub directory under share. If the sub directory doesn't exist, driver will create it
        """
        pulumi.set(__self__, "server", server)
        pulumi.set(__self__, "share", share)
        pulumi.set(__self__, "type", 'NFS')
        if mount_permissions is not None:
            pulumi.set(__self__, "mount_permissions", mount_permissions)
        if on_delete is not None:
            pulumi.set(__self__, "on_delete", on_delete)
        if sub_dir is not None:
            pulumi.set(__self__, "sub_dir", sub_dir)

    @property
    @pulumi.getter
    def server(self) -> builtins.str:
        """
        NFS Server
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter
    def share(self) -> builtins.str:
        """
        NFS share
        """
        return pulumi.get(self, "share")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Type of a storage class
        Expected value is 'NFS'.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="mountPermissions")
    def mount_permissions(self) -> Optional[builtins.str]:
        """
        Mounted folder permissions. Default is 0. If set as non-zero, driver will perform `chmod` after mount
        """
        return pulumi.get(self, "mount_permissions")

    @property
    @pulumi.getter(name="onDelete")
    def on_delete(self) -> Optional[builtins.str]:
        """
        The action to take when a NFS volume is deleted. Default is Delete
        """
        return pulumi.get(self, "on_delete")

    @property
    @pulumi.getter(name="subDir")
    def sub_dir(self) -> Optional[builtins.str]:
        """
        Sub directory under share. If the sub directory doesn't exist, driver will create it
        """
        return pulumi.get(self, "sub_dir")


@pulumi.output_type
class RwxStorageClassTypePropertiesResponse(dict):
    """
    The properties of RWX StorageClass
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "backingStorageClassName":
            suggest = "backing_storage_class_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RwxStorageClassTypePropertiesResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RwxStorageClassTypePropertiesResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RwxStorageClassTypePropertiesResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 backing_storage_class_name: builtins.str,
                 type: builtins.str):
        """
        The properties of RWX StorageClass
        :param builtins.str backing_storage_class_name: The backing storageclass used to create new storageclass
        :param builtins.str type: Type of a storage class
               Expected value is 'RWX'.
        """
        pulumi.set(__self__, "backing_storage_class_name", backing_storage_class_name)
        pulumi.set(__self__, "type", 'RWX')

    @property
    @pulumi.getter(name="backingStorageClassName")
    def backing_storage_class_name(self) -> builtins.str:
        """
        The backing storageclass used to create new storageclass
        """
        return pulumi.get(self, "backing_storage_class_name")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Type of a storage class
        Expected value is 'RWX'.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class SmbStorageClassTypePropertiesResponse(dict):
    """
    The properties of SMB StorageClass
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "subDir":
            suggest = "sub_dir"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SmbStorageClassTypePropertiesResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SmbStorageClassTypePropertiesResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SmbStorageClassTypePropertiesResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 source: builtins.str,
                 type: builtins.str,
                 domain: Optional[builtins.str] = None,
                 password: Optional[builtins.str] = None,
                 sub_dir: Optional[builtins.str] = None,
                 username: Optional[builtins.str] = None):
        """
        The properties of SMB StorageClass
        :param builtins.str source: SMB Source
        :param builtins.str type: Type of a storage class
               Expected value is 'SMB'.
        :param builtins.str domain: Server domain
        :param builtins.str password: Server password
        :param builtins.str sub_dir: Sub directory under share. If the sub directory doesn't exist, driver will create it
        :param builtins.str username: Server username
        """
        pulumi.set(__self__, "source", source)
        pulumi.set(__self__, "type", 'SMB')
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if sub_dir is not None:
            pulumi.set(__self__, "sub_dir", sub_dir)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def source(self) -> builtins.str:
        """
        SMB Source
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Type of a storage class
        Expected value is 'SMB'.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def domain(self) -> Optional[builtins.str]:
        """
        Server domain
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def password(self) -> Optional[builtins.str]:
        """
        Server password
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="subDir")
    def sub_dir(self) -> Optional[builtins.str]:
        """
        Sub directory under share. If the sub directory doesn't exist, driver will create it
        """
        return pulumi.get(self, "sub_dir")

    @property
    @pulumi.getter
    def username(self) -> Optional[builtins.str]:
        """
        Server username
        """
        return pulumi.get(self, "username")


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
                 created_at: Optional[builtins.str] = None,
                 created_by: Optional[builtins.str] = None,
                 created_by_type: Optional[builtins.str] = None,
                 last_modified_at: Optional[builtins.str] = None,
                 last_modified_by: Optional[builtins.str] = None,
                 last_modified_by_type: Optional[builtins.str] = None):
        """
        Metadata pertaining to creation and last modification of the resource.
        :param builtins.str created_at: The timestamp of resource creation (UTC).
        :param builtins.str created_by: The identity that created the resource.
        :param builtins.str created_by_type: The type of identity that created the resource.
        :param builtins.str last_modified_at: The timestamp of resource last modification (UTC)
        :param builtins.str last_modified_by: The identity that last modified the resource.
        :param builtins.str last_modified_by_type: The type of identity that last modified the resource.
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
    def created_at(self) -> Optional[builtins.str]:
        """
        The timestamp of resource creation (UTC).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> Optional[builtins.str]:
        """
        The identity that created the resource.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="createdByType")
    def created_by_type(self) -> Optional[builtins.str]:
        """
        The type of identity that created the resource.
        """
        return pulumi.get(self, "created_by_type")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> Optional[builtins.str]:
        """
        The timestamp of resource last modification (UTC)
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter(name="lastModifiedBy")
    def last_modified_by(self) -> Optional[builtins.str]:
        """
        The identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by")

    @property
    @pulumi.getter(name="lastModifiedByType")
    def last_modified_by_type(self) -> Optional[builtins.str]:
        """
        The type of identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by_type")


