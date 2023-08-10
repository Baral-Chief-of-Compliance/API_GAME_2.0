import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

export default defineNuxtPlugin((nuxtApp) => {
    const vuetify = createVuetify({
        ssr: true, //optional // detect if ssr is used
        components,
        directives
    })
    nuxtApp.vueApp.use(vuetify)
})