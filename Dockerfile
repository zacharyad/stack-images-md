# testing for deployment

FROM golang:1.22.1 as build
WORKDIR /
ADD . .
RUN make build
RUN make run

FROM scratch
COPY --from=build /stack-images-md /stack-images-md
EXPOSE 3000
CMD ["/main"]