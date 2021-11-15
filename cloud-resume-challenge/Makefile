.PHONY: build

build:
	sam build

deploy-infra:
	make build && aws-vault exec my-user --no-session -- sam deploy --no-confirm-changeset

deploy-site:
	aws-vault exec my-user --no-session -- aws s3 sync ./resume-oliver-bennett s3://cloud-resume-challenge-website


invoke-get:
	sam build && aws-vault exec my-user --no-session -- sam local invoke GetFunction

invoke-put:
	sam build && aws-vault exec my-user --no-session -- sam local invoke PutFunction