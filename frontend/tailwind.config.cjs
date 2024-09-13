/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,js,svelte}"],
  theme: {
    extend: {
      fontFamily: {
        poppins: ['Poppins']
      },
    },
  },
  plugins: [require("daisyui")],
}
