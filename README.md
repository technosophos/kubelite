# KubeLite: Minimal Kubernetes API

Perfect for when you need the Kubernetes APIs, but not the rest of Kubernetes.

This package provides the following:

- The v1 API, marhsalable to JSON or YAML
- Marshal and Unmarshal implementations for JSON or YAML

This library can be used to:

- Read or write Kubernetes manifests
- Bundle and unbundle Kubernetes manifests

This library is _not_ intended to...

- Link to (compile into) Kubernetes
- Provide a client
- Provide any of the core services of Kubernetes
- Perform full validation of a manifest. (This provides syntax
  validation only)

## Usage

```
manifests, err := codec.YAML(myBytes).All()

//...

for _, manifest := range manifests {

    // Find out what kind of object it is.
    obj, _ := manifest.Ref()
    println(obj.APIVersion)
    println(obj.Kind)
}

```

## Generating This Package from Upstream

Many parts of this package are generated computationally from the actual
Kubernetes repository. You can generate it yourself using `gen.sh`.

*Notice:* This project contains code from the Kubernetes project. All
code included is Copyright 2014,2015 The Kubernetes Authors, and is made
available under the Apache 2.0 license.
