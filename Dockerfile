# image for compiling binary
ARG BUILDER_IMAGE="golang:1.22.0"
# here we'll run binary app
ARG RUNNER_IMAGE="debian:bookworm-slim"

# Github access token to fetch private modules
# ARG GITHUB_TOKEN

FROM --platform=linux/amd64 ${BUILDER_IMAGE} as builder

RUN apt-get update && apt-get install -y \
    gcc \
    libc6 \
    git \
    libc-dev \
    librdkafka-dev

# configure git to work with private repos
# RUN echo -e "machine gitlab.com\nlogin gitlab-ci-token\npassword ${GITHUB_TOKEN}" > ~/.netrc

### copying project files
WORKDIR /app

# copy gomod 
COPY go.mod go.sum ./

# COPY the source code as the last step

COPY . .

ENV CGO_ENABLED=1

# creates build/main files
RUN make build

FROM --platform=linux/amd64 ${RUNNER_IMAGE}

EXPOSE 8000
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

# create config directory
RUN mkdir -p /etc/provider_mock

COPY --from=builder /app/build/provider_mock .


CMD ["./provider_mock"]
