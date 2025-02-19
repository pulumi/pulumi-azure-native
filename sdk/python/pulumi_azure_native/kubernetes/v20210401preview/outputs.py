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
from ._enums import *

__all__ = [
    'CredentialResultResponse',
    'HybridConnectionConfigResponse',
]

@pulumi.output_type
class CredentialResultResponse(dict):
    """
    The credential result response.
    """
    def __init__(__self__, *,
                 name: str,
                 value: str):
        """
        The credential result response.
        :param str name: The name of the credential.
        :param str value: Base64-encoded Kubernetes configuration file.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the credential.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Base64-encoded Kubernetes configuration file.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class HybridConnectionConfigResponse(dict):
    """
    Contains the REP (rendezvous endpoint) and “Sender” access token.
    """
    def __init__(__self__, *,
                 expiration_time: float,
                 hybrid_connection_name: str,
                 relay: str,
                 token: str):
        """
        Contains the REP (rendezvous endpoint) and “Sender” access token.
        :param float expiration_time: Timestamp when this token will be expired.
        :param str hybrid_connection_name: Name of the connection
        :param str relay: Name of the relay.
        :param str token: Sender access token
        """
        pulumi.set(__self__, "expiration_time", expiration_time)
        pulumi.set(__self__, "hybrid_connection_name", hybrid_connection_name)
        pulumi.set(__self__, "relay", relay)
        pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="expirationTime")
    def expiration_time(self) -> float:
        """
        Timestamp when this token will be expired.
        """
        return pulumi.get(self, "expiration_time")

    @property
    @pulumi.getter(name="hybridConnectionName")
    def hybrid_connection_name(self) -> str:
        """
        Name of the connection
        """
        return pulumi.get(self, "hybrid_connection_name")

    @property
    @pulumi.getter
    def relay(self) -> str:
        """
        Name of the relay.
        """
        return pulumi.get(self, "relay")

    @property
    @pulumi.getter
    def token(self) -> str:
        """
        Sender access token
        """
        return pulumi.get(self, "token")


