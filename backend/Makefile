SRCEXT := .go
SRCDIRS := $(patsubst ./%,%,$(shell find . -type d))
SRCLIST	:= $(foreach d,$(SRCDIRS),$(wildcard $(addprefix $(d)/*, $(SRCEXT))))

backend: backend_deps backend_build

backend_build: $(SRCLIST)
	go env
	go build

.PHONY: backend_deps backend_test

backend_deps:
	go get ./...
