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
    'GetCustomRecommendationResult',
    'AwaitableGetCustomRecommendationResult',
    'get_custom_recommendation',
    'get_custom_recommendation_output',
]

@pulumi.output_type
class GetCustomRecommendationResult:
    """
    Custom Recommendation
    """
    def __init__(__self__, assessment_key=None, azure_api_version=None, cloud_providers=None, description=None, display_name=None, id=None, name=None, query=None, remediation_description=None, security_issue=None, severity=None, system_data=None, type=None):
        if assessment_key and not isinstance(assessment_key, str):
            raise TypeError("Expected argument 'assessment_key' to be a str")
        pulumi.set(__self__, "assessment_key", assessment_key)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if cloud_providers and not isinstance(cloud_providers, list):
            raise TypeError("Expected argument 'cloud_providers' to be a list")
        pulumi.set(__self__, "cloud_providers", cloud_providers)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if query and not isinstance(query, str):
            raise TypeError("Expected argument 'query' to be a str")
        pulumi.set(__self__, "query", query)
        if remediation_description and not isinstance(remediation_description, str):
            raise TypeError("Expected argument 'remediation_description' to be a str")
        pulumi.set(__self__, "remediation_description", remediation_description)
        if security_issue and not isinstance(security_issue, str):
            raise TypeError("Expected argument 'security_issue' to be a str")
        pulumi.set(__self__, "security_issue", security_issue)
        if severity and not isinstance(severity, str):
            raise TypeError("Expected argument 'severity' to be a str")
        pulumi.set(__self__, "severity", severity)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="assessmentKey")
    def assessment_key(self) -> builtins.str:
        """
        The assessment metadata key used when an assessment is generated for this Recommendation.
        """
        return pulumi.get(self, "assessment_key")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="cloudProviders")
    def cloud_providers(self) -> Optional[Sequence[builtins.str]]:
        """
        List of all standard supported clouds.
        """
        return pulumi.get(self, "cloud_providers")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        The description to relate to the assessments generated by this Recommendation.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[builtins.str]:
        """
        The display name of the assessments generated by this Recommendation.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Resource Id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def query(self) -> Optional[builtins.str]:
        """
        KQL query representing the Recommendation results required.
        """
        return pulumi.get(self, "query")

    @property
    @pulumi.getter(name="remediationDescription")
    def remediation_description(self) -> Optional[builtins.str]:
        """
        The remediation description to relate to the assessments generated by this Recommendation.
        """
        return pulumi.get(self, "remediation_description")

    @property
    @pulumi.getter(name="securityIssue")
    def security_issue(self) -> Optional[builtins.str]:
        """
        The severity to relate to the assessments generated by this Recommendation.
        """
        return pulumi.get(self, "security_issue")

    @property
    @pulumi.getter
    def severity(self) -> Optional[builtins.str]:
        """
        The severity to relate to the assessments generated by this Recommendation.
        """
        return pulumi.get(self, "severity")

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
        Resource type
        """
        return pulumi.get(self, "type")


class AwaitableGetCustomRecommendationResult(GetCustomRecommendationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCustomRecommendationResult(
            assessment_key=self.assessment_key,
            azure_api_version=self.azure_api_version,
            cloud_providers=self.cloud_providers,
            description=self.description,
            display_name=self.display_name,
            id=self.id,
            name=self.name,
            query=self.query,
            remediation_description=self.remediation_description,
            security_issue=self.security_issue,
            severity=self.severity,
            system_data=self.system_data,
            type=self.type)


def get_custom_recommendation(custom_recommendation_name: Optional[builtins.str] = None,
                              scope: Optional[builtins.str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCustomRecommendationResult:
    """
    Get a specific custom recommendation for the requested scope by customRecommendationName

    Uses Azure REST API version 2024-08-01.


    :param builtins.str custom_recommendation_name: Name of the Custom Recommendation.
    :param builtins.str scope: The scope of the custom recommendation. Valid scopes are: management group (format: 'providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: 'subscriptions/{subscriptionId}'), or security connector (format: 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
    """
    __args__ = dict()
    __args__['customRecommendationName'] = custom_recommendation_name
    __args__['scope'] = scope
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:security:getCustomRecommendation', __args__, opts=opts, typ=GetCustomRecommendationResult).value

    return AwaitableGetCustomRecommendationResult(
        assessment_key=pulumi.get(__ret__, 'assessment_key'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        cloud_providers=pulumi.get(__ret__, 'cloud_providers'),
        description=pulumi.get(__ret__, 'description'),
        display_name=pulumi.get(__ret__, 'display_name'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        query=pulumi.get(__ret__, 'query'),
        remediation_description=pulumi.get(__ret__, 'remediation_description'),
        security_issue=pulumi.get(__ret__, 'security_issue'),
        severity=pulumi.get(__ret__, 'severity'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_custom_recommendation_output(custom_recommendation_name: Optional[pulumi.Input[builtins.str]] = None,
                                     scope: Optional[pulumi.Input[builtins.str]] = None,
                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCustomRecommendationResult]:
    """
    Get a specific custom recommendation for the requested scope by customRecommendationName

    Uses Azure REST API version 2024-08-01.


    :param builtins.str custom_recommendation_name: Name of the Custom Recommendation.
    :param builtins.str scope: The scope of the custom recommendation. Valid scopes are: management group (format: 'providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: 'subscriptions/{subscriptionId}'), or security connector (format: 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
    """
    __args__ = dict()
    __args__['customRecommendationName'] = custom_recommendation_name
    __args__['scope'] = scope
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:security:getCustomRecommendation', __args__, opts=opts, typ=GetCustomRecommendationResult)
    return __ret__.apply(lambda __response__: GetCustomRecommendationResult(
        assessment_key=pulumi.get(__response__, 'assessment_key'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        cloud_providers=pulumi.get(__response__, 'cloud_providers'),
        description=pulumi.get(__response__, 'description'),
        display_name=pulumi.get(__response__, 'display_name'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        query=pulumi.get(__response__, 'query'),
        remediation_description=pulumi.get(__response__, 'remediation_description'),
        security_issue=pulumi.get(__response__, 'security_issue'),
        severity=pulumi.get(__response__, 'severity'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
