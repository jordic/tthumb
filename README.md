


# Small thumbnailer

A wrap around imagemagick convert to generate ondisk thumbnails, 
for distinct profiles

convert initial.jpg -colorspace gray -level +10% +level-colors "#000000","#ffec00" -thumbnail 300x200^ -extent 300x200 -quality 82 init_thumb.jpg



