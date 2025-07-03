# easy diffusion bulk runner

I like to see what each of the paramaters do to a given prompt. This allows me to do that.

This is my 2nd version. The original was just crap!

The build file is network / machine dependent. I assume you can make your own.

```sh
#!/bin/sh

rm bin/auto-esd
if [ $? -ne 0 ]; then
    echo "Failed to remove old binary"
    exit 1
fi

go build -o bin/auto-esd
if [ $? -ne 0 ]; then
    echo "Build failed"
    exit 1
fi
echo "Build successful"

scp auto-esd user@machine:/home/bjones/easy-diffusion/
if [ $? -ne 0 ]; then
    echo "File transfer failed"
    exit 1
fi
echo "File transfer successful"
```