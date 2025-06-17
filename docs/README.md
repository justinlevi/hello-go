# /docs

Project documentation beyond the root README.md.

## Overview

The `/docs` directory contains additional documentation for the project including design documents, API specifications, deployment guides, and architecture decisions.

## Usage

Organize documentation by type:

```
docs/
├── api/           # API documentation
├── architecture/  # Architecture decisions and diagrams
├── deployment/    # Deployment guides
├── development/   # Development setup and guides
└── images/        # Images and diagrams for documentation
```

## Common Documents

- **API.md** - Detailed API documentation
- **ARCHITECTURE.md** - System architecture and design decisions
- **DEPLOYMENT.md** - How to deploy the application
- **DEVELOPMENT.md** - Development environment setup
- **TROUBLESHOOTING.md** - Common issues and solutions

## Guidelines

- Keep documentation close to code when possible
- Use this for documentation that doesn't fit elsewhere
- Include diagrams and images to clarify complex topics
- Keep documentation up to date with code changes
- Use Markdown for consistency

## Example Structure

```
docs/
├── API.md              # OpenAPI/Swagger spec or detailed API docs
├── ARCHITECTURE.md     # Why we made certain design choices
├── deployment/
│   ├── kubernetes.md   # K8s deployment guide
│   └── docker.md       # Docker deployment guide
└── images/
    └── architecture.png
```

## When to Use

Start using this directory when:
- Your README.md becomes too long
- You need multiple documentation files
- You have complex deployment procedures
- You want to document design decisions
- You need to include diagrams or images

## Documentation Best Practices

1. **Start Simple**: Don't over-document early
2. **Keep Current**: Outdated docs are worse than no docs
3. **Be Practical**: Focus on what developers need
4. **Use Examples**: Show, don't just tell
5. **Link Wisely**: Reference other docs and code

## Tools Integration

Consider generating documentation:
- `godoc` for Go package documentation
- OpenAPI/Swagger for API documentation
- PlantUML for diagrams
- MkDocs or Hugo for documentation sites