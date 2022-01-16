FROM gcr.io/distroless/static
WORKDIR /
COPY .target/linux_amd64/slack /slack
ENTRYPOINT [ "/slack" ]
