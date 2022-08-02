# create_provider only needs to be run once so versions can be created (per provider, NOT version)
create_provider() {
    curl \
        -Ss \
        --header "Authorization: Bearer $TF_TOKEN" \
        --header "Content-Type: application/vnd.api+json" \
        --request POST \
        --data "{\"data\":{\"type\":\"registry-providers\",\"attributes\":{\"name\":\"${PROVIDER_NAME}\",\"namespace\":\"${TF_ORG}\",\"registry-name\":\"private\"}}}" \
        $TERRAFORM_REGISTRY_URL
}

# print an error and exit if any passed vars are not defined
# check every passed arg to see if a variable is set with its name.
# print all missing variables and exit if any were missing
# example use: `require_vars VAR1 VAR2 var3`
function require_vars() {
    missing=0
    for required_var in "$@"; do
        if [[ -z ${!required_var} ]]; then
            echo "$required_var is required"
            missing=1
        fi
    done
    if [[ missing -eq 1 ]]; then
        exit 1
    fi
}

require_vars TF_TOKEN TF_ORG PROVIDER_NAME TERRAFORM_REGISTRY_URL
create_provider
