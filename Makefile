AppName=ac
BuildDir=build
PlistPath=./ios/Info.plist
IconName=MyIcon
IconPath=assets/${IconName}.icns

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

build: prepare build-linux build-mac

prepare:
	mkdir -p ${BuildDir}/macOS/${AppName}.app/Contents/MacOS
	mkdir -p ${BuildDir}/macOS/${AppName}.app/Contents/Resources
	mkdir -p ${BuildDir}/linux
	cp ${PlistPath} ${BuildDir}/macOS/${AppName}.app/Contents/Info.plist
	cp ${IconPath} ${BuildDir}/macOS/${AppName}.app/Contents/Resources/${IconName}.icns

build-mac:
	GOOS=darwin GOARCH=$(ARCH) go build -o ${BuildDir}/macOS/${AppName}.app/Contents/MacOS/${AppName} ./application

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ${BuildDir}/linux/${AppName} ./application

install-linux:
	@echo "Installing ${AppName} for the current user..."
	mkdir -p ~/bin
	cp ${BuildDir}/linux/${AppName} ~/bin/
	chmod +x ~/bin/${AppName}
	if [ ! -f ~/.bash_profile ] && [ ! -f ~/.bashrc ] && [ ! -f ~/.zshrc ]; then \
		echo 'export PATH="$$HOME/bin:$$PATH"' >> ~/.zshrc; \
		echo "No bash or zsh config found. Defaulted to ~/.zshrc"; \
	fi
	if [ -f ~/.bashrc ]; then \
		grep -qxF 'export PATH="$$HOME/bin:$$PATH"' ~/.bashrc || echo 'export PATH="$$HOME/bin:$$PATH"' >> ~/.bashrc; \
	elif [ -f ~/.bash_profile ]; then \
		grep -qxF 'export PATH="$$HOME/bin:$$PATH"' ~/.bash_profile || echo 'export PATH="$$HOME/bin:$$PATH"' >> ~/.bash_profile; \
	fi
	if [ -f ~/.zshrc ]; then \
		grep -qxF 'export PATH="$$HOME/bin:$$PATH"' ~/.zshrc || echo 'export PATH="$$HOME/bin:$$PATH"' >> ~/.zshrc; \
	fi
	@echo "Installation complete. Please open a new terminal or source your shell configuration file to update your PATH."

install-mac:
	@echo "Installing ${AppName} for the current user on macOS..."
	mkdir -p ~/bin
	cp ${BuildDir}/macOS/${AppName}.app/Contents/MacOS/${AppName} ~/bin/
	chmod +x ~/bin/${AppName}
	if [ ! -f ~/.bash_profile ] && [ ! -f ~/.bashrc ] && [ ! -f ~/.zshrc ]; then \
		echo 'export PATH="$$HOME/bin:$$PATH"' >> ~/.zshrc; \
		echo "No bash or zsh config found. Defaulted to ~/.zshrc"; \
	fi
	if [ -f ~/.bash_profile ]; then \
		grep -qxF 'export PATH="$$HOME/bin:$$PATH"' ~/.bash_profile || echo 'export PATH="$$HOME/bin:$$PATH"' >> ~/.bash_profile; \
	elif [ -f ~/.bashrc ]; then \
		grep -qxF 'export PATH="$$HOME/bin:$$PATH"' ~/.bashrc || echo 'export PATH="$$HOME/bin:$$PATH"' >> ~/.bashrc; \
	fi
	if [ -f ~/.zshrc ]; then \
		grep -qxF 'export PATH="$$HOME/bin:$$PATH"' ~/.zshrc || echo 'export PATH="$$HOME/bin:$$PATH"' >> ~/.zshrc; \
	fi
	@echo "Installation complete. Please open a new terminal or source your shell configuration file to update your PATH."

clean:
	rm -rf ${BuildDir}

mac: build-mac install-mac
linux: build-linux install-linux

.PHONY: all build prepare build-mac build-linux install install-linux install-mac clean mac linux
.DEFAULT: install
