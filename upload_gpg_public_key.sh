if [[ -z "$TF_TOKEN" ]]; then
    echo "TF_TOKEN is required"
    exit 1
fi

tf_org="cambia-demo" # name of organization in terraform cloud
key_id_or_email="doren.proctor@cambiahealth.com" # what was used when the gpg key was creatd
public_key=`gpg --armor --export "$key_id_or_email"` # get key from gpg
public_key=${public_key//$'\n'/\\n} # replace newlines with \n char

data="{\"data\":{\"type\":\"gpg-keys\",\"attributes\":{\"namespace\":\"$tf_org\",\"ascii-armor\":\"$public_key\"}}}"

curl \
    -Ss \
    --header "Authorization: Bearer $TF_TOKEN" \
    --header "Content-Type: application/vnd.api+json" \
    --request POST \
    --data "$data" \
    https://app.terraform.io/api/registry/private/v2/gpg-keys
