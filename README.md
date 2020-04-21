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

I cut out portions of JSON data returned in `search_all_by_hemisphere` because the file was too long. The API for the search, nevertheless, works.

## Search Example Queries

**Success Query:**

```graphql
{
  bug_and_fish {
    search_fish_north(item: "oarfish") {
      Name
      Sell
      Where
      StartTime
      EndTime
      Shadow
    }
  }
}
```

### Search ALL By Hemisphere

```graphql
{
  bug_and_fish {
    search_all_by_hemisphere(hemi: "north") {
      bugs {
        bug {
          Name
          InternalID
        }
      }
      fishes {
        fish {
          Name
          Sell
          Shadow
        }
      }
    }
  }
}
```

**cURL Examples (bash):**

`curl -i -H "Content-Type: application/json" -X POST -d '{"query": "query {bug_and_fish{search_fish_north(item: \"oarfish\") {Name, Sell, Where, Shadow}}}"}' http://localhost:4000/`

`curl -i -H "Content-Type: application/json" -X POST -d '{"query": "query {item{search_item(item: \"acoustic guitar\") {items{item{Name, Sell, PatternTitle, Variation}}}}}"}' http://localhost:4000/`

## Contributions

**All data** used in this server was taken from open source New Horizons data found [here](https://docs.google.com/spreadsheets/d/13d_LAJPlxMa_DubPTuirkIV4DERBMXbrWQsmSh8ReK4/).
