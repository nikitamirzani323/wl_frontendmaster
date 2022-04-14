module.exports = {
  content: ["./src/**/*.{html,js,svelte}"],
  theme: {
    extend: {},
  },
  plugins: [require('tailwind-scrollbar'),require("daisyui"),require("tailwindcss-animation-delay")],
  daisyui: {
    themes: [
      {
        light:{
          ...require("daisyui/src/colors/themes")["[data-theme=light]"],
          primary: "#1a73e8",
        }
      }
    ],
  },
}