plugins {
    application
    kotlin("jvm") version "1.3.70"
}

application {
    mainClassName = "monorepo.apigateway.MainKt"
}

repositories {
    jcenter()
}

dependencies {
    implementation("monorepo.library:number-lib:1.0")
    implementation("monorepo.library:string-lib:1.0")
    implementation(kotlin("stdlib"))
    testImplementation(kotlin("test-junit"))
}
