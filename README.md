# Building my personal website
My personal website serves as an overview of my academic (or otherwise software related) work.

## Inspiration
I looked at the websites of former supervisors:
- [Burcu Kulahcioglu Ozkan](https://burcuku.github.io/home/): I like this website a lot, especially the news section. 
  I decided to also have a news section on the front page: it is flexible enough for lots of different content, and gives an idea of what I have been doing most recently.
- [Soham Chakraborty](https://www.st.ewi.tudelft.nl/sschakraborty/): I like the table for displaying publications with dates. I will use that for the news section too.

## Content
The most important is to give an overview of:
- My current affiliation
- What I am working on
- My research interests
- Recent news related to my research (e.g. publications accepted, awards received)
- Links to my profiles on other platforms (primarily my TU/e profile page, email, orcid, Google Scholar, GitHub, LinkedIn)

Next to that, I would also like to provide:
- My CV
- A list of my publications
- Blog posts, if I can get into the habit of writing them.

Here is the basic design for the front page:

![](mockup.jpg)

## How *Not* To Build It 
Initially I built a very simple website using just HTML and [simplecss](https://simplecss.org/), but that has limitations:
- I have to write all the HTML manually. For longer bits of content, I would prefer something like markdown.
- I would like to programmatically generate some of the content, such as a list of publications.

I tried [Jekyll](https://jekyllrb.com/), which takes markdown as input and integrates nicely with GitHub pages.
I am sure it could be massaged to make it work the way I want, but I struggled to get there:
- Jekyll is a big project with lots of plugins, themes etc. I do not really understand how to get it to do what I want.
- It is implemented in Ruby, which I am not so comfortable with. Writing some custom code for it would be difficult for me.

## How To Build It
My website does not have to be very complicated.
I reckon that with a bit of code and some HTML/CSS, I can build it *almost* from scratch.
I know enough HTML to put it together, my one concern would be to have a CSS setup that works for different screen sizes and looks reasonably pretty.

I will write a small program to take in the following files, and from that create `index.html`:
- `index.template.html` for the basic structure of the website
- `config.json` for small bits of text and other configuration. 
- `about_me.md` for the text next to the profile picture.
- `news/*.md` news items. The format for files is `<year>-<month>-<date>-<name>.md`, e.g. `2023-08-30-defended-msc-thesis.md`.

### Settings for `config.json`
- URL: `wildarch.dev`
- profile picture
- platform links
- page title
- meta: SEO keywords
- meta: description
- meta: author name
- meta: author first name
- meta: author last name
- meta: username

For the first version, the CV can be a PDF that is manually built from https://github.com/wildarch/cv.
The list of publications can refer to google scholar.

### Recommended html metadata
See also:
- https://www.w3schools.com/tags/tag_meta.asp
- https://css-tricks.com/essential-meta-tags-social-media/

```html
<!-- Standard metadata -->
<meta name="keywords" content="<keyword 1>, <keyword 2>">
<meta name="description" content="<description>">
<meta name="author" content="<name>">

<!-- Facebook/OpenGraph -->
<meta property="og:title" content="<title>" />
<meta property="og:type" content="profile" />
<meta property="og:profile:first_name" content="<first name>" />
<meta property="og:profile:last_name" content="<last name>" />
<meta property="og:profile:username" content="<username>" />
<meta property="og:url" content="<url to my webpage>" />
<meta property="og:image" content="<profile picture url>" />

<!-- Twitter (reuses OpenGraph tags) -->
<meta name="twitter:card" content="summary">

<meta name="viewport" content="width=device-width, initial-scale=1.0">
```

Testing tools:
- Facebook: https://developers.facebook.com/tools/debug/
- Twitter: https://cards-dev.twitter.com/validator

**Note: The twitter validator requires me to login, but I don't have twitter, so I guess I won't support twitter.**

### Fonts
See also:
- https://www.w3schools.com/css/css_font_pairings.asp
- https://github.com/kevquirk/simple.css/blob/main/simple.css

I like the simplecss solution which does not require loading additional fonts, and I think it looks pretty good.
Other options would be:
- Helvetica and Garamond (web safe fonts)
- Merriweather and Open Sans (needs Google fonts)

### Tools
To implement the site generator, I need libraries for:
- Rendering markdown to HTML
- Rendering template files (for `index.template.html` -> `index.html`)
- parse JSON (config file)
- HTTP/JSON library for retrieving publications, or some specialized library.

I have used `go` for small tools like this before and liked the experience.
It also includes templating, JSON and HTTP in the standard library, so I only need to find a markdown library.
The two big ones appear to be:
- https://github.com/gomarkdown/markdown
- https://github.com/yuin/goldmark

`goldmark` looks more active, so I'll try that first.

## Testing
At a minimum, I want a dead-link checker.