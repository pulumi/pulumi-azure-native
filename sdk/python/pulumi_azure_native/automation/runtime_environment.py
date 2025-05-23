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

__all__ = ['RuntimeEnvironmentArgs', 'RuntimeEnvironment']

@pulumi.input_type
class RuntimeEnvironmentArgs:
    def __init__(__self__, *,
                 automation_account_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 default_packages: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 language: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 runtime_environment_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 version: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a RuntimeEnvironment resource.
        :param pulumi.Input[builtins.str] automation_account_name: The name of the automation account.
        :param pulumi.Input[builtins.str] resource_group_name: Name of an Azure Resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] default_packages: List of Default packages for Environment
        :param pulumi.Input[builtins.str] description: Gets or sets the description.
        :param pulumi.Input[builtins.str] language: Language of Runtime Environment
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] runtime_environment_name: The name of the Runtime Environment.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[builtins.str] version: Version of Language
        """
        pulumi.set(__self__, "automation_account_name", automation_account_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if default_packages is not None:
            pulumi.set(__self__, "default_packages", default_packages)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if language is not None:
            pulumi.set(__self__, "language", language)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if runtime_environment_name is not None:
            pulumi.set(__self__, "runtime_environment_name", runtime_environment_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="automationAccountName")
    def automation_account_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the automation account.
        """
        return pulumi.get(self, "automation_account_name")

    @automation_account_name.setter
    def automation_account_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "automation_account_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of an Azure Resource group.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="defaultPackages")
    def default_packages(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        List of Default packages for Environment
        """
        return pulumi.get(self, "default_packages")

    @default_packages.setter
    def default_packages(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "default_packages", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Gets or sets the description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def language(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Language of Runtime Environment
        """
        return pulumi.get(self, "language")

    @language.setter
    def language(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "language", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="runtimeEnvironmentName")
    def runtime_environment_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the Runtime Environment.
        """
        return pulumi.get(self, "runtime_environment_name")

    @runtime_environment_name.setter
    def runtime_environment_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "runtime_environment_name", value)

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
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Version of Language
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "version", value)


@pulumi.type_token("azure-native:automation:RuntimeEnvironment")
class RuntimeEnvironment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 automation_account_name: Optional[pulumi.Input[builtins.str]] = None,
                 default_packages: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 language: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 runtime_environment_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 version: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Definition of the Runtime Environment type.

        Uses Azure REST API version 2023-05-15-preview. In version 2.x of the Azure Native provider, it used API version 2023-05-15-preview.

        Other available API versions: 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] automation_account_name: The name of the automation account.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] default_packages: List of Default packages for Environment
        :param pulumi.Input[builtins.str] description: Gets or sets the description.
        :param pulumi.Input[builtins.str] language: Language of Runtime Environment
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] resource_group_name: Name of an Azure Resource group.
        :param pulumi.Input[builtins.str] runtime_environment_name: The name of the Runtime Environment.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[builtins.str] version: Version of Language
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RuntimeEnvironmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of the Runtime Environment type.

        Uses Azure REST API version 2023-05-15-preview. In version 2.x of the Azure Native provider, it used API version 2023-05-15-preview.

        Other available API versions: 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param RuntimeEnvironmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RuntimeEnvironmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 automation_account_name: Optional[pulumi.Input[builtins.str]] = None,
                 default_packages: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 language: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 runtime_environment_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 version: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RuntimeEnvironmentArgs.__new__(RuntimeEnvironmentArgs)

            if automation_account_name is None and not opts.urn:
                raise TypeError("Missing required property 'automation_account_name'")
            __props__.__dict__["automation_account_name"] = automation_account_name
            __props__.__dict__["default_packages"] = default_packages
            __props__.__dict__["description"] = description
            __props__.__dict__["language"] = language
            __props__.__dict__["location"] = location
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["runtime_environment_name"] = runtime_environment_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["version"] = version
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:automation/v20230515preview:RuntimeEnvironment"), pulumi.Alias(type_="azure-native:automation/v20241023:RuntimeEnvironment")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(RuntimeEnvironment, __self__).__init__(
            'azure-native:automation:RuntimeEnvironment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RuntimeEnvironment':
        """
        Get an existing RuntimeEnvironment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RuntimeEnvironmentArgs.__new__(RuntimeEnvironmentArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["default_packages"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["language"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["version"] = None
        return RuntimeEnvironment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="defaultPackages")
    def default_packages(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        List of Default packages for Environment
        """
        return pulumi.get(self, "default_packages")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Gets or sets the description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def language(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Language of Runtime Environment
        """
        return pulumi.get(self, "language")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

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
    @pulumi.getter
    def version(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Version of Language
        """
        return pulumi.get(self, "version")

