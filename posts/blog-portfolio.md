---
title: 'Building a Blog and Portfolio Website'
date: 2024-01-18
author: 'Brandon McKimmons'
description: 'The process and tools I chose to build this website after finding out that yet another open source project died too young.'
image: '/static/images/blog/building-blog.webp'
---

# Building a Blog & Portfolio

## The process and tools I chose to build this website.

The modern web development tool landscape for building a website is vast. There are so many options to choose from, it can be overwhelming.
Many Javascript frameworks have come and gone.

Nowadays you can choose between [React](https://react.dev/), [Vue](https://vuejs.org/), [Angular](https://angular.io/), [Svelte](https://svelte.dev/), [NextJS](https://nextjs.org/) and many more.
There's even things like [htmx](https://htmx.org/) which skirt the line between a framework and a library. It calls itself a library and I'm struggling to differentiate between the two myself, lol.
Then you have to choose between a static site generator or a server-side rendered application.

<img 
  src="/static/images/blog/Nextjs-logo.webp" 
  alt="NextJS Logo"
  class="w-full items-center"
/>

## Why I chose to use NextJS

Initially, I wanted to use [GatsbyJS](https://www.gatsbyjs.com/). I used it before and was familiar with it. I knew of NextJS as well, but never used it.
After doing some research, I decided to go with NextJS since Gatsby has basically been abandoned by the maintainers. Which while sad,
does seem to happen quite a lot with open source projects.

<img 
  src="/static/images/blog/crying-mike-jordan.webp" 
  alt="Crying Michael Jordan Meme & GatsbyJS is Dead" 
/>

The main thing I liked about Gatsby and now also NextJS was that it would give me the ability to use static site generation. It would also allow
me to use React, which I am very familiar with as well. Since React is probably the most widely used Javascript framework
(I can hear the React fanboi's screaming, "It's not a framework!") it's a wise choice to get accustomed to if I'm gunna look for a job in the industry at some point.

The ability to use Markdown to write my blog posts was also a huge plus. While it can be a bit wonky at times, it's still a lot easier than writing HTML
for every blog post. I also like that I can use React components in my Markdown files. This allows me to use things like PrismJS to highlight code blocks in the future.

## TailwindCSS and Component Libraries

<img 
  src="/static/images/blog/Tailwind_CSS_logo.webp" 
  alt="TailwindCSS Logo" 
/>

This would be my first time using [TailwindCSS](https://tailwindcss.com/). Giving even a cursory look at web dev YouTube channels, you'll see that Tailwind is often referenced. So I figured why not?
I'm glad I did. The time I've spent using it so far it's the first time I've felt CSS was intuitive. Of course, I've used things like Bootstrap and Material UI in the past but while
they are great, they still felt like I was fighting with CSS.

Then there's Tailwind components, [TailwindUI](https://tailwindui.com/). I did grab the navbar from their library,
it's free, well the one I used was anyways. I also gleaned some inspiration from the code examples they do share for other parts of my site. Looking back I probably would've went with something like
[ShadCN](https://ui.shadcn.com/) as that's completely free and really looks like a joy to use. Ah well, maybe next time.

## Good Ol' Linode, my best good buddy

<img
  src="/static/images/blog/Linode_updated_logo.webp" 
  alt="Linode Logo" 
/>

While everyone seems to be using the [Vercel](https://vercel.com/) platform for their NextJS apps, I decided to go with [Linode](https://www.linode.com/). I decided to be a bit of a boomer and go with a VPS. Why? Maybe I like pain.
Maybe I'm a masochist. Maybe I'm just a cheap bastard. Who knows? I do know that I like having full control over my server. I also like that I can host multiple sites on it. Also,
learning or relearning how to setup a server is a good skill to have. I'm not saying that Vercel is bad, it's obviously really good and easy. I just wanted to do it the hard way I guess.

## Conclusion

No thrilling conclusion here. Nothing to really knock your socks off at the end. Nothing to make you go, "Wow, I'm glad I read this." I just wanted to write a blog post about the tools I used to build this site.
If only to fill out the page so the footer would display properly without using h-screen. For now that's all. I hope you enjoyed this post. If you did, please consider sharing it with your friends. Just, please be nice!
