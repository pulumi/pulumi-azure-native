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

__all__ = [
    'GetTagDescriptionResult',
    'AwaitableGetTagDescriptionResult',
    'get_tag_description',
    'get_tag_description_output',
]

@pulumi.output_type
class GetTagDescriptionResult:
    """
    Contract details.
    """
    def __init__(__self__, description=None, display_name=None, external_docs_description=None, external_docs_url=None, id=None, name=None, type=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if external_docs_description and not isinstance(external_docs_description, str):
            raise TypeError("Expected argument 'external_docs_description' to be a str")
        pulumi.set(__self__, "external_docs_description", external_docs_description)
        if external_docs_url and not isinstance(external_docs_url, str):
            raise TypeError("Expected argument 'external_docs_url' to be a str")
        pulumi.set(__self__, "external_docs_url", external_docs_url)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the Tag.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        Tag name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="externalDocsDescription")
    def external_docs_description(self) -> Optional[str]:
        """
        Description of the external resources describing the tag.
        """
        return pulumi.get(self, "external_docs_description")

    @property
    @pulumi.getter(name="externalDocsUrl")
    def external_docs_url(self) -> Optional[str]:
        """
        Absolute URL of external resources describing the tag.
        """
        return pulumi.get(self, "external_docs_url")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type for API Management resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetTagDescriptionResult(GetTagDescriptionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTagDescriptionResult(
            description=self.description,
            display_name=self.display_name,
            external_docs_description=self.external_docs_description,
            external_docs_url=self.external_docs_url,
            id=self.id,
            name=self.name,
            type=self.type)


def get_tag_description(api_id: Optional[str] = None,
                        resource_group_name: Optional[str] = None,
                        service_name: Optional[str] = None,
                        tag_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTagDescriptionResult:
    """
    Get Tag description in scope of API


    :param str api_id: API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
    :param str resource_group_name: The name of the resource group.
    :param str service_name: The name of the API Management service.
    :param str tag_id: Tag identifier. Must be unique in the current API Management service instance.
    """
    __args__ = dict()
    __args__['apiId'] = api_id
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    __args__['tagId'] = tag_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:apimanagement/v20180601preview:getTagDescription', __args__, opts=opts, typ=GetTagDescriptionResult).value

    return AwaitableGetTagDescriptionResult(
        description=pulumi.get(__ret__, 'description'),
        display_name=pulumi.get(__ret__, 'display_name'),
        external_docs_description=pulumi.get(__ret__, 'external_docs_description'),
        external_docs_url=pulumi.get(__ret__, 'external_docs_url'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        type=pulumi.get(__ret__, 'type'))
def get_tag_description_output(api_id: Optional[pulumi.Input[str]] = None,
                               resource_group_name: Optional[pulumi.Input[str]] = None,
                               service_name: Optional[pulumi.Input[str]] = None,
                               tag_id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetTagDescriptionResult]:
    """
    Get Tag description in scope of API


    :param str api_id: API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
    :param str resource_group_name: The name of the resource group.
    :param str service_name: The name of the API Management service.
    :param str tag_id: Tag identifier. Must be unique in the current API Management service instance.
    """
    __args__ = dict()
    __args__['apiId'] = api_id
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    __args__['tagId'] = tag_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:apimanagement/v20180601preview:getTagDescription', __args__, opts=opts, typ=GetTagDescriptionResult)
    return __ret__.apply(lambda __response__: GetTagDescriptionResult(
        description=pulumi.get(__response__, 'description'),
        display_name=pulumi.get(__response__, 'display_name'),
        external_docs_description=pulumi.get(__response__, 'external_docs_description'),
        external_docs_url=pulumi.get(__response__, 'external_docs_url'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        type=pulumi.get(__response__, 'type')))
