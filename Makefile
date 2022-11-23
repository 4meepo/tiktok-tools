.PHONY : build
build :
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build  -o ./target/tiktok-tools-darwin-amd64-${version} .
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build  -o ./target/tiktok-tools-darwin-arm64-${version} .
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o ./target/tiktok-tools-windows-amd64-${version}.exe .
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o ./target/tiktok-tools-linux-amd64-${version} .

.PHONY : package
package :
	mkdir -p ./target/package
	zip ./target/package/tiktok-tools-darwin-amd64-${version}.zip ./target/tiktok-tools-darwin-amd64-${version}
	zip ./target/package/tiktok-tools-darwin-arm64-${version}.zip ./target/tiktok-tools-darwin-arm64-${version}
	zip ./target/package/tiktok-tools-windows-amd64-${version}.zip ./target/tiktok-tools-windows-amd64-${version}.exe
	zip ./target/package/tiktok-tools-linux-amd64-${version}.zip ./target/tiktok-tools-linux-amd64-${version}

.PHONY : clean
clean :
	rm -rf ./target