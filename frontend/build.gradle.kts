import org.jetbrains.compose.desktop.application.dsl.TargetFormat

plugins {
    kotlin("jvm") version "1.9.24"
    id("org.jetbrains.compose") version "1.7.0-rc01"
    id("org.jetbrains.kotlin.plugin.compose")
}

group = "com.techalert"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
    maven("https://maven.pkg.jetbrains.space/public/p/compose/dev")
    maven("https://packages.jetbrains.team/maven/p/kpm/public/")
    google()
}

dependencies {
    // Note, if you develop a library, you should use compose.desktop.common.
    // compose.desktop.currentOs should be used in launcher-sourceSet
    // (in a separate module for demo project and in testMain).
    // With compose.desktop.common you will also lose @Preview functionality
    implementation(compose.desktop.currentOs) {
        exclude(group = "org.jetbrains.compose.material")
    }

    implementation("org.jetbrains.jewel:jewel-int-ui-standalone-242:0.25.0")
    implementation("org.jetbrains.jewel:jewel-int-ui-decorated-window-242:0.25.0")
}

compose.desktop {
    application {
        mainClass = "MainKt"

        nativeDistributions {
            targetFormats(TargetFormat.Dmg, TargetFormat.Msi, TargetFormat.Deb)
            packageName = "TechAlert"
            packageVersion = "1.0.0"
        }
    }
}
