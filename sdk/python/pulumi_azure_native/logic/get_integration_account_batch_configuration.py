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

__all__ = [
    'GetIntegrationAccountBatchConfigurationResult',
    'AwaitableGetIntegrationAccountBatchConfigurationResult',
    'get_integration_account_batch_configuration',
    'get_integration_account_batch_configuration_output',
]

@pulumi.output_type
class GetIntegrationAccountBatchConfigurationResult:
    """
    The batch configuration resource definition.
    """
    def __init__(__self__, azure_api_version=None, id=None, location=None, name=None, properties=None, tags=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
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
    def id(self) -> builtins.str:
        """
        The resource id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> Optional[builtins.str]:
        """
        The resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Gets the resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.BatchConfigurationPropertiesResponse':
        """
        The batch configuration properties.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        The resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Gets the resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetIntegrationAccountBatchConfigurationResult(GetIntegrationAccountBatchConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIntegrationAccountBatchConfigurationResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            location=self.location,
            name=self.name,
            properties=self.properties,
            tags=self.tags,
            type=self.type)


def get_integration_account_batch_configuration(batch_configuration_name: Optional[builtins.str] = None,
                                                integration_account_name: Optional[builtins.str] = None,
                                                resource_group_name: Optional[builtins.str] = None,
                                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIntegrationAccountBatchConfigurationResult:
    """
    Get a batch configuration for an integration account.

    Uses Azure REST API version 2019-05-01.

    Other available API versions: 2016-06-01, 2018-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native logic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str batch_configuration_name: The batch configuration name.
    :param builtins.str integration_account_name: The integration account name.
    :param builtins.str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['batchConfigurationName'] = batch_configuration_name
    __args__['integrationAccountName'] = integration_account_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:logic:getIntegrationAccountBatchConfiguration', __args__, opts=opts, typ=GetIntegrationAccountBatchConfigurationResult).value

    return AwaitableGetIntegrationAccountBatchConfigurationResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        properties=pulumi.get(__ret__, 'properties'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_integration_account_batch_configuration_output(batch_configuration_name: Optional[pulumi.Input[builtins.str]] = None,
                                                       integration_account_name: Optional[pulumi.Input[builtins.str]] = None,
                                                       resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetIntegrationAccountBatchConfigurationResult]:
    """
    Get a batch configuration for an integration account.

    Uses Azure REST API version 2019-05-01.

    Other available API versions: 2016-06-01, 2018-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native logic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str batch_configuration_name: The batch configuration name.
    :param builtins.str integration_account_name: The integration account name.
    :param builtins.str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['batchConfigurationName'] = batch_configuration_name
    __args__['integrationAccountName'] = integration_account_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:logic:getIntegrationAccountBatchConfiguration', __args__, opts=opts, typ=GetIntegrationAccountBatchConfigurationResult)
    return __ret__.apply(lambda __response__: GetIntegrationAccountBatchConfigurationResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        properties=pulumi.get(__response__, 'properties'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
