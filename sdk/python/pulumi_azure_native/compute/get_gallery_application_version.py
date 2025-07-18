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
    'GetGalleryApplicationVersionResult',
    'AwaitableGetGalleryApplicationVersionResult',
    'get_gallery_application_version',
    'get_gallery_application_version_output',
]

@pulumi.output_type
class GetGalleryApplicationVersionResult:
    """
    Specifies information about the gallery Application Version that you want to create or update.
    """
    def __init__(__self__, azure_api_version=None, id=None, location=None, name=None, provisioning_state=None, publishing_profile=None, replication_status=None, safety_profile=None, system_data=None, tags=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
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
        if publishing_profile and not isinstance(publishing_profile, dict):
            raise TypeError("Expected argument 'publishing_profile' to be a dict")
        pulumi.set(__self__, "publishing_profile", publishing_profile)
        if replication_status and not isinstance(replication_status, dict):
            raise TypeError("Expected argument 'replication_status' to be a dict")
        pulumi.set(__self__, "replication_status", replication_status)
        if safety_profile and not isinstance(safety_profile, dict):
            raise TypeError("Expected argument 'safety_profile' to be a dict")
        pulumi.set(__self__, "safety_profile", safety_profile)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
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
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

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
        The provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publishingProfile")
    def publishing_profile(self) -> 'outputs.GalleryApplicationVersionPublishingProfileResponse':
        """
        The publishing profile of a gallery image version.
        """
        return pulumi.get(self, "publishing_profile")

    @property
    @pulumi.getter(name="replicationStatus")
    def replication_status(self) -> 'outputs.ReplicationStatusResponse':
        """
        This is the replication status of the gallery image version.
        """
        return pulumi.get(self, "replication_status")

    @property
    @pulumi.getter(name="safetyProfile")
    def safety_profile(self) -> Optional['outputs.GalleryApplicationVersionSafetyProfileResponse']:
        """
        The safety profile of the Gallery Application Version.
        """
        return pulumi.get(self, "safety_profile")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetGalleryApplicationVersionResult(GetGalleryApplicationVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGalleryApplicationVersionResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            publishing_profile=self.publishing_profile,
            replication_status=self.replication_status,
            safety_profile=self.safety_profile,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type)


def get_gallery_application_version(expand: Optional[builtins.str] = None,
                                    gallery_application_name: Optional[builtins.str] = None,
                                    gallery_application_version_name: Optional[builtins.str] = None,
                                    gallery_name: Optional[builtins.str] = None,
                                    resource_group_name: Optional[builtins.str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGalleryApplicationVersionResult:
    """
    Retrieves information about a gallery Application Version.

    Uses Azure REST API version 2024-03-03.

    Other available API versions: 2022-03-03, 2022-08-03, 2023-07-03. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str expand: The expand expression to apply on the operation.
    :param builtins.str gallery_application_name: The name of the gallery Application Definition to be retrieved.
    :param builtins.str gallery_application_version_name: The name of the gallery Application Version to be retrieved.
    :param builtins.str gallery_name: The name of the Shared Image Gallery.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['galleryApplicationName'] = gallery_application_name
    __args__['galleryApplicationVersionName'] = gallery_application_version_name
    __args__['galleryName'] = gallery_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:compute:getGalleryApplicationVersion', __args__, opts=opts, typ=GetGalleryApplicationVersionResult).value

    return AwaitableGetGalleryApplicationVersionResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        publishing_profile=pulumi.get(__ret__, 'publishing_profile'),
        replication_status=pulumi.get(__ret__, 'replication_status'),
        safety_profile=pulumi.get(__ret__, 'safety_profile'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_gallery_application_version_output(expand: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                           gallery_application_name: Optional[pulumi.Input[builtins.str]] = None,
                                           gallery_application_version_name: Optional[pulumi.Input[builtins.str]] = None,
                                           gallery_name: Optional[pulumi.Input[builtins.str]] = None,
                                           resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetGalleryApplicationVersionResult]:
    """
    Retrieves information about a gallery Application Version.

    Uses Azure REST API version 2024-03-03.

    Other available API versions: 2022-03-03, 2022-08-03, 2023-07-03. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str expand: The expand expression to apply on the operation.
    :param builtins.str gallery_application_name: The name of the gallery Application Definition to be retrieved.
    :param builtins.str gallery_application_version_name: The name of the gallery Application Version to be retrieved.
    :param builtins.str gallery_name: The name of the Shared Image Gallery.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['galleryApplicationName'] = gallery_application_name
    __args__['galleryApplicationVersionName'] = gallery_application_version_name
    __args__['galleryName'] = gallery_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:compute:getGalleryApplicationVersion', __args__, opts=opts, typ=GetGalleryApplicationVersionResult)
    return __ret__.apply(lambda __response__: GetGalleryApplicationVersionResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        publishing_profile=pulumi.get(__response__, 'publishing_profile'),
        replication_status=pulumi.get(__response__, 'replication_status'),
        safety_profile=pulumi.get(__response__, 'safety_profile'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
