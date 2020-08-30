all: build

TAG?=dev
FLAGS=
LDFLAGS?=-s
ENVVAR=CGO_ENABLED=0 GO111MODULE=off
GOOS?=linux
GOARCH?=amd64
#REGISTRY?=staging-k8s.gcr.io
REGISTRY?=registry.gitlab.com/frederic.boltz/k8s-autoscaler
BASEIMAGE?=gcr.io/distroless/static:latest
ifdef BUILD_TAGS
  TAGS_FLAG=--tags ${BUILD_TAGS}
  PROVIDER=-${BUILD_TAGS}
  FOR_PROVIDER=" for ${BUILD_TAGS}"
else
  TAGS_FLAG=
  PROVIDER=
  FOR_PROVIDER=
endif
ifdef LDFLAGS
  LDFLAGS_FLAG=--ldflags "${LDFLAGS}"
else
  LDFLAGS_FLAG=
endif

build: clean deps
	$(ENVVAR) GOOS=$(GOOS) go build ./... ${TAGS_FLAG}
	$(ENVVAR) GOOS=$(GOOS) go build -o ./out/cluster-autoscaler-$(GOOS)-$(GOARCH) ${TAGS_FLAG}

build-binary: clean deps
	$(ENVVAR) GOOS=linux go build -o ./out/cluster-autoscaler-linux-$(GOARCH) ${TAGS_FLAG}
	$(ENVVAR) GOOS=darwin go build -o ./out/cluster-autoscaler-darwin-$(GOARCH) ${TAGS_FLAG}

test-unit: clean deps build
	GO111MODULE=off go test --test.short -race ./... ${TAGS_FLAG}

dev-release: build-binary execute-release
	@echo "Release ${TAG}${FOR_PROVIDER} completed"

make-image:
ifdef BASEIMAGE
	docker build --pull --build-arg BASEIMAGE=${BASEIMAGE} \
	    -t ${REGISTRY}/cluster-autoscaler${PROVIDER}:${TAG} .
else
	docker build --pull -t ${REGISTRY}/cluster-autoscaler${PROVIDER}:${TAG} .
endif

push-image:
	./push_image.sh ${REGISTRY}/cluster-autoscaler${PROVIDER}:${TAG}

execute-release: make-image push-image

clean:
	rm -f ./out/cluster-autoscaler*

generate:
	go generate ./cloudprovider/aws

format:
	test -z "$$(find . -path ./vendor -prune -type f -o -name '*.go' -exec gofmt -s -d {} + | tee /dev/stderr)" || \
    test -z "$$(find . -path ./vendor -prune -type f -o -name '*.go' -exec gofmt -s -w {} + | tee /dev/stderr)"

docker-builder:
	docker build -t autoscaling-builder ../builder

build-in-docker: clean docker-builder
	docker run --rm -v `pwd`:/gopath/src/k8s.io/autoscaler/cluster-autoscaler/ autoscaling-builder:latest bash -c 'cd /gopath/src/k8s.io/autoscaler/cluster-autoscaler && BUILD_TAGS=${BUILD_TAGS} make build-binary'

release: build-in-docker execute-release
	@echo "Full in-docker release ${TAG}${FOR_PROVIDER} completed"

container: clean build-in-docker make-image
	@echo "Created in-docker image ${TAG}${FOR_PROVIDER}"

container-push: build-in-docker make-image
	docker push ${REGISTRY}/cluster-autoscaler${PROVIDER}:${TAG}
	@echo "Created in-docker image ${TAG}${FOR_PROVIDER}"

test-in-docker: clean docker-builder
	docker run --rm -v `pwd`:/gopath/src/k8s.io/autoscaler/cluster-autoscaler/ autoscaling-builder:latest bash -c 'cd /gopath/src/k8s.io/autoscaler/cluster-autoscaler && godep go test -race ./... ${TAGS_FLAG}'

.PHONY: all deps build test-unit clean format execute-release dev-release docker-builder build-in-docker release generate

