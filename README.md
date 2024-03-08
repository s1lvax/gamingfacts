# Gaming Facts

## Description

A simple api that returns a random interesting fact about videogames/gaming.

I wholeheartedly believe in learning-by-doing, so this is my first step into learning Go.

### Example Usage

```bash
curl https://gamingfacts.cfsilva.com/
```

Response

```json
{
  "fact": "‘Pac-Man’ was designed to attract female gamers, which led to the character’s round shape and love for eating."
}
```

### Limitations

The API is limited to `20 requests per minute`.

## Local Development

Clone the repository

```bash
git clone https://github.com/s1lvax/gamingfacts
```

Install dependencies

```bash
go mod download
```

Run API

```bash
go run main.go
```
