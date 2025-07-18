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

__all__ = ['WebAppFunctionArgs', 'WebAppFunction']

@pulumi.input_type
class WebAppFunctionArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 config: Optional[Any] = None,
                 config_href: Optional[pulumi.Input[builtins.str]] = None,
                 files: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 function_app_id: Optional[pulumi.Input[builtins.str]] = None,
                 function_name: Optional[pulumi.Input[builtins.str]] = None,
                 href: Optional[pulumi.Input[builtins.str]] = None,
                 invoke_url_template: Optional[pulumi.Input[builtins.str]] = None,
                 is_disabled: Optional[pulumi.Input[builtins.bool]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 language: Optional[pulumi.Input[builtins.str]] = None,
                 script_href: Optional[pulumi.Input[builtins.str]] = None,
                 script_root_path_href: Optional[pulumi.Input[builtins.str]] = None,
                 secrets_file_href: Optional[pulumi.Input[builtins.str]] = None,
                 test_data: Optional[pulumi.Input[builtins.str]] = None,
                 test_data_href: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a WebAppFunction resource.
        :param pulumi.Input[builtins.str] name: Site name.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the resource group to which the resource belongs.
        :param Any config: Config information.
        :param pulumi.Input[builtins.str] config_href: Config URI.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] files: File list.
        :param pulumi.Input[builtins.str] function_app_id: Function App ID.
        :param pulumi.Input[builtins.str] function_name: Function name.
        :param pulumi.Input[builtins.str] href: Function URI.
        :param pulumi.Input[builtins.str] invoke_url_template: The invocation URL
        :param pulumi.Input[builtins.bool] is_disabled: Gets or sets a value indicating whether the function is disabled
        :param pulumi.Input[builtins.str] kind: Kind of resource.
        :param pulumi.Input[builtins.str] language: The function language
        :param pulumi.Input[builtins.str] script_href: Script URI.
        :param pulumi.Input[builtins.str] script_root_path_href: Script root path URI.
        :param pulumi.Input[builtins.str] secrets_file_href: Secrets file URI.
        :param pulumi.Input[builtins.str] test_data: Test data used when testing via the Azure Portal.
        :param pulumi.Input[builtins.str] test_data_href: Test data URI.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if config is not None:
            pulumi.set(__self__, "config", config)
        if config_href is not None:
            pulumi.set(__self__, "config_href", config_href)
        if files is not None:
            pulumi.set(__self__, "files", files)
        if function_app_id is not None:
            pulumi.set(__self__, "function_app_id", function_app_id)
        if function_name is not None:
            pulumi.set(__self__, "function_name", function_name)
        if href is not None:
            pulumi.set(__self__, "href", href)
        if invoke_url_template is not None:
            pulumi.set(__self__, "invoke_url_template", invoke_url_template)
        if is_disabled is not None:
            pulumi.set(__self__, "is_disabled", is_disabled)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if language is not None:
            pulumi.set(__self__, "language", language)
        if script_href is not None:
            pulumi.set(__self__, "script_href", script_href)
        if script_root_path_href is not None:
            pulumi.set(__self__, "script_root_path_href", script_root_path_href)
        if secrets_file_href is not None:
            pulumi.set(__self__, "secrets_file_href", secrets_file_href)
        if test_data is not None:
            pulumi.set(__self__, "test_data", test_data)
        if test_data_href is not None:
            pulumi.set(__self__, "test_data_href", test_data_href)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[builtins.str]:
        """
        Site name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the resource group to which the resource belongs.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[Any]:
        """
        Config information.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[Any]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="configHref")
    def config_href(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Config URI.
        """
        return pulumi.get(self, "config_href")

    @config_href.setter
    def config_href(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "config_href", value)

    @property
    @pulumi.getter
    def files(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        File list.
        """
        return pulumi.get(self, "files")

    @files.setter
    def files(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "files", value)

    @property
    @pulumi.getter(name="functionAppId")
    def function_app_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Function App ID.
        """
        return pulumi.get(self, "function_app_id")

    @function_app_id.setter
    def function_app_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "function_app_id", value)

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Function name.
        """
        return pulumi.get(self, "function_name")

    @function_name.setter
    def function_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "function_name", value)

    @property
    @pulumi.getter
    def href(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Function URI.
        """
        return pulumi.get(self, "href")

    @href.setter
    def href(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "href", value)

    @property
    @pulumi.getter(name="invokeUrlTemplate")
    def invoke_url_template(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The invocation URL
        """
        return pulumi.get(self, "invoke_url_template")

    @invoke_url_template.setter
    def invoke_url_template(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "invoke_url_template", value)

    @property
    @pulumi.getter(name="isDisabled")
    def is_disabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Gets or sets a value indicating whether the function is disabled
        """
        return pulumi.get(self, "is_disabled")

    @is_disabled.setter
    def is_disabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "is_disabled", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Kind of resource.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def language(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The function language
        """
        return pulumi.get(self, "language")

    @language.setter
    def language(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "language", value)

    @property
    @pulumi.getter(name="scriptHref")
    def script_href(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Script URI.
        """
        return pulumi.get(self, "script_href")

    @script_href.setter
    def script_href(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "script_href", value)

    @property
    @pulumi.getter(name="scriptRootPathHref")
    def script_root_path_href(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Script root path URI.
        """
        return pulumi.get(self, "script_root_path_href")

    @script_root_path_href.setter
    def script_root_path_href(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "script_root_path_href", value)

    @property
    @pulumi.getter(name="secretsFileHref")
    def secrets_file_href(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Secrets file URI.
        """
        return pulumi.get(self, "secrets_file_href")

    @secrets_file_href.setter
    def secrets_file_href(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "secrets_file_href", value)

    @property
    @pulumi.getter(name="testData")
    def test_data(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Test data used when testing via the Azure Portal.
        """
        return pulumi.get(self, "test_data")

    @test_data.setter
    def test_data(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "test_data", value)

    @property
    @pulumi.getter(name="testDataHref")
    def test_data_href(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Test data URI.
        """
        return pulumi.get(self, "test_data_href")

    @test_data_href.setter
    def test_data_href(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "test_data_href", value)


@pulumi.type_token("azure-native:web:WebAppFunction")
class WebAppFunction(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[Any] = None,
                 config_href: Optional[pulumi.Input[builtins.str]] = None,
                 files: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 function_app_id: Optional[pulumi.Input[builtins.str]] = None,
                 function_name: Optional[pulumi.Input[builtins.str]] = None,
                 href: Optional[pulumi.Input[builtins.str]] = None,
                 invoke_url_template: Optional[pulumi.Input[builtins.str]] = None,
                 is_disabled: Optional[pulumi.Input[builtins.bool]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 language: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 script_href: Optional[pulumi.Input[builtins.str]] = None,
                 script_root_path_href: Optional[pulumi.Input[builtins.str]] = None,
                 secrets_file_href: Optional[pulumi.Input[builtins.str]] = None,
                 test_data: Optional[pulumi.Input[builtins.str]] = None,
                 test_data_href: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Function information.

        Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.

        Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param Any config: Config information.
        :param pulumi.Input[builtins.str] config_href: Config URI.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] files: File list.
        :param pulumi.Input[builtins.str] function_app_id: Function App ID.
        :param pulumi.Input[builtins.str] function_name: Function name.
        :param pulumi.Input[builtins.str] href: Function URI.
        :param pulumi.Input[builtins.str] invoke_url_template: The invocation URL
        :param pulumi.Input[builtins.bool] is_disabled: Gets or sets a value indicating whether the function is disabled
        :param pulumi.Input[builtins.str] kind: Kind of resource.
        :param pulumi.Input[builtins.str] language: The function language
        :param pulumi.Input[builtins.str] name: Site name.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[builtins.str] script_href: Script URI.
        :param pulumi.Input[builtins.str] script_root_path_href: Script root path URI.
        :param pulumi.Input[builtins.str] secrets_file_href: Secrets file URI.
        :param pulumi.Input[builtins.str] test_data: Test data used when testing via the Azure Portal.
        :param pulumi.Input[builtins.str] test_data_href: Test data URI.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebAppFunctionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Function information.

        Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.

        Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param WebAppFunctionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebAppFunctionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[Any] = None,
                 config_href: Optional[pulumi.Input[builtins.str]] = None,
                 files: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 function_app_id: Optional[pulumi.Input[builtins.str]] = None,
                 function_name: Optional[pulumi.Input[builtins.str]] = None,
                 href: Optional[pulumi.Input[builtins.str]] = None,
                 invoke_url_template: Optional[pulumi.Input[builtins.str]] = None,
                 is_disabled: Optional[pulumi.Input[builtins.bool]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 language: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 script_href: Optional[pulumi.Input[builtins.str]] = None,
                 script_root_path_href: Optional[pulumi.Input[builtins.str]] = None,
                 secrets_file_href: Optional[pulumi.Input[builtins.str]] = None,
                 test_data: Optional[pulumi.Input[builtins.str]] = None,
                 test_data_href: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebAppFunctionArgs.__new__(WebAppFunctionArgs)

            __props__.__dict__["config"] = config
            __props__.__dict__["config_href"] = config_href
            __props__.__dict__["files"] = files
            __props__.__dict__["function_app_id"] = function_app_id
            __props__.__dict__["function_name"] = function_name
            __props__.__dict__["href"] = href
            __props__.__dict__["invoke_url_template"] = invoke_url_template
            __props__.__dict__["is_disabled"] = is_disabled
            __props__.__dict__["kind"] = kind
            __props__.__dict__["language"] = language
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["script_href"] = script_href
            __props__.__dict__["script_root_path_href"] = script_root_path_href
            __props__.__dict__["secrets_file_href"] = secrets_file_href
            __props__.__dict__["test_data"] = test_data
            __props__.__dict__["test_data_href"] = test_data_href
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:web/v20160801:WebAppFunction"), pulumi.Alias(type_="azure-native:web/v20180201:WebAppFunction"), pulumi.Alias(type_="azure-native:web/v20181101:WebAppFunction"), pulumi.Alias(type_="azure-native:web/v20190801:WebAppFunction"), pulumi.Alias(type_="azure-native:web/v20200601:WebAppFunction"), pulumi.Alias(type_="azure-native:web/v20200901:WebAppFunction"), pulumi.Alias(type_="azure-native:web/v20201001:WebAppFunction"), pulumi.Alias(type_="azure-native:web/v20201201:WebAppFunction"), pulumi.Alias(type_="azure-native:web/v20210101:WebAppFunction"), pulumi.Alias(type_="azure-native:web/v20210115:WebAppFunction"), pulumi.Alias(type_="azure-native:web/v20210201:WebAppFunction"), pulumi.Alias(type_="azure-native:web/v20210301:WebAppFunction"), pulumi.Alias(type_="azure-native:web/v20220301:WebAppFunction"), pulumi.Alias(type_="azure-native:web/v20220901:WebAppFunction"), pulumi.Alias(type_="azure-native:web/v20230101:WebAppFunction"), pulumi.Alias(type_="azure-native:web/v20231201:WebAppFunction"), pulumi.Alias(type_="azure-native:web/v20240401:WebAppFunction"), pulumi.Alias(type_="azure-native:web/v20241101:WebAppFunction")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(WebAppFunction, __self__).__init__(
            'azure-native:web:WebAppFunction',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WebAppFunction':
        """
        Get an existing WebAppFunction resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WebAppFunctionArgs.__new__(WebAppFunctionArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["config"] = None
        __props__.__dict__["config_href"] = None
        __props__.__dict__["files"] = None
        __props__.__dict__["function_app_id"] = None
        __props__.__dict__["href"] = None
        __props__.__dict__["invoke_url_template"] = None
        __props__.__dict__["is_disabled"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["language"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["script_href"] = None
        __props__.__dict__["script_root_path_href"] = None
        __props__.__dict__["secrets_file_href"] = None
        __props__.__dict__["test_data"] = None
        __props__.__dict__["test_data_href"] = None
        __props__.__dict__["type"] = None
        return WebAppFunction(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[Optional[Any]]:
        """
        Config information.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="configHref")
    def config_href(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Config URI.
        """
        return pulumi.get(self, "config_href")

    @property
    @pulumi.getter
    def files(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        File list.
        """
        return pulumi.get(self, "files")

    @property
    @pulumi.getter(name="functionAppId")
    def function_app_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Function App ID.
        """
        return pulumi.get(self, "function_app_id")

    @property
    @pulumi.getter
    def href(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Function URI.
        """
        return pulumi.get(self, "href")

    @property
    @pulumi.getter(name="invokeUrlTemplate")
    def invoke_url_template(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The invocation URL
        """
        return pulumi.get(self, "invoke_url_template")

    @property
    @pulumi.getter(name="isDisabled")
    def is_disabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Gets or sets a value indicating whether the function is disabled
        """
        return pulumi.get(self, "is_disabled")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Kind of resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def language(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The function language
        """
        return pulumi.get(self, "language")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Resource Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="scriptHref")
    def script_href(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Script URI.
        """
        return pulumi.get(self, "script_href")

    @property
    @pulumi.getter(name="scriptRootPathHref")
    def script_root_path_href(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Script root path URI.
        """
        return pulumi.get(self, "script_root_path_href")

    @property
    @pulumi.getter(name="secretsFileHref")
    def secrets_file_href(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Secrets file URI.
        """
        return pulumi.get(self, "secrets_file_href")

    @property
    @pulumi.getter(name="testData")
    def test_data(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Test data used when testing via the Azure Portal.
        """
        return pulumi.get(self, "test_data")

    @property
    @pulumi.getter(name="testDataHref")
    def test_data_href(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Test data URI.
        """
        return pulumi.get(self, "test_data_href")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

