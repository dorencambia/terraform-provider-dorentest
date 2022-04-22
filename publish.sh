set -e # exit if any line fails

# change these
provider_name="dorentest"
tf_org="cambia-demo"
base_url="https://app.terraform.io/api/v2/organizations/$tf_org/registry-providers"

# get the tag of the current commit
version=`git tag --points-at`
version=${version#v} # remove the v at the beginning

# list of all generated zip files without the .zip suffix
zip_file_names="darwin_amd64 darwin_arm64 freebsd_386 freebsd_amd64 freebsd_arm freebsd_arm64 linux_386 linux_amd64 linux_arm linux_arm64 windows_386 windows_amd64 windows_arm windows_arm64"

# download all the files that GoReleaser created in GitHub releases
download_release_files() {
    base_download_url="https://github.com/dorencambia/terraform-provider-${provider_name}/releases/download"
    files="SHA256SUMS SHA256SUMS.sig"
    for zip_file in $zip_file_names; do
        files+=" $zip_file.zip"
    done
    for file in $files; do
        filename="terraform-provider-${provider_name}_${version}_${file}"
        curl -SsLO ${base_download_url}/v${version}/$filename
    done
}

# create_provider only needs to be run once so versions can be created (per provider, NOT version)
create_provider() {
    curl \
        -Ss \
        --header "Authorization: Bearer $TF_TOKEN" \
        --header "Content-Type: application/vnd.api+json" \
        --request POST \
        --data "{\"data\":{\"type\":\"registry-providers\",\"attributes\":{\"name\":\"${provider_name}\",\"namespace\":\"${tf_org}\",\"registry-name\":\"private\"}}}" \
        $base_url
}

# create the provider version and then upload the 2 SHA files
upload_sha_files() {
    create_provider_version() {
        curl \
        -Ss \
        --header "Authorization: Bearer $TF_TOKEN" \
        --header "Content-Type: application/vnd.api+json" \
        --request POST \
        --data "{\"data\":{\"type\":\"registry-provider-versions\",\"attributes\":{\"version\":\"$version\",\"key-id\":\"$TF_TOKEN\",\"protocols\":[\"5.0\"]}}}" \
        $base_url/private/$tf_org/$provider_name/versions
    }
    create_provider_version_resp=`create_provider_version`
    shasums_upload=`echo $create_provider_version_resp | jq -r '.data.links["shasums-upload"]'`
    shasums_sig_upload=`echo $create_provider_version_resp | jq -r '.data.links["shasums-sig-upload"]'`
    curl -Ss --upload-file ./terraform-provider-${provider_name}_${version}_SHA256SUMS $shasums_upload
    curl -Ss --upload-file ./terraform-provider-${provider_name}_${version}_SHA256SUMS.sig $shasums_sig_upload
}


# create provider platform for each type then upload the zipped binary along with its sha
upload_zip_files() {
    create_provider_platform() {
        os=$1
        arch=$2
        filename=$3
        sha_file="terraform-provider-${provider_name}_${version}_SHA256SUMS"
        shasum=`grep $filename $sha_file | cut -d ' ' -f1` # get sha from downloaded SHA256SUMS release file
        curl \
            -Ss \
            --header "Authorization: Bearer $TF_TOKEN" \
            --header "Content-Type: application/vnd.api+json" \
            --request POST \
            --data "{\"data\":{\"type\":\"registry-provider-version-platforms\",\"attributes\":{\"os\":\"${os}\",\"arch\":\"${arch}\",\"shasum\":\"${shasum}\",\"filename\":\"${filename}\"}}}" \
            $base_url/private/$tf_org/$provider_name/versions/$version/platforms
    }

    for name in $zip_file_names; do
        os=`echo $name | cut -d '_' -f1`
        arch=`echo $name | cut -d '_' -f2`
        filename="terraform-provider-${provider_name}_${version}_${os}_${arch}.zip"

        # create platform to upload the file
        resp=`create_provider_platform $os $arch $filename`
        # get url from reponse
        provider_binary_upload_url=`echo $resp | jq -r '.data.links["provider-binary-upload"]'`
        # upload the file
        curl -Ss -T $filename $provider_binary_upload_url
    done
}

_publish() {
    download_release_files
    upload_sha_files
    upload_zip_files
}

_publish