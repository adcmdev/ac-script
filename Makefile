AppName=ac
BuildDir=build

install:
	$(eval UNAME_S := $(shell uname -s))
	$(eval UNAME_M := $(shell uname -m))
	$(eval ARCH := amd64)
	@if [ "$(UNAME_S)" = "Darwin" ]; then \
		if [ "$(UNAME_M)" = "x86_64" ]; then \
			ARCH=amd64; \
		elif [ "$(UNAME_M)" = "arm64" ]; then \
			ARCH=arm64; \
		fi; \
		$(MAKE) build-mac ARCH=$$ARCH; \
		$(MAKE) install-mac; \
	elif [ "$(UNAME_S)" = "Linux" ]; then \
		$(MAKE) build-linux; \
		$(MAKE) install-linux; \
	else \
		echo "Unsupported OS"; \
	fi

all: build

build: build-linux build-mac

build-mac:
	GOOS=darwin GOARCH=$(ARCH) go build -o ${BuildDir}/macOS/${AppName} ./application

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ${BuildDir}/linux/${AppName} ./application

install-linux:
	@echo "Installing ${AppName} for the current user..."
	mkdir -p ~/bin
	cp ${BuildDir}/linux/${AppName} ~/bin/
	chmod +x ~/bin/${AppName}
	if ! grep -q 'export PATH="$$HOME/bin:$$PATH"' ~/.zshrc; then \
		echo 'export PATH="$$HOME/bin:$$PATH"' >> ~/.zshrc; \
		echo "The ${AppName} executable has been added to your PATH."; \
	else \
		echo "The ${AppName} executable is already in your PATH."; \
	fi
	@echo "Installation complete. Please open a new terminal or source your shell configuration file to update your PATH."

install-mac:
	@echo "Installing ${AppName} for the current user on macOS..."
	mkdir -p ~/bin
	cp ${BuildDir}/macOS/${AppName} ~/bin/
	chmod +x ~/bin/${AppName}
	if ! grep -q 'export PATH="$$HOME/bin:$$PATH"' ~/.zshrc; then \
		echo 'export PATH="$$HOME/bin:$$PATH"' >> ~/.zshrc; \
		echo "The ${AppName} executable has been added to your PATH."; \
	else \
		echo "The ${AppName} executable is already in your PATH."; \
	fi
	@echo "Installation complete. Please open a new terminal or source your shell configuration file to update your PATH."

clean:
	rm -rf ${BuildDir}

mac: build-mac install-mac
linux: build-linux install-linux

.PHONY: all build build-mac build-linux install install-linux install-mac clean mac linux
.DEFAULT: install
