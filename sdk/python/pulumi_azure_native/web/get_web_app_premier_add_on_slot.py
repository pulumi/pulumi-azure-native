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
    'GetWebAppPremierAddOnSlotResult',
    'AwaitableGetWebAppPremierAddOnSlotResult',
    'get_web_app_premier_add_on_slot',
    'get_web_app_premier_add_on_slot_output',
]

@pulumi.output_type
class GetWebAppPremierAddOnSlotResult:
    """
    Premier add-on.
    """
    def __init__(__self__, azure_api_version=None, id=None, kind=None, location=None, marketplace_offer=None, marketplace_publisher=None, name=None, product=None, sku=None, tags=None, type=None, vendor=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if marketplace_offer and not isinstance(marketplace_offer, str):
            raise TypeError("Expected argument 'marketplace_offer' to be a str")
        pulumi.set(__self__, "marketplace_offer", marketplace_offer)
        if marketplace_publisher and not isinstance(marketplace_publisher, str):
            raise TypeError("Expected argument 'marketplace_publisher' to be a str")
        pulumi.set(__self__, "marketplace_publisher", marketplace_publisher)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if product and not isinstance(product, str):
            raise TypeError("Expected argument 'product' to be a str")
        pulumi.set(__self__, "product", product)
        if sku and not isinstance(sku, str):
            raise TypeError("Expected argument 'sku' to be a str")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vendor and not isinstance(vendor, str):
            raise TypeError("Expected argument 'vendor' to be a str")
        pulumi.set(__self__, "vendor", vendor)

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
        Resource Id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def kind(self) -> Optional[builtins.str]:
        """
        Kind of resource. If the resource is an app, you can refer to https://github.com/Azure/app-service-linux-docs/blob/master/Things_You_Should_Know/kind_property.md#app-service-resource-kind-reference for details supported values for kind.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        Resource Location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="marketplaceOffer")
    def marketplace_offer(self) -> Optional[builtins.str]:
        """
        Premier add on Marketplace offer.
        """
        return pulumi.get(self, "marketplace_offer")

    @property
    @pulumi.getter(name="marketplacePublisher")
    def marketplace_publisher(self) -> Optional[builtins.str]:
        """
        Premier add on Marketplace publisher.
        """
        return pulumi.get(self, "marketplace_publisher")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Resource Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def product(self) -> Optional[builtins.str]:
        """
        Premier add on Product.
        """
        return pulumi.get(self, "product")

    @property
    @pulumi.getter
    def sku(self) -> Optional[builtins.str]:
        """
        Premier add on SKU.
        """
        return pulumi.get(self, "sku")

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
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def vendor(self) -> Optional[builtins.str]:
        """
        Premier add on Vendor.
        """
        return pulumi.get(self, "vendor")


class AwaitableGetWebAppPremierAddOnSlotResult(GetWebAppPremierAddOnSlotResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWebAppPremierAddOnSlotResult(
            azure_api_version=self.azure_api_version,
            id=self.id,
            kind=self.kind,
            location=self.location,
            marketplace_offer=self.marketplace_offer,
            marketplace_publisher=self.marketplace_publisher,
            name=self.name,
            product=self.product,
            sku=self.sku,
            tags=self.tags,
            type=self.type,
            vendor=self.vendor)


def get_web_app_premier_add_on_slot(name: Optional[builtins.str] = None,
                                    premier_add_on_name: Optional[builtins.str] = None,
                                    resource_group_name: Optional[builtins.str] = None,
                                    slot: Optional[builtins.str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWebAppPremierAddOnSlotResult:
    """
    Description for Gets a named add-on of an app.

    Uses Azure REST API version 2024-04-01.

    Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str name: Name of the app.
    :param builtins.str premier_add_on_name: Add-on name.
    :param builtins.str resource_group_name: Name of the resource group to which the resource belongs.
    :param builtins.str slot: Name of the deployment slot. If a slot is not specified, the API will get the named add-on for the production slot.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['premierAddOnName'] = premier_add_on_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['slot'] = slot
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:web:getWebAppPremierAddOnSlot', __args__, opts=opts, typ=GetWebAppPremierAddOnSlotResult).value

    return AwaitableGetWebAppPremierAddOnSlotResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        id=pulumi.get(__ret__, 'id'),
        kind=pulumi.get(__ret__, 'kind'),
        location=pulumi.get(__ret__, 'location'),
        marketplace_offer=pulumi.get(__ret__, 'marketplace_offer'),
        marketplace_publisher=pulumi.get(__ret__, 'marketplace_publisher'),
        name=pulumi.get(__ret__, 'name'),
        product=pulumi.get(__ret__, 'product'),
        sku=pulumi.get(__ret__, 'sku'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        vendor=pulumi.get(__ret__, 'vendor'))
def get_web_app_premier_add_on_slot_output(name: Optional[pulumi.Input[builtins.str]] = None,
                                           premier_add_on_name: Optional[pulumi.Input[builtins.str]] = None,
                                           resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                           slot: Optional[pulumi.Input[builtins.str]] = None,
                                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetWebAppPremierAddOnSlotResult]:
    """
    Description for Gets a named add-on of an app.

    Uses Azure REST API version 2024-04-01.

    Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str name: Name of the app.
    :param builtins.str premier_add_on_name: Add-on name.
    :param builtins.str resource_group_name: Name of the resource group to which the resource belongs.
    :param builtins.str slot: Name of the deployment slot. If a slot is not specified, the API will get the named add-on for the production slot.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['premierAddOnName'] = premier_add_on_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['slot'] = slot
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:web:getWebAppPremierAddOnSlot', __args__, opts=opts, typ=GetWebAppPremierAddOnSlotResult)
    return __ret__.apply(lambda __response__: GetWebAppPremierAddOnSlotResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        id=pulumi.get(__response__, 'id'),
        kind=pulumi.get(__response__, 'kind'),
        location=pulumi.get(__response__, 'location'),
        marketplace_offer=pulumi.get(__response__, 'marketplace_offer'),
        marketplace_publisher=pulumi.get(__response__, 'marketplace_publisher'),
        name=pulumi.get(__response__, 'name'),
        product=pulumi.get(__response__, 'product'),
        sku=pulumi.get(__response__, 'sku'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        vendor=pulumi.get(__response__, 'vendor')))
