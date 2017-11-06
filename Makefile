REPO=k8s-node-update-scheduler
GITHUB_USER=mad01
ORG_PATH=github.com/$(GITHUB_USER)
REPO_PATH=$(ORG_PATH)/$(REPO)

VERSION ?= $(shell ./hacks/git-version)
LD_FLAGS="-X main.Version=$(VERSION)  -extldflags \"-static\" "
version.Version=$(VERSION)
$( shell mkdir -p _bin )
$( shell mkdir -p _release )

export GOBIN=$(PWD)/_bin


default: build

clean:
	@rm -r _bin _release

test:
	@go test -v -i $(shell go list ./... | grep -v '/vendor/')
	@go test -v $(shell go list ./... | grep -v '/vendor/')

build: build/dev

build/dev:
	@go install -v -ldflags $(LD_FLAGS) 

build/release:
	@go build -v -o _release/$(REPO) -ldflags $(LD_FLAGS) 


docker/build:
	@docker build -t quay.io/$(GITHUB_USER)/$(REPO):$(VERSION) --file Dockerfile .

docker/push:
	@docker push quay.io/$(GITHUB_USER)/$(REPO):$(VERSION)

docker/login:
	@docker login -u $(QUAY_LOGIN) -p="$(QUAY_PASSWORD)" quay.io
