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
from ._enums import *
from ._inputs import *

__all__ = ['MediaGraphArgs', 'MediaGraph']

@pulumi.input_type
class MediaGraphArgs:
    def __init__(__self__, *,
                 account_name: pulumi.Input[str],
                 resource_group_name: pulumi.Input[str],
                 sinks: pulumi.Input[Sequence[pulumi.Input['MediaGraphAssetSinkArgs']]],
                 sources: pulumi.Input[Sequence[pulumi.Input['MediaGraphRtspSourceArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 media_graph_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MediaGraph resource.
        :param pulumi.Input[str] account_name: The Media Services account name.
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the Azure subscription.
        :param pulumi.Input[Sequence[pulumi.Input['MediaGraphAssetSinkArgs']]] sinks: Media Graph sinks.
        :param pulumi.Input[Sequence[pulumi.Input['MediaGraphRtspSourceArgs']]] sources: Media Graph sources.
        :param pulumi.Input[str] description: Media Graph description.
        :param pulumi.Input[str] media_graph_name: The Media Graph name.
        """
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "sinks", sinks)
        pulumi.set(__self__, "sources", sources)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if media_graph_name is not None:
            pulumi.set(__self__, "media_graph_name", media_graph_name)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Input[str]:
        """
        The Media Services account name.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[str]:
        """
        The name of the resource group within the Azure subscription.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def sinks(self) -> pulumi.Input[Sequence[pulumi.Input['MediaGraphAssetSinkArgs']]]:
        """
        Media Graph sinks.
        """
        return pulumi.get(self, "sinks")

    @sinks.setter
    def sinks(self, value: pulumi.Input[Sequence[pulumi.Input['MediaGraphAssetSinkArgs']]]):
        pulumi.set(self, "sinks", value)

    @property
    @pulumi.getter
    def sources(self) -> pulumi.Input[Sequence[pulumi.Input['MediaGraphRtspSourceArgs']]]:
        """
        Media Graph sources.
        """
        return pulumi.get(self, "sources")

    @sources.setter
    def sources(self, value: pulumi.Input[Sequence[pulumi.Input['MediaGraphRtspSourceArgs']]]):
        pulumi.set(self, "sources", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Media Graph description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="mediaGraphName")
    def media_graph_name(self) -> Optional[pulumi.Input[str]]:
        """
        The Media Graph name.
        """
        return pulumi.get(self, "media_graph_name")

    @media_graph_name.setter
    def media_graph_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "media_graph_name", value)


class MediaGraph(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 media_graph_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sinks: Optional[pulumi.Input[Sequence[pulumi.Input[Union['MediaGraphAssetSinkArgs', 'MediaGraphAssetSinkArgsDict']]]]] = None,
                 sources: Optional[pulumi.Input[Sequence[pulumi.Input[Union['MediaGraphRtspSourceArgs', 'MediaGraphRtspSourceArgsDict']]]]] = None,
                 __props__=None):
        """
        The Media Graph.

        Uses Azure REST API version 2020-02-01-preview. In version 1.x of the Azure Native provider, it used API version 2020-02-01-preview.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The Media Services account name.
        :param pulumi.Input[str] description: Media Graph description.
        :param pulumi.Input[str] media_graph_name: The Media Graph name.
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the Azure subscription.
        :param pulumi.Input[Sequence[pulumi.Input[Union['MediaGraphAssetSinkArgs', 'MediaGraphAssetSinkArgsDict']]]] sinks: Media Graph sinks.
        :param pulumi.Input[Sequence[pulumi.Input[Union['MediaGraphRtspSourceArgs', 'MediaGraphRtspSourceArgsDict']]]] sources: Media Graph sources.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MediaGraphArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The Media Graph.

        Uses Azure REST API version 2020-02-01-preview. In version 1.x of the Azure Native provider, it used API version 2020-02-01-preview.

        :param str resource_name: The name of the resource.
        :param MediaGraphArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MediaGraphArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 media_graph_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sinks: Optional[pulumi.Input[Sequence[pulumi.Input[Union['MediaGraphAssetSinkArgs', 'MediaGraphAssetSinkArgsDict']]]]] = None,
                 sources: Optional[pulumi.Input[Sequence[pulumi.Input[Union['MediaGraphRtspSourceArgs', 'MediaGraphRtspSourceArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MediaGraphArgs.__new__(MediaGraphArgs)

            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__.__dict__["account_name"] = account_name
            __props__.__dict__["description"] = description
            __props__.__dict__["media_graph_name"] = media_graph_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if sinks is None and not opts.urn:
                raise TypeError("Missing required property 'sinks'")
            __props__.__dict__["sinks"] = sinks
            if sources is None and not opts.urn:
                raise TypeError("Missing required property 'sources'")
            __props__.__dict__["sources"] = sources
            __props__.__dict__["created"] = None
            __props__.__dict__["last_modified"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:media/v20190901preview:MediaGraph"), pulumi.Alias(type_="azure-native:media/v20200201preview:MediaGraph")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(MediaGraph, __self__).__init__(
            'azure-native:media:MediaGraph',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'MediaGraph':
        """
        Get an existing MediaGraph resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = MediaGraphArgs.__new__(MediaGraphArgs)

        __props__.__dict__["created"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["last_modified"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["sinks"] = None
        __props__.__dict__["sources"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["type"] = None
        return MediaGraph(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[str]:
        """
        Date the Media Graph was created.
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Media Graph description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> pulumi.Output[str]:
        """
        Date the Media Graph was last modified.
        """
        return pulumi.get(self, "last_modified")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def sinks(self) -> pulumi.Output[Sequence['outputs.MediaGraphAssetSinkResponse']]:
        """
        Media Graph sinks.
        """
        return pulumi.get(self, "sinks")

    @property
    @pulumi.getter
    def sources(self) -> pulumi.Output[Sequence['outputs.MediaGraphRtspSourceResponse']]:
        """
        Media Graph sources.
        """
        return pulumi.get(self, "sources")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Media Graph state which indicates the resource allocation status for running the media graph pipeline.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

