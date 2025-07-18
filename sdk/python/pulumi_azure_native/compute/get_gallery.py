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
    'GetGalleryResult',
    'AwaitableGetGalleryResult',
    'get_gallery',
    'get_gallery_output',
]

@pulumi.output_type
class GetGalleryResult:
    """
    Specifies information about the Shared Image Gallery that you want to create or update.
    """
    def __init__(__self__, azure_api_version=None, description=None, id=None, identifier=None, identity=None, location=None, name=None, provisioning_state=None, sharing_profile=None, sharing_status=None, soft_delete_policy=None, system_data=None, tags=None, type=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identifier and not isinstance(identifier, dict):
            raise TypeError("Expected argument 'identifier' to be a dict")
        pulumi.set(__self__, "identifier", identifier)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if sharing_profile and not isinstance(sharing_profile, dict):
            raise TypeError("Expected argument 'sharing_profile' to be a dict")
        pulumi.set(__self__, "sharing_profile", sharing_profile)
        if sharing_status and not isinstance(sharing_status, dict):
            raise TypeError("Expected argument 'sharing_status' to be a dict")
        pulumi.set(__self__, "sharing_status", sharing_status)
        if soft_delete_policy and not isinstance(soft_delete_policy, dict):
            raise TypeError("Expected argument 'soft_delete_policy' to be a dict")
        pulumi.set(__self__, "soft_delete_policy", soft_delete_policy)
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
    def description(self) -> Optional[builtins.str]:
        """
        The description of this Shared Image Gallery resource. This property is updatable.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identifier(self) -> Optional['outputs.GalleryIdentifierResponse']:
        """
        Describes the gallery unique name.
        """
        return pulumi.get(self, "identifier")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.GalleryIdentityResponse']:
        """
        The identity of the gallery, if configured.
        """
        return pulumi.get(self, "identity")

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
    @pulumi.getter(name="sharingProfile")
    def sharing_profile(self) -> Optional['outputs.SharingProfileResponse']:
        """
        Profile for gallery sharing to subscription or tenant
        """
        return pulumi.get(self, "sharing_profile")

    @property
    @pulumi.getter(name="sharingStatus")
    def sharing_status(self) -> 'outputs.SharingStatusResponse':
        """
        Sharing status of current gallery.
        """
        return pulumi.get(self, "sharing_status")

    @property
    @pulumi.getter(name="softDeletePolicy")
    def soft_delete_policy(self) -> Optional['outputs.SoftDeletePolicyResponse']:
        """
        Contains information about the soft deletion policy of the gallery.
        """
        return pulumi.get(self, "soft_delete_policy")

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


class AwaitableGetGalleryResult(GetGalleryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGalleryResult(
            azure_api_version=self.azure_api_version,
            description=self.description,
            id=self.id,
            identifier=self.identifier,
            identity=self.identity,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            sharing_profile=self.sharing_profile,
            sharing_status=self.sharing_status,
            soft_delete_policy=self.soft_delete_policy,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type)


def get_gallery(expand: Optional[builtins.str] = None,
                gallery_name: Optional[builtins.str] = None,
                resource_group_name: Optional[builtins.str] = None,
                select: Optional[builtins.str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGalleryResult:
    """
    Retrieves information about a Shared Image Gallery.

    Uses Azure REST API version 2024-03-03.

    Other available API versions: 2022-03-03, 2022-08-03, 2023-07-03. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str expand: The expand query option to apply on the operation.
    :param builtins.str gallery_name: The name of the Shared Image Gallery.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str select: The select expression to apply on the operation.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['galleryName'] = gallery_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['select'] = select
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:compute:getGallery', __args__, opts=opts, typ=GetGalleryResult).value

    return AwaitableGetGalleryResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        identifier=pulumi.get(__ret__, 'identifier'),
        identity=pulumi.get(__ret__, 'identity'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        sharing_profile=pulumi.get(__ret__, 'sharing_profile'),
        sharing_status=pulumi.get(__ret__, 'sharing_status'),
        soft_delete_policy=pulumi.get(__ret__, 'soft_delete_policy'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_gallery_output(expand: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                       gallery_name: Optional[pulumi.Input[builtins.str]] = None,
                       resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                       select: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetGalleryResult]:
    """
    Retrieves information about a Shared Image Gallery.

    Uses Azure REST API version 2024-03-03.

    Other available API versions: 2022-03-03, 2022-08-03, 2023-07-03. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str expand: The expand query option to apply on the operation.
    :param builtins.str gallery_name: The name of the Shared Image Gallery.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str select: The select expression to apply on the operation.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['galleryName'] = gallery_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['select'] = select
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:compute:getGallery', __args__, opts=opts, typ=GetGalleryResult)
    return __ret__.apply(lambda __response__: GetGalleryResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        identifier=pulumi.get(__response__, 'identifier'),
        identity=pulumi.get(__response__, 'identity'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        sharing_profile=pulumi.get(__response__, 'sharing_profile'),
        sharing_status=pulumi.get(__response__, 'sharing_status'),
        soft_delete_policy=pulumi.get(__response__, 'soft_delete_policy'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
