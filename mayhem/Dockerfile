FROM golang:1.23 AS builder

WORKDIR /src
COPY . /src/go-plist
COPY mayhem/fuzz_native_test.go /src/go-plist/fuzz_native_test.go

WORKDIR /src/go-plist
RUN go test -c -fuzz=FuzzDecode -o /decoder-fuzz.libfuzzer .

FROM ubuntu:20.04
COPY --from=builder /decoder-fuzz.libfuzzer /decoder-fuzz.libfuzzer
