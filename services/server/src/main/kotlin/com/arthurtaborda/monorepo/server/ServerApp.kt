package com.arthurtaborda.monorepo.server

import com.arthurtaborda.monorepo.common.Commons.tellMeWhoWeAre
import com.arthurtaborda.monorepo.logging.Logger.info

object ServerApp {
    @JvmStatic
    fun main(args: Array<String>) {
        info("Hi, I will be the server when I will grow up!")
        info("For now I know our commons: " + tellMeWhoWeAre())
    }
}
