rm -rf output
mkdir -p output/bin

go build \
    -o output/bin/tornado \
    cmd/tornado/main.go
