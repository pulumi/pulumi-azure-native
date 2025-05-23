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

__all__ = ['AccountArgs', 'Account']

@pulumi.input_type
class AccountArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 account_name: Optional[pulumi.Input[builtins.str]] = None,
                 local_auth: Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 regional_affinity: Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]] = None,
                 reporting: Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]] = None,
                 scalable_execution: Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a Account resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] account_name: Name of account.
        :param pulumi.Input[Union[builtins.str, 'EnablementStatus']] local_auth: When enabled, this feature allows the workspace to use local auth (through service access token) for executing operations.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[Union[builtins.str, 'EnablementStatus']] regional_affinity: This property sets the connection region for Playwright client workers to cloud-hosted browsers. If enabled, workers connect to browsers in the closest Azure region, ensuring lower latency. If disabled, workers connect to browsers in the Azure region in which the workspace was initially created.
        :param pulumi.Input[Union[builtins.str, 'EnablementStatus']] reporting: When enabled, this feature allows the workspace to upload and display test results, including artifacts like traces and screenshots, in the Playwright portal. This enables faster and more efficient troubleshooting.
        :param pulumi.Input[Union[builtins.str, 'EnablementStatus']] scalable_execution: When enabled, Playwright client workers can connect to cloud-hosted browsers. This can increase the number of parallel workers for a test run, significantly minimizing test completion durations.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if account_name is not None:
            pulumi.set(__self__, "account_name", account_name)
        if local_auth is None:
            local_auth = 'Disabled'
        if local_auth is not None:
            pulumi.set(__self__, "local_auth", local_auth)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if regional_affinity is None:
            regional_affinity = 'Enabled'
        if regional_affinity is not None:
            pulumi.set(__self__, "regional_affinity", regional_affinity)
        if reporting is None:
            reporting = 'Enabled'
        if reporting is not None:
            pulumi.set(__self__, "reporting", reporting)
        if scalable_execution is None:
            scalable_execution = 'Enabled'
        if scalable_execution is not None:
            pulumi.set(__self__, "scalable_execution", scalable_execution)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

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
    @pulumi.getter(name="accountName")
    def account_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of account.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="localAuth")
    def local_auth(self) -> Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]]:
        """
        When enabled, this feature allows the workspace to use local auth (through service access token) for executing operations.
        """
        return pulumi.get(self, "local_auth")

    @local_auth.setter
    def local_auth(self, value: Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]]):
        pulumi.set(self, "local_auth", value)

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
    @pulumi.getter(name="regionalAffinity")
    def regional_affinity(self) -> Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]]:
        """
        This property sets the connection region for Playwright client workers to cloud-hosted browsers. If enabled, workers connect to browsers in the closest Azure region, ensuring lower latency. If disabled, workers connect to browsers in the Azure region in which the workspace was initially created.
        """
        return pulumi.get(self, "regional_affinity")

    @regional_affinity.setter
    def regional_affinity(self, value: Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]]):
        pulumi.set(self, "regional_affinity", value)

    @property
    @pulumi.getter
    def reporting(self) -> Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]]:
        """
        When enabled, this feature allows the workspace to upload and display test results, including artifacts like traces and screenshots, in the Playwright portal. This enables faster and more efficient troubleshooting.
        """
        return pulumi.get(self, "reporting")

    @reporting.setter
    def reporting(self, value: Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]]):
        pulumi.set(self, "reporting", value)

    @property
    @pulumi.getter(name="scalableExecution")
    def scalable_execution(self) -> Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]]:
        """
        When enabled, Playwright client workers can connect to cloud-hosted browsers. This can increase the number of parallel workers for a test run, significantly minimizing test completion durations.
        """
        return pulumi.get(self, "scalable_execution")

    @scalable_execution.setter
    def scalable_execution(self, value: Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]]):
        pulumi.set(self, "scalable_execution", value)

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


@pulumi.type_token("azure-native:azureplaywrightservice:Account")
class Account(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[builtins.str]] = None,
                 local_auth: Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 regional_affinity: Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]] = None,
                 reporting: Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 scalable_execution: Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        A Playwright service account resource.

        Uses Azure REST API version 2024-12-01. In version 2.x of the Azure Native provider, it used API version 2023-10-01-preview.

        Other available API versions: 2023-10-01-preview, 2024-02-01-preview, 2024-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azureplaywrightservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] account_name: Name of account.
        :param pulumi.Input[Union[builtins.str, 'EnablementStatus']] local_auth: When enabled, this feature allows the workspace to use local auth (through service access token) for executing operations.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[Union[builtins.str, 'EnablementStatus']] regional_affinity: This property sets the connection region for Playwright client workers to cloud-hosted browsers. If enabled, workers connect to browsers in the closest Azure region, ensuring lower latency. If disabled, workers connect to browsers in the Azure region in which the workspace was initially created.
        :param pulumi.Input[Union[builtins.str, 'EnablementStatus']] reporting: When enabled, this feature allows the workspace to upload and display test results, including artifacts like traces and screenshots, in the Playwright portal. This enables faster and more efficient troubleshooting.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Union[builtins.str, 'EnablementStatus']] scalable_execution: When enabled, Playwright client workers can connect to cloud-hosted browsers. This can increase the number of parallel workers for a test run, significantly minimizing test completion durations.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Playwright service account resource.

        Uses Azure REST API version 2024-12-01. In version 2.x of the Azure Native provider, it used API version 2023-10-01-preview.

        Other available API versions: 2023-10-01-preview, 2024-02-01-preview, 2024-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azureplaywrightservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param AccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[builtins.str]] = None,
                 local_auth: Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 regional_affinity: Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]] = None,
                 reporting: Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 scalable_execution: Optional[pulumi.Input[Union[builtins.str, 'EnablementStatus']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccountArgs.__new__(AccountArgs)

            __props__.__dict__["account_name"] = account_name
            if local_auth is None:
                local_auth = 'Disabled'
            __props__.__dict__["local_auth"] = local_auth
            __props__.__dict__["location"] = location
            if regional_affinity is None:
                regional_affinity = 'Enabled'
            __props__.__dict__["regional_affinity"] = regional_affinity
            if reporting is None:
                reporting = 'Enabled'
            __props__.__dict__["reporting"] = reporting
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if scalable_execution is None:
                scalable_execution = 'Enabled'
            __props__.__dict__["scalable_execution"] = scalable_execution
            __props__.__dict__["tags"] = tags
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["dashboard_uri"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:azureplaywrightservice/v20231001preview:Account"), pulumi.Alias(type_="azure-native:azureplaywrightservice/v20240201preview:Account"), pulumi.Alias(type_="azure-native:azureplaywrightservice/v20240801preview:Account"), pulumi.Alias(type_="azure-native:azureplaywrightservice/v20241201:Account")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Account, __self__).__init__(
            'azure-native:azureplaywrightservice:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Account':
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AccountArgs.__new__(AccountArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["dashboard_uri"] = None
        __props__.__dict__["local_auth"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["regional_affinity"] = None
        __props__.__dict__["reporting"] = None
        __props__.__dict__["scalable_execution"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return Account(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="dashboardUri")
    def dashboard_uri(self) -> pulumi.Output[builtins.str]:
        """
        The Playwright testing dashboard URI for the account resource.
        """
        return pulumi.get(self, "dashboard_uri")

    @property
    @pulumi.getter(name="localAuth")
    def local_auth(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        When enabled, this feature allows the workspace to use local auth (through service access token) for executing operations.
        """
        return pulumi.get(self, "local_auth")

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
        The status of the last operation.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="regionalAffinity")
    def regional_affinity(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        This property sets the connection region for Playwright client workers to cloud-hosted browsers. If enabled, workers connect to browsers in the closest Azure region, ensuring lower latency. If disabled, workers connect to browsers in the Azure region in which the workspace was initially created.
        """
        return pulumi.get(self, "regional_affinity")

    @property
    @pulumi.getter
    def reporting(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        When enabled, this feature allows the workspace to upload and display test results, including artifacts like traces and screenshots, in the Playwright portal. This enables faster and more efficient troubleshooting.
        """
        return pulumi.get(self, "reporting")

    @property
    @pulumi.getter(name="scalableExecution")
    def scalable_execution(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        When enabled, Playwright client workers can connect to cloud-hosted browsers. This can increase the number of parallel workers for a test run, significantly minimizing test completion durations.
        """
        return pulumi.get(self, "scalable_execution")

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
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

