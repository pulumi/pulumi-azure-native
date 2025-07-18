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

__all__ = ['WebAppSwiftVirtualNetworkConnectionArgs', 'WebAppSwiftVirtualNetworkConnection']

@pulumi.input_type
class WebAppSwiftVirtualNetworkConnectionArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 swift_supported: Optional[pulumi.Input[builtins.bool]] = None):
        """
        The set of arguments for constructing a WebAppSwiftVirtualNetworkConnection resource.
        :param pulumi.Input[builtins.str] name: Name of the app.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[builtins.str] kind: Kind of resource.
        :param pulumi.Input[builtins.str] subnet_resource_id: The Virtual Network subnet's resource ID. This is the subnet that this Web App will join. This subnet must have a delegation to Microsoft.Web/serverFarms defined first.
        :param pulumi.Input[builtins.bool] swift_supported: A flag that specifies if the scale unit this Web App is on supports Swift integration.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if subnet_resource_id is not None:
            pulumi.set(__self__, "subnet_resource_id", subnet_resource_id)
        if swift_supported is not None:
            pulumi.set(__self__, "swift_supported", swift_supported)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the app.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "name", value)

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
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Kind of resource.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="subnetResourceId")
    def subnet_resource_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Virtual Network subnet's resource ID. This is the subnet that this Web App will join. This subnet must have a delegation to Microsoft.Web/serverFarms defined first.
        """
        return pulumi.get(self, "subnet_resource_id")

    @subnet_resource_id.setter
    def subnet_resource_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "subnet_resource_id", value)

    @property
    @pulumi.getter(name="swiftSupported")
    def swift_supported(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        A flag that specifies if the scale unit this Web App is on supports Swift integration.
        """
        return pulumi.get(self, "swift_supported")

    @swift_supported.setter
    def swift_supported(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "swift_supported", value)


@pulumi.type_token("azure-native:web:WebAppSwiftVirtualNetworkConnection")
class WebAppSwiftVirtualNetworkConnection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 swift_supported: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        """
        Swift Virtual Network Contract. This is used to enable the new Swift way of doing virtual network integration.

        Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.

        Other available API versions: 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] kind: Kind of resource.
        :param pulumi.Input[builtins.str] name: Name of the app.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[builtins.str] subnet_resource_id: The Virtual Network subnet's resource ID. This is the subnet that this Web App will join. This subnet must have a delegation to Microsoft.Web/serverFarms defined first.
        :param pulumi.Input[builtins.bool] swift_supported: A flag that specifies if the scale unit this Web App is on supports Swift integration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebAppSwiftVirtualNetworkConnectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Swift Virtual Network Contract. This is used to enable the new Swift way of doing virtual network integration.

        Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.

        Other available API versions: 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param WebAppSwiftVirtualNetworkConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebAppSwiftVirtualNetworkConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 swift_supported: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebAppSwiftVirtualNetworkConnectionArgs.__new__(WebAppSwiftVirtualNetworkConnectionArgs)

            __props__.__dict__["kind"] = kind
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["subnet_resource_id"] = subnet_resource_id
            __props__.__dict__["swift_supported"] = swift_supported
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:web/v20180201:WebAppSwiftVirtualNetworkConnection"), pulumi.Alias(type_="azure-native:web/v20181101:WebAppSwiftVirtualNetworkConnection"), pulumi.Alias(type_="azure-native:web/v20190801:WebAppSwiftVirtualNetworkConnection"), pulumi.Alias(type_="azure-native:web/v20200601:WebAppSwiftVirtualNetworkConnection"), pulumi.Alias(type_="azure-native:web/v20200901:WebAppSwiftVirtualNetworkConnection"), pulumi.Alias(type_="azure-native:web/v20201001:WebAppSwiftVirtualNetworkConnection"), pulumi.Alias(type_="azure-native:web/v20201201:WebAppSwiftVirtualNetworkConnection"), pulumi.Alias(type_="azure-native:web/v20210101:WebAppSwiftVirtualNetworkConnection"), pulumi.Alias(type_="azure-native:web/v20210115:WebAppSwiftVirtualNetworkConnection"), pulumi.Alias(type_="azure-native:web/v20210201:WebAppSwiftVirtualNetworkConnection"), pulumi.Alias(type_="azure-native:web/v20210301:WebAppSwiftVirtualNetworkConnection"), pulumi.Alias(type_="azure-native:web/v20220301:WebAppSwiftVirtualNetworkConnection"), pulumi.Alias(type_="azure-native:web/v20220901:WebAppSwiftVirtualNetworkConnection"), pulumi.Alias(type_="azure-native:web/v20230101:WebAppSwiftVirtualNetworkConnection"), pulumi.Alias(type_="azure-native:web/v20231201:WebAppSwiftVirtualNetworkConnection"), pulumi.Alias(type_="azure-native:web/v20240401:WebAppSwiftVirtualNetworkConnection"), pulumi.Alias(type_="azure-native:web/v20241101:WebAppSwiftVirtualNetworkConnection")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(WebAppSwiftVirtualNetworkConnection, __self__).__init__(
            'azure-native:web:WebAppSwiftVirtualNetworkConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WebAppSwiftVirtualNetworkConnection':
        """
        Get an existing WebAppSwiftVirtualNetworkConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WebAppSwiftVirtualNetworkConnectionArgs.__new__(WebAppSwiftVirtualNetworkConnectionArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["subnet_resource_id"] = None
        __props__.__dict__["swift_supported"] = None
        __props__.__dict__["type"] = None
        return WebAppSwiftVirtualNetworkConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Kind of resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Resource Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="subnetResourceId")
    def subnet_resource_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The Virtual Network subnet's resource ID. This is the subnet that this Web App will join. This subnet must have a delegation to Microsoft.Web/serverFarms defined first.
        """
        return pulumi.get(self, "subnet_resource_id")

    @property
    @pulumi.getter(name="swiftSupported")
    def swift_supported(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        A flag that specifies if the scale unit this Web App is on supports Swift integration.
        """
        return pulumi.get(self, "swift_supported")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

