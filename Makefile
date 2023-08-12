bin: bin/preflight-netpath_darwin_amd64 bin/preflight-netpath_linux_amd64 bin/preflight-netpath_windows_amd64.exe
bin: bin/preflight-netpath_darwin_arm64 bin/preflight-netpath_linux_arm64 bin/preflight-netpath_windows_arm64.exe

bin/preflight-netpath_darwin_amd64:
	@mkdir -p bin
	@echo "Compiling preflight-netpath..."
	GOOS=darwin GOARCH=amd64 go build -o $@ cmd/preflight-netpath/*.go

bin/preflight-netpath_darwin_arm64:
	@mkdir -p bin
	@echo "Compiling preflight-netpath..."
	GOOS=darwin GOARCH=arm64 go build -o $@ cmd/preflight-netpath/*.go

bin/preflight-netpath_linux_amd64:
	@mkdir -p bin
	@echo "Compiling preflight-netpath..."
	GOOS=linux GOARCH=amd64 go build -o $@ cmd/preflight-netpath/*.go

bin/preflight-netpath_linux_arm64:
	@mkdir -p bin
	@echo "Compiling preflight-netpath..."
	GOOS=linux GOARCH=arm64 go build -o $@ cmd/preflight-netpath/*.go

bin/preflight-netpath_windows_amd64.exe:
	@mkdir -p bin
	@echo "Compiling preflight-netpath..."
	GOOS=windows GOARCH=amd64 go build -o $@ cmd/preflight-netpath/*.go

bin/preflight-netpath_windows_arm64.exe:
	@mkdir -p bin
	@echo "Compiling preflight-netpath..."
	GOOS=windows GOARCH=arm64 go build -o $@ cmd/preflight-netpath/*.go

.PHONY: install
install: bin
	@echo "Installing preflight-netpath..."
	@scp bin/preflight-netpath_$$(go env GOOS)_$$(go env GOARCH) /usr/local/bin/preflight-netpath