# Schema

Schema will be defined here for ease of use.

## Root

```graphql
type RootQuery {
  art: art
  clothes: clothes
  floor: flooring
  item: item
  music: music
  photos: photo
  poster: posters
  rugs: rug
  tool: tool
  villagers: villager
  wallpaper: wallpapers
}
```

## Art

```graphql
type art {
  query(query: String, glob: String = "f", limit: Int = 500): art_search!
}
```

```graphql
type art_search {
  art: [artObj!]!
}
```

```graphql
type artObj {
  Buy: Int
  Category: String
  Name: String
  Source: String
  Tag: String
  Type: [artType]
}
```

```graphql
type artType {
  Concept: [String]
  Genuine: String
  Image: String
  Sell: Int
}
```

  


