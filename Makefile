PROJECT_NAME="aoc-2023"
GITHUB_USER_NAME="moui72"

PROJECT_DIR=~/$(PROJECT_NAME)
MOD_DIR=$(PROJECT_DIR)/$(NAME)

ROOT_MOD_PATH=$(GITHUB_USER_NAME)/$(PROJECT_NAME)
NEW_MOD_PATH="$(ROOT_MOD_PATH)/$(NAME)"

new_mod:
	mkdir $(MOD_DIR)
	- cd $(MOD_DIR) && go mod init "$(NEW_MOD_PATH)"
	if ! test  -f $(MOD_DIR)/go.mod; then \
		cd $(PROJECT_DIR); make remove_mod; exit 1; \
	fi
	cp $(PROJECT_DIR)/go.template $(MOD_DIR)/main.go
	go work use ./$(NAME)

remove_mod:
	go work edit -dropuse ./$(NAME)
	rm -rf $(MOD_DIR)
