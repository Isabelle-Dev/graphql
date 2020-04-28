# Animal Crossing - New Horizons GraphQL Server

Server is set to run on `localhost:4000`.

## Installation

* Setup PostgreSQL
* Create and configure `.config` file using `example.config` as template
* Import data using `csv` files
* `go build -o isabelle-graphql.exe`
* `./isabelle-graphql.exe`

## CSV Data

All **csv data files** and **import code to Postgres** will be located in `csv`.

## Sample JSON Responses

**Sample JSON** responses can be found in `newhorizons/sample`.

## Search Example Queries

```graphql
query FloorDemo {
  floor {
    query(query: "buy:\"<= 3000 AND > 2000\" color:\"Gray AND Beige\"") {
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
query demoItem {
  item {
    query(query: "sell:\"< 2000 AND > 1500\" color:\"(Pink AND Green) OR Orange\"") {
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

```graphql
query DemoItem {
  item {
    query(query: " name:\"leaf\" tag:\"Plant\" color:\"Orange\" " glob:"t" limit: 30) {
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

**cURL Examples (bash):**

An updated **cURL** query can be found in `post.json`. I do not recommend making graphql requests using cURL, but if you must, it's easier to port requests using an external json file.

For example: `curl -H "Content-type:application/json" --data @post.json http://localhost:4000/`

## Contributions

* All Data is sourced from open-source New Horizons data found [here](https://docs.google.com/spreadsheets/d/13d_LAJPlxMa_DubPTuirkIV4DERBMXbrWQsmSh8ReK4/)

* High-resolution image links are provided by [Dodo Codes](https://acnhcdn.com/)
