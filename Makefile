# List of all projects
PROJECTS = \
	01-hello-world \
	02-variables-types \
	03-functions \
	04-control-flow \
	05-arrays-slices \
	06-maps \
	07-strings-runes \
	08-type-conversions \
	09-first-class-functions \
	10-structs \
	11-methods \
	12-interfaces \
	13-pointers \
	14-error-handling \
	15-defer-panic-recover \
	16-goroutines \
	17-channels \
	18-select \
	19-sync-primitives \
	20-sync-atomic \
	21-sync-map \
	22-context \
	23-packages \
	24-fmt-formatting \
	25-json \
	26-encoding \
	27-file-io \
	28-bytes-strconv \
	29-path-filepath \
	30-time \
	31-regexp \
	32-sort \
	33-http-server \
	34-net-tcp-udp \
	35-url \
	36-flag \
	37-log \
	38-os-exec \
	39-os-signal \
	40-templates \
	41-testing \
	42-race-detector \
	43-profiling \
	44-unsafe \
	45-go-embed \
	46-go-generate \
	47-build-tags \
	48-workspaces \
	49-reflection \
	50-generics \
	51-cgo \
	52-advanced-patterns

.PHONY: all help clean $(PROJECTS)

# Run all projects
all:
	@echo "╔══════════════════════════════════════════════════════════╗"
	@echo "║  Running all study projects...                             ║"
	@echo "╚══════════════════════════════════════════════════════════╝"
	@echo ""
	@for project in $(PROJECTS); do \
		echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"; \
		echo "▶ Running: $$project"; \
		echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"; \
		echo ""; \
		cd $$project && go run main.go || true; \
		cd ..; \
		echo ""; \
		echo ""; \
	done
	@echo "╔══════════════════════════════════════════════════════════╗"
	@echo "║  All projects have been executed!                        ║"
	@echo "╚══════════════════════════════════════════════════════════╝"

# Run a specific project
01-hello-world:
	@cd 01-hello-world && go run main.go

02-variables-types:
	@cd 02-variables-types && go run main.go

03-functions:
	@cd 03-functions && go run main.go

04-control-flow:
	@cd 04-control-flow && go run main.go

05-arrays-slices:
	@cd 05-arrays-slices && go run main.go

06-maps:
	@cd 06-maps && go run main.go

07-strings-runes:
	@cd 07-strings-runes && go run main.go

08-type-conversions:
	@cd 08-type-conversions && go run main.go

09-first-class-functions:
	@cd 09-first-class-functions && go run main.go

10-structs:
	@cd 10-structs && go run main.go

11-methods:
	@cd 11-methods && go run main.go

12-interfaces:
	@cd 12-interfaces && go run main.go

13-pointers:
	@cd 13-pointers && go run main.go

14-error-handling:
	@cd 14-error-handling && go run main.go

15-defer-panic-recover:
	@cd 15-defer-panic-recover && go run main.go

16-goroutines:
	@cd 16-goroutines && go run main.go

17-channels:
	@cd 17-channels && go run main.go

18-select:
	@cd 18-select && go run main.go

19-sync-primitives:
	@cd 19-sync-primitives && go run main.go

20-sync-atomic:
	@cd 20-sync-atomic && go run main.go

21-sync-map:
	@cd 21-sync-map && go run main.go

22-context:
	@cd 22-context && go run main.go

23-packages:
	@cd 23-packages && go run main.go

24-fmt-formatting:
	@cd 24-fmt-formatting && go run main.go

25-json:
	@cd 25-json && go run main.go

26-encoding:
	@cd 26-encoding && go run main.go

27-file-io:
	@cd 27-file-io && go run main.go

28-bytes-strconv:
	@cd 28-bytes-strconv && go run main.go

29-path-filepath:
	@cd 29-path-filepath && go run main.go

30-time:
	@cd 30-time && go run main.go

31-regexp:
	@cd 31-regexp && go run main.go

32-sort:
	@cd 32-sort && go run main.go

33-http-server:
	@cd 33-http-server && go run main.go

34-net-tcp-udp:
	@cd 34-net-tcp-udp && go run main.go

35-url:
	@cd 35-url && go run main.go

36-flag:
	@cd 36-flag && go run main.go

37-log:
	@cd 37-log && go run main.go

38-os-exec:
	@cd 38-os-exec && go run main.go

39-os-signal:
	@cd 39-os-signal && go run main.go

40-templates:
	@cd 40-templates && go run main.go

41-testing:
	@cd 41-testing && go test -v

42-race-detector:
	@cd 42-race-detector && go run -race main.go

43-profiling:
	@cd 43-profiling && go run main.go

44-unsafe:
	@cd 44-unsafe && go run main.go

45-go-embed:
	@cd 45-go-embed && go run main.go

46-go-generate:
	@cd 46-go-generate && go generate && go run main.go

47-build-tags:
	@cd 47-build-tags && go run main.go

48-workspaces:
	@cd 48-workspaces && go run main.go

49-reflection:
	@cd 49-reflection && go run main.go

50-generics:
	@cd 50-generics && go run main.go

51-cgo:
	@cd 51-cgo && go run main.go

52-advanced-patterns:
	@cd 52-advanced-patterns && go run main.go

# Show help
help:
	@echo "╔══════════════════════════════════════════════════════════╗"
	@echo "║  Available commands:                                      ║"
	@echo "╚══════════════════════════════════════════════════════════╝"
	@echo ""
	@echo "  make all                    - Run all projects"
	@echo "  make <project-name>         - Run a specific project"
	@echo "  make help                   - Show this help"
	@echo ""
	@echo "Available projects:"
	@echo ""
	@for project in $(PROJECTS); do \
		echo "  - $$project"; \
	done
	@echo ""

# Clean compiled files (optional)
clean:
	@for project in $(PROJECTS); do \
		rm -f $$project/$$project 2>/dev/null || true; \
	done
	@echo "Compiled files removed!"
