# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AssignmentType',
    'EnforcementMode',
    'OverrideKind',
    'ResourceIdentityType',
    'SelectorKind',
]


class AssignmentType(str, Enum):
    """
    The type of policy assignment. Possible values are NotSpecified, System, SystemHidden, and Custom. Immutable.
    """
    NOT_SPECIFIED = "NotSpecified"
    SYSTEM = "System"
    SYSTEM_HIDDEN = "SystemHidden"
    CUSTOM = "Custom"


class EnforcementMode(str, Enum):
    """
    The policy assignment enforcement mode. Possible values are Default and DoNotEnforce.
    """
    DEFAULT = "Default"
    """
    The policy effect is enforced during resource creation or update.
    """
    DO_NOT_ENFORCE = "DoNotEnforce"
    """
    The policy effect is not enforced during resource creation or update.
    """


class OverrideKind(str, Enum):
    """
    The override kind.
    """
    POLICY_EFFECT = "policyEffect"
    """
    It will override the policy effect type.
    """


class ResourceIdentityType(str, Enum):
    """
    The identity type. This is the only required field when adding a system or user assigned identity to a resource.
    """
    SYSTEM_ASSIGNED = "SystemAssigned"
    """
    Indicates that a system assigned identity is associated with the resource.
    """
    USER_ASSIGNED = "UserAssigned"
    """
    Indicates that a system assigned identity is associated with the resource.
    """
    NONE = "None"
    """
    Indicates that no identity is associated with the resource or that the existing identity should be removed.
    """


class SelectorKind(str, Enum):
    """
    The selector kind.
    """
    RESOURCE_LOCATION = "resourceLocation"
    """
    The selector kind to filter policies by the resource location.
    """
    RESOURCE_TYPE = "resourceType"
    """
    The selector kind to filter policies by the resource type.
    """
    RESOURCE_WITHOUT_LOCATION = "resourceWithoutLocation"
    """
    The selector kind to filter policies by the resource without location.
    """
    POLICY_DEFINITION_REFERENCE_ID = "policyDefinitionReferenceId"
    """
    The selector kind to filter policies by the policy definition reference ID.
    """
