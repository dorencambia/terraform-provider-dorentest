function upload_gpg_public_key() {
    public_key=$(gpg --armor --export "$KEY_ID_OR_EMAIL") # get key from gpg
    public_key=${public_key//$'\n'/\\n}                   # replace newlines with \n char

    data="{\"data\":{\"type\":\"gpg-keys\",\"attributes\":{\"namespace\":\"$TF_ORG\",\"ascii-armor\":\"$public_key\"}}}"

    curl \
        -Ss \
        --header "Authorization: Bearer $TF_TOKEN" \
        --header "Content-Type: application/vnd.api+json" \
        --request POST \
        --data "$data" \
        https://app.terraform.io/api/registry/private/v2/gpg-keys
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

# ensure these variables are set
require_vars TF_TOKEN TF_ORG
# run the logic
upload_gpg_public_key
