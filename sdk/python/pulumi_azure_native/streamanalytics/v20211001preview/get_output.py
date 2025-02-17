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
from ... import _utilities
from . import outputs

__all__ = [
    'GetOutputResult',
    'AwaitableGetOutputResult',
    'get_output',
    'get_output_output',
]

@pulumi.output_type
class GetOutputResult:
    """
    An output object, containing all information associated with the named output. All outputs are contained under a streaming job.
    """
    def __init__(__self__, datasource=None, diagnostics=None, etag=None, id=None, last_output_event_timestamps=None, name=None, serialization=None, size_window=None, time_window=None, type=None, watermark_settings=None):
        if datasource and not isinstance(datasource, dict):
            raise TypeError("Expected argument 'datasource' to be a dict")
        pulumi.set(__self__, "datasource", datasource)
        if diagnostics and not isinstance(diagnostics, dict):
            raise TypeError("Expected argument 'diagnostics' to be a dict")
        pulumi.set(__self__, "diagnostics", diagnostics)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_output_event_timestamps and not isinstance(last_output_event_timestamps, list):
            raise TypeError("Expected argument 'last_output_event_timestamps' to be a list")
        pulumi.set(__self__, "last_output_event_timestamps", last_output_event_timestamps)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if serialization and not isinstance(serialization, dict):
            raise TypeError("Expected argument 'serialization' to be a dict")
        pulumi.set(__self__, "serialization", serialization)
        if size_window and not isinstance(size_window, int):
            raise TypeError("Expected argument 'size_window' to be a int")
        pulumi.set(__self__, "size_window", size_window)
        if time_window and not isinstance(time_window, str):
            raise TypeError("Expected argument 'time_window' to be a str")
        pulumi.set(__self__, "time_window", time_window)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if watermark_settings and not isinstance(watermark_settings, dict):
            raise TypeError("Expected argument 'watermark_settings' to be a dict")
        pulumi.set(__self__, "watermark_settings", watermark_settings)

    @property
    @pulumi.getter
    def datasource(self) -> Optional[Any]:
        """
        Describes the data source that output will be written to. Required on PUT (CreateOrReplace) requests.
        """
        return pulumi.get(self, "datasource")

    @property
    @pulumi.getter
    def diagnostics(self) -> 'outputs.DiagnosticsResponse':
        """
        Describes conditions applicable to the Input, Output, or the job overall, that warrant customer attention.
        """
        return pulumi.get(self, "diagnostics")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        The current entity tag for the output. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Resource Id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastOutputEventTimestamps")
    def last_output_event_timestamps(self) -> Sequence['outputs.LastOutputEventTimestampResponse']:
        """
        A list of the last output event times for each output partition. The index of the array corresponds to the partition number.
        """
        return pulumi.get(self, "last_output_event_timestamps")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def serialization(self) -> Optional[Any]:
        """
        Describes how data from an input is serialized or how data is serialized when written to an output. Required on PUT (CreateOrReplace) requests.
        """
        return pulumi.get(self, "serialization")

    @property
    @pulumi.getter(name="sizeWindow")
    def size_window(self) -> Optional[int]:
        """
        The size window to constrain a Stream Analytics output to.
        """
        return pulumi.get(self, "size_window")

    @property
    @pulumi.getter(name="timeWindow")
    def time_window(self) -> Optional[str]:
        """
        The time frame for filtering Stream Analytics job outputs.
        """
        return pulumi.get(self, "time_window")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="watermarkSettings")
    def watermark_settings(self) -> Optional['outputs.OutputWatermarkPropertiesResponse']:
        """
        Settings which determine whether to send watermarks to downstream.
        """
        return pulumi.get(self, "watermark_settings")


class AwaitableGetOutputResult(GetOutputResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOutputResult(
            datasource=self.datasource,
            diagnostics=self.diagnostics,
            etag=self.etag,
            id=self.id,
            last_output_event_timestamps=self.last_output_event_timestamps,
            name=self.name,
            serialization=self.serialization,
            size_window=self.size_window,
            time_window=self.time_window,
            type=self.type,
            watermark_settings=self.watermark_settings)


def get_output(job_name: Optional[str] = None,
               output_name: Optional[str] = None,
               resource_group_name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOutputResult:
    """
    Gets details about the specified output.


    :param str job_name: The name of the streaming job.
    :param str output_name: The name of the output.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['jobName'] = job_name
    __args__['outputName'] = output_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:streamanalytics/v20211001preview:getOutput', __args__, opts=opts, typ=GetOutputResult).value

    return AwaitableGetOutputResult(
        datasource=pulumi.get(__ret__, 'datasource'),
        diagnostics=pulumi.get(__ret__, 'diagnostics'),
        etag=pulumi.get(__ret__, 'etag'),
        id=pulumi.get(__ret__, 'id'),
        last_output_event_timestamps=pulumi.get(__ret__, 'last_output_event_timestamps'),
        name=pulumi.get(__ret__, 'name'),
        serialization=pulumi.get(__ret__, 'serialization'),
        size_window=pulumi.get(__ret__, 'size_window'),
        time_window=pulumi.get(__ret__, 'time_window'),
        type=pulumi.get(__ret__, 'type'),
        watermark_settings=pulumi.get(__ret__, 'watermark_settings'))
def get_output_output(job_name: Optional[pulumi.Input[str]] = None,
                      output_name: Optional[pulumi.Input[str]] = None,
                      resource_group_name: Optional[pulumi.Input[str]] = None,
                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetOutputResult]:
    """
    Gets details about the specified output.


    :param str job_name: The name of the streaming job.
    :param str output_name: The name of the output.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['jobName'] = job_name
    __args__['outputName'] = output_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:streamanalytics/v20211001preview:getOutput', __args__, opts=opts, typ=GetOutputResult)
    return __ret__.apply(lambda __response__: GetOutputResult(
        datasource=pulumi.get(__response__, 'datasource'),
        diagnostics=pulumi.get(__response__, 'diagnostics'),
        etag=pulumi.get(__response__, 'etag'),
        id=pulumi.get(__response__, 'id'),
        last_output_event_timestamps=pulumi.get(__response__, 'last_output_event_timestamps'),
        name=pulumi.get(__response__, 'name'),
        serialization=pulumi.get(__response__, 'serialization'),
        size_window=pulumi.get(__response__, 'size_window'),
        time_window=pulumi.get(__response__, 'time_window'),
        type=pulumi.get(__response__, 'type'),
        watermark_settings=pulumi.get(__response__, 'watermark_settings')))
