all:
    windows-build
    macos-build
    linux-build

generate-assets:
    go-bindata -o assets.go data
    go fmt assets.go

windows-build:
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o mnemonic-windows-amd64.exe

macos-build:
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o mnemonic-apple-amd64

linux-build:
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mnemonic-linux-amd64

clean:
    rm -v mnemonic-windows-amd64.exe
    rm -v mnemonic-apple-amd64
    rm -v mnemonic-linux-amd64