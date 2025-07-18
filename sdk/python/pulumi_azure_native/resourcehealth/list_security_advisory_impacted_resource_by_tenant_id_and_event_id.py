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
    'ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult',
    'AwaitableListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult',
    'list_security_advisory_impacted_resource_by_tenant_id_and_event_id',
    'list_security_advisory_impacted_resource_by_tenant_id_and_event_id_output',
]

@pulumi.output_type
class ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult:
    """
    The List of eventImpactedResources operation response.
    """
    def __init__(__self__, next_link=None, value=None):
        if next_link and not isinstance(next_link, str):
            raise TypeError("Expected argument 'next_link' to be a str")
        pulumi.set(__self__, "next_link", next_link)
        if value and not isinstance(value, list):
            raise TypeError("Expected argument 'value' to be a list")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="nextLink")
    def next_link(self) -> Optional[builtins.str]:
        """
        The URI to fetch the next page of events. Call ListNext() with this URI to fetch the next page of impacted resource.
        """
        return pulumi.get(self, "next_link")

    @property
    @pulumi.getter
    def value(self) -> Sequence['outputs.EventImpactedResourceResponse']:
        """
        The list of eventImpactedResources.
        """
        return pulumi.get(self, "value")


class AwaitableListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult(ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult(
            next_link=self.next_link,
            value=self.value)


def list_security_advisory_impacted_resource_by_tenant_id_and_event_id(event_tracking_id: Optional[builtins.str] = None,
                                                                       filter: Optional[builtins.str] = None,
                                                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult:
    """
    Lists impacted resources in the tenant by an event (Security Advisory).

    Uses Azure REST API version 2024-02-01.

    Other available API versions: 2022-10-01, 2023-07-01-preview, 2023-10-01-preview, 2025-04-01, 2025-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resourcehealth [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str event_tracking_id: Event Id which uniquely identifies ServiceHealth event.
    :param builtins.str filter: The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
    """
    __args__ = dict()
    __args__['eventTrackingId'] = event_tracking_id
    __args__['filter'] = filter
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:resourcehealth:listSecurityAdvisoryImpactedResourceByTenantIdAndEventId', __args__, opts=opts, typ=ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult).value

    return AwaitableListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult(
        next_link=pulumi.get(__ret__, 'next_link'),
        value=pulumi.get(__ret__, 'value'))
def list_security_advisory_impacted_resource_by_tenant_id_and_event_id_output(event_tracking_id: Optional[pulumi.Input[builtins.str]] = None,
                                                                              filter: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                                                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult]:
    """
    Lists impacted resources in the tenant by an event (Security Advisory).

    Uses Azure REST API version 2024-02-01.

    Other available API versions: 2022-10-01, 2023-07-01-preview, 2023-10-01-preview, 2025-04-01, 2025-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resourcehealth [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str event_tracking_id: Event Id which uniquely identifies ServiceHealth event.
    :param builtins.str filter: The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
    """
    __args__ = dict()
    __args__['eventTrackingId'] = event_tracking_id
    __args__['filter'] = filter
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:resourcehealth:listSecurityAdvisoryImpactedResourceByTenantIdAndEventId', __args__, opts=opts, typ=ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult)
    return __ret__.apply(lambda __response__: ListSecurityAdvisoryImpactedResourceByTenantIdAndEventIdResult(
        next_link=pulumi.get(__response__, 'next_link'),
        value=pulumi.get(__response__, 'value')))
