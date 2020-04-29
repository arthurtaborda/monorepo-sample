plugins {
    kotlin("jvm") version "1.3.72"

    id("org.gradle.java-gradle-plugin")
}

repositories {
    mavenCentral()
    jcenter()
    maven { url = uri("https://plugins.gradle.org/m2/") }
}

dependencies {
    implementation(kotlin("stdlib"))
}

group = "com.arthurtaborda.gradle"

gradlePlugin {
    plugins {
        create("customPlugin") {
            id = "com.arthurtaborda.gradle.hello"
            implementationClass = "com.arthurtaborda.gradle.hello.HelloPlugin"
        }
    }
}
