FROM golang

# will copy the root of the repo, not dagger/
WORKDIR /src
COPY . .

# pre-fetch dependencies while we're at it
RUN go mod download

# magefiles/ exists in the root of the repo, so during dev you just run e.g.
# mage test
RUN go install github.com/magefile/mage
ENTRYPOINT ["mage"]
