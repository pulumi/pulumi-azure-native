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
    'AppServiceCertificateArgs',
    'AppServiceCertificateArgsDict',
]

MYPY = False

if not MYPY:
    class AppServiceCertificateArgsDict(TypedDict):
        """
        Key Vault container for a certificate that is purchased through Azure.
        """
        key_vault_id: NotRequired[pulumi.Input[builtins.str]]
        """
        Key Vault resource Id.
        """
        key_vault_secret_name: NotRequired[pulumi.Input[builtins.str]]
        """
        Key Vault secret name.
        """
elif False:
    AppServiceCertificateArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class AppServiceCertificateArgs:
    def __init__(__self__, *,
                 key_vault_id: Optional[pulumi.Input[builtins.str]] = None,
                 key_vault_secret_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        Key Vault container for a certificate that is purchased through Azure.
        :param pulumi.Input[builtins.str] key_vault_id: Key Vault resource Id.
        :param pulumi.Input[builtins.str] key_vault_secret_name: Key Vault secret name.
        """
        if key_vault_id is not None:
            pulumi.set(__self__, "key_vault_id", key_vault_id)
        if key_vault_secret_name is not None:
            pulumi.set(__self__, "key_vault_secret_name", key_vault_secret_name)

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Key Vault resource Id.
        """
        return pulumi.get(self, "key_vault_id")

    @key_vault_id.setter
    def key_vault_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "key_vault_id", value)

    @property
    @pulumi.getter(name="keyVaultSecretName")
    def key_vault_secret_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Key Vault secret name.
        """
        return pulumi.get(self, "key_vault_secret_name")

    @key_vault_secret_name.setter
    def key_vault_secret_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "key_vault_secret_name", value)


