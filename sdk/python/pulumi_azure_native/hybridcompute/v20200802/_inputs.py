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
    'LocationDataArgs',
    'LocationDataArgsDict',
    'MachineIdentityArgs',
    'MachineIdentityArgsDict',
]

MYPY = False

if not MYPY:
    class LocationDataArgsDict(TypedDict):
        """
        Metadata pertaining to the geographic location of the resource.
        """
        name: pulumi.Input[str]
        """
        A canonical name for the geographic or physical location.
        """
        city: NotRequired[pulumi.Input[str]]
        """
        The city or locality where the resource is located.
        """
        country_or_region: NotRequired[pulumi.Input[str]]
        """
        The country or region where the resource is located
        """
        district: NotRequired[pulumi.Input[str]]
        """
        The district, state, or province where the resource is located.
        """
elif False:
    LocationDataArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class LocationDataArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 city: Optional[pulumi.Input[str]] = None,
                 country_or_region: Optional[pulumi.Input[str]] = None,
                 district: Optional[pulumi.Input[str]] = None):
        """
        Metadata pertaining to the geographic location of the resource.
        :param pulumi.Input[str] name: A canonical name for the geographic or physical location.
        :param pulumi.Input[str] city: The city or locality where the resource is located.
        :param pulumi.Input[str] country_or_region: The country or region where the resource is located
        :param pulumi.Input[str] district: The district, state, or province where the resource is located.
        """
        pulumi.set(__self__, "name", name)
        if city is not None:
            pulumi.set(__self__, "city", city)
        if country_or_region is not None:
            pulumi.set(__self__, "country_or_region", country_or_region)
        if district is not None:
            pulumi.set(__self__, "district", district)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        A canonical name for the geographic or physical location.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def city(self) -> Optional[pulumi.Input[str]]:
        """
        The city or locality where the resource is located.
        """
        return pulumi.get(self, "city")

    @city.setter
    def city(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "city", value)

    @property
    @pulumi.getter(name="countryOrRegion")
    def country_or_region(self) -> Optional[pulumi.Input[str]]:
        """
        The country or region where the resource is located
        """
        return pulumi.get(self, "country_or_region")

    @country_or_region.setter
    def country_or_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "country_or_region", value)

    @property
    @pulumi.getter
    def district(self) -> Optional[pulumi.Input[str]]:
        """
        The district, state, or province where the resource is located.
        """
        return pulumi.get(self, "district")

    @district.setter
    def district(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "district", value)


if not MYPY:
    class MachineIdentityArgsDict(TypedDict):
        type: NotRequired[pulumi.Input[str]]
        """
        The identity type.
        """
elif False:
    MachineIdentityArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class MachineIdentityArgs:
    def __init__(__self__, *,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: The identity type.
        """
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The identity type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


