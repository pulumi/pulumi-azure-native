# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'MigrationMode',
]


class MigrationMode(str, Enum):
    """
    There are two types of migration modes Online and Offline
    """
    OFFLINE = "Offline"
    ONLINE = "Online"
