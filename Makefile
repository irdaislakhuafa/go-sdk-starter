# old/current project package
APP_PACKAGE_OLD:=$(strip github.com/irdaislakhuafa/go-sdk-starter)

# change value of this project package as you wish
APP_PACKAGE_NEW:=$(strip github.com/irdaislakhuafa/go-sdk-starter)

repackage-linux:
	@ find . -type f -name "*.go" -exec sed -i "s@$(APP_PACKAGE_OLD)@$(APP_PACKAGE_NEW)@g" {} +
repackage-macos:
	@ find . -type f -name "*.go" -exec sed -i '' "s@$(APP_PACKAGE_OLD)@$(APP_PACKAGE_NEW)@g" {} +
