# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'CloudType',
]


@pulumi.type_token("azure-native:hybridcloud:CloudType")
class CloudType(builtins.str, Enum):
    """
    The cloud connector type.
    """
    AWS = "AWS"
