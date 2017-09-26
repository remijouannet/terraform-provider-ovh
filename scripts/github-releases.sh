#!/bin/bash

set +e

version=$(git describe --abbrev=0 --tags)

which github-release || echo 'please install the tool github-releases'

github-release info \
    --user remijouannet \
    --repo terraform-provider-ovh \
    --tag $version || echo "the release doesn't exist"

if [ $? != 0 ]
then
    github-release release \
        --user remijouannet \
        --repo terraform-provider-ovh \
        --tag $version \
        --draft \
        --pre-release \
        --name "$version-hyper-alpha-yolo-experimental" \
        --description "risks of explosion" \
        --target $version || echo "failed to create release for $version"
fi

cd pkg/

rm -f *.zip
 
ls | while read binary
do
    echo "zipping $binary"
    zip -r $binary.zip $binary
    echo "upload $binary"
    github-release upload \
        --user remijouannet \
        --name "$binary.zip" \
        --repo terraform-provider-ovh \
        --file "$binary.zip" \
        --replace \
        --tag $version || echo "failed to upload $binary"
done
