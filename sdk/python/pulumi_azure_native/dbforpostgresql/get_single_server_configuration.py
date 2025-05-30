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
    'GetSingleServerConfigurationResult',
    'AwaitableGetSingleServerConfigurationResult',
    'get_single_server_configuration',
    'get_single_server_configuration_output',
]

@pulumi.output_type
class GetSingleServerConfigurationResult:
    """
    Represents a Configuration.
    """
    def __init__(__self__, allowed_values=None, azure_api_version=None, data_type=None, default_value=None, description=None, id=None, name=None, source=None, type=None, value=None):
        if allowed_values and not isinstance(allowed_values, str):
            raise TypeError("Expected argument 'allowed_values' to be a str")
        pulumi.set(__self__, "allowed_values", allowed_values)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if data_type and not isinstance(data_type, str):
            raise TypeError("Expected argument 'data_type' to be a str")
        pulumi.set(__self__, "data_type", data_type)
        if default_value and not isinstance(default_value, str):
            raise TypeError("Expected argument 'default_value' to be a str")
        pulumi.set(__self__, "default_value", default_value)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if source and not isinstance(source, str):
            raise TypeError("Expected argument 'source' to be a str")
        pulumi.set(__self__, "source", source)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="allowedValues")
    def allowed_values(self) -> builtins.str:
        """
        Allowed values of the configuration.
        """
        return pulumi.get(self, "allowed_values")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="dataType")
    def data_type(self) -> builtins.str:
        """
        Data type of the configuration.
        """
        return pulumi.get(self, "data_type")

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> builtins.str:
        """
        Default value of the configuration.
        """
        return pulumi.get(self, "default_value")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        Description of the configuration.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def source(self) -> Optional[builtins.str]:
        """
        Source of the configuration.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> Optional[builtins.str]:
        """
        Value of the configuration.
        """
        return pulumi.get(self, "value")


class AwaitableGetSingleServerConfigurationResult(GetSingleServerConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSingleServerConfigurationResult(
            allowed_values=self.allowed_values,
            azure_api_version=self.azure_api_version,
            data_type=self.data_type,
            default_value=self.default_value,
            description=self.description,
            id=self.id,
            name=self.name,
            source=self.source,
            type=self.type,
            value=self.value)


def get_single_server_configuration(configuration_name: Optional[builtins.str] = None,
                                    resource_group_name: Optional[builtins.str] = None,
                                    server_name: Optional[builtins.str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSingleServerConfigurationResult:
    """
    Gets information about a configuration of server.

    Uses Azure REST API version 2017-12-01.


    :param builtins.str configuration_name: The name of the server configuration.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['configurationName'] = configuration_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:dbforpostgresql:getSingleServerConfiguration', __args__, opts=opts, typ=GetSingleServerConfigurationResult).value

    return AwaitableGetSingleServerConfigurationResult(
        allowed_values=pulumi.get(__ret__, 'allowed_values'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        data_type=pulumi.get(__ret__, 'data_type'),
        default_value=pulumi.get(__ret__, 'default_value'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        source=pulumi.get(__ret__, 'source'),
        type=pulumi.get(__ret__, 'type'),
        value=pulumi.get(__ret__, 'value'))
def get_single_server_configuration_output(configuration_name: Optional[pulumi.Input[builtins.str]] = None,
                                           resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                           server_name: Optional[pulumi.Input[builtins.str]] = None,
                                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSingleServerConfigurationResult]:
    """
    Gets information about a configuration of server.

    Uses Azure REST API version 2017-12-01.


    :param builtins.str configuration_name: The name of the server configuration.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['configurationName'] = configuration_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:dbforpostgresql:getSingleServerConfiguration', __args__, opts=opts, typ=GetSingleServerConfigurationResult)
    return __ret__.apply(lambda __response__: GetSingleServerConfigurationResult(
        allowed_values=pulumi.get(__response__, 'allowed_values'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        data_type=pulumi.get(__response__, 'data_type'),
        default_value=pulumi.get(__response__, 'default_value'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        source=pulumi.get(__response__, 'source'),
        type=pulumi.get(__response__, 'type'),
        value=pulumi.get(__response__, 'value')))
