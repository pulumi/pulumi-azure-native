# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'PrivateEndpointServiceConnectionStatus',
]


class PrivateEndpointServiceConnectionStatus(str, Enum):
    """
    Indicates whether the connection has been approved, rejected or removed by the Relay Namespace owner.
    """
    PENDING = "Pending"
    APPROVED = "Approved"
    REJECTED = "Rejected"
    DISCONNECTED = "Disconnected"
