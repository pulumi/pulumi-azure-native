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
    'GetApiGatewayResult',
    'AwaitableGetApiGatewayResult',
    'get_api_gateway',
    'get_api_gateway_output',
]

@pulumi.output_type
class GetApiGatewayResult:
    """
    A single API Management gateway resource in List or Get response.
    """
    def __init__(__self__, azure_api_version=None, backend=None, configuration_api=None, created_at_utc=None, etag=None, frontend=None, id=None, location=None, name=None, provisioning_state=None, sku=None, system_data=None, tags=None, target_provisioning_state=None, type=None, virtual_network_type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if backend and not isinstance(backend, dict):
            raise TypeError("Expected argument 'backend' to be a dict")
        pulumi.set(__self__, "backend", backend)
        if configuration_api and not isinstance(configuration_api, dict):
            raise TypeError("Expected argument 'configuration_api' to be a dict")
        pulumi.set(__self__, "configuration_api", configuration_api)
        if created_at_utc and not isinstance(created_at_utc, str):
            raise TypeError("Expected argument 'created_at_utc' to be a str")
        pulumi.set(__self__, "created_at_utc", created_at_utc)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if frontend and not isinstance(frontend, dict):
            raise TypeError("Expected argument 'frontend' to be a dict")
        pulumi.set(__self__, "frontend", frontend)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if target_provisioning_state and not isinstance(target_provisioning_state, str):
            raise TypeError("Expected argument 'target_provisioning_state' to be a str")
        pulumi.set(__self__, "target_provisioning_state", target_provisioning_state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if virtual_network_type and not isinstance(virtual_network_type, str):
            raise TypeError("Expected argument 'virtual_network_type' to be a str")
        pulumi.set(__self__, "virtual_network_type", virtual_network_type)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def backend(self) -> Optional['outputs.BackendConfigurationResponse']:
        """
        Information regarding how the gateway should integrate with backend systems.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="configurationApi")
    def configuration_api(self) -> Optional['outputs.GatewayConfigurationApiResponse']:
        """
        Information regarding the Configuration API of the API Management gateway. This is only applicable for API gateway with Standard SKU.
        """
        return pulumi.get(self, "configuration_api")

    @property
    @pulumi.getter(name="createdAtUtc")
    def created_at_utc(self) -> builtins.str:
        """
        Creation UTC date of the API Management gateway.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
        """
        return pulumi.get(self, "created_at_utc")

    @property
    @pulumi.getter
    def etag(self) -> builtins.str:
        """
        ETag of the resource.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def frontend(self) -> Optional['outputs.FrontendConfigurationResponse']:
        """
        Information regarding how the gateway should be exposed.
        """
        return pulumi.get(self, "frontend")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

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
        The current provisioning state of the API Management gateway which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> 'outputs.ApiManagementGatewaySkuPropertiesResponse':
        """
        SKU properties of the API Management gateway.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Metadata pertaining to creation and last modification of the resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetProvisioningState")
    def target_provisioning_state(self) -> builtins.str:
        """
        The provisioning state of the API Management gateway, which is targeted by the long running operation started on the gateway.
        """
        return pulumi.get(self, "target_provisioning_state")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Resource type for API Management resource is set to Microsoft.ApiManagement.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualNetworkType")
    def virtual_network_type(self) -> Optional[builtins.str]:
        """
        The type of VPN in which API Management gateway needs to be configured in. 
        """
        return pulumi.get(self, "virtual_network_type")


class AwaitableGetApiGatewayResult(GetApiGatewayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApiGatewayResult(
            azure_api_version=self.azure_api_version,
            backend=self.backend,
            configuration_api=self.configuration_api,
            created_at_utc=self.created_at_utc,
            etag=self.etag,
            frontend=self.frontend,
            id=self.id,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            sku=self.sku,
            system_data=self.system_data,
            tags=self.tags,
            target_provisioning_state=self.target_provisioning_state,
            type=self.type,
            virtual_network_type=self.virtual_network_type)


def get_api_gateway(gateway_name: Optional[builtins.str] = None,
                    resource_group_name: Optional[builtins.str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApiGatewayResult:
    """
    Gets an API Management gateway resource description.

    Uses Azure REST API version 2024-06-01-preview.

    Other available API versions: 2023-09-01-preview, 2024-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str gateway_name: The name of the API Management gateway.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['gatewayName'] = gateway_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:apimanagement:getApiGateway', __args__, opts=opts, typ=GetApiGatewayResult).value

    return AwaitableGetApiGatewayResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        backend=pulumi.get(__ret__, 'backend'),
        configuration_api=pulumi.get(__ret__, 'configuration_api'),
        created_at_utc=pulumi.get(__ret__, 'created_at_utc'),
        etag=pulumi.get(__ret__, 'etag'),
        frontend=pulumi.get(__ret__, 'frontend'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        sku=pulumi.get(__ret__, 'sku'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        target_provisioning_state=pulumi.get(__ret__, 'target_provisioning_state'),
        type=pulumi.get(__ret__, 'type'),
        virtual_network_type=pulumi.get(__ret__, 'virtual_network_type'))
def get_api_gateway_output(gateway_name: Optional[pulumi.Input[builtins.str]] = None,
                           resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetApiGatewayResult]:
    """
    Gets an API Management gateway resource description.

    Uses Azure REST API version 2024-06-01-preview.

    Other available API versions: 2023-09-01-preview, 2024-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str gateway_name: The name of the API Management gateway.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['gatewayName'] = gateway_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:apimanagement:getApiGateway', __args__, opts=opts, typ=GetApiGatewayResult)
    return __ret__.apply(lambda __response__: GetApiGatewayResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        backend=pulumi.get(__response__, 'backend'),
        configuration_api=pulumi.get(__response__, 'configuration_api'),
        created_at_utc=pulumi.get(__response__, 'created_at_utc'),
        etag=pulumi.get(__response__, 'etag'),
        frontend=pulumi.get(__response__, 'frontend'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        sku=pulumi.get(__response__, 'sku'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        target_provisioning_state=pulumi.get(__response__, 'target_provisioning_state'),
        type=pulumi.get(__response__, 'type'),
        virtual_network_type=pulumi.get(__response__, 'virtual_network_type')))
