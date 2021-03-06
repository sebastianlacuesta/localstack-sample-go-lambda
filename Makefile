.PHONY: install update package clean main

ENDPOINT_URL=http://localhost:4566
FUNCTION=sample-lambda
ROLE=role

update: function.zip
	aws lambda update-function-code \
		--endpoint-url $(ENDPOINT_URL) \
		--function-name $(FUNCTION) \
		--zip-file fileb://function.zip

install: function.zip
	aws lambda create-function \
		--endpoint-url $(ENDPOINT_URL) \
		--runtime go1.x \
		--role $(ROLE) \
		--function-name $(FUNCTION) \
		--handler main \
		--zip-file fileb://function.zip

package: function.zip

function.zip: main
	zip function.zip main

main:
	go build -gcflags='all=-N -l' main.go

clean:
	rm -rf main function.zip
