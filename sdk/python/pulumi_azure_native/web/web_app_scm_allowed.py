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

__all__ = ['WebAppScmAllowedArgs', 'WebAppScmAllowed']

@pulumi.input_type
class WebAppScmAllowedArgs:
    def __init__(__self__, *,
                 allow: pulumi.Input[builtins.bool],
                 name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 kind: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a WebAppScmAllowed resource.
        :param pulumi.Input[builtins.bool] allow: <code>true</code> to allow access to a publishing method; otherwise, <code>false</code>.
        :param pulumi.Input[builtins.str] name: Name of the app.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[builtins.str] kind: Kind of resource.
        """
        pulumi.set(__self__, "allow", allow)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)

    @property
    @pulumi.getter
    def allow(self) -> pulumi.Input[builtins.bool]:
        """
        <code>true</code> to allow access to a publishing method; otherwise, <code>false</code>.
        """
        return pulumi.get(self, "allow")

    @allow.setter
    def allow(self, value: pulumi.Input[builtins.bool]):
        pulumi.set(self, "allow", value)

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


@pulumi.type_token("azure-native:web:WebAppScmAllowed")
class WebAppScmAllowed(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow: Optional[pulumi.Input[builtins.bool]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Publishing Credentials Policies parameters.

        Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.

        Other available API versions: 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] allow: <code>true</code> to allow access to a publishing method; otherwise, <code>false</code>.
        :param pulumi.Input[builtins.str] kind: Kind of resource.
        :param pulumi.Input[builtins.str] name: Name of the app.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the resource group to which the resource belongs.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebAppScmAllowedArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Publishing Credentials Policies parameters.

        Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.

        Other available API versions: 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param WebAppScmAllowedArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebAppScmAllowedArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow: Optional[pulumi.Input[builtins.bool]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebAppScmAllowedArgs.__new__(WebAppScmAllowedArgs)

            if allow is None and not opts.urn:
                raise TypeError("Missing required property 'allow'")
            __props__.__dict__["allow"] = allow
            __props__.__dict__["kind"] = kind
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:web/v20190801:WebAppScmAllowed"), pulumi.Alias(type_="azure-native:web/v20200601:WebAppScmAllowed"), pulumi.Alias(type_="azure-native:web/v20200901:WebAppScmAllowed"), pulumi.Alias(type_="azure-native:web/v20201001:WebAppScmAllowed"), pulumi.Alias(type_="azure-native:web/v20201201:WebAppScmAllowed"), pulumi.Alias(type_="azure-native:web/v20210101:WebAppScmAllowed"), pulumi.Alias(type_="azure-native:web/v20210115:WebAppScmAllowed"), pulumi.Alias(type_="azure-native:web/v20210201:WebAppScmAllowed"), pulumi.Alias(type_="azure-native:web/v20210301:WebAppScmAllowed"), pulumi.Alias(type_="azure-native:web/v20220301:WebAppScmAllowed"), pulumi.Alias(type_="azure-native:web/v20220901:WebAppScmAllowed"), pulumi.Alias(type_="azure-native:web/v20230101:WebAppScmAllowed"), pulumi.Alias(type_="azure-native:web/v20231201:WebAppScmAllowed"), pulumi.Alias(type_="azure-native:web/v20240401:WebAppScmAllowed"), pulumi.Alias(type_="azure-native:web/v20241101:WebAppScmAllowed")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(WebAppScmAllowed, __self__).__init__(
            'azure-native:web:WebAppScmAllowed',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WebAppScmAllowed':
        """
        Get an existing WebAppScmAllowed resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WebAppScmAllowedArgs.__new__(WebAppScmAllowedArgs)

        __props__.__dict__["allow"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["type"] = None
        return WebAppScmAllowed(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def allow(self) -> pulumi.Output[builtins.bool]:
        """
        <code>true</code> to allow access to a publishing method; otherwise, <code>false</code>.
        """
        return pulumi.get(self, "allow")

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
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

