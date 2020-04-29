rootProject.name = "monorepo"

// order is important, don't change it
includeBuild("tools/gradle-plugins/gradle-hello-plugin")
includeBuild("services/client")
includeBuild("services/server")
includeBuild("libraries/common")
includeBuild("libraries/logging")

// buildCache {
//     local {
//         isEnabled = false
//     }
//     remote<HttpBuildCache> {
//         isEnabled = true
//         isPush = true
//         url = uri("http://localhost:5071/cache/")
//     }
// }
//
// plugins {
//     id("com.gradle.enterprise").version("3.2.1")
// }
//
// gradleEnterprise {
//     buildScan {
//         termsOfServiceUrl = "https://gradle.com/terms-of-service"
//         termsOfServiceAgree = "yes"
//     }
// }


