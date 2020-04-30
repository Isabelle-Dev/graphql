# Animal Crossing - New Horizons GraphQL Server

Server and **endpoint** is available at [acnhgraphql.com](acnhgraphql.com).

**I'm still setting up Linux server, so not all features are available yet!**

_I will be writing more detailed guides on how to use the API after I setup the server. Thanks for being patient!_

## Server Status

- [x] Item
- [x] Wallpaper
- [x] Floor
- [x] Clothes
- [ ] Music
- [ ] Photos
- [ ] Posters
- [ ] Rug
- [ ] Tools
- [ ] Villager

## Self-Hosting Installation

Want to host the server yourself? You can!

- `go get -u github.com/Isabelle-Dev/graphql`
- Setup PostgreSQL
- Create and configure `.config` file using `example.config` as template
- Import data using `csv` files
- `go build -o graphql.exe`
- `./graphql.exe`

## CSV Data

All **csv data files** and **import code to Postgres** are located in `csv`.

## Sample JSON Responses

**Sample JSON** responses can be found in `newhorizons/sample`.

## Example Queries

```graphql
query FloorDemo {
  floor {
    query(query: "buy:\"<= 3000 AND > 2000\" color:\"gray AND beige\"") {
      floors {
        Color
        Catalog
        Concepts
        Sell
        SourceNotes
        Tag
        Image
        VFX
        Buy
        Name
      }
    }
  }
}
```

```graphql
query DemoItem {
  item {
    query(query: " name:\"leaf\" tag:\"plant\" color:\"orange\" " glob:"t" limit: 30) {
      items {
        Name
        Buy
        Sell
        HHAConcepts
        HHASet
        HHASeries
        Variants {
          Colors
        }
        Tag
      }
    }
  }
}
```

```graphql
query DemoItemV2 {
  item {
    query(query: "sell:\"< 2000 AND > 1500\" color:\"(pink AND green) OR orange\"") {
      items {
        Name
        Variants {
          Colors
          Img
          Pattern
        }
        Buy
        Sell
        Source
        KitCost
        HHASet
        HHASeries
        HHAConcepts
        Interact
        PatternTitle
        PatternCustomize
      }
    }
  }
}
```

**cURL Example:**

An updated **cURL** query can be found in `post.json`. I do not recommend making graphql requests using cURL, but if you must, it's easier to port requests using an external json file.

`curl -H "Content-type:application/json" --data @post.json https://acnhgraphql.com`

## Table Migrations

### Item Table (item)

**Import Script:** `importItem.psql`

- Housewares - `housewares_clean.csv`
- Misc - `misc_clean.csv`
- Wall-hanging - `wall-hanging_clean.csv`

### Wallpaper Table (wallpaper)

**Import Script:** `importWallpaper.psql`

- Wallpapers - `wallpaper_clean.csv`

### Floor Table (floor)

**Import Script:** `importFloor.psql`

- Floors - `floor_clean.csv`

### Clothing Table (clothes)

**Import Script:** `importClothes.psql`

- Accessories - `accessories_clean.csv`
- Bags - `bags_clean.csv`
- Bottoms - `bottoms_clean.csv`
- Tops - `tops_clean.csv`
- Dresses - `dresses_clean.csv`
- Headwear - `headwear_clean.csv`
- Shoes - `shoes_clean.csv`
- Socks - `socks_clean.csv`

## Contributions

- All Data is sourced from open-source New Horizons data found [here](https://docs.google.com/spreadsheets/d/13d_LAJPlxMa_DubPTuirkIV4DERBMXbrWQsmSh8ReK4/)

- High-resolution image links are provided by [Dodo Codes](https://acnhcdn.com/)
