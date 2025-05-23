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
    'GetMetricsConfigurationResult',
    'AwaitableGetMetricsConfigurationResult',
    'get_metrics_configuration',
    'get_metrics_configuration_output',
]

@pulumi.output_type
class GetMetricsConfigurationResult:
    def __init__(__self__, azure_api_version=None, collection_interval=None, detailed_status=None, detailed_status_message=None, disabled_metrics=None, enabled_metrics=None, etag=None, extended_location=None, id=None, location=None, name=None, provisioning_state=None, system_data=None, tags=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if collection_interval and not isinstance(collection_interval, float):
            raise TypeError("Expected argument 'collection_interval' to be a float")
        pulumi.set(__self__, "collection_interval", collection_interval)
        if detailed_status and not isinstance(detailed_status, str):
            raise TypeError("Expected argument 'detailed_status' to be a str")
        pulumi.set(__self__, "detailed_status", detailed_status)
        if detailed_status_message and not isinstance(detailed_status_message, str):
            raise TypeError("Expected argument 'detailed_status_message' to be a str")
        pulumi.set(__self__, "detailed_status_message", detailed_status_message)
        if disabled_metrics and not isinstance(disabled_metrics, list):
            raise TypeError("Expected argument 'disabled_metrics' to be a list")
        pulumi.set(__self__, "disabled_metrics", disabled_metrics)
        if enabled_metrics and not isinstance(enabled_metrics, list):
            raise TypeError("Expected argument 'enabled_metrics' to be a list")
        pulumi.set(__self__, "enabled_metrics", enabled_metrics)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if extended_location and not isinstance(extended_location, dict):
            raise TypeError("Expected argument 'extended_location' to be a dict")
        pulumi.set(__self__, "extended_location", extended_location)
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
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="collectionInterval")
    def collection_interval(self) -> builtins.float:
        """
        The interval in minutes by which metrics will be collected.
        """
        return pulumi.get(self, "collection_interval")

    @property
    @pulumi.getter(name="detailedStatus")
    def detailed_status(self) -> builtins.str:
        """
        The more detailed status of the metrics configuration.
        """
        return pulumi.get(self, "detailed_status")

    @property
    @pulumi.getter(name="detailedStatusMessage")
    def detailed_status_message(self) -> builtins.str:
        """
        The descriptive message about the current detailed status.
        """
        return pulumi.get(self, "detailed_status_message")

    @property
    @pulumi.getter(name="disabledMetrics")
    def disabled_metrics(self) -> Sequence[builtins.str]:
        """
        The list of metrics that are available for the cluster but disabled at the moment.
        """
        return pulumi.get(self, "disabled_metrics")

    @property
    @pulumi.getter(name="enabledMetrics")
    def enabled_metrics(self) -> Optional[Sequence[builtins.str]]:
        """
        The list of metric names that have been chosen to be enabled in addition to the core set of enabled metrics.
        """
        return pulumi.get(self, "enabled_metrics")

    @property
    @pulumi.getter
    def etag(self) -> builtins.str:
        """
        Resource ETag.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="extendedLocation")
    def extended_location(self) -> 'outputs.ExtendedLocationResponse':
        """
        The extended location of the cluster associated with the resource.
        """
        return pulumi.get(self, "extended_location")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
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
        The provisioning state of the metrics configuration.
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


class AwaitableGetMetricsConfigurationResult(GetMetricsConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMetricsConfigurationResult(
            azure_api_version=self.azure_api_version,
            collection_interval=self.collection_interval,
            detailed_status=self.detailed_status,
            detailed_status_message=self.detailed_status_message,
            disabled_metrics=self.disabled_metrics,
            enabled_metrics=self.enabled_metrics,
            etag=self.etag,
            extended_location=self.extended_location,
            id=self.id,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type)


def get_metrics_configuration(cluster_name: Optional[builtins.str] = None,
                              metrics_configuration_name: Optional[builtins.str] = None,
                              resource_group_name: Optional[builtins.str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMetricsConfigurationResult:
    """
    Get metrics configuration of the provided cluster.

    Uses Azure REST API version 2025-02-01.

    Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str cluster_name: The name of the cluster.
    :param builtins.str metrics_configuration_name: The name of the metrics configuration for the cluster.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['metricsConfigurationName'] = metrics_configuration_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:networkcloud:getMetricsConfiguration', __args__, opts=opts, typ=GetMetricsConfigurationResult).value

    return AwaitableGetMetricsConfigurationResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        collection_interval=pulumi.get(__ret__, 'collection_interval'),
        detailed_status=pulumi.get(__ret__, 'detailed_status'),
        detailed_status_message=pulumi.get(__ret__, 'detailed_status_message'),
        disabled_metrics=pulumi.get(__ret__, 'disabled_metrics'),
        enabled_metrics=pulumi.get(__ret__, 'enabled_metrics'),
        etag=pulumi.get(__ret__, 'etag'),
        extended_location=pulumi.get(__ret__, 'extended_location'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_metrics_configuration_output(cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                                     metrics_configuration_name: Optional[pulumi.Input[builtins.str]] = None,
                                     resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMetricsConfigurationResult]:
    """
    Get metrics configuration of the provided cluster.

    Uses Azure REST API version 2025-02-01.

    Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str cluster_name: The name of the cluster.
    :param builtins.str metrics_configuration_name: The name of the metrics configuration for the cluster.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['metricsConfigurationName'] = metrics_configuration_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:networkcloud:getMetricsConfiguration', __args__, opts=opts, typ=GetMetricsConfigurationResult)
    return __ret__.apply(lambda __response__: GetMetricsConfigurationResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        collection_interval=pulumi.get(__response__, 'collection_interval'),
        detailed_status=pulumi.get(__response__, 'detailed_status'),
        detailed_status_message=pulumi.get(__response__, 'detailed_status_message'),
        disabled_metrics=pulumi.get(__response__, 'disabled_metrics'),
        enabled_metrics=pulumi.get(__response__, 'enabled_metrics'),
        etag=pulumi.get(__response__, 'etag'),
        extended_location=pulumi.get(__response__, 'extended_location'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
