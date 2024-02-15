---
title: 'Serving Images and Other Media From a CDN in NextJS'
date: 2024-01-31
author: 'Brandon McKimmons'
description: 'Switching to serving images and other media from a CDN and where the real gains come from.'
image: '/static/images/blog/image_serve_CDN.webp'
---

# Serving Images and Other Media From a CDN in NextJS

## Switching to serving images and other media from a CDN and where the real gains come from

In my journey creating this website and blog I was having real issues with image load times. Especially the photos in my image sliders.
There was a 5-6 second delay before the images would load. Granted, they were pretty high resolution images, but still, that's a long time to wait for an image to load.
Since CDN's are pretty much the standard on any kind of project of size or significance I decided to look into using one.

## What is a CDN?

For those laypeople out there, a CDN is a Content Delivery Network. It's a network of servers that are distributed across the globe. They are used to serve content to users based on their geographic location. 
This means that the content is served from a server that is closest to the user. This results in faster load times and a better user experience.

<img
  src='/static/images/blog/imgix-logo.webp'
  alt='TailwindCSS Logo'
/>

After doing some initial research I was able to find a free option that would work for me. [Imgix](https://www.imgix.com/) is a service that will host your images and serve them from a CDN.
They have a free tier that includes 1,000 origin images, infinite transformations, included bandwidth, two unique sources and two unique users. Which for my
use case was absolutely perfect. I signed up for an account and started the process of moving my images over to Imgix. In their setup guide they require a public facing
URL for your images. Since NextJS doesn't expose the images in the public folder by default I had to go another route. 

## Akamai Object Storage

Luckily, [Linode](https://www.linode.com/) has a [Object Storage](https://www.linode.com/lp/object-storage/?promo=sitelin100-02162023&promo_value=100&promo_length=60&utm_source=google&utm_medium=cpc&utm_campaign=builder_pilot&utm_term=g_aud-2187602242764:kwd-959451010310_e_linode%20object%20storage%20pricing&utm_content=675097361193&gad_source=1&gclid=CjwKCAiA_OetBhAtEiwAPTeQZ7nYbOZoajK7ACeJa0ApcO8-D-TUtf5KBvVLieBK60U8PnZ-u8MWWRoCOzYQAvD_BwE) 
service that I was able to use to host my images. I created a bucket and uploaded all of my images to it. I will say here that initially I didn't make a seperate folder for the images and so I had to go back and do that later on.
Which was kind of a hassle and super tedious because that's when I changed all the filenames to be descriptive instead of the generic names my camera gives them. 
However, once that was done I was able to continue with the Imgix setup.

Now with my Imgix address in hand I was able to start the process of moving my images over to Imgix in the code. I started with the images in my image slider. After looking through the Imgix API, I was able to
find some options that would be best to use. I also saw that it would probably be best to use the [@imgix/js-core](https://github.com/imgix/js-core) library to help with things.
For a more detailed rundown of the complete process please refer to the previous [post](https://brandonmckimmons.com/blog/image-slider) about the image slider's creation.

After I had the photos in the image slider being served from Imgix I was able to see the difference in load times. The images were faster but not as fast as I was expecting.
This led me down a rabbit hole of research and testing. I did find that Imgix was serving the images properly sized for all the various devices. So, they were much smaller on 
mobile devices and larger on desktops. This was a massive improvement. I also found that Imgix was degrading the quality of the images as needed. Still, though with those and other optimizations load times were still more than a couple of seconds.

## Where the real gains come from

In the end the biggest improvements came from two completely different sources. First, in my falling down the rabbit hole I was able to find out that my [NGINX](https://www.nginx.com/) config file wasn't using TLS1.2 or TLS1.3.
This alone made a huge difference in load times. The second improvement came from actually upgrading my [Linode](https://www.linode.com/) server. I was using the cheapest server they had, the nanode and it was just not able to keep up with the demand.
After upgrading to the next tier up, the Linode 2GB, I was able to see the images load in less than a second. This was a massive improvement and I was very happy with the results. Now, this was on my gigabit ethernet and so actual load times will vary
depending on the user's internet connection. Still, I was very happy with the results and I'm glad that I was able to get the images loading faster.

You can do all the optimizations you want but if the server can't keep up then it's all for naught. I'm glad that I was able to find the real source of the problem and fix it. I'm also glad that I was able to get some experience with a CDN and see some of the benefits it provides. 
This whole process took more than a few days. From the initial research to the final implementation it was a lot of work. From SEO optimization to image optimization to server optimization. Needless to say, there was a lot of optimizing and an equal amount of coffee involved.
