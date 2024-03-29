generate-api:
	oapi-codegen \
		-config $(PWD)/api/http/oapi.cfg.yaml \
		-package api \
		$(PWD)/api/http/api.openapi.yaml \
		> $(PWD)/api/http/openapi.gen.go

sqlc-gen:
	docker run -u $(UID):$(GID) -v $(PWD):/app -w /app kjconroy/sqlc generate