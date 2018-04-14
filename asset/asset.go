package asset

//go:generate rm -f bindata.go
//go:generate go-bindata -ignore=asset.go -ignore=.empty -pkg=$GOPACKAGE -nometadata ./...
