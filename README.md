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

```json
{
  "data": {
    "search": null
  },
  "errors": [
    {
      "message": "search(): sunset moth not found",
      "locations": [
        {
          "line": 3,
          "column": 5
        }
      ],
      "path": [
        "search",
        "search_bug_north"
      ]
    }
  ]
}
```

**Success Query:**

```graphql
{
  bug_and_fish {
    search_bug_north(item:"ladybug"){
      Name
      Sell
    }
  }
}
```

**Response:**

```json
{
  "data": {
    "bug_and_fish": {
      "search_bug_north": {
        "Name": "ladybug",
        "Sell": 200
      }
    }
  }
}
```

**cURL Example (bash):**

`curl -i -H "Content-Type: application/json" -X POST -d '{"query": "query {search{search_bug_north(item: \"cricket\") {Sell,Where}}}"}' http://localhost:4000/`
