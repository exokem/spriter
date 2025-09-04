# Spriter

Shield: [![CC BY-NC-SA 4.0][cc-by-nc-sa-shield]][cc-by-nc-sa]

## Overview
This is a tool for bulk downloading assets from [The Spriters Resource](https://www.spriters-resource.com/).

This tool provides a scanning feature, which scans the list of available game assets on the website and caches that information so it can be used for searching without HTTP requests.

Scanning is only required if you want to use the search functionality. Searching is only really necessary if you don't know the URL of the game assets you are looking for.

## Usage: Downloading

`spriter get <path>`
> Downloads all of the assets at `https://www.spriters-resource.com<path>`.
> 
> e.g. `/ds_dsi/dgmnworldds/` (must include leading slash) 

## Usage: Scanning

Scanning is required before using search functionality.
These commands scan the site and index information about what is available.

`spriter scan consoles`
> Scans available consoles.

`spriter scan games --meta` `scan -m`
> Fetches and caches game and page counts.

`spriter scan games --all` `scan -a`
> **WARNING**: This command may take longer than 10 minutes to finish.
> 
> Scans all game pages and caches available game information.

`spriter scan games --page <page-number>`
> Fetches and caches all games for a specific page number.

`spriter scan games --pages <first-page-number> <last-page-number>`
> Fetches and caches all games for a specific range of pages (inclusive).

## Usage: Searching (Requires Scanning)
Search the game cache by name or console. Requires scanning to work.

The search algorithm is very basic so it won't work flawlessly all of the time.

`spriter search --console <console>`
> Filters results by console name only.

`spriter search --game <game>`
> Filters results by game name only.

`spriter search --game <game> --console <console>`
> Filters results by game and console.

## Compiling

`go build`

## License

This work is licensed under a
[Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License][cc-by-nc-sa].

[![CC BY-NC-SA 4.0][cc-by-nc-sa-image]][cc-by-nc-sa]

[cc-by-nc-sa]: http://creativecommons.org/licenses/by-nc-sa/4.0/
[cc-by-nc-sa-image]: https://licensebuttons.net/l/by-nc-sa/4.0/88x31.png
[cc-by-nc-sa-shield]: https://img.shields.io/badge/License-CC%20BY--NC--SA%204.0-lightgrey.svg