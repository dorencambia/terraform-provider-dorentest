# required environment variables
# TF_ORG: name of the organization in terraform cloud
# PROVIDER_NAME: name of the provider being published
# KEY_ID_OR_EMAIL: what was used when the gpg key was creatd
# TERRAFORM_REGISTRY_URL: url to private terraform registry
# GITHUB_REPO_URL: url to the github repo hosting the provider


set -e # exit if any line fails

version=$(git tag --points-at) # get the tag of the current commit
version=${version#v}           # remove the v at the beginning

# list of all generated zip files without the .zip suffix
# build is real fast when only building one version
# set this to match the combinations in .goreleaser.yml
zip_file_names="darwin_amd64 darwin_arm64 freebsd_386 freebsd_amd64 freebsd_arm freebsd_arm64 linux_386 linux_amd64 linux_arm linux_arm64 windows_386 windows_amd64 windows_arm windows_arm64"
# zip_file_names="linux_amd64"


# download all the files that GoReleaser created in GitHub releases
download_release_files() {
    echo "### downloading release files ###"
    base_download_url="${GITHUB_REPO_URL}/releases/download"
    files="SHA256SUMS SHA256SUMS.sig"
    for zip_file in $zip_file_names; do
        files+=" $zip_file.zip"
    done
    for file in $files; do
        filename="terraform-provider-${PROVIDER_NAME}_${version}_${file}"
        curl -SsLO ${base_download_url}/v${version}/$filename
    done
}

# create the provider version and then upload the 2 SHA files
upload_sha_files() {
    echo "### uploading sha files ###"
    create_provider_version() {
        curl \
            -Ss \
            --header "Authorization: Bearer $TF_TOKEN" \
            --header "Content-Type: application/vnd.api+json" \
            --request POST \
            --data "{\"data\":{\"type\":\"registry-provider-versions\",\"attributes\":{\"version\":\"$version\",\"key-id\":\"$GPG_KEY_ID\",\"protocols\":[\"5.0\"]}}}" \
            $TERRAFORM_REGISTRY_URL/private/$TF_ORG/$PROVIDER_NAME/versions
    }
    create_provider_version_resp=$(create_provider_version)
    shasums_upload=$(echo $create_provider_version_resp | jq -r '.data.links["shasums-upload"]')
    shasums_sig_upload=$(echo $create_provider_version_resp | jq -r '.data.links["shasums-sig-upload"]')
    curl -Ss --upload-file ./terraform-provider-${PROVIDER_NAME}_${version}_SHA256SUMS $shasums_upload
    curl -Ss --upload-file ./terraform-provider-${PROVIDER_NAME}_${version}_SHA256SUMS.sig $shasums_sig_upload
}

# create provider platform for each type then upload the zipped binary along with its sha
upload_zip_files() {
    echo "### uploading zip files ###"
    create_provider_platform() {
        os=$1
        arch=$2
        filename=$3
        sha_file="terraform-provider-${PROVIDER_NAME}_${version}_SHA256SUMS"
        shasum=$(grep $filename $sha_file | cut -d ' ' -f1) # get sha from downloaded SHA256SUMS release file
        curl \
            -Ss \
            --header "Authorization: Bearer $TF_TOKEN" \
            --header "Content-Type: application/vnd.api+json" \
            --request POST \
            --data "{\"data\":{\"type\":\"registry-provider-version-platforms\",\"attributes\":{\"os\":\"${os}\",\"arch\":\"${arch}\",\"shasum\":\"${shasum}\",\"filename\":\"${filename}\"}}}" \
            $TERRAFORM_REGISTRY_URL/private/$TF_ORG/$PROVIDER_NAME/versions/$version/platforms
    }

    for name in $zip_file_names; do
        os=$(echo $name | cut -d '_' -f1)
        arch=$(echo $name | cut -d '_' -f2)
        filename="terraform-provider-${PROVIDER_NAME}_${version}_${os}_${arch}.zip"

        # create platform to upload the file
        resp=$(create_provider_platform $os $arch $filename)
        # get url from reponse
        provider_binary_upload_url=$(echo $resp | jq -r '.data.links["provider-binary-upload"]')
        # upload the file
        curl -Ss -T $filename $provider_binary_upload_url
        echo "### uploaded file $filename ###"
    done
}

check_required_env_vars() {
    required="TF_TOKEN GPG_KEY_ID"
    for x in $required; do
        if [[ -z "${!x}" ]]; then
            echo value required for ${x}
            exit 1
        fi
    done
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

_publish() {
    require_vars TF_TOKEN GPG_KEY_ID TF_ORG PROVIDER_NAME TERRAFORM_REGISTRY_URL GITHUB_REPO_URL
    download_release_files
    upload_sha_files
    upload_zip_files
}
_publish
