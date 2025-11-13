# Serialization Formats Benchmarks

This repository contains benchmarking code for different serialization formats across programming languages.

Sample data came from [Farcaster](https://github.com/farcasterxyz/hub).

## Tested

### Serialization Formats

    * [Protocol Buffers](https://protobuf.dev)
    * [FlatBuffers](https://google.github.io/flatbuffers/)
    * [Cap'n Proto](https://capnproto.org/)
    * [Diarkis Puffer](https://diarkis.io/)

### Programming Languages

    * Go
    * ...

## Results

> **Note**
> Smaller is better

### Go (13 Nov 2025)

    * Benchmarks are CPU bound (no disk operations)
    * OS: macOS 26.1
    * CPU: Apple M3
    * Puffer: [diarkis/puffer](https://github.com/Diarkis/puffer) v0.1.0
    * Protobuf: [golang/protobuf](https://github.com/golang/protobuf) v1.36.10
    * Protobuf (gogofaster): [gogoprotobuf](https://github.com/gogo/protobuf) v1.3.2
    * FlatBuffers: [google/flatbuffers](https://github.com/google/flatbuffers/tree/master/go) v1.12.1
    * ~~Capnp: [capnproto/go-capnproto2](https://github.com/capnproto/go-capnproto2) v3.1.0-alpha.2~~

| Test                              | Puffer      | Protobuf     | Protobuf (gogofaster) | FlatBuffers    |
| --------------------------------- | ----------- | -------------| --------------------- | -------------- |
| Encode (ns/op)                    | 215.6n ± 4% | 425.9n ± 4%  | **167.0n ± 6%**       | 454.7n ± 1%    |
| Encode (allocs/op)                | 7.000 ± 0%  | 8.000 ± 0%   | **6.000 ± 0%**        | 8.000 ± 0%     |
| Decode (allocs/op)                | 16.00 ± 0%  | 14.00 ± 0%   | 12.00 ± 0%            | **0.000 ± 0%** |
| Wire format size (bytes)          | **297**     | 299          | 299                   | 432            |
| Wire format size, gzipped (bytes) | **322**     | 325          | 324                   | 412            |

**Notes**:

    * FlatBuffers does not require much ram. Protobuf (gogofaster) does not require much CPU.
    * FlatBuffers was created "for game development and other performance-critical applications" at Google as an alternative to Protobuf.
    * See: https://eltonminetto.dev/en/post/2024-08-05-json-vs-flatbuffers-vs-protobuf/
    * Puffer lacks MANY features compared to the other benchmarked languages.
