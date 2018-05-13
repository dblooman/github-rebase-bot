#!/bin/sh

/github-rebase-bot \
 -repos "$GITHUB_REPOS" \
 -public-dns https://$PUBLIC_DNS \
 -merge-label "$GITHUB_MERGE_LABEL" \
 -addr :8080
