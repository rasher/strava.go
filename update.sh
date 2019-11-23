#!/bin/bash
CODEGEN_URL=http://central.maven.org/maven2/io/swagger/swagger-codegen-cli/2.4.9/swagger-codegen-cli-2.4.9.jar

pushd "$(dirname "$0")"

git ls-files |grep -v "$(basename "$0")" |xargs git rm -f

wget $CODEGEN_URL -O swagger-codegen-cli.jar

java -jar swagger-codegen-cli.jar generate -DpackageName=strava -DhideGenerationTimestamp=false --input-spec https://developers.strava.com/swagger/swagger.json --lang go --output .
rm -f swagger-codegen-cli.jar

git add .

popd
