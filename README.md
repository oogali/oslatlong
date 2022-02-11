oslatlong
=========

[Quickly conceived app](https://news.ycombinator.com/item?id=30301118) to query OSM's [Nominatim](https://nominatim.openstreetmap.org/ui/search.html) geocoding [API](https://nominatim.org/release-docs/latest/api/Overview/).

## How do I install it?

### Homebrew (Mac)
```
brew install oogali/tap/oslatlong
```

## How do I use it?

### Examples

#### Return all matching coordinates
```
oslatlong "Berlin, Germany"
```

#### Return only the first matching coordinate
```
oslatlong -n 1 "Berlin, Germany"
```

#### Return all matching coordinates, but prefix with the query
```
oslatlong -s "Berlin, Germany"
```

## Why did you make this?

I needed a distraction and this seemed to scratch that itch.
