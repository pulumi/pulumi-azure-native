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

__all__ = ['NotificationChannelArgs', 'NotificationChannel']

@pulumi.input_type
class NotificationChannelArgs:
    def __init__(__self__, *,
                 lab_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 email_recipient: Optional[pulumi.Input[builtins.str]] = None,
                 events: Optional[pulumi.Input[Sequence[pulumi.Input['EventArgs']]]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 notification_locale: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 web_hook_url: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a NotificationChannel resource.
        :param pulumi.Input[builtins.str] lab_name: The name of the lab.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] description: Description of notification.
        :param pulumi.Input[builtins.str] email_recipient: The email recipient to send notifications to (can be a list of semi-colon separated email addresses).
        :param pulumi.Input[Sequence[pulumi.Input['EventArgs']]] events: The list of event for which this notification is enabled.
        :param pulumi.Input[builtins.str] location: The location of the resource.
        :param pulumi.Input[builtins.str] name: The name of the NotificationChannel
        :param pulumi.Input[builtins.str] notification_locale: The locale to use when sending a notification (fallback for unsupported languages is EN).
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: The tags of the resource.
        :param pulumi.Input[builtins.str] web_hook_url: The webhook URL to send notifications to.
        """
        pulumi.set(__self__, "lab_name", lab_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if email_recipient is not None:
            pulumi.set(__self__, "email_recipient", email_recipient)
        if events is not None:
            pulumi.set(__self__, "events", events)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notification_locale is not None:
            pulumi.set(__self__, "notification_locale", notification_locale)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if web_hook_url is not None:
            pulumi.set(__self__, "web_hook_url", web_hook_url)

    @property
    @pulumi.getter(name="labName")
    def lab_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the lab.
        """
        return pulumi.get(self, "lab_name")

    @lab_name.setter
    def lab_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "lab_name", value)

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
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description of notification.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="emailRecipient")
    def email_recipient(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The email recipient to send notifications to (can be a list of semi-colon separated email addresses).
        """
        return pulumi.get(self, "email_recipient")

    @email_recipient.setter
    def email_recipient(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "email_recipient", value)

    @property
    @pulumi.getter
    def events(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventArgs']]]]:
        """
        The list of event for which this notification is enabled.
        """
        return pulumi.get(self, "events")

    @events.setter
    def events(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventArgs']]]]):
        pulumi.set(self, "events", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The location of the resource.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the NotificationChannel
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationLocale")
    def notification_locale(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The locale to use when sending a notification (fallback for unsupported languages is EN).
        """
        return pulumi.get(self, "notification_locale")

    @notification_locale.setter
    def notification_locale(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "notification_locale", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="webHookUrl")
    def web_hook_url(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The webhook URL to send notifications to.
        """
        return pulumi.get(self, "web_hook_url")

    @web_hook_url.setter
    def web_hook_url(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "web_hook_url", value)


@pulumi.type_token("azure-native:devtestlab:NotificationChannel")
class NotificationChannel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 email_recipient: Optional[pulumi.Input[builtins.str]] = None,
                 events: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventArgs', 'EventArgsDict']]]]] = None,
                 lab_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 notification_locale: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 web_hook_url: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        A notification.

        Uses Azure REST API version 2018-09-15. In version 2.x of the Azure Native provider, it used API version 2018-09-15.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: Description of notification.
        :param pulumi.Input[builtins.str] email_recipient: The email recipient to send notifications to (can be a list of semi-colon separated email addresses).
        :param pulumi.Input[Sequence[pulumi.Input[Union['EventArgs', 'EventArgsDict']]]] events: The list of event for which this notification is enabled.
        :param pulumi.Input[builtins.str] lab_name: The name of the lab.
        :param pulumi.Input[builtins.str] location: The location of the resource.
        :param pulumi.Input[builtins.str] name: The name of the NotificationChannel
        :param pulumi.Input[builtins.str] notification_locale: The locale to use when sending a notification (fallback for unsupported languages is EN).
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: The tags of the resource.
        :param pulumi.Input[builtins.str] web_hook_url: The webhook URL to send notifications to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NotificationChannelArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A notification.

        Uses Azure REST API version 2018-09-15. In version 2.x of the Azure Native provider, it used API version 2018-09-15.

        :param str resource_name: The name of the resource.
        :param NotificationChannelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NotificationChannelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 email_recipient: Optional[pulumi.Input[builtins.str]] = None,
                 events: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventArgs', 'EventArgsDict']]]]] = None,
                 lab_name: Optional[pulumi.Input[builtins.str]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 notification_locale: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 web_hook_url: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NotificationChannelArgs.__new__(NotificationChannelArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["email_recipient"] = email_recipient
            __props__.__dict__["events"] = events
            if lab_name is None and not opts.urn:
                raise TypeError("Missing required property 'lab_name'")
            __props__.__dict__["lab_name"] = lab_name
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["notification_locale"] = notification_locale
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["web_hook_url"] = web_hook_url
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["created_date"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["unique_identifier"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:devtestlab/v20160515:NotificationChannel"), pulumi.Alias(type_="azure-native:devtestlab/v20180915:NotificationChannel")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(NotificationChannel, __self__).__init__(
            'azure-native:devtestlab:NotificationChannel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NotificationChannel':
        """
        Get an existing NotificationChannel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NotificationChannelArgs.__new__(NotificationChannelArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["created_date"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["email_recipient"] = None
        __props__.__dict__["events"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["notification_locale"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["unique_identifier"] = None
        __props__.__dict__["web_hook_url"] = None
        return NotificationChannel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="createdDate")
    def created_date(self) -> pulumi.Output[builtins.str]:
        """
        The creation date of the notification channel.
        """
        return pulumi.get(self, "created_date")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Description of notification.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="emailRecipient")
    def email_recipient(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The email recipient to send notifications to (can be a list of semi-colon separated email addresses).
        """
        return pulumi.get(self, "email_recipient")

    @property
    @pulumi.getter
    def events(self) -> pulumi.Output[Optional[Sequence['outputs.EventResponse']]]:
        """
        The list of event for which this notification is enabled.
        """
        return pulumi.get(self, "events")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationLocale")
    def notification_locale(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The locale to use when sending a notification (fallback for unsupported languages is EN).
        """
        return pulumi.get(self, "notification_locale")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        The provisioning status of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="uniqueIdentifier")
    def unique_identifier(self) -> pulumi.Output[builtins.str]:
        """
        The unique immutable identifier of a resource (Guid).
        """
        return pulumi.get(self, "unique_identifier")

    @property
    @pulumi.getter(name="webHookUrl")
    def web_hook_url(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The webhook URL to send notifications to.
        """
        return pulumi.get(self, "web_hook_url")

