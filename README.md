# Spriter

## Overview
This is a tool for bulk downloading assets from [The Spriters Resource](https://www.spriters-resource.com/).

This tool provides a scanning feature, which scans the list of available game assets on the website and caches that information so it can be used for searching without HTTP requests.

Scanning is only required if you want to use the search functionality. Searching is only really necessary if you don't know the URL of the game assets you are looking for.

## Scanning (Optional)

Scanning is required before using search functionality.
These commands scan the site and index information about what is available.

`scan consoles`
> Scans available consoles.

`scan games --meta` `scan -m`
> Fetches and caches game and page counts.

`scan games --all` `scan -a`
> **WARNING**: This command may take longer than 10 minutes to finish.
> 
> Scans all game pages and caches available game information.

`scan games --page <page-number>`
> Fetches and caches all games for a specific page number.

`scan games --pages <first-page-number> <last-page-number>`
> Fetches and caches all games for a specific range of pages (inclusive).

## Searching (Optional - requires Scanning)
Search the game cache by name or console. Requires scanning to work.

The search algorithm is very basic so it won't work flawlessly all of the time.

`search --console <console>`
> Filters results by console name only.

`search --game <game>`
> Filters results by game name only.

`search --game <game> --console <console>`
> Filters results by game and console.

### Downloading

`get <path>`
> Downloads all of the assets at `https://www.spriters-resource.com<path>`.
> 
> e.g. `/ds_dsi/dgmnworldds/` (must include leading slash) 
