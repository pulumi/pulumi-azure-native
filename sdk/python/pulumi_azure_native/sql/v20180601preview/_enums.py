# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'SecurityAlertPolicyState',
]


class SecurityAlertPolicyState(str, Enum):
    """
    Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database.
    """
    NEW = "New"
    ENABLED = "Enabled"
    DISABLED = "Disabled"
