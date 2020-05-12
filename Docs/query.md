# Query

All GraphQL queries to the server needs separate query parameters besides the graphQL schema.

Consider the following graphQL request to retrieve an item based on color:

```graphql
query ItemExample {
  item {
    query(query: "color: \"blue\" ") {
      items {
        Name
        HHAConcepts
        Variants {
          Colors
        }
      }
    }
  }
}
```

## Syntax

In the example above, we can see a line: `query(query: "color: \"blue\" ")`.

`query:` signifies that whatever comes after it is a parameter to set the variable. All variables in graphql need to be enclosed in quotes, hence the
outer quotes.

This graphql server specifically allows parsing, so you need to add another set of escaped quotes around filter requirements.

For example: `"color: \"this needs escaped quotes\" "`

The outer non-escaped quotes belong to the `query:` parameter while the inner escaped quotes (`\"`) belong to the `color:` parameter.

## Filtering

`color: \"blue\"` specifically tells the server to filter all items to have the **blue** color tag.

We can also ask for more specific filters such as:

```graphql
query ItemExample2 {
  item {
    query(query: "color: \"(blue AND white) OR (brown AND beige)\" ") {
      items {
        Name
        HHAConcepts
        Variants {
          Colors
        }
      }
    }
  }
}
```

The query above will filter items to have either `blue and white` or `brown and beige` colors.

We can also combine filters...

```graphql
query ItemExample3 {
  item {
    query(query: "color: \"(blue AND white) OR (brown AND beige)\" buy:\"2000\" ") {
      items {
        Name
        HHAConcepts
        Buy
        Variants {
          Colors
        }
      }
    }
  }
}
```

This query will filter items which are either `blue and white` **or** `brown and beige` which have a buy price of 2000 bells.

For a more in-depth list on filter options, read the doc on filters.

## Operators

The server is able to parse the following operators:

- `AND`
- `OR`
- `<`
- `>`
- `<=`
- `>=`

**AND** and **OR** are logical operators, and **<**, **>**, **<=**, and **>=** are used for number comparisons.

Any field that is set without an operator is assumed to be **=**.

For example, `buy:\"2000\"` will assume you're looking for items **equal to** 2000 bells.

The in-built parser parses commands from left to right.

Consider the following query:

`query: "color: \"blue OR white AND purple\"  buy: \">= 6000\"  "`

The server will interpret that to find items that are either **blue** or **white and purple** which has a buy price of 6000 bells and over.

To avoid ambiguous parsing, you can enclose options in parentheses.

For example, `query: "color: \"(blue OR white) AND purple\"  buy: \">= 6000\"  "`, will filter items that are either **blue and purple** or
**white and purple**  and are greater or equal to 6000 bells in buy price.

## Optionals

The server recognizes the following optional parameters:

- `glob`
- `limit`
- `month`

### Glob

**Glob** will return results that looks for **LIKE** results in the `name` filter parameter.

For example,

```graphql
query ItemGlobExample {
  item {
    query(query: "name: \"coffee\"", glob: "t") {
      items {
        Name
      }
    }
  }
}
```

will return this response:

```json
{
  "data": {
    "item": {
      "query": {
        "items": [
          {
            "Name": "coffee cup"
          },
          {
            "Name": "coffee grinder"
          }
        ]
      }
    }
  }
}
```

**Glob's** default value is always set to _f (false)_. Please note that **glob** only works for `name` filters.

### Limit

**Limit** limits how many max entries the response will return. This is particularly helpful for top level filtering.

_**NOTE:**_ I do not recommend using `limit` on tables with multiple variation entries such as the `item`, `clothing`, `photos`, and `tools` tables since limit does not ignore duplicate entries in the database.

**Limit** has a default value of 500 except on the Item table where the default is 1000.

For example,

```graphql
query ItemGlobLimitExample {
  item {
    query(query: "name: \"coffee\"", glob: "t", limit: 1) {
      items {
        Name
      }
    }
  }
}
```

will return the following response:

```json
{
  "data": {
    "item": {
      "query": {
        "items": [
          {
            "Name": "coffee cup"
          }
        ]
      }
    }
  }
}
```

### Month

**Month** is a special optional parameter that only shows up on **bug** and **fish** table querying. This parameter
is mainly used to specify the month which to return the available times the bug or fish is available.

The default value will return all times for the particular bug or fish entry.

`month` values should be shortened to three character abbreviations. For example, **January** should be represented as **Jan**.

The following query,

```graphql
query DemoFish {
  fishes {
    query(query: "name: \"arapaima\"", month: "Jun") {
      fishes {
        Name
        MonthTime {
          Month
          Time
        }
        Colors
        CatchUp
        Sell
        Shadow
        Image
        Location
        TotalCatches
      }
    }
  }
}
```

will return the following response:

```json
{
  "data": {
    "fishes": {
      "query": {
        "fishes": [
          {
            "CatchUp": "yes",
            "Colors": [
              "black",
              "blue"
            ],
            "Image": "https://acnhcdn.com/latest/MenuIcon/Fish36.png",
            "Location": "river",
            "MonthTime": [
              {
                "Month": "NH Jun",
                "Time": "4 PM to 9 AM"
              },
              {
                "Month": "SH Jun",
                "Time": "NA"
              }
            ],
            "Name": "arapaima",
            "Sell": 10000,
            "Shadow": "xx-large",
            "TotalCatches": 50
          }
        ]
      }
    }
  }
}
```
