plugins {
    application
    kotlin("jvm") version "1.3.70"
}

application {
    mainClassName = "monorepo.emailservice.MainKt"
}

repositories {
    jcenter()
}

dependencies {
    implementation("monorepo.library:string-lib:1.0")
    implementation(kotlin("stdlib"))
}
