# Reproducible build example
Examples of reproducible builds for Rust, Go, and Go w/CGO_ENABLED=1. 

## Build Go

### CGO
    docker build -f Dockerfile.go .
    docker run -it --rm -v $(pwd)/go_cgo:$(pwd)/go_cgo -w $(pwd)/go_cgo -e CGO_ENABLED=1 e98460803fa4 build -v --trimpath --ldflags="-s -w -extldflags=-static" -o bin0

### Native
    docker run -it --rm -v $(pwd)/go:$(pwd)/go -w $(pwd)/go -e CGO_ENABLED=0 e98460803fa4 build -v --trimpath --ldflags="-s -w" -o bin0

## Build Rust w/Docker
    docker build -f Dockerfile.rust .
    docker run -it --rm -v $(pwd)/rust:$(pwd)/rust -w $(pwd)/rust d14447efd914 build --release 

## Compare Binaries
    # build two binaries named bin0 and bin1
    diff bin0 bin1

    # Compare Sha256 digests
    sha256sum -b bin0
    sha256sum -b bin1

    # View dependencies
    ldd bin0

    # View BuildID, dependency and target.
    file bin0 | tr , '\n'

## Wishlist
- Nix example
     nix build -f . --repeat 20 --enforce-determinism --diff-hook $(nix-build '<nixpkgs>' -A diffoscope) --run-diff-hook
- Guix example
