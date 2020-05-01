# Schema

Schema will be defined here for ease of use.

## Contents

- [Root](#root)
- [Art](#art)
- [Clothes](#clothes)
- [Floor](#floor)
- [Item](#item)
- [Music](#music)
- [Photos](#photos)
- [Posters](#posters)
- [Rugs](#rugs)
- [Tools](#tools)
- [Villager](#villager)
- [Wallpaper](#wallpaper)

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

## Clothes

```graphql
type clothes {
  query(query: String, glob: String = "f", limit: Int = 500): search_clothes!
}
```

```graphql
type search_clothes {
  clothes: [ClothesResult]!
}
```

```graphql
type ClothesResult {
  Buy: Int
  DIY: String
  Name: String
  Sell: Int
  Source: String
  SourceNotes: String
  Themes: [String]
  Variants: [clothesvariants]
}
```

```graphql
type clothesvariants {
  Color: [String]
  Image: String
  Variation: String
}
```

## Floor

```graphql 
type flooring {
  query(query: String, glob: String = "f", limit: Int = 500): floor_search!
}
```

```graphql
type floor_search {
  floors: [floor!]!
}
```

```graphql
type floor {
  Buy: Int
  Catalog: String
  Color: [String]
  Concepts: [String]
  DIY: String
  Image: String
  Name: String
  Sell: Int
  Source: String
  SourceNotes: String
  Tag: String
  VFX: String
}
```

## Item

```graphql
type item {
  query(query: String, glob: String = "f", limit: Int = 500): search_by_item!
}
```

```graphql
type search_by_item {
  items: [ItemResult!]!
}
```

```graphql
type ItemResult {
  BodyCustomize: String
  BodyTitle: String
  Buy: Int
  Concepts: [String]!
  DIY: String
  HHASeries: String
  HHASet: String
  Interact: String
  KitCost: Int
  Name: String
  PatternCustomize: String
  PatternTitle: String
  Sell: Int
  Source: String
  Tag: String
  Variants: [Variant]!
}
```

```graphql
type Variant {
  Colors: [String]!
  Img: String
  Pattern: String
}
```

## Music

```graphql
type music {
  query(query: String, glob: String = "f", limit: Int = 500): music_search!
}
```

```graphql
type music_search {
  music: [musicObj!]!
}
```

```graphql
type musicObj {
  Buy: Int
  Color: [String]
  Image: String
  Name: String
  Sell: Int
  Source: String
  SourceNotes: String
}
```

## Photos

```graphql
type photo {
  query(query: String, glob: String = "f", limit: Int = 500): photo_search!
}
```

```graphql
type photo_search {
  photos: [photoObj!]!
}
```

```graphql
type photoObj {
  BodyTitle: String
  Buy: Int
  Customize: String
  DIY: String
  KitCost: Int
  Name: String
  PatternTitle: String
  Sell: Int
  Source: String
  Variants: [variantObj]
}
```

```graphql
type variantObj {
  Image: String
  Pattern: String
}
```

## Posters

```graphql
type posters {
  query(query: String, glob: String = "f", limit: Int = 500): poster_search!
}
```

```graphql
type poster_search {
  posters: [posterObj!]!
}
```

```graphql
type posterObj {
  Buy: Int
  Color: [String]
  Image: String
  Name: String
  Sell: Int
  Source: String
}
```

## Rugs

```graphql
type rug {
  query(query: String, glob: String = "f", limit: Int = 500): rug_search!
}
```

```graphql
type rug_search {
  rugs: [rugObj!]!
}
```

```graphql
type rugObj {
  Buy: Int
  Catalog: String
  Color: [String]
  Concepts: [String]
  DIY: String
  HHASeries: String
  Image: String
  Name: String
  Sell: Int
  Source: String
  SourceNotes: String
  Tag: String
}
```

## Tools

```graphql
type tool {
  query(query: String, glob: String = "f", limit: Int = 500): tool_search!
}
```

```graphql
type tool_search {
  tools: [toolObj!]!
}
```

```graphql
type toolObj {
  BodyTitle: String
  Buy: Int
  Customize: String
  DIY: String
  KitCost: Int
  Name: String
  Sell: Int
  Set: String
  Source: [String]
  SourceNotes: String
  Uses: Int
  Variant: [tool_variant]
}
```

```graphql
type tool_variant {
  Image: String
  Variation: String
}
```

## Villager

```graphql
type villager {
  query(query: String, glob: String = "f", limit: Int = 500): villager_search!
}
```

```graphql
type villager_search {
  villagers: [villagerObj!]!
}
```

```graphql
type villagerObj {
  Birthday: String
  Catchphrase: String
  Color: [String]
  Gender: String
  Image: String
  Name: String
  Personality: String
  Species: String
  Style: [String]
}
```

## Wallpaper

```graphql
type wallpapers {
  query(query: String, glob: String = "f", limit: Int = 500): search_wallpaper!
}
```

```graphql
type search_wallpaper {
  wallpapers: [WallpaperResult!]!
}
```

```graphql
type WallpaperResult {
  Buy: Int
  CeilingType: String
  Color: [String]!
  Concepts: [String]!
  CurtainInfo: Curtain
  DIY: String
  HHASeries: String
  Image: String
  Name: String
  Sell: Int
  Source: String
  SourceNotes: String
  Tag: String
  VFXInfo: VFX
  WindowInfo: Window
}
```

```graphql
type Curtain {
  CurtainColor: String
  CurtainType: String
}
```

```graphql
type VFX {
  VFX: String
  VFXType: String
}
```

```graphql
type Window {
  Color: String
  Type: String
}
```
