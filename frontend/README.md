# sv

Everything you need to build a Svelte project, powered by [`sv`](https://github.com/sveltejs/cli).

## Creating a project

If you're seeing this, you've probably already done this step. Congrats!

```sh
# create a new project in the current directory
npx sv create

# create a new project in my-app
npx sv create my-app
```

## Developing

Once you've created a project and installed dependencies with `npm install` (or `pnpm install` or `yarn`), start a development server:

```sh
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

## Building

To create a production version of your app:

```sh
npm run build
```

You can preview the production build with `npm run preview`.

> To deploy your app, you may need to install an [adapter](https://svelte.dev/docs/kit/adapters) for your target environment.


UX prompt

Product Requirement: AI SaaS Software - Homepage

Purpose: To provide users with an engaging and informative overview of our AI software offerings, showcasing features and benefits to encourage exploration and engagement.

UI Components: 
- Navigation Bar: Global navigation for product sections; includes links to features, pricing, and contact.
- Hero Section: Prominent feature area with a headline, subheadline, and call-to-action button.
- Feature Highlights: Grid or list of key features with icons and brief descriptions.
- Testimonials: Carousel or list view featuring customer reviews and ratings.
- Footer: Additional links, contact information, and social media icons.

Visual Style: 
- Theme: Dark theme with optional light mode
- Primary color: Indigo #6366F1
- Secondary color: Purple #8B5CF6
- Accent color: Cyan #06B6D4
- Error/Alert: Red #DF3F40
- Spacing: Consistent 20px outer padding, 16px gutter spacing between items
- Borders: 1px solid light gray #E3E6EA on cards and input fields; slightly rounded corners (6px radius)
- Typography: Sans-serif, medium font weight (500) for headings, regular (400) for body, base size 16px
- Icons/images: Simple, filled vector icons for navigation and actions; illustrative flat images used occasionally for empty states