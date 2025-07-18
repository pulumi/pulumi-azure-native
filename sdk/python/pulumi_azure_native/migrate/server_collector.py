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
from ._inputs import *

__all__ = ['ServerCollectorArgs', 'ServerCollector']

@pulumi.input_type
class ServerCollectorArgs:
    def __init__(__self__, *,
                 project_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 e_tag: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input['CollectorPropertiesArgs']] = None,
                 server_collector_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ServerCollector resource.
        :param pulumi.Input[builtins.str] project_name: Name of the Azure Migrate project.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the Azure Resource Group that project is part of.
        :param pulumi.Input[builtins.str] server_collector_name: Unique name of a Server collector within a project.
        """
        pulumi.set(__self__, "project_name", project_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if e_tag is not None:
            pulumi.set(__self__, "e_tag", e_tag)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)
        if server_collector_name is not None:
            pulumi.set(__self__, "server_collector_name", server_collector_name)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the Azure Migrate project.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the Azure Resource Group that project is part of.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="eTag")
    def e_tag(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "e_tag")

    @e_tag.setter
    def e_tag(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "e_tag", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input['CollectorPropertiesArgs']]:
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input['CollectorPropertiesArgs']]):
        pulumi.set(self, "properties", value)

    @property
    @pulumi.getter(name="serverCollectorName")
    def server_collector_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Unique name of a Server collector within a project.
        """
        return pulumi.get(self, "server_collector_name")

    @server_collector_name.setter
    def server_collector_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "server_collector_name", value)


@pulumi.type_token("azure-native:migrate:ServerCollector")
class ServerCollector(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 e_tag: Optional[pulumi.Input[builtins.str]] = None,
                 project_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['CollectorPropertiesArgs', 'CollectorPropertiesArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 server_collector_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Uses Azure REST API version 2019-10-01. In version 2.x of the Azure Native provider, it used API version 2019-10-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] project_name: Name of the Azure Migrate project.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the Azure Resource Group that project is part of.
        :param pulumi.Input[builtins.str] server_collector_name: Unique name of a Server collector within a project.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServerCollectorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Uses Azure REST API version 2019-10-01. In version 2.x of the Azure Native provider, it used API version 2019-10-01.

        :param str resource_name: The name of the resource.
        :param ServerCollectorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerCollectorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 e_tag: Optional[pulumi.Input[builtins.str]] = None,
                 project_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['CollectorPropertiesArgs', 'CollectorPropertiesArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 server_collector_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServerCollectorArgs.__new__(ServerCollectorArgs)

            __props__.__dict__["e_tag"] = e_tag
            if project_name is None and not opts.urn:
                raise TypeError("Missing required property 'project_name'")
            __props__.__dict__["project_name"] = project_name
            __props__.__dict__["properties"] = properties
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["server_collector_name"] = server_collector_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:migrate/v20191001:ServerCollector"), pulumi.Alias(type_="azure-native:migrate/v20230315:ServerCollector"), pulumi.Alias(type_="azure-native:migrate/v20230315:ServerCollectorsOperation"), pulumi.Alias(type_="azure-native:migrate/v20230401preview:ServerCollector"), pulumi.Alias(type_="azure-native:migrate/v20230401preview:ServerCollectorsOperation"), pulumi.Alias(type_="azure-native:migrate/v20230501preview:ServerCollector"), pulumi.Alias(type_="azure-native:migrate/v20230501preview:ServerCollectorsOperation"), pulumi.Alias(type_="azure-native:migrate/v20230909preview:ServerCollector"), pulumi.Alias(type_="azure-native:migrate/v20230909preview:ServerCollectorsOperation"), pulumi.Alias(type_="azure-native:migrate/v20240101preview:ServerCollector"), pulumi.Alias(type_="azure-native:migrate/v20240101preview:ServerCollectorsOperation"), pulumi.Alias(type_="azure-native:migrate/v20240115:ServerCollector"), pulumi.Alias(type_="azure-native:migrate/v20240303preview:ServerCollector"), pulumi.Alias(type_="azure-native:migrate:ServerCollectorsOperation")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ServerCollector, __self__).__init__(
            'azure-native:migrate:ServerCollector',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ServerCollector':
        """
        Get an existing ServerCollector resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ServerCollectorArgs.__new__(ServerCollectorArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["e_tag"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["properties"] = None
        __props__.__dict__["type"] = None
        return ServerCollector(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="eTag")
    def e_tag(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "e_tag")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output['outputs.CollectorPropertiesResponse']:
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "type")

