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
    'GetDeploymentSafeguardResult',
    'AwaitableGetDeploymentSafeguardResult',
    'get_deployment_safeguard',
    'get_deployment_safeguard_output',
]

@pulumi.output_type
class GetDeploymentSafeguardResult:
    """
    Deployment Safeguards
    """
    def __init__(__self__, azure_api_version=None, e_tag=None, excluded_namespaces=None, id=None, level=None, name=None, provisioning_state=None, system_data=None, system_excluded_namespaces=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if e_tag and not isinstance(e_tag, str):
            raise TypeError("Expected argument 'e_tag' to be a str")
        pulumi.set(__self__, "e_tag", e_tag)
        if excluded_namespaces and not isinstance(excluded_namespaces, list):
            raise TypeError("Expected argument 'excluded_namespaces' to be a list")
        pulumi.set(__self__, "excluded_namespaces", excluded_namespaces)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if level and not isinstance(level, str):
            raise TypeError("Expected argument 'level' to be a str")
        pulumi.set(__self__, "level", level)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if system_excluded_namespaces and not isinstance(system_excluded_namespaces, list):
            raise TypeError("Expected argument 'system_excluded_namespaces' to be a list")
        pulumi.set(__self__, "system_excluded_namespaces", system_excluded_namespaces)
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
    @pulumi.getter(name="eTag")
    def e_tag(self) -> builtins.str:
        """
        If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
        """
        return pulumi.get(self, "e_tag")

    @property
    @pulumi.getter(name="excludedNamespaces")
    def excluded_namespaces(self) -> Optional[Sequence[builtins.str]]:
        """
        User defined list of namespaces to exclude from Deployment Safeguards. Deployments in these namespaces will not be checked against any safeguards
        """
        return pulumi.get(self, "excluded_namespaces")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def level(self) -> builtins.str:
        """
        The deployment safeguards level. Possible values are Warn and Enforce
        """
        return pulumi.get(self, "level")

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
        Provisioning State
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
    @pulumi.getter(name="systemExcludedNamespaces")
    def system_excluded_namespaces(self) -> Sequence[builtins.str]:
        """
        System defined list of namespaces excluded from Deployment Safeguards. These are determined by the underlying provider (such as AKS), and cannot be changed. Deployments in these namespaces will not be checked
        """
        return pulumi.get(self, "system_excluded_namespaces")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetDeploymentSafeguardResult(GetDeploymentSafeguardResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeploymentSafeguardResult(
            azure_api_version=self.azure_api_version,
            e_tag=self.e_tag,
            excluded_namespaces=self.excluded_namespaces,
            id=self.id,
            level=self.level,
            name=self.name,
            provisioning_state=self.provisioning_state,
            system_data=self.system_data,
            system_excluded_namespaces=self.system_excluded_namespaces,
            type=self.type)


def get_deployment_safeguard(resource_uri: Optional[builtins.str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeploymentSafeguardResult:
    """
    Fetch a deployment safeguard by name

    Uses Azure REST API version 2025-04-02-preview.

    Other available API versions: 2025-04-01, 2025-05-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_uri: The fully qualified Azure Resource manager identifier of the resource.
    """
    __args__ = dict()
    __args__['resourceUri'] = resource_uri
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:containerservice:getDeploymentSafeguard', __args__, opts=opts, typ=GetDeploymentSafeguardResult).value

    return AwaitableGetDeploymentSafeguardResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        e_tag=pulumi.get(__ret__, 'e_tag'),
        excluded_namespaces=pulumi.get(__ret__, 'excluded_namespaces'),
        id=pulumi.get(__ret__, 'id'),
        level=pulumi.get(__ret__, 'level'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        system_data=pulumi.get(__ret__, 'system_data'),
        system_excluded_namespaces=pulumi.get(__ret__, 'system_excluded_namespaces'),
        type=pulumi.get(__ret__, 'type'))
def get_deployment_safeguard_output(resource_uri: Optional[pulumi.Input[builtins.str]] = None,
                                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDeploymentSafeguardResult]:
    """
    Fetch a deployment safeguard by name

    Uses Azure REST API version 2025-04-02-preview.

    Other available API versions: 2025-04-01, 2025-05-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_uri: The fully qualified Azure Resource manager identifier of the resource.
    """
    __args__ = dict()
    __args__['resourceUri'] = resource_uri
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:containerservice:getDeploymentSafeguard', __args__, opts=opts, typ=GetDeploymentSafeguardResult)
    return __ret__.apply(lambda __response__: GetDeploymentSafeguardResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        e_tag=pulumi.get(__response__, 'e_tag'),
        excluded_namespaces=pulumi.get(__response__, 'excluded_namespaces'),
        id=pulumi.get(__response__, 'id'),
        level=pulumi.get(__response__, 'level'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        system_data=pulumi.get(__response__, 'system_data'),
        system_excluded_namespaces=pulumi.get(__response__, 'system_excluded_namespaces'),
        type=pulumi.get(__response__, 'type')))
