# How to Contribute
The project may look complex at first, but it's not. You can try to understand how the project works, then you can contribute by adding new features, refactoring code architecture to match standards
or fixing issues.

## How you can help?
- You can refactor the frontend design to make it simpler and easier to use.
- You can help by creating more documentations for the project.
- It's hard to write tests for a project like that, because it uses filesystem and we need to generate example files etc to test the general functionality. Because of that for now
folderhost does not have tests yet. It would be nice to help us or give an idea about how we can start to creating a tests. If you have a good idea we're ready!
- Project still does not have automatic CI. You can help to integrate it after setting tests!
- You can contribute for the Services feature.
- You can add comments to explain the code better!
- And much more.

## How to start

### Development dependencies:
- [Go language](https://go.dev/)
- [Git](https://git-scm.com/)
- [Node.js](https://nodejs.org/en)
- IDE / Code editor like [VS Code](https://code.visualstudio.com/)
- [MinGW](https://www.msys2.org/) / GCC compiler for the some go libraries. It depends on what OS you're using, usually some linux distros have it built-in. If you're on Linux make sure to use MinGW to create windows builds. Otherwise you can't
create Windows builds.
- Make to use Makefile (Recommended)
- [Docker](https://www.docker.com/) (Optional)

### Getting started
1. Create a fork on GitHub.
2. Clone the repository to your computer:
```sh
git clone https://github.com/YOUR-USERNAME/folderhost.git
cd folderhost
```
3. Install the dependencies:
```bash
# Using Makefile (Recommended)
make setup
```
```bash
# Or install them manually
go mod tidy
go mod download
cd web && npm install
cd ..
```
4. Start the servers:
```bash
# Frontend
cd web && npm run dev
```
```bash
# Main server
go run main.go
```
```bash
# Docker example
docker compose build
docker compose up
```
5. Open your browser
  - Frontend: http://localhost:5173
  - Main Server: http://localhost:5000

## Testing
While comprehensive testing is challenging due to filesystem dependencies, we welcome:
- Unit tests for utility functions
- Mock-based tests for filesystem operations
- Integration test ideas and proposals
- Test infrastructure setup help

## Commit naming
Try to use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) as much as possible.
### Why use Conventional Commits?
- Automatically generating CHANGELOGs.
- Automatically determining a semantic version bump (based on the types of commits landed).
- Communicating the nature of changes to teammates, the public, and other stakeholders.
- Triggering build and publish processes.
- Making it easier for people to contribute to your projects, by allowing them to explore a more structured commit history.

### Examples:
- ```feat: add file upload service```
- ```fix: resolve memory leak in file handler```
- ```docs: update API documentation```
- ```refactor: frontend design```
- ```chore: add comments```
