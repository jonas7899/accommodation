null :=
space := ${null} ${null}
repl := @
#dash := -
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_path := $(subst ${space},${repl},${mkfile_path})
project_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))
#project_dir := $(subst ${repl},${dash},${project_dir})

$(info mkfile_path is [${mkfile_path}])
$(info project_dir is [${project_dir}])

BINARY_NAME=$(project_dir)
SRC_DIR := ./src/
BIN_DIR := ./bin/
ifeq ($(OS),Windows_NT)
    BINARY_NAME := $(BINARY_NAME).exe
endif

git build:
	go build -o ${BIN_DIR}${BINARY_NAME} ${SRC_DIR}main.go

run_go:
	go run ${SRC_DIR}main.go

run:
	./${BIN_DIR}${BINARY_NAME}
 
build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}
