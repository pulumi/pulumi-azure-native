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

__all__ = ['TransformArgs', 'Transform']

@pulumi.input_type
class TransformArgs:
    def __init__(__self__, *,
                 account_name: pulumi.Input[builtins.str],
                 outputs: pulumi.Input[Sequence[pulumi.Input['TransformOutputArgs']]],
                 resource_group_name: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 transform_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Transform resource.
        :param pulumi.Input[builtins.str] account_name: The Media Services account name.
        :param pulumi.Input[Sequence[pulumi.Input['TransformOutputArgs']]] outputs: An array of one or more TransformOutputs that the Transform should generate.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group within the Azure subscription.
        :param pulumi.Input[builtins.str] description: An optional verbose description of the Transform.
        :param pulumi.Input[builtins.str] transform_name: The Transform name.
        """
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "outputs", outputs)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if transform_name is not None:
            pulumi.set(__self__, "transform_name", transform_name)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Input[builtins.str]:
        """
        The Media Services account name.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter
    def outputs(self) -> pulumi.Input[Sequence[pulumi.Input['TransformOutputArgs']]]:
        """
        An array of one or more TransformOutputs that the Transform should generate.
        """
        return pulumi.get(self, "outputs")

    @outputs.setter
    def outputs(self, value: pulumi.Input[Sequence[pulumi.Input['TransformOutputArgs']]]):
        pulumi.set(self, "outputs", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group within the Azure subscription.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        An optional verbose description of the Transform.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="transformName")
    def transform_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Transform name.
        """
        return pulumi.get(self, "transform_name")

    @transform_name.setter
    def transform_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "transform_name", value)


@pulumi.type_token("azure-native:media:Transform")
class Transform(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 outputs: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TransformOutputArgs', 'TransformOutputArgsDict']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 transform_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        A Transform encapsulates the rules or instructions for generating desired outputs from input media, such as by transcoding or by extracting insights. After the Transform is created, it can be applied to input media by creating Jobs.

        Uses Azure REST API version 2022-07-01. In version 2.x of the Azure Native provider, it used API version 2022-07-01.

        Other available API versions: 2018-03-30-preview, 2018-06-01-preview, 2018-07-01, 2020-05-01, 2021-06-01, 2021-11-01, 2022-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native media [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] account_name: The Media Services account name.
        :param pulumi.Input[builtins.str] description: An optional verbose description of the Transform.
        :param pulumi.Input[Sequence[pulumi.Input[Union['TransformOutputArgs', 'TransformOutputArgsDict']]]] outputs: An array of one or more TransformOutputs that the Transform should generate.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group within the Azure subscription.
        :param pulumi.Input[builtins.str] transform_name: The Transform name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TransformArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Transform encapsulates the rules or instructions for generating desired outputs from input media, such as by transcoding or by extracting insights. After the Transform is created, it can be applied to input media by creating Jobs.

        Uses Azure REST API version 2022-07-01. In version 2.x of the Azure Native provider, it used API version 2022-07-01.

        Other available API versions: 2018-03-30-preview, 2018-06-01-preview, 2018-07-01, 2020-05-01, 2021-06-01, 2021-11-01, 2022-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native media [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param TransformArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TransformArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 outputs: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TransformOutputArgs', 'TransformOutputArgsDict']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 transform_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TransformArgs.__new__(TransformArgs)

            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__.__dict__["account_name"] = account_name
            __props__.__dict__["description"] = description
            if outputs is None and not opts.urn:
                raise TypeError("Missing required property 'outputs'")
            __props__.__dict__["outputs"] = outputs
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["transform_name"] = transform_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["created"] = None
            __props__.__dict__["last_modified"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:media/v20180330preview:Transform"), pulumi.Alias(type_="azure-native:media/v20180601preview:Transform"), pulumi.Alias(type_="azure-native:media/v20180701:Transform"), pulumi.Alias(type_="azure-native:media/v20200501:Transform"), pulumi.Alias(type_="azure-native:media/v20210601:Transform"), pulumi.Alias(type_="azure-native:media/v20211101:Transform"), pulumi.Alias(type_="azure-native:media/v20220501preview:Transform"), pulumi.Alias(type_="azure-native:media/v20220701:Transform")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Transform, __self__).__init__(
            'azure-native:media:Transform',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Transform':
        """
        Get an existing Transform resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TransformArgs.__new__(TransformArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["created"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["last_modified"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["outputs"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return Transform(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[builtins.str]:
        """
        The UTC date and time when the Transform was created, in 'YYYY-MM-DDThh:mm:ssZ' format.
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        An optional verbose description of the Transform.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> pulumi.Output[builtins.str]:
        """
        The UTC date and time when the Transform was last updated, in 'YYYY-MM-DDThh:mm:ssZ' format.
        """
        return pulumi.get(self, "last_modified")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def outputs(self) -> pulumi.Output[Sequence['outputs.TransformOutputResponse']]:
        """
        An array of one or more TransformOutputs that the Transform should generate.
        """
        return pulumi.get(self, "outputs")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        The system metadata relating to this resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

