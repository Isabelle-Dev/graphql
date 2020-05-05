# Schema

The schema for the server will be defined here for ease of use.

Further info on graphql types and schemas can be found [here](https://graphql.org/learn/schema/).

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
- [Fencing](#fencing)
- [Construction](#construction)
- [Nook Miles](#nook-miles)
- [Umbrella](#umbrella)
- [Other](#other)
- [Bug](#bug)
- [Fish](#fish)
- [Fossil](#fossil)
- [Reactions](#reactions)

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
  Concepts: [String]
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
  clothes: [clothesObj]!
}
```

```graphql
type ClothesObj {
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
  Colors: [String]
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
  floors: [floorObj!]!
}
```

```graphql
type floorObj {
  Buy: Int
  Catalog: String
  Colors: [String]
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
  query(query: String, glob: String = "f", limit: Int = 1000): item_search!
}
```

```graphql
type item_search {
  items: [itemObj!]!
}
```

```graphql
type itemObj {
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
  Image: String
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
  Colors: [String]
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
  Colors: [String]
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
  Colors: [String]
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
  Sources: [String]
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
  Colors: [String]
  Gender: String
  Hobby: String
  HouseImage: String
  Image: String
  Name: String
  Personality: String
  Species: String
  Styles: [String]
}
```

## Wallpaper

```graphql
type wallpapers {
  query(query: String, glob: String = "f", limit: Int = 500): wallpaper_search!
}
```

```graphql
type wallpaper_search {
  wallpapers: [wallpaperObj!]!
}
```

```graphql
type wallpaperObj {
  Buy: Int
  CeilingType: String
  Colors: [String]!
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

## Fencing

```graphql
type fencing {
  query(query: String, glob: String = "f", limit: Int = 500): fence_search!
}
```

```graphql
type fence_search {
  fences: [fenceObj]!
}
```

```graphql
type fenceObj {
  Buy: Int
  DIY: String
  Image: String
  Name: String
  Sell: Int
  Source: String
  SourceNotes: String
}
```

## Construction

```graphql
type construction {
  query(query: String, glob: String = "f", limit: Int = 500): construction_search!
}
```

```graphql
type construction_search {
  construction: [constructionObj]!
}
```

```graphql
type constructionObj {
  Buy: Int
  Category: String
  Image: String
  Name: String
  Source: String
}
```

## Nook Miles

```graphql
type nookmiles {
  query(query: String, glob: String = "f", limit: Int = 500): nookmile_search!
}
```

```graphql
type nookmile_search {
  nook_mile: [nookmileObj!]!
}
```

```graphql
type nookmileObj {
  Category: String
  Image: String
  Name: String
  NookMiles: Int
}
```

## Umbrella

```graphql
type umbrellas {
  query(query: String, glob: String = "f", limit: Int = 500): umbrella_search!
}
```

```graphql
type umbrella_search {
  umbrellas: [umbrellaObj!]!
}
```

```graphql
type umbrellaObj {
  Buy: Int
  Catalog: String
  Colors: [String]
  DIY: String
  Image: String
  Name: String
  Sell: Int
  Source: String
  SourceNotes: String
}
```

## Other

```graphql
type other {
  query(query: String, glob: String = "f", limit: Int = 500): other_search!
}
```

```graphql
type other_search {
  others: [otherObj!]!
}
```

```graphql
type otherObj {
  Buy: Int
  DIY: String
  Image: String
  Name: String
  Sell: Int
  Source: String
  SourceNotes: String
  Tag: String
}
```

## Bug

```graphql
type bugs {
  query(query: String, glob: String = "f", limit: Int = 500): bug_search!
}
```

```graphql
type bug_search {
  bugs: [bugObj]!
}
```

```graphql
type bugObj {
  Colors: [String]
  EndTime: String
  Image: String
  Location: String
  Name: String
  Rarity: String
  Sell: Int
  StartTime: String
  Weather: String
}
```

## Fish

```graphql
type fishes {
  query(query: String, glob: String = "f", limit: Int = 500): fish_search!
}
```

```graphql
type fish_search {
  fishes: [fishObj!]!
}
```

```graphql
type fishObj {
  CatchUp: String
  Colors: [String]
  EndTime: String
  Image: String
  Location: String
  Name: String
  Rarity: String
  Sell: Int
  Shadow: String
  StartTime: String
}
```

## Fossil

```graphql
type fossil {
  query(query: String, glob: String = "f", limit: Int = 500): fossil_search!
}
```

```graphql
type fossil_search {
  fossils: [fossilObj]!
}
```

```graphql
type fossilObj {
  Buy: Int
  Catalog: String
  Colors: [String]
  Image: String
  Interact: String
  Name: String
  Sell: Int
  Source: String
}
```

## Reactions

```graphql
type reaction {
  query(query: String, glob: String = "f", limit: Int = 500): reaction_search!
}
```

```graphql
type reaction_search {
  reactions: [reactionObj]!
}
```

```graphql
type reactionObj {
  Image: String
  Name: String
  Source: String
}
```
