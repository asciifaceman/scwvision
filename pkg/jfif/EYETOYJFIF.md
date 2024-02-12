# Eyetoy JFIF/JPG

As of this writing, relatively early in the project, when a request to read is made over USB the first 2 bytes received back appear to be as such:

```
ff ff ff 50 00 d7 1f 6a  
bf 00 03 00 00 00 10 4b
```

And a JPEG follows immediately after starting with the classic marker `FF D8`