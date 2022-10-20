OPERATING_SYSTEM ?= linux
ANDROID_API ?= 26
ANDROID_TOOLCHAIN_DIR ?= $(ANDROID_NDK_HOME)/toolchains/llvm/prebuilt/$(OPERATING_SYSTEM)-x86_64

build:
	CC="$(ANDROID_TOOLCHAIN_DIR)/bin/armv7a-linux-androideabi$(ANDROID_API)-clang" \
	CXX="$(ANDROID_TOOLCHAIN_DIR)/bin/armv7a-linux-androideabi$(ANDROID_API)-clang++" \
	GOOS=android \
	GOARCH=arm \
	CGO_ENABLED=1 \
	go build -buildmode=c-shared -o android/app/src/main/jniLibs/armeabi-v7a/libvulkandraw.so

build64:
	CC="$(ANDROID_TOOLCHAIN_DIR)/bin/aarch64-linux-android$(ANDROID_API)-clang" \
	CXX="$(ANDROID_TOOLCHAIN_DIR)/bin/aarch64-linux-android$(ANDROID_API)-clang++" \
	GOOS=android \
	GOARCH=arm64 \
	CGO_ENABLED=1 \
	go build -buildmode=c-shared -o android/app/src/main/jniLibs/arm64-v8a/libvulkandraw.so

xxx:
	echo notes
#	cp "$(ANDROID_TOOLCHAIN_DIR)/sysroot/usr/lib/arm-linux-androideabi/26/libvulkan.so" "./android/app/src/main/jniLibs/armeabi-v7a/"
