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
**white and purple**  and are greater than 6000 bells in buy price.


