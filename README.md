# dnd-shop-generator
A golang application that will use weights to generate a list of foundryvtt items

```bash
go run main.go --toyaml=food.txt
```

copy/paste contents to the food.yaml file (might want to drop the weights down from 100)

```bash
go run main.go --yaml=food.yaml --num=30
```