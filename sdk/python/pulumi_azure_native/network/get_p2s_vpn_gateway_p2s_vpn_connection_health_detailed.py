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

__all__ = [
    'GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult',
    'AwaitableGetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult',
    'get_p2s_vpn_gateway_p2s_vpn_connection_health_detailed',
    'get_p2s_vpn_gateway_p2s_vpn_connection_health_detailed_output',
]

@pulumi.output_type
class GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult:
    """
    P2S Vpn connection detailed health written to sas url.
    """
    def __init__(__self__, sas_url=None):
        if sas_url and not isinstance(sas_url, str):
            raise TypeError("Expected argument 'sas_url' to be a str")
        pulumi.set(__self__, "sas_url", sas_url)

    @property
    @pulumi.getter(name="sasUrl")
    def sas_url(self) -> Optional[builtins.str]:
        """
        Returned sas url of the blob to which the p2s vpn connection detailed health will be written.
        """
        return pulumi.get(self, "sas_url")


class AwaitableGetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult(GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult(
            sas_url=self.sas_url)


def get_p2s_vpn_gateway_p2s_vpn_connection_health_detailed(gateway_name: Optional[builtins.str] = None,
                                                           output_blob_sas_url: Optional[builtins.str] = None,
                                                           resource_group_name: Optional[builtins.str] = None,
                                                           vpn_user_names_filter: Optional[Sequence[builtins.str]] = None,
                                                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult:
    """
    Gets the sas url to get the connection health detail of P2S clients of the virtual wan P2SVpnGateway in the specified resource group.

    Uses Azure REST API version 2024-05-01.

    Other available API versions: 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str gateway_name: The name of the P2SVpnGateway.
    :param builtins.str output_blob_sas_url: The sas-url to download the P2S Vpn connection health detail.
    :param builtins.str resource_group_name: The name of the resource group.
    :param Sequence[builtins.str] vpn_user_names_filter: The list of p2s vpn user names whose p2s vpn connection detailed health to retrieve for.
    """
    __args__ = dict()
    __args__['gatewayName'] = gateway_name
    __args__['outputBlobSasUrl'] = output_blob_sas_url
    __args__['resourceGroupName'] = resource_group_name
    __args__['vpnUserNamesFilter'] = vpn_user_names_filter
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:network:getP2sVpnGatewayP2sVpnConnectionHealthDetailed', __args__, opts=opts, typ=GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult).value

    return AwaitableGetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult(
        sas_url=pulumi.get(__ret__, 'sas_url'))
def get_p2s_vpn_gateway_p2s_vpn_connection_health_detailed_output(gateway_name: Optional[pulumi.Input[builtins.str]] = None,
                                                                  output_blob_sas_url: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                                                  resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                                                  vpn_user_names_filter: Optional[pulumi.Input[Optional[Sequence[builtins.str]]]] = None,
                                                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult]:
    """
    Gets the sas url to get the connection health detail of P2S clients of the virtual wan P2SVpnGateway in the specified resource group.

    Uses Azure REST API version 2024-05-01.

    Other available API versions: 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str gateway_name: The name of the P2SVpnGateway.
    :param builtins.str output_blob_sas_url: The sas-url to download the P2S Vpn connection health detail.
    :param builtins.str resource_group_name: The name of the resource group.
    :param Sequence[builtins.str] vpn_user_names_filter: The list of p2s vpn user names whose p2s vpn connection detailed health to retrieve for.
    """
    __args__ = dict()
    __args__['gatewayName'] = gateway_name
    __args__['outputBlobSasUrl'] = output_blob_sas_url
    __args__['resourceGroupName'] = resource_group_name
    __args__['vpnUserNamesFilter'] = vpn_user_names_filter
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:network:getP2sVpnGatewayP2sVpnConnectionHealthDetailed', __args__, opts=opts, typ=GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult)
    return __ret__.apply(lambda __response__: GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult(
        sas_url=pulumi.get(__response__, 'sas_url')))
