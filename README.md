# Project stack-images-md

One link to get tech-stack logos as one image from a single get request.

This can be used inside a markdown image link in order to easily get stack images all in one image file through a single get request. 

Simply follow the url with a hyphen seperated list of techstack names and a line of logos will be sent back. The addition to fuzzy finding of the logos means that any close spelling to the desired techstack will still general the proper logo, e.g. "rect" will still get you the "react" logo. Don't sweat the typos, but if the typos are too far off, or just happens to not be found by the fuzzy search logic then you will get a "404" image spliced into the returned images.

## Current Custimizations:

 - Using the endpoint to preface the stack names with a "row x col" (4x2) creating a grid formation you specify in the qrid endpoint. (see below)

## Future Custimizations:

 - Additional customization to make the logos sepia or black&white, as well as ability to change background color and playing additional images on stop of stack logos <em>(an example would by to allow a box on each logo that can by checked of not in order to say the stack you know but only checke the stack you are using in a given project)</em>


## Getting Started

For markdown

```markdown
![Alt image text for tech stack logos](http://s8wwggk.5.161.48.105.sslip.io/react-go/golang-react)

```

Browser

```markdown
http://localhost:8080/golang-linux

```

## Endpoints

### To get a single line of stack logos

<em>e.g. below react logo followed by golang logo</em>
```txt
http://localhost:8080/golang-linux

```

#### Single Line Output:
![Alt image text for tech stack logos](https://stackimages.xyz/golang-linux)

### To get stack logos in defined grid layout

e.g. six linux logos in a 3 row, 2 column grid


```txt 
https://stackimages.xyz/2x3/linux-linux-linux-linux-linux-linux

```  
#### Grid Output:
![six linux logos in a 2 by 3 grid](https://stackimages.xyz/3x2/linux-linux-linux-linux-linux-linux)
