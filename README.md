# cut-tool

> This is a simple version of the Unix tool `cut`

`go run main.go -f 2 sample.tsv`

`go run main.go -f 1 -d , fourchords.csv | head -n5`

`go run main.go -f '1 2' sample.tsv`

`go run main.go -f 1,2 sample.tsv`

`go run main.go -d , -f "1 2" fourchords.csv | head -n5`

`tail -n5 fourchords.csv | go run main.go -d , -f 1,2`

`tail -n5 fourchords.csv | go run main.go -d , -f 1,2 -`

`go run main.go -f 2 -d , fourchords.csv | uniq | wc -l`