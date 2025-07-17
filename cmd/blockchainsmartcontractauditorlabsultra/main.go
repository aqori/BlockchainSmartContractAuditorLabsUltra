// cmd/blockchainsmartcontractauditorlabsultra/main.go
package main

import (
"flag"
"log"
"os"

"blockchainsmartcontractauditorlabsultra/internal/blockchainsmartcontractauditorlabsultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainsmartcontractauditorlabsultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
