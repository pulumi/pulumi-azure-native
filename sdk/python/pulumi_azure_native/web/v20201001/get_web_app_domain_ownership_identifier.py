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
    'GetWebAppDomainOwnershipIdentifierResult',
    'AwaitableGetWebAppDomainOwnershipIdentifierResult',
    'get_web_app_domain_ownership_identifier',
    'get_web_app_domain_ownership_identifier_output',
]

@pulumi.output_type
class GetWebAppDomainOwnershipIdentifierResult:
    """
    A domain specific resource identifier.
    """
    def __init__(__self__, id=None, kind=None, name=None, system_data=None, type=None, value=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Resource Id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind of resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        The system metadata relating to this resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        String representation of the identity.
        """
        return pulumi.get(self, "value")


class AwaitableGetWebAppDomainOwnershipIdentifierResult(GetWebAppDomainOwnershipIdentifierResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWebAppDomainOwnershipIdentifierResult(
            id=self.id,
            kind=self.kind,
            name=self.name,
            system_data=self.system_data,
            type=self.type,
            value=self.value)


def get_web_app_domain_ownership_identifier(domain_ownership_identifier_name: Optional[str] = None,
                                            name: Optional[str] = None,
                                            resource_group_name: Optional[str] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWebAppDomainOwnershipIdentifierResult:
    """
    Get domain ownership identifier for web app.


    :param str domain_ownership_identifier_name: Name of domain ownership identifier.
    :param str name: Name of the app.
    :param str resource_group_name: Name of the resource group to which the resource belongs.
    """
    __args__ = dict()
    __args__['domainOwnershipIdentifierName'] = domain_ownership_identifier_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:web/v20201001:getWebAppDomainOwnershipIdentifier', __args__, opts=opts, typ=GetWebAppDomainOwnershipIdentifierResult).value

    return AwaitableGetWebAppDomainOwnershipIdentifierResult(
        id=pulumi.get(__ret__, 'id'),
        kind=pulumi.get(__ret__, 'kind'),
        name=pulumi.get(__ret__, 'name'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'),
        value=pulumi.get(__ret__, 'value'))
def get_web_app_domain_ownership_identifier_output(domain_ownership_identifier_name: Optional[pulumi.Input[str]] = None,
                                                   name: Optional[pulumi.Input[str]] = None,
                                                   resource_group_name: Optional[pulumi.Input[str]] = None,
                                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetWebAppDomainOwnershipIdentifierResult]:
    """
    Get domain ownership identifier for web app.


    :param str domain_ownership_identifier_name: Name of domain ownership identifier.
    :param str name: Name of the app.
    :param str resource_group_name: Name of the resource group to which the resource belongs.
    """
    __args__ = dict()
    __args__['domainOwnershipIdentifierName'] = domain_ownership_identifier_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:web/v20201001:getWebAppDomainOwnershipIdentifier', __args__, opts=opts, typ=GetWebAppDomainOwnershipIdentifierResult)
    return __ret__.apply(lambda __response__: GetWebAppDomainOwnershipIdentifierResult(
        id=pulumi.get(__response__, 'id'),
        kind=pulumi.get(__response__, 'kind'),
        name=pulumi.get(__response__, 'name'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type'),
        value=pulumi.get(__response__, 'value')))
