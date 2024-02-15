---
title: 'Creating an Image Slider with NextJS and TailwindCSS'
date: 2024-01-19
author: 'Brandon McKimmons'
description: 'A short story of the trials and tribulations of creating an image slider in NextJS with TailwindCSS.'
image: '/static/images/blog/image_slider.webp'
---

# Creating an Image Slider with NextJS and TailwindCSS

## A short story of the trials and tribulations of creating an image slider in NextJS with TailwindCSS.

As I mentioned in my previous post about the tools I chose to build this site and why, looking back it would've been a wise
choice to go with [ShadCN](https://ui.shadcn.com/) instead of hand coding everything right off the dome with Tailwind.
However, I made my bed and now it was time to lay it.

I had a pretty good idea of what I wanted the image slider to look like. I wanted a fully responsive slider that would
display each image in the biggest size for each device. I also wanted to be able to click on the image and have it open
to full screen, displaying the image in all it's glory. With those requirements in mind, I set out to build it.

## Github Copilot, he loves to fly and it codes

<img 
  src="/static/images/blog/GitHub_Copilot_logo.webp" 
  alt="TailwindCSS Logo" 
/>

It should come as no surprise that I used [Github Copilot](https://github.com/features/copilot) to help me build this. In general, there are a few issues with using
Copilot or [ChatGPT](https://chat.openai.com/) to generate code. The first is that it's not always correct. The second is that it's not always the best way to do things.
But most importantly, since it's knowledge stopped somewhere back in 2021 I believe, it doesn't know about the latest version of just about anything
in the web development world. So, it's not going to be able to help you with the latest version of [NextJS](https://nextjs.org/) or [TailwindCSS](https://tailwindcss.com/). That can be a problem.

However, it can be a pretty good starting point and really where it shines is asking it to explain things to you. For example, I asked it to explain
just about every error that gets flagged or that [NPM](https://www.npmjs.com/) throws at me. Although, I must say that the error messages from NPM and more so NextJS are really good.
They even give you a link to the documentation that explains the error and how to fix it. Which is really nice and something I would've killed for in the past.

## The trials and tribulations of the image slider

<img 
  src="/static/images/blog/Heroicons.webp" 
  alt="TailwindCSS Logo" 
/>

So, I began by asking Copilot to give me the code to create and image slider using the same JSON file that I had setup to create the image cards on the home page.
Out came it's result and I went to refining it and making it look how I wanted it to. I imported [Heroicons](https://heroicons.com/) to the page since I used that
for the navbar as well. The ole left and right arrows in tow I went about placing them in the center of each image and on their respective sides. Then I started playing
with different colors to see what would look best. White was the obvious choice but I did try some browns and blues but they tended to blend in with the images too much.
So, I went with white and added a hover effect that changed the opacity to show that it was clickable.

The real issue came with the responsive nature of the image slider. Whenever the image sizes changed, depending on the devices the arrows would move out from over the images.
At the time I had another "card" as the background and it used the color white as well. This was obviously a problem. I really spent too much time on this issue. It really annoyed
me that for one I couldn't get the arrows to stay over the image and two that essentially the white arrows would disappear into the white background. I tried many different things
to fix this. I tried changing the color of the arrows like I mentioned before. I watched [YouTube](https://www.youtube.com/) after YouTube video on image sliders.

Don't get me started on web dev YouTube videos either.
The big secret about those videos is that they show very specific use cases and they don't really explain how to do things. They just show you how they did it. Which is fine but the real
clincher on why I didn't like their videos is they all use images of the same ratio. Like, sure someone is gunna have an image slider with only images of the same resolution and dimensions, lol.
This is actually a problem I noticed with game dev videos as well. They show you how they did things, for their specific use case and it's never the way you'd actually do things. Making the entire
video pointless in my opinion.

## The Solution

Actually, the solution was pretty simple and I'm sure many of you are screaming it at your screen right now. I just needed to change the background color to the image slider from white to really
anything else. I ended up going with a color pallette that I found on what I could've sworn was an article on Davey & Krista but clicking that link
yields an, "Oops! It looks like this page is missing." message. However, looking at the little Google blurb for the link it does mention something about Benjamin Moore's 2023 color of the year.
Which was Raspberry Blush and I distinctly remember that color as being from the site I got my palette from. Because I tried using that color along with the rest of the palette it belonged to but ended up
going with another. Anyways, I digress.

With the color palette in hand and in use on the site you're reading this on, I was able to get the image slider to look how I wanted it to. With the very light brown as the card's background the white arrows
showed up perfectly no matter what the resolution of the device it was being viewed on. The next step was to display the images in their full glory when clicked on. I came up with this approach to get it done:


<pre>
<code class="language-typescript">
const images = [
    {
      src: 'Astrophotography/Blue_Orion_Shot1',
      alt: 'The Orion constellation with a blue hue, 1st photo',
    },
    {
      src: 'Astrophotography/Blue_Starry_Night_Shot1',
      alt: 'A random shot of a section of the night sky with a blue hue',
    },
    {
      src: 'Astrophotography/Purple_Orion_Shot1',
      alt: 'The Orion constellation with a purple hue, 1st photo',
    },
    {
      src: 'Astrophotography/Purple_Orion_Shot2',
      alt: 'The Orion constellation with a purple hue, 2nd photo',
    },
    {
      src: 'Astrophotography/Purple_Orion_Shot3',
      alt: 'The Orion constellation with a purple hue, 3rd photo',
    },
    {
      src: 'Astrophotography/Starry_Night_Sky_with_Clouds_Shot1',
      alt: 'A colorful photo of a cloudy night sky, 1st photo',
    },
    {
      src: 'Astrophotography/Starry_Night_Sky_with_Clouds_Shot2',
      alt: 'A colorful photo of a cloudy night sky, 2nd photo',
    },
    {
      src: 'Astrophotography/Starry_Night_Sky_with_Clouds_Shot3',
      alt: 'A colorful photo of a cloudy night sky, 3rd photo',
    },
    {
      src: 'Astrophotography/Starry_Night_Sky_with_Clouds_Shot4',
      alt: 'A colorful photo of a cloudy night sky, 4th photo',
    },
    {
      src: 'Astrophotography/Night_Streetlight_Gate',
      alt: 'A photo of a street light behind a gate with the night sky in the background',
    },
  ];

  const imgixClient = new ImgixClient({
    domain: 'brandonmckimmons-nextjs-563476088.imgix.net',
  });

  const imgUrl = (image: { src: string; alt: string }) =>
    imgixClient.buildURL(`${image.src}.webp`, {
      fit: 'fill', // fill mode
      auto: 'format,compress', // auto format and compress
      lossless: 1,
      // ... other Imgix parameters
    });
</code>
</pre>


I created an array of the images with both a `src` and an `alt` for SEO purposes and then created a function to build the URL for the images using the `imgixClient` helper to configure some default settings.
`imgixClient` is a helper that I got from the [Imgix](https://www.imgix.com/) documentation that recommends to use [@imgix/js-core](https://www.npmjs.com/package/@imgix/js-core) to generate the image URLs.

I'll talk more about Imgix in another post but for now I'll just say that it's a really cool service that I'm using to serve the images on this site. It's really easy to setup and they do have a free tier which I'm using.
Also, don't worry I have since changed the Imgix domain to something else and am using secured tokens so there's no funny business.

Anyways, initially I hadn't set a `src` and an `alt` for each image and just had one property which was the original generic file name my camera spits out.
This is something Imgix specifically advises against doing, again for SEO purposes. While going back and making these changes wasn't too much of a headache it still was annoying.
It may have ended up being the same amount of work in the end but I would've preferred to do all the original, descriptive adjustments in the first place instead of having to go back and revise everything.
Once you get in the groove though, I personally just want to get the thing working and so at the time was okay with going back and revising or so I thought at the time, lol.

At any rate, I then used the `imgUrl` function to build the URL for the photos in the image slider. To show the rest of my process more thoroughly I'll show the code for how each
photo in the image slider is selected:

<pre>
<code class="language-typescript">
const [currentImageIndex, setCurrentImageIndex] = useState(0);

  const nextImage = () => {
    if (images) {
      setCurrentImageIndex((currentImageIndex + 1) % images.length);
    }
  };

  const prevImage = () => {
    if (images) {
      setCurrentImageIndex(
        (currentImageIndex - 1 + images.length) % images.length
      );
    }
  };
</code>
</pre>

Here I'm using the `useState` hook to set the current image index and then using the `nextImage` and `prevImage` functions to change the current image index.
Pretty straightforward stuff. Now I'll show the code for how the images are displayed:

<pre>
<code class="language-typescript">
Image
  className='object-cover max-h-svh max-w-min px-3 py-3 z-10'
  src={imgUrl(images[currentImageIndex])}
  alt={images[currentImageIndex].alt}
  sizes='(min-width: 1280px) 1256px, (min-width: 1040px) 744px, (min-width: 780px) 648px, calc(100vw - 24px)'
  style={{
    objectFit: 'contain',
  }}
  width={7028}
  height={4688}
  onClick={toggleModal}
  placeholder='blur'
  blurDataURL={
    'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNkqAcAAIUAgUW0RjgAAAAASUVORK5CYII='
  }
  priority
</code>
</pre>

Then here I'm using the `Image` component from NextJS to display the images. I'm using the `imgUrl` function to get the URL for the image and then using `currentImageIndex` to select the current image.
I got the proper `sizes` vaule using the [RespImageLint](https://ausi.github.io/respimagelint/) tool.
I've only had one instance that the sizes it recommended weren't correct or perhaps in my specific use case it wasn't displaying the images how I wanted it. Overall though I think you can trust the values it gives you.
I'm also setting the `priority` as well as the `width` and `height` properties to the original dimensions of the first photo the image slider will display. This is important to accomidate for the LCP (Largest Contentful Paint) metric
and potential layout shift.

These are things, again, too help with your lighthouse scores and so potentially your SEO. I'm also using the `placeholder` and `blurDataURL` properties that I'd hoped would display a blurred version of the image while it's loading.
However, I couldn't get this to work and so I'm not sure if it's a bug or if I'm doing something wrong and I did try using [Plaiceholder](https://plaiceholder.co/docs) which NextJS recommends. I'll have to look into it more later. 
I'm also using the `onClick` property to open a modal to display the image in full screen. This is done using the `useState` hook again this time to set the modal state and then using the `toggleModal` function to change the modal state.

<pre>
<code class="language-typescript">
const [isModalOpen, setIsModalOpen] = useState<boolean>(false);

  const toggleModal = () => {
    setIsModalOpen(!isModalOpen);
  };
</code>
</pre>

## This Is The End

That's pretty much it. There was some adjusting with making things fully responsive but it now displays fullscreen on mobile devices as well as on the desktop.
Which is exactly what I wanted. You can watch a million YouTube videos but at the end of the day it just ends up being way better if you try coming up with the
solution yourself. You'll learn a lot more and you'll be able to solve problems on your own. Which is really what being a developer is all about. Solving problems.

Yes, I did use Copilot but I'm not at all embarrassed about that. AI is here to stay and either you're going to embrace the new tools we have or you aren't. It's
a personal choice and if you can get by without it then more power to you. But I like new and useful tools. I like that NPM and NextJS tell me exactly what the issue
is and gives me the link to solve it. These are good things. At some point I'd like to setup a local, open source AI and train that for my specific use cases. But that's
a story for another day.
