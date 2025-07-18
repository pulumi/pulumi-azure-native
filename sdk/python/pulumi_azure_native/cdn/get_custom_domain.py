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
    'GetCustomDomainResult',
    'AwaitableGetCustomDomainResult',
    'get_custom_domain',
    'get_custom_domain_output',
]

@pulumi.output_type
class GetCustomDomainResult:
    """
    Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
    """
    def __init__(__self__, azure_api_version=None, custom_https_parameters=None, custom_https_provisioning_state=None, custom_https_provisioning_substate=None, host_name=None, id=None, name=None, provisioning_state=None, resource_state=None, system_data=None, type=None, validation_data=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if custom_https_parameters and not isinstance(custom_https_parameters, dict):
            raise TypeError("Expected argument 'custom_https_parameters' to be a dict")
        pulumi.set(__self__, "custom_https_parameters", custom_https_parameters)
        if custom_https_provisioning_state and not isinstance(custom_https_provisioning_state, str):
            raise TypeError("Expected argument 'custom_https_provisioning_state' to be a str")
        pulumi.set(__self__, "custom_https_provisioning_state", custom_https_provisioning_state)
        if custom_https_provisioning_substate and not isinstance(custom_https_provisioning_substate, str):
            raise TypeError("Expected argument 'custom_https_provisioning_substate' to be a str")
        pulumi.set(__self__, "custom_https_provisioning_substate", custom_https_provisioning_substate)
        if host_name and not isinstance(host_name, str):
            raise TypeError("Expected argument 'host_name' to be a str")
        pulumi.set(__self__, "host_name", host_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if resource_state and not isinstance(resource_state, str):
            raise TypeError("Expected argument 'resource_state' to be a str")
        pulumi.set(__self__, "resource_state", resource_state)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if validation_data and not isinstance(validation_data, str):
            raise TypeError("Expected argument 'validation_data' to be a str")
        pulumi.set(__self__, "validation_data", validation_data)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="customHttpsParameters")
    def custom_https_parameters(self) -> Optional[Any]:
        """
        Certificate parameters for securing custom HTTPS
        """
        return pulumi.get(self, "custom_https_parameters")

    @property
    @pulumi.getter(name="customHttpsProvisioningState")
    def custom_https_provisioning_state(self) -> builtins.str:
        """
        Provisioning status of the custom domain.
        """
        return pulumi.get(self, "custom_https_provisioning_state")

    @property
    @pulumi.getter(name="customHttpsProvisioningSubstate")
    def custom_https_provisioning_substate(self) -> builtins.str:
        """
        Provisioning substate shows the progress of custom HTTPS enabling/disabling process step by step.
        """
        return pulumi.get(self, "custom_https_provisioning_substate")

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> builtins.str:
        """
        The host name of the custom domain. Must be a domain name.
        """
        return pulumi.get(self, "host_name")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Provisioning status of Custom Https of the custom domain.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceState")
    def resource_state(self) -> builtins.str:
        """
        Resource status of the custom domain.
        """
        return pulumi.get(self, "resource_state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Read only system data
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="validationData")
    def validation_data(self) -> Optional[builtins.str]:
        """
        Special validation or data may be required when delivering CDN to some regions due to local compliance reasons. E.g. ICP license number of a custom domain is required to deliver content in China.
        """
        return pulumi.get(self, "validation_data")


class AwaitableGetCustomDomainResult(GetCustomDomainResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCustomDomainResult(
            azure_api_version=self.azure_api_version,
            custom_https_parameters=self.custom_https_parameters,
            custom_https_provisioning_state=self.custom_https_provisioning_state,
            custom_https_provisioning_substate=self.custom_https_provisioning_substate,
            host_name=self.host_name,
            id=self.id,
            name=self.name,
            provisioning_state=self.provisioning_state,
            resource_state=self.resource_state,
            system_data=self.system_data,
            type=self.type,
            validation_data=self.validation_data)


def get_custom_domain(custom_domain_name: Optional[builtins.str] = None,
                      endpoint_name: Optional[builtins.str] = None,
                      profile_name: Optional[builtins.str] = None,
                      resource_group_name: Optional[builtins.str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCustomDomainResult:
    """
    Gets an existing custom domain within an endpoint.

    Uses Azure REST API version 2024-09-01.

    Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str custom_domain_name: Name of the custom domain within an endpoint.
    :param builtins.str endpoint_name: Name of the endpoint under the profile which is unique globally.
    :param builtins.str profile_name: Name of the CDN profile which is unique within the resource group.
    :param builtins.str resource_group_name: Name of the Resource group within the Azure subscription.
    """
    __args__ = dict()
    __args__['customDomainName'] = custom_domain_name
    __args__['endpointName'] = endpoint_name
    __args__['profileName'] = profile_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:cdn:getCustomDomain', __args__, opts=opts, typ=GetCustomDomainResult).value

    return AwaitableGetCustomDomainResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        custom_https_parameters=pulumi.get(__ret__, 'custom_https_parameters'),
        custom_https_provisioning_state=pulumi.get(__ret__, 'custom_https_provisioning_state'),
        custom_https_provisioning_substate=pulumi.get(__ret__, 'custom_https_provisioning_substate'),
        host_name=pulumi.get(__ret__, 'host_name'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        resource_state=pulumi.get(__ret__, 'resource_state'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'),
        validation_data=pulumi.get(__ret__, 'validation_data'))
def get_custom_domain_output(custom_domain_name: Optional[pulumi.Input[builtins.str]] = None,
                             endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                             profile_name: Optional[pulumi.Input[builtins.str]] = None,
                             resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCustomDomainResult]:
    """
    Gets an existing custom domain within an endpoint.

    Uses Azure REST API version 2024-09-01.

    Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str custom_domain_name: Name of the custom domain within an endpoint.
    :param builtins.str endpoint_name: Name of the endpoint under the profile which is unique globally.
    :param builtins.str profile_name: Name of the CDN profile which is unique within the resource group.
    :param builtins.str resource_group_name: Name of the Resource group within the Azure subscription.
    """
    __args__ = dict()
    __args__['customDomainName'] = custom_domain_name
    __args__['endpointName'] = endpoint_name
    __args__['profileName'] = profile_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:cdn:getCustomDomain', __args__, opts=opts, typ=GetCustomDomainResult)
    return __ret__.apply(lambda __response__: GetCustomDomainResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        custom_https_parameters=pulumi.get(__response__, 'custom_https_parameters'),
        custom_https_provisioning_state=pulumi.get(__response__, 'custom_https_provisioning_state'),
        custom_https_provisioning_substate=pulumi.get(__response__, 'custom_https_provisioning_substate'),
        host_name=pulumi.get(__response__, 'host_name'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        resource_state=pulumi.get(__response__, 'resource_state'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type'),
        validation_data=pulumi.get(__response__, 'validation_data')))
