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

__all__ = ['CopilotSettingArgs', 'CopilotSetting']

@pulumi.input_type
class CopilotSettingArgs:
    def __init__(__self__, *,
                 access_control_enabled: pulumi.Input[builtins.bool]):
        """
        The set of arguments for constructing a CopilotSetting resource.
        :param pulumi.Input[builtins.bool] access_control_enabled: Boolean indicating if role-based access control is enabled for copilot in this tenant.
        """
        pulumi.set(__self__, "access_control_enabled", access_control_enabled)

    @property
    @pulumi.getter(name="accessControlEnabled")
    def access_control_enabled(self) -> pulumi.Input[builtins.bool]:
        """
        Boolean indicating if role-based access control is enabled for copilot in this tenant.
        """
        return pulumi.get(self, "access_control_enabled")

    @access_control_enabled.setter
    def access_control_enabled(self, value: pulumi.Input[builtins.bool]):
        pulumi.set(self, "access_control_enabled", value)


@pulumi.type_token("azure-native:portalservices:CopilotSetting")
class CopilotSetting(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_control_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        """
        The copilot settings tenant resource definition.

        Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2024-04-01-preview.

        Other available API versions: 2024-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native portalservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] access_control_enabled: Boolean indicating if role-based access control is enabled for copilot in this tenant.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CopilotSettingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The copilot settings tenant resource definition.

        Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2024-04-01-preview.

        Other available API versions: 2024-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native portalservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param CopilotSettingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CopilotSettingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_control_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CopilotSettingArgs.__new__(CopilotSettingArgs)

            if access_control_enabled is None and not opts.urn:
                raise TypeError("Missing required property 'access_control_enabled'")
            __props__.__dict__["access_control_enabled"] = access_control_enabled
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:portalservices/v20240401:CopilotSetting"), pulumi.Alias(type_="azure-native:portalservices/v20240401preview:CopilotSetting")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(CopilotSetting, __self__).__init__(
            'azure-native:portalservices:CopilotSetting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CopilotSetting':
        """
        Get an existing CopilotSetting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CopilotSettingArgs.__new__(CopilotSettingArgs)

        __props__.__dict__["access_control_enabled"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return CopilotSetting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessControlEnabled")
    def access_control_enabled(self) -> pulumi.Output[builtins.bool]:
        """
        Boolean indicating if role-based access control is enabled for copilot in this tenant.
        """
        return pulumi.get(self, "access_control_enabled")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

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
        The status of the last provisioning operation performed on the resource.
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
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

