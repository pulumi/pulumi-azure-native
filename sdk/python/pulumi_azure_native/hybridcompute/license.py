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

__all__ = ['LicenseInitArgs', 'License']

@pulumi.input_type
class LicenseInitArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 license_details: Optional[pulumi.Input['LicenseDetailsArgs']] = None,
                 license_name: Optional[pulumi.Input[builtins.str]] = None,
                 license_type: Optional[pulumi.Input[Union[builtins.str, 'LicenseType']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a License resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input['LicenseDetailsArgs'] license_details: Describes the properties of a License.
        :param pulumi.Input[builtins.str] license_name: The name of the license.
        :param pulumi.Input[Union[builtins.str, 'LicenseType']] license_type: The type of the license resource.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[builtins.str] tenant_id: Describes the tenant id.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if license_details is not None:
            pulumi.set(__self__, "license_details", license_details)
        if license_name is not None:
            pulumi.set(__self__, "license_name", license_name)
        if license_type is not None:
            pulumi.set(__self__, "license_type", license_type)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group. The name is case insensitive.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="licenseDetails")
    def license_details(self) -> Optional[pulumi.Input['LicenseDetailsArgs']]:
        """
        Describes the properties of a License.
        """
        return pulumi.get(self, "license_details")

    @license_details.setter
    def license_details(self, value: Optional[pulumi.Input['LicenseDetailsArgs']]):
        pulumi.set(self, "license_details", value)

    @property
    @pulumi.getter(name="licenseName")
    def license_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the license.
        """
        return pulumi.get(self, "license_name")

    @license_name.setter
    def license_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "license_name", value)

    @property
    @pulumi.getter(name="licenseType")
    def license_type(self) -> Optional[pulumi.Input[Union[builtins.str, 'LicenseType']]]:
        """
        The type of the license resource.
        """
        return pulumi.get(self, "license_type")

    @license_type.setter
    def license_type(self, value: Optional[pulumi.Input[Union[builtins.str, 'LicenseType']]]):
        pulumi.set(self, "license_type", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The geo-location where the resource lives
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
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Describes the tenant id.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "tenant_id", value)


@pulumi.type_token("azure-native:hybridcompute:License")
class License(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 license_details: Optional[pulumi.Input[Union['LicenseDetailsArgs', 'LicenseDetailsArgsDict']]] = None,
                 license_name: Optional[pulumi.Input[builtins.str]] = None,
                 license_type: Optional[pulumi.Input[Union[builtins.str, 'LicenseType']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Describes a license in a hybrid machine.

        Uses Azure REST API version 2024-07-10. In version 2.x of the Azure Native provider, it used API version 2023-06-20-preview.

        Other available API versions: 2023-06-20-preview, 2023-10-03-preview, 2024-03-31-preview, 2024-05-20-preview, 2024-07-31-preview, 2024-09-10-preview, 2024-11-10-preview, 2025-01-13, 2025-02-19-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridcompute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['LicenseDetailsArgs', 'LicenseDetailsArgsDict']] license_details: Describes the properties of a License.
        :param pulumi.Input[builtins.str] license_name: The name of the license.
        :param pulumi.Input[Union[builtins.str, 'LicenseType']] license_type: The type of the license resource.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[builtins.str] tenant_id: Describes the tenant id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LicenseInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Describes a license in a hybrid machine.

        Uses Azure REST API version 2024-07-10. In version 2.x of the Azure Native provider, it used API version 2023-06-20-preview.

        Other available API versions: 2023-06-20-preview, 2023-10-03-preview, 2024-03-31-preview, 2024-05-20-preview, 2024-07-31-preview, 2024-09-10-preview, 2024-11-10-preview, 2025-01-13, 2025-02-19-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridcompute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param LicenseInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LicenseInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 license_details: Optional[pulumi.Input[Union['LicenseDetailsArgs', 'LicenseDetailsArgsDict']]] = None,
                 license_name: Optional[pulumi.Input[builtins.str]] = None,
                 license_type: Optional[pulumi.Input[Union[builtins.str, 'LicenseType']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LicenseInitArgs.__new__(LicenseInitArgs)

            __props__.__dict__["license_details"] = license_details
            __props__.__dict__["license_name"] = license_name
            __props__.__dict__["license_type"] = license_type
            __props__.__dict__["location"] = location
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:hybridcompute/v20230620preview:License"), pulumi.Alias(type_="azure-native:hybridcompute/v20231003preview:License"), pulumi.Alias(type_="azure-native:hybridcompute/v20240331preview:License"), pulumi.Alias(type_="azure-native:hybridcompute/v20240520preview:License"), pulumi.Alias(type_="azure-native:hybridcompute/v20240710:License"), pulumi.Alias(type_="azure-native:hybridcompute/v20240731preview:License"), pulumi.Alias(type_="azure-native:hybridcompute/v20240910preview:License"), pulumi.Alias(type_="azure-native:hybridcompute/v20241110preview:License"), pulumi.Alias(type_="azure-native:hybridcompute/v20250113:License"), pulumi.Alias(type_="azure-native:hybridcompute/v20250219preview:License")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(License, __self__).__init__(
            'azure-native:hybridcompute:License',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'License':
        """
        Get an existing License resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LicenseInitArgs.__new__(LicenseInitArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["license_details"] = None
        __props__.__dict__["license_type"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["tenant_id"] = None
        __props__.__dict__["type"] = None
        return License(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="licenseDetails")
    def license_details(self) -> pulumi.Output[Optional['outputs.LicenseDetailsResponse']]:
        """
        Describes the properties of a License.
        """
        return pulumi.get(self, "license_details")

    @property
    @pulumi.getter(name="licenseType")
    def license_type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The type of the license resource.
        """
        return pulumi.get(self, "license_type")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        The provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Describes the tenant id.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

