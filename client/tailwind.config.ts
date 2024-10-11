/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/app/**/*.{js,ts,jsx,tsx}",
    "./src/components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontSize: {
        base: "18px",
      },
      colors: {
        customGray: "#E2E8F0",
        customBorderGray: "#CBD5E1",
        customBgGray: "#F8FAFC",
        primary: "#F8FAFC",
        primaryForeground: "#ffffff",
        destructive: "#dc2626",
        destructiveForeground: "#ffffff",
        background: "#f3f4f6",
        accent: "#4b5563",
        accentForeground: "#f9fafb",
        ring: "#3b82f6",
        input: "#9ca3af",
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
