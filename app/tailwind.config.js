/** @type {import('tailwindcss').Config} */
module.exports = {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  content: [],
  darkMode: true,
  theme: {
    extend: {
      colors: {
        'ghdark': '#0d1117',
      }
      // => @media (min-width: 1536px) { ... }
    },
  },
  daisyui: {
    // themes: [
    //   {
    //     mytheme: {

    //       "primary": "#1e40af",

    //       "secondary": "#e0f2fe",

    //       "accent": "#60a5fa",

    //       "neutral": "#3b82f6",

    //       "base-100": "#2563eb",

    //       "info": "#2463eb",

    //       "success": "#16a249",

    //       "warning": "#db7706",

    //       "error": "#dc2828",
    //     },
    //   },
    // ],

    themes: [
      "light",
      "dark",
      "cupcake",
      "bumblebee",
      "emerald",
      "corporate",
      "synthwave",
      "retro",
      "cyberpunk",
      "valentine",
      "halloween",
      "garden",
      "forest",
      "aqua",
      "lofi",
      "pastel",
      "fantasy",
      "wireframe",
      "black",
      "luxury",
      "dracula",
      "cmyk",
      "autumn",
      "business",
      "acid",
      "lemonade",
      "night",
      "coffee",
      "winter",
      "dim",
      "nord",
      "sunset",
    ]
  },
  base: false,
  plugins: [require("daisyui")],
}

