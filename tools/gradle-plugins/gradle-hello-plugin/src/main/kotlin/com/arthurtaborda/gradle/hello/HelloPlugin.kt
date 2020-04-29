package com.arthurtaborda.gradle.hello

import org.gradle.api.Plugin
import org.gradle.api.Project

internal class HelloPlugin : Plugin<Project> {
    override fun apply(project: Project) {
        project.task("hello").doLast {
            project.logger.quiet("Hello World!")
        }
    }
}
