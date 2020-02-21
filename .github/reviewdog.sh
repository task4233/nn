#!/bin/sh

# set root directory by 1st argument
PKGROOT=$1
if [ "$1" = "" ]; then
    PKGROOT="."
fi

REVIEWDOG_ARG="-reporter='github-pr-review'"
if [ "$CI_PULL_REQUEST" = "" ]; then
    REVIEWDOG_ARG="-diff='git diff master'"
fi

golint $(go list $PKGROOT/...) | eval reviewdog -f=golint $REVIEWDOG_ARG

gsc -tests=false \
