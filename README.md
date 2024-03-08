# Eiger

## Description

Implementation of custom rdiff.

## Build

```
go build
```

## Usage

### Signature
```
./eiger signature INPUT_FILE OUTPUT_FILE [CHUNK_SIZE]
```
By default `chunk_size = len(input) ** 0.5`

### Delta
```
./eiger delta INPUT_FILE SIGNATURE_FILE OUTPUT_FILE
```

### Patch

```
./eiger patch BASE_FILE DELTA_FILE OUTPUT_FILE
```

## Testing

#### Unit tests
```
docker compose up test
```

#### Sample test
```
docker compose up sample
```

## References
* https://www.andrew.cmu.edu/course/15-749/READINGS/required/cas/tridgell96.pdf
* https://www.youtube.com/watch?v=X3Stha8pxXc
* https://stackoverflow.com/questions/1535017/rolling-checksums-in-the-rsync-algorithm
