FROM mcr.microsoft.com/dotnet/sdk:10.0.404 AS production

WORKDIR /workspace
COPY src/templates/csharp/global.json src/templates/csharp/Directory.Build.props ./src/templates/csharp/
COPY src/templates/csharp/src/ ./src/templates/csharp/src/
RUN cd src/templates/csharp \
    && dotnet restore src/DDDTemplate.csproj --locked-mode \
    && dotnet build src/DDDTemplate.csproj --no-restore --configuration Release

FROM production AS tests
RUN apt-get update \
    && apt-get install -y --no-install-recommends bash \
    && rm -rf /var/lib/apt/lists/*
COPY src/templates/csharp/DDDTemplate.Tests.csproj src/templates/csharp/packages.lock.json ./src/templates/csharp/
COPY src/templates/csharp/.config/ ./src/templates/csharp/.config/
RUN cd src/templates/csharp && dotnet restore DDDTemplate.Tests.csproj --locked-mode
COPY src/templates/csharp/tests/ ./src/templates/csharp/tests/
COPY scripts/common/ ./scripts/common/
COPY scripts/ci/checks/test-runners/ ./scripts/ci/checks/test-runners/
CMD ["bash", "/workspace/scripts/ci/checks/test-runners/test-unit.sh"]
