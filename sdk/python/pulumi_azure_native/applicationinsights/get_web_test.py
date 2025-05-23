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
    'GetWebTestResult',
    'AwaitableGetWebTestResult',
    'get_web_test',
    'get_web_test_output',
]

@pulumi.output_type
class GetWebTestResult:
    """
    An Application Insights WebTest definition.
    """
    def __init__(__self__, azure_api_version=None, configuration=None, description=None, enabled=None, frequency=None, id=None, kind=None, location=None, locations=None, name=None, provisioning_state=None, request=None, retry_enabled=None, synthetic_monitor_id=None, tags=None, timeout=None, type=None, validation_rules=None, web_test_kind=None, web_test_name=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if configuration and not isinstance(configuration, dict):
            raise TypeError("Expected argument 'configuration' to be a dict")
        pulumi.set(__self__, "configuration", configuration)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if frequency and not isinstance(frequency, int):
            raise TypeError("Expected argument 'frequency' to be a int")
        pulumi.set(__self__, "frequency", frequency)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if locations and not isinstance(locations, list):
            raise TypeError("Expected argument 'locations' to be a list")
        pulumi.set(__self__, "locations", locations)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if request and not isinstance(request, dict):
            raise TypeError("Expected argument 'request' to be a dict")
        pulumi.set(__self__, "request", request)
        if retry_enabled and not isinstance(retry_enabled, bool):
            raise TypeError("Expected argument 'retry_enabled' to be a bool")
        pulumi.set(__self__, "retry_enabled", retry_enabled)
        if synthetic_monitor_id and not isinstance(synthetic_monitor_id, str):
            raise TypeError("Expected argument 'synthetic_monitor_id' to be a str")
        pulumi.set(__self__, "synthetic_monitor_id", synthetic_monitor_id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if timeout and not isinstance(timeout, int):
            raise TypeError("Expected argument 'timeout' to be a int")
        pulumi.set(__self__, "timeout", timeout)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if validation_rules and not isinstance(validation_rules, dict):
            raise TypeError("Expected argument 'validation_rules' to be a dict")
        pulumi.set(__self__, "validation_rules", validation_rules)
        if web_test_kind and not isinstance(web_test_kind, str):
            raise TypeError("Expected argument 'web_test_kind' to be a str")
        pulumi.set(__self__, "web_test_kind", web_test_kind)
        if web_test_name and not isinstance(web_test_name, str):
            raise TypeError("Expected argument 'web_test_name' to be a str")
        pulumi.set(__self__, "web_test_name", web_test_name)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def configuration(self) -> Optional['outputs.WebTestPropertiesResponseConfiguration']:
        """
        An XML configuration specification for a WebTest.
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        User defined description for this WebTest.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[builtins.bool]:
        """
        Is the test actively being monitored.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def frequency(self) -> Optional[builtins.int]:
        """
        Interval in seconds between test runs for this WebTest. Default value is 300.
        """
        return pulumi.get(self, "frequency")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Azure resource Id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def kind(self) -> Optional[builtins.str]:
        """
        The kind of WebTest that this web test watches. Choices are ping, multistep and standard.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def locations(self) -> Sequence['outputs.WebTestGeolocationResponse']:
        """
        A list of where to physically run the tests from to give global coverage for accessibility of your application.
        """
        return pulumi.get(self, "locations")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Azure resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Current state of this component, whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def request(self) -> Optional['outputs.WebTestPropertiesResponseRequest']:
        """
        The collection of request properties
        """
        return pulumi.get(self, "request")

    @property
    @pulumi.getter(name="retryEnabled")
    def retry_enabled(self) -> Optional[builtins.bool]:
        """
        Allow for retries should this WebTest fail.
        """
        return pulumi.get(self, "retry_enabled")

    @property
    @pulumi.getter(name="syntheticMonitorId")
    def synthetic_monitor_id(self) -> builtins.str:
        """
        Unique ID of this WebTest. This is typically the same value as the Name field.
        """
        return pulumi.get(self, "synthetic_monitor_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def timeout(self) -> Optional[builtins.int]:
        """
        Seconds until this WebTest will timeout and fail. Default value is 30.
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Azure resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="validationRules")
    def validation_rules(self) -> Optional['outputs.WebTestPropertiesResponseValidationRules']:
        """
        The collection of validation rule properties
        """
        return pulumi.get(self, "validation_rules")

    @property
    @pulumi.getter(name="webTestKind")
    def web_test_kind(self) -> builtins.str:
        """
        The kind of web test this is, valid choices are ping, multistep and standard.
        """
        return pulumi.get(self, "web_test_kind")

    @property
    @pulumi.getter(name="webTestName")
    def web_test_name(self) -> builtins.str:
        """
        User defined name if this WebTest.
        """
        return pulumi.get(self, "web_test_name")


class AwaitableGetWebTestResult(GetWebTestResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWebTestResult(
            azure_api_version=self.azure_api_version,
            configuration=self.configuration,
            description=self.description,
            enabled=self.enabled,
            frequency=self.frequency,
            id=self.id,
            kind=self.kind,
            location=self.location,
            locations=self.locations,
            name=self.name,
            provisioning_state=self.provisioning_state,
            request=self.request,
            retry_enabled=self.retry_enabled,
            synthetic_monitor_id=self.synthetic_monitor_id,
            tags=self.tags,
            timeout=self.timeout,
            type=self.type,
            validation_rules=self.validation_rules,
            web_test_kind=self.web_test_kind,
            web_test_name=self.web_test_name)


def get_web_test(resource_group_name: Optional[builtins.str] = None,
                 web_test_name: Optional[builtins.str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWebTestResult:
    """
    Get a specific Application Insights web test definition.

    Uses Azure REST API version 2022-06-15.

    Other available API versions: 2015-05-01, 2018-05-01-preview, 2020-10-05-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native applicationinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str web_test_name: The name of the Application Insights WebTest resource.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['webTestName'] = web_test_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:applicationinsights:getWebTest', __args__, opts=opts, typ=GetWebTestResult).value

    return AwaitableGetWebTestResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        configuration=pulumi.get(__ret__, 'configuration'),
        description=pulumi.get(__ret__, 'description'),
        enabled=pulumi.get(__ret__, 'enabled'),
        frequency=pulumi.get(__ret__, 'frequency'),
        id=pulumi.get(__ret__, 'id'),
        kind=pulumi.get(__ret__, 'kind'),
        location=pulumi.get(__ret__, 'location'),
        locations=pulumi.get(__ret__, 'locations'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        request=pulumi.get(__ret__, 'request'),
        retry_enabled=pulumi.get(__ret__, 'retry_enabled'),
        synthetic_monitor_id=pulumi.get(__ret__, 'synthetic_monitor_id'),
        tags=pulumi.get(__ret__, 'tags'),
        timeout=pulumi.get(__ret__, 'timeout'),
        type=pulumi.get(__ret__, 'type'),
        validation_rules=pulumi.get(__ret__, 'validation_rules'),
        web_test_kind=pulumi.get(__ret__, 'web_test_kind'),
        web_test_name=pulumi.get(__ret__, 'web_test_name'))
def get_web_test_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                        web_test_name: Optional[pulumi.Input[builtins.str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetWebTestResult]:
    """
    Get a specific Application Insights web test definition.

    Uses Azure REST API version 2022-06-15.

    Other available API versions: 2015-05-01, 2018-05-01-preview, 2020-10-05-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native applicationinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str web_test_name: The name of the Application Insights WebTest resource.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['webTestName'] = web_test_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:applicationinsights:getWebTest', __args__, opts=opts, typ=GetWebTestResult)
    return __ret__.apply(lambda __response__: GetWebTestResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        configuration=pulumi.get(__response__, 'configuration'),
        description=pulumi.get(__response__, 'description'),
        enabled=pulumi.get(__response__, 'enabled'),
        frequency=pulumi.get(__response__, 'frequency'),
        id=pulumi.get(__response__, 'id'),
        kind=pulumi.get(__response__, 'kind'),
        location=pulumi.get(__response__, 'location'),
        locations=pulumi.get(__response__, 'locations'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        request=pulumi.get(__response__, 'request'),
        retry_enabled=pulumi.get(__response__, 'retry_enabled'),
        synthetic_monitor_id=pulumi.get(__response__, 'synthetic_monitor_id'),
        tags=pulumi.get(__response__, 'tags'),
        timeout=pulumi.get(__response__, 'timeout'),
        type=pulumi.get(__response__, 'type'),
        validation_rules=pulumi.get(__response__, 'validation_rules'),
        web_test_kind=pulumi.get(__response__, 'web_test_kind'),
        web_test_name=pulumi.get(__response__, 'web_test_name')))
