/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
      "./templates/**/*.html",
      "./templates/components/*.html",
      "./templates/pages/*.html",
      "./templates/layout/*.html",



      './node_modules/flowbite/**/*.js'
  ],
    theme: {
      extend: {},
    },
    plugins: [
      require('flowbite/plugin')
  ]
  }

