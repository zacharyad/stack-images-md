# Project stack-images-md
( !!! work in progress getting more logos and some other features !!! )

One Link to Fetch Tech-Stack Logos in a Single GET Request

## Overview

Stack-Images-MD is a self-hosted solution, coded in Go, that enables users to easily obtain a composite image of tech-stack logos through a single GET request. This service is SSL-certified, ensuring security usage. Ideal for enhancing markdown documentation, Stack-Images-MD simplifies the process of displaying a series of technology logos with minimal effort.

## Features

### Current Customizations

- **Grid Formation**: Specify a grid layout by prefixing the stack names with a "row x col" (e.g., 4x2). This creates a customizable grid of logos (usage show below)

### Future Customizations
- **images for light or dark mode options**: specify that you will have a light/dark or mixed background and images best for those will be selected.
- **Image Filters**: Options to render logos in sepia or black & white.
- **Background Customization**: Ability to change the background color of the logo images.
- **Overlay Additional Images**: For example, a checkbox overlay on each logo to indicate familiarity or usage in a specific project.
- **Caching and Stack names**: Less server load with caching and ability to type a common tech stack to get a prebuild image of all logos of that give tech stack

## Self-Hosting and Security

Stack-Images-MD is self-hosted with custom domain (stackimages.xyz). The platform is SSL-certified, ensuring secure communication over custom domains. By leveraging Go's performance and reliability, Stack-Images-MD delivers a robust and efficient solution for tech-stack logo delivery through http.

## Getting Started

To integrate Stack-Images-MD into your markdown documentation or access it via a browser, follow the instructions below.


## Getting Started

For markdown

```markdown
![Alt image text for tech stack logos](http://stackimages.xyz/l/react-javascript-node)

```

Browser

```markdown
http://stackimages.xyz/l/react-javascript-node

```

## Endpoints

### To get a single line of stack logos

<em>e.g. below linux, golang, and javascipt logo in a row</em>
```txt
http://stackimages.xyz/l/react-javascript-node

```

#### Single Line Output:
![Alt image text for tech stack logos](https://stackimages.xyz/l/react-javascript-node)

### To get stack logos in defined grid layout

e.g. six linux logos in a 3 row, 4 column grid


```txt 
[stack logos in a 3 by 5 4 grid](https://www.stackimages.xyz/l/3x4/node-dart-d3-django-sequelize-node-deno-css-react-ember-elm-grunt)

```  
#### Grid Output:
![stack logos in a 3 x 4 grid](https://www.stackimages.xyz/l/3x4/node-dart-d3-django-sequelize-node-deno-css-react-ember-elm-grunt)

#### How to Help
##### Contribute images of stack logos by:

- Create logo image that is a png file, clear background, 300x300, and named with the stack technologies name (no hyphens, spaces, or capitol letters as this is what the applicaton draws from when selecting inside URL).
- Fork the repo, add the image(s), and create a pull request.
- Wait for review. Once approved all I have to do is redeploy and your new logo will be available to all.
  
##
### Conclusion
##

Stack-Images-MD provides a seamless way to incorporate a line or grid of tech-stack logos into your markdown documentation or web pages. Show off any projects tech stack logos with ease and professionalism using Stack-Images-MD.
