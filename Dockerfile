FROM golang:1.22.0 as builder
ADD . /build
WORKDIR /build
RUN make -j4

FROM alpine as putter
COPY --from=builder /build/build/kubecog-helper .
ENTRYPOINT [ "mv", "kubecog-helper", "/custom-tools/" ]
