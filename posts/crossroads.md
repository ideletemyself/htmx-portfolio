---
title: 'Abandoning the CDN and Moving Away From NextJS'
date: 2024-02-05
author: 'Brandon McKimmons'
description: 'Why I decided to abandon the CDN and serve images statically as well as move from NextJS to HTMX and Go.'
image: '/static/images/blog/man_at_crossroads.webp'
---

# Abandoning the CDN and Moving Away From NextJS

## Why I decided to abandon the CDN and serve images statically as well as move from NextJS to HTMX and Go.

To start off I will say that a lot of what I'm going to be talking about will be based off of what I said in my previous [post](https://brandonmckimmons.com/blog/cdn-media) about serving
images from a CDN. I have since changed my mind on that and have decided to serve images statically again. Why is this?
Well, in that previous post I mentioned that I saw the most benefits from actually increasing the RAM and CPU power of my [Linode](https://www.linode.com/) webserver.
Although [Imgix](https://imgix.com/) was loading them quicker, how much quicker was kind of debatable. Perhaps for people in other parts of the world it would be a lot quicker, but for me in the US, it wasn't that much of a difference.

## To swipe or not to swipe with SwiperJS 

<img
  src='/static/images/blog/swiperJS_logo.webp'
  alt='SwiperJS Logo'
/>

In my testing and building I had become aware of a Javascript library called [SwiperJS](https://swiperjs.com/) which allows you to swipe
through your images on mobile or drag them on desktop. It also provides an easy way to show thumbnails of the images in your gallery.
While I love Swiper and what it provides it does have its drawbacks. The biggest one being that it is a bit heavy and can slow down your site.
I spent many an hour trying this and that, scouring the [Github](https://github.com/nolimits4web/swiper/issues) issues for the library and trying to figure out if there was any way to make it faster.
Especially in regards to LCP, [Largest Contentful Paint](https://developer.chrome.com/docs/lighthouse/performance/lighthouse-largest-contentful-paint).

I saw people having the same issues as me and I did take note of the fact some people said it was due to them and potentially myself not optimizing the images properly.
However, I don't think that's the case since [NextJS](https://nextjs.org/) itself optimizes the images for you pretty well. Also, I was loading them as well with Imgix configurations for optimal speed
with things like setting compression and format. I was also taking care to make sure my images were in Webp format and none of them were over 10mb in size.

Still, I was seeing a lot of issues with LCP and Swiper. I decided to go ahead and for testing purposes see if loading the images statically would offer any benefits.
After a little retooling of my code and a few hours of testing I was able to determine that loading the images statically was basically the same as serving them via Imgix.

Again, it may be slightly slower but sometimes it honestly feels like it's pretty much the same speed. Added to that, the LCP score was exactly the same either way.
I know since I'm using the thumbnails module of Swiper that essentially I'm loading many more images at once and therefore that's why I'm getting the increased load times overall.

Maybe looking into doing further deferred loading of the images in some way would help. It could very well be a skill issue on my part. I'm not sure.
If that's the case though, exactly how deep do I want to go down the rabbit hole of trying to figure out how to make Swiper faster?

Another kind of big issue with Imgix or really any other CDN was their bandwidth limits. I have only been using them for a couple of weeks, maybe three weeks at most and I was already 50Gb out of 100Gb used.
If I started getting any kind of traffic that would be a problem. I'd have to pay more for bandwidth and I'm not really interested in doing that. I'd rather just pay for a bigger server if that's the case.

<img
  src='/static/images/blog/plaiceholder_logo@96px.webp'
  alt='Plaiceholder Logo'
/>

I will also mention that I did try getting [Plaiceholder](https://plaiceholder.co/) to work with NextJS but because of how it needs to run on the server side and the amount of images I have to convert
it was monumentally slow. Like omg this is crazy how long this is taking, slow. The fact I couldn't get a blurry placeholder using almost any method and NextJS was really quite annoying and frustrating.

## What NextJS Does Well and What It Doesn't

If all you want to do is prototype quickly and have a site that is pretty fast, NextJS is great. It's easy to use and has a lot of features that make it a great choice for a lot of people.
However, if you happen to have a lot of niche needs and you're trying to do something that's a bit more complex, it can be a bit of a headache. While their documentation is great for getting up and going quickly,
it does lack the kind of in depth explanations that one would want for again, more niche and complex situations.

I know I'm not alone in that feeling either. I've only been using it more than a month
and I've seen a lot of people on [Reddit](https://www.reddit.com/r/nextjs/) and [Stack Overflow](https://stackoverflow.com/questions/tagged/next.js) either asking questions on how to do something or complaining about how it's not working for them.
Often the Stack Overflow questions are left unanswered or are answered with a "I'm having the same issue" type of response. Same with the Github issues.

Like I said, if all you want is to prototype quickly or just toss up a basic website that's not too complex then NextJS is really a great option. I just know for me personally,
at this point in the project I can already see that it's probably not the solution for me going forward.

## HTMX high power tools for HTML and Golang

<img
  src='/static/images/blog/htmx_logo1.webp'
  alt='HTMX Logo'
/>

I've heard a lot about [HTMX](https://htmx.org/) and I've been really impressed with what it can do. It's a great way to add interactivity to your site without having to write a lot of Javascript.
The real benefit I feel though is that you aren't reliant on some framework to work it's magic and leave you kind of confused at how it's accomplishing things or why something isn't working when it should.
That's really the thing that turned me off to web development years ago and it's only been recently that I'm seeing others take note and start looking for other solutions that aren't so esoteric and complex.

Also, having your framework of choice change quite often is really annoying as well. Just when you get used to doing things a certain way, boom it's time to change the entire thing up.
Which usually there's valid reasons for making those changes but it's still not ideal and it's something I feel the Javascript world has been stuck in for almost a decade now.

I won't deny that NextJS does a lot of things for you right out of the box and so I'll have to duplicate some of those things myself. I just feel that in the long run it'll be better
for me to have more control over what's happening and not have to worry about some framework changing things up on me. As well as it will be a far greater
learning experience for me to do things myself and not rely on a framework to do it for me.

<img
  src='/static/images/blog/320px-Go_Logo_Blue.webp'
  alt='HTMX Logo'
/>

Which is what I got back into this whole programming gig for in the first place. I want to make things and make them well and having to resort to just saying, "Well, that's just how it works" is not something I'm willing to do.
Not if I want to make quality products and tools for people to use. [Go](https://golang.org/) seems like a solid choice for a backend server. It's either between that or [Rust](https://www.rust-lang.org/) really.
Honestly, I won't know if I've made a good or bad choice until I've tried it out and seen how it works for me.

## Aaaaaaaanyways

I hope you'll stick around and see how things go. I've decided I'll still be adding a music and video section once I have some stuff to show here on this website. I'll also be continuing
to blog here and will be documenting my progress throughout. I'm excited to see how things go. Thanks for reading and I'll see you in the next one.
