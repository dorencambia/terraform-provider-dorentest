# run `source ./set_env_vars.sh` to set these variables in your environment

# name of the organization in terraform cloud
export TF_ORG="my-terraform-org"

# name of the provider being published
export PROVIDER_NAME="my-private-provider"

# what was used when the gpg key was created
export KEY_ID_OR_EMAIL="email-or-id-used-when-creating-gpg-key@example.com"

# url to private registry
export TERRAFORM_REGISTRY_URL="https://app.terraform.io/api/v2/organizations/$TF_ORG/registry-providers"

# url to the github repo hosting the provider
export GITHUB_REPO_URL="https://github.com/my-github-user/my-private-terraform-provider-repo"