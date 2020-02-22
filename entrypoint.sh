#!/bin/sh

set -e

# -------------
# Environments
# -------------

RUN=$1
PKGROOT=$2
SEND_COMMENT=$3
COMMENT=""
SUCCESS=0

# if not set, assign default value
if [ "$2" = "" ]; then
    PKGROOT="."
fi
if [ "$3" = "" ]; then
    SEND_COMMENT="true"
fi

PKGNAME=$(go list $PKGROOT/...)

# ------------
# Functions
# ------------

# send_comment sends ${COMMENT} to pull request
# this function uses ${GITHUB_TOKEN}, ${COMMENT} and ${GITHUB_EVENT_PATH}
send_comment() {
    PAYLOAD=$(echo '{}' | jq --arg body "${COMMENT}" '.body = $body')
    COMMENTS_URL=$(cat ${GITHUB_EVENT_PATH} | jq -r .pull_request.comments_url)
    curl -s -S -H "Authorization: token ${GITHUB_TOKEN}" --header "Content-Type: application/json" --data "{PAYLOAD}" "${COMMENTS_URL}" > /dev/null
}

module_download() {
    # if not exist go.mod
    if [ ! -e go.mod ]; then
	go mod init
    fi

    # if finished in error status, exit 1
    go mod download
    if [ $? -ne 0 ]; then
	exit 1;
    fi
}

run_gofmt() {
    set +e

    # TODO
    # impelements workflow

    set -e

    # exit successfully
    if [ ${SUCCESS} -eq 0]; then
	return
    fi

    # TODO
    # implements applying comment
}


run_goimports() {
        set +e

    # TODO
    # impelements workflow

    set -e

    # exit successfully
    if [ ${SUCCESS} -eq 0]; then
	return
    fi

    # TODO
    # implements applying comment
    
}

# golint for checking coding style
run_golint() {
    set +e

    # TODO
    # impelements workflow

    set -e

    # exit successfully
    if [ ${SUCCESS} -eq 0]; then
	return
    fi

    # TODO
    # implements applying comment
    
}
    
# gsc for static analysis
run_gsc() {
}

# staticcheck is like a go vet
run_staticcheck() {
    set +e

    # TODO
    # impelements workflow

    set -e

    # exit successfully
    if [ ${SUCCESS} -eq 0]; then
	return
    fi

    # TODO
    # implements applying comment
}

# -------------
# Main
# ------------
case ${RUN} in
    "fmt" )
	module_download
	run_gofmt
	;;
    "imports" )
	module_download
	run_goimports
	;;
    "lint" )
	module_download
	run_golint
	;;
    "gsc" )
	module_download
	run_gsc
	;;
    "staticcheck" )
	module_download
	run_staticcheck
	;;
    * )
	echo "Invalid command." >&2
	exit 1
esac

if [ ${SUCCESS} -ne 0 ]; then
    echo "Check failed." >&2
    echo "${COMMENT}" >&2
    if [ "${SEND_COMMENT}" = "true" ]; then
	send_comment
    fi
fi

exit ${SUCCESS}
