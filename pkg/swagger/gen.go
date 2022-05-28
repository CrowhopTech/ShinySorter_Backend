package swagger

//go:generate rm -rf server client
//go:generate mkdir -p server client
//go:generate swagger generate server --quiet --target server --name shiny-sorter --spec swagger.yaml --exclude-main
//go:generate swagger generate client --quiet --name shiny-sorter --spec swagger.yaml
