## URL Shortener

The basic concept URL Shortener by mapping the long URL to shortener, and unique. When a user accesses the short URL, the system looks up the origin long URL and redirect them.

### How It Works
- User submits a long URL
- System generates a short, unique identifier (e.g `kmzw87a`)
- The mapping between short URL and long URL is stored on storage (or database)
- When user visits the short URL, the system looks up to the mapping. If the mapping is exist, they will be redirected to the origin URL

## Dependencies
- Gin
- Redis

## Routing

TODO