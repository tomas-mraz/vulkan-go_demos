
# Mention the cmake version name.
cmake_minimum_required(VERSION 3.4.1)

project( VulkanDraw )


# Application
add_library(vulkandraw SHARED IMPORTED)

set_target_properties( vulkandraw
        PROPERTIES IMPORTED_LOCATION
        ${CMAKE_SOURCE_DIR}/src/main/jniLibs/${ANDROID_ABI}/libvulkandraw.so
)


# Vulkan validation layers
#add_library(libVkLayer_khronos_validation SHARED IMPORTED)
#
#set_target_properties( libVkLayer_khronos_validation
#       PROPERTIES IMPORTED_LOCATION
#       ${CMAKE_SOURCE_DIR}/src/main/jniLibs/${ANDROID_ABI}/libVkLayer_khronos_validation.so
#)


find_library( vulkan-lib
        vulkan
        )


find_library( android-lib
        android
        )


find_library( log-lib
        log
        )


target_link_libraries( vulkandraw
        INTERFACE vulkan-lib
        android-lib
        log-lib
        )
