# Go + gRPC + OpenTelemetry Example

An example of using [gRPC](https://grpc.io/) and [OpenTelemetry](https://opentelemetry.io/) in Go.

How to run the example:
1. `jaeger-start.bash` (pulls and) starts [Jaeger](https://www.jaegertracing.io/), opens the Jaeger UI.
2. `server.bash` runs the example Notes server.
3. `client.bash` runs the example Notes client.
4. `grpcui.bash` (installs and) runs [gRPC UI](https://github.com/fullstorydev/grpcui).
5. Now you can send requests to the server, either via the client or gRPC UI. The server returns a note for IDs 1, 2, 3; and a NotFound error for any other ID. You will be able to see the traces in the Jaeger UI.
6. `jaeger-stop.bash` stops and removes Jaeger.
