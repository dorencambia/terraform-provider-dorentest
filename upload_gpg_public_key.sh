if [[ -z "$TF_TOKEN" ]]; then
    echo "TF_TOKEN is required"
    exit 1
fi

TF_ORG="cambia-demo"                                  # name of organization in terraform cloud
KEY_ID_OR_EMAIL="doren.proctor@cambiahealth.com"      # what was used when the gpg key was creatd
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
