FROM gitpod/workspace-full-vnc

# Added OS packages
RUN apt-get update
# Per g3n readme https://github.com/g3n/engine#ubuntudebian-like
RUN apt-get install xorg-dev libgl1-mesa-dev libopenal1 libopenal-dev libvorbis0a libvorbis-dev libvorbisfile3 -y