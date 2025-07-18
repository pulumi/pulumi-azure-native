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

__all__ = [
    'GetWebAppFunctionResult',
    'AwaitableGetWebAppFunctionResult',
    'get_web_app_function',
    'get_web_app_function_output',
]

@pulumi.output_type
class GetWebAppFunctionResult:
    """
    Function information.
    """
    def __init__(__self__, azure_api_version=None, config=None, config_href=None, files=None, function_app_id=None, href=None, id=None, invoke_url_template=None, is_disabled=None, kind=None, language=None, name=None, script_href=None, script_root_path_href=None, secrets_file_href=None, test_data=None, test_data_href=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if config and not isinstance(config, dict):
            raise TypeError("Expected argument 'config' to be a dict")
        pulumi.set(__self__, "config", config)
        if config_href and not isinstance(config_href, str):
            raise TypeError("Expected argument 'config_href' to be a str")
        pulumi.set(__self__, "config_href", config_href)
        if files and not isinstance(files, dict):
            raise TypeError("Expected argument 'files' to be a dict")
        pulumi.set(__self__, "files", files)
        if function_app_id and not isinstance(function_app_id, str):
            raise TypeError("Expected argument 'function_app_id' to be a str")
        pulumi.set(__self__, "function_app_id", function_app_id)
        if href and not isinstance(href, str):
            raise TypeError("Expected argument 'href' to be a str")
        pulumi.set(__self__, "href", href)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if invoke_url_template and not isinstance(invoke_url_template, str):
            raise TypeError("Expected argument 'invoke_url_template' to be a str")
        pulumi.set(__self__, "invoke_url_template", invoke_url_template)
        if is_disabled and not isinstance(is_disabled, bool):
            raise TypeError("Expected argument 'is_disabled' to be a bool")
        pulumi.set(__self__, "is_disabled", is_disabled)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if language and not isinstance(language, str):
            raise TypeError("Expected argument 'language' to be a str")
        pulumi.set(__self__, "language", language)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if script_href and not isinstance(script_href, str):
            raise TypeError("Expected argument 'script_href' to be a str")
        pulumi.set(__self__, "script_href", script_href)
        if script_root_path_href and not isinstance(script_root_path_href, str):
            raise TypeError("Expected argument 'script_root_path_href' to be a str")
        pulumi.set(__self__, "script_root_path_href", script_root_path_href)
        if secrets_file_href and not isinstance(secrets_file_href, str):
            raise TypeError("Expected argument 'secrets_file_href' to be a str")
        pulumi.set(__self__, "secrets_file_href", secrets_file_href)
        if test_data and not isinstance(test_data, str):
            raise TypeError("Expected argument 'test_data' to be a str")
        pulumi.set(__self__, "test_data", test_data)
        if test_data_href and not isinstance(test_data_href, str):
            raise TypeError("Expected argument 'test_data_href' to be a str")
        pulumi.set(__self__, "test_data_href", test_data_href)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def config(self) -> Optional[Any]:
        """
        Config information.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="configHref")
    def config_href(self) -> Optional[builtins.str]:
        """
        Config URI.
        """
        return pulumi.get(self, "config_href")

    @property
    @pulumi.getter
    def files(self) -> Optional[Mapping[str, builtins.str]]:
        """
        File list.
        """
        return pulumi.get(self, "files")

    @property
    @pulumi.getter(name="functionAppId")
    def function_app_id(self) -> Optional[builtins.str]:
        """
        Function App ID.
        """
        return pulumi.get(self, "function_app_id")

    @property
    @pulumi.getter
    def href(self) -> Optional[builtins.str]:
        """
        Function URI.
        """
        return pulumi.get(self, "href")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Resource Id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="invokeUrlTemplate")
    def invoke_url_template(self) -> Optional[builtins.str]:
        """
        The invocation URL
        """
        return pulumi.get(self, "invoke_url_template")

    @property
    @pulumi.getter(name="isDisabled")
    def is_disabled(self) -> Optional[builtins.bool]:
        """
        Gets or sets a value indicating whether the function is disabled
        """
        return pulumi.get(self, "is_disabled")

    @property
    @pulumi.getter
    def kind(self) -> Optional[builtins.str]:
        """
        Kind of resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def language(self) -> Optional[builtins.str]:
        """
        The function language
        """
        return pulumi.get(self, "language")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Resource Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="scriptHref")
    def script_href(self) -> Optional[builtins.str]:
        """
        Script URI.
        """
        return pulumi.get(self, "script_href")

    @property
    @pulumi.getter(name="scriptRootPathHref")
    def script_root_path_href(self) -> Optional[builtins.str]:
        """
        Script root path URI.
        """
        return pulumi.get(self, "script_root_path_href")

    @property
    @pulumi.getter(name="secretsFileHref")
    def secrets_file_href(self) -> Optional[builtins.str]:
        """
        Secrets file URI.
        """
        return pulumi.get(self, "secrets_file_href")

    @property
    @pulumi.getter(name="testData")
    def test_data(self) -> Optional[builtins.str]:
        """
        Test data used when testing via the Azure Portal.
        """
        return pulumi.get(self, "test_data")

    @property
    @pulumi.getter(name="testDataHref")
    def test_data_href(self) -> Optional[builtins.str]:
        """
        Test data URI.
        """
        return pulumi.get(self, "test_data_href")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetWebAppFunctionResult(GetWebAppFunctionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWebAppFunctionResult(
            azure_api_version=self.azure_api_version,
            config=self.config,
            config_href=self.config_href,
            files=self.files,
            function_app_id=self.function_app_id,
            href=self.href,
            id=self.id,
            invoke_url_template=self.invoke_url_template,
            is_disabled=self.is_disabled,
            kind=self.kind,
            language=self.language,
            name=self.name,
            script_href=self.script_href,
            script_root_path_href=self.script_root_path_href,
            secrets_file_href=self.secrets_file_href,
            test_data=self.test_data,
            test_data_href=self.test_data_href,
            type=self.type)


def get_web_app_function(function_name: Optional[builtins.str] = None,
                         name: Optional[builtins.str] = None,
                         resource_group_name: Optional[builtins.str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWebAppFunctionResult:
    """
    Description for Get function information by its ID for web site, or a deployment slot.

    Uses Azure REST API version 2024-04-01.

    Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str function_name: Function name.
    :param builtins.str name: Site name.
    :param builtins.str resource_group_name: Name of the resource group to which the resource belongs.
    """
    __args__ = dict()
    __args__['functionName'] = function_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:web:getWebAppFunction', __args__, opts=opts, typ=GetWebAppFunctionResult).value

    return AwaitableGetWebAppFunctionResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        config=pulumi.get(__ret__, 'config'),
        config_href=pulumi.get(__ret__, 'config_href'),
        files=pulumi.get(__ret__, 'files'),
        function_app_id=pulumi.get(__ret__, 'function_app_id'),
        href=pulumi.get(__ret__, 'href'),
        id=pulumi.get(__ret__, 'id'),
        invoke_url_template=pulumi.get(__ret__, 'invoke_url_template'),
        is_disabled=pulumi.get(__ret__, 'is_disabled'),
        kind=pulumi.get(__ret__, 'kind'),
        language=pulumi.get(__ret__, 'language'),
        name=pulumi.get(__ret__, 'name'),
        script_href=pulumi.get(__ret__, 'script_href'),
        script_root_path_href=pulumi.get(__ret__, 'script_root_path_href'),
        secrets_file_href=pulumi.get(__ret__, 'secrets_file_href'),
        test_data=pulumi.get(__ret__, 'test_data'),
        test_data_href=pulumi.get(__ret__, 'test_data_href'),
        type=pulumi.get(__ret__, 'type'))
def get_web_app_function_output(function_name: Optional[pulumi.Input[builtins.str]] = None,
                                name: Optional[pulumi.Input[builtins.str]] = None,
                                resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetWebAppFunctionResult]:
    """
    Description for Get function information by its ID for web site, or a deployment slot.

    Uses Azure REST API version 2024-04-01.

    Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str function_name: Function name.
    :param builtins.str name: Site name.
    :param builtins.str resource_group_name: Name of the resource group to which the resource belongs.
    """
    __args__ = dict()
    __args__['functionName'] = function_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:web:getWebAppFunction', __args__, opts=opts, typ=GetWebAppFunctionResult)
    return __ret__.apply(lambda __response__: GetWebAppFunctionResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        config=pulumi.get(__response__, 'config'),
        config_href=pulumi.get(__response__, 'config_href'),
        files=pulumi.get(__response__, 'files'),
        function_app_id=pulumi.get(__response__, 'function_app_id'),
        href=pulumi.get(__response__, 'href'),
        id=pulumi.get(__response__, 'id'),
        invoke_url_template=pulumi.get(__response__, 'invoke_url_template'),
        is_disabled=pulumi.get(__response__, 'is_disabled'),
        kind=pulumi.get(__response__, 'kind'),
        language=pulumi.get(__response__, 'language'),
        name=pulumi.get(__response__, 'name'),
        script_href=pulumi.get(__response__, 'script_href'),
        script_root_path_href=pulumi.get(__response__, 'script_root_path_href'),
        secrets_file_href=pulumi.get(__response__, 'secrets_file_href'),
        test_data=pulumi.get(__response__, 'test_data'),
        test_data_href=pulumi.get(__response__, 'test_data_href'),
        type=pulumi.get(__response__, 'type')))
