# run `source ./set_env_vars.sh` to set these variables in your environment

export TF_ORG="cambia-demo" # name of the organization in terraform cloud
export PROVIDER_NAME="dorentest" # name of the provider being published
export KEY_ID_OR_EMAIL="doren.proctor@cambiahealth.com" # what was used when the gpg key was creatd
export BASE_URL="https://app.terraform.io/api/v2/organizations/$TF_ORG/registry-providers" # url to private registry