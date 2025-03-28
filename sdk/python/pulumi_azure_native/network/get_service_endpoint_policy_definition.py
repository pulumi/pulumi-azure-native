# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    'GetServiceEndpointPolicyDefinitionResult',
    'AwaitableGetServiceEndpointPolicyDefinitionResult',
    'get_service_endpoint_policy_definition',
    'get_service_endpoint_policy_definition_output',
]

@pulumi.output_type
class GetServiceEndpointPolicyDefinitionResult:
    """
    Service Endpoint policy definitions.
    """
    def __init__(__self__, description=None, etag=None, id=None, name=None, provisioning_state=None, service=None, service_resources=None, type=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if service and not isinstance(service, str):
            raise TypeError("Expected argument 'service' to be a str")
        pulumi.set(__self__, "service", service)
        if service_resources and not isinstance(service_resources, list):
            raise TypeError("Expected argument 'service_resources' to be a list")
        pulumi.set(__self__, "service_resources", service_resources)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A description for this rule. Restricted to 140 chars.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the service endpoint policy definition resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def service(self) -> Optional[str]:
        """
        Service endpoint name.
        """
        return pulumi.get(self, "service")

    @property
    @pulumi.getter(name="serviceResources")
    def service_resources(self) -> Optional[Sequence[str]]:
        """
        A list of service resources.
        """
        return pulumi.get(self, "service_resources")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetServiceEndpointPolicyDefinitionResult(GetServiceEndpointPolicyDefinitionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceEndpointPolicyDefinitionResult(
            description=self.description,
            etag=self.etag,
            id=self.id,
            name=self.name,
            provisioning_state=self.provisioning_state,
            service=self.service,
            service_resources=self.service_resources,
            type=self.type)


def get_service_endpoint_policy_definition(resource_group_name: Optional[str] = None,
                                           service_endpoint_policy_definition_name: Optional[str] = None,
                                           service_endpoint_policy_name: Optional[str] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceEndpointPolicyDefinitionResult:
    """
    Get the specified service endpoint policy definitions from service endpoint policy.

    Uses Azure REST API version 2023-02-01.

    Other available API versions: 2018-07-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.


    :param str resource_group_name: The name of the resource group.
    :param str service_endpoint_policy_definition_name: The name of the service endpoint policy definition name.
    :param str service_endpoint_policy_name: The name of the service endpoint policy name.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceEndpointPolicyDefinitionName'] = service_endpoint_policy_definition_name
    __args__['serviceEndpointPolicyName'] = service_endpoint_policy_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:network:getServiceEndpointPolicyDefinition', __args__, opts=opts, typ=GetServiceEndpointPolicyDefinitionResult).value

    return AwaitableGetServiceEndpointPolicyDefinitionResult(
        description=pulumi.get(__ret__, 'description'),
        etag=pulumi.get(__ret__, 'etag'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        service=pulumi.get(__ret__, 'service'),
        service_resources=pulumi.get(__ret__, 'service_resources'),
        type=pulumi.get(__ret__, 'type'))
def get_service_endpoint_policy_definition_output(resource_group_name: Optional[pulumi.Input[str]] = None,
                                                  service_endpoint_policy_definition_name: Optional[pulumi.Input[str]] = None,
                                                  service_endpoint_policy_name: Optional[pulumi.Input[str]] = None,
                                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetServiceEndpointPolicyDefinitionResult]:
    """
    Get the specified service endpoint policy definitions from service endpoint policy.

    Uses Azure REST API version 2023-02-01.

    Other available API versions: 2018-07-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.


    :param str resource_group_name: The name of the resource group.
    :param str service_endpoint_policy_definition_name: The name of the service endpoint policy definition name.
    :param str service_endpoint_policy_name: The name of the service endpoint policy name.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceEndpointPolicyDefinitionName'] = service_endpoint_policy_definition_name
    __args__['serviceEndpointPolicyName'] = service_endpoint_policy_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:network:getServiceEndpointPolicyDefinition', __args__, opts=opts, typ=GetServiceEndpointPolicyDefinitionResult)
    return __ret__.apply(lambda __response__: GetServiceEndpointPolicyDefinitionResult(
        description=pulumi.get(__response__, 'description'),
        etag=pulumi.get(__response__, 'etag'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        service=pulumi.get(__response__, 'service'),
        service_resources=pulumi.get(__response__, 'service_resources'),
        type=pulumi.get(__response__, 'type')))
