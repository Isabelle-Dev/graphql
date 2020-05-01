# Filters

Each query table has filter options based on database rows. This document will list all the different filter options for each table used
in the graphql server.

The following filters must be paired with an integer number:

- `kitcost`
- `buy`
- `sell`
- `uses`
- `month`

Everything else uses regular character strings.

`tag`, `theme`, and `name` (with the glob parameter set to true) will apply glob filters. All other parameters must have values
matching an existing value in the database.

For example, setting color to `avocado green` will result in the server responding `entries not found` since there are no `avocado green`
color codes in the database.

## Contents

- [Item](#item)
- [Wallpaper](#wallpaper)
- [Floor](#floor)
- [Clothes](#clothes)
- [Music](#music)
- [Photos](#photos)
- [Posters](#posters)
- [Rugs](#rugs)
- [Tools](#tools)
- [Villagers](#villagers)


### Item

- `name` - Name
- `variation` - variation color
- `kitcost` - kit cost
- `buy` - buy price in bells
- `sell` - sell price in bells
- `color` - object color codes
- `concept` - HHA Concept
- `tag` - object tag
- `catalog` - For sale or Not for sale

### Wallpaper

- `name` - Name
- `buy` - buy price in bells
- `sell` - sell price in bells
- `color` - object color codes
- `concept` - HHA Concept
- `tag` - object tag
- `catalog` - For sale or Not for sale

### Floor

- `name` - Name
- `buy` - buy price in bells
- `sell` - sell price in bells
- `color` - object color codes
- `concept` - HHA Concept
- `tag` - object tag
- `catalog` - For sale or Not for sale

### Clothes

- `name` - Name
- `variation` - variation color
- `diy` - diy status
- `buy` - buy price in bells
- `sell` - sell price in bells
- `color` - object color codes
- `style` - clothing style
- `theme` - label theme
- `catalog` - For sale or Not for sale

### Music

- `name` - Name
- `buy` - buy price in bells
- `sell` - sell price in bells
- `color` - object color codes
- `catalog` - For sale or Not for sale

### Photos

- `name` - Name
- `variation` - variation color

### Posters

- `name` - Name
- `color` - object color codes

### Rugs

- `name` - Name
- `buy` - buy price in bells
- `sell` - sell price in bells
- `color` - object color codes
- `catalog` - For sale or Not for sale
- `concept` - HHA Concept
- `tag` - object tag

### Tools

- `name` - Name
- `variation` - variation color
- `kitcost` - kit cost
- `uses` - number of usage
- `buy` - buy price in bells
- `sell` - sell price in bells

### Villagers

- `name` - Name
- `species` - villager species
- `gender` - villager gender
- `personality` - villager personality
- `style` - villager dressing style
- `color` - villager color codes
- `month` - villager birthday month
