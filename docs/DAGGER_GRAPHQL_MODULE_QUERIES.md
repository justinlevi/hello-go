# Working GraphQL Queries for hello-go Dagger Module

This document contains tested GraphQL queries that successfully retrieve functions from the hello-go Dagger module.

## Query Module Functions with Basic Info

This query retrieves all functions with their names and descriptions:

```bash
echo '{"query":"{ moduleSource(refString: \".\") { asModule { name description objects { asObject { name functions { name description } } } } } }"}' | \
  dagger run sh -c 'curl -s \
    -u $DAGGER_SESSION_TOKEN: \
    -H "content-type:application/json" \
    -d @- \
    http://127.0.0.1:$DAGGER_SESSION_PORT/query' | jq '.'
```

**Result:**
- Module name: `hello-go`
- Object name: `HelloGo`
- Functions retrieved:
  - `ci` - CI runs the complete CI pipeline with all checks
  - `format` - Format checks code formatting with gofmt
  - `vet` - Vet runs go vet for static analysis
  - `test` - Test runs tests with race detection and coverage
  - `build` - Build creates binaries for multiple platforms
  - `docker` - Docker builds and tests the Docker image
  - `checkFormat` - CheckFormat runs only the formatting check
  - `checkVet` - CheckVet runs only the static analysis
  - `runTests` - RunTests runs only the tests
  - `buildBinaries` - BuildBinaries runs only the build step
  - `buildDocker` - BuildDocker runs only the Docker build and test

## Query Module Functions with Full Details

This query includes function arguments, types, and return types:

```bash
echo '{"query":"{ moduleSource(refString: \".\") { asModule { name description objects { asObject { name functions { name description args { name description defaultValue typeDef { kind asObject { name } asScalar { name } asList { elementTypeDef { kind asObject { name } asScalar { name } } } } } returnType { kind asObject { name } asScalar { name } asList { elementTypeDef { kind asObject { name } asScalar { name } } } } } } } } } }"}' | \
  dagger run sh -c 'curl -s \
    -u $DAGGER_SESSION_TOKEN: \
    -H "content-type:application/json" \
    -d @- \
    http://127.0.0.1:$DAGGER_SESSION_PORT/query' | jq '.'
```

**Result includes:**
- All function names and descriptions
- Arguments for each function:
  - Most functions take a `source` argument of type `Directory`
  - `graphQlexplorer` takes an optional `port` argument (integer, default: 8080)
- Return types:
  - Most functions return `VOID_KIND` (no return value)
  - `test` and `runTests` return `STRING_KIND` (coverage information)
  - `graphQlexplorer` returns a `Service` object

## Key Query Patterns

### 1. Using moduleSource
The `moduleSource` query is the primary way to load a module by reference:
```graphql
moduleSource(refString: ".") {
  asModule {
    # Module fields
  }
}
```

### 2. Module Structure
The module has this hierarchy:
- Module (`hello-go`)
  - Objects (TypeDef array)
    - Object (`HelloGo`)
      - Functions (array)
        - Function details (name, args, returnType, etc.)

### 3. TypeDef Kinds
When querying type definitions, use the `kind` field to determine the type:
- `OBJECT_KIND` - Complex types like Directory, Service
- `STRING_KIND` - String return values
- `INTEGER_KIND` - Integer arguments
- `VOID_KIND` - No return value

### 4. Accessing Type Details
Depending on the `kind`, use the appropriate field:
- `asObject { name }` for object types
- `asScalar { name }` for scalar types
- `asList { elementTypeDef { ... } }` for list types

## Alternative Approaches

### Get Available Query Fields
To explore what queries are available:

```bash
echo '{"query":"{ __schema { queryType { fields { name } } } }"}' | \
  dagger run sh -c 'curl -s \
    -u $DAGGER_SESSION_TOKEN: \
    -H "content-type:application/json" \
    -d @- \
    http://127.0.0.1:$DAGGER_SESSION_PORT/query' | jq -r '.data.__schema.queryType.fields[].name' | sort
```

### Introspection Query
To understand the structure of specific types:

```bash
echo '{"query":"{ __type(name: \"Module\") { fields { name type { name kind } } } }"}' | \
  dagger run sh -c 'curl -s \
    -u $DAGGER_SESSION_TOKEN: \
    -H "content-type:application/json" \
    -d @- \
    http://127.0.0.1:$DAGGER_SESSION_PORT/query' | jq '.'
```

## Notes

1. The `currentModule` query doesn't work in this context because we're not running inside a module.
2. Direct module queries like `{ helloGo { ... } }` don't work - you must use `moduleSource` to load the module first.
3. The module must be loaded from the current directory using `refString: "."`.
4. All queries must be run inside `dagger run` to have access to the session environment variables.