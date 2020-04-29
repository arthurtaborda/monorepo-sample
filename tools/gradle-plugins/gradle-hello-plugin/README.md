# Hello plugin

## How to build

```
./gradlew clean build
```

## How to use

Include build in your **settings.xml**.

```
includeBuild("${REPOSITORY_ROOT}/tools/gradle-plugins/gradle-hello-plugin")
```

Apply plugin in your **build.gradle**.

```
plugins {
    id("com.arthurtaborda.gradle.hello")
}
```
