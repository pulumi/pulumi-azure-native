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

__all__ = ['ShareArgs', 'Share']

@pulumi.input_type
class ShareArgs:
    def __init__(__self__, *,
                 account_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 share_kind: Optional[pulumi.Input[Union[builtins.str, 'ShareKind']]] = None,
                 share_name: Optional[pulumi.Input[builtins.str]] = None,
                 terms: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Share resource.
        :param pulumi.Input[builtins.str] account_name: The name of the share account.
        :param pulumi.Input[builtins.str] resource_group_name: The resource group name.
        :param pulumi.Input[builtins.str] description: Share description.
        :param pulumi.Input[Union[builtins.str, 'ShareKind']] share_kind: Share kind.
        :param pulumi.Input[builtins.str] share_name: The name of the share.
        :param pulumi.Input[builtins.str] terms: Share terms.
        """
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if share_kind is not None:
            pulumi.set(__self__, "share_kind", share_kind)
        if share_name is not None:
            pulumi.set(__self__, "share_name", share_name)
        if terms is not None:
            pulumi.set(__self__, "terms", terms)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the share account.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The resource group name.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Share description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="shareKind")
    def share_kind(self) -> Optional[pulumi.Input[Union[builtins.str, 'ShareKind']]]:
        """
        Share kind.
        """
        return pulumi.get(self, "share_kind")

    @share_kind.setter
    def share_kind(self, value: Optional[pulumi.Input[Union[builtins.str, 'ShareKind']]]):
        pulumi.set(self, "share_kind", value)

    @property
    @pulumi.getter(name="shareName")
    def share_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the share.
        """
        return pulumi.get(self, "share_name")

    @share_name.setter
    def share_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "share_name", value)

    @property
    @pulumi.getter
    def terms(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Share terms.
        """
        return pulumi.get(self, "terms")

    @terms.setter
    def terms(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "terms", value)


@pulumi.type_token("azure-native:datashare:Share")
class Share(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 share_kind: Optional[pulumi.Input[Union[builtins.str, 'ShareKind']]] = None,
                 share_name: Optional[pulumi.Input[builtins.str]] = None,
                 terms: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        A share data transfer object.

        Uses Azure REST API version 2021-08-01. In version 2.x of the Azure Native provider, it used API version 2021-08-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] account_name: The name of the share account.
        :param pulumi.Input[builtins.str] description: Share description.
        :param pulumi.Input[builtins.str] resource_group_name: The resource group name.
        :param pulumi.Input[Union[builtins.str, 'ShareKind']] share_kind: Share kind.
        :param pulumi.Input[builtins.str] share_name: The name of the share.
        :param pulumi.Input[builtins.str] terms: Share terms.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ShareArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A share data transfer object.

        Uses Azure REST API version 2021-08-01. In version 2.x of the Azure Native provider, it used API version 2021-08-01.

        :param str resource_name: The name of the resource.
        :param ShareArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ShareArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 share_kind: Optional[pulumi.Input[Union[builtins.str, 'ShareKind']]] = None,
                 share_name: Optional[pulumi.Input[builtins.str]] = None,
                 terms: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ShareArgs.__new__(ShareArgs)

            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__.__dict__["account_name"] = account_name
            __props__.__dict__["description"] = description
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["share_kind"] = share_kind
            __props__.__dict__["share_name"] = share_name
            __props__.__dict__["terms"] = terms
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["user_email"] = None
            __props__.__dict__["user_name"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:datashare/v20181101preview:Share"), pulumi.Alias(type_="azure-native:datashare/v20191101:Share"), pulumi.Alias(type_="azure-native:datashare/v20200901:Share"), pulumi.Alias(type_="azure-native:datashare/v20201001preview:Share"), pulumi.Alias(type_="azure-native:datashare/v20210801:Share")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Share, __self__).__init__(
            'azure-native:datashare:Share',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Share':
        """
        Get an existing Share resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ShareArgs.__new__(ShareArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["share_kind"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["terms"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["user_email"] = None
        __props__.__dict__["user_name"] = None
        return Share(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        Time at which the share was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Share description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the azure resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Gets or sets the provisioning state
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="shareKind")
    def share_kind(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Share kind.
        """
        return pulumi.get(self, "share_kind")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        System Data of the Azure resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def terms(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Share terms.
        """
        return pulumi.get(self, "terms")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Type of the azure resource
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userEmail")
    def user_email(self) -> pulumi.Output[builtins.str]:
        """
        Email of the user who created the resource
        """
        return pulumi.get(self, "user_email")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the user who created the resource
        """
        return pulumi.get(self, "user_name")

