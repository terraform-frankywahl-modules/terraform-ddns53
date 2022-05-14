TF_DOCS = docker run --rm --volume $(shell pwd):/terraform-docs quay.io/terraform-docs/terraform-docs:latest

.PHONY: readme
readme:
	make -C _examples/ readme
	$(TF_DOCS) markdown /terraform-docs --hide-empty > README.md
