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
from .. import _utilities
from . import outputs

__all__ = [
    'GetCertificateProfileResult',
    'AwaitableGetCertificateProfileResult',
    'get_certificate_profile',
    'get_certificate_profile_output',
]

@pulumi.output_type
class GetCertificateProfileResult:
    """
    Certificate profile resource.
    """
    def __init__(__self__, city=None, common_name=None, country=None, enhanced_key_usage=None, id=None, identity_validation_id=None, include_city=None, include_country=None, include_postal_code=None, include_state=None, include_street_address=None, name=None, organization=None, organization_unit=None, postal_code=None, profile_type=None, provisioning_state=None, state=None, status=None, street_address=None, system_data=None, type=None):
        if city and not isinstance(city, str):
            raise TypeError("Expected argument 'city' to be a str")
        pulumi.set(__self__, "city", city)
        if common_name and not isinstance(common_name, str):
            raise TypeError("Expected argument 'common_name' to be a str")
        pulumi.set(__self__, "common_name", common_name)
        if country and not isinstance(country, str):
            raise TypeError("Expected argument 'country' to be a str")
        pulumi.set(__self__, "country", country)
        if enhanced_key_usage and not isinstance(enhanced_key_usage, str):
            raise TypeError("Expected argument 'enhanced_key_usage' to be a str")
        pulumi.set(__self__, "enhanced_key_usage", enhanced_key_usage)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity_validation_id and not isinstance(identity_validation_id, str):
            raise TypeError("Expected argument 'identity_validation_id' to be a str")
        pulumi.set(__self__, "identity_validation_id", identity_validation_id)
        if include_city and not isinstance(include_city, bool):
            raise TypeError("Expected argument 'include_city' to be a bool")
        pulumi.set(__self__, "include_city", include_city)
        if include_country and not isinstance(include_country, bool):
            raise TypeError("Expected argument 'include_country' to be a bool")
        pulumi.set(__self__, "include_country", include_country)
        if include_postal_code and not isinstance(include_postal_code, bool):
            raise TypeError("Expected argument 'include_postal_code' to be a bool")
        pulumi.set(__self__, "include_postal_code", include_postal_code)
        if include_state and not isinstance(include_state, bool):
            raise TypeError("Expected argument 'include_state' to be a bool")
        pulumi.set(__self__, "include_state", include_state)
        if include_street_address and not isinstance(include_street_address, bool):
            raise TypeError("Expected argument 'include_street_address' to be a bool")
        pulumi.set(__self__, "include_street_address", include_street_address)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization and not isinstance(organization, str):
            raise TypeError("Expected argument 'organization' to be a str")
        pulumi.set(__self__, "organization", organization)
        if organization_unit and not isinstance(organization_unit, str):
            raise TypeError("Expected argument 'organization_unit' to be a str")
        pulumi.set(__self__, "organization_unit", organization_unit)
        if postal_code and not isinstance(postal_code, str):
            raise TypeError("Expected argument 'postal_code' to be a str")
        pulumi.set(__self__, "postal_code", postal_code)
        if profile_type and not isinstance(profile_type, str):
            raise TypeError("Expected argument 'profile_type' to be a str")
        pulumi.set(__self__, "profile_type", profile_type)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if street_address and not isinstance(street_address, str):
            raise TypeError("Expected argument 'street_address' to be a str")
        pulumi.set(__self__, "street_address", street_address)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def city(self) -> str:
        """
        Used as L in the certificate subject name.
        """
        return pulumi.get(self, "city")

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> str:
        """
        Used as CN in the certificate subject name.
        """
        return pulumi.get(self, "common_name")

    @property
    @pulumi.getter
    def country(self) -> str:
        """
        Used as C in the certificate subject name.
        """
        return pulumi.get(self, "country")

    @property
    @pulumi.getter(name="enhancedKeyUsage")
    def enhanced_key_usage(self) -> str:
        """
        Enhanced key usage of the certificate.
        """
        return pulumi.get(self, "enhanced_key_usage")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="identityValidationId")
    def identity_validation_id(self) -> Optional[str]:
        """
        Identity validation id used for the certificate subject name.
        """
        return pulumi.get(self, "identity_validation_id")

    @property
    @pulumi.getter(name="includeCity")
    def include_city(self) -> Optional[bool]:
        """
        Whether to include L in the certificate subject name. Applicable only for private trust, private trust ci profile types
        """
        return pulumi.get(self, "include_city")

    @property
    @pulumi.getter(name="includeCountry")
    def include_country(self) -> Optional[bool]:
        """
        Whether to include C in the certificate subject name. Applicable only for private trust, private trust ci profile types
        """
        return pulumi.get(self, "include_country")

    @property
    @pulumi.getter(name="includePostalCode")
    def include_postal_code(self) -> Optional[bool]:
        """
        Whether to include PC in the certificate subject name.
        """
        return pulumi.get(self, "include_postal_code")

    @property
    @pulumi.getter(name="includeState")
    def include_state(self) -> Optional[bool]:
        """
        Whether to include S in the certificate subject name. Applicable only for private trust, private trust ci profile types
        """
        return pulumi.get(self, "include_state")

    @property
    @pulumi.getter(name="includeStreetAddress")
    def include_street_address(self) -> Optional[bool]:
        """
        Whether to include STREET in the certificate subject name.
        """
        return pulumi.get(self, "include_street_address")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def organization(self) -> str:
        """
        Used as O in the certificate subject name.
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="organizationUnit")
    def organization_unit(self) -> str:
        """
        Used as OU in the private trust certificate subject name.
        """
        return pulumi.get(self, "organization_unit")

    @property
    @pulumi.getter(name="postalCode")
    def postal_code(self) -> str:
        """
        Used as PC in the certificate subject name.
        """
        return pulumi.get(self, "postal_code")

    @property
    @pulumi.getter(name="profileType")
    def profile_type(self) -> str:
        """
        Profile type of the certificate.
        """
        return pulumi.get(self, "profile_type")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Status of the current operation on certificate profile.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Used as S in the certificate subject name.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the certificate profile.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="streetAddress")
    def street_address(self) -> str:
        """
        Used as STREET in the certificate subject name.
        """
        return pulumi.get(self, "street_address")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetCertificateProfileResult(GetCertificateProfileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCertificateProfileResult(
            city=self.city,
            common_name=self.common_name,
            country=self.country,
            enhanced_key_usage=self.enhanced_key_usage,
            id=self.id,
            identity_validation_id=self.identity_validation_id,
            include_city=self.include_city,
            include_country=self.include_country,
            include_postal_code=self.include_postal_code,
            include_state=self.include_state,
            include_street_address=self.include_street_address,
            name=self.name,
            organization=self.organization,
            organization_unit=self.organization_unit,
            postal_code=self.postal_code,
            profile_type=self.profile_type,
            provisioning_state=self.provisioning_state,
            state=self.state,
            status=self.status,
            street_address=self.street_address,
            system_data=self.system_data,
            type=self.type)


def get_certificate_profile(account_name: Optional[str] = None,
                            profile_name: Optional[str] = None,
                            resource_group_name: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCertificateProfileResult:
    """
    Get details of a certificate profile.

    Uses Azure REST API version 2024-02-05-preview.

    Other available API versions: 2024-09-30-preview.


    :param str account_name: Trusted Signing account name.
    :param str profile_name: Certificate profile name.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['profileName'] = profile_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:codesigning:getCertificateProfile', __args__, opts=opts, typ=GetCertificateProfileResult).value

    return AwaitableGetCertificateProfileResult(
        city=pulumi.get(__ret__, 'city'),
        common_name=pulumi.get(__ret__, 'common_name'),
        country=pulumi.get(__ret__, 'country'),
        enhanced_key_usage=pulumi.get(__ret__, 'enhanced_key_usage'),
        id=pulumi.get(__ret__, 'id'),
        identity_validation_id=pulumi.get(__ret__, 'identity_validation_id'),
        include_city=pulumi.get(__ret__, 'include_city'),
        include_country=pulumi.get(__ret__, 'include_country'),
        include_postal_code=pulumi.get(__ret__, 'include_postal_code'),
        include_state=pulumi.get(__ret__, 'include_state'),
        include_street_address=pulumi.get(__ret__, 'include_street_address'),
        name=pulumi.get(__ret__, 'name'),
        organization=pulumi.get(__ret__, 'organization'),
        organization_unit=pulumi.get(__ret__, 'organization_unit'),
        postal_code=pulumi.get(__ret__, 'postal_code'),
        profile_type=pulumi.get(__ret__, 'profile_type'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        state=pulumi.get(__ret__, 'state'),
        status=pulumi.get(__ret__, 'status'),
        street_address=pulumi.get(__ret__, 'street_address'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_certificate_profile_output(account_name: Optional[pulumi.Input[str]] = None,
                                   profile_name: Optional[pulumi.Input[str]] = None,
                                   resource_group_name: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCertificateProfileResult]:
    """
    Get details of a certificate profile.

    Uses Azure REST API version 2024-02-05-preview.

    Other available API versions: 2024-09-30-preview.


    :param str account_name: Trusted Signing account name.
    :param str profile_name: Certificate profile name.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['profileName'] = profile_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:codesigning:getCertificateProfile', __args__, opts=opts, typ=GetCertificateProfileResult)
    return __ret__.apply(lambda __response__: GetCertificateProfileResult(
        city=pulumi.get(__response__, 'city'),
        common_name=pulumi.get(__response__, 'common_name'),
        country=pulumi.get(__response__, 'country'),
        enhanced_key_usage=pulumi.get(__response__, 'enhanced_key_usage'),
        id=pulumi.get(__response__, 'id'),
        identity_validation_id=pulumi.get(__response__, 'identity_validation_id'),
        include_city=pulumi.get(__response__, 'include_city'),
        include_country=pulumi.get(__response__, 'include_country'),
        include_postal_code=pulumi.get(__response__, 'include_postal_code'),
        include_state=pulumi.get(__response__, 'include_state'),
        include_street_address=pulumi.get(__response__, 'include_street_address'),
        name=pulumi.get(__response__, 'name'),
        organization=pulumi.get(__response__, 'organization'),
        organization_unit=pulumi.get(__response__, 'organization_unit'),
        postal_code=pulumi.get(__response__, 'postal_code'),
        profile_type=pulumi.get(__response__, 'profile_type'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        state=pulumi.get(__response__, 'state'),
        status=pulumi.get(__response__, 'status'),
        street_address=pulumi.get(__response__, 'street_address'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
