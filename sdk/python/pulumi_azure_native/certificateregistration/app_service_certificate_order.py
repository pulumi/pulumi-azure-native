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
from ._enums import *
from ._inputs import *

__all__ = ['AppServiceCertificateOrderArgs', 'AppServiceCertificateOrder']

@pulumi.input_type
class AppServiceCertificateOrderArgs:
    def __init__(__self__, *,
                 product_type: pulumi.Input['CertificateProductType'],
                 resource_group_name: pulumi.Input[builtins.str],
                 auto_renew: Optional[pulumi.Input[builtins.bool]] = None,
                 certificate_order_name: Optional[pulumi.Input[builtins.str]] = None,
                 certificates: Optional[pulumi.Input[Mapping[str, pulumi.Input['AppServiceCertificateArgs']]]] = None,
                 csr: Optional[pulumi.Input[builtins.str]] = None,
                 distinguished_name: Optional[pulumi.Input[builtins.str]] = None,
                 key_size: Optional[pulumi.Input[builtins.int]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 validity_in_years: Optional[pulumi.Input[builtins.int]] = None):
        """
        The set of arguments for constructing a AppServiceCertificateOrder resource.
        :param pulumi.Input['CertificateProductType'] product_type: Certificate product type.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[builtins.bool] auto_renew: <code>true</code> if the certificate should be automatically renewed when it expires; otherwise, <code>false</code>.
        :param pulumi.Input[builtins.str] certificate_order_name: Name of the certificate order.
        :param pulumi.Input[Mapping[str, pulumi.Input['AppServiceCertificateArgs']]] certificates: State of the Key Vault secret.
        :param pulumi.Input[builtins.str] csr: Last CSR that was created for this order.
        :param pulumi.Input[builtins.str] distinguished_name: Certificate distinguished name.
        :param pulumi.Input[builtins.int] key_size: Certificate key size.
        :param pulumi.Input[builtins.str] kind: Kind of resource. If the resource is an app, you can refer to https://github.com/Azure/app-service-linux-docs/blob/master/Things_You_Should_Know/kind_property.md#app-service-resource-kind-reference for details supported values for kind.
        :param pulumi.Input[builtins.str] location: Resource Location.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[builtins.int] validity_in_years: Duration in years (must be 1).
        """
        pulumi.set(__self__, "product_type", product_type)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if auto_renew is None:
            auto_renew = True
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if certificate_order_name is not None:
            pulumi.set(__self__, "certificate_order_name", certificate_order_name)
        if certificates is not None:
            pulumi.set(__self__, "certificates", certificates)
        if csr is not None:
            pulumi.set(__self__, "csr", csr)
        if distinguished_name is not None:
            pulumi.set(__self__, "distinguished_name", distinguished_name)
        if key_size is None:
            key_size = 2048
        if key_size is not None:
            pulumi.set(__self__, "key_size", key_size)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if validity_in_years is None:
            validity_in_years = 1
        if validity_in_years is not None:
            pulumi.set(__self__, "validity_in_years", validity_in_years)

    @property
    @pulumi.getter(name="productType")
    def product_type(self) -> pulumi.Input['CertificateProductType']:
        """
        Certificate product type.
        """
        return pulumi.get(self, "product_type")

    @product_type.setter
    def product_type(self, value: pulumi.Input['CertificateProductType']):
        pulumi.set(self, "product_type", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the resource group to which the resource belongs.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        <code>true</code> if the certificate should be automatically renewed when it expires; otherwise, <code>false</code>.
        """
        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter(name="certificateOrderName")
    def certificate_order_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the certificate order.
        """
        return pulumi.get(self, "certificate_order_name")

    @certificate_order_name.setter
    def certificate_order_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "certificate_order_name", value)

    @property
    @pulumi.getter
    def certificates(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['AppServiceCertificateArgs']]]]:
        """
        State of the Key Vault secret.
        """
        return pulumi.get(self, "certificates")

    @certificates.setter
    def certificates(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['AppServiceCertificateArgs']]]]):
        pulumi.set(self, "certificates", value)

    @property
    @pulumi.getter
    def csr(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Last CSR that was created for this order.
        """
        return pulumi.get(self, "csr")

    @csr.setter
    def csr(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "csr", value)

    @property
    @pulumi.getter(name="distinguishedName")
    def distinguished_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Certificate distinguished name.
        """
        return pulumi.get(self, "distinguished_name")

    @distinguished_name.setter
    def distinguished_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "distinguished_name", value)

    @property
    @pulumi.getter(name="keySize")
    def key_size(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Certificate key size.
        """
        return pulumi.get(self, "key_size")

    @key_size.setter
    def key_size(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "key_size", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Kind of resource. If the resource is an app, you can refer to https://github.com/Azure/app-service-linux-docs/blob/master/Things_You_Should_Know/kind_property.md#app-service-resource-kind-reference for details supported values for kind.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Resource Location.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="validityInYears")
    def validity_in_years(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Duration in years (must be 1).
        """
        return pulumi.get(self, "validity_in_years")

    @validity_in_years.setter
    def validity_in_years(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "validity_in_years", value)


@pulumi.type_token("azure-native:certificateregistration:AppServiceCertificateOrder")
class AppServiceCertificateOrder(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew: Optional[pulumi.Input[builtins.bool]] = None,
                 certificate_order_name: Optional[pulumi.Input[builtins.str]] = None,
                 certificates: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union['AppServiceCertificateArgs', 'AppServiceCertificateArgsDict']]]]] = None,
                 csr: Optional[pulumi.Input[builtins.str]] = None,
                 distinguished_name: Optional[pulumi.Input[builtins.str]] = None,
                 key_size: Optional[pulumi.Input[builtins.int]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 product_type: Optional[pulumi.Input['CertificateProductType']] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 validity_in_years: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        """
        SSL certificate purchase order.

        Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.

        Other available API versions: 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native certificateregistration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] auto_renew: <code>true</code> if the certificate should be automatically renewed when it expires; otherwise, <code>false</code>.
        :param pulumi.Input[builtins.str] certificate_order_name: Name of the certificate order.
        :param pulumi.Input[Mapping[str, pulumi.Input[Union['AppServiceCertificateArgs', 'AppServiceCertificateArgsDict']]]] certificates: State of the Key Vault secret.
        :param pulumi.Input[builtins.str] csr: Last CSR that was created for this order.
        :param pulumi.Input[builtins.str] distinguished_name: Certificate distinguished name.
        :param pulumi.Input[builtins.int] key_size: Certificate key size.
        :param pulumi.Input[builtins.str] kind: Kind of resource. If the resource is an app, you can refer to https://github.com/Azure/app-service-linux-docs/blob/master/Things_You_Should_Know/kind_property.md#app-service-resource-kind-reference for details supported values for kind.
        :param pulumi.Input[builtins.str] location: Resource Location.
        :param pulumi.Input['CertificateProductType'] product_type: Certificate product type.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[builtins.int] validity_in_years: Duration in years (must be 1).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppServiceCertificateOrderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        SSL certificate purchase order.

        Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.

        Other available API versions: 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native certificateregistration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param AppServiceCertificateOrderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppServiceCertificateOrderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew: Optional[pulumi.Input[builtins.bool]] = None,
                 certificate_order_name: Optional[pulumi.Input[builtins.str]] = None,
                 certificates: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union['AppServiceCertificateArgs', 'AppServiceCertificateArgsDict']]]]] = None,
                 csr: Optional[pulumi.Input[builtins.str]] = None,
                 distinguished_name: Optional[pulumi.Input[builtins.str]] = None,
                 key_size: Optional[pulumi.Input[builtins.int]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 product_type: Optional[pulumi.Input['CertificateProductType']] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 validity_in_years: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppServiceCertificateOrderArgs.__new__(AppServiceCertificateOrderArgs)

            if auto_renew is None:
                auto_renew = True
            __props__.__dict__["auto_renew"] = auto_renew
            __props__.__dict__["certificate_order_name"] = certificate_order_name
            __props__.__dict__["certificates"] = certificates
            __props__.__dict__["csr"] = csr
            __props__.__dict__["distinguished_name"] = distinguished_name
            if key_size is None:
                key_size = 2048
            __props__.__dict__["key_size"] = key_size
            __props__.__dict__["kind"] = kind
            __props__.__dict__["location"] = location
            if product_type is None and not opts.urn:
                raise TypeError("Missing required property 'product_type'")
            __props__.__dict__["product_type"] = product_type
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            if validity_in_years is None:
                validity_in_years = 1
            __props__.__dict__["validity_in_years"] = validity_in_years
            __props__.__dict__["app_service_certificate_not_renewable_reasons"] = None
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["contact"] = None
            __props__.__dict__["domain_verification_token"] = None
            __props__.__dict__["expiration_time"] = None
            __props__.__dict__["intermediate"] = None
            __props__.__dict__["is_private_key_external"] = None
            __props__.__dict__["last_certificate_issuance_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["next_auto_renewal_time_stamp"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["root"] = None
            __props__.__dict__["serial_number"] = None
            __props__.__dict__["signed_certificate"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:certificateregistration/v20150801:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-native:certificateregistration/v20180201:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-native:certificateregistration/v20190801:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-native:certificateregistration/v20200601:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-native:certificateregistration/v20200901:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-native:certificateregistration/v20201001:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-native:certificateregistration/v20201201:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-native:certificateregistration/v20210101:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-native:certificateregistration/v20210115:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-native:certificateregistration/v20210201:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-native:certificateregistration/v20210301:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-native:certificateregistration/v20220301:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-native:certificateregistration/v20220901:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-native:certificateregistration/v20230101:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-native:certificateregistration/v20231201:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-native:certificateregistration/v20240401:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-native:certificateregistration/v20241101:AppServiceCertificateOrder")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AppServiceCertificateOrder, __self__).__init__(
            'azure-native:certificateregistration:AppServiceCertificateOrder',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AppServiceCertificateOrder':
        """
        Get an existing AppServiceCertificateOrder resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AppServiceCertificateOrderArgs.__new__(AppServiceCertificateOrderArgs)

        __props__.__dict__["app_service_certificate_not_renewable_reasons"] = None
        __props__.__dict__["auto_renew"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["certificates"] = None
        __props__.__dict__["contact"] = None
        __props__.__dict__["csr"] = None
        __props__.__dict__["distinguished_name"] = None
        __props__.__dict__["domain_verification_token"] = None
        __props__.__dict__["expiration_time"] = None
        __props__.__dict__["intermediate"] = None
        __props__.__dict__["is_private_key_external"] = None
        __props__.__dict__["key_size"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["last_certificate_issuance_time"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["next_auto_renewal_time_stamp"] = None
        __props__.__dict__["product_type"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["root"] = None
        __props__.__dict__["serial_number"] = None
        __props__.__dict__["signed_certificate"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["validity_in_years"] = None
        return AppServiceCertificateOrder(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appServiceCertificateNotRenewableReasons")
    def app_service_certificate_not_renewable_reasons(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        Reasons why App Service Certificate is not renewable at the current moment.
        """
        return pulumi.get(self, "app_service_certificate_not_renewable_reasons")

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        <code>true</code> if the certificate should be automatically renewed when it expires; otherwise, <code>false</code>.
        """
        return pulumi.get(self, "auto_renew")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def certificates(self) -> pulumi.Output[Optional[Mapping[str, 'outputs.AppServiceCertificateResponse']]]:
        """
        State of the Key Vault secret.
        """
        return pulumi.get(self, "certificates")

    @property
    @pulumi.getter
    def contact(self) -> pulumi.Output['outputs.CertificateOrderContactResponse']:
        """
        Contact info
        """
        return pulumi.get(self, "contact")

    @property
    @pulumi.getter
    def csr(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Last CSR that was created for this order.
        """
        return pulumi.get(self, "csr")

    @property
    @pulumi.getter(name="distinguishedName")
    def distinguished_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Certificate distinguished name.
        """
        return pulumi.get(self, "distinguished_name")

    @property
    @pulumi.getter(name="domainVerificationToken")
    def domain_verification_token(self) -> pulumi.Output[builtins.str]:
        """
        Domain verification token.
        """
        return pulumi.get(self, "domain_verification_token")

    @property
    @pulumi.getter(name="expirationTime")
    def expiration_time(self) -> pulumi.Output[builtins.str]:
        """
        Certificate expiration time.
        """
        return pulumi.get(self, "expiration_time")

    @property
    @pulumi.getter
    def intermediate(self) -> pulumi.Output['outputs.CertificateDetailsResponse']:
        """
        Intermediate certificate.
        """
        return pulumi.get(self, "intermediate")

    @property
    @pulumi.getter(name="isPrivateKeyExternal")
    def is_private_key_external(self) -> pulumi.Output[builtins.bool]:
        """
        <code>true</code> if private key is external; otherwise, <code>false</code>.
        """
        return pulumi.get(self, "is_private_key_external")

    @property
    @pulumi.getter(name="keySize")
    def key_size(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        Certificate key size.
        """
        return pulumi.get(self, "key_size")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Kind of resource. If the resource is an app, you can refer to https://github.com/Azure/app-service-linux-docs/blob/master/Things_You_Should_Know/kind_property.md#app-service-resource-kind-reference for details supported values for kind.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="lastCertificateIssuanceTime")
    def last_certificate_issuance_time(self) -> pulumi.Output[builtins.str]:
        """
        Certificate last issuance time.
        """
        return pulumi.get(self, "last_certificate_issuance_time")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        Resource Location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Resource Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nextAutoRenewalTimeStamp")
    def next_auto_renewal_time_stamp(self) -> pulumi.Output[builtins.str]:
        """
        Time stamp when the certificate would be auto renewed next
        """
        return pulumi.get(self, "next_auto_renewal_time_stamp")

    @property
    @pulumi.getter(name="productType")
    def product_type(self) -> pulumi.Output[builtins.str]:
        """
        Certificate product type.
        """
        return pulumi.get(self, "product_type")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Status of certificate order.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def root(self) -> pulumi.Output['outputs.CertificateDetailsResponse']:
        """
        Root certificate.
        """
        return pulumi.get(self, "root")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> pulumi.Output[builtins.str]:
        """
        Current serial number of the certificate.
        """
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter(name="signedCertificate")
    def signed_certificate(self) -> pulumi.Output['outputs.CertificateDetailsResponse']:
        """
        Signed certificate.
        """
        return pulumi.get(self, "signed_certificate")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        Current order status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="validityInYears")
    def validity_in_years(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        Duration in years (must be 1).
        """
        return pulumi.get(self, "validity_in_years")

