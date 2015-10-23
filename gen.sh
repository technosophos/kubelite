#!/bin/bash

prefix="https://raw.githubusercontent.com/kubernetes/kubernetes"
branch="master"

unversioned="$prefix/$branch/pkg/api/unversioned/types.go"
v1="$prefix/$branch/pkg/api/v1/types.go"
res="$prefix/$branch/pkg/api/resource/quantity.go"

echo "Fetching unversioned/types.go"
curl -o unversioned/types.go -SsL $unversioned
echo "Fetching v1/types.go"
curl -o v1/types.go -SsL $v1
#curl -o resource/quantity.go -SsL $res

echo "Rewriting imports on v1/types.go"
perl -pi -e "s|k8s\.io/kubernetes/pkg/api/unversioned|github.com/technosophos/kubelite/unversioned|g"  v1/types.go
perl -pi -e "s|k8s\.io/kubernetes/pkg/api/resource|github.com/technosophos/kubelite/resource|g"  v1/types.go
perl -pi -e "s|k8s\.io/kubernetes/pkg/types|github.com/technosophos/kubelite/types|g"  v1/types.go
perl -pi -e "s|\"k8s\.io/kubernetes/pkg/util\"||g" v1/types.go
perl -pi -e "s|\"k8s\.io/kubernetes/pkg/runtime\"||g" v1/types.go
perl -pi -e "s|util\.IntOrString|types.IntOrString|g" v1/types.go
perl -pi -e "s|runtime\.RawExtension|types.RawExtension|g" v1/types.go
