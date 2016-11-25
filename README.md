# SunSpec tools for Go

This package contains Go data types representing the [SunSpec][] information model.
Subpackages implement particular use-cases that rely on this domain model:

 * [/core](./core) contains the types used by the remainder of the library
 * [/smdx](./smdx) contains types that represent (XML) SMDX models defined by the Sunspec specification
 * [/generators](./generators) implements code generators that transform the [SMDX models][SMDX] into Go code
 * [/models](./models) contains 1 generated package for each SMDX model
 * [/xml](./xml) contains utilities for exchanging SunSpec data using the XML format
 * [/modbus](#) contains utilities for talking to SunSpec devices via Modbus RTU
 * [/server](#) is an example HTTP server that receives SunSpec XML data

[SunSpec]: http://sunspec.org/
[SMDX]: https://github.com/sunspec/models

## Generated code

This package uses typesafe representations of each model type. To avoid
excessive manual maintenance, structs for each type of model defined by SunSpec
are generated from the [SMDX model files][SMDX]. To regenerate this code, first
initialise the spec submodule with:

    git submodule update --init spec

Then run the generators:

    go generate ./models
