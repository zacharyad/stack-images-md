# testing for deployment

FROM golang:1.22 as build
WORKDIR /
COPY . .
RUN make build && make run

FROM scratch
COPY --from=build /stack-images-md /stack-images-md
EXPOSE 3000
CMD ["/stack-images-md"]