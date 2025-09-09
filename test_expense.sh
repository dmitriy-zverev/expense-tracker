# !/bin/zsh

rm ./data/expenses.json && touch ./data/expenses.json
go run . add --description "item 1" --amount 100 --category "cat 1"
go run . add --description "item 2" --amount 200 --category "cat 2"
go run . add --description "item 3" --amount 5000 --category "cat 3"
go run . list