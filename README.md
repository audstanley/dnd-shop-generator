# dnd-shop-generator
A golang application that will use weights to generate a list of foundryvtt items

### Prep
install golang
```bash
# build the binary from within the project folder
go build .
```


### Step 1
copy uuids from a foundry journal into a text file.

### Step 2

```bash
./dnd-shop-generator --toyaml=food.txt
```

### Step 3
copy/paste contents from generated_output_items_from_textfile.yaml to a food.yaml file

### Step 3 (optional)
(might want to drop the weights down from 100)

### Step 4

```bash
./dnd-shop-generator --yaml=food.yaml --num=30
```

output to put into a foundry journal:

```
@UUID[Item.jWAjwLnXrKQ2LHml]{Bearberry}
@UUID[Item.ONFpSSLQwkfliem9]{Beet}
@UUID[Item.efnKO9MmiYFRzz46]{Bitterroot}
@UUID[Item.cLz727ZNuHLWtwfQ]{Broccoli}
@UUID[Item.kOL73LeUfa1y0z49]{Cauliflower}
@UUID[Item.4mtkqtrD1nOd84HG]{Cheese}
@UUID[Item.CPSHd4OTFAIXyBA4]{Cherry}
@UUID[Item.TVsfbYaDn6axKTL9]{Date Fruit}
@UUID[Item.pYF4acOH5msSFEzb]{Death Cap}
@UUID[Item.Fm63zTZTn2GwiCGu]{Dragon Fruit}
@UUID[Item.fSN6ifxNHPNEwECK]{Durian}
@UUID[Item.HMKPQdQdJg7ViN9a]{Grapes}
@UUID[Item.15YFKzLbgsdO0rIN]{Green Apple}
@UUID[Item.pKZqgl4gdnHRj53t]{Green Grapes}
@UUID[Item.eIZ1kW1LDoiITEEV]{Green Pepper}
@UUID[Item.Bdvq0nX4Hvx4oyiF]{Kiwifruit}
@UUID[Item.aINSxAHatji8iUBQ]{Olives}
@UUID[Item.OZ3aMZMB5u1qanvJ]{Onion}
@UUID[Item.XBtyWgiGynUYQUfm]{Peanut}
@UUID[Item.MrpDHntVyOOKIu4v]{Pear}
@UUID[Item.0EqA9ji9s11WyjbZ]{Plum}
@UUID[Item.0hYy2ESxYPqXtPEC]{Raspberry}
@UUID[Item.Gw9IV9TgZ5GOr5IQ]{Spinach}
@UUID[Item.dcEdOPTB8MBVHpnC]{Sunberries}
```