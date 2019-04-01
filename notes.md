# notes

project console:

    https://console.cloud.google.com/home/dashboard?folder=&project=poc-ns-api

windows:

    $env:GOOGLE_APPLICATION_CREDENTIALS="secrets/poc-ns-api-02835f9471c9.json"
    go run cmd/addstory/addstory.go
    go run cmd/addstories/addstories.go path/to/stories.json

bash:

    export GOOGLE_APPLICATION_CREDENTIALS="secrets/poc-ns-api-02835f9471c9.json"
    go run cmd/addstory/addstory.go