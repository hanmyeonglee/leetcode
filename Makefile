GO_FILES := $(filter-out sample.go,$(wildcard *.go))

.PHONY: prepare run git $(GO_FILES)

URL ?= $(or $(url),$(word 2,$(MAKECMDGOALS)))
FILE ?= $(or $(file),$(word 2,$(MAKECMDGOALS)))

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

run:
	@file="$(FILE)"; \
	if [ -z "$$file" ]; then \
		echo "Usage: make run <filename>"; \
		exit 1; \
	fi; \
	if [ ! -f "$$file" ]; then \
		echo "Error: File '$$file' does not exist."; \
		exit 1; \
	fi; \
	gofmt -s -w "$$file" && \
	go vet "$$file" && \
	go run "$$file"

git:
	@file="$(FILE)"; \
	if [ -z "$$file" ]; then \
		echo "Usage: make git <filename>"; \
		exit 1; \
	fi; \
	if [ ! -f "$$file" ]; then \
		echo "Error: File '$$file' does not exist."; \
		exit 1; \
	fi; \
	gofmt -s -w "$$file" && \
	go vet "$$file" && \
	name=$$(basename "$$file" .go); \
	msg="solve https://leetcode.com/problems/$$name"; \
	git add "$$file" && \
	git commit -m "$$msg" && \
	git push

ifneq ($(GO_FILES),)
$(GO_FILES):
	@:
endif

%:
	@:
