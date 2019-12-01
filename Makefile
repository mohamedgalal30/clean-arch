.PHONY: release major minor patch

release:
	@git checkout master
	@git merge dev
	@git push origin master
	@git push origin master --tags
	@git checkout dev
	@git push origin dev
	@git push origin dev --tags

major:
	@bumpversion major

minor:
	@bumpversion minor

patch:
	@bumpversion patch

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ticket-api .
	docker build -t ticket-api .
	docker tag ticket-api fly365dotcom/ticket-api:latest
	docker rmi ticket-api

push:
	docker push fly365dotcom/ticket-api:latest