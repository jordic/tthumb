


# Thumbnail server in go

Simple and high performance thumbnail server written in go.

## Key points

- Simple (100LOC)
- No dependencies ( only a readable installation of Imagemagick )
- Uses imagemagick command to generate thumbnails
- Cached thumbnails from disk. Internally go uses file_serve to get maximum performance
- http server from go stdlib

## Operation

Create a presets.json wich each of the params you provide to imagemagic for creating an image

```bash
convert initial.jpg -colorspace gray -level +10% +level-colors "#000000","#ffec00" -thumbnail 300x200^ -extent 300x200 -quality 82 init_thumb.jpg
```
preset wil be:

```json
{
    "red300": "-colorspace gray -level +10% +level-colors '#000000','#ffec00' -thumbnail 350x250^ -extent 350x250 -quality 82"
}
```
Launch the server and you can access generated thumbnails from,

http://yourlocaladdress/pathtoimage/image?p=red300


## Docker

There is a Dockerfile provided to build a self contained image that can run everywhere :)

## Improvements

- Mount remote storage services from the app. (This way will also be allowed to be used as a cloud cache)






