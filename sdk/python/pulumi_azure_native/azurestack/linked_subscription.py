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

__all__ = ['LinkedSubscriptionArgs', 'LinkedSubscription']

@pulumi.input_type
class LinkedSubscriptionArgs:
    def __init__(__self__, *,
                 linked_subscription_id: pulumi.Input[builtins.str],
                 registration_resource_id: pulumi.Input[builtins.str],
                 resource_group: pulumi.Input[builtins.str],
                 linked_subscription_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[Union[builtins.str, 'Location']]] = None):
        """
        The set of arguments for constructing a LinkedSubscription resource.
        :param pulumi.Input[builtins.str] linked_subscription_id: The identifier associated with the device subscription.
        :param pulumi.Input[builtins.str] registration_resource_id: The identifier associated with the device registration.
        :param pulumi.Input[builtins.str] resource_group: Name of the resource group.
        :param pulumi.Input[builtins.str] linked_subscription_name: Name of the Linked Subscription resource.
        :param pulumi.Input[Union[builtins.str, 'Location']] location: Location of the resource.
        """
        pulumi.set(__self__, "linked_subscription_id", linked_subscription_id)
        pulumi.set(__self__, "registration_resource_id", registration_resource_id)
        pulumi.set(__self__, "resource_group", resource_group)
        if linked_subscription_name is not None:
            pulumi.set(__self__, "linked_subscription_name", linked_subscription_name)
        if location is not None:
            pulumi.set(__self__, "location", location)

    @property
    @pulumi.getter(name="linkedSubscriptionId")
    def linked_subscription_id(self) -> pulumi.Input[builtins.str]:
        """
        The identifier associated with the device subscription.
        """
        return pulumi.get(self, "linked_subscription_id")

    @linked_subscription_id.setter
    def linked_subscription_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "linked_subscription_id", value)

    @property
    @pulumi.getter(name="registrationResourceId")
    def registration_resource_id(self) -> pulumi.Input[builtins.str]:
        """
        The identifier associated with the device registration.
        """
        return pulumi.get(self, "registration_resource_id")

    @registration_resource_id.setter
    def registration_resource_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "registration_resource_id", value)

    @property
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> pulumi.Input[builtins.str]:
        """
        Name of the resource group.
        """
        return pulumi.get(self, "resource_group")

    @resource_group.setter
    def resource_group(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group", value)

    @property
    @pulumi.getter(name="linkedSubscriptionName")
    def linked_subscription_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the Linked Subscription resource.
        """
        return pulumi.get(self, "linked_subscription_name")

    @linked_subscription_name.setter
    def linked_subscription_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "linked_subscription_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[Union[builtins.str, 'Location']]]:
        """
        Location of the resource.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[Union[builtins.str, 'Location']]]):
        pulumi.set(self, "location", value)


@pulumi.type_token("azure-native:azurestack:LinkedSubscription")
class LinkedSubscription(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 linked_subscription_id: Optional[pulumi.Input[builtins.str]] = None,
                 linked_subscription_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[Union[builtins.str, 'Location']]] = None,
                 registration_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Linked Subscription information.

        Uses Azure REST API version 2020-06-01-preview. In version 2.x of the Azure Native provider, it used API version 2020-06-01-preview.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] linked_subscription_id: The identifier associated with the device subscription.
        :param pulumi.Input[builtins.str] linked_subscription_name: Name of the Linked Subscription resource.
        :param pulumi.Input[Union[builtins.str, 'Location']] location: Location of the resource.
        :param pulumi.Input[builtins.str] registration_resource_id: The identifier associated with the device registration.
        :param pulumi.Input[builtins.str] resource_group: Name of the resource group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LinkedSubscriptionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Linked Subscription information.

        Uses Azure REST API version 2020-06-01-preview. In version 2.x of the Azure Native provider, it used API version 2020-06-01-preview.

        :param str resource_name: The name of the resource.
        :param LinkedSubscriptionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LinkedSubscriptionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 linked_subscription_id: Optional[pulumi.Input[builtins.str]] = None,
                 linked_subscription_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[Union[builtins.str, 'Location']]] = None,
                 registration_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LinkedSubscriptionArgs.__new__(LinkedSubscriptionArgs)

            if linked_subscription_id is None and not opts.urn:
                raise TypeError("Missing required property 'linked_subscription_id'")
            __props__.__dict__["linked_subscription_id"] = linked_subscription_id
            __props__.__dict__["linked_subscription_name"] = linked_subscription_name
            __props__.__dict__["location"] = location
            if registration_resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'registration_resource_id'")
            __props__.__dict__["registration_resource_id"] = registration_resource_id
            if resource_group is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group'")
            __props__.__dict__["resource_group"] = resource_group
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["device_connection_status"] = None
            __props__.__dict__["device_id"] = None
            __props__.__dict__["device_link_state"] = None
            __props__.__dict__["device_object_id"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["last_connected_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["tags"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:azurestack/v20200601preview:LinkedSubscription")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(LinkedSubscription, __self__).__init__(
            'azure-native:azurestack:LinkedSubscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'LinkedSubscription':
        """
        Get an existing LinkedSubscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LinkedSubscriptionArgs.__new__(LinkedSubscriptionArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["device_connection_status"] = None
        __props__.__dict__["device_id"] = None
        __props__.__dict__["device_link_state"] = None
        __props__.__dict__["device_object_id"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["last_connected_time"] = None
        __props__.__dict__["linked_subscription_id"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["registration_resource_id"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return LinkedSubscription(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="deviceConnectionStatus")
    def device_connection_status(self) -> pulumi.Output[builtins.str]:
        """
        The status of the remote management connection of the Azure Stack device.
        """
        return pulumi.get(self, "device_connection_status")

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Output[builtins.str]:
        """
        The identifier of the Azure Stack device for remote management.
        """
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter(name="deviceLinkState")
    def device_link_state(self) -> pulumi.Output[builtins.str]:
        """
        The connection state of the Azure Stack device.
        """
        return pulumi.get(self, "device_link_state")

    @property
    @pulumi.getter(name="deviceObjectId")
    def device_object_id(self) -> pulumi.Output[builtins.str]:
        """
        The object identifier associated with the Azure Stack device connecting to Azure.
        """
        return pulumi.get(self, "device_object_id")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The entity tag used for optimistic concurrency when modifying the resource.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[builtins.str]:
        """
        The kind of the resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="lastConnectedTime")
    def last_connected_time(self) -> pulumi.Output[builtins.str]:
        """
        The last remote management connection time for the Azure Stack device connected to the linked subscription resource.
        """
        return pulumi.get(self, "last_connected_time")

    @property
    @pulumi.getter(name="linkedSubscriptionId")
    def linked_subscription_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The identifier associated with the device subscription.
        """
        return pulumi.get(self, "linked_subscription_id")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        Location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="registrationResourceId")
    def registration_resource_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The identifier associated with the device registration.
        """
        return pulumi.get(self, "registration_resource_id")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Metadata pertaining to creation and last modification of the resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Custom tags for the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Type of Resource.
        """
        return pulumi.get(self, "type")

