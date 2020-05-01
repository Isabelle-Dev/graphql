<p align="center">
  <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/1/17/GraphQL_Logo.svg/1024px-GraphQL_Logo.svg.png" height="200" width="200">
</p>

# Animal Crossing - New Horizons GraphQL Server

[acnhgraphql.com](https://acnhgraphql.com/) is a Animal Crossing New Horizons [GraphQL](https://graphql.org/) API server. It handles API requests and provides an interactive GraphQL IDE to sample queries.

You can access both the IDE and endpoint at [https://acnhgraphql.com](https://acnhgraphql.com).

**Why GraphQL?**

I chose GraphQL for several reasons!

1. I wanted to learn it!
2. I'm writing an API for all kinds of developers, and I don't know what data they need - graphql handles data selection very well
3. GraphQL allows me to let you customize options very well (i.e. check out Docs for Filters!)

**Ready to get started?**

Read the API documentation [here](#documentation)!

- If you're new to GraphQL, I suggest you read all the Doc pages
- If you've used GraphQL before, you can jump right into the _Queries_ and _Filter_ Docs

**I'm still setting up the Linux server (and coding), so not all features are available yet!**

## Table of Contents

- [Server Database Status](#server-database-status)
- [Self-Hosting Installation](#self-hosting-installation)
- [Sample JSON Responses](#sample-json-responses)
- [Documentation](#documentation)
- [cURL](#curl)
- [Bugs](#bugs)
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
| Fish and Bugs      | :x:                   | Fish and Bug Entries             |
| Fossils            | :x:                   | Fossil Entries                   |
| Fencing            | :x:                   | Fence Entries                    |
| Umbrellas          | :x:                   | Umbrellas (TBA)                  |
| Recipes            | :x:                   | Recipe Entries (TBA)             |
| Art                | :x:                   | Art Entries                      |
| Construction       | :x:                   | Construction Project Entries     |
| Nook Miles         | :x:                   | Nook Miles Options               |
| Other              | :x:                   | Other Misc. Entries              |

## Self-Hosting Installation

Want to host the server yourself?

- `go get -u github.com/Isabelle-Dev/graphql`
- Setup PostgreSQL
- Create and configure `.config` file using `example.config` as template
- Import data using `csv` files
- `go build -o graphql.exe`
- `./graphql.exe`

All **csv data files** and **Postgres migration code** are located in the `csv` directory.

## Sample JSON Responses

**Sample JSON** responses can be found in the `newhorizons/sample` directory.

You can also play around with different query options by visiting [https://acnhgraphql.com/](https://acnhgraphql.com/)

The endpoint itself renders GraphiQL - a GraphQL IDE.

## Documentation

| Doc Type   | Documentation                           |
|:-----------|:---------------------------------------:|
| GraphQL    | [:book:](Docs/graphql.md)               |
| Queries    | [:book:](Docs/query.md)                 |
| Filters    | [:book:](Docs/filters.md)               |

## Example Queries

**See Documentation For More Info**

```graphql
query FloorDemo {
  floor {
    query(query: "buy:\"<= 3000 AND > 2000\" color:\"gray AND beige\"") {
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
query DemoItem {
  item {
    query(query: " name:\"leaf\" tag:\"plant\" color:\"orange\" " glob:"t" limit: 30) {
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

```graphql
query DemoItemV2 {
  item {
    query(query: "sell:\"< 2000 AND > 1500\" color:\"(pink AND green) OR orange\"") {
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

## cURL

Example **cURL** queries can be found in `post.json` and `post.graphql`. I do not recommend making graphql requests using pure cURL, but if you must, it's easier to port requests using an external file.

**Example cURL With JSON File:**

`curl -H "Content-type:application/json" --data @post.json https://acnhgraphql.com`

**Example cURL With graphql File:**

`curl -H "Content-type:application/graphql" --data @post.graphql https://acnhgraphql.com`

## Bugs

Oops, bugs are unintentional (I promise!). If you find one, please open an issue with a description of what the bug is.

## Contributing

If you would like to help develop, follow the steps below:

- Fork the repo
- Create your own feature branch
- Commit your changes and push to the new branch
- Open a pull request

## License

Isabelle-Dev graphql server is licensed under the **MIT License**.

See [License](LICENSE) for more details.

## External Contributions

- All data is sourced from New Horizons data found [here](https://docs.google.com/spreadsheets/d/13d_LAJPlxMa_DubPTuirkIV4DERBMXbrWQsmSh8ReK4/)

- High-resolution image links and hosting is provided by [Dodo Codes](https://acnhcdn.com/)
