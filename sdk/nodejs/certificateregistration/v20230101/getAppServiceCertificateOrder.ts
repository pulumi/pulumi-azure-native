// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Description for Get a certificate order.
 */
export function getAppServiceCertificateOrder(args: GetAppServiceCertificateOrderArgs, opts?: pulumi.InvokeOptions): Promise<GetAppServiceCertificateOrderResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azure-native:certificateregistration/v20230101:getAppServiceCertificateOrder", {
        "certificateOrderName": args.certificateOrderName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAppServiceCertificateOrderArgs {
    /**
     * Name of the certificate order..
     */
    certificateOrderName: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    resourceGroupName: string;
}

/**
 * SSL certificate purchase order.
 */
export interface GetAppServiceCertificateOrderResult {
    /**
     * Reasons why App Service Certificate is not renewable at the current moment.
     */
    readonly appServiceCertificateNotRenewableReasons: string[];
    /**
     * <code>true</code> if the certificate should be automatically renewed when it expires; otherwise, <code>false</code>.
     */
    readonly autoRenew?: boolean;
    /**
     * State of the Key Vault secret.
     */
    readonly certificates?: {[key: string]: outputs.certificateregistration.v20230101.AppServiceCertificateResponse};
    /**
     * Contact info
     */
    readonly contact: outputs.certificateregistration.v20230101.CertificateOrderContactResponse;
    /**
     * Last CSR that was created for this order.
     */
    readonly csr?: string;
    /**
     * Certificate distinguished name.
     */
    readonly distinguishedName?: string;
    /**
     * Domain verification token.
     */
    readonly domainVerificationToken: string;
    /**
     * Certificate expiration time.
     */
    readonly expirationTime: string;
    /**
     * Resource Id.
     */
    readonly id: string;
    /**
     * Intermediate certificate.
     */
    readonly intermediate: outputs.certificateregistration.v20230101.CertificateDetailsResponse;
    /**
     * <code>true</code> if private key is external; otherwise, <code>false</code>.
     */
    readonly isPrivateKeyExternal: boolean;
    /**
     * Certificate key size.
     */
    readonly keySize?: number;
    /**
     * Kind of resource.
     */
    readonly kind?: string;
    /**
     * Certificate last issuance time.
     */
    readonly lastCertificateIssuanceTime: string;
    /**
     * Resource Location.
     */
    readonly location: string;
    /**
     * Resource Name.
     */
    readonly name: string;
    /**
     * Time stamp when the certificate would be auto renewed next
     */
    readonly nextAutoRenewalTimeStamp: string;
    /**
     * Certificate product type.
     */
    readonly productType: string;
    /**
     * Status of certificate order.
     */
    readonly provisioningState: string;
    /**
     * Root certificate.
     */
    readonly root: outputs.certificateregistration.v20230101.CertificateDetailsResponse;
    /**
     * Current serial number of the certificate.
     */
    readonly serialNumber: string;
    /**
     * Signed certificate.
     */
    readonly signedCertificate: outputs.certificateregistration.v20230101.CertificateDetailsResponse;
    /**
     * Current order status.
     */
    readonly status: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * Duration in years (must be 1).
     */
    readonly validityInYears?: number;
}
/**
 * Description for Get a certificate order.
 */
export function getAppServiceCertificateOrderOutput(args: GetAppServiceCertificateOrderOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAppServiceCertificateOrderResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azure-native:certificateregistration/v20230101:getAppServiceCertificateOrder", {
        "certificateOrderName": args.certificateOrderName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAppServiceCertificateOrderOutputArgs {
    /**
     * Name of the certificate order..
     */
    certificateOrderName: pulumi.Input<string>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    resourceGroupName: pulumi.Input<string>;
}
