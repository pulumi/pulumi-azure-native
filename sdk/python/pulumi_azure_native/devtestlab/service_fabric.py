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
from . import outputs

__all__ = ['ServiceFabricArgs', 'ServiceFabric']

@pulumi.input_type
class ServiceFabricArgs:
    def __init__(__self__, *,
                 lab_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 user_name: pulumi.Input[builtins.str],
                 environment_id: Optional[pulumi.Input[builtins.str]] = None,
                 external_service_fabric_id: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a ServiceFabric resource.
        :param pulumi.Input[builtins.str] lab_name: The name of the lab.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] user_name: The name of the user profile.
        :param pulumi.Input[builtins.str] environment_id: The resource id of the environment under which the service fabric resource is present
        :param pulumi.Input[builtins.str] external_service_fabric_id: The backing service fabric resource's id
        :param pulumi.Input[builtins.str] location: The location of the resource.
        :param pulumi.Input[builtins.str] name: The name of the ServiceFabric
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: The tags of the resource.
        """
        pulumi.set(__self__, "lab_name", lab_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "user_name", user_name)
        if environment_id is not None:
            pulumi.set(__self__, "environment_id", environment_id)
        if external_service_fabric_id is not None:
            pulumi.set(__self__, "external_service_fabric_id", external_service_fabric_id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="labName")
    def lab_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the lab.
        """
        return pulumi.get(self, "lab_name")

    @lab_name.setter
    def lab_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "lab_name", value)

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
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the user profile.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The resource id of the environment under which the service fabric resource is present
        """
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter(name="externalServiceFabricId")
    def external_service_fabric_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The backing service fabric resource's id
        """
        return pulumi.get(self, "external_service_fabric_id")

    @external_service_fabric_id.setter
    def external_service_fabric_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "external_service_fabric_id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The location of the resource.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the ServiceFabric
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("azure-native:devtestlab:ServiceFabric")
class ServiceFabric(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 environment_id: Optional[pulumi.Input[builtins.str]] = None,
                 external_service_fabric_id: Optional[pulumi.Input[builtins.str]] = None,
                 lab_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 user_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        A Service Fabric.

        Uses Azure REST API version 2018-09-15. In version 2.x of the Azure Native provider, it used API version 2018-09-15.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] environment_id: The resource id of the environment under which the service fabric resource is present
        :param pulumi.Input[builtins.str] external_service_fabric_id: The backing service fabric resource's id
        :param pulumi.Input[builtins.str] lab_name: The name of the lab.
        :param pulumi.Input[builtins.str] location: The location of the resource.
        :param pulumi.Input[builtins.str] name: The name of the ServiceFabric
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: The tags of the resource.
        :param pulumi.Input[builtins.str] user_name: The name of the user profile.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceFabricArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Service Fabric.

        Uses Azure REST API version 2018-09-15. In version 2.x of the Azure Native provider, it used API version 2018-09-15.

        :param str resource_name: The name of the resource.
        :param ServiceFabricArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceFabricArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 environment_id: Optional[pulumi.Input[builtins.str]] = None,
                 external_service_fabric_id: Optional[pulumi.Input[builtins.str]] = None,
                 lab_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 user_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceFabricArgs.__new__(ServiceFabricArgs)

            __props__.__dict__["environment_id"] = environment_id
            __props__.__dict__["external_service_fabric_id"] = external_service_fabric_id
            if lab_name is None and not opts.urn:
                raise TypeError("Missing required property 'lab_name'")
            __props__.__dict__["lab_name"] = lab_name
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
            __props__.__dict__["applicable_schedule"] = None
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["unique_identifier"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:devtestlab/v20180915:ServiceFabric")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ServiceFabric, __self__).__init__(
            'azure-native:devtestlab:ServiceFabric',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ServiceFabric':
        """
        Get an existing ServiceFabric resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ServiceFabricArgs.__new__(ServiceFabricArgs)

        __props__.__dict__["applicable_schedule"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["environment_id"] = None
        __props__.__dict__["external_service_fabric_id"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["unique_identifier"] = None
        return ServiceFabric(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicableSchedule")
    def applicable_schedule(self) -> pulumi.Output['outputs.ApplicableScheduleResponse']:
        """
        The applicable schedule for the virtual machine.
        """
        return pulumi.get(self, "applicable_schedule")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The resource id of the environment under which the service fabric resource is present
        """
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter(name="externalServiceFabricId")
    def external_service_fabric_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The backing service fabric resource's id
        """
        return pulumi.get(self, "external_service_fabric_id")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        The provisioning status of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="uniqueIdentifier")
    def unique_identifier(self) -> pulumi.Output[builtins.str]:
        """
        The unique immutable identifier of a resource (Guid).
        """
        return pulumi.get(self, "unique_identifier")

