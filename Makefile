.PHONY : build
build :
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build  -o ./target/tiktok-tools-darwin-amd64 .
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build  -o ./target/tiktok-tools-darwin-arm64 .
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o ./target/tiktok-tools-windows-amd64.exe .
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o ./target/tiktok-tools-linux-amd64 .




.PHONY : clean
clean :
	rm -rf ./target