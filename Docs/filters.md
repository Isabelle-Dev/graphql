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
- [Art](#art)
- [Fencing](#fencing)
- [Construction](#construction)
- [Nook Miles](#nook-miles)
- [Umbrella](#umbrella)
- [Other](#other)
- [Bug](#bug)
- [Fish](#fish)
- [Fossil](#fossil)
- [Reactions](#reactions)

### Item

- `name` - Name
- `variation` - variation color
- `kitcost` - kit cost
- `buy` - buy price in bells
- `sell` - sell price in bells
- `color` - object color codes
- `concept` - HHA Concept
- `tag` - object tag
- `catalog` - fs (for sale) or nfs (not for sale)

### Wallpaper

- `name` - Name
- `buy` - buy price in bells
- `sell` - sell price in bells
- `color` - object color codes
- `concept` - HHA Concept
- `tag` - object tag
- `catalog` - fs (for sale) or nfs (not for sale)

### Floor

- `name` - Name
- `buy` - buy price in bells
- `sell` - sell price in bells
- `color` - object color codes
- `concept` - HHA Concept
- `tag` - object tag
- `catalog` - fs (for sale) or nfs (not for sale)

### Clothes

- `name` - Name
- `variation` - variation color
- `diy` - diy status
- `buy` - buy price in bells
- `sell` - sell price in bells
- `color` - object color codes
- `style` - clothing style
- `theme` - label theme
- `catalog` - fs (for sale) or nfs (not for sale)

### Music

- `name` - Name
- `buy` - buy price in bells
- `sell` - sell price in bells
- `color` - object color codes
- `catalog` - fs (for sale) or nfs (not for sale)

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
- `catalog` - **fs** (For sale) or **nfs** (Not for sale)
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
- `hobby` - villager's hobby
- `style` - villager dressing style
- `color` - villager color codes
- `month` - villager birthday month

### Art

- `name` - Name
- `tag` - object tag

### Fencing

- `name` - Name
- `buy` - buy price in bells
- `sell` - sell price in bells

### Construction

- `name` - Name
- `buy` - buy price in bells
- `category` - category type (i.e. bridge, door, mailbox, etc)

### Nook Miles

- `name` - Name
- `nookmiles` - nook miles amount

### Umbrella

- `name` - Name
- `buy` - buy price in bells
- `sell` - sell price in bells
- `color` - colors of item
- `catalog` - **fs** or **nfs** for sale status
- `diy` - **yes** or **no** for diy status

### Other

- `name` - Name
- `buy` - buy price in bells
- `sell` - sell price in bells
- `tag` - item tag

### Bug

- `name` - Name
- `sell` - sell price in bells
- `rarity` - rarity status
- `color` - color of bug

**Special: Month Status**

- `NHJan` ~ `NHDec` - `yes` or `no`
- `SHJan` ~ `SHDec` - `yes` or `no`

### Fish

- `name` - Name
- `sell` - sell price in bells
- `rarity` - rarity status
- `color` - color of fish
- `catchup` - `yes` or `no`
- `shadow` - fish shadow size in water

**Special: Month Status**

- `NHJan` ~ `NHDec` - `yes` or `no`
- `SHJan` ~ `SHDec` - `yes` or `no`

### Fossil

- `name` - Name
- `sell` - sell price in bells
- `color` - color of fossil
- `interact` - interact status (`yes` or `no`)

### Reactions

- `name` - Name
- `source` - villager personality you get reaction from
