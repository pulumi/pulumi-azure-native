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
    'GetWebAppDeploymentSlotResult',
    'AwaitableGetWebAppDeploymentSlotResult',
    'get_web_app_deployment_slot',
    'get_web_app_deployment_slot_output',
]

@pulumi.output_type
class GetWebAppDeploymentSlotResult:
    """
    User credentials used for publishing activity.
    """
    def __init__(__self__, active=None, author=None, author_email=None, azure_api_version=None, deployer=None, details=None, end_time=None, id=None, kind=None, message=None, name=None, start_time=None, status=None, type=None):
        if active and not isinstance(active, bool):
            raise TypeError("Expected argument 'active' to be a bool")
        pulumi.set(__self__, "active", active)
        if author and not isinstance(author, str):
            raise TypeError("Expected argument 'author' to be a str")
        pulumi.set(__self__, "author", author)
        if author_email and not isinstance(author_email, str):
            raise TypeError("Expected argument 'author_email' to be a str")
        pulumi.set(__self__, "author_email", author_email)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if deployer and not isinstance(deployer, str):
            raise TypeError("Expected argument 'deployer' to be a str")
        pulumi.set(__self__, "deployer", deployer)
        if details and not isinstance(details, str):
            raise TypeError("Expected argument 'details' to be a str")
        pulumi.set(__self__, "details", details)
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if message and not isinstance(message, str):
            raise TypeError("Expected argument 'message' to be a str")
        pulumi.set(__self__, "message", message)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)
        if status and not isinstance(status, int):
            raise TypeError("Expected argument 'status' to be a int")
        pulumi.set(__self__, "status", status)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def active(self) -> Optional[builtins.bool]:
        """
        True if deployment is currently active, false if completed and null if not started.
        """
        return pulumi.get(self, "active")

    @property
    @pulumi.getter
    def author(self) -> Optional[builtins.str]:
        """
        Who authored the deployment.
        """
        return pulumi.get(self, "author")

    @property
    @pulumi.getter(name="authorEmail")
    def author_email(self) -> Optional[builtins.str]:
        """
        Author email.
        """
        return pulumi.get(self, "author_email")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def deployer(self) -> Optional[builtins.str]:
        """
        Who performed the deployment.
        """
        return pulumi.get(self, "deployer")

    @property
    @pulumi.getter
    def details(self) -> Optional[builtins.str]:
        """
        Details on deployment.
        """
        return pulumi.get(self, "details")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[builtins.str]:
        """
        End time.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Resource Id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def kind(self) -> Optional[builtins.str]:
        """
        Kind of resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def message(self) -> Optional[builtins.str]:
        """
        Details about deployment status.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Resource Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[builtins.str]:
        """
        Start time.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def status(self) -> Optional[builtins.int]:
        """
        Deployment status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetWebAppDeploymentSlotResult(GetWebAppDeploymentSlotResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWebAppDeploymentSlotResult(
            active=self.active,
            author=self.author,
            author_email=self.author_email,
            azure_api_version=self.azure_api_version,
            deployer=self.deployer,
            details=self.details,
            end_time=self.end_time,
            id=self.id,
            kind=self.kind,
            message=self.message,
            name=self.name,
            start_time=self.start_time,
            status=self.status,
            type=self.type)


def get_web_app_deployment_slot(id: Optional[builtins.str] = None,
                                name: Optional[builtins.str] = None,
                                resource_group_name: Optional[builtins.str] = None,
                                slot: Optional[builtins.str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWebAppDeploymentSlotResult:
    """
    Description for Get a deployment by its ID for an app, or a deployment slot.

    Uses Azure REST API version 2024-04-01.

    Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str id: Deployment ID.
    :param builtins.str name: Name of the app.
    :param builtins.str resource_group_name: Name of the resource group to which the resource belongs.
    :param builtins.str slot: Name of the deployment slot. If a slot is not specified, the API gets a deployment for the production slot.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['slot'] = slot
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:web:getWebAppDeploymentSlot', __args__, opts=opts, typ=GetWebAppDeploymentSlotResult).value

    return AwaitableGetWebAppDeploymentSlotResult(
        active=pulumi.get(__ret__, 'active'),
        author=pulumi.get(__ret__, 'author'),
        author_email=pulumi.get(__ret__, 'author_email'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        deployer=pulumi.get(__ret__, 'deployer'),
        details=pulumi.get(__ret__, 'details'),
        end_time=pulumi.get(__ret__, 'end_time'),
        id=pulumi.get(__ret__, 'id'),
        kind=pulumi.get(__ret__, 'kind'),
        message=pulumi.get(__ret__, 'message'),
        name=pulumi.get(__ret__, 'name'),
        start_time=pulumi.get(__ret__, 'start_time'),
        status=pulumi.get(__ret__, 'status'),
        type=pulumi.get(__ret__, 'type'))
def get_web_app_deployment_slot_output(id: Optional[pulumi.Input[builtins.str]] = None,
                                       name: Optional[pulumi.Input[builtins.str]] = None,
                                       resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                       slot: Optional[pulumi.Input[builtins.str]] = None,
                                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetWebAppDeploymentSlotResult]:
    """
    Description for Get a deployment by its ID for an app, or a deployment slot.

    Uses Azure REST API version 2024-04-01.

    Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str id: Deployment ID.
    :param builtins.str name: Name of the app.
    :param builtins.str resource_group_name: Name of the resource group to which the resource belongs.
    :param builtins.str slot: Name of the deployment slot. If a slot is not specified, the API gets a deployment for the production slot.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['slot'] = slot
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:web:getWebAppDeploymentSlot', __args__, opts=opts, typ=GetWebAppDeploymentSlotResult)
    return __ret__.apply(lambda __response__: GetWebAppDeploymentSlotResult(
        active=pulumi.get(__response__, 'active'),
        author=pulumi.get(__response__, 'author'),
        author_email=pulumi.get(__response__, 'author_email'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        deployer=pulumi.get(__response__, 'deployer'),
        details=pulumi.get(__response__, 'details'),
        end_time=pulumi.get(__response__, 'end_time'),
        id=pulumi.get(__response__, 'id'),
        kind=pulumi.get(__response__, 'kind'),
        message=pulumi.get(__response__, 'message'),
        name=pulumi.get(__response__, 'name'),
        start_time=pulumi.get(__response__, 'start_time'),
        status=pulumi.get(__response__, 'status'),
        type=pulumi.get(__response__, 'type')))
