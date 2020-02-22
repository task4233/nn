#!/bin/sh

# set root directory by 1st argument
PKGROOT=$1
if [ "$1" = "" ]; then
    PKGROOT="."
fi

PKGNAME=$(go list $PKGROOT/...)

REVIEWDOG_ARG="-reporter='github-pr-review'"
if [ "$CI_PULL_REQUEST" = "" ]; then
    REVIEWDOG_ARG="-diff='git diff master'"
fi

golint $PKGNAME \
    | eval reviewdog -f=golint $REVIEWDOG_ARG

gsc $PKGNAME \
    | eval reviewdog -f=golint -name="gsc" $REVIEWDOG_ARG

staticcheck $PKGNAME \
    | eval reviewdog -f=golint -name="staticcheck" $REVIEWDOG_ARG


