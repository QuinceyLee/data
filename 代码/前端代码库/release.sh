#! /bin/bash
build() {
    if [ ! -d "./_release/" ]; then
        mkdir -p ./_release/
    else
        rm -rf ./_release/
        mkdir -p ./_release/
    fi
    npm run build:prod
    zip -r $1.zip ./dist
    mv ./$1.zip ./_release

}
#-------------------------------------------------------
# Build
#-------------------------------------------------------
VERSION="$(git describe --tags $(git rev-list --tags --max-count=1))"
HASH=$(git rev-list --tags --max-count=1)
cat >./VERSION <<EOF
$VERSION
EOF
echo "----------------------------------->"
echo ">>>>   Build version: ${VERSION}"
echo "----------------------------------->"
build ${VERSION}
echo "---------------------------------------------->"
echo ">>>>   Build version: ${VERSION} successfully!"
echo "---------------------------------------------->"
