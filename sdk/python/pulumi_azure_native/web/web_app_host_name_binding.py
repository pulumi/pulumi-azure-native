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
from ._enums import *

__all__ = ['WebAppHostNameBindingArgs', 'WebAppHostNameBinding']

@pulumi.input_type
class WebAppHostNameBindingArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 azure_resource_name: Optional[pulumi.Input[builtins.str]] = None,
                 azure_resource_type: Optional[pulumi.Input['AzureResourceType']] = None,
                 custom_host_name_dns_record_type: Optional[pulumi.Input['CustomHostNameDnsRecordType']] = None,
                 domain_id: Optional[pulumi.Input[builtins.str]] = None,
                 host_name: Optional[pulumi.Input[builtins.str]] = None,
                 host_name_type: Optional[pulumi.Input['HostNameType']] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 site_name: Optional[pulumi.Input[builtins.str]] = None,
                 ssl_state: Optional[pulumi.Input['SslState']] = None,
                 thumbprint: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a WebAppHostNameBinding resource.
        :param pulumi.Input[builtins.str] name: Name of the app.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[builtins.str] azure_resource_name: Azure resource name.
        :param pulumi.Input['AzureResourceType'] azure_resource_type: Azure resource type.
        :param pulumi.Input['CustomHostNameDnsRecordType'] custom_host_name_dns_record_type: Custom DNS record type.
        :param pulumi.Input[builtins.str] domain_id: Fully qualified ARM domain resource URI.
        :param pulumi.Input[builtins.str] host_name: Hostname in the hostname binding.
        :param pulumi.Input['HostNameType'] host_name_type: Hostname type.
        :param pulumi.Input[builtins.str] kind: Kind of resource.
        :param pulumi.Input[builtins.str] site_name: App Service app name.
        :param pulumi.Input['SslState'] ssl_state: SSL type
        :param pulumi.Input[builtins.str] thumbprint: SSL certificate thumbprint
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if azure_resource_name is not None:
            pulumi.set(__self__, "azure_resource_name", azure_resource_name)
        if azure_resource_type is not None:
            pulumi.set(__self__, "azure_resource_type", azure_resource_type)
        if custom_host_name_dns_record_type is not None:
            pulumi.set(__self__, "custom_host_name_dns_record_type", custom_host_name_dns_record_type)
        if domain_id is not None:
            pulumi.set(__self__, "domain_id", domain_id)
        if host_name is not None:
            pulumi.set(__self__, "host_name", host_name)
        if host_name_type is not None:
            pulumi.set(__self__, "host_name_type", host_name_type)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if site_name is not None:
            pulumi.set(__self__, "site_name", site_name)
        if ssl_state is not None:
            pulumi.set(__self__, "ssl_state", ssl_state)
        if thumbprint is not None:
            pulumi.set(__self__, "thumbprint", thumbprint)

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
    @pulumi.getter(name="azureResourceName")
    def azure_resource_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Azure resource name.
        """
        return pulumi.get(self, "azure_resource_name")

    @azure_resource_name.setter
    def azure_resource_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "azure_resource_name", value)

    @property
    @pulumi.getter(name="azureResourceType")
    def azure_resource_type(self) -> Optional[pulumi.Input['AzureResourceType']]:
        """
        Azure resource type.
        """
        return pulumi.get(self, "azure_resource_type")

    @azure_resource_type.setter
    def azure_resource_type(self, value: Optional[pulumi.Input['AzureResourceType']]):
        pulumi.set(self, "azure_resource_type", value)

    @property
    @pulumi.getter(name="customHostNameDnsRecordType")
    def custom_host_name_dns_record_type(self) -> Optional[pulumi.Input['CustomHostNameDnsRecordType']]:
        """
        Custom DNS record type.
        """
        return pulumi.get(self, "custom_host_name_dns_record_type")

    @custom_host_name_dns_record_type.setter
    def custom_host_name_dns_record_type(self, value: Optional[pulumi.Input['CustomHostNameDnsRecordType']]):
        pulumi.set(self, "custom_host_name_dns_record_type", value)

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Fully qualified ARM domain resource URI.
        """
        return pulumi.get(self, "domain_id")

    @domain_id.setter
    def domain_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "domain_id", value)

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Hostname in the hostname binding.
        """
        return pulumi.get(self, "host_name")

    @host_name.setter
    def host_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "host_name", value)

    @property
    @pulumi.getter(name="hostNameType")
    def host_name_type(self) -> Optional[pulumi.Input['HostNameType']]:
        """
        Hostname type.
        """
        return pulumi.get(self, "host_name_type")

    @host_name_type.setter
    def host_name_type(self, value: Optional[pulumi.Input['HostNameType']]):
        pulumi.set(self, "host_name_type", value)

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
    @pulumi.getter(name="siteName")
    def site_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        App Service app name.
        """
        return pulumi.get(self, "site_name")

    @site_name.setter
    def site_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "site_name", value)

    @property
    @pulumi.getter(name="sslState")
    def ssl_state(self) -> Optional[pulumi.Input['SslState']]:
        """
        SSL type
        """
        return pulumi.get(self, "ssl_state")

    @ssl_state.setter
    def ssl_state(self, value: Optional[pulumi.Input['SslState']]):
        pulumi.set(self, "ssl_state", value)

    @property
    @pulumi.getter
    def thumbprint(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        SSL certificate thumbprint
        """
        return pulumi.get(self, "thumbprint")

    @thumbprint.setter
    def thumbprint(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "thumbprint", value)


@pulumi.type_token("azure-native:web:WebAppHostNameBinding")
class WebAppHostNameBinding(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 azure_resource_name: Optional[pulumi.Input[builtins.str]] = None,
                 azure_resource_type: Optional[pulumi.Input['AzureResourceType']] = None,
                 custom_host_name_dns_record_type: Optional[pulumi.Input['CustomHostNameDnsRecordType']] = None,
                 domain_id: Optional[pulumi.Input[builtins.str]] = None,
                 host_name: Optional[pulumi.Input[builtins.str]] = None,
                 host_name_type: Optional[pulumi.Input['HostNameType']] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 site_name: Optional[pulumi.Input[builtins.str]] = None,
                 ssl_state: Optional[pulumi.Input['SslState']] = None,
                 thumbprint: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        A hostname binding object.

        Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.

        Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] azure_resource_name: Azure resource name.
        :param pulumi.Input['AzureResourceType'] azure_resource_type: Azure resource type.
        :param pulumi.Input['CustomHostNameDnsRecordType'] custom_host_name_dns_record_type: Custom DNS record type.
        :param pulumi.Input[builtins.str] domain_id: Fully qualified ARM domain resource URI.
        :param pulumi.Input[builtins.str] host_name: Hostname in the hostname binding.
        :param pulumi.Input['HostNameType'] host_name_type: Hostname type.
        :param pulumi.Input[builtins.str] kind: Kind of resource.
        :param pulumi.Input[builtins.str] name: Name of the app.
        :param pulumi.Input[builtins.str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[builtins.str] site_name: App Service app name.
        :param pulumi.Input['SslState'] ssl_state: SSL type
        :param pulumi.Input[builtins.str] thumbprint: SSL certificate thumbprint
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebAppHostNameBindingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A hostname binding object.

        Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.

        Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param WebAppHostNameBindingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebAppHostNameBindingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 azure_resource_name: Optional[pulumi.Input[builtins.str]] = None,
                 azure_resource_type: Optional[pulumi.Input['AzureResourceType']] = None,
                 custom_host_name_dns_record_type: Optional[pulumi.Input['CustomHostNameDnsRecordType']] = None,
                 domain_id: Optional[pulumi.Input[builtins.str]] = None,
                 host_name: Optional[pulumi.Input[builtins.str]] = None,
                 host_name_type: Optional[pulumi.Input['HostNameType']] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 site_name: Optional[pulumi.Input[builtins.str]] = None,
                 ssl_state: Optional[pulumi.Input['SslState']] = None,
                 thumbprint: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebAppHostNameBindingArgs.__new__(WebAppHostNameBindingArgs)

            __props__.__dict__["azure_resource_name"] = azure_resource_name
            __props__.__dict__["azure_resource_type"] = azure_resource_type
            __props__.__dict__["custom_host_name_dns_record_type"] = custom_host_name_dns_record_type
            __props__.__dict__["domain_id"] = domain_id
            __props__.__dict__["host_name"] = host_name
            __props__.__dict__["host_name_type"] = host_name_type
            __props__.__dict__["kind"] = kind
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["site_name"] = site_name
            __props__.__dict__["ssl_state"] = ssl_state
            __props__.__dict__["thumbprint"] = thumbprint
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["virtual_ip"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:web/v20150801:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20160801:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20180201:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20181101:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20190801:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20200601:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20200901:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20201001:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20201201:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20210101:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20210115:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20210201:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20210301:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20220301:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20220901:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20230101:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20231201:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20240401:WebAppHostNameBinding"), pulumi.Alias(type_="azure-native:web/v20241101:WebAppHostNameBinding")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(WebAppHostNameBinding, __self__).__init__(
            'azure-native:web:WebAppHostNameBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WebAppHostNameBinding':
        """
        Get an existing WebAppHostNameBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WebAppHostNameBindingArgs.__new__(WebAppHostNameBindingArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["azure_resource_name"] = None
        __props__.__dict__["azure_resource_type"] = None
        __props__.__dict__["custom_host_name_dns_record_type"] = None
        __props__.__dict__["domain_id"] = None
        __props__.__dict__["host_name_type"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["site_name"] = None
        __props__.__dict__["ssl_state"] = None
        __props__.__dict__["thumbprint"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["virtual_ip"] = None
        return WebAppHostNameBinding(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="azureResourceName")
    def azure_resource_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Azure resource name.
        """
        return pulumi.get(self, "azure_resource_name")

    @property
    @pulumi.getter(name="azureResourceType")
    def azure_resource_type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Azure resource type.
        """
        return pulumi.get(self, "azure_resource_type")

    @property
    @pulumi.getter(name="customHostNameDnsRecordType")
    def custom_host_name_dns_record_type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Custom DNS record type.
        """
        return pulumi.get(self, "custom_host_name_dns_record_type")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Fully qualified ARM domain resource URI.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter(name="hostNameType")
    def host_name_type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Hostname type.
        """
        return pulumi.get(self, "host_name_type")

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
    @pulumi.getter(name="siteName")
    def site_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        App Service app name.
        """
        return pulumi.get(self, "site_name")

    @property
    @pulumi.getter(name="sslState")
    def ssl_state(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        SSL type
        """
        return pulumi.get(self, "ssl_state")

    @property
    @pulumi.getter
    def thumbprint(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        SSL certificate thumbprint
        """
        return pulumi.get(self, "thumbprint")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualIP")
    def virtual_ip(self) -> pulumi.Output[builtins.str]:
        """
        Virtual IP address assigned to the hostname if IP based SSL is enabled.
        """
        return pulumi.get(self, "virtual_ip")

