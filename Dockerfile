# testing for deployment

FROM golang:1.22 as build
WORKDIR /cmd/api
COPY . .
RUN make build && make run

FROM scratch
COPY --from=build /stack-images-md /stack-images-md
EXPOSE 8000
CMD ["/stack-images-md"]