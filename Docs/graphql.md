# GraphQL

GraphQL is a runtime query language for APIs. It is designed to retrieve ONLY data that was requested by the client.

For example, consider the following GraphQL request for a specific villager called _Agnes_.

We can write a query asking for the villager _Name_:

```graphql
query VillagerExample {
  villagers {
    query(query: "name:\"Agnes\"") {
      villagers {
        Name
      }
    }
  }
}
```

And the response returned to us will look like:

```json
{
  "data": {
    "villagers": {
      "query": {
        "villagers": [
          {
            "Name": "agnes"
          }
        ]
      }
    }
  }
}
```

What if we wanted more data fields besides just the villager name?

In this case, we can just add more resource properties defined by the `schema`. All of the _schemas_ defined in this server is located in
`newhorizons`. From that folder, we can see a file called `villager_model.go` which defines the GraphQL schema for **villager** type objects.

```go
// Villager represents a villager entry graphql is able to parse and
// display
type Villager struct {
	Name        string
	Image       string
	Species     string
	Gender      string
	Personality string
	Birthday    string
	Catchphrase string
	Style       []string
	Color       []string
}
```

From the schema, we can see that there are **eight** other query fields we can use.

For example, we can query for a villager's name, species, and gender at the same time.

```graphql
query VillagerExample2 {
  villagers {
    query(query: "name:\"Agnes\"") {
      villagers {
        Name
        Species
        Gender
      }
    }
  }
}
```

And the response returned will look like:

```json
{
  "data": {
    "villagers": {
      "query": {
        "villagers": [
          {
            "Gender": "female",
            "Name": "agnes",
            "Species": "pig"
          }
        ]
      }
    }
  }
}
```

Now we can see that gender and species were also returned in the response!

Besides looking at schema definitions, you can also use `GraphiQL`, a graphQL interface for editing and testing queries.

The link to this server's GraphiQL interface is the same as the API endpoint: [https://acnhgraphql.com](https://acnhgraphql.com).

The `Docs` tab in the upper right hand corner is able to show you all schema definitions and query options for the server.

[![Image from Gyazo](https://i.gyazo.com/1e66cd80b5c2f693a41f5ec786615ad7.png)](https://gyazo.com/1e66cd80b5c2f693a41f5ec786615ad7)

That's it for a general introduction to GraphQL!

If you're curious on some of the other cooler features of GraphQL, make sure to check out their official [website](https://graphql.org/)!
