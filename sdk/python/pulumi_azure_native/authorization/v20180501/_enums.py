# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'PolicyMode',
    'PolicyType',
]


class PolicyMode(str, Enum):
    """
    The policy definition mode. Possible values are NotSpecified, Indexed, and All.
    """
    NOT_SPECIFIED = "NotSpecified"
    INDEXED = "Indexed"
    ALL = "All"


class PolicyType(str, Enum):
    """
    The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
    """
    NOT_SPECIFIED = "NotSpecified"
    BUILT_IN = "BuiltIn"
    CUSTOM = "Custom"
