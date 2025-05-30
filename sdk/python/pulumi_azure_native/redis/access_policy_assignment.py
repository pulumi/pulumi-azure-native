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

__all__ = ['AccessPolicyAssignmentArgs', 'AccessPolicyAssignment']

@pulumi.input_type
class AccessPolicyAssignmentArgs:
    def __init__(__self__, *,
                 access_policy_name: pulumi.Input[builtins.str],
                 cache_name: pulumi.Input[builtins.str],
                 object_id: pulumi.Input[builtins.str],
                 object_id_alias: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 access_policy_assignment_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a AccessPolicyAssignment resource.
        :param pulumi.Input[builtins.str] access_policy_name: The name of the access policy that is being assigned
        :param pulumi.Input[builtins.str] cache_name: The name of the Redis cache.
        :param pulumi.Input[builtins.str] object_id: Object Id to assign access policy to
        :param pulumi.Input[builtins.str] object_id_alias: User friendly name for object id. Also represents username for token based authentication
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] access_policy_assignment_name: The name of the access policy assignment.
        """
        pulumi.set(__self__, "access_policy_name", access_policy_name)
        pulumi.set(__self__, "cache_name", cache_name)
        pulumi.set(__self__, "object_id", object_id)
        pulumi.set(__self__, "object_id_alias", object_id_alias)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if access_policy_assignment_name is not None:
            pulumi.set(__self__, "access_policy_assignment_name", access_policy_assignment_name)

    @property
    @pulumi.getter(name="accessPolicyName")
    def access_policy_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the access policy that is being assigned
        """
        return pulumi.get(self, "access_policy_name")

    @access_policy_name.setter
    def access_policy_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "access_policy_name", value)

    @property
    @pulumi.getter(name="cacheName")
    def cache_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the Redis cache.
        """
        return pulumi.get(self, "cache_name")

    @cache_name.setter
    def cache_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "cache_name", value)

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> pulumi.Input[builtins.str]:
        """
        Object Id to assign access policy to
        """
        return pulumi.get(self, "object_id")

    @object_id.setter
    def object_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "object_id", value)

    @property
    @pulumi.getter(name="objectIdAlias")
    def object_id_alias(self) -> pulumi.Input[builtins.str]:
        """
        User friendly name for object id. Also represents username for token based authentication
        """
        return pulumi.get(self, "object_id_alias")

    @object_id_alias.setter
    def object_id_alias(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "object_id_alias", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group. The name is case insensitive.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="accessPolicyAssignmentName")
    def access_policy_assignment_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the access policy assignment.
        """
        return pulumi.get(self, "access_policy_assignment_name")

    @access_policy_assignment_name.setter
    def access_policy_assignment_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "access_policy_assignment_name", value)


@pulumi.type_token("azure-native:redis:AccessPolicyAssignment")
class AccessPolicyAssignment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_policy_assignment_name: Optional[pulumi.Input[builtins.str]] = None,
                 access_policy_name: Optional[pulumi.Input[builtins.str]] = None,
                 cache_name: Optional[pulumi.Input[builtins.str]] = None,
                 object_id: Optional[pulumi.Input[builtins.str]] = None,
                 object_id_alias: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Response to an operation on access policy assignment

        Uses Azure REST API version 2024-11-01.

        Other available API versions: 2023-05-01-preview, 2023-08-01, 2024-03-01, 2024-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native redis [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] access_policy_assignment_name: The name of the access policy assignment.
        :param pulumi.Input[builtins.str] access_policy_name: The name of the access policy that is being assigned
        :param pulumi.Input[builtins.str] cache_name: The name of the Redis cache.
        :param pulumi.Input[builtins.str] object_id: Object Id to assign access policy to
        :param pulumi.Input[builtins.str] object_id_alias: User friendly name for object id. Also represents username for token based authentication
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccessPolicyAssignmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Response to an operation on access policy assignment

        Uses Azure REST API version 2024-11-01.

        Other available API versions: 2023-05-01-preview, 2023-08-01, 2024-03-01, 2024-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native redis [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param AccessPolicyAssignmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccessPolicyAssignmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_policy_assignment_name: Optional[pulumi.Input[builtins.str]] = None,
                 access_policy_name: Optional[pulumi.Input[builtins.str]] = None,
                 cache_name: Optional[pulumi.Input[builtins.str]] = None,
                 object_id: Optional[pulumi.Input[builtins.str]] = None,
                 object_id_alias: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccessPolicyAssignmentArgs.__new__(AccessPolicyAssignmentArgs)

            __props__.__dict__["access_policy_assignment_name"] = access_policy_assignment_name
            if access_policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'access_policy_name'")
            __props__.__dict__["access_policy_name"] = access_policy_name
            if cache_name is None and not opts.urn:
                raise TypeError("Missing required property 'cache_name'")
            __props__.__dict__["cache_name"] = cache_name
            if object_id is None and not opts.urn:
                raise TypeError("Missing required property 'object_id'")
            __props__.__dict__["object_id"] = object_id
            if object_id_alias is None and not opts.urn:
                raise TypeError("Missing required property 'object_id_alias'")
            __props__.__dict__["object_id_alias"] = object_id_alias
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:cache/v20230501preview:AccessPolicyAssignment"), pulumi.Alias(type_="azure-native:cache/v20230801:AccessPolicyAssignment"), pulumi.Alias(type_="azure-native:cache/v20240301:AccessPolicyAssignment"), pulumi.Alias(type_="azure-native:cache/v20240401preview:AccessPolicyAssignment"), pulumi.Alias(type_="azure-native:cache/v20241101:AccessPolicyAssignment"), pulumi.Alias(type_="azure-native:cache:AccessPolicyAssignment"), pulumi.Alias(type_="azure-native:redis/v20230501preview:AccessPolicyAssignment"), pulumi.Alias(type_="azure-native:redis/v20230801:AccessPolicyAssignment"), pulumi.Alias(type_="azure-native:redis/v20240301:AccessPolicyAssignment"), pulumi.Alias(type_="azure-native:redis/v20240401preview:AccessPolicyAssignment"), pulumi.Alias(type_="azure-native:redis/v20241101:AccessPolicyAssignment")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AccessPolicyAssignment, __self__).__init__(
            'azure-native:redis:AccessPolicyAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AccessPolicyAssignment':
        """
        Get an existing AccessPolicyAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AccessPolicyAssignmentArgs.__new__(AccessPolicyAssignmentArgs)

        __props__.__dict__["access_policy_name"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["object_id"] = None
        __props__.__dict__["object_id_alias"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["type"] = None
        return AccessPolicyAssignment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessPolicyName")
    def access_policy_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the access policy that is being assigned
        """
        return pulumi.get(self, "access_policy_name")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> pulumi.Output[builtins.str]:
        """
        Object Id to assign access policy to
        """
        return pulumi.get(self, "object_id")

    @property
    @pulumi.getter(name="objectIdAlias")
    def object_id_alias(self) -> pulumi.Output[builtins.str]:
        """
        User friendly name for object id. Also represents username for token based authentication
        """
        return pulumi.get(self, "object_id_alias")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Provisioning state of an access policy assignment set
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

