# Provider Mock

## This is repo to mock provider response, right now this is just barebone and totally static but it has potential.

### Development

1. Use `go1.22.0`
2. Develop as usual, create each package for each provider if needed
3. Merge or push to `main` will trigger build and push to ECR

### Deploying

1. Ask Fijar/DevOps to ssh to `opensearch-staging` EC2 instance
2. Go to `~/provider_mock`
3. Run `docker compose pull`
4. Run `docker compose up -d`


## TODO
- [ ] Full CI/CD automation
- [ ] Make dynamic request, response, url

