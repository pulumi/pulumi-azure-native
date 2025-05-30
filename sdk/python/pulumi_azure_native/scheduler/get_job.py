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
    'GetJobResult',
    'AwaitableGetJobResult',
    'get_job',
    'get_job_output',
]

@pulumi.output_type
class GetJobResult:
    def __init__(__self__, azure_api_version=None, id=None, name=None, properties=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)
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
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Gets the job resource identifier.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Gets the job resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.JobPropertiesResponse':
        """
        Gets or sets the job properties.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Gets the job resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetJobResult(GetJobResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetJobResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_job(job_collection_name: Optional[builtins.str] = None,
            job_name: Optional[builtins.str] = None,
            resource_group_name: Optional[builtins.str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetJobResult:
    """
    Gets a job.

    Uses Azure REST API version 2016-03-01.


    :param builtins.str job_collection_name: The job collection name.
    :param builtins.str job_name: The job name.
    :param builtins.str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['jobCollectionName'] = job_collection_name
    __args__['jobName'] = job_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:scheduler:getJob', __args__, opts=opts, typ=GetJobResult).value

    return AwaitableGetJobResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        properties=pulumi.get(__ret__, 'properties'),
        type=pulumi.get(__ret__, 'type'))
def get_job_output(job_collection_name: Optional[pulumi.Input[builtins.str]] = None,
                   job_name: Optional[pulumi.Input[builtins.str]] = None,
                   resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetJobResult]:
    """
    Gets a job.

    Uses Azure REST API version 2016-03-01.


    :param builtins.str job_collection_name: The job collection name.
    :param builtins.str job_name: The job name.
    :param builtins.str resource_group_name: The resource group name.
    """
    __args__ = dict()
    __args__['jobCollectionName'] = job_collection_name
    __args__['jobName'] = job_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:scheduler:getJob', __args__, opts=opts, typ=GetJobResult)
    return __ret__.apply(lambda __response__: GetJobResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        properties=pulumi.get(__response__, 'properties'),
        type=pulumi.get(__response__, 'type')))
