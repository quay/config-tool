include .env

PID_FILE=/tmp/go_id.pid

all: 

generate: 
	go run ./utils/generate/gen.go

install:

	go install ./...

test:
	go test ./pkg/lib/fieldgroups/...

# Used to bring up production container
run-local-prod:
	docker build -t config-app:latest -f Dockerfile . && docker run -p 7070:8080 -v ${CONFIG_MOUNT}:/conf -ti config-app:latest editor --config-dir=/conf --password=password --operator-endpoint=${OPERATOR_ENDPOINT}

# Used to bring up dev container
build-local-dev:
	docker build -t config-app:dev -f Dockerfile.dev .

run-local-dev:
	docker run -p 7070:8080 \
	-v ${PWD}/pkg/lib/editor/js:/jssrc/js \
	-v ${PWD}/pkg/lib/editor/editor.go:/jssrc/editor.go \
	-v ${PWD}/:/go/src/config-tool \
	-v ${CONFIG_MOUNT}:/conf \
	-v ${CT_PRIVATE_KEY}:/tls/localhost.key \
	-v ${CT_PUBLIC_KEY}:/tls/localhost.crt \
	-e CONFIG_TOOL_PRIVATE_KEY=/tls/localhost.key \
	-e CONFIG_TOOL_PUBLIC_KEY=/tls/localhost.crt \
	-e DEBUGLOG=true \
	-ti config-app:dev

run-local-dev-setup:
	docker run -p 7070:8080 \
	-v ${PWD}/pkg/lib/editor/js:/jssrc/js \
	-v ${PWD}/pkg/lib/editor/editor.go:/jssrc/editor.go \
	-v ${PWD}/:/go/src/config-tool \
	-v ${CT_PRIVATE_KEY}:/tls/localhost.key \
	-v ${CT_PUBLIC_KEY}:/tls/localhost.crt \
	-e CONFIG_TOOL_PRIVATE_KEY=/tls/localhost.key \
	-e CONFIG_TOOL_PUBLIC_KEY=/tls/localhost.crt \
	-e DEBUGLOG=true \
	-ti config-app:dev


swagger:
	swag init -g pkg/lib/editor/editor.go