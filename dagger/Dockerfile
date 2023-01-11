FROM ruby

# will copy the root of the repo, not dagger/
WORKDIR /src
COPY . .

# default to the target that lists the checks to run
ENTRYPOINT ["./dagger/ci"]
