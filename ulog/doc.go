// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package ulog is the Logging package.
//
// package ulog provides an API wrapper around the logging library
// zap (https://github.com/uber-go/zap). package ulog uses the builder pattern
// to instantiate the logger. With
// LogBuilder you can perform pre-initialization
// setup by injecting configuration, custom logger, and log level prior to building
// the usable
// ulog.Log object.
//
// ulog.Log interface provides a few benefits:
//
// • Decouple services from the logger used underneath the framework
//
// • Easy to use API for logging
//
// • Easily swappable backend logger without changing the service
//
// Sample usage
//
//   package main
//
//   import "go.uber.org/fx/ulog"
//
//   func main() {
//     // Configure logger with configuration preferred by your service
//     logConfig := ulog.Configuration{}
//     logBuilder := ulog.Builder().WithConfiguration(&logConfig)
//
//     // Build ulog.Log from logBuilder
//     log := lobBuilder.Build()
//
//     // Use logger in your service
//     log.Info("Message describing logging reason", "key", "value")
//   }
//
// Context
//
// It is very common that in addition to logging a string message, it is desirable
// to provide additional information: customer uuid, tracing id, etc.
//
//
// For that very reason, the logging methods (Info,Warn, Debug, etc) take
// additional parameters as key value pairs.
//
//
// Retaining Context
//
// Sometimes the same context is used over and over in a logger. For example
// service name, shard id, module name, etc. For this very reason
// With()functionality exists which will return a new instance of the logger with
// that information baked in so it doesn't have to be provided
// for each logging call.
//
//
// For example, the following piece of code:
//
//   package main
//
//   import "go.uber.org/fx/ulog"
//
//   func main() {
//     log := ulog.Logger()
//     log.Info("My info message")
//     log.Info("Info with context", "customer_id", 1234)
//
//     richLog := log.With("shard_id", 3, "levitation", true)
//     richLog.Info("Rich info message")
//     richLog.Info("Even richer", "more_info", []int{1, 2, 3})
//   }
//
// Produces this output:
//
//   {"level":"info","ts":1479946972.102394,"msg":"My info message"}
//   {"level":"info","ts":1479946972.1024208,"msg":"Info with context","customer_id":1234}
//   {"level":"info","ts":1479946972.1024246,"msg":"Rich info message","shard_id":3,"levitation":true}
//   {"level":"info","ts":1479946972.1024623,"msg":"Even richer","shard_id":3,"levitation":true,"more_info":[1,2,3]}
//
// Configuration
//
// ulog configuration can be defined in multiple ways:
//
// Writing the struct yourself
//
//   loggingConfig := ulog.Configuration{
//     Stdout: true,
//   }
//
// Configuration defined in YAML
//
//   logging:
//     stdout: true
//     level: debug
//
// Benchmarks
//
// Current performance benchmark data with ulog interface,
// ulog baseLogger struct, and zap.Logger
//
//   BenchmarkUlogWithoutFields-8                    10000000               223 ns/op               0 B/op          0 allocs/op
//   BenchmarkUlogWithFieldsLogIFace-8                1000000              1237 ns/op            1005 B/op         18 allocs/op
//   BenchmarkUlogWithFieldsBaseLoggerStruct-8        1000000              1096 ns/op             748 B/op         17 allocs/op
//   BenchmarkUlogWithFieldsZapLogger-8               2000000               664 ns/op             514 B/op          1 allocs/op
//   BenchmarkUlogLiteWithFields-8                    3000000               489 ns/op             249 B/op          6 allocs/op
//   BenchmarkUlogSentry-8                            3000000               407 ns/op             128 B/op          4 allocs/op
//   BenchmarkUlogSentryWith-8                        1000000              1535 ns/op            1460 B/op         12 allocs/op
//
// Sentry
//
// ulog has a seamless integration with Sentry. For out-of-the-box usage
// just include this in your configuration yaml:
//
//
//   logging:
//     sentry:
//       dsn: http://user:secret@your.sentry.dsn/project
//
// If you'd like to take more control over the Sentry integration, see
// sentry.Configuration
//
// For example, to turn off Stacktrace generation:
//
//   logging:
//     sentry:
//       dsn: http://user:secret@your.sentry.dsn/project
//       trace:
//         disabled: true
//
// To set up Sentry in code use sentry.Hook. For example:
//
//   import (
//     "github.com/uber-go/zap"
//     "go.uber.org/fx/ulog/ulog"
//     "go.uber.org/fx/ulog/sentry"
//     "go.uber.org/fx/service"
//   )
//
//   func main() {
//     h := sentry.New(MY_DSN, MinLevel(zap.InfoLevel), DisableTraces())
//     l := ulog.Builder().WithSentryHook(h).Build()
//     svc, err := service.WithLogger(l).Build()
//   }
//
//
package ulog
