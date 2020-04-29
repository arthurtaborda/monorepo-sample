setOf("clean", "assemble").forEach { taskName ->
    tasks.create(taskName) {
        gradle.includedBuilds.forEach { build ->
            dependsOn(gradle.includedBuild(build.name).task(":$taskName"))
        }
    }
}
