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

__all__ = [
    'GetEndpointResult',
    'AwaitableGetEndpointResult',
    'get_endpoint',
    'get_endpoint_output',
]

@pulumi.output_type
class GetEndpointResult:
    """
    CDN endpoint is the entity within a CDN profile containing configuration information such as origin, protocol, content caching and delivery behavior. The CDN endpoint uses the URL format <endpointname>.azureedge.net.
    """
    def __init__(__self__, azure_api_version=None, content_types_to_compress=None, custom_domains=None, default_origin_group=None, delivery_policy=None, geo_filters=None, host_name=None, id=None, is_compression_enabled=None, is_http_allowed=None, is_https_allowed=None, location=None, name=None, optimization_type=None, origin_groups=None, origin_host_header=None, origin_path=None, origins=None, probe_path=None, provisioning_state=None, query_string_caching_behavior=None, resource_state=None, system_data=None, tags=None, type=None, url_signing_keys=None, web_application_firewall_policy_link=None):
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if content_types_to_compress and not isinstance(content_types_to_compress, list):
            raise TypeError("Expected argument 'content_types_to_compress' to be a list")
        pulumi.set(__self__, "content_types_to_compress", content_types_to_compress)
        if custom_domains and not isinstance(custom_domains, list):
            raise TypeError("Expected argument 'custom_domains' to be a list")
        pulumi.set(__self__, "custom_domains", custom_domains)
        if default_origin_group and not isinstance(default_origin_group, dict):
            raise TypeError("Expected argument 'default_origin_group' to be a dict")
        pulumi.set(__self__, "default_origin_group", default_origin_group)
        if delivery_policy and not isinstance(delivery_policy, dict):
            raise TypeError("Expected argument 'delivery_policy' to be a dict")
        pulumi.set(__self__, "delivery_policy", delivery_policy)
        if geo_filters and not isinstance(geo_filters, list):
            raise TypeError("Expected argument 'geo_filters' to be a list")
        pulumi.set(__self__, "geo_filters", geo_filters)
        if host_name and not isinstance(host_name, str):
            raise TypeError("Expected argument 'host_name' to be a str")
        pulumi.set(__self__, "host_name", host_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_compression_enabled and not isinstance(is_compression_enabled, bool):
            raise TypeError("Expected argument 'is_compression_enabled' to be a bool")
        pulumi.set(__self__, "is_compression_enabled", is_compression_enabled)
        if is_http_allowed and not isinstance(is_http_allowed, bool):
            raise TypeError("Expected argument 'is_http_allowed' to be a bool")
        pulumi.set(__self__, "is_http_allowed", is_http_allowed)
        if is_https_allowed and not isinstance(is_https_allowed, bool):
            raise TypeError("Expected argument 'is_https_allowed' to be a bool")
        pulumi.set(__self__, "is_https_allowed", is_https_allowed)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if optimization_type and not isinstance(optimization_type, str):
            raise TypeError("Expected argument 'optimization_type' to be a str")
        pulumi.set(__self__, "optimization_type", optimization_type)
        if origin_groups and not isinstance(origin_groups, list):
            raise TypeError("Expected argument 'origin_groups' to be a list")
        pulumi.set(__self__, "origin_groups", origin_groups)
        if origin_host_header and not isinstance(origin_host_header, str):
            raise TypeError("Expected argument 'origin_host_header' to be a str")
        pulumi.set(__self__, "origin_host_header", origin_host_header)
        if origin_path and not isinstance(origin_path, str):
            raise TypeError("Expected argument 'origin_path' to be a str")
        pulumi.set(__self__, "origin_path", origin_path)
        if origins and not isinstance(origins, list):
            raise TypeError("Expected argument 'origins' to be a list")
        pulumi.set(__self__, "origins", origins)
        if probe_path and not isinstance(probe_path, str):
            raise TypeError("Expected argument 'probe_path' to be a str")
        pulumi.set(__self__, "probe_path", probe_path)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if query_string_caching_behavior and not isinstance(query_string_caching_behavior, str):
            raise TypeError("Expected argument 'query_string_caching_behavior' to be a str")
        pulumi.set(__self__, "query_string_caching_behavior", query_string_caching_behavior)
        if resource_state and not isinstance(resource_state, str):
            raise TypeError("Expected argument 'resource_state' to be a str")
        pulumi.set(__self__, "resource_state", resource_state)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if url_signing_keys and not isinstance(url_signing_keys, list):
            raise TypeError("Expected argument 'url_signing_keys' to be a list")
        pulumi.set(__self__, "url_signing_keys", url_signing_keys)
        if web_application_firewall_policy_link and not isinstance(web_application_firewall_policy_link, dict):
            raise TypeError("Expected argument 'web_application_firewall_policy_link' to be a dict")
        pulumi.set(__self__, "web_application_firewall_policy_link", web_application_firewall_policy_link)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="contentTypesToCompress")
    def content_types_to_compress(self) -> Optional[Sequence[builtins.str]]:
        """
        List of content types on which compression applies. The value should be a valid MIME type.
        """
        return pulumi.get(self, "content_types_to_compress")

    @property
    @pulumi.getter(name="customDomains")
    def custom_domains(self) -> Sequence['outputs.DeepCreatedCustomDomainResponse']:
        """
        The custom domains under the endpoint.
        """
        return pulumi.get(self, "custom_domains")

    @property
    @pulumi.getter(name="defaultOriginGroup")
    def default_origin_group(self) -> Optional['outputs.ResourceReferenceResponse']:
        """
        A reference to the origin group.
        """
        return pulumi.get(self, "default_origin_group")

    @property
    @pulumi.getter(name="deliveryPolicy")
    def delivery_policy(self) -> Optional['outputs.EndpointPropertiesUpdateParametersResponseDeliveryPolicy']:
        """
        A policy that specifies the delivery rules to be used for an endpoint.
        """
        return pulumi.get(self, "delivery_policy")

    @property
    @pulumi.getter(name="geoFilters")
    def geo_filters(self) -> Optional[Sequence['outputs.GeoFilterResponse']]:
        """
        List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
        """
        return pulumi.get(self, "geo_filters")

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> builtins.str:
        """
        The host name of the endpoint structured as {endpointName}.{DNSZone}, e.g. contoso.azureedge.net
        """
        return pulumi.get(self, "host_name")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isCompressionEnabled")
    def is_compression_enabled(self) -> Optional[builtins.bool]:
        """
        Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
        """
        return pulumi.get(self, "is_compression_enabled")

    @property
    @pulumi.getter(name="isHttpAllowed")
    def is_http_allowed(self) -> Optional[builtins.bool]:
        """
        Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
        """
        return pulumi.get(self, "is_http_allowed")

    @property
    @pulumi.getter(name="isHttpsAllowed")
    def is_https_allowed(self) -> Optional[builtins.bool]:
        """
        Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
        """
        return pulumi.get(self, "is_https_allowed")

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="optimizationType")
    def optimization_type(self) -> Optional[builtins.str]:
        """
        Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
        """
        return pulumi.get(self, "optimization_type")

    @property
    @pulumi.getter(name="originGroups")
    def origin_groups(self) -> Optional[Sequence['outputs.DeepCreatedOriginGroupResponse']]:
        """
        The origin groups comprising of origins that are used for load balancing the traffic based on availability.
        """
        return pulumi.get(self, "origin_groups")

    @property
    @pulumi.getter(name="originHostHeader")
    def origin_host_header(self) -> Optional[builtins.str]:
        """
        The host header value sent to the origin with each request. This property at Endpoint is only allowed when endpoint uses single origin and can be overridden by the same property specified at origin.If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
        """
        return pulumi.get(self, "origin_host_header")

    @property
    @pulumi.getter(name="originPath")
    def origin_path(self) -> Optional[builtins.str]:
        """
        A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
        """
        return pulumi.get(self, "origin_path")

    @property
    @pulumi.getter
    def origins(self) -> Sequence['outputs.DeepCreatedOriginResponse']:
        """
        The source of the content being delivered via CDN.
        """
        return pulumi.get(self, "origins")

    @property
    @pulumi.getter(name="probePath")
    def probe_path(self) -> Optional[builtins.str]:
        """
        Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path. This property is only relevant when using a single origin.
        """
        return pulumi.get(self, "probe_path")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Provisioning status of the endpoint.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="queryStringCachingBehavior")
    def query_string_caching_behavior(self) -> Optional[builtins.str]:
        """
        Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
        """
        return pulumi.get(self, "query_string_caching_behavior")

    @property
    @pulumi.getter(name="resourceState")
    def resource_state(self) -> builtins.str:
        """
        Resource status of the endpoint.
        """
        return pulumi.get(self, "resource_state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Read only system data
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="urlSigningKeys")
    def url_signing_keys(self) -> Optional[Sequence['outputs.UrlSigningKeyResponse']]:
        """
        List of keys used to validate the signed URL hashes.
        """
        return pulumi.get(self, "url_signing_keys")

    @property
    @pulumi.getter(name="webApplicationFirewallPolicyLink")
    def web_application_firewall_policy_link(self) -> Optional['outputs.EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink']:
        """
        Defines the Web Application Firewall policy for the endpoint (if applicable)
        """
        return pulumi.get(self, "web_application_firewall_policy_link")


class AwaitableGetEndpointResult(GetEndpointResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEndpointResult(
            azure_api_version=self.azure_api_version,
            content_types_to_compress=self.content_types_to_compress,
            custom_domains=self.custom_domains,
            default_origin_group=self.default_origin_group,
            delivery_policy=self.delivery_policy,
            geo_filters=self.geo_filters,
            host_name=self.host_name,
            id=self.id,
            is_compression_enabled=self.is_compression_enabled,
            is_http_allowed=self.is_http_allowed,
            is_https_allowed=self.is_https_allowed,
            location=self.location,
            name=self.name,
            optimization_type=self.optimization_type,
            origin_groups=self.origin_groups,
            origin_host_header=self.origin_host_header,
            origin_path=self.origin_path,
            origins=self.origins,
            probe_path=self.probe_path,
            provisioning_state=self.provisioning_state,
            query_string_caching_behavior=self.query_string_caching_behavior,
            resource_state=self.resource_state,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type,
            url_signing_keys=self.url_signing_keys,
            web_application_firewall_policy_link=self.web_application_firewall_policy_link)


def get_endpoint(endpoint_name: Optional[builtins.str] = None,
                 profile_name: Optional[builtins.str] = None,
                 resource_group_name: Optional[builtins.str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEndpointResult:
    """
    Gets an existing CDN endpoint with the specified endpoint name under the specified subscription, resource group and profile.

    Uses Azure REST API version 2024-09-01.

    Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str endpoint_name: Name of the endpoint under the profile which is unique globally.
    :param builtins.str profile_name: Name of the CDN profile which is unique within the resource group.
    :param builtins.str resource_group_name: Name of the Resource group within the Azure subscription.
    """
    __args__ = dict()
    __args__['endpointName'] = endpoint_name
    __args__['profileName'] = profile_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:cdn:getEndpoint', __args__, opts=opts, typ=GetEndpointResult).value

    return AwaitableGetEndpointResult(
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        content_types_to_compress=pulumi.get(__ret__, 'content_types_to_compress'),
        custom_domains=pulumi.get(__ret__, 'custom_domains'),
        default_origin_group=pulumi.get(__ret__, 'default_origin_group'),
        delivery_policy=pulumi.get(__ret__, 'delivery_policy'),
        geo_filters=pulumi.get(__ret__, 'geo_filters'),
        host_name=pulumi.get(__ret__, 'host_name'),
        id=pulumi.get(__ret__, 'id'),
        is_compression_enabled=pulumi.get(__ret__, 'is_compression_enabled'),
        is_http_allowed=pulumi.get(__ret__, 'is_http_allowed'),
        is_https_allowed=pulumi.get(__ret__, 'is_https_allowed'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        optimization_type=pulumi.get(__ret__, 'optimization_type'),
        origin_groups=pulumi.get(__ret__, 'origin_groups'),
        origin_host_header=pulumi.get(__ret__, 'origin_host_header'),
        origin_path=pulumi.get(__ret__, 'origin_path'),
        origins=pulumi.get(__ret__, 'origins'),
        probe_path=pulumi.get(__ret__, 'probe_path'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        query_string_caching_behavior=pulumi.get(__ret__, 'query_string_caching_behavior'),
        resource_state=pulumi.get(__ret__, 'resource_state'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        url_signing_keys=pulumi.get(__ret__, 'url_signing_keys'),
        web_application_firewall_policy_link=pulumi.get(__ret__, 'web_application_firewall_policy_link'))
def get_endpoint_output(endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                        profile_name: Optional[pulumi.Input[builtins.str]] = None,
                        resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetEndpointResult]:
    """
    Gets an existing CDN endpoint with the specified endpoint name under the specified subscription, resource group and profile.

    Uses Azure REST API version 2024-09-01.

    Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str endpoint_name: Name of the endpoint under the profile which is unique globally.
    :param builtins.str profile_name: Name of the CDN profile which is unique within the resource group.
    :param builtins.str resource_group_name: Name of the Resource group within the Azure subscription.
    """
    __args__ = dict()
    __args__['endpointName'] = endpoint_name
    __args__['profileName'] = profile_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:cdn:getEndpoint', __args__, opts=opts, typ=GetEndpointResult)
    return __ret__.apply(lambda __response__: GetEndpointResult(
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        content_types_to_compress=pulumi.get(__response__, 'content_types_to_compress'),
        custom_domains=pulumi.get(__response__, 'custom_domains'),
        default_origin_group=pulumi.get(__response__, 'default_origin_group'),
        delivery_policy=pulumi.get(__response__, 'delivery_policy'),
        geo_filters=pulumi.get(__response__, 'geo_filters'),
        host_name=pulumi.get(__response__, 'host_name'),
        id=pulumi.get(__response__, 'id'),
        is_compression_enabled=pulumi.get(__response__, 'is_compression_enabled'),
        is_http_allowed=pulumi.get(__response__, 'is_http_allowed'),
        is_https_allowed=pulumi.get(__response__, 'is_https_allowed'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        optimization_type=pulumi.get(__response__, 'optimization_type'),
        origin_groups=pulumi.get(__response__, 'origin_groups'),
        origin_host_header=pulumi.get(__response__, 'origin_host_header'),
        origin_path=pulumi.get(__response__, 'origin_path'),
        origins=pulumi.get(__response__, 'origins'),
        probe_path=pulumi.get(__response__, 'probe_path'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        query_string_caching_behavior=pulumi.get(__response__, 'query_string_caching_behavior'),
        resource_state=pulumi.get(__response__, 'resource_state'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        url_signing_keys=pulumi.get(__response__, 'url_signing_keys'),
        web_application_firewall_policy_link=pulumi.get(__response__, 'web_application_firewall_policy_link')))
