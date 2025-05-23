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
    'GetRestApiPollerDataConnectorResult',
    'AwaitableGetRestApiPollerDataConnectorResult',
    'get_rest_api_poller_data_connector',
    'get_rest_api_poller_data_connector_output',
]

@pulumi.output_type
class GetRestApiPollerDataConnectorResult:
    """
    Represents Rest Api Poller data connector.
    """
    def __init__(__self__, add_on_attributes=None, auth=None, azure_api_version=None, connector_definition_name=None, data_type=None, dcr_config=None, etag=None, id=None, is_active=None, kind=None, name=None, paging=None, request=None, response=None, system_data=None, type=None):
        if add_on_attributes and not isinstance(add_on_attributes, dict):
            raise TypeError("Expected argument 'add_on_attributes' to be a dict")
        pulumi.set(__self__, "add_on_attributes", add_on_attributes)
        if auth and not isinstance(auth, dict):
            raise TypeError("Expected argument 'auth' to be a dict")
        pulumi.set(__self__, "auth", auth)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if connector_definition_name and not isinstance(connector_definition_name, str):
            raise TypeError("Expected argument 'connector_definition_name' to be a str")
        pulumi.set(__self__, "connector_definition_name", connector_definition_name)
        if data_type and not isinstance(data_type, str):
            raise TypeError("Expected argument 'data_type' to be a str")
        pulumi.set(__self__, "data_type", data_type)
        if dcr_config and not isinstance(dcr_config, dict):
            raise TypeError("Expected argument 'dcr_config' to be a dict")
        pulumi.set(__self__, "dcr_config", dcr_config)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_active and not isinstance(is_active, bool):
            raise TypeError("Expected argument 'is_active' to be a bool")
        pulumi.set(__self__, "is_active", is_active)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if paging and not isinstance(paging, dict):
            raise TypeError("Expected argument 'paging' to be a dict")
        pulumi.set(__self__, "paging", paging)
        if request and not isinstance(request, dict):
            raise TypeError("Expected argument 'request' to be a dict")
        pulumi.set(__self__, "request", request)
        if response and not isinstance(response, dict):
            raise TypeError("Expected argument 'response' to be a dict")
        pulumi.set(__self__, "response", response)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="addOnAttributes")
    def add_on_attributes(self) -> Optional[Mapping[str, builtins.str]]:
        """
        The add on attributes. The key name will become attribute name (a column) and the value will become the attribute value in the payload.
        """
        return pulumi.get(self, "add_on_attributes")

    @property
    @pulumi.getter
    def auth(self) -> Any:
        """
        The a authentication model.
        """
        return pulumi.get(self, "auth")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="connectorDefinitionName")
    def connector_definition_name(self) -> builtins.str:
        """
        The connector definition name (the dataConnectorDefinition resource id).
        """
        return pulumi.get(self, "connector_definition_name")

    @property
    @pulumi.getter(name="dataType")
    def data_type(self) -> Optional[builtins.str]:
        """
        The Log Analytics table destination.
        """
        return pulumi.get(self, "data_type")

    @property
    @pulumi.getter(name="dcrConfig")
    def dcr_config(self) -> Optional['outputs.DCRConfigurationResponse']:
        """
        The DCR related properties.
        """
        return pulumi.get(self, "dcr_config")

    @property
    @pulumi.getter
    def etag(self) -> Optional[builtins.str]:
        """
        Etag of the azure resource
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> Optional[builtins.bool]:
        """
        Indicates whether the connector is active or not.
        """
        return pulumi.get(self, "is_active")

    @property
    @pulumi.getter
    def kind(self) -> builtins.str:
        """
        The kind of the data connector
        Expected value is 'RestApiPoller'.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def paging(self) -> Optional['outputs.RestApiPollerRequestPagingConfigResponse']:
        """
        The paging configuration.
        """
        return pulumi.get(self, "paging")

    @property
    @pulumi.getter
    def request(self) -> 'outputs.RestApiPollerRequestConfigResponse':
        """
        The request configuration.
        """
        return pulumi.get(self, "request")

    @property
    @pulumi.getter
    def response(self) -> Optional['outputs.CcpResponseConfigResponse']:
        """
        The response configuration.
        """
        return pulumi.get(self, "response")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetRestApiPollerDataConnectorResult(GetRestApiPollerDataConnectorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRestApiPollerDataConnectorResult(
            add_on_attributes=self.add_on_attributes,
            auth=self.auth,
            azure_api_version=self.azure_api_version,
            connector_definition_name=self.connector_definition_name,
            data_type=self.data_type,
            dcr_config=self.dcr_config,
            etag=self.etag,
            id=self.id,
            is_active=self.is_active,
            kind=self.kind,
            name=self.name,
            paging=self.paging,
            request=self.request,
            response=self.response,
            system_data=self.system_data,
            type=self.type)


def get_rest_api_poller_data_connector(data_connector_id: Optional[builtins.str] = None,
                                       resource_group_name: Optional[builtins.str] = None,
                                       workspace_name: Optional[builtins.str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRestApiPollerDataConnectorResult:
    """
    Gets a data connector.

    Uses Azure REST API version 2024-09-01.


    :param builtins.str data_connector_id: Connector ID
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['dataConnectorId'] = data_connector_id
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:securityinsights:getRestApiPollerDataConnector', __args__, opts=opts, typ=GetRestApiPollerDataConnectorResult).value

    return AwaitableGetRestApiPollerDataConnectorResult(
        add_on_attributes=pulumi.get(__ret__, 'add_on_attributes'),
        auth=pulumi.get(__ret__, 'auth'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        connector_definition_name=pulumi.get(__ret__, 'connector_definition_name'),
        data_type=pulumi.get(__ret__, 'data_type'),
        dcr_config=pulumi.get(__ret__, 'dcr_config'),
        etag=pulumi.get(__ret__, 'etag'),
        id=pulumi.get(__ret__, 'id'),
        is_active=pulumi.get(__ret__, 'is_active'),
        kind=pulumi.get(__ret__, 'kind'),
        name=pulumi.get(__ret__, 'name'),
        paging=pulumi.get(__ret__, 'paging'),
        request=pulumi.get(__ret__, 'request'),
        response=pulumi.get(__ret__, 'response'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_rest_api_poller_data_connector_output(data_connector_id: Optional[pulumi.Input[builtins.str]] = None,
                                              resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                              workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRestApiPollerDataConnectorResult]:
    """
    Gets a data connector.

    Uses Azure REST API version 2024-09-01.


    :param builtins.str data_connector_id: Connector ID
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['dataConnectorId'] = data_connector_id
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:securityinsights:getRestApiPollerDataConnector', __args__, opts=opts, typ=GetRestApiPollerDataConnectorResult)
    return __ret__.apply(lambda __response__: GetRestApiPollerDataConnectorResult(
        add_on_attributes=pulumi.get(__response__, 'add_on_attributes'),
        auth=pulumi.get(__response__, 'auth'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        connector_definition_name=pulumi.get(__response__, 'connector_definition_name'),
        data_type=pulumi.get(__response__, 'data_type'),
        dcr_config=pulumi.get(__response__, 'dcr_config'),
        etag=pulumi.get(__response__, 'etag'),
        id=pulumi.get(__response__, 'id'),
        is_active=pulumi.get(__response__, 'is_active'),
        kind=pulumi.get(__response__, 'kind'),
        name=pulumi.get(__response__, 'name'),
        paging=pulumi.get(__response__, 'paging'),
        request=pulumi.get(__response__, 'request'),
        response=pulumi.get(__response__, 'response'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
