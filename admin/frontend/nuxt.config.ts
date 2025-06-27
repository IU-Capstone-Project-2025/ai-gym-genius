import tailwindcss from "@tailwindcss/vite";

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    compatibilityDate: '2025-06-25',
    modules: ["@nuxt/ui", "@nuxtjs/storybook"],
    css: ['@/assets/css/main.css'],
    vite: {
        plugins: [
            tailwindcss(),
        ],
    },
    ui: {
        fonts: false
    }
})