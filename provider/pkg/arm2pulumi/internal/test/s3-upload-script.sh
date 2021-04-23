#!/bin/bash

currDateTime=$(date +"%Y-%m-%d_%H-%M-%S")
newSummaryName="summary_${currDateTime}.json"
echo "New file to be uploaded: ${newSummaryName}"

cd test-results
sed '1d' summary.json | sed '$d' | sed -E 's/},/}/' > $newSummaryName
aws s3 cp $newSummaryName s3://temp-test-data-bucket-3b9ef45/summaries/$newSummaryName