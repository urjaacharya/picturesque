# picturesque

picturesque is a favicon generator meant to be used with static website generator (for eg. Hugo). It generates favicons of specified sizes and generates HTML script (goes in `head` eventually) with the favicon declarations along with the site webmanifest file.

## Build from source

1. Get the repository:
`go get -u github.com/urjaacharya/picturesque`
2. Navigate to the root of the repo you downloaded above (i.e picturesque) and run `go build`
    
## Usage

### Arguments JSON

Specify the input arguments in a `.json` file. Following is the sample json:
```
{
  "input_image": <Path to the input image>,
  "output": {
    "images_path": <Path to save output images to>,
    "html": {
      "path": <Path to save output HTML script>,
      "name": <Name of output HTML script>
    }
  },
  "link": {
    "href_prefix": <HREF_PREFIX>,
    "href_suffix": <HREF_SUFFIX>
  },
  "site_webmanifest": {
    "background_color": <hexcode of background color to be used in webmanifest file>,
    "name": <site name used in webmanifest file>,
    "short_name": <short name for site used in webmanifest file>,
    "theme_color": <hexcode of site theme color to be used in webmanifest file>
  },
  "icons": {
    <Name of output favicon>: { "width": <width of outout favicon>, "height": <height of output favicon>, "rel": <Array of rel attributes> },
  }
}
```
In the above json, `HREF_PREFIX` and `HREF_SUFFIX` are prefix and suffix that can be added to location of each favicon. For example, if these both are empty strings, the `link` tag for each favicon would look something like the following:
```
<link rel="icon" type="image/png" href="myicon.png"/>
```
If you are planning to use the generated HTML script for a Hugo project, the `link` tags for each favicon would look like the following:
```
<link rel="icon" type="image/png" href="{{ `favicon-32.png` | relURL }}"/>
```
Therefore, to add these extra information in addition to favicon location, use `HREF_PREFIX` as ``"{{ `"`` and `HREF_SUFFIX` as ``"` | relURL }}"``. 
An example of sample json is also provided in `data/sample-input.json`.

### Running `picturesque` in commandline

#### For Linux/Unix system:

```
./picturesque --inputArgs data/sample-input.json
```

#### For Windows system:

```
./picturesque.exe --inputArgs data/sample-input.json
```

