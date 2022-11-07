#!/bin/bash
set -e
CODEGEN_VERSION=6.2.1
CODEGEN_URL=https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/${CODEGEN_VERSION}/openapi-generator-cli-${CODEGEN_VERSION}.jar
JAR=openapi-generator-cli.jar

pushd "$(dirname "$0")"

if [ $(git ls-files |grep -v "$(basename "$0")" |wc -l) -gt 0 ]; then
    git ls-files |grep -v "$(basename "$0")" |xargs git rm --ignore-unmatch -f
fi

wget $CODEGEN_URL -O $JAR

java -jar $JAR generate \
    --git-host github.com \
    --git-repo-id strava.go \
    --git-user-id rasher \
    --input-spec https://developers.strava.com/swagger/swagger.json \
    --generator-name go \
    --additional-properties=packageName=strava,hideGenerationTimestamp=false \
    --skip-validate-spec \
    --output .
rm -f $JAR

git add .

popd
TZ=UTC printf "Strava API as of %s\n" "$(date +%Y-%m-%d\ %H:%M:%S\ %Z)"
