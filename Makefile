TF_DOCS = terraform-docs

ifeq ($(DOCKER), true)
	TF_DOCS = docker run --rm --volume $(shell pwd):/terraform-docs --workdir /terraform-docs quay.io/terraform-docs/terraform-docs:latest
endif

.PHONY: default
default: help

%-docker: ## Run a make command using docker
	@DOCKER=true $(MAKE) $*

.PHONY: help
help: ## Show this help
	@echo "Usage: make <target>"
	@echo
	@echo "Targets:"
	@grep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: readme
readme: ## Generate the README.md file
	make -C _examples/ readme
	$(TF_DOCS) markdown . --hide-empty > README.md
