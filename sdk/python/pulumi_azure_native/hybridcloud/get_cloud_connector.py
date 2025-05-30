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
    'GetCloudConnectorResult',
    'AwaitableGetCloudConnectorResult',
    'get_cloud_connector',
    'get_cloud_connector_output',
]

@pulumi.output_type
class GetCloudConnectorResult:
    """
    Cloud Connector resource.
    """
    def __init__(__self__, account_id=None, azure_api_version=None, cloud_type=None, etag=None, id=None, location=None, name=None, provisioning_state=None, system_data=None, tags=None, type=None):
        if account_id and not isinstance(account_id, str):
            raise TypeError("Expected argument 'account_id' to be a str")
        pulumi.set(__self__, "account_id", account_id)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if cloud_type and not isinstance(cloud_type, str):
            raise TypeError("Expected argument 'cloud_type' to be a str")
        pulumi.set(__self__, "cloud_type", cloud_type)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
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
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[builtins.str]:
        """
        Account identifier of the remote cloud.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="cloudType")
    def cloud_type(self) -> Optional[builtins.str]:
        """
        The cloud connector type.
        """
        return pulumi.get(self, "cloud_type")

    @property
    @pulumi.getter
    def etag(self) -> builtins.str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        The provisioning state of the cloud connector resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
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
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetCloudConnectorResult(GetCloudConnectorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudConnectorResult(
            account_id=self.account_id,
            azure_api_version=self.azure_api_version,
            cloud_type=self.cloud_type,
            etag=self.etag,
            id=self.id,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type)


def get_cloud_connector(cloud_connector_name: Optional[builtins.str] = None,
                        resource_group_name: Optional[builtins.str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudConnectorResult:
    """
    Gets the specified cloud connector in a specified resource group.

    Uses Azure REST API version 2023-01-01-preview.


    :param builtins.str cloud_connector_name: The name of the cloud connector resource
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['cloudConnectorName'] = cloud_connector_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:hybridcloud:getCloudConnector', __args__, opts=opts, typ=GetCloudConnectorResult).value

    return AwaitableGetCloudConnectorResult(
        account_id=pulumi.get(__ret__, 'account_id'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        cloud_type=pulumi.get(__ret__, 'cloud_type'),
        etag=pulumi.get(__ret__, 'etag'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_cloud_connector_output(cloud_connector_name: Optional[pulumi.Input[builtins.str]] = None,
                               resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCloudConnectorResult]:
    """
    Gets the specified cloud connector in a specified resource group.

    Uses Azure REST API version 2023-01-01-preview.


    :param builtins.str cloud_connector_name: The name of the cloud connector resource
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['cloudConnectorName'] = cloud_connector_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:hybridcloud:getCloudConnector', __args__, opts=opts, typ=GetCloudConnectorResult)
    return __ret__.apply(lambda __response__: GetCloudConnectorResult(
        account_id=pulumi.get(__response__, 'account_id'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        cloud_type=pulumi.get(__response__, 'cloud_type'),
        etag=pulumi.get(__response__, 'etag'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
