# Rpi CCTV

A Raspberry Pi based CCTV camera with Pan & Tilt that can be controlled remotely

# Setup

As of now Rpi CCTV relies on `motion` for relaying rpi camera images. Ensure
you have motion installed and running on the raspberry pi.

    $ sudo apt-get install motion
    $ go generate && go build && sudo ./rpicctv

Then navigate to http://<ip-address-of-rpi>:3000/cctv to see the live feed.
You can use the arrow buttons or keyboard arrow keys to move the camera.
