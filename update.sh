#!/bin/bash
set -e
CODEGEN_URL=https://repo.maven.apache.org/maven2/org/openapitools/openapi-generator-cli/4.2.3/openapi-generator-cli-4.2.3.jar
JAR=openapi-generator-cli.jar

pushd "$(dirname "$0")"

if [ $(git ls-files |grep -v "$(basename "$0")" |wc -l) -gt 0 ]; then
    git ls-files |grep -v "$(basename "$0")" |xargs git rm --ignore-unmatch -f
fi

wget $CODEGEN_URL -O $JAR

java -jar $JAR generate \
    --input-spec https://developers.strava.com/swagger/swagger.json \
    --generator-name go \
    --additional-properties=packageName=strava \
    --additional-properties=hideGenerationTimestamp=false \
    --skip-validate-spec \
    --output .
rm -f $JAR

git add .

popd
