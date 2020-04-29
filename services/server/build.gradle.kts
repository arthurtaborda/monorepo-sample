plugins {
    application

    id("com.arthurtaborda.gradle.hello")
    kotlin("jvm") version "1.3.72"
}

repositories {
    jcenter()
}

group = "com.arthurtaborda.monorepo"

dependencies {
    implementation(kotlin("stdlib"))

    implementation("com.arthurtaborda.monorepo:logging")
    implementation("com.arthurtaborda.monorepo:common")
}

application {
    mainClassName = "com.arthurtaborda.monorepo.server.ServerApp"
}
