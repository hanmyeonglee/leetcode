.PHONY: prepare

URL ?= $(or $(url),$(word 2,$(MAKECMDGOALS)))

prepare:
	@url="$(URL)"; \
	if [ -z "$$url" ]; then \
		read -p "LeetCode URL: " url; \
	fi; \
	if [ -z "$$url" ]; then \
		echo "Usage: make prepare <leetcode-url> (or make prepare URL=<leetcode-url>)"; \
		exit 1; \
	fi; \
	name=$$(echo "$$url" | sed -E 's|[?#].*$$||; s|/+$$||; s|.*/problems/([^/]+).*|\1|; s|.*/||'); \
	file="$${name}.go"; \
	if [ -e "$$file" ]; then \
		echo "Error: $$file already exists."; \
		exit 1; \
	fi; \
	cp sample.go "$$file"; \
	echo "Created $$file from sample.go"

%:
	@:
