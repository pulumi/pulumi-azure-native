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
    'GetRouteFilterRuleResult',
    'AwaitableGetRouteFilterRuleResult',
    'get_route_filter_rule',
    'get_route_filter_rule_output',
]

@pulumi.output_type
class GetRouteFilterRuleResult:
    """
    Route Filter Rule Resource.
    """
    def __init__(__self__, access=None, azure_api_version=None, communities=None, etag=None, id=None, location=None, name=None, provisioning_state=None, route_filter_rule_type=None):
        if access and not isinstance(access, str):
            raise TypeError("Expected argument 'access' to be a str")
        pulumi.set(__self__, "access", access)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if communities and not isinstance(communities, list):
            raise TypeError("Expected argument 'communities' to be a list")
        pulumi.set(__self__, "communities", communities)
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
        if route_filter_rule_type and not isinstance(route_filter_rule_type, str):
            raise TypeError("Expected argument 'route_filter_rule_type' to be a str")
        pulumi.set(__self__, "route_filter_rule_type", route_filter_rule_type)

    @property
    @pulumi.getter
    def access(self) -> builtins.str:
        """
        The access type of the rule.
        """
        return pulumi.get(self, "access")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def communities(self) -> Sequence[builtins.str]:
        """
        The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020'].
        """
        return pulumi.get(self, "communities")

    @property
    @pulumi.getter
    def etag(self) -> builtins.str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> Optional[builtins.str]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        The name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        The provisioning state of the route filter rule resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="routeFilterRuleType")
    def route_filter_rule_type(self) -> builtins.str:
        """
        The rule type of the rule.
        """
        return pulumi.get(self, "route_filter_rule_type")


class AwaitableGetRouteFilterRuleResult(GetRouteFilterRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouteFilterRuleResult(
            access=self.access,
            azure_api_version=self.azure_api_version,
            communities=self.communities,
            etag=self.etag,
            id=self.id,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            route_filter_rule_type=self.route_filter_rule_type)


def get_route_filter_rule(resource_group_name: Optional[builtins.str] = None,
                          route_filter_name: Optional[builtins.str] = None,
                          rule_name: Optional[builtins.str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouteFilterRuleResult:
    """
    Gets the specified rule from a route filter.

    Uses Azure REST API version 2024-05-01.

    Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group.
    :param builtins.str route_filter_name: The name of the route filter.
    :param builtins.str rule_name: The name of the rule.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['routeFilterName'] = route_filter_name
    __args__['ruleName'] = rule_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:network:getRouteFilterRule', __args__, opts=opts, typ=GetRouteFilterRuleResult).value

    return AwaitableGetRouteFilterRuleResult(
        access=pulumi.get(__ret__, 'access'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        communities=pulumi.get(__ret__, 'communities'),
        etag=pulumi.get(__ret__, 'etag'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        route_filter_rule_type=pulumi.get(__ret__, 'route_filter_rule_type'))
def get_route_filter_rule_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                 route_filter_name: Optional[pulumi.Input[builtins.str]] = None,
                                 rule_name: Optional[pulumi.Input[builtins.str]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRouteFilterRuleResult]:
    """
    Gets the specified rule from a route filter.

    Uses Azure REST API version 2024-05-01.

    Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group.
    :param builtins.str route_filter_name: The name of the route filter.
    :param builtins.str rule_name: The name of the rule.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['routeFilterName'] = route_filter_name
    __args__['ruleName'] = rule_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:network:getRouteFilterRule', __args__, opts=opts, typ=GetRouteFilterRuleResult)
    return __ret__.apply(lambda __response__: GetRouteFilterRuleResult(
        access=pulumi.get(__response__, 'access'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        communities=pulumi.get(__response__, 'communities'),
        etag=pulumi.get(__response__, 'etag'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        route_filter_rule_type=pulumi.get(__response__, 'route_filter_rule_type')))
