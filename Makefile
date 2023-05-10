PROJECT_NAME := kubectl-cluster
OUTPUT_DIR := /usr/local/bin

build:
	go build -x -o $(PROJECT_NAME)

install:
	mv $(PROJECT_NAME) $(OUTPUT_DIR)

#uninstall:
#	rm $(OUTPUT_DIR)/$(PROJECT_NAME)