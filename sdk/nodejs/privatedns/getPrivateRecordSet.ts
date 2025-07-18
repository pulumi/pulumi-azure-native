// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Gets a record set.
 *
 * Uses Azure REST API version 2024-06-01.
 *
 * Other available API versions: 2018-09-01, 2020-01-01, 2020-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native privatedns [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getPrivateRecordSet(args: GetPrivateRecordSetArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateRecordSetResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:privatedns:getPrivateRecordSet", {
        "privateZoneName": args.privateZoneName,
        "recordType": args.recordType,
        "relativeRecordSetName": args.relativeRecordSetName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPrivateRecordSetArgs {
    /**
     * The name of the DNS zone (without a terminating dot).
     */
    privateZoneName: string;
    /**
     * The type of DNS record in this record set.
     */
    recordType: string;
    /**
     * The name of the record set, relative to the name of the zone.
     */
    relativeRecordSetName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: string;
}

/**
 * Describes a DNS record set (a collection of DNS records with the same name and type) in a Private DNS zone.
 */
export interface GetPrivateRecordSetResult {
    /**
     * The list of A records in the record set.
     */
    readonly aRecords?: outputs.privatedns.ARecordResponse[];
    /**
     * The list of AAAA records in the record set.
     */
    readonly aaaaRecords?: outputs.privatedns.AaaaRecordResponse[];
    /**
     * The Azure API version of the resource.
     */
    readonly azureApiVersion: string;
    /**
     * The CNAME record in the record set.
     */
    readonly cnameRecord?: outputs.privatedns.CnameRecordResponse;
    /**
     * The ETag of the record set.
     */
    readonly etag?: string;
    /**
     * Fully qualified domain name of the record set.
     */
    readonly fqdn: string;
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * Is the record set auto-registered in the Private DNS zone through a virtual network link?
     */
    readonly isAutoRegistered: boolean;
    /**
     * The metadata attached to the record set.
     */
    readonly metadata?: {[key: string]: string};
    /**
     * The list of MX records in the record set.
     */
    readonly mxRecords?: outputs.privatedns.MxRecordResponse[];
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The list of PTR records in the record set.
     */
    readonly ptrRecords?: outputs.privatedns.PtrRecordResponse[];
    /**
     * The SOA record in the record set.
     */
    readonly soaRecord?: outputs.privatedns.SoaRecordResponse;
    /**
     * The list of SRV records in the record set.
     */
    readonly srvRecords?: outputs.privatedns.SrvRecordResponse[];
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    readonly systemData: outputs.privatedns.SystemDataResponse;
    /**
     * The TTL (time-to-live) of the records in the record set.
     */
    readonly ttl?: number;
    /**
     * The list of TXT records in the record set.
     */
    readonly txtRecords?: outputs.privatedns.TxtRecordResponse[];
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
/**
 * Gets a record set.
 *
 * Uses Azure REST API version 2024-06-01.
 *
 * Other available API versions: 2018-09-01, 2020-01-01, 2020-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native privatedns [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export function getPrivateRecordSetOutput(args: GetPrivateRecordSetOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetPrivateRecordSetResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:privatedns:getPrivateRecordSet", {
        "privateZoneName": args.privateZoneName,
        "recordType": args.recordType,
        "relativeRecordSetName": args.relativeRecordSetName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPrivateRecordSetOutputArgs {
    /**
     * The name of the DNS zone (without a terminating dot).
     */
    privateZoneName: pulumi.Input<string>;
    /**
     * The type of DNS record in this record set.
     */
    recordType: pulumi.Input<string>;
    /**
     * The name of the record set, relative to the name of the zone.
     */
    relativeRecordSetName: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
