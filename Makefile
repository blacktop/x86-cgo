REPO=blacktop
NAME=x86-cgo
CLI=github.com/blacktop/x86-cgo/cmd/disass
CUR_VERSION=$(shell svu current)
NEXT_VERSION=$(shell svu patch)


.PHONY: download
download:
	@echo "Fetching Dependancies"
	git clone https://github.com/intelxed/xed.git disassemble/xed
	git clone https://github.com/intelxed/mbuild.git disassemble/mbuild

	@echo "Reverting commit 9bdeca6d77065e5f1b23891655a26e510ffae74a"
	cd disassemble/xed && git revert 9bdeca6d77065e5f1b23891655a26e510ffae74a --no-edit

.PHONY: build-xed
build-xed:
	rm -rf disassemble/xedInc

	@echo "Building XED"
	# xed now assumes building in a subdir like ./build
	mkdir -p disassemble/build
	# this runs very well on Linux; but it sometimes has problems with Windows And Mac
	cd disassemble/build && \
	../xed/mfile.py -j 9 --static --extra-flags=-fPIC --opt=3 --no-encoder install --install-dir=../xedKit && \
	../xed/mfile.py -c

	@echo "Setting Up"
	mv disassemble/xedKit/include/xed disassemble/xedInc
	mkdir -p disassemble/lib

.PHONY: mac
mac: build-xed
	mv disassemble/xedKit/lib/libxed.a disassemble/lib/libxed_macos.a

.PHONY: build-deps
build-deps: ## Install the build dependencies
	@echo " > Installing build deps"
	brew install go goreleaser
	go get -u github.com/crazy-max/xgo

.PHONY: build
build: ## Build disass locally
	@echo " > Building locally"
	CGO_ENABLED=1 go build -o disass.${NEXT_VERSION} ./cmd/disass

.PHONY: test
test: ## Run disass on hello-mte
	@echo " > disassembling hello-mte\n"
	@dist/x86-cgo_darwin_amd64/disass  ../../Proteas/hello-mte/hello-mte _test

.PHONY: dry_release
dry_release: ## Run goreleaser without releasing/pushing artifacts to github
	@echo " > Creating Pre-release Build ${NEXT_VERSION}"
	@goreleaser build --rm-dist --skip-validate --snapshot -f .goreleaser.mac.yml

.PHONY: release
release: ## Create a new release from the NEXT_VERSION
	@echo " > Creating Release ${NEXT_VERSION}"
	@hack/make/release ${NEXT_VERSION}

.PHONY: cross
cross: ## Create xgo releases
	@echo " > Creating xgo releases"
	@mkdir -p dist/xgo
	@cd dist/xgo; xgo --targets=*/amd64 -go 1.16.5 -ldflags='-s -w' -out disass-${NEXT_VERSION} ${CLI}

clean: ## Clean up artifacts
	@echo " > Cleaning"
	rm -rf dist
	rm -rf disassemble/xedKit disassemble/build

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help