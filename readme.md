# NAWST
## Nicey's Adequate Web Statistics Tool

A simple tool that tracks visitor numbers for my sites in a privacy friendly way.

### What is sent?

``` json
{
    "tz": "Your Timezone (will be used for appx location, i.e. country)",
    "dm": "Domain you visited",
    "pt": "Page on the domain you visited",
    "lw": "Have you visited this domain this week?",
    "vb": "Have you visited this domain at all?"
}
```

NAWST stores a single value in the local storage of your browser,
which is the date you last visited this site.