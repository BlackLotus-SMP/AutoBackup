#!/bin/bash

BUILD_DIR=$(dirname "$0")/build
mkdir -p "$BUILD_DIR"
cd "$BUILD_DIR" || exit

sum="sha1sum"

export GO111MODULE=on
echo "Setting GO111MODULE to" $GO111MODULE

if ! hash sha1sum 2>/dev/null; then
	if ! hash shasum 2>/dev/null; then
		echo "I can't see 'sha1sum' or 'shasum'"
		echo "Please install one of them!"
		exit
	fi
	sum="shasum"
fi

UPX=false
if hash upx 2>/dev/null; then
	UPX=true
fi

VERSION=$(date -u +%Y%m%d)
LDFLAGS="-X main.VERSION=$VERSION -s -w"
GCFLAGS=""

# AMD64
OSES=(linux darwin windows freebsd)
for os in "${OSES[@]}"; do
	suffix=""
	if [ "$os" == "windows" ]
	then
		suffix=".exe"
	fi
	env CGO_ENABLED=0 GOOS="$os" GOARCH=amd64 go build -ldflags "$LDFLAGS" -gcflags "$GCFLAGS" -o build_"${os}"_amd64"${suffix}" ..
	if $UPX; then upx -9 build_"${os}"_amd64"${suffix}";fi
	tar -zcf backup-"${os}"-amd64-"${VERSION}".tar.gz build_"${os}"_amd64"${suffix}"
	$sum backup-"${os}"-amd64-"${VERSION}".tar.gz
done

# 386
OSES=(linux windows)
for os in "${OSES[@]}"; do
	suffix=""
	if [ "$os" == "windows" ]
	then
		suffix=".exe"
	fi
	env CGO_ENABLED=0 GOOS="$os" GOARCH=386 go build -ldflags "$LDFLAGS" -gcflags "$GCFLAGS" -o build_"${os}"_386"${suffix}" ..
	if $UPX; then upx -9 build_"${os}"_386"${suffix}";fi
	tar -zcf backup-"${os}"-386-"${VERSION}".tar.gz build_"${os}"_386"${suffix}"
	$sum backup-"${os}"-386-"${VERSION}".tar.gz
done

#Apple M1 device
env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "$LDFLAGS" -gcflags "$GCFLAGS" -o build_darwin_arm64 ..
tar -zcf backup-darwin-arm64-"${VERSION}".tar.gz build_darwin_arm64
$sum backup-darwin-arm64-"${VERSION}".tar.gz

# ARM
ARMS=(5 6 7)
for v in "${ARMS[@]}"; do
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM="$v" go build -ldflags "$LDFLAGS" -gcflags "$GCFLAGS" -o build_linux_arm"${v}" ..
if $UPX; then upx -9 build_linux_arm"${v}";fi
tar -zcf backup-linux-arm"${v}"-"${VERSION}".tar.gz build_linux_arm"${v}"
$sum backup-linux-arm"${v}"-"${VERSION}".tar.gz
done

# ARM64
env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "$LDFLAGS" -gcflags "$GCFLAGS" -o build_linux_arm64 ..
if $UPX; then upx -9 build_linux_arm64;fi
tar -zcf backup-linux-arm64-"${VERSION}".tar.gz build_linux_arm64
$sum backup-linux-arm64-"${VERSION}".tar.gz

# MIPS32LE
env CGO_ENABLED=0 GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags "$LDFLAGS" -gcflags "$GCFLAGS" -o build_linux_mipsle ..
env CGO_ENABLED=0 GOOS=linux GOARCH=mips GOMIPS=softfloat go build -ldflags "$LDFLAGS" -gcflags "$GCFLAGS" -o build_linux_mips ..

if $UPX; then upx -9 build_linux_mips* backup_linux_mips*;fi
tar -zcf backup-linux-mipsle-"${VERSION}".tar.gz build_linux_mipsle
tar -zcf backup-linux-mips-"${VERSION}".tar.gz build_linux_mips
$sum backup-linux-mipsle-"${VERSION}".tar.gz
$sum backup-linux-mips-"${VERSION}".tar.gz

rm -rf ./*.tar.gz
