group="monorepo.library"

plugins {
    kotlin("jvm") version "1.3.70"
}

repositories {
    jcenter()
}

dependencies {
    implementation(kotlin("stdlib"))
    testImplementation(kotlin("test-junit"))
}
