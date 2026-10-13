# Go Koans

A Next.js web project for learning Go through short code questions. Includes a responsive homepage and a first slice exercise with a revealable explanation. Uses TypeScript and Biome for development checks.

## Develop

Use Node.js 24 and npm.

```sh
npm ci
npm run dev
```

Edit `app/page.tsx` to add exercises and `app/globals.css` to change the design. The starter uses system fonts and has no external services or environment variables.

## Check and run a production build

```sh
npm run lint
npm run typecheck
npm run build
npm start
```

## Deploy to Vercel

Import `ssemakov/go-koans` in Vercel and select the `main` branch.

- Framework: **Next.js** (automatically detected).
- Root directory: repository root.
- Node.js: **24.x**.
- Install command: `npm ci`.
- Build command: `npm run build`.
- Output directory: leave the Next.js default.

No custom Vercel configuration is required. Vercel builds the project when you deploy; local production validation does not publish a deployment.
