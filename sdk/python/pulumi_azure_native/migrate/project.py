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
from ._enums import *
from ._inputs import *

__all__ = ['ProjectArgs', 'Project']

@pulumi.input_type
class ProjectArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 e_tag: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 project_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input['ProjectPropertiesArgs']] = None,
                 tags: Optional[Any] = None):
        """
        The set of arguments for constructing a Project resource.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the Azure Resource Group that project is part of.
        :param pulumi.Input[builtins.str] e_tag: For optimistic concurrency control.
        :param pulumi.Input[builtins.str] location: Azure location in which project is created.
        :param pulumi.Input[builtins.str] project_name: Name of the Azure Migrate project.
        :param pulumi.Input['ProjectPropertiesArgs'] properties: Properties of the project.
        :param Any tags: Tags provided by Azure Tagging service.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if e_tag is not None:
            pulumi.set(__self__, "e_tag", e_tag)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

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
        """
        For optimistic concurrency control.
        """
        return pulumi.get(self, "e_tag")

    @e_tag.setter
    def e_tag(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "e_tag", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Azure location in which project is created.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the Azure Migrate project.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input['ProjectPropertiesArgs']]:
        """
        Properties of the project.
        """
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input['ProjectPropertiesArgs']]):
        pulumi.set(self, "properties", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[Any]:
        """
        Tags provided by Azure Tagging service.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[Any]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("azure-native:migrate:Project")
class Project(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 e_tag: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 project_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['ProjectPropertiesArgs', 'ProjectPropertiesArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[Any] = None,
                 __props__=None):
        """
        Azure Migrate Project.

        Uses Azure REST API version 2019-10-01. In version 2.x of the Azure Native provider, it used API version 2019-10-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] e_tag: For optimistic concurrency control.
        :param pulumi.Input[builtins.str] location: Azure location in which project is created.
        :param pulumi.Input[builtins.str] project_name: Name of the Azure Migrate project.
        :param pulumi.Input[Union['ProjectPropertiesArgs', 'ProjectPropertiesArgsDict']] properties: Properties of the project.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the Azure Resource Group that project is part of.
        :param Any tags: Tags provided by Azure Tagging service.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Azure Migrate Project.

        Uses Azure REST API version 2019-10-01. In version 2.x of the Azure Native provider, it used API version 2019-10-01.

        :param str resource_name: The name of the resource.
        :param ProjectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 e_tag: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 project_name: Optional[pulumi.Input[builtins.str]] = None,
                 properties: Optional[pulumi.Input[Union['ProjectPropertiesArgs', 'ProjectPropertiesArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[Any] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectArgs.__new__(ProjectArgs)

            __props__.__dict__["e_tag"] = e_tag
            __props__.__dict__["location"] = location
            __props__.__dict__["project_name"] = project_name
            __props__.__dict__["properties"] = properties
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:migrate/v20180202:Project"), pulumi.Alias(type_="azure-native:migrate/v20191001:Project"), pulumi.Alias(type_="azure-native:migrate/v20230315:AssessmentProjectsOperation"), pulumi.Alias(type_="azure-native:migrate/v20230315:Project"), pulumi.Alias(type_="azure-native:migrate/v20230401preview:AssessmentProjectsOperation"), pulumi.Alias(type_="azure-native:migrate/v20230401preview:Project"), pulumi.Alias(type_="azure-native:migrate/v20230501preview:AssessmentProjectsOperation"), pulumi.Alias(type_="azure-native:migrate/v20230501preview:Project"), pulumi.Alias(type_="azure-native:migrate/v20230909preview:AssessmentProjectsOperation"), pulumi.Alias(type_="azure-native:migrate/v20230909preview:Project"), pulumi.Alias(type_="azure-native:migrate/v20240101preview:AssessmentProjectsOperation"), pulumi.Alias(type_="azure-native:migrate/v20240101preview:Project"), pulumi.Alias(type_="azure-native:migrate/v20240115:Project"), pulumi.Alias(type_="azure-native:migrate/v20240303preview:Project"), pulumi.Alias(type_="azure-native:migrate:AssessmentProjectsOperation")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Project, __self__).__init__(
            'azure-native:migrate:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Project':
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ProjectArgs.__new__(ProjectArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["e_tag"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["properties"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return Project(resource_name, opts=opts, __props__=__props__)

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
        """
        For optimistic concurrency control.
        """
        return pulumi.get(self, "e_tag")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Azure location in which project is created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the project.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output['outputs.ProjectPropertiesResponse']:
        """
        Properties of the project.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Any]]:
        """
        Tags provided by Azure Tagging service.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Type of the object = [Microsoft.Migrate/assessmentProjects].
        """
        return pulumi.get(self, "type")

