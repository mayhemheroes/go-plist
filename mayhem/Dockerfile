FROM golang:1.17.13 as gofuzzbuilder
ARG project='go-plist'
ARG fuzzTargetPath=''

ADD . ${project}
WORKDIR ${project}/${fuzzTargetPath}
RUN go get -u github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN GO_ENABLED=0 go-fuzz-build -libfuzzer -o /.a .

FROM silkeh/clang:latest as sanitizer
COPY --from=gofuzzbuilder /.a .a
RUN clang -fsanitize=fuzzer /.a -o /.libfuzzer

FROM ubuntu:20.04
COPY --from=sanitizer /.libfuzzer /
