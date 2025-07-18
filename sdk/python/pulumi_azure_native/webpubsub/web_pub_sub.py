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

__all__ = ['WebPubSubArgs', 'WebPubSub']

@pulumi.input_type
class WebPubSubArgs:
    def __init__(__self__, *,
                 resource_group_name: pulumi.Input[builtins.str],
                 disable_aad_auth: Optional[pulumi.Input[builtins.bool]] = None,
                 disable_local_auth: Optional[pulumi.Input[builtins.bool]] = None,
                 identity: Optional[pulumi.Input['ManagedIdentityArgs']] = None,
                 kind: Optional[pulumi.Input[Union[builtins.str, 'ServiceKind']]] = None,
                 live_trace_configuration: Optional[pulumi.Input['LiveTraceConfigurationArgs']] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 network_acls: Optional[pulumi.Input['WebPubSubNetworkACLsArgs']] = None,
                 public_network_access: Optional[pulumi.Input[builtins.str]] = None,
                 region_endpoint_enabled: Optional[pulumi.Input[builtins.str]] = None,
                 resource_log_configuration: Optional[pulumi.Input['ResourceLogConfigurationArgs']] = None,
                 resource_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_stopped: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input['ResourceSkuArgs']] = None,
                 socket_io: Optional[pulumi.Input['WebPubSubSocketIOSettingsArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 tls: Optional[pulumi.Input['WebPubSubTlsSettingsArgs']] = None):
        """
        The set of arguments for constructing a WebPubSub resource.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.bool] disable_aad_auth: DisableLocalAuth
               Enable or disable aad auth
               When set as true, connection with AuthType=aad won't work.
        :param pulumi.Input[builtins.bool] disable_local_auth: DisableLocalAuth
               Enable or disable local auth with AccessKey
               When set as true, connection with AccessKey=xxx won't work.
        :param pulumi.Input['ManagedIdentityArgs'] identity: A class represent managed identities used for request and response
        :param pulumi.Input[Union[builtins.str, 'ServiceKind']] kind: The kind of the service
        :param pulumi.Input['LiveTraceConfigurationArgs'] live_trace_configuration: Live trace configuration of a Microsoft.SignalRService resource.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input['WebPubSubNetworkACLsArgs'] network_acls: Network ACLs for the resource
        :param pulumi.Input[builtins.str] public_network_access: Enable or disable public network access. Default to "Enabled".
               When it's Enabled, network ACLs still apply.
               When it's Disabled, public network access is always disabled no matter what you set in network ACLs.
        :param pulumi.Input[builtins.str] region_endpoint_enabled: Enable or disable the regional endpoint. Default to "Enabled".
               When it's Disabled, new connections will not be routed to this endpoint, however existing connections will not be affected.
               This property is replica specific. Disable the regional endpoint without replica is not allowed.
        :param pulumi.Input['ResourceLogConfigurationArgs'] resource_log_configuration: Resource log configuration of a Microsoft.SignalRService resource.
        :param pulumi.Input[builtins.str] resource_name: The name of the resource.
        :param pulumi.Input[builtins.str] resource_stopped: Stop or start the resource.  Default to "False".
               When it's true, the data plane of the resource is shutdown.
               When it's false, the data plane of the resource is started.
        :param pulumi.Input['ResourceSkuArgs'] sku: The billing information of the resource.
        :param pulumi.Input['WebPubSubSocketIOSettingsArgs'] socket_io: SocketIO settings for the resource
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input['WebPubSubTlsSettingsArgs'] tls: TLS settings for the resource
        """
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if disable_aad_auth is None:
            disable_aad_auth = False
        if disable_aad_auth is not None:
            pulumi.set(__self__, "disable_aad_auth", disable_aad_auth)
        if disable_local_auth is None:
            disable_local_auth = False
        if disable_local_auth is not None:
            pulumi.set(__self__, "disable_local_auth", disable_local_auth)
        if identity is not None:
            pulumi.set(__self__, "identity", identity)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if live_trace_configuration is not None:
            pulumi.set(__self__, "live_trace_configuration", live_trace_configuration)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if network_acls is not None:
            pulumi.set(__self__, "network_acls", network_acls)
        if public_network_access is None:
            public_network_access = 'Enabled'
        if public_network_access is not None:
            pulumi.set(__self__, "public_network_access", public_network_access)
        if region_endpoint_enabled is None:
            region_endpoint_enabled = 'Enabled'
        if region_endpoint_enabled is not None:
            pulumi.set(__self__, "region_endpoint_enabled", region_endpoint_enabled)
        if resource_log_configuration is not None:
            pulumi.set(__self__, "resource_log_configuration", resource_log_configuration)
        if resource_name is not None:
            pulumi.set(__self__, "resource_name", resource_name)
        if resource_stopped is None:
            resource_stopped = 'false'
        if resource_stopped is not None:
            pulumi.set(__self__, "resource_stopped", resource_stopped)
        if sku is not None:
            pulumi.set(__self__, "sku", sku)
        if socket_io is not None:
            pulumi.set(__self__, "socket_io", socket_io)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tls is not None:
            pulumi.set(__self__, "tls", tls)

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
    @pulumi.getter(name="disableAadAuth")
    def disable_aad_auth(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        DisableLocalAuth
        Enable or disable aad auth
        When set as true, connection with AuthType=aad won't work.
        """
        return pulumi.get(self, "disable_aad_auth")

    @disable_aad_auth.setter
    def disable_aad_auth(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "disable_aad_auth", value)

    @property
    @pulumi.getter(name="disableLocalAuth")
    def disable_local_auth(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        DisableLocalAuth
        Enable or disable local auth with AccessKey
        When set as true, connection with AccessKey=xxx won't work.
        """
        return pulumi.get(self, "disable_local_auth")

    @disable_local_auth.setter
    def disable_local_auth(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "disable_local_auth", value)

    @property
    @pulumi.getter
    def identity(self) -> Optional[pulumi.Input['ManagedIdentityArgs']]:
        """
        A class represent managed identities used for request and response
        """
        return pulumi.get(self, "identity")

    @identity.setter
    def identity(self, value: Optional[pulumi.Input['ManagedIdentityArgs']]):
        pulumi.set(self, "identity", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[Union[builtins.str, 'ServiceKind']]]:
        """
        The kind of the service
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[Union[builtins.str, 'ServiceKind']]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="liveTraceConfiguration")
    def live_trace_configuration(self) -> Optional[pulumi.Input['LiveTraceConfigurationArgs']]:
        """
        Live trace configuration of a Microsoft.SignalRService resource.
        """
        return pulumi.get(self, "live_trace_configuration")

    @live_trace_configuration.setter
    def live_trace_configuration(self, value: Optional[pulumi.Input['LiveTraceConfigurationArgs']]):
        pulumi.set(self, "live_trace_configuration", value)

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
    @pulumi.getter(name="networkACLs")
    def network_acls(self) -> Optional[pulumi.Input['WebPubSubNetworkACLsArgs']]:
        """
        Network ACLs for the resource
        """
        return pulumi.get(self, "network_acls")

    @network_acls.setter
    def network_acls(self, value: Optional[pulumi.Input['WebPubSubNetworkACLsArgs']]):
        pulumi.set(self, "network_acls", value)

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Enable or disable public network access. Default to "Enabled".
        When it's Enabled, network ACLs still apply.
        When it's Disabled, public network access is always disabled no matter what you set in network ACLs.
        """
        return pulumi.get(self, "public_network_access")

    @public_network_access.setter
    def public_network_access(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "public_network_access", value)

    @property
    @pulumi.getter(name="regionEndpointEnabled")
    def region_endpoint_enabled(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Enable or disable the regional endpoint. Default to "Enabled".
        When it's Disabled, new connections will not be routed to this endpoint, however existing connections will not be affected.
        This property is replica specific. Disable the regional endpoint without replica is not allowed.
        """
        return pulumi.get(self, "region_endpoint_enabled")

    @region_endpoint_enabled.setter
    def region_endpoint_enabled(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region_endpoint_enabled", value)

    @property
    @pulumi.getter(name="resourceLogConfiguration")
    def resource_log_configuration(self) -> Optional[pulumi.Input['ResourceLogConfigurationArgs']]:
        """
        Resource log configuration of a Microsoft.SignalRService resource.
        """
        return pulumi.get(self, "resource_log_configuration")

    @resource_log_configuration.setter
    def resource_log_configuration(self, value: Optional[pulumi.Input['ResourceLogConfigurationArgs']]):
        pulumi.set(self, "resource_log_configuration", value)

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "resource_name")

    @resource_name.setter
    def resource_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "resource_name", value)

    @property
    @pulumi.getter(name="resourceStopped")
    def resource_stopped(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Stop or start the resource.  Default to "False".
        When it's true, the data plane of the resource is shutdown.
        When it's false, the data plane of the resource is started.
        """
        return pulumi.get(self, "resource_stopped")

    @resource_stopped.setter
    def resource_stopped(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "resource_stopped", value)

    @property
    @pulumi.getter
    def sku(self) -> Optional[pulumi.Input['ResourceSkuArgs']]:
        """
        The billing information of the resource.
        """
        return pulumi.get(self, "sku")

    @sku.setter
    def sku(self, value: Optional[pulumi.Input['ResourceSkuArgs']]):
        pulumi.set(self, "sku", value)

    @property
    @pulumi.getter(name="socketIO")
    def socket_io(self) -> Optional[pulumi.Input['WebPubSubSocketIOSettingsArgs']]:
        """
        SocketIO settings for the resource
        """
        return pulumi.get(self, "socket_io")

    @socket_io.setter
    def socket_io(self, value: Optional[pulumi.Input['WebPubSubSocketIOSettingsArgs']]):
        pulumi.set(self, "socket_io", value)

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
    @pulumi.getter
    def tls(self) -> Optional[pulumi.Input['WebPubSubTlsSettingsArgs']]:
        """
        TLS settings for the resource
        """
        return pulumi.get(self, "tls")

    @tls.setter
    def tls(self, value: Optional[pulumi.Input['WebPubSubTlsSettingsArgs']]):
        pulumi.set(self, "tls", value)


@pulumi.type_token("azure-native:webpubsub:WebPubSub")
class WebPubSub(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_aad_auth: Optional[pulumi.Input[builtins.bool]] = None,
                 disable_local_auth: Optional[pulumi.Input[builtins.bool]] = None,
                 identity: Optional[pulumi.Input[Union['ManagedIdentityArgs', 'ManagedIdentityArgsDict']]] = None,
                 kind: Optional[pulumi.Input[Union[builtins.str, 'ServiceKind']]] = None,
                 live_trace_configuration: Optional[pulumi.Input[Union['LiveTraceConfigurationArgs', 'LiveTraceConfigurationArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 network_acls: Optional[pulumi.Input[Union['WebPubSubNetworkACLsArgs', 'WebPubSubNetworkACLsArgsDict']]] = None,
                 public_network_access: Optional[pulumi.Input[builtins.str]] = None,
                 region_endpoint_enabled: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_log_configuration: Optional[pulumi.Input[Union['ResourceLogConfigurationArgs', 'ResourceLogConfigurationArgsDict']]] = None,
                 resource_name_: Optional[pulumi.Input[builtins.str]] = None,
                 resource_stopped: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input[Union['ResourceSkuArgs', 'ResourceSkuArgsDict']]] = None,
                 socket_io: Optional[pulumi.Input[Union['WebPubSubSocketIOSettingsArgs', 'WebPubSubSocketIOSettingsArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 tls: Optional[pulumi.Input[Union['WebPubSubTlsSettingsArgs', 'WebPubSubTlsSettingsArgsDict']]] = None,
                 __props__=None):
        """
        A class represent a resource.

        Uses Azure REST API version 2024-03-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.

        Other available API versions: 2023-02-01, 2023-03-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native webpubsub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] disable_aad_auth: DisableLocalAuth
               Enable or disable aad auth
               When set as true, connection with AuthType=aad won't work.
        :param pulumi.Input[builtins.bool] disable_local_auth: DisableLocalAuth
               Enable or disable local auth with AccessKey
               When set as true, connection with AccessKey=xxx won't work.
        :param pulumi.Input[Union['ManagedIdentityArgs', 'ManagedIdentityArgsDict']] identity: A class represent managed identities used for request and response
        :param pulumi.Input[Union[builtins.str, 'ServiceKind']] kind: The kind of the service
        :param pulumi.Input[Union['LiveTraceConfigurationArgs', 'LiveTraceConfigurationArgsDict']] live_trace_configuration: Live trace configuration of a Microsoft.SignalRService resource.
        :param pulumi.Input[builtins.str] location: The geo-location where the resource lives
        :param pulumi.Input[Union['WebPubSubNetworkACLsArgs', 'WebPubSubNetworkACLsArgsDict']] network_acls: Network ACLs for the resource
        :param pulumi.Input[builtins.str] public_network_access: Enable or disable public network access. Default to "Enabled".
               When it's Enabled, network ACLs still apply.
               When it's Disabled, public network access is always disabled no matter what you set in network ACLs.
        :param pulumi.Input[builtins.str] region_endpoint_enabled: Enable or disable the regional endpoint. Default to "Enabled".
               When it's Disabled, new connections will not be routed to this endpoint, however existing connections will not be affected.
               This property is replica specific. Disable the regional endpoint without replica is not allowed.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Union['ResourceLogConfigurationArgs', 'ResourceLogConfigurationArgsDict']] resource_log_configuration: Resource log configuration of a Microsoft.SignalRService resource.
        :param pulumi.Input[builtins.str] resource_name_: The name of the resource.
        :param pulumi.Input[builtins.str] resource_stopped: Stop or start the resource.  Default to "False".
               When it's true, the data plane of the resource is shutdown.
               When it's false, the data plane of the resource is started.
        :param pulumi.Input[Union['ResourceSkuArgs', 'ResourceSkuArgsDict']] sku: The billing information of the resource.
        :param pulumi.Input[Union['WebPubSubSocketIOSettingsArgs', 'WebPubSubSocketIOSettingsArgsDict']] socket_io: SocketIO settings for the resource
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Resource tags.
        :param pulumi.Input[Union['WebPubSubTlsSettingsArgs', 'WebPubSubTlsSettingsArgsDict']] tls: TLS settings for the resource
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebPubSubArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A class represent a resource.

        Uses Azure REST API version 2024-03-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.

        Other available API versions: 2023-02-01, 2023-03-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native webpubsub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param WebPubSubArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebPubSubArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_aad_auth: Optional[pulumi.Input[builtins.bool]] = None,
                 disable_local_auth: Optional[pulumi.Input[builtins.bool]] = None,
                 identity: Optional[pulumi.Input[Union['ManagedIdentityArgs', 'ManagedIdentityArgsDict']]] = None,
                 kind: Optional[pulumi.Input[Union[builtins.str, 'ServiceKind']]] = None,
                 live_trace_configuration: Optional[pulumi.Input[Union['LiveTraceConfigurationArgs', 'LiveTraceConfigurationArgsDict']]] = None,
                 location: Optional[pulumi.Input[builtins.str]] = None,
                 network_acls: Optional[pulumi.Input[Union['WebPubSubNetworkACLsArgs', 'WebPubSubNetworkACLsArgsDict']]] = None,
                 public_network_access: Optional[pulumi.Input[builtins.str]] = None,
                 region_endpoint_enabled: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_log_configuration: Optional[pulumi.Input[Union['ResourceLogConfigurationArgs', 'ResourceLogConfigurationArgsDict']]] = None,
                 resource_name_: Optional[pulumi.Input[builtins.str]] = None,
                 resource_stopped: Optional[pulumi.Input[builtins.str]] = None,
                 sku: Optional[pulumi.Input[Union['ResourceSkuArgs', 'ResourceSkuArgsDict']]] = None,
                 socket_io: Optional[pulumi.Input[Union['WebPubSubSocketIOSettingsArgs', 'WebPubSubSocketIOSettingsArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 tls: Optional[pulumi.Input[Union['WebPubSubTlsSettingsArgs', 'WebPubSubTlsSettingsArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebPubSubArgs.__new__(WebPubSubArgs)

            if disable_aad_auth is None:
                disable_aad_auth = False
            __props__.__dict__["disable_aad_auth"] = disable_aad_auth
            if disable_local_auth is None:
                disable_local_auth = False
            __props__.__dict__["disable_local_auth"] = disable_local_auth
            __props__.__dict__["identity"] = identity
            __props__.__dict__["kind"] = kind
            __props__.__dict__["live_trace_configuration"] = live_trace_configuration
            __props__.__dict__["location"] = location
            __props__.__dict__["network_acls"] = network_acls
            if public_network_access is None:
                public_network_access = 'Enabled'
            __props__.__dict__["public_network_access"] = public_network_access
            if region_endpoint_enabled is None:
                region_endpoint_enabled = 'Enabled'
            __props__.__dict__["region_endpoint_enabled"] = region_endpoint_enabled
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["resource_log_configuration"] = resource_log_configuration
            __props__.__dict__["resource_name"] = resource_name_
            if resource_stopped is None:
                resource_stopped = 'false'
            __props__.__dict__["resource_stopped"] = resource_stopped
            __props__.__dict__["sku"] = sku
            __props__.__dict__["socket_io"] = socket_io
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tls"] = tls
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["external_ip"] = None
            __props__.__dict__["host_name"] = None
            __props__.__dict__["host_name_prefix"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["private_endpoint_connections"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["public_port"] = None
            __props__.__dict__["server_port"] = None
            __props__.__dict__["shared_private_link_resources"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["version"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:webpubsub/v20210401preview:WebPubSub"), pulumi.Alias(type_="azure-native:webpubsub/v20210601preview:WebPubSub"), pulumi.Alias(type_="azure-native:webpubsub/v20210901preview:WebPubSub"), pulumi.Alias(type_="azure-native:webpubsub/v20211001:WebPubSub"), pulumi.Alias(type_="azure-native:webpubsub/v20220801preview:WebPubSub"), pulumi.Alias(type_="azure-native:webpubsub/v20230201:WebPubSub"), pulumi.Alias(type_="azure-native:webpubsub/v20230301preview:WebPubSub"), pulumi.Alias(type_="azure-native:webpubsub/v20230601preview:WebPubSub"), pulumi.Alias(type_="azure-native:webpubsub/v20230801preview:WebPubSub"), pulumi.Alias(type_="azure-native:webpubsub/v20240101preview:WebPubSub"), pulumi.Alias(type_="azure-native:webpubsub/v20240301:WebPubSub"), pulumi.Alias(type_="azure-native:webpubsub/v20240401preview:WebPubSub"), pulumi.Alias(type_="azure-native:webpubsub/v20240801preview:WebPubSub"), pulumi.Alias(type_="azure-native:webpubsub/v20241001preview:WebPubSub"), pulumi.Alias(type_="azure-native:webpubsub/v20250101preview:WebPubSub")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(WebPubSub, __self__).__init__(
            'azure-native:webpubsub:WebPubSub',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WebPubSub':
        """
        Get an existing WebPubSub resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WebPubSubArgs.__new__(WebPubSubArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["disable_aad_auth"] = None
        __props__.__dict__["disable_local_auth"] = None
        __props__.__dict__["external_ip"] = None
        __props__.__dict__["host_name"] = None
        __props__.__dict__["host_name_prefix"] = None
        __props__.__dict__["identity"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["live_trace_configuration"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["network_acls"] = None
        __props__.__dict__["private_endpoint_connections"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["public_network_access"] = None
        __props__.__dict__["public_port"] = None
        __props__.__dict__["region_endpoint_enabled"] = None
        __props__.__dict__["resource_log_configuration"] = None
        __props__.__dict__["resource_stopped"] = None
        __props__.__dict__["server_port"] = None
        __props__.__dict__["shared_private_link_resources"] = None
        __props__.__dict__["sku"] = None
        __props__.__dict__["socket_io"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["tls"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["version"] = None
        return WebPubSub(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="disableAadAuth")
    def disable_aad_auth(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        DisableLocalAuth
        Enable or disable aad auth
        When set as true, connection with AuthType=aad won't work.
        """
        return pulumi.get(self, "disable_aad_auth")

    @property
    @pulumi.getter(name="disableLocalAuth")
    def disable_local_auth(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        DisableLocalAuth
        Enable or disable local auth with AccessKey
        When set as true, connection with AccessKey=xxx won't work.
        """
        return pulumi.get(self, "disable_local_auth")

    @property
    @pulumi.getter(name="externalIP")
    def external_ip(self) -> pulumi.Output[builtins.str]:
        """
        The publicly accessible IP of the resource.
        """
        return pulumi.get(self, "external_ip")

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> pulumi.Output[builtins.str]:
        """
        FQDN of the service instance.
        """
        return pulumi.get(self, "host_name")

    @property
    @pulumi.getter(name="hostNamePrefix")
    def host_name_prefix(self) -> pulumi.Output[builtins.str]:
        """
        Deprecated.
        """
        return pulumi.get(self, "host_name_prefix")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.ManagedIdentityResponse']]:
        """
        A class represent managed identities used for request and response
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The kind of the service
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="liveTraceConfiguration")
    def live_trace_configuration(self) -> pulumi.Output[Optional['outputs.LiveTraceConfigurationResponse']]:
        """
        Live trace configuration of a Microsoft.SignalRService resource.
        """
        return pulumi.get(self, "live_trace_configuration")

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
    @pulumi.getter(name="networkACLs")
    def network_acls(self) -> pulumi.Output[Optional['outputs.WebPubSubNetworkACLsResponse']]:
        """
        Network ACLs for the resource
        """
        return pulumi.get(self, "network_acls")

    @property
    @pulumi.getter(name="privateEndpointConnections")
    def private_endpoint_connections(self) -> pulumi.Output[Sequence['outputs.PrivateEndpointConnectionResponse']]:
        """
        Private endpoint connections to the resource.
        """
        return pulumi.get(self, "private_endpoint_connections")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Enable or disable public network access. Default to "Enabled".
        When it's Enabled, network ACLs still apply.
        When it's Disabled, public network access is always disabled no matter what you set in network ACLs.
        """
        return pulumi.get(self, "public_network_access")

    @property
    @pulumi.getter(name="publicPort")
    def public_port(self) -> pulumi.Output[builtins.int]:
        """
        The publicly accessible port of the resource which is designed for browser/client side usage.
        """
        return pulumi.get(self, "public_port")

    @property
    @pulumi.getter(name="regionEndpointEnabled")
    def region_endpoint_enabled(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Enable or disable the regional endpoint. Default to "Enabled".
        When it's Disabled, new connections will not be routed to this endpoint, however existing connections will not be affected.
        This property is replica specific. Disable the regional endpoint without replica is not allowed.
        """
        return pulumi.get(self, "region_endpoint_enabled")

    @property
    @pulumi.getter(name="resourceLogConfiguration")
    def resource_log_configuration(self) -> pulumi.Output[Optional['outputs.ResourceLogConfigurationResponse']]:
        """
        Resource log configuration of a Microsoft.SignalRService resource.
        """
        return pulumi.get(self, "resource_log_configuration")

    @property
    @pulumi.getter(name="resourceStopped")
    def resource_stopped(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Stop or start the resource.  Default to "False".
        When it's true, the data plane of the resource is shutdown.
        When it's false, the data plane of the resource is started.
        """
        return pulumi.get(self, "resource_stopped")

    @property
    @pulumi.getter(name="serverPort")
    def server_port(self) -> pulumi.Output[builtins.int]:
        """
        The publicly accessible port of the resource which is designed for customer server side usage.
        """
        return pulumi.get(self, "server_port")

    @property
    @pulumi.getter(name="sharedPrivateLinkResources")
    def shared_private_link_resources(self) -> pulumi.Output[Sequence['outputs.SharedPrivateLinkResourceResponse']]:
        """
        The list of shared private link resources.
        """
        return pulumi.get(self, "shared_private_link_resources")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.ResourceSkuResponse']]:
        """
        The billing information of the resource.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="socketIO")
    def socket_io(self) -> pulumi.Output[Optional['outputs.WebPubSubSocketIOSettingsResponse']]:
        """
        SocketIO settings for the resource
        """
        return pulumi.get(self, "socket_io")

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
    def tls(self) -> pulumi.Output[Optional['outputs.WebPubSubTlsSettingsResponse']]:
        """
        TLS settings for the resource
        """
        return pulumi.get(self, "tls")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[builtins.str]:
        """
        Version of the resource. Probably you need the same or higher version of client SDKs.
        """
        return pulumi.get(self, "version")

