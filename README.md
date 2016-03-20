# SunSpec tools for Go

This package contains Go data types representing the [SunSpec][] information model.
Subpackages implement particular use-cases that rely on this domain model:

 * [gosunspec/modbus](#) contains utilities for talking to SunSpec devices via Modbus RTU
 * [gosunspec/xml](./xml) contains utilities for exchanging SunSpec data using the XML format
 * [gosunspec/server](#) is an example HTTP server that receives SunSpec XML data

[SunSpec]: http://sunspec.org/
