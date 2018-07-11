FROM golang:1.10.2-alpine3.7
RUN apk --no-cache --update add git
COPY . /go/src/github.com/depop/github-rebase-bot
RUN cd /go/src/github.com/depop/github-rebase-bot && go install

FROM alpine:3.7

RUN apk --no-cache --update add ca-certificates git curl jq && update-ca-certificates

ENV GITHUB_TOKEN="" \
    GITHUB_OWNER="" \
    GITHUB_REPOS="" \
    GITHUB_MERGE_LABEL="LGTM" \
    PUBLIC_DNS=""

COPY --from=0 /go/bin/github-rebase-bot /
ADD startup.sh /

ENTRYPOINT ["/startup.sh"]
