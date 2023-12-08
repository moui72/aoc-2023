ROOT_DIR=~/aoc-2023
MOD_DIR="$(ROOT_DIR)/$(NAME)"

new_mod:
	mkdir $(MOD_DIR)
	cd $(MOD_DIR); ls; go mod init .
	cd $(MOD_DIR); go work use ./$(NAME)

remove_mod:
	rm -rf $(MOD_DIR)
