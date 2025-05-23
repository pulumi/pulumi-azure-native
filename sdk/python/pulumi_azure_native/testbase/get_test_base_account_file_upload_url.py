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
from ._enums import *

__all__ = [
    'GetTestBaseAccountFileUploadUrlResult',
    'AwaitableGetTestBaseAccountFileUploadUrlResult',
    'get_test_base_account_file_upload_url',
    'get_test_base_account_file_upload_url_output',
]

@pulumi.output_type
class GetTestBaseAccountFileUploadUrlResult:
    """
    The URL response
    """
    def __init__(__self__, blob_path=None, upload_url=None):
        if blob_path and not isinstance(blob_path, str):
            raise TypeError("Expected argument 'blob_path' to be a str")
        pulumi.set(__self__, "blob_path", blob_path)
        if upload_url and not isinstance(upload_url, str):
            raise TypeError("Expected argument 'upload_url' to be a str")
        pulumi.set(__self__, "upload_url", upload_url)

    @property
    @pulumi.getter(name="blobPath")
    def blob_path(self) -> builtins.str:
        """
        The blob path of the uploaded package. It will be used as the 'blobPath' property of PackageResource.
        """
        return pulumi.get(self, "blob_path")

    @property
    @pulumi.getter(name="uploadUrl")
    def upload_url(self) -> builtins.str:
        """
        The URL used for uploading the package.
        """
        return pulumi.get(self, "upload_url")


class AwaitableGetTestBaseAccountFileUploadUrlResult(GetTestBaseAccountFileUploadUrlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTestBaseAccountFileUploadUrlResult(
            blob_path=self.blob_path,
            upload_url=self.upload_url)


def get_test_base_account_file_upload_url(blob_name: Optional[builtins.str] = None,
                                          resource_group_name: Optional[builtins.str] = None,
                                          resource_type: Optional[Union[builtins.str, 'FileUploadResourceType']] = None,
                                          test_base_account_name: Optional[builtins.str] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTestBaseAccountFileUploadUrlResult:
    """
    Gets the file upload URL of a Test Base Account.

    Uses Azure REST API version 2023-11-01-preview.

    Other available API versions: 2022-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native testbase [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str blob_name: The custom file name of the uploaded blob.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param Union[builtins.str, 'FileUploadResourceType'] resource_type: Resource type for file uploading.
    :param builtins.str test_base_account_name: The resource name of the Test Base Account.
    """
    __args__ = dict()
    __args__['blobName'] = blob_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceType'] = resource_type
    __args__['testBaseAccountName'] = test_base_account_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:testbase:getTestBaseAccountFileUploadUrl', __args__, opts=opts, typ=GetTestBaseAccountFileUploadUrlResult).value

    return AwaitableGetTestBaseAccountFileUploadUrlResult(
        blob_path=pulumi.get(__ret__, 'blob_path'),
        upload_url=pulumi.get(__ret__, 'upload_url'))
def get_test_base_account_file_upload_url_output(blob_name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                                 resource_type: Optional[pulumi.Input[Optional[Union[builtins.str, 'FileUploadResourceType']]]] = None,
                                                 test_base_account_name: Optional[pulumi.Input[builtins.str]] = None,
                                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetTestBaseAccountFileUploadUrlResult]:
    """
    Gets the file upload URL of a Test Base Account.

    Uses Azure REST API version 2023-11-01-preview.

    Other available API versions: 2022-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native testbase [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str blob_name: The custom file name of the uploaded blob.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param Union[builtins.str, 'FileUploadResourceType'] resource_type: Resource type for file uploading.
    :param builtins.str test_base_account_name: The resource name of the Test Base Account.
    """
    __args__ = dict()
    __args__['blobName'] = blob_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceType'] = resource_type
    __args__['testBaseAccountName'] = test_base_account_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:testbase:getTestBaseAccountFileUploadUrl', __args__, opts=opts, typ=GetTestBaseAccountFileUploadUrlResult)
    return __ret__.apply(lambda __response__: GetTestBaseAccountFileUploadUrlResult(
        blob_path=pulumi.get(__response__, 'blob_path'),
        upload_url=pulumi.get(__response__, 'upload_url')))
