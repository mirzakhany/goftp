export ROOT=$(realpath $(dir $(firstword $(MAKEFILE_LIST))))
export APPNAME=goftp
export COMPANY_DOMAIN=mirzakhany
export SRC_PATH=$(COMPANY_DOMAIN)/$(APPNAME)
export BIN=$(ROOT)/bin
export DB_USER?=root
export RUSER=mohsen
export DB_NAME=mohsen
export GOBIN?=$(BIN)
export LINTER_BIN=$(shell which gometalinter)
export GO=$(shell which go)
export GIT=$(shell which git)
export GOPATH=$(abspath $(ROOT)/../../../..)
export LONG_HASH?=$(shell git log -n1 --pretty="format:%H" | cat)
export SHORT_HASH?=$(shell git log -n1 --pretty="format:%h"| cat)
export COMMIT_DATE?=$(shell git log -n1 --date="format:%D-%H-%I-%S" --pretty="format:%cd"| sed -e "s/\//-/g")
export IMP_DATE=$(shell date +%Y%m%d)
export COMMIT_COUNT?=$(shell git rev-list HEAD --count| cat)
export BUILD_DATE=$(shell date "+%D/%H/%I/%S"| sed -e "s/\//-/g")
export FLAGS="-X $(SRC_PATH)/commands/version.hash=$(LONG_HASH) -X $(SRC_PATH)/commands/version.short=$(SHORT_HASH) -X $(SRC_PATH)/commands/version.date=$(COMMIT_DATE) -X $(SRC_PATH)/commands/version.count=$(COMMIT_COUNT) -X $(SRC_PATH)/commands/version.build=$(BUILD_DATE)"
export LDARG=-ldflags $(FLAGS)
export BUILD=cd $(ROOT) && $(GO) install -v $(LDARG) -tags=jsoniter

all:
	$(BUILD) ./...

run: all
	$(ROOT)/bin/$(APPNAME)

needroot :
	@[ "$(shell id -u)" -eq "0" ] || exit 1

export LINTER=$(LINTER_BIN)
export LINTERCMD=$(LINTER) -e ".*.gen.go" -e ".*_test.go" -e "$(COMPANY_DOMAIN)/$(APP_NAME)/vendor/.*" --cyclo-over=19  --sort=path --disable-all --line-length=120 --deadline=100s --enable=structcheck --enable=deadcode --enable=gocyclo --enable=ineffassign --enable=golint --enable=goimports --enable=errcheck --enable=varcheck --enable=goconst --enable=megacheck --enable=misspell

lint: $(LINTER)
	$(LINTERCMD) $(ROOT)/commands/...

metalinter_install:
	$(GO) get -v github.com/alecthomas/gometalinter
	$(GO) install -v github.com/alecthomas/gometalinter
	$(LINTER) --install

$(LINTER):
	@[ -f $(LINTER) ] || make -f $(ROOT)/Makefile metalinter
