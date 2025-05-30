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

__all__ = ['LinkedServiceArgs', 'LinkedService']

@pulumi.input_type
class LinkedServiceArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 workspace_name: pulumi.Input[builtins.str],
                 linked_service_name: Optional[pulumi.Input[builtins.str]] = None,
                 provisioning_state: Optional[pulumi.Input[Union[builtins.str, 'LinkedServiceEntityStatus']]] = None,
                 resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 write_access_resource_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a LinkedService resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] workspace_name: The name of the workspace.
        :param pulumi.Input[builtins.str] linked_service_name: Name of the linkedServices resource
        :param pulumi.Input[Union[builtins.str, 'LinkedServiceEntityStatus']] provisioning_state: The provisioning state of the linked service.
        :param pulumi.Input[builtins.str] resource_id: The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require read access
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[builtins.str] write_access_resource_id: The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require write access
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "workspace_name", workspace_name)
        if linked_service_name is not None:
            pulumi.set(__self__, "linked_service_name", linked_service_name)
        if provisioning_state is not None:
            pulumi.set(__self__, "provisioning_state", provisioning_state)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if write_access_resource_id is not None:
            pulumi.set(__self__, "write_access_resource_id", write_access_resource_id)

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
    @pulumi.getter(name="workspaceName")
    def workspace_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the workspace.
        """
        return pulumi.get(self, "workspace_name")

    @workspace_name.setter
    def workspace_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "workspace_name", value)

    @property
    @pulumi.getter(name="linkedServiceName")
    def linked_service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the linkedServices resource
        """
        return pulumi.get(self, "linked_service_name")

    @linked_service_name.setter
    def linked_service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "linked_service_name", value)

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[pulumi.Input[Union[builtins.str, 'LinkedServiceEntityStatus']]]:
        """
        The provisioning state of the linked service.
        """
        return pulumi.get(self, "provisioning_state")

    @provisioning_state.setter
    def provisioning_state(self, value: Optional[pulumi.Input[Union[builtins.str, 'LinkedServiceEntityStatus']]]):
        pulumi.set(self, "provisioning_state", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require read access
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="writeAccessResourceId")
    def write_access_resource_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require write access
        """
        return pulumi.get(self, "write_access_resource_id")

    @write_access_resource_id.setter
    def write_access_resource_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "write_access_resource_id", value)


@pulumi.type_token("azure-native:operationalinsights:LinkedService")
class LinkedService(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 linked_service_name: Optional[pulumi.Input[builtins.str]] = None,
                 provisioning_state: Optional[pulumi.Input[Union[builtins.str, 'LinkedServiceEntityStatus']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                 write_access_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The top level Linked service resource container.

        Uses Azure REST API version 2023-09-01. In version 2.x of the Azure Native provider, it used API version 2020-08-01.

        Other available API versions: 2015-11-01-preview, 2019-08-01-preview, 2020-03-01-preview, 2020-08-01, 2025-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native operationalinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] linked_service_name: Name of the linkedServices resource
        :param pulumi.Input[Union[builtins.str, 'LinkedServiceEntityStatus']] provisioning_state: The provisioning state of the linked service.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] resource_id: The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require read access
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[builtins.str] workspace_name: The name of the workspace.
        :param pulumi.Input[builtins.str] write_access_resource_id: The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require write access
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LinkedServiceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The top level Linked service resource container.

        Uses Azure REST API version 2023-09-01. In version 2.x of the Azure Native provider, it used API version 2020-08-01.

        Other available API versions: 2015-11-01-preview, 2019-08-01-preview, 2020-03-01-preview, 2020-08-01, 2025-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native operationalinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param LinkedServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LinkedServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 linked_service_name: Optional[pulumi.Input[builtins.str]] = None,
                 provisioning_state: Optional[pulumi.Input[Union[builtins.str, 'LinkedServiceEntityStatus']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                 write_access_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LinkedServiceArgs.__new__(LinkedServiceArgs)

            __props__.__dict__["linked_service_name"] = linked_service_name
            __props__.__dict__["provisioning_state"] = provisioning_state
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["resource_id"] = resource_id
            __props__.__dict__["tags"] = tags
            if workspace_name is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_name'")
            __props__.__dict__["workspace_name"] = workspace_name
            __props__.__dict__["write_access_resource_id"] = write_access_resource_id
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:operationalinsights/v20151101preview:LinkedService"), pulumi.Alias(type_="azure-native:operationalinsights/v20190801preview:LinkedService"), pulumi.Alias(type_="azure-native:operationalinsights/v20200301preview:LinkedService"), pulumi.Alias(type_="azure-native:operationalinsights/v20200801:LinkedService"), pulumi.Alias(type_="azure-native:operationalinsights/v20230901:LinkedService"), pulumi.Alias(type_="azure-native:operationalinsights/v20250201:LinkedService")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(LinkedService, __self__).__init__(
            'azure-native:operationalinsights:LinkedService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'LinkedService':
        """
        Get an existing LinkedService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LinkedServiceArgs.__new__(LinkedServiceArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["resource_id"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["write_access_resource_id"] = None
        return LinkedService(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The provisioning state of the linked service.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require read access
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="writeAccessResourceId")
    def write_access_resource_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require write access
        """
        return pulumi.get(self, "write_access_resource_id")

