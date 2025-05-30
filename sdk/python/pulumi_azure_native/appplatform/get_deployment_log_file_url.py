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
    'GetDeploymentLogFileUrlResult',
    'AwaitableGetDeploymentLogFileUrlResult',
    'get_deployment_log_file_url',
    'get_deployment_log_file_url_output',
]

@pulumi.output_type
class GetDeploymentLogFileUrlResult:
    """
    Log file URL payload
    """
    def __init__(__self__, url=None):
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def url(self) -> builtins.str:
        """
        URL of the log file
        """
        return pulumi.get(self, "url")


class AwaitableGetDeploymentLogFileUrlResult(GetDeploymentLogFileUrlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeploymentLogFileUrlResult(
            url=self.url)


def get_deployment_log_file_url(app_name: Optional[builtins.str] = None,
                                deployment_name: Optional[builtins.str] = None,
                                resource_group_name: Optional[builtins.str] = None,
                                service_name: Optional[builtins.str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeploymentLogFileUrlResult:
    """
    Get deployment log file URL

    Uses Azure REST API version 2024-01-01-preview.

    Other available API versions: 2023-05-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appplatform [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str app_name: The name of the App resource.
    :param builtins.str deployment_name: The name of the Deployment resource.
    :param builtins.str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param builtins.str service_name: The name of the Service resource.
    """
    __args__ = dict()
    __args__['appName'] = app_name
    __args__['deploymentName'] = deployment_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:appplatform:getDeploymentLogFileUrl', __args__, opts=opts, typ=GetDeploymentLogFileUrlResult).value

    return AwaitableGetDeploymentLogFileUrlResult(
        url=pulumi.get(__ret__, 'url'))
def get_deployment_log_file_url_output(app_name: Optional[pulumi.Input[builtins.str]] = None,
                                       deployment_name: Optional[pulumi.Input[builtins.str]] = None,
                                       resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                       service_name: Optional[pulumi.Input[builtins.str]] = None,
                                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDeploymentLogFileUrlResult]:
    """
    Get deployment log file URL

    Uses Azure REST API version 2024-01-01-preview.

    Other available API versions: 2023-05-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appplatform [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str app_name: The name of the App resource.
    :param builtins.str deployment_name: The name of the Deployment resource.
    :param builtins.str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param builtins.str service_name: The name of the Service resource.
    """
    __args__ = dict()
    __args__['appName'] = app_name
    __args__['deploymentName'] = deployment_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:appplatform:getDeploymentLogFileUrl', __args__, opts=opts, typ=GetDeploymentLogFileUrlResult)
    return __ret__.apply(lambda __response__: GetDeploymentLogFileUrlResult(
        url=pulumi.get(__response__, 'url')))
