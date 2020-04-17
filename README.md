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
  search {
    search_bug_north(item: "common butterfly") {
      Name
      Sell
      Image
      StartTime
    }
  }
}
```

**Response:**

```json
{
  "data": {
    "search": {
      "search_bug_north": {
        "Image": "https://i.imgur.com/UJiH4E9.png",
        "Name": "common butterfly",
        "Sell": 160,
        "StartTime": "4:00 AM"
      }
    }
  }
}
```
