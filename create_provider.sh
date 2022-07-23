# create_provider only needs to be run once so versions can be created (per provider, NOT version)
create_provider() {
    curl \
        -Ss \
        --header "Authorization: Bearer $TF_TOKEN" \
        --header "Content-Type: application/vnd.api+json" \
        --request POST \
        --data "{\"data\":{\"type\":\"registry-providers\",\"attributes\":{\"name\":\"${PROVIDER_NAME}\",\"namespace\":\"${TF_ORG}\",\"registry-name\":\"private\"}}}" \
        $BASE_URL
}

# run the function
create_provider
