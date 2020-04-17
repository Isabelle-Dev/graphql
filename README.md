# Animal Crossing - New Horizons GraphQL Server

Server is set to run on `localhost:4000`.

## Search

### Search By Item (Bug or Fish)

**Query:**

```graphql
{
  search {
    search_bug_or_fish(item: "ladybug", tablename: "bug_and_fish") {
      Name
      Time
      Location
      SellPrice
    }
  }
}
```

**Response:**

```json
{
  "data": {
    "search": {
      "search_bug_or_fish": {
        "Location": "On_flowers",
        "Name": "ladybug",
        "SellPrice": 200,
        "Time": "8_AM_to_5_PM"
      }
    }
  }
}
```
