FROM mcr.microsoft.com/dotnet/aspnet:5.0 AS Final

RUN apt-get update && apt-get install -y screen \
    curl \
libleveldb-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /neo-cli
COPY neo-cli .
RUN chmod +x entry.sh

ENTRYPOINT ["/neo-cli/entry.sh" ]