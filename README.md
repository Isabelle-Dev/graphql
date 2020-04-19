# Animal Crossing - New Horizons GraphQL Server

Server is set to run on `localhost:4000`.

## Installation

* Setup PostgreSQL
* Create and configure `.config` file using `example.config` as template
* Import data using `csv` files
* `go build`
* `./isabelle-graphql.exe`

## CSV Data

All **csv data files** and **import code to Postgres** will be located in `csv`.

## Sample JSON Responses

**Sample JSON** responses can be found in `newhorizons/sample`.

I cut out portions of JSON data returned in `search_all_by_hemisphere` because the file was too long. The API for the search, nevertheless, works.

## Search

### Search By Item (Bug or Fish)

**Failed Query:**

```graphql
{
  search {
    search_bug_north(item: "sunset moth") {
      Name
      Sell
      Image
      StartTime
    }
  }
}
```

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

**cURL Example (bash):**

`curl -i -H "Content-Type: application/json" -X POST -d '{"query": "query {search{search_bug_north(item: \"cricket\") {Sell,Where}}}"}' http://localhost:4000/`
