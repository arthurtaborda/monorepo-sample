plugins {
    kotlin("jvm") version "1.3.72"

    id("com.arthurtaborda.gradle.hello")
}

repositories {
    jcenter()
}

group = "com.arthurtaborda.monorepo"

dependencies {
    implementation(kotlin("stdlib"))
    
    implementation("com.arthurtaborda.monorepo:common")
}
