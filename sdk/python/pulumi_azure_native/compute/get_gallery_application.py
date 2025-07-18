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
    'GetGalleryApplicationResult',
    'AwaitableGetGalleryApplicationResult',
    'get_gallery_application',
    'get_gallery_application_output',
]

@pulumi.output_type
class GetGalleryApplicationResult:
    """
    Specifies information about the gallery Application Definition that you want to create or update.
    """
    def __init__(__self__, azure_api_version=None, custom_actions=None, description=None, end_of_life_date=None, eula=None, id=None, location=None, name=None, privacy_statement_uri=None, release_note_uri=None, supported_os_type=None, system_data=None, tags=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if custom_actions and not isinstance(custom_actions, list):
            raise TypeError("Expected argument 'custom_actions' to be a list")
        pulumi.set(__self__, "custom_actions", custom_actions)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if end_of_life_date and not isinstance(end_of_life_date, str):
            raise TypeError("Expected argument 'end_of_life_date' to be a str")
        pulumi.set(__self__, "end_of_life_date", end_of_life_date)
        if eula and not isinstance(eula, str):
            raise TypeError("Expected argument 'eula' to be a str")
        pulumi.set(__self__, "eula", eula)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if privacy_statement_uri and not isinstance(privacy_statement_uri, str):
            raise TypeError("Expected argument 'privacy_statement_uri' to be a str")
        pulumi.set(__self__, "privacy_statement_uri", privacy_statement_uri)
        if release_note_uri and not isinstance(release_note_uri, str):
            raise TypeError("Expected argument 'release_note_uri' to be a str")
        pulumi.set(__self__, "release_note_uri", release_note_uri)
        if supported_os_type and not isinstance(supported_os_type, str):
            raise TypeError("Expected argument 'supported_os_type' to be a str")
        pulumi.set(__self__, "supported_os_type", supported_os_type)
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
    @pulumi.getter(name="customActions")
    def custom_actions(self) -> Optional[Sequence['outputs.GalleryApplicationCustomActionResponse']]:
        """
        A list of custom actions that can be performed with all of the Gallery Application Versions within this Gallery Application.
        """
        return pulumi.get(self, "custom_actions")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        The description of this gallery Application Definition resource. This property is updatable.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endOfLifeDate")
    def end_of_life_date(self) -> Optional[builtins.str]:
        """
        The end of life date of the gallery Application Definition. This property can be used for decommissioning purposes. This property is updatable.
        """
        return pulumi.get(self, "end_of_life_date")

    @property
    @pulumi.getter
    def eula(self) -> Optional[builtins.str]:
        """
        The Eula agreement for the gallery Application Definition.
        """
        return pulumi.get(self, "eula")

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
    @pulumi.getter(name="privacyStatementUri")
    def privacy_statement_uri(self) -> Optional[builtins.str]:
        """
        The privacy statement uri.
        """
        return pulumi.get(self, "privacy_statement_uri")

    @property
    @pulumi.getter(name="releaseNoteUri")
    def release_note_uri(self) -> Optional[builtins.str]:
        """
        The release note uri.
        """
        return pulumi.get(self, "release_note_uri")

    @property
    @pulumi.getter(name="supportedOSType")
    def supported_os_type(self) -> builtins.str:
        """
        This property allows you to specify the supported type of the OS that application is built for. Possible values are: **Windows,** **Linux.**
        """
        return pulumi.get(self, "supported_os_type")

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


class AwaitableGetGalleryApplicationResult(GetGalleryApplicationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGalleryApplicationResult(
            azure_api_version=self.azure_api_version,
            custom_actions=self.custom_actions,
            description=self.description,
            end_of_life_date=self.end_of_life_date,
            eula=self.eula,
            id=self.id,
            location=self.location,
            name=self.name,
            privacy_statement_uri=self.privacy_statement_uri,
            release_note_uri=self.release_note_uri,
            supported_os_type=self.supported_os_type,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type)


def get_gallery_application(gallery_application_name: Optional[builtins.str] = None,
                            gallery_name: Optional[builtins.str] = None,
                            resource_group_name: Optional[builtins.str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGalleryApplicationResult:
    """
    Retrieves information about a gallery Application Definition.

    Uses Azure REST API version 2024-03-03.

    Other available API versions: 2022-03-03, 2022-08-03, 2023-07-03. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str gallery_application_name: The name of the gallery Application Definition to be retrieved.
    :param builtins.str gallery_name: The name of the Shared Image Gallery.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['galleryApplicationName'] = gallery_application_name
    __args__['galleryName'] = gallery_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:compute:getGalleryApplication', __args__, opts=opts, typ=GetGalleryApplicationResult).value

    return AwaitableGetGalleryApplicationResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        custom_actions=pulumi.get(__ret__, 'custom_actions'),
        description=pulumi.get(__ret__, 'description'),
        end_of_life_date=pulumi.get(__ret__, 'end_of_life_date'),
        eula=pulumi.get(__ret__, 'eula'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        privacy_statement_uri=pulumi.get(__ret__, 'privacy_statement_uri'),
        release_note_uri=pulumi.get(__ret__, 'release_note_uri'),
        supported_os_type=pulumi.get(__ret__, 'supported_os_type'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_gallery_application_output(gallery_application_name: Optional[pulumi.Input[builtins.str]] = None,
                                   gallery_name: Optional[pulumi.Input[builtins.str]] = None,
                                   resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetGalleryApplicationResult]:
    """
    Retrieves information about a gallery Application Definition.

    Uses Azure REST API version 2024-03-03.

    Other available API versions: 2022-03-03, 2022-08-03, 2023-07-03. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str gallery_application_name: The name of the gallery Application Definition to be retrieved.
    :param builtins.str gallery_name: The name of the Shared Image Gallery.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['galleryApplicationName'] = gallery_application_name
    __args__['galleryName'] = gallery_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:compute:getGalleryApplication', __args__, opts=opts, typ=GetGalleryApplicationResult)
    return __ret__.apply(lambda __response__: GetGalleryApplicationResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        custom_actions=pulumi.get(__response__, 'custom_actions'),
        description=pulumi.get(__response__, 'description'),
        end_of_life_date=pulumi.get(__response__, 'end_of_life_date'),
        eula=pulumi.get(__response__, 'eula'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        privacy_statement_uri=pulumi.get(__response__, 'privacy_statement_uri'),
        release_note_uri=pulumi.get(__response__, 'release_note_uri'),
        supported_os_type=pulumi.get(__response__, 'supported_os_type'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
