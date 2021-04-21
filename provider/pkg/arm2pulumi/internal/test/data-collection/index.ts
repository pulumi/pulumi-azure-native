import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";
import * as mime from "mime";
import { Client as RedshiftClient } from "pg";

// Create an AWS S3 to store our coverage resutls
const dataBucket = new aws.s3.Bucket("data-bucket");

// Directory where the results are currenlty located
const dataDir = "../";

const dataFile = "test-results/summary.json" // file that contains summary of results
const JSONPathFile = "test-results/summary_JSONPath.json" // file that contains JSONPath structure of results

const dataFilePath = dataDir + dataFile;
const JSONPathFilePath = dataDir + JSONPathFile;

// Create an S3 object stored in `dataBucket` for the summary of results
const dataObject = new aws.s3.BucketObject(dataFilePath, {
    bucket: dataBucket,                                   // reference the s3.Bucket object
    source: new pulumi.asset.FileAsset(dataFilePath),     // use FileAsset to point to a file
    contentType: mime.getType(dataFilePath) || undefined, // set the MIME type of the file, TODO: hardcode to JSON?
});

// Create an S3 object stored in `dataBucket` for the JSONPath file
const JSONPathObject = new aws.s3.BucketObject(JSONPathFilePath, {
    bucket: dataBucket,                               // reference the s3.Bucket object
    source: new pulumi.asset.FileAsset(JSONPathFilePath),     // use FileAsset to point to a file
    contentType: mime.getType(JSONPathFilePath) || undefined, // set the MIME type of the file, TODO: hardcode to JSON?
});

// Set the access policy for the bucket so all objects are readable
const bucketPolicy = new aws.s3.BucketPolicy("bucketPolicy", {
    bucket: dataBucket.id, // refer to the bucket created earlier
    policy: dataBucket.arn.apply(bucketArn => JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Effect: "Allow",
            Principal: "*",
            Action: [
                "s3:GetObject",
            ],
            Resource: [
                `${bucketArn}/*`,
            ],
        }],
    })),
});


// const s3JSONPathSource: string = `s3://${dataBucket.id.apply(v => `prefix${v}suffix`)}/${JSONPathFile}`;

const name: string = "temp-test-";

// Pulled from lines 44-71: https://github.com/pulumi/home/blob/master/infrastructure/datawarehouse/datawarehouse.ts
const spectrumRole = new aws.iam.Role(`${name}warehouse-redshiftspectrum`, {
    assumeRolePolicy: aws.iam.assumeRolePolicyForPrincipal({ Service: "redshift.amazonaws.com" }),
});

const spectrumPermissions = new aws.iam.RolePolicy(`${name}warehouse-redshiftspectrum`, {
    role: spectrumRole.id,
    // See https://docs.aws.amazon.com/redshift/latest/dg/c-spectrum-iam-policies.html for details on IAM
    // policies for Redshift Spectrum.
    policy: {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Action": [
                    // Allow import from S3
                    "s3:Get*",
                    "s3:List*",
                    // Allow access to data catalogs in Athena and Glue
                    // "athena:*",
                    // "glue:GetTable",
                    // "glue:GetDatabase",
                    // "glue:GetDatabases",
                    // "glue:CreateDatabase",
                ],
                "Resource": "*",
            },
        ],
    },
});

// Pulled from lines 131-147: https://github.com/pulumi/home/blob/master/infrastructure/datawarehouse/datawarehouse.ts
const redshift = new aws.redshift.Cluster("default", {
    clusterIdentifier: "tf-redshift-cluster",
    clusterType: "single-node",
    databaseName: "mydb",
    masterPassword: "tesT!ng123", // temporary value that is hardcoded to make sure correct password is passed to the lambda
    masterUsername: "foo",
    nodeType: "dc2.large",
    iamRoles: [spectrumRole.arn],
    skipFinalSnapshot: true,
    //     vpcSecurityGroupIds: [redshiftSecurityGroup.id], <-- Is not having this set up causing the authentication issue?
});

// Currently need to manually create table with following query:
// DROP TABLE IF EXISTS coverage_results;

// CREATE TABLE IF NOT EXISTS coverage_results(
// 	sourceElement VARCHAR(10000),
// 	sourceToken VARCHAR(10000),
//   	severity VARCHAR(1000),
//   	description VARCHAR(10000),
//   	template VARCHAR(1000)
// 	);


const redshiftPort = redshift.port.apply(p => p || 5439);

const fiveMinInSecs = 5 * 60;

// Returns a lambda that automatically copies any new summary of results file added to `dataBucket`
// into the redshift cluster.
// Based on lines 50-108: https://github.com/pulumi/home/blob/master/infrastructure/datawarehouse/alblogs.ts
function copyToRedshiftLambda() {
    return new aws.lambda.CallbackFunction<aws.s3.BucketEvent, void>("copy-results-to-redshift", {
        policies: [ // TODO: Not sure if these policies are necessary
            aws.iam.ManagedPolicy.LambdaFullAccess,
            aws.iam.ManagedPolicies.AWSLambdaVPCAccessExecutionRole, 
        ],
        callback: async (e) => {
            const s3JSONPathSource: string = `s3://${dataBucket.id.get()}/${JSONPathFile}`;
            console.log(e.Records || []);
            for (const rec of e.Records || []) {
                const [ buck, key ] = [ rec.s3.bucket.name, rec.s3.object.key ];
                if (key !== JSONPathFile) {
                    console.log(`Copying data from ${key} to Redshift`);
                    const client = new RedshiftClient({
                        user: redshift.masterUsername.get(),
                        host: redshift.dnsName.get(),
                        database: redshift.databaseName.get(),
                        password: redshift.masterPassword.get(),
                        port: redshiftPort.get(),
                        // Tried hardcoding values to see if that would get a result (did not)
                        // user: "foo",
                        // host: "tf-redshift-cluster.cizihz8m4fgu.us-west-2.redshift.amazonaws.com",
                        // database: "mydb",
                        // password: "tesT!ng123",
                        // port: redshiftPort.get(),
                    });
                    
                    console.log("Connecting to cluster...");
                    await client.connect();
                    const s3DataSource: string = `s3://${buck}/${key}`;
                    
                    console.log("Copying S3 data from:")
                    console.log(s3DataSource);
                    console.log("using:");
                    console.log(s3JSONPathSource);
                    
                    const copyQuery = 
                    `COPY coverage_results
                    from '${s3DataSource}'
                    credentials 'aws_iam_role=${spectrumRole.arn}'
                    json '${s3JSONPathSource}';`
                    // json 'auto';`

                    console.log("trying to copy:");

                    const copyRes = await client.query(copyQuery);
                    
                    console.log("S3 copy operation response:");
                    console.log(JSON.stringify(copyRes));
                    
                    await client.end();
                    console.log("done");
                }
                console.log(`Hello from Lambda -- got an S3 Object: ${buck}/${key}`);
                
            }
        },
        timeout: fiveMinInSecs,
    })
}

dataBucket.onObjectCreated("copy-results-to-redshift", copyToRedshiftLambda());

// Export the name of the bucket
export const bucketName = dataBucket.id;
export const bucketDomain = dataBucket.bucketDomainName;
export const redshiftId = redshift.id;
// export const redshiftDNS = redshift.dnsName;
export const redshiftDBName = redshift.databaseName;
export const spectrumRoleARN = spectrumRole.arn;

export const redshiftUser = redshift.masterUsername;
export const redshiftHost = redshift.dnsName;
export const redshiftDatabase = redshift.databaseName;
export const redshiftPassword = redshift.masterPassword; 
export const redshift_Port = redshiftPort;