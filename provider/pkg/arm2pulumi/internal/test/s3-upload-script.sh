#!/bin/bash

currDateTime=$(date +"%Y-%m-%d_%H-%M-%S")

# Makes sure that each summary has a unique name in S3
newSummaryName="summary_${currDateTime}.json"
echo "New file to be uploaded: ${newSummaryName}"

# URI of S3 bucket and location to upload to
s3KeyName="s3://arm2pulumi-coverage-results-236fe75/summaries/${newSummaryName}"

cd test-results
# Edit JSON summary file to be copiable into Redshift
# Changing file from a list of JSON objects to group of JSON objects. Done by 
# removing first ("[") and last line ("]") and replacing all "}," with "}".
sed '1d' summary.json | sed '$d' | sed -E 's/},/}/' > $newSummaryName
aws s3 cp $newSummaryName $s3KeyName