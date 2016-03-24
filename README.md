# SunSpec tools for Go

This package contains Go data types representing the [SunSpec][] information model.
Subpackages implement particular use-cases that rely on this domain model:

 * [/core](./core) contains the SunSpec model block definitions and functions for working with them
 * [/modbus](#) contains utilities for talking to SunSpec devices via Modbus RTU
 * [/xml](./xml) contains utilities for exchanging SunSpec data using the XML format
 * [/generators](./generators) implements the code that transforms the [SMDX models][] into Go code for the above modules
 * [/server](#) is an example HTTP server that receives SunSpec XML data

[SunSpec]: http://sunspec.org/
[SMDX models]: https://github.com/sunspec/models
