<p align="center">
  <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/1/17/GraphQL_Logo.svg/1024px-GraphQL_Logo.svg.png" height="200" width="200">
  <h1 align="center"> ACNH - GraphQL Server </h1>
  <p align="center">An open source Animal Crossing New Horizons GraphQL API Server & Endpoint </p>
</p>

ACNH - GraphQL is a Animal Crossing New Horizons [GraphQL](https://graphql.org/) API server. It handles API requests and supports an interactive GraphQL IDE to sample queries.

---

:exclamation: NOTE :exclamation:

This project is no longer maintained as of 11/04/2020. Please disregard mentions of hosted API server and
database update status.

Although this project is not maintained anymore, feel free to fork from this repo and update it for your needs.

---

**Why GraphQL?**

I chose GraphQL for several reasons!

1. I wanted to learn it!
2. I'm writing an API for all kinds of developers, and I don't know what data they need - graphql handles data selection very well
3. GraphQL allows me to let you customize options very well (i.e. check out Docs for Filters!)

**Ready to get started?**

Read the API documentation [here](#documentation)!

- If you're new to GraphQL, I suggest you read all the Doc pages
- If you've used GraphQL before, you can skip the GraphQL Doc

## Table of Contents

- [Server Database Status](#server-database-status)
- [Self-Hosting Installation](#self-hosting-installation)
- [Sample JSON Responses](#sample-json-responses)
- [Documentation](#documentation)
- [Example Queries](#example-queries)
- [cURL](#curl)
- [Contributing](#contributing)
- [License](#license)

## Server Database Status

| Icon               |  Status           |
|:------------------:|:------------------|
| :heavy_check_mark: | OK                |
| :x:                | NOT READY         |
| :warning:          | TO BE UPDATED     |

| Table              | Status                | Summary                          |
|:-------------------|:---------------------:|:---------------------------------|
| Item               | :heavy_check_mark:    | Item-related Entries             |
| Wallpaper          | :heavy_check_mark:    | Wallpaper Entries                |
| Floor              | :heavy_check_mark:    | Flooring Entries                 |
| Clothes            | :heavy_check_mark:    | Clothing Entries                 |
| Music              | :heavy_check_mark:    | KK Slider Music Entries          |
| Photos             | :heavy_check_mark:    | Villager Photo Entries           |
| Posters            | :heavy_check_mark:    | Villager Poster Entries          |
| Rug                | :heavy_check_mark:    | Rug Entries                      |
| Tools              | :heavy_check_mark:    | Tool Entries                     |
| Villager           | :heavy_check_mark:    | Villager Entries                 |
| Bugs               | :heavy_check_mark:    | Bug Entries                      |
| Fishes             | :heavy_check_mark:    | Fish Entries                     |
| Fossils            | :heavy_check_mark:    | Fossil Entries                   |
| Fencing            | :heavy_check_mark:    | Fence Entries                    |
| Umbrellas          | :heavy_check_mark:    | Umbrellas                        |
| Art                | :heavy_check_mark:    | Art Entries                      |
| Construction       | :heavy_check_mark:    | Construction Project Entries     |
| Nook Miles         | :heavy_check_mark:    | Nook Miles Options               |
| Reactions          | :heavy_check_mark:    | Reaction Types                   |
| Other              | :heavy_check_mark:    | Other Misc. Entries              |

Check [issues](https://github.com/Isabelle-Dev/graphql/issues) for detailed descriptions on **TO BE UPDATED** tables.

## Self-Hosting Installation

Want to host the server yourself?

- Run `go get -u github.com/Isabelle-Dev/graphql` in terminal
- Setup [PostgreSQL](https://www.postgresql.org/)
- Create and configure `.config` file using `example.config` as template
- Import data using `csv` files
- `go build -o graphql.exe`
- `./graphql.exe`

All **csv data files** can be found in the `csv` directory. Postgres migration code (for linux) can be found in `csv/linux`. `csv/master` contains the master excel data sheet used for data importation and cleaning.

## Sample JSON Responses

**Sample JSON** responses can be found in the `newhorizons/sample` directory.

You can also play around with different query options by visiting [https://acnhgraphql.com](https://acnhgraphql.com/)

The endpoint itself renders GraphiQL - a GraphQL IDE.

## Documentation

| Doc Type   | Documentation                           |
|:-----------|:---------------------------------------:|
| GraphQL    | [:book:](Docs/graphql.md)               |
| Queries    | [:book:](Docs/query.md)                 |
| Filters    | [:book:](Docs/filters.md)               |
| Schema     | [:book:](Docs/schema.md)                |

## Example Queries

:exclamation: **See Documentation For More Info** :exclamation:

```graphql
query FloorDemo {
  floor {
    query(query: "buy:\"<= 3000 AND > 2000\" color:\"gray AND beige\"", limit: 3) {
      floors {
        Colors
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
query ItemDemo {
  item {
    query(query: " name:\"leaf\" tag:\"plant\" color:\"orange\" ", glob:"t") {
      items {
        Name
        Buy
        Sell
        Concepts
        HHASet
        HHASeries
        Variants {
          Colors
          Pattern
          Image
          Variation
        }
        Tag
      }
    }
  }
}
```

```graphql
query ArtDemo {
  art {
    query(query: "buy: \"4980\" tag: \"picture\" ") {
      art {
        Name
        Buy
        Tag
        Category
        Source
        Type {
          Concepts
          Genuine
          Sell
          Image
        }
      }
    }
  }
}
```

## cURL

Example **cURL** queries can be found at [post.json](post.json) and [post.graphql](post.graphql). I do not recommend making graphql requests using pure cURL, but if you must, it's easier to port requests using an external file.

**Example cURL With JSON File:**

`curl -H "Content-Type:application/json" --data @post.json https://acnhgraphql.com`

**Example cURL With graphql File:**

`curl -H "Content-Type:application/graphql" --data @post.graphql https://acnhgraphql.com`

## Contributing

See [CONTRIBUTING](.github/CONTRIBUTING.md) for more details.

## License

Isabelle-Dev graphql server is licensed under the **MIT License**.

See [License](LICENSE) for more details.

## External Contributions

- All data is sourced from New Horizons data found [here](https://docs.google.com/spreadsheets/d/13d_LAJPlxMa_DubPTuirkIV4DERBMXbrWQsmSh8ReK4/)

- High-resolution image links are provided by [Dodo Codes](https://acnhcdn.com/)
