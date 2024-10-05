/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/app/**/*.{js,ts,jsx,tsx}",
    "./src/components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        customGray: "#E2E8F0",
        customBorderGray: "#CBD5E1",
        customBgGray: "#F8FAFC",
      },
      boxShadow: {
        customBlockShadow:
          "0px 4px 8px -2px rgba(23, 23, 23, 0.10), 0px 2px 4px -2px rgba(23, 23, 23, 0.06)",
        customBlockShadowHover:
          "0px 10px 15px -3px rgba(23, 23, 23, 0.15), 6px 8px 10px -2px rgba(23, 23, 23, 0.10)",
      },
      gridTemplateColumns: {
        "20": "repeat(20, minmax(0, 1fr))",
      },
    },
  },
  plugins: [],
};
