PKGNAME    = proc-maps
DESTDIR   ?=
PREFIX    ?= /usr
BINDIR     = $(PREFIX)/bin

GOPROJROOT  = $(GOSRC)/$(PROJREPO)

GOLDFLAGS   = -ldflags "-s -w"
GOCC        = CGO_ENABLED=0 go
GOFMT       = $(GOCC) fmt -x
GOGET       = $(GOCC) get $(GOLDFLAGS)
GOBUILD     = $(GOCC) build -v $(GOLDFLAGS)
GOTEST      = $(GOCC) test
GOVET       = $(GOCC) vet
GOINSTALL   = $(GOCC) install $(GOLDFLAGS)

include Makefile.waterlog

GOLINT = golint -set_exit_status
STATICCHECK = staticcheck

all: build

build:
	@$(call stage,BUILD)
	@$(GOBUILD)
	@$(call pass,BUILD)

test: build
	@$(call stage,TEST)
	@$(GOTEST) ./...
	@$(call pass,TEST)

setup:
	@$(call stage,EXTRA-DEPS)
	@$(GOCC) mod download
	@$(call pass,EXTRA-DEPS)

validate: setup
	@$(call stage,FORMAT)
	@$(GOFMT) ./...
	@$(call pass,FORMAT)
	@$(call stage,VET)
	@$(call task,Running 'go vet'...)
	@$(GOVET) ./...
	@$(call pass,VET)
	@$(call stage,LINT)
	@$(call task,Running 'golint'...)
	@$(GOLINT) ./...
	@$(call task,Running 'staticcheck'...)
	@$(STATICCHECK) ./...
	@$(call pass,LINT)

install:
	@$(call stage,INSTALL)
	install -Dm 00755 $(PKGNAME) $(DESTDIR)$(BINDIR)/$(PKGNAME)
	@$(call pass,INSTALL)

uninstall:
	@$(call stage,UNINSTALL)
	rm -f $(DESTDIR)$(BINDIR)/$(PKGNAME)
	@$(call pass,UNINSTALL)

clean:
	@$(call stage,CLEAN)
	@$(call task,Removing executable...)
	@rm $(PKGNAME)
	@$(call pass,CLEAN)
