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

# run_reviewdog executes golint, gsc, staticcheck by reviewdog
# this function sends comment to pull request if it has errors
run_reviewdog() {
    set +e
    
    # golint for checking coding style
    golint $PKGNAME \
	| eval reviewdog -f=golint $REVIEWDOG_ARG
    SUCCESS=$?

    if [ ${SUCCESS} -ne 0]; then
	
	if [ "${SEND_COMMENT}" = "true" ]; then
	    COMMENT="golint failed
\`\`\`
${OUTPUT}				
\`\`\`
"	
	fi
    fi
    
    # gsc for static analysis
    gsc $PKGNAME \
	| eval reviewdog -f=golint -name="gsc" $REVIEWDOG_ARG
    
    
    
    # staticcheck is like a go vet
    staticcheck $PKGNAME \
	| eval reviewdog -f=golint -name="staticcheck" $REVIEWDOG_ARG

    set -e

    
}

# -------------
# Main
# ------------
case ${RUN} in
    "errcheck" )
	module_download
	exec_errcheck
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
