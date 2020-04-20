# picturesque
picturesque is a tool that creates favicons of the sizes you specify and generates HTML script with the favicon declarations along with the site webmanifest file. 

**Usage**

Specify the input arguments in a `.json` file. Following is the sample json:
```
{
  "input_image": <Path to the input image>,
  "output": {
    "images_path": <Path to save output favicons to>,
    "html": {
      "path": <Path to save output HTML script>,
      "name": <Name of output HTML script>
    }
  },
  "link": {
    "href_prefix": <content to be appended before location of all favicons>,
    "href_suffix": <content to be appended after location of all favicons>
  },
  "site_webmanifest": {
    "background_color": <hexcode of background color to be used in webmanifest file>,
    "name": <name used in webmanifest file>,
    "short_name": <name used in webmanifest file>,
    "theme_color": <hexcode of theme color to be used in webmanifest file>
  },
  "icons": {
    <Name of output favicon>: { "width": <width of outout favicon>, "height": <height of output favicon>, "rel": <Array of rel attributes> },
  }
}
```
An example of sample json is also provided in `data/sample-input.json`

1. Run directly using binary executables:
    -  The binary files for *nix and Windows system are in `bin` folder.

        *For Linux/Unix system:*

        Run the following from the terminal:
`./picturesque --inputArgs data/sample-input.json`

        *For Windows system:*

        Run the following from the terminal:
`./picturesque.exe --inputArgs data/sample-input.json`

2. Generate executables and run it:
    - Get the repository:
`go get -u github.com/urjaacharya/picturesque`
    - Navigate to the root of the repo you downloaded above (i.e picturesque). Now, build the repository: 
`go build`
    - Now, run the generated executables following the steps listed in 1.
